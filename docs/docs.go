// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Tim Seitzinger",
            "email": "tseitzinger@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/images": {
            "get": {
                "description": "Responds with the list of all images as JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "images"
                ],
                "summary": "Get Images array",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Image"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Takes a image JSON and store in DB. Return saved JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "images"
                ],
                "summary": "Store a new image",
                "parameters": [
                    {
                        "description": "Image Directory",
                        "name": "image",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Image"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Image"
                        }
                    }
                }
            }
        },
        "/images/{image_id}": {
            "get": {
                "description": "Returns the image whose ID value matches the ImageID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "images"
                ],
                "summary": "Get single image by ImageID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search image by ImageID",
                        "name": "image_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Image"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete the image whose ID value matches the image_id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "images"
                ],
                "summary": "Delete single image by image_id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "delete image by image_id",
                        "name": "image_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Image"
                        }
                    }
                }
            },
            "patch": {
                "description": "Updates the image whose ID value matches the image_id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "images"
                ],
                "summary": "Update single image by image_id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "update image by image_id",
                        "name": "image_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Image JSON",
                        "name": "image",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Image"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Image"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Image": {
            "type": "object",
            "properties": {
                "imageDateTime": {
                    "type": "string"
                },
                "imageDay": {
                    "type": "integer"
                },
                "imageDirLocation": {
                    "type": "string"
                },
                "imageFileName": {
                    "type": "string"
                },
                "imageHeight": {
                    "type": "integer"
                },
                "imageID": {
                    "type": "integer"
                },
                "imageLat": {
                    "type": "string"
                },
                "imageLon": {
                    "type": "string"
                },
                "imageMonth": {
                    "type": "integer"
                },
                "imageSize": {
                    "type": "string"
                },
                "imageType": {
                    "type": "string"
                },
                "imageWidth": {
                    "type": "integer"
                },
                "imageYear": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Images API",
	Description:      "An Image API in Go using Gin framework.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
