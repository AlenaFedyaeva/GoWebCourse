// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Alena Fedyaeva",
            "email": "placeholder@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "show all posts in blog",
                "produces": [
                    "text/html"
                ],
                "summary": "PagePostList"
            }
        },
        "/create": {
            "get": {
                "description": "show page for new post",
                "produces": [
                    "text/html"
                ],
                "summary": "PageNewPost"
            },
            "post": {
                "description": "update posts \u0026 redirect to /",
                "produces": [
                    "text/html"
                ],
                "summary": "FormPageNewPost"
            }
        },
        "/delete/:id": {
            "post": {
                "description": "delete post and redirect to /",
                "produces": [
                    "text/html"
                ],
                "summary": "ButtonDeletePost"
            }
        },
        "/edit/:id": {
            "get": {
                "description": "show page for new post",
                "produces": [
                    "text/html"
                ],
                "summary": "PageUpdatePost"
            },
            "put": {
                "description": "show page for new post",
                "produces": [
                    "text/html"
                ],
                "summary": "FormUpdatePost"
            }
        },
        "/post/:id": {
            "get": {
                "description": "render one post with id",
                "produces": [
                    "text/html"
                ],
                "summary": "PageGetPost",
                "parameters": [
                    {
                        "type": "string",
                        "description": "task list id",
                        "name": "id",
                        "in": "query"
                    }
                ]
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Posts / my blog",
	Description: "This is blog",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
