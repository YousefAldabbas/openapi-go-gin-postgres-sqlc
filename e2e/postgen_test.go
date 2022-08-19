package e2e_test

import (
	"bytes"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/postgen"
	"github.com/google/go-cmp/cmp"
)

func setupTests(t *testing.T) {
	os.Setenv("IS_TESTING", "1")

	cmd := exec.Command(
		"../bin/build",
		"generate-tests-api",
	)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Logf("combined out:\n%s\n", string(out))
		t.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func TestHandlerPostProcessing(t *testing.T) {
	t.Parallel()

	setupTests(t)

	const baseDir = "testdata/postgen/openapi_generator"

	cases := []struct {
		Name string
		Dir  string
	}{
		{
			"Merging",
			"merge_changes",
		},
		{
			"NameClashing",
			"name_clashing",
		},
	}

	for _, test := range cases {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			conf := &postgen.Conf{
				CurrentHandlersDir: path.Join(baseDir, test.Dir, "internal/handlers"),
				GenHandlersDir:     path.Join(baseDir, test.Dir, "internal/gen"),
				OutHandlersDir:     path.Join(baseDir, test.Dir, "got"),
				OutServicesDir:     path.Join(baseDir, test.Dir, "internal/services"),
			}

			err := os.RemoveAll(conf.OutHandlersDir)
			if err != nil {
				t.Fatal(err)
			}
			var stderr bytes.Buffer
			og := postgen.NewOpenapiGenerator(conf, &stderr)

			err = og.Generate()
			if err != nil {
				s := getStderr(t, path.Join(baseDir, test.Dir, "want"))
				if diff := cmp.Diff(s, stderr.String()); s != "" && diff != "" {
					t.Fatalf("stderr differed (-want +got):\n%s", diff)
				}
			}

			pconf := &printer.Config{Mode: printer.TabIndent | printer.UseSpaces, Tabwidth: 8}
			ff, err := filepath.Glob(path.Join(conf.OutHandlersDir, "/*"))
			if err != nil {
				t.Fatalf("OutHandlersDir glob failed")
			}
			for _, f := range ff {
				basename := path.Base(f)
				wp := path.Join(baseDir, test.Dir, "want", basename)
				gp := path.Join(conf.OutHandlersDir, basename)
				wantBlob, _ := os.ReadFile(wp)
				gotBlob, _ := os.ReadFile(gp)
				want := &bytes.Buffer{}
				got := &bytes.Buffer{}

				printContent(t, string(wantBlob), pconf, want)
				printContent(t, string(gotBlob), pconf, got)

				if diff := cmp.Diff(want.String(), got.String()); diff != "" {
					t.Errorf("strings differed (-want +got):\n%s", diff)
				}
			}
		})
	}
}

// printContent normalizes source code and prints to a dest.
func printContent(t *testing.T, content string, pconf *printer.Config, dest io.Writer) {
	t.Helper()
	fset := token.NewFileSet()

	r, err := parser.ParseFile(fset, "", content, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}
	if err := pconf.Fprint(buf, fset, r); err != nil {
		panic(err)
	}

	out, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}

	dest.Write(out)
}
