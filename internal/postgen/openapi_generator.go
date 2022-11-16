package postgen

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/dave/dst"
	"github.com/getkin/kin-openapi/openapi3"
)

type (
	Dir string
	Tag string
)

type Method struct {
	// Name is the method identifier.
	Name string
	// Decl is the method declaration node.
	Decl *dst.FuncDecl
}

var (
	handlerRegex     = regexp.MustCompile("api_(.*).go")
	operationIDRegex = regexp.MustCompile("^[a-zA-Z0-9]*$")
)

func contains[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}

	return false
}

func ensureUnique[T comparable](s []T) error {
	set := make(map[T]struct{})
	for _, element := range s {
		if _, ok := set[element]; ok {
			return fmt.Errorf("element %T not unique", element)
		}
		set[element] = struct{}{}
	}

	return nil
}

/*
Rationale:
- Users can still add new methods to the struct.
In case of generated methods conflicting with existing ones, generation will stop.
- If we need a new route that cant be defined in the spec, e.g. fileserver,
we purposely want that out of the generated handler struct,
so its clear that its outside the spec.
It can still remain in handlers/* as long as its not api_*(!_test).go, e.g. fileserver.go
and it can still follow the same handlers struct pattern for all we care, it wont be touched.
*/
type openapiGenerator struct {
	conf       *Conf
	stderr     io.Writer
	spec       string
	operations map[string][]string
}

// NewOpenapiGenerator returns a new postgen openapiGenerator.
func NewOpenapiGenerator(conf *Conf, stderr io.Writer, spec string) *openapiGenerator {
	operations := make(map[string][]string)

	return &openapiGenerator{
		conf:       conf,
		stderr:     stderr,
		spec:       spec,
		operations: operations,
	}
}

// analyzeSpec ensures specific rules for codegen are met and extracts necessary data.
func (o *openapiGenerator) analyzeSpec() error {
	schemaBlob, err := os.ReadFile(o.spec)
	if err != nil {
		return fmt.Errorf("error opening schema file: %w", err)
	}

	sl := openapi3.NewLoader()

	openapi, err := sl.LoadFromData(schemaBlob)
	if err != nil {
		return fmt.Errorf("error loading openapi spec: %w", err)
	}

	if err = openapi.Validate(sl.Context); err != nil {
		return fmt.Errorf("error validating openapi spec: %w", err)
	}

	for path, pi := range openapi.Paths {
		ops := pi.Operations()
		for method, v := range ops {
			if v.OperationID == "" {
				return fmt.Errorf("path %q: method %q: operationId is required for postgen", path, method)
			}

			if !operationIDRegex.MatchString(v.OperationID) {
				return fmt.Errorf("path %q: method %q: operationId %q does not match pattern %q",
					path, method, v.OperationID, operationIDRegex.String())
			}

			if len(v.Tags) > 1 {
				return fmt.Errorf("path %q: method %q: at most one tag is permitted for postgen", path, method)
			}

			t := "default"
			if len(v.Tags) > 0 {
				t = strings.ToLower(v.Tags[0])
			}

			o.operations[t] = append(o.operations[t], v.OperationID)
		}

		for t, opIDs := range o.operations {
			sort.Slice(opIDs, func(i, j int) bool {
				return opIDs[i] < opIDs[j]
			})
			o.operations[t] = opIDs
		}
	}

	return nil
}

func (o *openapiGenerator) Generate() error {
	err := o.analyzeSpec()
	if err != nil {
		return fmt.Errorf("invalid spec: %w", err)
	}

	src, err := o.generateOpIDs()
	if err != nil {
		return fmt.Errorf("could not generate operation IDs: %w", err)
	}
	fname := path.Join(string(o.conf.OutHandlersDir), "operation_ids.gen.go")

	f, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o660)
	if err != nil {
		return fmt.Errorf("could not open %s: %w", fname, err)
	}

	if _, err = f.Write(src); err != nil {
		return fmt.Errorf("could not write opId template: %w", err)
	}

	return nil
}

// generateOpIDs fills in a template with all operation IDs to a dest.
func (o *openapiGenerator) generateOpIDs() ([]byte, error) {
	funcs := template.FuncMap{
		// "stringsJoin": func(elems []string, prefix string, suffix string, sep string) string {
		// 	for i, e := range elems {
		// 		elems[i] = prefix + e + suffix
		// 	}

		// 	return strings.Join(elems, sep)
		// },
	}

	t := template.Must(template.New("").Funcs(funcs).Parse(`// Code generated by postgen. DO NOT EDIT.

package rest

{{ $first := true }}
type operationID string

const ({{range $tag, $opIDs := .Operations}}
// Operation IDs for the '{{$tag}}' tag.

{{range $opIDs -}}
{{.}}                operationID = "{{.}}"
{{end -}}
{{end}}
	)

	`))

	buf := &bytes.Buffer{}

	params := map[string]interface{}{
		"Operations": o.operations,
	}

	if err := t.Execute(buf, params); err != nil {
		return nil, fmt.Errorf("could not execute template: %w", err)
	}

	src, err := format.Source(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("could not format opId template: %w", err)
	}

	return src, nil
}
