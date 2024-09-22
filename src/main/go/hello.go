package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()
	router.Delims("{[{", "}]}")

	router.GET("/tests", getTests)
	router.POST("/tests", postTest)
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "test.html", gin.H{
			"title": "Main website",
		})
	})

	err := router.Run(":8080")
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
