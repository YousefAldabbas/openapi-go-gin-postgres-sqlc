// Code generated by openapi-generator. DO NOT EDIT.

package handlers

import (
	"net/http"

	gen "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/gen"
	services "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/gin-gonic/gin"
)

// StoreApi handles routes with the StoreApi tag.
type StoreApi struct {
	svc services.StoreApi
	// add your own services, etc. as required
}

// NewStoreApi returns a new handler for StoreApi.
// Edit as required
// TODO rewriting handler methods based on current postgen:
// see https://eli.thegreenplace.net/2021/rewriting-go-source-code-with-ast-tooling/
// simpler solutions based on drawbacks (complicated, comments not attached to nodes):
// - https://github.com/dave/dst
// - https://github.com/uber-go/gopatch
func NewStoreApi(svc services.StoreApi) *StoreApi {
	return &StoreApi{
		svc: svc,
	}
}

// Register connects the handlers to a router.
func (t *StoreApi) Register(r *gin.Engine) {
	gen.RegisterRoute(r, gen.Route{
		Name:        "DeleteOrder",
		Method:      http.MethodDelete,
		Pattern:     "/v2/store/order/:orderId",
		HandlerFunc: t.DeleteOrder,
	})
	gen.RegisterRoute(r, gen.Route{
		Name:        "GetInventory",
		Method:      http.MethodGet,
		Pattern:     "/v2/store/inventory",
		HandlerFunc: t.GetInventory,
	})
	gen.RegisterRoute(r, gen.Route{
		Name:        "GetOrderById",
		Method:      http.MethodGet,
		Pattern:     "/v2/store/order/:orderId",
		HandlerFunc: t.GetOrderById,
	})
	gen.RegisterRoute(r, gen.Route{
		Name:        "PlaceOrder",
		Method:      http.MethodPost,
		Pattern:     "/v2/store/order",
		HandlerFunc: t.PlaceOrder,
	})
}

// DeleteOrder delete purchase order by id.
func (t *StoreApi) DeleteOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetInventory returns pet inventories by status.
func (t *StoreApi) GetInventory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetOrderById find purchase order by id.
func (t *StoreApi) GetOrderById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// PlaceOrder place an order for a pet.
func (t *StoreApi) PlaceOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
