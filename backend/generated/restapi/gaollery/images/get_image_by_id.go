// Code generated by go-swagger; DO NOT EDIT.

package images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetImageByIDHandlerFunc turns a function with the right signature into a get image by Id handler
type GetImageByIDHandlerFunc func(GetImageByIDParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetImageByIDHandlerFunc) Handle(params GetImageByIDParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetImageByIDHandler interface for that can handle valid get image by Id params
type GetImageByIDHandler interface {
	Handle(GetImageByIDParams, interface{}) middleware.Responder
}

// NewGetImageByID creates a new http.Handler for the get image by Id operation
func NewGetImageByID(ctx *middleware.Context, handler GetImageByIDHandler) *GetImageByID {
	return &GetImageByID{Context: ctx, Handler: handler}
}

/*
	GetImageByID swagger:route GET /images/{id} images getImageById

# Get image by id

Get image by id
*/
type GetImageByID struct {
	Context *middleware.Context
	Handler GetImageByIDHandler
}

func (o *GetImageByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetImageByIDParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
