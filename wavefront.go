package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

// Service for wavefront
type Service interface {
	GetAWSMetricsList(ctx context.Context, data interface{}, metric string) error
}

// Client for wavefront
type Client struct {
	token string
}

var (
	baseURL = "https://okta.wavefront.com/chart/metrics/all?&l=1000"
)

var _ Service = (*Client)(nil)

// New wavefront client
func New() (*Client, error) {
	token := viper.GetString("wavefront.token")
	if token == "" {
		return nil, errors.New("wavefront.token is needed, please add this to your $HOME/.okta/config.yaml or export CKP_WAVEFRONT_TOKEN")
	}
	return &Client{token: token}, nil
}

// GetAWSMetricsList gets all the aws sub root metrics under aws root metrics namespace
func (c *Client) GetAWSMetricsList(ctx context.Context, data interface{}, metric string) error {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf(baseURL+"&q=%s", metric), nil)
	if err != nil {
		return err
	}

	return c.sendRequest(ctx, req, data)
}

func (c *Client) sendRequest(ctx context.Context, req *http.Request, data interface{}) error {
	ctx, cancelFunc := context.WithTimeout(ctx, 30*time.Second)
	defer cancelFunc()

	err := ensureAuth(req, c.token)
	if err != nil {
		return err
	}

	err = getRequest(ctx, req, data)
	if errors.Is(err, context.DeadlineExceeded) {
		return fmt.Errorf("request time out for %s", err.Error())
	}

	if err != nil {
		return err
	}

	return nil
}

func getRequest(ctx context.Context, req *http.Request, data interface{}) error {
	client := &http.Client{}

	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		if req.Method != http.MethodDelete {
			return json.NewDecoder(resp.Body).Decode(data)
		}
		return nil
	}

	if resp.StatusCode == http.StatusNotFound && req.Method == http.MethodDelete {
		return nil
	}

	var respBody bytes.Buffer
	if resp.Body != nil {
		io.Copy(&respBody, resp.Body)
	}

	return fmt.Errorf("error response: %d, %s", resp.StatusCode, respBody.String())
}

func ensureAuth(req *http.Request, token string) error {
	if req.Header == nil {
		req.Header = map[string][]string{}
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("content-type", "application/json")
	return nil
}
