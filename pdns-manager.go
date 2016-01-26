package main

import (
	"github.com/abiosoft/ishell"
	"github.com/epicagency/pdns-manager/commands"
)

func main() {
	// create new shell.
	// by default, new shell includes 'exit', 'help' and 'clear' commands.
	shell := ishell.New()

	// display welcome info.
	shell.Println("PowerDNS Manager")

	commands.Init(shell)

	// start shell
	shell.Start()

	//zone := new(pdns.Zone)
	//zone.DNSSEC = false
	//zone.Kind = "Master"
	//zone.LastCheck = 0
	//zone.Name = "test.com"
	//zone.Type = "Zone"
	//zone.Records = make([]*pdns.Record, 6)
	//zone.Nameservers = []string{}

	//zone.Records[0] = new(pdns.Record)
	//zone.Records[0].Name = "test.com"
	//zone.Records[0].Type = "SOA"
	//zone.Records[0].Content = "dns1.epic-sys.io. info.epic.net. 2016012001 3600 600 2600 60"
	//zone.Records[0].Priority = 0
	//zone.Records[0].Disabled = false
	//zone.Records[0].TTL = 300
	//zone.Records[1] = new(pdns.Record)
	//zone.Records[1].Name = "test.com"
	//zone.Records[1].Type = "NS"
	//zone.Records[1].Content = "dns1.epic-sys.io"
	//zone.Records[1].Priority = 0
	//zone.Records[1].Disabled = false
	//zone.Records[1].TTL = 300
	//zone.Records[2] = new(pdns.Record)
	//zone.Records[2].Name = "test.com"
	//zone.Records[2].Type = "NS"
	//zone.Records[2].Content = "dns2.epic-sys.io"
	//zone.Records[2].Priority = 0
	//zone.Records[2].Disabled = false
	//zone.Records[2].TTL = 300
	//zone.Records[3] = new(pdns.Record)
	//zone.Records[3].Name = "test.com"
	//zone.Records[3].Type = "NS"
	//zone.Records[3].Content = "dns3.epic-sys.io"
	//zone.Records[3].Priority = 0
	//zone.Records[3].Disabled = false
	//zone.Records[3].TTL = 300
	//zone.Records[4] = new(pdns.Record)
	//zone.Records[4].Name = "test.com"
	//zone.Records[4].Type = "A"
	//zone.Records[4].Content = "127.0.0.1"
	//zone.Records[4].Priority = 0
	//zone.Records[4].Disabled = false
	//zone.Records[4].TTL = 300
	//zone.Records[5] = new(pdns.Record)
	//zone.Records[5].Name = "www.test.com"
	//zone.Records[5].Type = "CNAME"
	//zone.Records[5].Content = "test.com"
	//zone.Records[5].Priority = 0
	//zone.Records[5].Disabled = false
	//zone.Records[5].TTL = 300

	//request := gorequest.New()
	//request.Type("json")
	//request.Post("http://dns1.epic-sys.io/servers/localhost/zones")
	//request.Set("X-API-Key", "bisque.tutelage.organist.payment")
	//request.SendStruct(zone)
	//_, body, errs := request.End()
	//spew.Dump(body)
	//spew.Dump(errs)

}
