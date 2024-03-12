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
cp ./config/.env.develop.tmpl ./config/.env
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
migrations  | Connection established successfully
migrations  | Documents inserted. IDs: [ObjectID("65f0575382f1be93d94ae2c6") ObjectID("65f0575382f1be93d94ae2c7") ObjectID("65f0575382f1be93d94ae2c8") ObjectID("65f0575382f1be93d94ae2c9") ObjectID("65f0575382f1be93d94ae2ca")]
migrations  | Connection to MongoDB closed.
migrations exited with code 0
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
✔ Container simulation  Started                                                                                                      0.0s 
?       github.com/henriquemarlon/hipercongo/cmd/app    [no test files]
?       github.com/henriquemarlon/hipercongo/cmd/simulation     [no test files]
?       github.com/henriquemarlon/hipercongo/internal/infra/kafka       [no test files]
?       github.com/henriquemarlon/hipercongo/internal/infra/repository  [no test files]
?       github.com/henriquemarlon/hipercongo/internal/infra/web [no test files]
?       github.com/henriquemarlon/hipercongo/internal/usecase   [no test files]
?       github.com/henriquemarlon/hipercongo/tools      [no test files]
=== RUN   TestNewAlert
--- PASS: TestNewAlert (0.00s)
=== RUN   TestNewLog
--- PASS: TestNewLog (0.00s)
=== RUN   TestEntropy
--- PASS: TestEntropy (0.00s)
=== RUN   TestNewSensor
--- PASS: TestNewSensor (0.00s)
=== RUN   TestNewSensorPayload
--- PASS: TestNewSensorPayload (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/henriquemarlon/hipercongo/internal/domain/entity     0.005s  coverage: 100.0% of statements
=== RUN   TestMqttIntegration
2024/03/12 10:33:48 Selecting all sensors from the MongoDB collection sensors
--- PASS: TestMqttIntegration (152.76s)
PASS
coverage: [no statements]
ok      github.com/henriquemarlon/hipercongo/test       152.771s        coverage: [no statements]
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
![output_coverage](https://github.com/henriquemarlon/hipercongo/assets/89201795/95767c00-44dd-4852-9d63-956e9947c4c6)

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
![sensor_unit](https://github.com/henriquemarlon/hipercongo/assets/89201795/52d2e87a-6d7b-4dea-9232-2b93309180c9)
![log_unit](https://github.com/henriquemarlon/hipercongo/assets/89201795/896254be-cf8a-4883-93ba-5e5172d97ba1)
![alert_unit](https://github.com/henriquemarlon/hipercongo/assets/89201795/45be50cf-0fcc-487e-b15b-83736ada0307)

### Testes de integração:

## Demonstração do sistema

[^1]: A estrutura de pastas escolhida para este projeto está de acordo com as convenções e padrões utilizados pela comunidade de desenvolvedores Golang.

[^2]: As entidades, repositórios e usecases estão de acordo com os padrões previstos para a arquitetura hexagonal.
