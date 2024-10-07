package templates

import (
	"net/http"

	"github.com/Richard87/goallery/pkg/router"
	"github.com/a-h/templ"
)

func NewController() router.RouteMapper {
	component := hello("John")

	return func(router *http.ServeMux) {
		router.Handle("/", templ.Handler(component))
	}
}
