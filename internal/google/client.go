package google

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type Client struct {
	cx         string
	key        string
	httpClient *http.Client
}

func NewClient(cx string, key string) *Client {
	httpClient := http.DefaultClient
	if proxyURL := os.Getenv("https_proxy"); proxyURL != "" {
		proxy, _ := url.Parse(proxyURL)
		httpClient.Transport = &http.Transport{
			Proxy:           http.ProxyURL(proxy),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		fmt.Println("USE https_proxy", proxyURL)
	}
	return &Client{
		cx:         cx,
		key:        key,
		httpClient: httpClient,
	}
}

func (c *Client) SetHTTPClient(clt *http.Client) {
	c.httpClient = clt
}

func (c *Client) Search(ctx context.Context, req *Request, resp *Response) error {
	values := make(url.Values)
	values.Set("cx", c.cx)
	values.Set("key", c.key)
	values.Set("q", req.Query)
	if req.SearchType != "" {
		values.Set("searchType", req.SearchType)
	}
	var buf bytes.Buffer
	buf.WriteString(GATEWAY)
	buf.WriteByte('?')
	buf.WriteString(values.Encode())
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, buf.String(), nil)
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(httpResp.Body)
		return fmt.Errorf("google search API returned an error: %s, \n%s", httpResp.Status, string(body))
	}
	if err := json.NewDecoder(httpResp.Body).Decode(resp); err != nil {
		return err
	}
	return nil
}
