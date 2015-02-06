package main

import "fmt"

// Todo .
type Todo struct {
	text      string
	completed bool
}

func (todo *Todo) complete() {
	todo.completed = true
}

// TodoList .
type TodoList struct {
	idIncrement int
	items       map[int]*Todo
}

func (tl *TodoList) addItem(text string) {
	todo := &Todo{text, false}

	if tl.items == nil {
		tl.idIncrement = 0
		tl.items = make(map[int]*Todo)
	}

	tl.idIncrement++
	tl.items[tl.idIncrement] = todo
}

func (tl *TodoList) getItem(id int) *Todo {
	todo := tl.items[id]

	return todo
}

func main() {
	todoList := TodoList{}
	todoList.addItem("Make coffee")
	todoList.addItem("Drink coffee")

	todoList.getItem(1).complete()
	todoList.getItem(2).complete()

	for _, todo := range todoList.items {
		fmt.Println(todo)
	}
}
