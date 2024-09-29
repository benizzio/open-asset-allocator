package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {

	var port = os.Getenv("PORT")
	var rootHTMLPath = os.Getenv("ROOT_HTML_PATH")

	var router = gin.Default()

	//router.Static("/static", "../web-static")
	router.StaticFile("/", rootHTMLPath)

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
