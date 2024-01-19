package api

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"go.uber.org/zap"
)

var ErrFailedResponse = errors.New("failed to get response")

func (c *Client) GetAge(ctx context.Context, name string) (*AgeResponse, error) {
	base, err := url.Parse(c.cfg.AgeURL)
	if err != nil {
		c.log.Error("failed to parse url", zap.String("url", c.cfg.AgeURL), zap.Error(err))
		return nil, err
	}

	queryParams := url.Values{}
	queryParams.Add("name", name)

	base.RawQuery = queryParams.Encode()

	c.log.Info("created request", zap.String("url", base.String()))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, base.String(), nil)
	if err != nil {
		c.log.Error("failed to create new request", zap.Error(err))
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		c.log.Error("failed to make request", zap.Error(err))
		return nil, err

	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		errBody, _ := io.ReadAll(resp.Body)
		c.log.Error("failed to get response", zap.Int("http code", resp.StatusCode), zap.ByteString("err body", errBody))
		return nil, ErrFailedResponse
	}

	var dto AgeResponse
	err = json.NewDecoder(resp.Body).Decode(&dto)
	if err != nil {
		c.log.Error("failed to decode response", zap.Error(err))
		return nil, err
	}

	c.log.Debug("decoded age dto", zap.Any("response", dto))

	return &dto, nil
}

func (c *Client) GetGender(ctx context.Context, name string) (*GenderResponse, error) {
	base, err := url.Parse(c.cfg.GenderURL)
	if err != nil {
		c.log.Error("failed to parse url", zap.String("url", c.cfg.AgeURL), zap.Error(err))
		return nil, err
	}

	queryParams := url.Values{}
	queryParams.Add("name", name)

	base.RawQuery = queryParams.Encode()

	c.log.Info("created request", zap.String("url", base.String()))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, base.String(), nil)
	if err != nil {
		c.log.Error("failed to create new request", zap.Error(err))
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		c.log.Error("failed to make request", zap.Error(err))
		return nil, err

	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		errBody, _ := io.ReadAll(resp.Body)
		c.log.Error("failed to get response", zap.Int("http code", resp.StatusCode), zap.ByteString("err body", errBody))
		return nil, ErrFailedResponse
	}

	var dto GenderResponse
	err = json.NewDecoder(resp.Body).Decode(&dto)
	if err != nil {
		c.log.Error("failed to decode response", zap.Error(err))
		return nil, err
	}

	c.log.Debug("decoded gender dto", zap.Any("response", dto))

	return &dto, nil
}

func (c *Client) GetNationality(ctx context.Context, name string) (*NationalityResponse, error) {
	base, err := url.Parse(c.cfg.NationalityURL)
	if err != nil {
		c.log.Error("failed to parse url", zap.String("url", c.cfg.AgeURL), zap.Error(err))
		return nil, err
	}

	queryParams := url.Values{}
	queryParams.Add("name", name)

	base.RawQuery = queryParams.Encode()

	c.log.Info("created request", zap.String("url", base.String()))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, base.String(), nil)
	if err != nil {
		c.log.Error("failed to create new request", zap.Error(err))
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		c.log.Error("failed to make request", zap.Error(err))
		return nil, err

	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		errBody, _ := io.ReadAll(resp.Body)
		c.log.Error("failed to get response", zap.Int("http code", resp.StatusCode), zap.ByteString("err body", errBody))
		return nil, ErrFailedResponse
	}

	var dto NationalityResponse
	err = json.NewDecoder(resp.Body).Decode(&dto)
	if err != nil {
		c.log.Error("failed to decode response", zap.Error(err))
		return nil, err
	}

	c.log.Debug("decoded gender dto", zap.Any("response", dto))

	return &dto, nil
}
