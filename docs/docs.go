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
        "contact": {
            "name": "API Support"
        },
        "license": {
            "name": "Apache 2.0"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/accounts": {
            "post": {
                "description": "Create Account",
                "tags": [
                    "Account"
                ],
                "summary": "Create Account",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string",
                            "example": "{\"device_id\":\"string\",\"session_id\":\"string\",\"platform\":\"int\",\"app_version\":\"string\"}"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/param.AddAccountResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/param.BadRequestHttpError"
                        }
                    },
                    "401": {
                        "description": "unauthorized",
                        "schema": {
                            "$ref": "#/definitions/param.UnAuthorizedHttpError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/param.UnprocessableEntity"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/param.InternalError"
                        }
                    }
                }
            }
        },
        "/app-config": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get App Config",
                "tags": [
                    "AppConfig"
                ],
                "summary": "Get App Config",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/param.GetAppConfigResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/param.BadRequestHttpError"
                        }
                    },
                    "401": {
                        "description": "unauthorized",
                        "schema": {
                            "$ref": "#/definitions/param.UnAuthorizedHttpError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/param.UnprocessableEntity"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/param.InternalError"
                        }
                    }
                }
            }
        },
        "/characters": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "List of characters",
                "tags": [
                    "Characters"
                ],
                "summary": "List of characters",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "version of your app",
                        "name": "version",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/param.ListCharactersResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/param.BadRequestHttpError"
                        }
                    },
                    "401": {
                        "description": "unauthorized",
                        "schema": {
                            "$ref": "#/definitions/param.UnAuthorizedHttpError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/param.UnprocessableEntity"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/param.InternalError"
                        }
                    }
                }
            }
        },
        "/configs": {
            "get": {
                "description": "Get Configs",
                "tags": [
                    "Configs"
                ],
                "summary": "Get Configs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/param.GetConfigsResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/param.InternalError"
                        }
                    }
                }
            }
        },
        "/contexts": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "List of Contexts",
                "tags": [
                    "Context"
                ],
                "summary": "List of Contexts",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "request body",
                        "name": "version",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/param.ListContextsResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/param.BadRequestHttpError"
                        }
                    },
                    "401": {
                        "description": "unauthorized",
                        "schema": {
                            "$ref": "#/definitions/param.UnAuthorizedHttpError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/param.UnprocessableEntity"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/param.InternalError"
                        }
                    }
                }
            }
        },
        "/feedbacks": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "List of Feedbacks",
                "tags": [
                    "Feedback"
                ],
                "summary": "List of Feedbacks",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "request body",
                        "name": "feedback_type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/param.BadRequestHttpError"
                        }
                    },
                    "401": {
                        "description": "unauthorized",
                        "schema": {
                            "$ref": "#/definitions/param.UnAuthorizedHttpError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/param.UnprocessableEntity"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/param.InternalError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Add Feedback",
                "tags": [
                    "Feedback"
                ],
                "summary": "Add Feedback",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string",
                            "example": "{\"account_id\":\"uint\",\"device_id\":\"string\",\"feedback_type\":\"int\",\"reported_account_id\":\"uint\",\"comment\":\"string\"}"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/param.BadRequestHttpError"
                        }
                    },
                    "401": {
                        "description": "unauthorized",
                        "schema": {
                            "$ref": "#/definitions/param.UnAuthorizedHttpError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/param.UnprocessableEntity"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/param.InternalError"
                        }
                    }
                }
            }
        },
        "/health-check": {
            "get": {
                "description": "Health check",
                "tags": [
                    "Health check"
                ],
                "summary": "Health check",
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/param.BadRequestHttpError"
                        }
                    }
                }
            }
        },
        "/languages": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "List of languages",
                "tags": [
                    "languages"
                ],
                "summary": "List of languages",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "version of your app",
                        "name": "version",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/param.ListLanguagesResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/param.BadRequestHttpError"
                        }
                    },
                    "401": {
                        "description": "unauthorized",
                        "schema": {
                            "$ref": "#/definitions/param.UnAuthorizedHttpError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/param.UnprocessableEntity"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/param.InternalError"
                        }
                    }
                }
            }
        },
        "/lets-talk": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Lets Talk",
                "tags": [
                    "Talk"
                ],
                "summary": "Lets Talk",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string",
                            "example": "{\"language_id\":\"string\",\"character_id\":\"int\",\"device_id\":\"string\",\"timestamp\":\"string\",\"account_id\":\"string\"}"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/param.LetsTalkResponse"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/param.BadRequestHttpError"
                        }
                    },
                    "401": {
                        "description": "unauthorized",
                        "schema": {
                            "$ref": "#/definitions/param.UnAuthorizedHttpError"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/param.NotFound"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/param.UnprocessableEntity"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/param.InternalError"
                        }
                    }
                }
            }
        },
        "/subscribe/{token}": {
            "post": {
                "description": "Subscribe",
                "tags": [
                    "Socket"
                ],
                "summary": "Subscribe",
                "parameters": [
                    {
                        "type": "string",
                        "description": "request param",
                        "name": "token",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/param.LetsTalkResponse"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/param.BadRequestHttpError"
                        }
                    },
                    "401": {
                        "description": "unauthorized",
                        "schema": {
                            "$ref": "#/definitions/param.UnAuthorizedHttpError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/param.UnprocessableEntity"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/param.InternalError"
                        }
                    }
                }
            }
        },
        "/token": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Refresh Tokens",
                "tags": [
                    "Token"
                ],
                "summary": "Refresh Tokens",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/param.RefreshTokensResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/param.BadRequestHttpError"
                        }
                    },
                    "401": {
                        "description": "unauthorized",
                        "schema": {
                            "$ref": "#/definitions/param.UnAuthorizedHttpError"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/param.NotFound"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/param.UnprocessableEntity"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/param.InternalError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.AppConfig": {
            "type": "object",
            "properties": {
                "app_version": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "latest_character_version": {
                    "type": "integer"
                },
                "latest_context_version": {
                    "type": "integer"
                },
                "latest_language_version": {
                    "type": "integer"
                },
                "platform": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.Character": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "gender": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "is_deleted": {
                    "type": "boolean"
                },
                "mood": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "version": {
                    "type": "integer"
                }
            }
        },
        "model.Context": {
            "type": "object",
            "properties": {
                "context": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "frame": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "is_deleted": {
                    "type": "boolean"
                },
                "page": {
                    "type": "integer"
                },
                "type": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "version": {
                    "type": "integer"
                }
            }
        },
        "model.Language": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_deleted": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "version": {
                    "type": "integer"
                }
            }
        },
        "param.AccountTokens": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "param.AddAccountResponse": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "integer"
                },
                "tokens": {
                    "$ref": "#/definitions/param.AccountTokens"
                }
            }
        },
        "param.BadRequestHttpError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "code=400, message=bad request"
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        },
        "param.FeedbackUnit": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "value": {
                    "type": "integer"
                }
            }
        },
        "param.GetAppConfigResponse": {
            "type": "object",
            "properties": {
                "app_config": {
                    "$ref": "#/definitions/model.AppConfig"
                }
            }
        },
        "param.GetConfigsResponse": {
            "type": "object",
            "properties": {
                "feedbacks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/param.FeedbackUnit"
                    }
                },
                "images": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/param.ImageUnit"
                    }
                },
                "images_static_url": {
                    "type": "string"
                },
                "lets_talk_delay_in_seconds": {
                    "type": "integer"
                },
                "lets_talk_waiting_timeout_in_seconds": {
                    "type": "integer"
                },
                "platforms": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/param.PlatformUnit"
                    }
                }
            }
        },
        "param.ImageUnit": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "param.InternalError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "code=500, message=internal server error"
                },
                "message": {
                    "type": "string",
                    "example": "nil pointer"
                }
            }
        },
        "param.LetsTalkResponse": {
            "type": "object",
            "properties": {
                "waiting_timeout_in_seconds": {
                    "type": "integer"
                }
            }
        },
        "param.ListCharactersResponse": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Character"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "param.ListContextsResponse": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Context"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "param.ListLanguagesResponse": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Language"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "param.NotFound": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "code=404, message=not found"
                },
                "message": {
                    "type": "string",
                    "example": "requested object not found"
                }
            }
        },
        "param.PlatformUnit": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "value": {
                    "type": "integer"
                }
            }
        },
        "param.RefreshTokensResponse": {
            "type": "object",
            "properties": {
                "tokens": {
                    "$ref": "#/definitions/param.AccountTokens"
                }
            }
        },
        "param.UnAuthorizedHttpError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "code=401, message=invalid or expired jwt"
                },
                "message": {
                    "type": "string",
                    "example": "unauthorized"
                }
            }
        },
        "param.UnprocessableEntity": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "code=422, message=Unprocessable Entity"
                },
                "message": {
                    "type": "string",
                    "example": "request parameters are not valid"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "185.255.88.17:8095",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Thanos",
	Description:      "This is Thanos api docs.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
