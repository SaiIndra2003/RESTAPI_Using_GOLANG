package main

import (
	"fmt"
)


type todo struct{
	ID string
	Item string
	Completed bool
}


func main(){
	var todos = make([]todo,0)
	todos = append(todos,todo{"1","Clean Room",false})
	todos = append(todos,todo{"2","Wash Clothes",false})
	todos = append(todos,todo{"3","Home Work",false})
	fmt.Println(todos)
}