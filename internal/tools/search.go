package tools

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"strconv"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/bububa/mcp-google-search/entity"
	"github.com/bububa/mcp-google-search/internal/google"
)

func SearchTool(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return searchTool(ctx, req, "")
}

func searchTool(ctx context.Context, req mcp.CallToolRequest, searchType string) (*mcp.CallToolResult, error) {
	var query string
	mp := req.GetArguments()
	if v, ok := mp["query"]; ok {
		if str, ok := v.(string); ok {
			query = str
		}
	}
	if query == "" {
		return nil, errors.New("missing query parameter")
	}
	searchReq := &google.Request{
		Query:      query,
		SearchType: searchType,
	}
	if v, ok := mp["num"]; ok {
		if str, ok := v.(string); ok {
			searchReq.Num, _ = strconv.Atoi(str)
		}
	}
	resp := new(google.Response)
	clt := google.NewClient(os.Getenv("GOOGLE_SEARCH_CX"), os.Getenv("GOOGLE_SEARCH_KEY"))
	if err := clt.Search(ctx, searchReq, resp); err != nil {
		return nil, err
	}
	list := make([]any, 0, len(resp.Items))
	for _, v := range resp.Items {
		if searchType == "image" {
			var img entity.Image
			convertImage(&img, &v)
			list = append(list, img)
		} else {
			var page entity.Webpage
			convertWebpage(&page, &v)
			list = append(list, page)
		}
	}
	bs, _ := json.Marshal(list)
	return mcp.NewToolResultText(string(bs)), nil
}

func convertWebpage(dist *entity.Webpage, src *google.Item) {
	dist.Title = src.Title
	dist.Snippet = src.Snippet
	dist.URL = src.Link
}

func convertImage(dist *entity.Image, src *google.Item) {
	dist.Title = src.Title
	dist.URL = src.Link
	dist.Mime = src.Mime
	if img := src.Image; img != nil {
		dist.ContextLink = img.ContextLink
		dist.Width = img.Width
		dist.Height = img.Height
		dist.ByteSize = img.ByteSize
	}
}
