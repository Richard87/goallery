// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"time"
)

const (
	BearerScopes = "bearer.Scopes"
)

// AuthRequest defines model for AuthRequest.
type AuthRequest struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// AuthResponse defines model for AuthResponse.
type AuthResponse struct {
	Token string `json:"token"`
}

// Image defines model for Image.
type Image struct {
	Created     time.Time    `json:"created"`
	Description string       `json:"description"`
	Features    ImageFeature `json:"features"`
	Filename    string       `json:"filename"`
	Height      int64        `json:"height"`
	Id          string       `json:"id"`
	Mime        string       `json:"mime"`
	Size        int64        `json:"size"`
	Src         string       `json:"src"`
	Tags        []string     `json:"tags"`
	Updated     time.Time    `json:"updated"`
	Width       int64        `json:"width"`
}

// ImageFeature defines model for ImageFeature.
type ImageFeature struct {
	PluginBlurryimage *string `json:"plugin.blurryimage,omitempty"`
}

// ProblemDetails defines model for ProblemDetails.
type ProblemDetails struct {
	Detail *string `json:"detail,omitempty"`
	Status int32   `json:"status"`
	Title  string  `json:"title"`
}

// GetTokenJSONRequestBody defines body for GetToken for application/json ContentType.
type GetTokenJSONRequestBody = AuthRequest
