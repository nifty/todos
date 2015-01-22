package main

import "fmt"

var todos = make(map[int]string)
var id = 0

func main() {
	addItem("Make coffee")
	addItem("Drink coffee")

	fmt.Println(todos)

	removeItem(1)
	addItem("Clean cup")

	fmt.Println(todos)
}

func addItem(item string) {
	id += 1
	todos[id] = item
}

func removeItem(id int) {
	delete(todos, id)
}
