package entity

type Webpage struct {
	Title   string `json:"title,omitempty"`
	URL     string `json:"url,omitempty"`
	Snippet string `json:"snippet,omitempty"`
}

type Image struct {
	Title       string `json:"title,omitempty"`
	URL         string `json:"url,omitempty"`
	ContextLink string `json:"context_link,omitempty"`
	Mime        string `json:"mime,omitempty"`
	Width       int    `json:"width,omitempty"`
	Height      int    `json:"height,omitempty"`
	ByteSize    int64  `json:"byte_size,omitempty"`
}
