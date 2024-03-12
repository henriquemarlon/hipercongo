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

Novamente, todos os comandos necessários estão sendo abstraídos por um arquivo Makefile. Se você tiver curiosidade para saber o que o comando abaixo faz, basta conferir [aqui]().

#### Comando:

```bash
make coverage 
```

#### Output:


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
![sensor_entity](https://github.com/henriquemarlon/hipercongo/assets/89201795/e64f1191-0d1e-4ab0-b651-39264bd21090)

### Mensageria:
![kafka](https://github.com/henriquemarlon/hipercongo/assets/89201795/c9da37fa-896c-40d5-9199-2ee0d986400d)

### Servidor WEB:
![web](https://github.com/henriquemarlon/hipercongo/assets/89201795/4b15e627-dc83-413a-aa4a-a5e2e8abffce)

### Banco de dados:

![sensor_log](https://github.com/henriquemarlon/hipercongo/assets/89201795/c33bf0c5-4e2d-4084-a7b0-50d1fe8441da)

![sensor](https://github.com/henriquemarlon/hipercongo/assets/89201795/9de0254f-45cb-40b1-8b61-1844f28f5c98)

### Visualização:

## Desenvolvimento orientado a testes

### Testes unitários:

### Testes de integração:

## Demonstração do sistema

[^1]: A estrutura de pastas escolhida para este projeto está de acordo com as convenções e padrões utilizados pela comunidade de desenvolvedores Golang.

[^2]: As entidades, repositórios e usecases estão de acordo com os padrões previstos para a arquitetura hexagonal.
