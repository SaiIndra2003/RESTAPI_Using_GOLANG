package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)


type todo struct{
	ID string `json:"id"` // define litereals to convert in json format
	Item string `json:"title"`
	Completed bool `json:"completed"`
}

var todos = make([]todo,0)

func getTodos(context *gin.Context){ //gin.context contains information regarding http req
	context.IndentedJSON(http.StatusOK,todos) // converts our slice to json using the indentaions we definedd at struct
}




func main(){
	todos = append(todos,todo{"1","Clean Room",false})
	todos = append(todos,todo{"2","Wash Clothes",false})
	todos = append(todos,todo{"3","Home Work",false})
	fmt.Println(todos)


	router := gin.Default()

	router.GET("/todos",getTodos)

	router.Run("localhost:8080")
}