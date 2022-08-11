// Code generated by openapi-generator. DO NOT EDIT.

package handlers

import (
	"net/http"

	gen "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/gen"
	services "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/gin-gonic/gin"
)

// FakeApi handles routes with the FakeApi tag.
type FakeApi struct {
	svc services.FakeApi
	// add your own services, etc. as required
}

// NewFakeApi returns a new handler for FakeApi.
// Edit as required
// TODO rewriting handler methods based on current postgen:
// see https://eli.thegreenplace.net/2021/rewriting-go-source-code-with-ast-tooling/
// simpler solutions based on drawbacks (complicated, comments not attached to nodes):
// - https://github.com/dave/dst
// - https://github.com/uber-go/gopatch
func NewFakeApi(svc services.FakeApi) *FakeApi {
	return &FakeApi{
		svc: svc,
	}
}

// Register connects the handlers to a router.
func (t *FakeApi) Register(r *gin.Engine) {
	gen.RegisterRoute(r, gen.Route{
		Name:        "FakeDataFile",
		Method:      http.MethodGet,
		Pattern:     "/v2/fake/data_file",
		HandlerFunc: t.FakeDataFile,
	})
}

// FakeDataFile test data_file to ensure it's escaped correctly.
func (t *FakeApi) FakeDataFile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
