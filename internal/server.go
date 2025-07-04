package internal

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/bububa/mcp-google-search/internal/tools"
)

const (
	AppName    = "mcp-google-search"
	AppVersion = "1.0.0"
)

func StartServer(port int) {
	// Create MCP server
	s := server.NewMCPServer(
		AppName,
		AppVersion,
		server.WithLogging(),
		server.WithToolCapabilities(true),
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

	if port == 0 {
		// Start server via stdio
		if err := server.ServeStdio(s); err != nil {
			os.Exit(1)
		}
		return
	}
	// srv := server.NewSSEServer(s)
	// if err := srv.Start(fmt.Sprintf(":%d", port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
	// 	log.Fatalln(err)
	// }
	//
	mux := http.NewServeMux()
	registerHealthAndVersion(mux)
	srv := server.NewStreamableHTTPServer(s)
	mux.Handle("/", srv)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalln(err)
	}

	// if err := srv.Start(fmt.Sprintf(":%d", port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
	// 	log.Fatalln(err)
	// }
}

// registerHealthAndVersion adds the /health and /version endpoints.
func registerHealthAndVersion(mux *http.ServeMux) {
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// Matching the documented output from the project's README.
		fmt.Fprint(w, `{"status":"ok"}`)
	})

	mux.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"name":"%s","version":"%s"}`, AppName, AppVersion)
	})
}
