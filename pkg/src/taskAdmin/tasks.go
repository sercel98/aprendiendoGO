package main

import "fmt"

type taskList struct {
	tasks []*task
}

type task struct {
	name        string
	description string
	completed   bool
}

func (t *taskList) addTaskToList(newTask *task) {
	t.tasks = append(t.tasks, newTask)
}

func (t *taskList) deleteTaskToList(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *taskList) printList() {
	for _, task := range t.tasks {
		fmt.Println("Task:", task.name)
		fmt.Println("Description:", task.description)
	}
}

func (t *taskList) printCompletedTasks() {
	for _, task := range t.tasks {
		if task.completed {
			fmt.Println("Task:", task.name)
			fmt.Println("Description:", task.description)
		}
	}
}

func (t *task) completeTask() {
	t.completed = true
}

func (t *task) updateName(newName string) {
	t.name = newName
}

func (t *task) updateDescription(newDescription string) {
	t.name = newDescription
}

func main() {

	t1 := &task{
		name:        "Manual de usuario de Hippocampus",
		description: "Es necesario realizar el manual de usuario durante el fin de semana",
	}

	t2 := &task{
		name:        "Reunion semanal",
		description: "Martes a las 9:30",
	}

	t3 := &task{
		name:        "Just a test",
		description: "Martes a las 2 pm",
	}

	lista := &taskList{
		tasks: []*task{
			t1, t2, t3,
		},
	}

	lista.tasks[0].completeTask()

	lista.printCompletedTasks()

}
