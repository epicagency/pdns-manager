package commands

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/epicagency/pdns-manager/pdns"
	"github.com/parnurzeal/gorequest"
)

func touch(args ...string) (string, error) {
	shell.ShowPrompt(false)
	defer shell.ShowPrompt(true)

	_, bytes, errs := gorequest.
		New().
		Get(fmt.Sprintf("http://dns1.epic-sys.io/servers/localhost/zones/%s", args[0])).
		Set("X-API-Key", "bisque.tutelage.organist.payment").
		EndBytes()

	if errs != nil {
		for err := range errs {
			shell.Println(err)
		}
	}

	zone := new(pdns.Zone)
	err := json.Unmarshal(bytes, zone)
	if err != nil {
		return "", err
	}
	var soa_record *pdns.Record
	for _, record := range zone.Records {
		if record.Type == "SOA" {
			soa_record = record
			break
		}
	}
	if soa_record == nil {
		shell.Println("No SOA record found!")
		return "", nil
	}
	shell.Println(fmt.Sprintf("Current serial: %d", zone.Serial))
	for {
		shell.Print("New serial: ")
		serial := shell.ReadLine()
		if serial == "" {
			shell.Println("No serial provided, canceling")
			return "", nil
		}
		iserial, err := strconv.Atoi(serial)
		if err != nil {
			shell.Println("Not a valid serial")
			continue
		}
		if iserial < zone.Serial {
			shell.Print("Serial is below current, are you sure? [y/n]")
			confirm := shell.ReadLine()
			if confirm == "n" || confirm == "N" {
				continue
			}
		}
		shell.Println("Updating serial...")
		return "", nil
	}
}
