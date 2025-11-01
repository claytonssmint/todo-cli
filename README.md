# todo-cli

CLI simples para gerenciar tarefas (todo) escrito em Go.

## Descrição
Aplicativo de linha de comando para adicionar, listar, marcar como concluída e remover tarefas. As tarefas são persistidas em um arquivo JSON.

## Requisitos
- Go 1.20+ (ou versão compatível)

## Instalação
Do diretório do projeto:
```bash
go build -o todo ./   # compila binário `todo`
# ou para instalar no GOPATH/bin
go install ./...
```

## Uso
Sintaxe:
```
todo <comando> [argumentos]
```

Comandos:
- `todo add "<título da tarefa>"` — adiciona nova tarefa (use aspas para títulos com espaços).
- `todo list` — lista todas as tarefas.
- `todo done <ID>` — marca tarefa como concluída.
- `todo remove <ID>` — remove tarefa.

Exemplos:
```bash
./todo add "Comprar leite"
./todo list
./todo done 1
./todo remove 1
```

## Testes
Rode os testes do pacote `task`:
```bash
go test ./task -v
```
Se os testes dependerem do arquivo de dados, crie o diretório/arquivo inicial:
```bash
mkdir -p data
echo "[]" > data/tasks.json
```

## Notas de desenvolvimento
- Exemplo de criar uma tarefa via terminal, go run main.go add "Estudar Golang 
# todo-cli
