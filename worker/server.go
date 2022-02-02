package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rwestlund/gotex"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type tplate int16

const (
	File tplate = iota
	Book
	Pdf
)

type data struct {
	Name     string `json:"name"`
	ID       string `json:"id"`
	Template tplate `json:"template"`
}

func getFile(c *gin.Context) {
	name := c.Param("name")

	path := "archive/" + strings.ToLower(name)

	if _, err := os.Stat(path + ".txt"); err == nil {
		c.File(path + ".txt")
	} else {
		if _, errPdf := os.Stat(path + ".pdf"); errPdf == nil {
			c.File(path + ".pdf")
		} else {
			c.Status(http.StatusNotFound)
		}
	}
}

func createFile(c *gin.Context) {
	var reqData data

	if err := c.BindJSON(&reqData); err != nil || reqData.Name == "" || reqData.ID == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	templateContent := readTemplate(reqData.Template)

	t := template.Must(template.New("gen").Parse(templateContent))
	buf := &bytes.Buffer{}
	if templateErr := t.Execute(buf, reqData); templateErr != nil {
		c.String(http.StatusInternalServerError, templateErr.Error())
		return
	}

	path := "archive/" + strings.ToLower(reqData.Name)

	if reqData.Template == Pdf {
		path += ".pdf"

		doc := buf.String()

		doc = strings.ReplaceAll(doc, "#(#", "{")
		doc = strings.ReplaceAll(doc, "#)#", "}")

		fmt.Println(doc)

		pdf, pdfErr := gotex.Render(doc, gotex.Options{
			Runs: 1,
		})
		if pdfErr != nil {
			log.Println("render failed ", pdfErr)
		}

		writePdfErr := ioutil.WriteFile(path, pdf, 0644)
		if writePdfErr != nil {
			c.String(http.StatusInternalServerError, writePdfErr.Error())
			return
		}
	} else {
		path += ".txt"
		writeErr := ioutil.WriteFile(path, buf.Bytes(), 0644)
		if writeErr != nil {
			c.String(http.StatusInternalServerError, writeErr.Error())
			return
		}
	}

	c.Status(http.StatusOK)
}

func readTemplate(t tplate) string {
	var temp string

	switch t {
	case File:
		temp = "file.template"
		break
	case Book:
		temp = "book.template"
		break
	case Pdf:
		temp = "pdf.template"
	default:
		return ""
	}

	content, fileErr := ioutil.ReadFile("template/" + temp)
	if fileErr != nil {
		log.Fatal(fileErr)
	}

	out := string(content)

	if t == Pdf {
		out = strings.ReplaceAll(out, "{", "#(#")
		out = strings.ReplaceAll(out, "}", "#)#")
		out = strings.ReplaceAll(out, "[[", "{{")
		out = strings.ReplaceAll(out, "]]", "}}")
	}

	return out
}

func main() {
	router := gin.Default()

	router.GET("/archive/:name", getFile)

	router.POST("/archive", createFile)

	err := router.Run("0.0.0.0:8080")
	if err != nil {
		return
	}
}
