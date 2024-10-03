package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	var port = os.Getenv("PORT")
	var webStaticContentPath = os.Getenv("WEB_STATIC_CONTENT_PATH")
	var webStatiComponentsPath = os.Getenv("WEB_STATIC_COMPONENTS_PATH")
	var rootHTMLFilename = os.Getenv("ROOT_HTML_FILENAME")

	var router = gin.Default()

	router.Static(webStatiComponentsPath, webStaticContentPath+webStatiComponentsPath)
	router.StaticFile("/", webStaticContentPath+"/"+rootHTMLFilename)

	router.GET("/:filepath", func(c *gin.Context) {
		file := c.Param("filepath")
		if strings.HasSuffix(file, ".js") || strings.HasSuffix(file, ".js.map") {
			c.File(filepath.Join(webStaticContentPath, file))
		} else {
			c.Status(http.StatusNotFound)
		}
	})

	router.GET("/api/tests", getTests)
	router.POST("/api/tests", postTest)

	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}

type test struct {
	TestField1 string `json:"testField1"`
	TestField2 string `json:"testField2"`
}

var testVar = []test{
	{TestField1: "test1", TestField2: "test2"},
	{TestField1: "test3", TestField2: "test4"},
}

func getTests(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, testVar)
}

func postTest(context *gin.Context) {
	var newTest test
	err := context.BindJSON(&newTest)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	testVar = append(testVar, newTest)
	context.IndentedJSON(http.StatusCreated, newTest)
}
