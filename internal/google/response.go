package google

type Response struct {
	Items []Item `json:"items,omitempty"`
}

type Item struct {
	Title       string `json:"title,omitempty"`
	Link        string `json:"link,omitempty"`
	DisplayLink string `json:"displayLink,omitempty"`
	Snippet     string `json:"snippet,omitempty"`
	Mime        string `json:"mime,omitempty"`
	Image       *Image `json:"image,omitempty"`
}

type Image struct {
	ContextLink         string `json:"contextLink,omitempty"`
	Height              int    `json:"height,omitempty"`
	Width               int    `json:"width,omitempty"`
	ByteSize            int64  `json:"byteSize,omitempty"`
	ThumbnailLink       string `json:"thumbnailLink,omitempty"`
	ThumbnailLinkHeight int    `json:"thumbnailLinkHeight,omitempty"`
	ThumbnailLinkWidth  int    `json:"thumbnailLinkWidth,omitempty"`
}
