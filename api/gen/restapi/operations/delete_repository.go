// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/treeverse/lakefs/api/gen/models"
)

// DeleteRepositoryHandlerFunc turns a function with the right signature into a delete repository handler
type DeleteRepositoryHandlerFunc func(DeleteRepositoryParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteRepositoryHandlerFunc) Handle(params DeleteRepositoryParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// DeleteRepositoryHandler interface for that can handle valid delete repository params
type DeleteRepositoryHandler interface {
	Handle(DeleteRepositoryParams, *models.User) middleware.Responder
}

// NewDeleteRepository creates a new http.Handler for the delete repository operation
func NewDeleteRepository(ctx *middleware.Context, handler DeleteRepositoryHandler) *DeleteRepository {
	return &DeleteRepository{Context: ctx, Handler: handler}
}

/*DeleteRepository swagger:route DELETE /repositories/{repositoryId} deleteRepository

delete repository

*/
type DeleteRepository struct {
	Context *middleware.Context
	Handler DeleteRepositoryHandler
}

func (o *DeleteRepository) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteRepositoryParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.User
	if uprinc != nil {
		principal = uprinc.(*models.User) // this is really a models.User, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}