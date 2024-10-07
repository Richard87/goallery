package swagger

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/Richard87/goallery/api"
	"github.com/Richard87/goallery/pkg/router"
)

//go:embed embed
var swagfs embed.FS

func NewController() router.RouteMapper {
	spec := api.PathToRawSpec("swagger")
	static, _ := fs.Sub(swagfs, "embed")

	return func(router *http.ServeMux) {
		router.HandleFunc("GET /swagger/swagger.json", func(w http.ResponseWriter, _ *http.Request) {
			bytes, _ := spec["swagger"]()
			_, _ = w.Write(bytes)
		})
		router.Handle("/swagger/", http.StripPrefix("/swagger", http.FileServer(http.FS(static))))
	}
}
