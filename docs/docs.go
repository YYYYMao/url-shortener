// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/urls": {
            "post": {
                "description": "create short url",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "urls"
                ],
                "summary": "create short url",
                "parameters": [
                    {
                        "description": "url and expire time 2022-12-30T15:03:43.4Z",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/urls.UrlParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/urls.UrlResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/resHandler.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/resHandler.ErrResponse"
                        }
                    }
                }
            }
        },
        "/{url_id}": {
            "get": {
                "description": "redirect url",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "urls"
                ],
                "summary": "redirect url",
                "parameters": [
                    {
                        "type": "string",
                        "description": "url_id",
                        "name": "url_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "302": {
                        "description": "Found"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/resHandler.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/resHandler.ErrResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "resHandler.ErrResponse": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "urls.UrlParam": {
            "type": "object",
            "required": [
                "expireAt",
                "url"
            ],
            "properties": {
                "expireAt": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "urls.UrlResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "shortUrl": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Url Shortener",
	Description:      "Url Shortener API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
