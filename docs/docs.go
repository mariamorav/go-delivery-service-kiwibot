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
        "/bots/assign-order": {
            "post": {
                "description": "This endpoint allows you assign an available bot to an a pending order.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bots"
                ],
                "summary": "assign a bot to an order",
                "parameters": [
                    {
                        "description": "The body request to assign a bot",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/BodyAssignBot"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/AssignBotToOrderResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/bots/new": {
            "post": {
                "description": "This endpoint allows you to create a new bot with available status as default.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bots"
                ],
                "summary": "create a new bot",
                "parameters": [
                    {
                        "description": "The body request to create a Bot",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/CreateBotRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/CreateBotResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/bots/{zone_id}": {
            "get": {
                "description": "This endpoint allows you to list the bots located in the zone_id submitted.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bots"
                ],
                "summary": "get bots by a zone_id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id of the zone where you want to find bots",
                        "name": "zone_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Start point of documents search",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Limit number of results returned by search",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/GetBotsByZoneResponse"
                        }
                    },
                    "404": {
                        "description": "there are no results for your search",
                        "schema": {
                            "$ref": "#/definitions/MessageNotFound"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/deliveries": {
            "get": {
                "description": "This endpoint allows you to list deliveries paginated.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deliveries"
                ],
                "summary": "get deliveries paginated",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Start point of documents search",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Limit number of results returned by search",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "options: 'asc' or 'desc' to order the documents by creation_date. Default is 'asc'",
                        "name": "order",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/GetDeliveriesResponse"
                        }
                    },
                    "404": {
                        "description": "there are no results for your search",
                        "schema": {
                            "$ref": "#/definitions/MessageNotFound"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/deliveries/assing-pending-orders": {
            "get": {
                "description": "This endpoint allows you to assign available bots to pending orders.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deliveries"
                ],
                "summary": "assign bots to all pending orders",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/AssignBotsToAllPendingOrdersResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/deliveries/new": {
            "post": {
                "description": "This endpoint allows you to create a new delivery with pending state as default.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deliveries"
                ],
                "summary": "create a new delivery",
                "parameters": [
                    {
                        "description": "The body request to create a delivery",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/CreateDeliveryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/CreateDeliveryResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/deliveries/{id}": {
            "get": {
                "description": "This endpoint allows you to get a delivery by the id.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deliveries"
                ],
                "summary": "get delivery by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id of the delivery",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/GetBotsByZoneResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "AssignBotToOrderResponse": {
            "type": "object",
            "properties": {
                "bot": {
                    "$ref": "#/definitions/models.Bot"
                },
                "message": {
                    "type": "string"
                },
                "order": {
                    "$ref": "#/definitions/models.Delivery"
                }
            }
        },
        "AssignBotsToAllPendingOrdersResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "orders_unassigned": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Delivery"
                    }
                }
            }
        },
        "BodyAssignBot": {
            "type": "object",
            "required": [
                "order_id"
            ],
            "properties": {
                "order_id": {
                    "description": "The order id of delivery pending to assign\nexample: \"QgDz4iJONRfbHGSvqmSl\"\nrequired: true",
                    "type": "string"
                }
            }
        },
        "CreateBotRequest": {
            "type": "object",
            "properties": {
                "location": {
                    "description": "Location coordinates of the bot",
                    "$ref": "#/definitions/models.Location"
                },
                "zone_id": {
                    "description": "ZoneId of the zone where the bot is located",
                    "type": "string"
                }
            }
        },
        "CreateBotResponse": {
            "type": "object",
            "properties": {
                "bot": {
                    "$ref": "#/definitions/models.Bot"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "CreateDeliveryRequest": {
            "type": "object",
            "required": [
                "zone_id"
            ],
            "properties": {
                "dropoff": {
                    "$ref": "#/definitions/models.Dropoff"
                },
                "pickup": {
                    "$ref": "#/definitions/models.Pickup"
                },
                "zone_id": {
                    "type": "string"
                }
            }
        },
        "CreateDeliveryResponse": {
            "type": "object",
            "properties": {
                "delivery": {
                    "$ref": "#/definitions/models.Delivery"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "GetBotsByZoneResponse": {
            "type": "object",
            "properties": {
                "next_url": {
                    "type": "string"
                },
                "previous_url": {
                    "type": "string"
                },
                "results": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Bot"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "GetDeliveriesResponse": {
            "type": "object",
            "properties": {
                "next_url": {
                    "type": "string"
                },
                "previous_url": {
                    "type": "string"
                },
                "results": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Delivery"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "MessageNotFound": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Bot": {
            "type": "object",
            "required": [
                "location",
                "zone_id"
            ],
            "properties": {
                "id": {
                    "type": "string"
                },
                "location": {
                    "$ref": "#/definitions/models.Location"
                },
                "status": {
                    "type": "string"
                },
                "zone_id": {
                    "type": "string"
                }
            }
        },
        "models.Delivery": {
            "type": "object",
            "required": [
                "dropoff",
                "pickup",
                "zone_id"
            ],
            "properties": {
                "creation_date": {
                    "type": "string"
                },
                "dropoff": {
                    "$ref": "#/definitions/models.Dropoff"
                },
                "id": {
                    "type": "string"
                },
                "pickup": {
                    "$ref": "#/definitions/models.Pickup"
                },
                "state": {
                    "type": "string"
                },
                "zone_id": {
                    "type": "string"
                }
            }
        },
        "models.Dropoff": {
            "type": "object",
            "required": [
                "dropoff_lat",
                "dropoff_lon"
            ],
            "properties": {
                "dropoff_lat": {
                    "type": "number"
                },
                "dropoff_lon": {
                    "type": "number"
                }
            }
        },
        "models.Location": {
            "type": "object",
            "required": [
                "lat",
                "lon"
            ],
            "properties": {
                "lat": {
                    "type": "number"
                },
                "lon": {
                    "type": "number"
                }
            }
        },
        "models.Pickup": {
            "type": "object",
            "required": [
                "pickup_lat",
                "pickup_lon"
            ],
            "properties": {
                "pickup_lat": {
                    "type": "number"
                },
                "pickup_lon": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
