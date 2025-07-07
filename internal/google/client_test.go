package google

import (
	"context"
	"encoding/json"
	"os"
	"testing"
)

func TestSearch(t *testing.T) {
	cx := os.Getenv("GOOGLE_SEARCH_CX")
	key := os.Getenv("GOOGLE_SEARCH_KEY")
	query := "rabbit"
	clt := NewClient(cx, key)
	ret := new(Response)
	if err := clt.Search(context.Background(), &Request{
		Query: query,
	}, ret); err != nil {
		t.Error(err)
		return
	}
	{
		bs, _ := json.MarshalIndent(ret, "", "  ")
		t.Log(string(bs))
	}
}

func TestImageSearch(t *testing.T) {
	cx := os.Getenv("GOOGLE_SEARCH_CX")
	key := os.Getenv("GOOGLE_SEARCH_KEY")
	query := "rabbit"
	clt := NewClient(cx, key)
	ret := new(Response)
	if err := clt.Search(context.Background(), &Request{
		Query:      query,
		SearchType: "image",
	}, ret); err != nil {
		t.Error(err)
		return
	}
	{
		bs, _ := json.MarshalIndent(ret, "", "  ")
		t.Log(string(bs))
	}
}
