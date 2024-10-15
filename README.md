# Goallery

Goallery is a simple image gallery written in Go. It is designed to be used as a standalone application with a wide array of backends and data-storage. 

## Features
- [ ] Create thumbnails (wip, make it into a `ImageFeaturePlugin`)
- [ ] Scale images and set cache headers
- [ ] Frontend Gallery
- [ ] Pluggable image features
- [ ] Authentication
- [ ] Persistent db storage (sqlite, postgres, mysql, etc)
- [ ] Pluggable storage backends (s3, azure storage, google cloud storage, etc)
- [ ] Upload image(s)
- [ ] Create albums
- [ ] face recognition
- [ ] Use chromem-go for storage
- [ ] Use nats with leader election for high availability
- [ ] use tensorflow with goface for face recognition (https://github.com/jdeng/goface)

## Contributing

When adding a feature, start with updating the swagger.json with the new details. Then run `make generate` to generate the new code.
You probably want to modify the `/definitions/ImageFeature/properties` with your new feature. Then create a new `ImageFeaturePlugin` to generate data.

## TO RUN
`make run` run the frontend and backend respectively. By default it currently looks for images in `./photos` in the root folder
