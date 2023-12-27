package frontend

import (
	"embed"
	"io/fs"
)

//go:generate npm run build
//go:embed dist
var frontend embed.FS

// FS returns a FS with SwaggerUI files in root
func FS() fs.FS {
	rootFS, err := fs.Sub(frontend, "dist")
	if err != nil {
		panic(err)
	}
	return rootFS
}
