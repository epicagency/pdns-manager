package pdns

type ZoneItem struct {
	DNSSEC         bool   `json:"dnssec"`
	Id             string `json:"id"`
	Kind           string `json:"kind"`
	LastCheck      int    `json:"last_check"`
	Name           string `json:"name"`
	NotifiedSerial int    `json:"notified_serial"`
	Serial         int    `json:"serial"`
	URL            string `json:"url"`
	//Masters: [],
}

type Zone struct {
	Serial         int       `json:"serial"`
	SOAEditApi     string    `json:"soa_edit_api"`
	SOAEdit        string    `json:"soa_edit"`
	URL            string    `json:"url"`
	Records        []*Record `json:"records"`
	Type           string    `json:"type"`
	Comments       []string  `json:"comments"`
	Name           string    `json:"name"`
	NotifiedSerial int       `json:"notified_serial"`
	Masters        []string  `json:"masters"`
	Nameservers    []string  `json:"nameservers"`
	Kind           string    `json:"kind"`
	LastCheck      int       `json:"last_check"`
	DNSSEC         bool      `json:"dnssec"`
}
