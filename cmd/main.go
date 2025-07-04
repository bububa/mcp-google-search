package cmd

import (
	"flag"

	"github.com/bububa/mcp-google-search/internal"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 0, "sse server port")
	flag.Parse()
	internal.StartServer(port)
}
