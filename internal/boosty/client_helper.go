package boosty

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"

	"boosty/internal/boosty/endpoint"
	"boosty/internal/boosty/model"
	"boosty/internal/logger"
	"github.com/grafov/m3u8"
)

var (
	ErrUserUnauthorized = errors.New("unauthorized or token has been expired")
	ErrInternalServer   = errors.New("internal server error")
	ErrPostNotFound     = errors.New("post not found")
)

func (c *Client) GetPostIterator(a Args) func(ctx context.Context) (model.Posts, error) {
	args := a
	done := false
	return func(ctx context.Context) (model.Posts, error) {
		if done {
			return nil, ErrPostNotFound
		}

		body, err := c.sendRequest(ctx, endpoint.GetPosts, args.QueryParams())
		if err != nil {
			return nil, err
		}

		var data model.V1GetPostsResponse
		if err := json.NewDecoder(body).Decode(&data); err != nil {
			return nil, fmt.Errorf("parse body: %w", err)
		}

		args.Offset = data.Extra.Offset
		done = data.Extra.IsLast

		var res = make([]*model.Post, 0, len(data.Data))
		for i := 0; i < len(data.Data); i++ {
			res = append(res, &data.Data[i])
		}

		return res, nil
	}
}

func (c *Client) GetPosts(ctx context.Context, args Args) (model.Posts, error) {
	body, err := c.sendRequest(ctx, endpoint.GetPosts, args.QueryParams())
	if err != nil {
		return nil, err
	}

	var data model.V1GetPostsResponse
	if err := json.NewDecoder(body).Decode(&data); err != nil {
		return nil, fmt.Errorf("parse body: %w", err)
	}

	var res = make([]*model.Post, 0, len(data.Data))
	for i := 0; i < len(data.Data); i++ {
		res = append(res, &data.Data[i])
	}

	return res, nil
}

func (c *Client) GetBlog(ctx context.Context) (*model.Blog, error) {
	var res model.V1GetBlogResponse

	body, err := c.sendRequest(ctx, endpoint.GetBlog, nil)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(body).Decode(&res); err != nil {
		return nil, fmt.Errorf("parse body: %w", err)
	}

	return &model.Blog{
		Title:        res.Title,
		URL:          res.BlogUrl,
		Stats:        res.Count,
		AccessRights: res.AccessRights,
		IsSubscribed: res.IsSubscribed,
	}, nil
}

func (c *Client) sendRequest(ctx context.Context, e endpoint.Endpoint, values url.Values) (io.Reader, error) {
	requestUrl := c.getEndpoint(e)

	logger.Debug(ctx, "sendRequest", "endpoint", requestUrl)

	req := c.http.R().SetQueryParamsFromValues(values)
	if c.debug {
		req.EnableTrace()
	}
	if c.token != "" {
		req.SetAuthToken(c.token)
	}

	resp, err := req.Get(requestUrl)
	if c.debug {
		ti := resp.Request.TraceInfo()
		logger.Debug(ctx, "sendRequest",
			"code", resp.StatusCode(),
			"status", resp.Status(),
			"proto", resp.Proto(),
			"time", resp.Time(),
			"receivedAt", resp.ReceivedAt(),
			"requestAttempt", ti.RequestAttempt,
			"remoteAddr", ti.RemoteAddr.String(),
		)
	}
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	switch resp.StatusCode() {
	case 200:
		return bytes.NewReader(resp.Body()), nil
	case 401:
		return nil, ErrUserUnauthorized
	case 400, 403, 500:
		return nil, ErrInternalServer
	default:
		return nil, fmt.Errorf("do request: %v", resp.Status())
	}
}

func (c *Client) GetM3u8MasterPlaylist(url string) (*m3u8.MasterPlaylist, error) {
	resp, err := c.http.R().Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, errors.New(resp.Status())
	}
	p, t, err := m3u8.Decode(*bytes.NewBuffer(resp.Body()), false)
	if err != nil {
		return nil, err
	}
	if t != m3u8.MASTER {
		return nil, errors.New("not master playlist")
	}
	return p.(*m3u8.MasterPlaylist), err
}
