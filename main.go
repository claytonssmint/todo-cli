package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/claytonssmint/todo-cli/task"
)

func main() {
	// Esse código gerencia a interface de linha de comando para o aplicativo de tarefas
	if len(os.Args) < 2 {
		fmt.Println("Uso: todo <comando> [argumentos]")
		fmt.Println("Comandos disponíveis: add, list, done, remove")
		return
	}

	// Obtém o comando fornecido pelo usuário
	command := os.Args[1]

	// Executa a ação correspondente ao comando
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Uso: todo add <título da tarefa>")
			return
		}
		title := os.Args[2]
		if err := task.Add(title); err != nil {
			fmt.Println("Erro ao adicionar tarefa:", err)
		} else {
			fmt.Println("✅ Tarefa adicionada com sucesso!")
		}

	case "list":
		_ = task.List()

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Uso: todo done <ID da tarefa>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		if err := task.Complete(id); err != nil {
			fmt.Println("Erro ao marcar tarefa como concluída:", err)
		} else {
			fmt.Println("✅ Tarefa marcada como concluída!")
		}

	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Uso: todo remove <ID da tarefa>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		if err := task.Remove(id); err != nil {
			fmt.Println("Erro ao remover tarefa:", err)
		} else {
			fmt.Println("✅ Tarefa removida com sucesso!")
		}

	default:
		fmt.Println("Comando desconhecido:", command)
	}
}
