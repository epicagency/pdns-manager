package commands

import (
	"os"
	"text/template"

	"github.com/epicagency/pdns-manager/pdns"
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
{{range $i, $record := .Zone.Records}}{{if eq $.Filter "" $record.Type}}[{{printf "%3d" $i}}]  {{printf $fmt $record.Name}} {{printf "%5d" $record.TTL}} IN {{printf $fmt2 $record.Type}} {{printf "%3d" $record.Priority}} {{$record.Content}} {{if $record.Disabled}}(Disabled){{end}}
{{end}}{{end}}
`

	zone, errs := pdns.GetZone(args[0])

	if errs != nil {
		for err := range errs {
			shell.Println(err)
		}
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
		Filter        string
	}{
		zone,
		max_length,
		max_type_length,
		"",
	}
	if len(args) > 1 && args[1] != "ANY" {
		data.Filter = args[1]
	}

	tmpl, err := template.New("zone").Parse(zone_template)
	if err != nil {
		return "", err
	}
	tmpl.Execute(os.Stdout, data)
	return "", nil
}
