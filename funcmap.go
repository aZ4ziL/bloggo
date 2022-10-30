package main

import (
	"bytes"
	"log"

	"github.com/aZ4ziL/bloggo/models"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func markdown(content string) string {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	var buf bytes.Buffer
	if err := md.Convert([]byte(content), &buf); err != nil {
		log.Println(err.Error())
		return ""
	}
	return buf.String()
}

// Length return length of array
func lengthComment(objs []interface{}) int {
	var result int
	for range objs {
		result += 1
	}
	return result
}

// GetUserByID
func getFullNameByID(id uint) string {
	user, _ := models.GetUserByID(id)
	return user.FirstName + " " + user.LastName
}
