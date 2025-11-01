package task

import "fmt"

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	CreatedAt string `json:"created_at"`
}

func Add(title string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	newTask := Task{
		ID:        len(tasks) + 1,
		Title:     title,
		Completed: false,
		CreatedAt: "", // Pode ser preenchido com a data atual se necessário
	}

	tasks = append(tasks, newTask)
	return SaveTasks(tasks)
}

func List() error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("Nenhuma tarefa encontrada")
		return nil
	}
	// Exibe as tarefas, indicando se estão concluídas ou não
	for _, t := range tasks {
		status := "❌ não concluída"
		if t.Completed {
			status = "✅ concluída"
		}
		fmt.Printf("%d - %s [%s]\n", t.ID, t.Title, status)
	}
	return nil
}

func Complete(id int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Completed = true
			return SaveTasks(tasks)
		}
	}
	return fmt.Errorf("tarefa com ID %d não encontrada", id)
}

func Remove(id int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	var newTasks []Task
	found := false // Indica se a tarefa foi encontrada
	for _, t := range tasks {
		if t.ID != id {
			newTasks = append(newTasks, t)
		} else {
			found = true // Marca que a tarefa foi encontrada
		}
	}
	if !found {
		return fmt.Errorf("tarefa com ID %d não encontrada", id)
	}
	return SaveTasks(newTasks)
}
