package commands

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/epicagency/pdns-manager/pdns"
)

func zone_touch(args ...string) (string, error) {
	shell.ShowPrompt(false)
	defer shell.ShowPrompt(true)

	zone, errs := pdns.GetZone(args[0])
	if errs != nil {
		for err := range errs {
			shell.Println(err)
		}
		return "", nil
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

	var serial string
	for {
		shell.Print("New serial: ")
		serial = shell.ReadLine()
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

		break
	}

	soa_record.Content = strings.Replace(soa_record.Content, strconv.Itoa(zone.Serial), serial, 1)

	shell.Println(fmt.Sprintf("New SOA: %s", soa_record.Content))

	errs = zone.UpdateRecords([]*pdns.Record{soa_record})
	if errs != nil {
		for err := range errs {
			shell.Println(err)
		}
	}
	return "", nil
}
