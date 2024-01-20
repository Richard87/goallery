// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is the API Specification for Goallery.",
    "title": "Goallery",
    "version": "1.23.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/auth/login": {
      "post": {
        "security": [],
        "description": "Get JWT token",
        "tags": [
          "auth"
        ],
        "summary": "Get JWT token",
        "operationId": "getToken",
        "parameters": [
          {
            "description": "Credentials",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/AuthRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/AuthResponse"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/ProblemDetails"
            }
          }
        }
      }
    },
    "/images": {
      "get": {
        "description": "Get all images",
        "tags": [
          "images"
        ],
        "summary": "Get all images",
        "operationId": "getImages",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Image"
              }
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/ProblemDetails"
            }
          }
        }
      }
    },
    "/images/{id}": {
      "get": {
        "description": "Get image by id",
        "tags": [
          "images"
        ],
        "summary": "Get image by id",
        "operationId": "getImageById",
        "parameters": [
          {
            "type": "string",
            "description": "Image id",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Image"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/ProblemDetails"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "AuthRequest": {
      "type": "object",
      "required": [
        "username",
        "password"
      ],
      "properties": {
        "password": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "AuthResponse": {
      "type": "object",
      "required": [
        "token"
      ],
      "properties": {
        "token": {
          "type": "string"
        }
      }
    },
    "Image": {
      "type": "object",
      "required": [
        "id",
        "filename",
        "description",
        "mime",
        "tags",
        "created",
        "updated",
        "size",
        "width",
        "height",
        "features"
      ],
      "properties": {
        "created": {
          "type": "string",
          "format": "date-time"
        },
        "description": {
          "type": "string"
        },
        "features": {
          "type": "object",
          "$ref": "#/definitions/ImageFeature"
        },
        "filename": {
          "type": "string"
        },
        "height": {
          "type": "integer",
          "format": "int64"
        },
        "id": {
          "type": "string"
        },
        "mime": {
          "type": "string"
        },
        "size": {
          "type": "integer",
          "format": "int64"
        },
        "tags": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "updated": {
          "type": "string",
          "format": "date-time"
        },
        "width": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "ImageFeature": {
      "type": "object",
      "properties": {
        "plugin.blurryimage": {
          "type": "string"
        }
      }
    },
    "ImageList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Image"
      }
    },
    "ProblemDetails": {
      "type": "object",
      "required": [
        "status",
        "title"
      ],
      "properties": {
        "detail": {
          "type": "string"
        },
        "status": {
          "type": "integer",
          "format": "int32"
        },
        "title": {
          "type": "string"
        }
      }
    },
    "User": {
      "type": "object",
      "required": [
        "id",
        "username",
        "password",
        "token"
      ],
      "properties": {
        "id": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "token": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  },
  "security": [
    {
      "bearer": []
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is the API Specification for Goallery.",
    "title": "Goallery",
    "version": "1.23.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/auth/login": {
      "post": {
        "security": [],
        "description": "Get JWT token",
        "tags": [
          "auth"
        ],
        "summary": "Get JWT token",
        "operationId": "getToken",
        "parameters": [
          {
            "description": "Credentials",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/AuthRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/AuthResponse"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/ProblemDetails"
            }
          }
        }
      }
    },
    "/images": {
      "get": {
        "description": "Get all images",
        "tags": [
          "images"
        ],
        "summary": "Get all images",
        "operationId": "getImages",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Image"
              }
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/ProblemDetails"
            }
          }
        }
      }
    },
    "/images/{id}": {
      "get": {
        "description": "Get image by id",
        "tags": [
          "images"
        ],
        "summary": "Get image by id",
        "operationId": "getImageById",
        "parameters": [
          {
            "type": "string",
            "description": "Image id",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Image"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/ProblemDetails"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "AuthRequest": {
      "type": "object",
      "required": [
        "username",
        "password"
      ],
      "properties": {
        "password": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "AuthResponse": {
      "type": "object",
      "required": [
        "token"
      ],
      "properties": {
        "token": {
          "type": "string"
        }
      }
    },
    "Image": {
      "type": "object",
      "required": [
        "id",
        "filename",
        "description",
        "mime",
        "tags",
        "created",
        "updated",
        "size",
        "width",
        "height",
        "features"
      ],
      "properties": {
        "created": {
          "type": "string",
          "format": "date-time"
        },
        "description": {
          "type": "string"
        },
        "features": {
          "type": "object",
          "$ref": "#/definitions/ImageFeature"
        },
        "filename": {
          "type": "string"
        },
        "height": {
          "type": "integer",
          "format": "int64"
        },
        "id": {
          "type": "string"
        },
        "mime": {
          "type": "string"
        },
        "size": {
          "type": "integer",
          "format": "int64"
        },
        "tags": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "updated": {
          "type": "string",
          "format": "date-time"
        },
        "width": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "ImageFeature": {
      "type": "object",
      "properties": {
        "plugin.blurryimage": {
          "type": "string"
        }
      }
    },
    "ImageList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Image"
      }
    },
    "ProblemDetails": {
      "type": "object",
      "required": [
        "status",
        "title"
      ],
      "properties": {
        "detail": {
          "type": "string"
        },
        "status": {
          "type": "integer",
          "format": "int32"
        },
        "title": {
          "type": "string"
        }
      }
    },
    "User": {
      "type": "object",
      "required": [
        "id",
        "username",
        "password",
        "token"
      ],
      "properties": {
        "id": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "token": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  },
  "security": [
    {
      "bearer": []
    }
  ]
}`))
}
