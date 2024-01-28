package boosty

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strconv"

	"boosty/internal/clients/boosty/endpoint"
	"boosty/internal/clients/boosty/models"
	"boosty/pkg/logger"
)

func (c *Client) GetPosts(ctx context.Context, limit int) (models.Posts, error) {
	values := url.Values{}
	values.Add("limit", strconv.Itoa(limit))

	body, err := c.sendRequest(ctx, endpoint.GetPosts, values)
	if err != nil {
		return nil, err
	}

	var data models.V1GetPostsResponse
	if err := json.NewDecoder(body).Decode(&data); err != nil {
		return nil, fmt.Errorf("parse body: %w", err)
	}

	var res = make([]*models.Post, 0, len(data.Data))
	for i := 0; i < len(data.Data); i++ {
		res = append(res, &data.Data[i])
	}
	return res, nil
}

func (c *Client) GetBlog(ctx context.Context) (*models.Blog, error) {
	var res models.V1GetBlogResponse

	body, err := c.sendRequest(ctx, endpoint.GetBlog, nil)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(body).Decode(&res); err != nil {
		return nil, fmt.Errorf("parse body: %w", err)
	}

	return &models.Blog{
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
