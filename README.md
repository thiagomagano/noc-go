# Separador de Times

## Descrição

Este programa em Go foi desenvolvido para criar times equilibrados automaticamente. Ele recebe uma lista de jogadores com seus respectivos níveis de habilidade e os divide em dois times balanceados, garantindo que ambos tenham um nível total de habilidade similar e quantidade de jogadores aproximadamente igual.

## Funcionalidades

- Carrega jogadores de um arquivo JSON
- Divide jogadores em dois times equilibrados com base em seus níveis de habilidade
- Salva os times resultantes em formato JSON ou Markdown
- Adiciona timestamp no nome do arquivo de saída para fácil rastreamento

## Requisitos

- Go 1.23.4 ou superior

## Instalação

Clone este repositório:

```bash
git clone https://github.com/thiagomagano/go/separador-de-times.git
cd separador-de-times
```

## Como Usar

### Executar com configurações padrão:

```bash
go run main.go
```

Isso carregará os jogadores do arquivo `exemplo_jogadores.json` e salvará os times em formato JSON.

### Opções de linha de comando:

```bash
go run main.go -arquivo [caminho_do_arquivo] -formato [formato]
```

Onde:
- `-arquivo`: Caminho para o arquivo JSON com os dados dos jogadores (padrão: "exemplo_jogadores.json")
- `-formato`: Formato de saída, "json" ou "md" (padrão: "json")

### Exemplo:

```bash
go run main.go -arquivo meus_jogadores.json -formato md
```

## Formato do Arquivo de Entrada

O arquivo de entrada deve ser um JSON com a seguinte estrutura:

```json
[
    {"nome": "Jogador1", "nivel": 5},
    {"nome": "Jogador2", "nivel": 3},
    {"nome": "Jogador3", "nivel": 4}
]
```

## Saída

### JSON

O programa salva os times em um arquivo JSON com a seguinte estrutura:

```json
{
  "time1": {
    "jogadores": [
      {"nome": "Jogador1", "nivel": 5},
      {"nome": "Jogador3", "nivel": 3}
    ],
    "nivel_total": 8,
    "total_jogadores": 2
  },
  "time2": {
    "jogadores": [
      {"nome": "Jogador2", "nivel": 3},
      {"nome": "Jogador4", "nivel": 5}
    ],
    "nivel_total": 8,
    "total_jogadores": 2
  }
}
```

### Markdown

Em formato Markdown, a saída terá a seguinte estrutura:

```markdown
# *Atenção aos times!*


*Time Branco(2)*:
- Jogador1
- Jogador3

*Time Preto (2)*:
- Jogador2
- Jogador4
```

## Estrutura do Projeto

- `main.go`: Código principal do programa
- `exemplo_jogadores.json`: Exemplo de arquivo de entrada
- `times/`: Diretório onde os arquivos de saída são salvos
