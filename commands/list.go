package commands

import (
	"encoding/json"
	"fmt"

	"github.com/epicagency/pdns-manager/pdns"
	"github.com/parnurzeal/gorequest"
)

func list(args ...string) (string, error) {

	_, bytes, errs := gorequest.
		New().
		Get("http://dns1.epic-sys.io/servers/localhost/zones").
		Set("X-API-Key", "bisque.tutelage.organist.payment").
		EndBytes()

	if errs != nil {
		for err := range errs {
			shell.Println(err)
		}
	}

	zones := make([]pdns.ZoneItem, 0)
	err := json.Unmarshal(bytes, &zones)
	if err != nil {
		return "", err
	}

	for _, zone := range zones {
		shell.Println(fmt.Sprintf("%s: %s (%d)", zone.Id, zone.Kind, zone.NotifiedSerial))
	}

	return "", nil
}
