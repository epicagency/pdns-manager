package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"text/template"

	"github.com/epicagency/pdns-manager/pdns"
	"github.com/parnurzeal/gorequest"
)

func show(args ...string) (string, error) {

	const zone_template = `
{{with .Zone}}{{.Name}}
  Comments       : {{.Comments}}
  DNSSEC         : {{.DNSSEC}}
  Kind           : {{.Kind}}
  LastCheck      : {{.LastCheck}}
  Masters        : {{.Masters}}
  Nameservers    : {{.Nameservers}}
  NotifiedSerial : {{.NotifiedSerial}}
  Serial         : {{.Serial}}
  Type           : {{.Type}}
  URL            : {{.URL}}{{end}}

{{$fmt := printf "%%-%ds" .MaxLength}}
{{$fmt2 := printf "%%-%ds" .MaxTypeLength}}
  Records:
{{range $i, $record := .Zone.Records}}[{{printf "%3d" $i}}]  {{printf $fmt $record.Name}} {{printf "%5d" $record.TTL}} IN {{printf $fmt2 $record.Type}} {{printf "%3d" $record.Priority}} {{$record.Content}} {{if $record.Disabled}}(Disabled){{end}}
{{end}}
`

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
	max_length := 0
	max_type_length := 0
	for _, record := range zone.Records {
		if len(record.Name) > max_length {
			max_length = len(record.Name)
		}
		if len(record.Type) > max_type_length {
			max_type_length = len(record.Type)
		}
	}
	data := struct {
		Zone          *pdns.Zone
		MaxLength     int
		MaxTypeLength int
	}{
		zone,
		max_length,
		max_type_length,
	}

	tmpl, err := template.New("zone").Parse(zone_template)
	tmpl.Execute(os.Stdout, data)
	return "", nil
}
