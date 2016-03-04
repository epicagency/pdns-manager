package commands

import (
	"strconv"

	"github.com/epicagency/pdns-manager/pdns"
)

func record_add(args ...string) (string, error) {
	shell.ShowPrompt(false)
	defer shell.ShowPrompt(true)

	zone, errs := pdns.GetZone(args[0])
	if errs != nil {
		for err := range errs {
			shell.Println(err)
		}
		return "", nil
	}

	record := new(pdns.Record)
	record.TTL = 300
	record.Priority = 0
	record.Disabled = false

	if len(args) < 4 {
		shell.Print("Name: ")
		name := shell.ReadLine()
		if name == "" {
			return "name can't be empty", nil
		}
		record.Name = name

		shell.Print("Type: ")
		typ := shell.ReadLine()
		if typ == "" {
			return "type can't be empty", nil
		}
		record.Type = typ

		shell.Print("Content: ")
		content := shell.ReadLine()
		if content == "" {
			return "content can't be empty", nil
		}
		record.Content = content

		shell.Print("TTL [300]: ")
		ttl := shell.ReadLine()
		if ttl != "" {
			record.TTL, _ = strconv.Atoi(ttl)
		}

		shell.Print("Priority [0]: ")
		prio := shell.ReadLine()
		if prio != "" {
			record.Priority, _ = strconv.Atoi(prio)
		}

		shell.Print("Disabled [false]?: ")
		dis := shell.ReadLine()
		record.Disabled = (dis == "y")
	} else {
		record.Name = args[1]
		record.Type = args[2]
		record.Content = args[3]
		if len(args) > 4 {
			record.TTL, _ = strconv.Atoi(args[4])
		}
		if len(args) > 5 {
			record.Priority, _ = strconv.Atoi(args[5])
		}
		if len(args) > 6 {
			record.Disabled = (args[6] == "y")
		}
	}

	shell.Print("Do you really want to add this record?? [y/n] ")
	confirm := shell.ReadLine()
	if confirm != "y" && confirm != "Y" {
		return "", nil
	}

	records := make([]*pdns.Record, 0, 5)
	for _, rec := range zone.Records {
		if rec.Name == record.Name && rec.Type == record.Type {
			records = append(records, rec)
		}
	}
	records = append(records, record)

	errs = zone.UpdateRecords(records)
	if errs != nil {
		for err := range errs {
			shell.Println(err)
		}
	}
	return "", nil
}
