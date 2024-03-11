# Hipercongo

O código contido neste repositório representa o sistema de simulação, mensageria, visualização (Metabase), servidor web e banco de dados NoSQL ( MongoDB ) de um projeto de **Hiperconectividade para Cidades Inteligentes**. Este projeto foi construído conforme as [golang-standards](https://github.com/golang-standards/project-layout) [^1]. Ademais, também foi implementada, dado requisito de alta escalabilidade definido para o projeto, a arquitetura [hexagonal](https://netflixtechblog.com/ready-for-changes-with-hexagonal-architecture-b315ec967749) [^2]

## Requisitos:

Antes de continuar, é necessário instalar as dependências e criar os serviços listados para a execução dos comandos posteriores. Para isso siga as seguintes instruções:

- ee
- ee
- ee
- ee

## Como rodar o sistema:

Abaixo estão as possíveis interações e as instruções de como realizá-las.

#### Definir as variáveis de ambiente

###### Comando:
```shell

```

###### Output:
```shell

```

> [!NOTE]
>

#### Rodar as migrations:

###### Comando:
```shell

```

###### Output:
```shell

```

> [!NOTE]
>

#### Rodar testes:

Aqui, todos os comandos necessários estão sendo abstraídos por um arquivo Makefile. Se você tiver curiosidade para saber o que o comando abaixo faz, basta conferir [aqui]().

###### Comando:

```shell
make test
```

###### Output:

```shell

```

> [!NOTE]
> - No meio do processo, é necessário subir um broker local para realizar os testes de transmissão de mensagens entre os tópicos.

#### Rodar a visualização da cobertura de testes:

Novamente, todos os comandos necessários estão sendo abstraídos por um arquivo Makefile. Se você tiver curiosidade para saber o que o comando abaixo faz, basta conferir [aqui](https://github.com/Inteli-College/2024-T0002-EC09-G04/blob/main/backend/Makefile#L21).

###### Comando:

```bash
make coverage 
```

###### Output:
```shell

```

> [!NOTE]
>  - Este comando está criando, a partir do arquivo `coverage_sheet.md`, uma visualização da cobertura de testes nos principais arquivos Go.

#### Rodar o sistema:

Mais uma vez, todos os comandos necessários estão sendo abstraídos por um arquivo Makefile. Se você tiver curiosidade para saber o que o comando abaixo faz, basta conferir [aqui]().

###### Comando:

```bash
make run
```

###### Output:

```shell

```

> [!NOTE]
>  - Este comando está subindo todos os serviços presentes no arquivo compose.yml. São eles, o broker local, a simulação e a api-test que está sendo usada, por hora apenas para mostrar o log do que está sendo transmitido pela simulação.

[^1]: A estrutura de pastas escolhida para este projeto está de acordo com as convenções e padrões utilizados pela comunidade de desenvolvedores Golang.

[^2]: As entidades, repositórios e usecases estão de acordo com os padrões previstos para a arquitetura hexagonal.
