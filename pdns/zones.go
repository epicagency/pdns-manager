package pdns

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/parnurzeal/gorequest"
)

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

func GetZone(name string) (*Zone, []error) {
	_, bytes, errs := gorequest.
		New().
		Get(fmt.Sprintf("%s/zones/%s", os.Getenv("API_URL"), name)).
		Set("X-API-Key", os.Getenv("API_KEY")).
		EndBytes()

	if errs != nil {
		return nil, errs
	}

	zone := new(Zone)
	err := json.Unmarshal(bytes, zone)
	if err != nil {
		return nil, []error{err}
	}
	return zone, nil
}

func GetZones() ([]ZoneItem, []error) {
	_, bytes, errs := gorequest.
		New().
		Get(fmt.Sprintf("%s/zones", os.Getenv("API_URL"))).
		Set("X-API-Key", os.Getenv("API_KEY")).
		EndBytes()

	if errs != nil {
		return nil, errs
	}
	zones := make([]ZoneItem, 0)
	err := json.Unmarshal(bytes, &zones)
	if err != nil {
		return nil, []error{err}
	}
	return zones, nil
}

func (z *Zone) UpdateRecords(rs []*Record) []error {
	c := RRsetContainer{
		RRsets: []*RRset{
			&RRset{
				Name: rs[0].Name, Type: rs[0].Type, ChangeType: "REPLACE",
				Records: rs,
			},
		},
	}

	_, bytes, errs := gorequest.
		New().
		Patch(fmt.Sprintf("%s/zones/%s", os.Getenv("API_URL"), z.Name)).
		Set("X-API-Key", os.Getenv("API_KEY")).
		Send(c).
		EndBytes()

	if errs != nil {
		return errs
	}
	err := json.Unmarshal(bytes, z)
	if err != nil {
		return []error{err}
	}

	return nil
}

func (z *Zone) DeleteRecord(r *Record) []error {
	c := RRsetContainer{
		RRsets: []*RRset{
			&RRset{
				Name: r.Name, Type: r.Type, ChangeType: "DELETE",
				Records: []*Record{r},
			},
		},
	}

	_, bytes, errs := gorequest.
		New().
		Patch(fmt.Sprintf("%s/zones/%s", os.Getenv("API_URL"), z.Name)).
		Set("X-API-Key", os.Getenv("API_KEY")).
		Send(c).
		EndBytes()

	if errs != nil {
		return errs
	}
	err := json.Unmarshal(bytes, z)
	if err != nil {
		return []error{err}
	}

	return nil
}
