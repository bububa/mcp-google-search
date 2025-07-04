package main

import (
	"os"
	"strconv"

	"github.com/bububa/mcp-google-search/internal"
)

func main() {
	port, _ := strconv.Atoi(os.Getenv("GOOGLE_SEARCH_PORT"))
	internal.StartServer(port)
}
