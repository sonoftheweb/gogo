package main

import (
	"fmt"
	"note/input"
	"note/interfaces"
	"note/note"
	"note/todo"
)

func main() {
	title, content := getNoteData()
	todoText := input.GetInput("Todo Content:")

	todoData, err := todo.New(todoText)
	if handleErrorWithErr(err) {
		return
	}
	err = outputData(todoData)
	if handleErrorWithReturn(err) {
		return
	}

	userNote, err := note.New(title, content)
	if handleErrorWithErr(err) {
		return
	}
	err = outputData(userNote)
	if handleErrorWithReturn(err) {
		return
	}
}

func handleErrorWithReturn(err error) bool {
	if err != nil {
		return true
	}
	return false
}

func handleErrorWithErr(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
		return true
	}
	return false
}

func outputData(data interfaces.Outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data interfaces.Saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving note failed")
		return err
	}
	fmt.Println("Note saved successfully")
	return nil
}

func getNoteData() (string, string) {
	title := input.GetInput("Note Title:")
	content := input.GetInput("Note Content:")

	return title, content
}
