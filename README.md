# Hipercongo

O código contido neste repositório representa o sistema de simulação, mensageria, visualização (Metabase), servidor web e banco de dados NoSQL ( MongoDB ) de um projeto de **Hiperconectividade para Cidades Inteligentes**. Este projeto foi construído conforme as [golang-standards](https://github.com/golang-standards/project-layout) [^1]. Ademais, também foi implementada, dado requisito de alta escalabilidade definido para o projeto, a arquitetura [hexagonal](https://netflixtechblog.com/ready-for-changes-with-hexagonal-architecture-b315ec967749) [^2]

## Dependências e Serviços

Antes de continuar, é necessário instalar as dependências e criar os serviços listados para a execução dos comandos posteriores. Para isso siga as seguintes instruções:

- ee
- ee
- ee
- ee

## Como rodar o sistema

Siga as intruções abaixo para rodar o sistema junto a todos os seus recortes, simulação, mensageria, banco de dados e vicualização com o Metabase.

### Definir as variáveis de ambiente:

#### Comando:
```shell
make env
```

#### Output:
```shell

```

> [!NOTE]
> - Note

### Rodar as migrações:

#### Comando:
```shell
make migrations
```

#### Output:
```shell

```

> [!NOTE]
> - Note

### Rodar testes:

Aqui, todos os comandos necessários estão sendo abstraídos por um arquivo Makefile. Se você tiver curiosidade para saber o que o comando abaixo faz, basta conferir [aqui]().

#### Comando:

```shell
make tests
```

#### Output:

```shell

```

> [!NOTE]
> - Note

### Rodar a visualização da cobertura de testes:

Novamente, todos os comandos necessários estão sendo abstraídos por um arquivo Makefile. Se você tiver curiosidade para saber o que o comando abaixo faz, basta conferir [aqui](https://github.com/Inteli-College/2024-T0002-EC09-G04/blob/main/backend/Makefile#L21).

#### Comando:

```bash
make coverage 
```

#### Output:
```shell

```

> [!NOTE]
>  - Este comando está criando, a partir do arquivo `coverage_sheet.md`, uma visualização da cobertura de testes nos principais arquivos Go.

### Rodar o sistema:

Mais uma vez, todos os comandos necessários estão sendo abstraídos por um arquivo Makefile. Se você tiver curiosidade para saber o que o comando abaixo faz, basta conferir [aqui]().

#### Comando:

```bash
make run
```

#### Output:

```shell

```

> [!NOTE]
>  - Note

## Características do sistema

### Simulação:

### Mensageria:

### Servidor WEB:

### Banco de dados:

### Visualização:

## Desenvolvimento orientado a testes

### Testes unitários:

### Testes de integração:

## Demonstração do sistema

[^1]: A estrutura de pastas escolhida para este projeto está de acordo com as convenções e padrões utilizados pela comunidade de desenvolvedores Golang.

[^2]: As entidades, repositórios e usecases estão de acordo com os padrões previstos para a arquitetura hexagonal.
