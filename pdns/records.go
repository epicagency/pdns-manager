package pdns

type Record struct {
	TTL      int    `json:"ttl"`
	Type     string `json:"type"`
	Content  string `json:"content"`
	Name     string `json:"name"`
	Priority int    `json:"priority"`
	Disabled bool   `json:"disabled"`
}
