package ui

import (
	"bytes"
	_ "embed"
	"html/template"
	"io/ioutil"
	"log"
)

//go:embed template/template.tsx
var templateTxt string

func Create() {
	// Generate phase: template.Execute
	var buf bytes.Buffer
	t := template.Must(template.New("meta-txt").Parse(templateTxt))
	params := map[string]interface{}{}
	t.Execute(&buf, params)

	// Output phase
	var out bytes.Buffer
	out.Write(buf.Bytes())
	// body, _ := format.Source(out.Bytes())
	// if err != nil {
	// 	// The user can compile the output to see the error.
	// 	log.Fatalf("warning: internal error: invalid Go generated: %s", err)
	// }

	if err := ioutil.WriteFile("ui/template.tsx", out.Bytes(), 0644); err != nil {
		log.Fatalf("writing output: %s", err)
	}
}
