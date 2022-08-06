package main

import (
	"fmt"
	"os"
	"path"
	"sort"

	postgen "github.com/danicc097/openapi-go-gin-postgres-sqlc/postgen"
)

func main() {
	cwd, _ := os.Getwd()
	genHandlers := postgen.ParseHandlers(path.Join(cwd, "internal/gen/api_*.go"))
	localHandlers := postgen.ParseHandlers(path.Join(cwd, "internal/handlers/api_*.go"))

	missingHandlers := []postgen.Handler{}
	for k, v := range genHandlers {
		if _, ok := localHandlers[k]; !ok {
			missingHandlers = append(missingHandlers, v)
		}
	}
	sort.Slice(missingHandlers, func(i, j int) bool {
		return missingHandlers[i].OperationId < missingHandlers[j].OperationId
	})

	fmt.Printf("Generating non-implemented route handlers: %s\n", missingHandlers)
	outPath := path.Join(cwd, "internal/handlers/not_implemented.gen.go")
	f, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}
	postgen.GenerateHandlers(missingHandlers, f)
}
