// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
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
        "/Data/DeleteData": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "testswager",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data"
                ],
                "summary": "Delete",
                "parameters": [
                    {
                        "type": "string",
                        "default": "th",
                        "description": "(en, th)",
                        "name": "Accept-Language",
                        "in": "header"
                    },
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/delete.Delete"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "410": {
                        "description": "Gone",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    }
                }
            }
        },
        "/Data/DeleteStatus": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "testswager",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data"
                ],
                "summary": "DeleteStatus",
                "parameters": [
                    {
                        "type": "string",
                        "default": "th",
                        "description": "(en, th)",
                        "name": "Accept-Language",
                        "in": "header"
                    },
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/deletestatus.DeleteStatus"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "410": {
                        "description": "Gone",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    }
                }
            }
        },
        "/Data/InsertData": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "testswager",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data"
                ],
                "summary": "InsertData",
                "parameters": [
                    {
                        "type": "string",
                        "default": "th",
                        "description": "(en, th)",
                        "name": "Accept-Language",
                        "in": "header"
                    },
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/insertData.AddInformationRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "410": {
                        "description": "Gone",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    }
                }
            }
        },
        "/Data/SelectGroup": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "testswager",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data"
                ],
                "summary": "SelectGroup",
                "parameters": [
                    {
                        "type": "string",
                        "default": "th",
                        "description": "(en, th)",
                        "name": "Accept-Language",
                        "in": "header"
                    },
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/selectGroup.SelectGroup"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "410": {
                        "description": "Gone",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    }
                }
            }
        },
        "/Data/UpDateData": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "testswager",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data"
                ],
                "summary": "Update",
                "parameters": [
                    {
                        "type": "string",
                        "default": "th",
                        "description": "(en, th)",
                        "name": "Accept-Language",
                        "in": "header"
                    },
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/updateData.UpdateData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "410": {
                        "description": "Gone",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    }
                }
            }
        },
        "/login/login": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "testswager",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login\u0026logout"
                ],
                "summary": "login",
                "parameters": [
                    {
                        "type": "string",
                        "default": "th",
                        "description": "(en, th)",
                        "name": "Accept-Language",
                        "in": "header"
                    },
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/login.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AlComUserModel"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "410": {
                        "description": "Gone",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    }
                }
            }
        },
        "/test/logout": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "testswager",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login\u0026logout"
                ],
                "summary": "logout",
                "parameters": [
                    {
                        "type": "string",
                        "default": "th",
                        "description": "(en, th)",
                        "name": "Accept-Language",
                        "in": "header"
                    },
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/logout.LogoutRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "410": {
                        "description": "Gone",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    }
                }
            }
        },
        "/test/refreshTK": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "testswager",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login\u0026logout"
                ],
                "summary": "refreshtoken",
                "parameters": [
                    {
                        "type": "string",
                        "default": "th",
                        "description": "(en, th)",
                        "name": "Accept-Language",
                        "in": "header"
                    },
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/refreshTK.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AlComUserModel"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "410": {
                        "description": "Gone",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    }
                }
            }
        },
        "/upload/Upload": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "เฉพาะแบบร่างโครงการ สำหรับอัพโหลดไฟล์ หมายเหตุ : บังคับฟิล project_id ทุกครั้ง (ไม่สามารถ required ผ่าน swagger ได้เนื่องจากติดข้อจำกัด)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data"
                ],
                "summary": "อัพโหลดไฟล์",
                "parameters": [
                    {
                        "type": "string",
                        "description": "(en, th)",
                        "name": "Accept-Language",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "ID โครงการ",
                        "name": "project_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID Item งบประมาณ",
                        "name": "budget_id",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "input file to upload1",
                        "name": "fileupload[0]",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "input file to upload2",
                        "name": "fileupload[1]",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.SwaggerInfoResult"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.SwaggerInfoResult"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "config.SwaggerInfoResult": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "delete.Delete": {
            "type": "object",
            "properties": {
                "config_id": {
                    "type": "string",
                    "example": "0001"
                },
                "group_id": {
                    "type": "string",
                    "example": "group1"
                }
            }
        },
        "deletestatus.DeleteStatus": {
            "type": "object",
            "properties": {
                "config_id": {
                    "type": "string",
                    "example": "0001"
                },
                "group_id": {
                    "type": "string",
                    "example": "group1"
                },
                "log_date": {
                    "type": "string",
                    "example": "2006-01-02"
                },
                "log_time": {
                    "type": "string",
                    "example": "00:00:00"
                },
                "log_user": {
                    "type": "string",
                    "example": "66861"
                },
                "node_status": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "insertData.AddInformationRequest": {
            "type": "object",
            "properties": {
                "config_id": {
                    "type": "string",
                    "example": "0001"
                },
                "create_date": {
                    "type": "string",
                    "example": "2022-01-01"
                },
                "create_time": {
                    "type": "string",
                    "example": "00:00:00"
                },
                "create_user": {
                    "type": "string",
                    "example": "66861"
                },
                "group_id": {
                    "type": "string",
                    "example": "group1"
                },
                "log_date": {
                    "type": "string",
                    "example": "2006-01-02"
                },
                "log_status": {
                    "type": "string",
                    "example": "xxxxxxxxxx"
                },
                "log_time": {
                    "type": "string",
                    "example": "00:00:00"
                },
                "log_user": {
                    "type": "string",
                    "example": "66861"
                },
                "node_desc": {
                    "type": "string",
                    "example": "xxxxxxxxxx"
                },
                "node_name_en": {
                    "type": "string",
                    "example": "en"
                },
                "node_name_th": {
                    "type": "string",
                    "example": "th"
                },
                "node_ref_value": {
                    "type": "string",
                    "example": "xxxxxxxxxx"
                },
                "node_ref_value2": {
                    "type": "string",
                    "example": "xxxxxxxxxx"
                },
                "node_sort": {
                    "type": "integer",
                    "example": 1
                },
                "node_status": {
                    "type": "integer",
                    "example": 1
                },
                "node_url": {
                    "type": "string",
                    "example": "/test/newurl"
                }
            }
        },
        "login.LoginRequest": {
            "type": "object",
            "properties": {
                "encode": {
                    "type": "string",
                    "example": "c433z2w284c2o3w385v3o3i24474e4b4h4"
                }
            }
        },
        "logout.LogoutRequest": {
            "type": "object",
            "properties": {
                "log_user_id": {
                    "type": "string",
                    "example": "xxxxxxxxxx"
                }
            }
        },
        "models.AlComUserModel": {
            "type": "object",
            "properties": {
                "Last_active_ip": {
                    "type": "string"
                },
                "access_add": {
                    "type": "string"
                },
                "access_audit": {
                    "type": "string"
                },
                "access_delete": {
                    "type": "string"
                },
                "access_display": {
                    "type": "string"
                },
                "access_edit": {
                    "type": "string"
                },
                "access_operation": {
                    "type": "string"
                },
                "access_payment": {
                    "type": "string"
                },
                "access_post": {
                    "type": "string"
                },
                "access_quality": {
                    "type": "string"
                },
                "access_report": {
                    "type": "string"
                },
                "access_super_user": {
                    "type": "string"
                },
                "address": {
                    "type": "string"
                },
                "applicant_id": {
                    "type": "string"
                },
                "date_time_ex": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "last_active_date": {
                    "type": "string"
                },
                "last_active_time": {
                    "type": "string"
                },
                "last_chande_pwd": {
                    "type": "string"
                },
                "last_chande_pwd_time": {
                    "type": "string"
                },
                "level_priv": {
                    "type": "string"
                },
                "log_date": {
                    "type": "string"
                },
                "log_status": {
                    "type": "string"
                },
                "log_time": {
                    "type": "string"
                },
                "log_user": {
                    "type": "string"
                },
                "mobile_no": {
                    "type": "string"
                },
                "operate_date": {
                    "type": "string"
                },
                "operate_time": {
                    "type": "string"
                },
                "operate_user": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "pin_user": {
                    "type": "string"
                },
                "pre_name": {
                    "type": "string"
                },
                "sur_name": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                },
                "user_status": {
                    "type": "integer"
                }
            }
        },
        "models.Message": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "refreshTK.LoginRequest": {
            "type": "object",
            "properties": {
                "encode": {
                    "type": "string",
                    "example": "c433z2w284c2o3w385v3o3i24474e4b4h4"
                },
                "user_id": {
                    "type": "string",
                    "example": "66866"
                }
            }
        },
        "selectGroup.SelectGroup": {
            "type": "object",
            "properties": {
                "group_id": {
                    "type": "string",
                    "example": "group1"
                },
                "node_desc": {
                    "type": "string",
                    "example": "XXXXXXXX"
                },
                "node_name_en": {
                    "type": "string",
                    "example": "XXXXXXXX"
                },
                "node_name_th": {
                    "type": "string",
                    "example": "XXXXXXXX"
                },
                "node_ref_value": {
                    "type": "string",
                    "example": "XXXXXXXX"
                },
                "node_ref_value2": {
                    "type": "string",
                    "example": "XXXXXXXX"
                },
                "node_url": {
                    "type": "string",
                    "example": "XXXXXXXX"
                }
            }
        },
        "updateData.UpdateData": {
            "type": "object",
            "properties": {
                "config_id": {
                    "type": "string",
                    "example": "0001"
                },
                "create_date": {
                    "type": "string",
                    "example": "2022-01-01"
                },
                "create_time": {
                    "type": "string",
                    "example": "00:00:00"
                },
                "create_user": {
                    "type": "string",
                    "example": "66861"
                },
                "group_id": {
                    "type": "string",
                    "example": "group1"
                },
                "log_date": {
                    "type": "string",
                    "example": "2006-01-02"
                },
                "log_status": {
                    "type": "string",
                    "example": "newupdate"
                },
                "log_time": {
                    "type": "string",
                    "example": "00:00:00"
                },
                "log_user": {
                    "type": "string",
                    "example": "66861"
                },
                "node_desc": {
                    "type": "string",
                    "example": "newupdate"
                },
                "node_name_en": {
                    "type": "string",
                    "example": "en"
                },
                "node_name_th": {
                    "type": "string",
                    "example": "th"
                },
                "node_ref_value": {
                    "type": "string",
                    "example": "newupdate"
                },
                "node_ref_value2": {
                    "type": "string",
                    "example": "newupdate"
                },
                "node_sort": {
                    "type": "integer",
                    "example": 1
                },
                "node_status": {
                    "type": "integer",
                    "example": 1
                },
                "node_url": {
                    "type": "string",
                    "example": "/test/newurl"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "ใส่ค่า Bearer เว้นวรรคและตามด้วย TOKEN  ex(Bearer ?????????)"
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