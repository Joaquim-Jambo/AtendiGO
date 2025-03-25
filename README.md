# AtendiGO

🟢 **AtendiGO** - Sistema de gerenciamento de filas eficiente e intuitivo em Go (Golang).

## 🚀 Funcionalidades
- ➕ Adiciona pessoas à fila com verificação de duplicidade.
- ⚠️ Alerta automático caso a pessoa já esteja na fila.
- 🟢 Atendimento claro e organizado.
- ❗️ Identifica se não há mais pessoas na fila.

## 🛠️ Como Executar

### 1. Clone o repositório
```bash
git clone <URL_DO_REPOSITORIO>
cd <NOME_DO_REPOSITORIO>
```

### 2. Execute o projeto
Certifique-se de ter o Go instalado.

```bash
go run main.go
```

## 📋 Exemplo de Uso
```bash
João adicionado a fila
Joaquim adicionado a fila
Maria adicionado a fila
Maria já está na fila ⚠️
➡️ Começando o atendimento...
🟢 Atendendo: João
🟢 Atendendo: Joaquim
🟢 Atendendo: Maria
❗️ Não há pessoas na fila para atendimento
```

## 📄 Estrutura do Código
- `adicionarPessoa`: Adiciona uma pessoa à fila com verificação de duplicidade.
- `atenderPessoa`: Atende a próxima pessoa na fila ou informa que não há mais pessoas.

## 📦 Requisitos
- Go 1.21+

## 🤝 Contribuição
Sinta-se à vontade para contribuir com melhorias ou novas funcionalidades. Basta fazer um **fork** do projeto e enviar um **pull request**. 😊

## 📧 Contato
Joaquim Fariti Nacassi Jambo - [joaquimjambo12@gmail.com](mailto:joaquimjambo12@gmail.com)
