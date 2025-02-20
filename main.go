package main

import (
	"errors"
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

func addTodo(context *gin.Context){
	var newTodo todo 
	err := context.BindJSON(&newTodo)
	if err != nil{
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated,newTodo)
}


func getTodo(context *gin.Context){
	id := context.Param("id")
	
	todo, err := getTodoById(id)
	if err == nil{
		context.IndentedJSON(http.StatusOK,todo)
		return
	} else {
		context.IndentedJSON(http.StatusNotFound,gin.H{"message":"Todo not Found"})
		return
	}
}
	
func getTodoById(id string)(*todo, error){
		for i, t := range todos{
			if t.ID == id{
				return &todos[i], nil
			}
		}
		return nil, errors.New("todo not found")
}

func main(){
	todos = append(todos,todo{"1","Clean Room",false})
	todos = append(todos,todo{"2","Wash Clothes",false})
	todos = append(todos,todo{"3","Home Work",false})
	fmt.Println(todos)


	router := gin.Default()

	router.GET("/todos",getTodos)
	router.POST("/todos",addTodo)
	router.GET("/todos/:id",getTodo)

	router.Run("localhost:8080")
}