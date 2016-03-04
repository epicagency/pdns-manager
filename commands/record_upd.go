package commands

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/epicagency/pdns-manager/pdns"
)

func record_upd(args ...string) (string, error) {
	shell.ShowPrompt(false)
	defer shell.ShowPrompt(true)

	zone, errs := pdns.GetZone(args[0])
	if errs != nil {
		for err := range errs {
			shell.Println(err)
		}
		return "", nil
	}
	record_id, err := strconv.Atoi(args[1])
	if err != nil {
		return "", err
	}
	if record_id < 0 || record_id >= len(zone.Records) {
		return "", errors.New("Index out of bounds")
	}

	record := zone.Records[record_id]
	records := make([]*pdns.Record, 0, 5)
	for _, rec := range zone.Records {
		if rec.Name == record.Name && rec.Type == record.Type {
			records = append(records, rec)
		}
	}

	shell.Println(fmt.Sprintf("Name: %s", record.Name))
	shell.Println(fmt.Sprintf("Type: %s", record.Type))
	shell.Print(fmt.Sprintf("New content [%s]: ", record.Content))
	content := shell.ReadLine()
	if content != "" {
		record.Content = content
	}
	shell.Print(fmt.Sprintf("New TTL [%d]: ", record.TTL))
	ttl := shell.ReadLine()
	if ttl != "" {
		record.TTL, _ = strconv.Atoi(ttl)
	}
	shell.Print(fmt.Sprintf("New priority [%d]: ", record.Priority))
	prio := shell.ReadLine()
	if prio != "" {
		record.Priority, _ = strconv.Atoi(prio)
	}
	shell.Print(fmt.Sprintf("Disabled [%t]?: ", record.Disabled))
	dis := shell.ReadLine()
	if dis != "" {
		record.Disabled = (dis == "y")
	}
	shell.Print("Do you really want to update this record?? [y/n] ")
	confirm := shell.ReadLine()
	if confirm != "y" && confirm != "Y" {
		return "", nil
	}

	errs = zone.UpdateRecords(records)
	if errs != nil {
		for err := range errs {
			shell.Println(err)
		}
	}
	return "", nil
}
