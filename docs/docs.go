// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/": {
            "get": {
                "description": "Home",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Home"
                ],
                "summary": "Home based on parameter",
                "operationId": "home",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Base"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.Home"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Base"
                        }
                    }
                }
            }
        },
        "/auth/google": {
            "get": {
                "description": "Auth Google Redirection",
                "tags": [
                    "auth"
                ],
                "summary": "Get Auth Google Redirection based on parameter",
                "operationId": "get-auth-google",
                "responses": {
                    "307": {
                        "description": "Temporary Redirect"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/auth/google/callback": {
            "get": {
                "description": "Auth Google Callback",
                "tags": [
                    "auth"
                ],
                "summary": "Get Auth Google Callback based on parameter",
                "operationId": "get-auth-google-callback",
                "responses": {
                    "307": {
                        "description": "Temporary Redirect"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/auth/me": {
            "get": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "Auth Me",
                "tags": [
                    "auth"
                ],
                "summary": "Get Auth Me based on parameter",
                "operationId": "get-auth-me",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "Unauthorized"
                    }
                }
            }
        },
        "/cctv": {
            "get": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cctv"
                ],
                "summary": "Get all CCTV records",
                "operationId": "get-all-cctv",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Base"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.CCTVItem"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Base"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "message": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.CCTVItem": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "latitude": {
                    "type": "number"
                },
                "link": {
                    "type": "string"
                },
                "longitude": {
                    "type": "number"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "response.Base": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string",
                    "example": "Message!"
                }
            }
        },
        "response.Home": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "github": {
                    "type": "string"
                },
                "linkedin": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "website": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "AccessToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "BE Tilik Jalan",
	Description:      "This is an API documentation of BE Tilik Jalan",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
