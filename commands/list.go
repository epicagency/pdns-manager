package commands

import (
	"fmt"

	"github.com/epicagency/pdns-manager/pdns"
)

func list(args ...string) (string, error) {
	zones, errs := pdns.GetZones()
	if errs != nil {
		for err := range errs {
			shell.Println(err)
		}
	}

	for _, zone := range zones {
		shell.Println(fmt.Sprintf("%s: %s (%d)", zone.Id, zone.Kind, zone.NotifiedSerial))
	}

	return "", nil
}
