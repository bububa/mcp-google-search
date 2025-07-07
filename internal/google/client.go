package google

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	cx         string
	key        string
	httpClient *http.Client
}

func NewClient(cx string, key string) *Client {
	return &Client{
		cx:         cx,
		key:        key,
		httpClient: http.DefaultClient,
	}
}

func (c *Client) SetHTTPClient(clt *http.Client) {
	c.httpClient = clt
}

func (c *Client) Search(ctx context.Context, req *Request, resp *Response) error {
	values := make(url.Values)
	values.Set("cx", c.cx)
	values.Set("key", c.key)
	values.Set("query", req.Query)
	if req.SearchType != "" {
		values.Set("searchType", req.SearchType)
	}
	var buf bytes.Buffer
	buf.WriteString(GATEWAY)
	buf.WriteByte('?')
	buf.WriteString(values.Encode())
	fmt.Println(buf.String())
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, buf.String(), nil)
	if err != nil {
		return err
	}
	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(httpResp.Body)
		return fmt.Errorf("google search API returned an error: %s, \n%s", httpResp.Status, string(body))
	}
	buf.Reset()
	if err := json.NewDecoder(&buf).Decode(resp); err != nil {
		return err
	}
	return nil
}
