package boosty

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"strconv"

	"boosty/internal/boosty/endpoint"
	"boosty/pkg/logger"
	"github.com/grafov/m3u8"
)

func (c *Client) GetPosts(ctx context.Context, limit int) (Posts, error) {
	values := url.Values{}
	values.Add("limit", strconv.Itoa(limit))

	body, err := c.sendRequest(ctx, endpoint.GetPosts, values)
	if err != nil {
		return nil, err
	}

	var data V1GetPostsResponse
	if err := json.NewDecoder(body).Decode(&data); err != nil {
		return nil, fmt.Errorf("parse body: %w", err)
	}

	var res = make([]*Post, 0, len(data.Data))
	for i := 0; i < len(data.Data); i++ {
		res = append(res, &data.Data[i])
	}

	return res, nil
}

func (c *Client) GetBlog(ctx context.Context) (*Blog, error) {
	var res V1GetBlogResponse

	body, err := c.sendRequest(ctx, endpoint.GetBlog, nil)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(body).Decode(&res); err != nil {
		return nil, fmt.Errorf("parse body: %w", err)
	}

	return &Blog{
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

	return bytes.NewReader(resp.Body()), nil
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
