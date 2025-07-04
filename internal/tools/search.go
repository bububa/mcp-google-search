package tools

import (
	"context"
	"encoding/json"
	"errors"
	"os"

	"github.com/mark3labs/mcp-go/mcp"

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
	resp := new(google.Response)
	clt := google.NewClient(os.Getenv("GOOGLE_SEARCH_CS"), os.Getenv("GOOGLE_SEARCH_KEY"))
	if err := clt.Search(ctx, &google.Request{
		Query:      query,
		SearchType: searchType,
	}, resp); err != nil {
		return nil, err
	}
	bs, _ := json.Marshal(resp.Items)
	return mcp.NewToolResultText(string(bs)), nil
}
