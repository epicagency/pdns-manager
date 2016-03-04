package commands

const help_text = `
PowerDNS commands:

- zlist
- zshow [zone-name]
- ztouch [zone-name] (updates SOA record)
- zcreate [zone-name]
- drop [zone-name]
- radd [zone-name] {name} {type} {content} {ttl} {priority} {disabled}
- rdel [zone-name] [index]
- rupd [zone-name] [index]

`

func help(args ...string) (string, error) {
	return help_text, nil
}
