package commands

const help_text = `
PowerDNS commands:

- list
- show [zone-name]
- touch [zone-name] (updates SOA record)
- create [zone-name]
- drop [zone-name]
- add_record [zone-name]
- update_record [zone-name]
- delete_record [zone-name]

`

func help(args ...string) (string, error) {
	return help_text, nil
}
