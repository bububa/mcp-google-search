# mcp-google-search Google Search MCP server implementation in Golang

[![Go Reference](https://pkg.go.dev/badge/github.com/bububa/mcp-google-search.svg)](https://pkg.go.dev/github.com/bububa/mcp-google-search)
[![Go](https://github.com/bububa/mcp-google-search/actions/workflows/go.yml/badge.svg)](https://github.com/bububa/mcp-google-search/actions/workflows/go.yml)
[![goreleaser](https://github.com/bububa/mcp-google-search/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/bububa/mcp-google-search/actions/workflows/goreleaser.yml)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/bububa/mcp-google-search.svg)](https://github.com/bububa/mcp-google-search)
[![GoReportCard](https://goreportcard.com/badge/github.com/bububa/mcp-google-search)](https://goreportcard.com/report/github.com/bububa/mcp-google-search)
[![GitHub license](https://img.shields.io/github/license/bububa/mcp-google-search.svg)](https://github.com/bububa/mcp-google-search/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/bububa/mcp-google-search.svg)](https://gitHub.com/bububa/mcp-google-search/releases/)

## Installation

```bash
go install github.com/bububa/mcp-google-search
```

## Setup

Set environment variable

- GOOGLE_SEARCH_CX
  Google custom search ID
- GOOGLE_SEARCH_KEY
  Google custom search API key

## Usage

### Start stdio server

```bash
mcp-google-search
```

### Start streamable server

```bash
mcp-google-search -port=8080
```

## Schema

### Webpage Search

- Function Name:
  search
- Input:

```json
{
  "query": "search query",
  "num": "max number of search results"
}
```

- Output:

```json
[
  {
    "title": "webpage title",
    "snippet": "webpage snippet",
    "url": "webpage link"
  }
]
```

### Image Search

- Function Name:
  image_search
- Input:

```json
{
  "query": "search query",
  "num": "max number of search results"
}
```

- Output:

```json
[
  {
    "title": "image title",
    "url": "image link",
    "mime": "image mimetype",
    "context_link": "image context_link",
    "width": "image width",
    "height": "image height",
    "byte_size": "image_bytesize"
  }
]
```
