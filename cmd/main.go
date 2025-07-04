package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/bububa/mcp-google-search/internal/tools"
)

func main() {
	// Create MCP server
	s := server.NewMCPServer(
		"mcp-google-search",
		"1.0.0",
		server.WithLogging(),
		server.WithToolCapabilities(false),
	)

	// Register search tools
	searchTool := mcp.NewTool("search",
		mcp.WithDescription("search web pages through google"),
		mcp.WithString("query",
			mcp.Description("The search terms to query."),
			mcp.Required(),
		),
	)

	// Register imageSearch tools
	imageSearchTool := mcp.NewTool("image_search",
		mcp.WithDescription("search images through google"),
		mcp.WithString("query",
			mcp.Description("The search terms to query."),
			mcp.Required(),
		),
	)

	// Add tool handler
	s.AddTool(searchTool, tools.SearchTool)
	s.AddTool(imageSearchTool, tools.SearchImageTool)

	var (
		host string
		port int
	)
	flag.StringVar(&host, "host", "", "sse server host")
	flag.IntVar(&port, "port", 0, "sse server port")
	flag.Parse()
	if host == "" || port == 0 {
		// Start server via stdio
		if err := server.ServeStdio(s); err != nil {
			os.Exit(1)
		}
		return
	}
	startCtx := context.Background()
	srv := server.NewSSEServer(s)
	stopCtx, stop := signal.NotifyContext(startCtx, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := srv.Start(fmt.Sprintf("%s:%d", host, port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln(err)
		}
	}()
	<-stopCtx.Done()
	stop()
	ctx, cancel := context.WithTimeout(startCtx, 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}
}
