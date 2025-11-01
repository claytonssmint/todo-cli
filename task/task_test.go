package task_test

import (
	"os"
	"testing"

	"github.com/claytonssmint/todo-cli/task"
)

func TestAdd(t *testing.T) {
	// cria diretório temporário para armazenar o arquivo de tarefas durante o teste
	tempDir := t.TempDir()
	originalPath := task.FilePath

	// muda o caminho do arquivo de tarefas para o diretório temporário
	task.FilePath = tempDir + "/tasks.json"
	defer func() {
		// restaura o caminho original após o teste
		task.FilePath = originalPath
	}()

	t.Run("Adicionar nova tarefa", func(t *testing.T) {
		err := task.Add("Testar aplicação")
		if err != nil {
			t.Errorf("Erro ao adicionar tarefa: %v", err)
		}
	})

	t.Run("Erro ao carregar tarefas", func(t *testing.T) {
		// Simula erro ao carregar tarefas renomeando o arquivo temporário

		badFile := tempDir + "/bad_tasks.json"
		f, err := os.Create(badFile)
		if err != nil {
			t.Fatalf("Erro ao criar arquivo ruim: %v", err)
		}
		f.Close()
		task.FilePath = badFile + "/invalid" // Força erro de leitura

		err = task.Add("Tarefa inválida")
		if err == nil {
			t.Error("Esperava erro ao carregar tarefas, mas não ocorreu")
		}
	})

}

func TestList(t *testing.T) {
	// cria diretório temporário para armazenar o arquivo de tarefas durante o teste
	tempDir := t.TempDir()
	originalPath := task.FilePath

	// muda o caminho do arquivo de tarefas para o diretório temporário
	task.FilePath = tempDir + "/tasks.json"
	defer func() {
		// restaura o caminho original após o teste
		task.FilePath = originalPath
	}()

	t.Run("Listar tarefas vazias", func(t *testing.T) {
		err := task.List()
		if err != nil {
			t.Errorf("Erro ao listar tarefas: %v", err)
		}
	})

	t.Run("Listar tarefas com itens", func(t *testing.T) {
		_ = task.Add("Tarefa 1")
		_ = task.Add("Tarefa 2")
		err := task.List()
		if err != nil {
			t.Errorf("Erro ao listar tarefas: %v", err)
		}
	})

	t.Run("Listar tarefas concluídas", func(t *testing.T) {
		_ = task.Complete(1)
		err := task.List()
		if err != nil {
			t.Errorf("Erro ao listar tarefas: %v", err)
		}
	})

	t.Run("Erro ao carregar tarefas", func(t *testing.T) {
		// Simula erro ao carregar tarefas renomeando o arquivo temporário
		badFile := tempDir + "/bad_tasks.json"
		f, err := os.Create(badFile)
		if err != nil {
			t.Fatalf("Erro ao criar arquivo ruim: %v", err)
		}
		f.Close()
		task.FilePath = badFile + "/invalid" // Força erro de leitura
		err = task.List()
		if err == nil {
			t.Error("Esperava erro ao carregar tarefas, mas não ocorreu")
		}
	})

}

func TestRemove(t *testing.T) {
	// cria diretório temporário para armazenar o arquivo de tarefas durante o teste
	tempDir := t.TempDir()
	originalPath := task.FilePath

	// muda o caminho do arquivo de tarefas para o diretório temporário
	task.FilePath = tempDir + "/tasks.json"
	defer func() {
		// restaura o caminho original após o teste
		task.FilePath = originalPath
	}()

	t.Run("Remover tarefa existente", func(t *testing.T) {
		_ = task.Add("Tarefa para remover")
		err := task.Remove(1)
		if err != nil {
			t.Errorf("Erro ao remover tarefa: %v", err)
		}
	})

	t.Run("Remover tarefa inexistente", func(t *testing.T) {
		err := task.Remove(999)
		if err == nil {
			t.Error("Esperava erro ao remover tarefa inexistente, mas não ocorreu")
		}
	})

	t.Run("Erro ao carregar tarefas", func(t *testing.T) {
		// Simula erro ao carregar tarefas renomeando o arquivo temporário
		badFile := tempDir + "/bad_tasks.json"
		f, err := os.Create(badFile)
		if err != nil {
			t.Fatalf("Erro ao criar arquivo ruim: %v", err)
		}
		f.Close()
		task.FilePath = badFile + "/invalid" // Força erro de leitura
		err = task.Remove(1)
		if err == nil {
			t.Error("Esperava erro ao carregar tarefas, mas não ocorreu")
		}
	})
}
