package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
)

func SearchImageTool(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return searchTool(ctx, req, "image")
}
