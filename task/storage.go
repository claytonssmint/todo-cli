package task

import (
	"encoding/json"
	"errors"
	"os"
)

var FilePath = "data/tasks.json"

func LoadTasks() ([]Task, error) {
	// Implementação para carregar tarefas do arquivo JSON
	file, err := os.ReadFile(FilePath)
	if errors.Is(err, os.ErrNotExist) {
		return []Task{}, nil // Retorna lista vazia se o arquivo não existir
	}
	if err != nil {
		return nil, err
	}

	var tasks []Task // serve para armazenar as tarefas carregadas
	if err := json.Unmarshal(file, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func SaveTasks(tasks []Task) error {
	// Implementação para salvar tarefas no arquivo JSON
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(FilePath, data, 0644)
}
