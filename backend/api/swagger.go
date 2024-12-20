// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RWTY/bNhD9KwLbW2VrP4IedMs2SODm0KBZoIeFD7Q4lriVSGZI7VpZ+L8XGkoryaJj",
	"Z5uiyU0WZ0Zv3nuc8RPLdGW0AuUsS5+YzQqoOD2+rl3xJ3yqwbr2p0FtAJ0EOjTc2keNon12jQGWMutQ",
	"qpztY1ZbQMUrCBzuY4bwqZYIgqV3Q5lR0jruk/TmHjLXVvRYrNHKwhyM03+DOv0xHxYqv6p4HqibIXAH",
	"1ONWY8UdS5ngDhZOVsDieeMCbIbSOKlVkJgtcFejL/4zwpal7KdkECDp2E8Iz1sfTHmyhCOExqwAmRdu",
	"AlIq9+urAaBUDnLANliGFavkkeJWfoYzS1vMgiUcz6lh6aCy4Qj/giPyhvxjxNfx/iiFK86CeWCJXuGp",
	"dCOhRtw/M00sdpx1DPnuu16HBnpkRz3Xazy/X2WdS7XclDViI3t7zhrvwrY8gynH57rrLc9gLsH+BF7K",
	"mmHeXY4wjoyxuwq/b47EN8H4A+V2l4wqUxnKCZH8AfWmhOoNOC5LO4cs6CDsfMddbQ9NdX0V9L6Trjxj",
	"3HU1+/g54va7kNUoXfOxVcuj3ABHQNJXsZQVwAUgi5kfCTQcNcrPvDNvr6WR76EVs732aqt9v6MBxW4L",
	"aSNpI1dA9PrDKvpoIJNbmVGhaKsxeqd5WQI2y2fMKevfsZg9AFpf6nJ5db28aKnQBhQ3kqXsenmxbLUx",
	"3BXUR8JrVySlziWNR6P9ZpmCegcu+v2v28hPa6qHBGgl/Oltd4B+N91o0dCw1sqBooLcmLLrIrm3fhZ7",
	"85+6GuOdR8RNsf2GIEA5yUvLxso6rIGk9huKmr26uPjGsLr1F8D1x/uW+Vff8IsHFyfwzRsuooGqwbcs",
	"vVvHzNZVxbEJ6OnXwR1rvdDegN0i0wJyUItO0MVGi2bRebt9pvIJDUFiNocjtuFlGXVhAd+s+pN/JdP5",
	"EzYwVr9P4SZSTTjsteperEdCJE9S7L+oBsVFmyaibRmW46ZZCZoQyCtwgO3XDmtRoC9C468dJ8Pwo/fT",
	"mxiPqDqcx+v/8JZ2qv8IKk+1OSVzIvSjKjUXR/V+0wV8UfQ+6HtVnrAnudxOdXhe/xupOK29+ZL3qfcG",
	"8pfmGvXiVPuQ/7KrypemP8LGfG3uD7WBjphz7vppjeF/1926dY8FfOh9WmPJUpZwI5OHS7anNWY84IUY",
	"/mz+L0T8EwAA//+5NgAY0w8AAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
