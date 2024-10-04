package swagger

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/Richard87/goallery/api"
)

//go:embed embed
var swagfs embed.FS

func Handler() http.Handler {
	spec := api.PathToRawSpec("swagger")
	static, _ := fs.Sub(swagfs, "embed")
	mux := http.NewServeMux()
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, _ *http.Request) {
		bytes, _ := spec["swagger"]()
		_, _ = w.Write(bytes)
	})
	mux.Handle("/", http.FileServer(http.FS(static)))
	return mux
}
