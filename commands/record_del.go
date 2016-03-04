package commands

import (
	"errors"
	"strconv"

	"github.com/epicagency/pdns-manager/pdns"
)

func record_del(args ...string) (string, error) {
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
		if rec != record && rec.Name == record.Name && rec.Type == record.Type {
			records = append(records, rec)
		}
	}

	shell.Print("Do you really want to delete this record?? [y/n] ")
	confirm := shell.ReadLine()
	if confirm != "y" && confirm != "Y" {
		return "", nil
	}

	if len(records) > 1 {
		errs = zone.UpdateRecords(records)
	} else {
		errs = zone.DeleteRecord(record)
	}
	if errs != nil {
		for err := range errs {
			shell.Println(err)
		}
	}
	return "", nil
}
