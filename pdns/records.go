package pdns

type Record struct {
	TTL      int    `json:"ttl"`
	Type     string `json:"type"`
	Content  string `json:"content"`
	Name     string `json:"name"`
	Priority int    `json:"priority"`
	Disabled bool   `json:"disabled"`
}

type RRset struct {
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	ChangeType string    `json:"changetype"`
	Records    []*Record `json:"records"`
}

type RRsetContainer struct {
	RRsets []*RRset `json:"rrsets"`
}
