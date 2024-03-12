# Hipercongo

O código contido neste repositório representa o sistema de simulação, mensageria, visualização (Metabase), servidor web e banco de dados NoSQL ( MongoDB ) de um projeto de **Hiperconectividade para Cidades Inteligentes**. Este projeto foi construído conforme as [golang-standards](https://github.com/golang-standards/project-layout) [^1]. Ademais, também foi implementada, dado requisito de alta escalabilidade definido para o projeto, a arquitetura [hexagonal](https://netflixtechblog.com/ready-for-changes-with-hexagonal-architecture-b315ec967749) [^2]

## Dependências e Serviços

Antes de continuar, é necessário instalar as dependências e criar os serviços listados para a execução dos comandos posteriores. Para isso siga as seguintes instruções:

- Cluster Kafka [Confluent Cloud](https://docs.confluent.io/cloud/current/clusters/create-cluster.html#create-ak-clusters)
- Cluster MQTT [HiveMQ Cloud](https://www.hivemq.com/article/step-by-step-guide-using-hivemq-cloud-starter-iot/)
- Docker engine [Install Docker Engine on Ubuntu
](https://docs.docker.com/engine/install/ubuntu/)
- Build Essential [https://itsfoss.com/build-essential-ubuntu/](https://itsfoss.com/build-essential-ubuntu/)

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
simulation  | 2024/03/12 14:03:36 Published: {"sensor_id":"65f042d4b6b032e6b873d320","data":{"co":10,"co2":488,"mp10":120,"mp25":68,"no2":571,"rad":644},"timestamp":"2024-03-12T14:03:36.65569959Z"}, on topic: sensors
simulation  | 2024/03/12 14:03:36 Published: {"sensor_id":"65f0575382f1be93d94ae2ca","data":{"co":8,"co2":485,"mp10":118,"mp25":59,"no2":577,"rad":641},"timestamp":"2024-03-12T14:03:36.656197113Z"}, on topic: sensors
simulation  | 2024/03/12 14:03:36 Published: {"sensor_id":"65f0503d1c30b0c1222d3b99","data":{"co":6,"co2":488,"mp10":117,"mp25":66,"no2":575,"rad":656},"timestamp":"2024-03-12T14:03:36.656816178Z"}, on topic: sensors
simulation  | 2024/03/12 14:03:36 Published: {"sensor_id":"65f052819aaa09e60a915e1f","data":{"co":10,"co2":504,"mp10":133,"mp25":59,"no2":567,"rad":649},"timestamp":"2024-03-12T14:03:36.657669417Z"}, on topic: sensors
simulation  | 2024/03/12 14:03:36 Published: {"sensor_id":"65f0575382f1be93d94ae2c7","data":{"co":9,"co2":484,"mp10":125,"mp25":61,"no2":562,"rad":638},"timestamp":"2024-03-12T14:03:36.657775884Z"}, on topic: sensors
simulation  | 2024/03/12 14:03:36 Published: {"sensor_id":"65f042d4b6b032e6b873d31f","data":{"co":5,"co2":487,"mp10":126,"mp25":58,"no2":574,"rad":626},"timestamp":"2024-03-12T14:03:36.658284966Z"}, on topic: sensors
simulation  | 2024/03/12 14:03:36 Published: {"sensor_id":"65f052819aaa09e60a915e1e","data":{"co":5,"co2":490,"mp10":134,"mp25":61,"no2":561,"rad":650},"timestamp":"2024-03-12T14:03:36.658663385Z"}, on topic: sensors
simulation  | 2024/03/12 14:03:36 Published: {"sensor_id":"65f041fe2cb67aaec8b1ea69","data":{"co":7,"co2":485,"mp10":128,"mp25":59,"no2":570,"rad":624},"timestamp":"2024-03-12T14:03:36.658784378Z"}, on topic: sensors
simulation  | 2024/03/12 14:03:36 Published: {"sensor_id":"65ef0e8426d3e11550dcc2e6","data":{"co":8,"co2":500,"mp10":125,"mp25":63,"no2":565,"rad":641},"timestamp":"2024-03-12T14:03:36.661013772Z"}, on topic: sensors
simulation  | 2024/03/12 14:03:36 Published: {"sensor_id":"65f0503d1c30b0c1222d3b9c","data":{"co":6,"co2":483,"mp10":129,"mp25":62,"no2":583,"rad":634},"timestamp":"2024-03-12T14:03:36.661013772Z"}, on topic: sensors
simulation  | 2024/03/12 14:03:36 Published: {"sensor_id":"65f041fe2cb67aaec8b1ea6c","data":{"co":8,"co2":517,"mp10":130,"mp25":66,"no2":562,"rad":636},"timestamp":"2024-03-12T14:03:36.661014855Z"}, on topic: sensors
simulation  | 2024/03/12 14:03:36 Published: {"sensor_id":"65f0503d1c30b0c1222d3b98","data":{"co":6,"co2":505,"mp10":116,"mp25":59,"no2":549,"rad":654},"timestamp":"2024-03-12T14:03:36.661306844Z"}, on topic: sensors
simulation  | 2024/03/12 14:03:36 Published: {"sensor_id":"65f0575382f1be93d94ae2c8","data":{"co":8,"co2":513,"mp10":124,"mp25":67,"no2":572,"rad":624},"timestamp":"2024-03-12T14:03:36.66132675Z"}, on topic: sensors
simulation  | 2024/03/12 14:03:36 Published: {"sensor_id":"65f0503d1c30b0c1222d3b9b","data":{"co":6,"co2":494,"mp10":121,"mp25":66,"no2":580,"rad":626},"timestamp":"2024-03-12T14:03:36.6612992Z"}, on topic: sensors
simulation  | 2024/03/12 14:03:36 Published: {"sensor_id":"65f052819aaa09e60a915e1d","data":{"co":8,"co2":512,"mp10":118,"mp25":57,"no2":580,"rad":660},"timestamp":"2024-03-12T14:03:36.661400558Z"}, on topic: sensors
metabase-1  | 2024-03-12 14:03:43,218 INFO metabase.util :: Maximum memory available to JVM: 1.9 GB
metabase-1  | 2024-03-12 14:03:47,354 INFO util.encryption :: Saved credentials encryption is DISABLED for this Metabase instance. 🔓 
metabase-1  |  For more information, see https://metabase.com/docs/latest/operations-guide/encrypting-database-details-at-rest.html
metabase-1  | 2024-03-12 14:03:56,673 INFO driver.impl :: Registered abstract driver :sql  🚚
metabase-1  | 2024-03-12 14:03:56,691 INFO driver.impl :: Registered abstract driver :sql-jdbc (parents: [:sql]) 🚚
metabase-1  | 2024-03-12 14:03:56,699 INFO metabase.util :: Load driver :sql-jdbc took 122.6 ms
metabase-1  | 2024-03-12 14:03:56,701 INFO driver.impl :: Registered driver :h2 (parents: [:sql-jdbc]) 🚚
metabase-1  | 2024-03-12 14:03:56,992 INFO driver.impl :: Registered driver :mysql (parents: [:sql-jdbc]) 🚚
metabase-1  | 2024-03-12 14:03:57,033 INFO driver.impl :: Registered driver :postgres (parents: [:sql-jdbc]) 🚚
metabase-1  | 2024-03-12 14:03:58,955 INFO metabase.core :: 
metabase-1  | Metabase v0.48.8 (a900c85) 
metabase-1  | 
metabase-1  | Copyright © 2024 Metabase, Inc. 
metabase-1  | 
metabase-1  | Metabase Enterprise Edition extensions are NOT PRESENT.
metabase-1  | 2024-03-12 14:03:58,965 INFO metabase.core :: Starting Metabase in STANDALONE mode
metabase-1  | 2024-03-12 14:03:59,029 INFO metabase.server :: Launching Embedded Jetty Webserver with config:
metabase-1  |  {:port 3000, :host "0.0.0.0"}
metabase-1  | 
metabase-1  | 2024-03-12 14:03:59,098 INFO metabase.core :: Starting Metabase version v0.48.8 (a900c85) ...
metabase-1  | 2024-03-12 14:03:59,102 INFO metabase.core :: System info:
metabase-1  |  {"file.encoding" "UTF-8",
metabase-1  |  "java.runtime.name" "OpenJDK Runtime Environment",
metabase-1  |  "java.runtime.version" "11.0.22+7",
metabase-1  |  "java.vendor" "Eclipse Adoptium",
metabase-1  |  "java.vendor.url" "https://adoptium.net/",
metabase-1  |  "java.version" "11.0.22",
metabase-1  |  "java.vm.name" "OpenJDK 64-Bit Server VM",
metabase-1  |  "java.vm.version" "11.0.22+7",
metabase-1  |  "os.name" "Linux",
metabase-1  |  "os.version" "6.5.11-linuxkit",
metabase-1  |  "user.language" "en",
metabase-1  |  "user.timezone" "GMT"}
metabase-1  | 
metabase-1  | 2024-03-12 14:03:59,104 INFO metabase.plugins :: Loading plugins in /plugins...
metabase-1  | 2024-03-12 14:03:59,256 INFO util.files :: Extract file /modules/bigquery-cloud-sdk.metabase-driver.jar -> /plugins/bigquery-cloud-sdk.metabase-driver.jar
metabase-1  | 2024-03-12 14:03:59,410 INFO util.files :: Extract file /modules/redshift.metabase-driver.jar -> /plugins/redshift.metabase-driver.jar
metabase-1  | 2024-03-12 14:03:59,416 INFO util.files :: Extract file /modules/mongo.metabase-driver.jar -> /plugins/mongo.metabase-driver.jar
metabase-1  | 2024-03-12 14:03:59,432 INFO util.files :: Extract file /modules/presto-jdbc.metabase-driver.jar -> /plugins/presto-jdbc.metabase-driver.jar
metabase-1  | 2024-03-12 14:03:59,480 INFO util.files :: Extract file /modules/oracle.metabase-driver.jar -> /plugins/oracle.metabase-driver.jar
metabase-1  | 2024-03-12 14:03:59,483 INFO util.files :: Extract file /modules/sqlite.metabase-driver.jar -> /plugins/sqlite.metabase-driver.jar
metabase-1  | 2024-03-12 14:03:59,507 INFO util.files :: Extract file /modules/sqlserver.metabase-driver.jar -> /plugins/sqlserver.metabase-driver.jar
metabase-1  | 2024-03-12 14:03:59,516 INFO util.files :: Extract file /modules/sparksql.metabase-driver.jar -> /plugins/sparksql.metabase-driver.jar
metabase-1  | 2024-03-12 14:03:59,564 INFO util.files :: Extract file /modules/druid.metabase-driver.jar -> /plugins/druid.metabase-driver.jar
metabase-1  | 2024-03-12 14:03:59,569 INFO util.files :: Extract file /modules/vertica.metabase-driver.jar -> /plugins/vertica.metabase-driver.jar
metabase-1  | 2024-03-12 14:03:59,571 INFO util.files :: Extract file /modules/athena.metabase-driver.jar -> /plugins/athena.metabase-driver.jar
metabase-1  | 2024-03-12 14:03:59,644 INFO util.files :: Extract file /modules/googleanalytics.metabase-driver.jar -> /plugins/googleanalytics.metabase-driver.jar
metabase-1  | 2024-03-12 14:03:59,653 INFO util.files :: Extract file /modules/snowflake.metabase-driver.jar -> /plugins/snowflake.metabase-driver.jar
metabase-1  | 2024-03-12 14:04:00,087 DEBUG plugins.lazy-loaded-driver :: Registering lazy loading driver :snowflake...
metabase-1  | 2024-03-12 14:04:00,089 INFO driver.impl :: Registered driver :snowflake (parents: [:sql-jdbc]) 🚚
metabase-1  | 2024-03-12 14:04:00,102 DEBUG plugins.lazy-loaded-driver :: Registering lazy loading driver :athena...
metabase-1  | 2024-03-12 14:04:00,104 INFO driver.impl :: Registered driver :athena (parents: [:sql-jdbc]) 🚚
metabase-1  | 2024-03-12 14:04:00,118 DEBUG plugins.lazy-loaded-driver :: Registering lazy loading driver :mongo...
metabase-1  | 2024-03-12 14:04:00,119 INFO driver.impl :: Registered driver :mongo  🚚
metabase-1  | 2024-03-12 14:04:00,127 DEBUG plugins.lazy-loaded-driver :: Registering lazy loading driver :redshift...
metabase-1  | 2024-03-12 14:04:00,128 INFO driver.impl :: Registered driver :redshift (parents: [:postgres]) 🚚
metabase-1  | 2024-03-12 14:04:00,161 DEBUG plugins.lazy-loaded-driver :: Registering lazy loading driver :hive-like...
metabase-1  | 2024-03-12 14:04:00,163 INFO driver.impl :: Registered abstract driver :hive-like (parents: [:sql-jdbc]) 🚚
metabase-1  | 2024-03-12 14:04:00,164 DEBUG plugins.lazy-loaded-driver :: Registering lazy loading driver :sparksql...
metabase-1  | 2024-03-12 14:04:00,165 INFO driver.impl :: Registered driver :sparksql (parents: [:hive-like]) 🚚
metabase-1  | 2024-03-12 14:04:00,169 DEBUG plugins.lazy-loaded-driver :: Registering lazy loading driver :sqlite...
metabase-1  | 2024-03-12 14:04:00,170 INFO driver.impl :: Registered driver :sqlite (parents: [:sql-jdbc]) 🚚
metabase-1  | 2024-03-12 14:04:00,175 DEBUG plugins.lazy-loaded-driver :: Registering lazy loading driver :sqlserver...
metabase-1  | 2024-03-12 14:04:00,176 INFO driver.impl :: Registered driver :sqlserver (parents: [:sql-jdbc]) 🚚
metabase-1  | 2024-03-12 14:04:00,184 INFO plugins.dependencies :: Metabase cannot initialize plugin Metabase Oracle Driver due to required dependencies. Metabase requires the Oracle JDBC driver in order to connect to Oracle databases, but we can't ship it as part of Metabase due to licensing restrictions. See https://metabase.com/docs/latest/administration-guide/databases/oracle.html for more details.
metabase-1  | 
metabase-1  | 2024-03-12 14:04:00,187 INFO plugins.dependencies :: Metabase Oracle Driver dependency {:class oracle.jdbc.OracleDriver} satisfied? false
metabase-1  | 2024-03-12 14:04:00,188 INFO plugins.dependencies :: Plugins with unsatisfied deps: ["Metabase Oracle Driver"]
metabase-1  | 2024-03-12 14:04:00,192 INFO plugins.dependencies :: Metabase cannot initialize plugin Metabase Vertica Driver due to required dependencies. Metabase requires the Vertica JDBC driver in order to connect to Vertica databases, but we can't ship it as part of Metabase due to licensing restrictions. See https://metabase.com/docs/latest/administration-guide/databases/vertica.html for more details.
metabase-1  | 
metabase-1  | 2024-03-12 14:04:00,193 INFO plugins.dependencies :: Metabase Vertica Driver dependency {:class com.vertica.jdbc.Driver} satisfied? false
metabase-1  | 2024-03-12 14:04:00,194 INFO plugins.dependencies :: Plugins with unsatisfied deps: ["Metabase Oracle Driver" "Metabase Vertica Driver"]
metabase-1  | 2024-03-12 14:04:00,198 DEBUG plugins.lazy-loaded-driver :: Registering lazy loading driver :druid...
metabase-1  | 2024-03-12 14:04:00,199 INFO driver.impl :: Registered driver :druid  🚚
metabase-1  | 2024-03-12 14:04:00,205 DEBUG plugins.lazy-loaded-driver :: Registering lazy loading driver :googleanalytics...
metabase-1  | 2024-03-12 14:04:00,206 INFO driver.impl :: Registered driver :googleanalytics  🚚
metabase-1  | 2024-03-12 14:04:00,218 DEBUG plugins.lazy-loaded-driver :: Registering lazy loading driver :bigquery-cloud-sdk...
metabase-1  | 2024-03-12 14:04:00,219 INFO driver.impl :: Registered driver :bigquery-cloud-sdk (parents: [:sql]) 🚚
metabase-1  | 2024-03-12 14:04:00,234 DEBUG plugins.lazy-loaded-driver :: Registering lazy loading driver :presto-jdbc...
metabase-1  | 2024-03-12 14:04:00,235 INFO driver.impl :: Registered driver :presto-jdbc (parents: [:sql-jdbc]) 🚚
metabase-1  | 2024-03-12 14:04:00,261 INFO metabase.core :: Setting up and migrating Metabase DB. Please sit tight, this may take a minute...
metabase-1  | 2024-03-12 14:04:00,264 INFO db.setup :: Verifying postgres Database Connection ...
metabase-1  | 2024-03-12 14:04:02,893 INFO db.setup :: Successfully verified PostgreSQL 16.1 application database connection. ✅
metabase-1  | 2024-03-12 14:04:02,894 INFO db.setup :: Checking if a database downgrade is required...
app         | 2024/03/12 14:04:04 Message on sensors_log_queue[4]@4743: {"sensor_id":"65f0503d1c30b0c1222d3b99","data":{"co":9,"co2":492,"mp10":118,"mp25":68,"no2":567,"rad":658},"timestamp":"2024-03-12T13:48:02.880256815Z"}
app         | 2024/03/12 14:04:04 Message on sensors_log_queue[4]@4744: {"sensor_id":"65f042d4b6b032e6b873d321","data":{"co":9,"co2":490,"mp10":126,"mp25":65,"no2":564,"rad":628},"timestamp":"2024-03-12T13:48:02.879628481Z"}
app         | 2024/03/12 14:04:04 Inserting log into the MongoDB collection with id: &{ObjectID("65f060d4ea77e1a1709a9dc5")}
app         | 2024/03/12 14:04:04 Message on sensors_log_queue[4]@4745: {"sensor_id":"65f052819aaa09e60a915e1c","data":{"co":7,"co2":483,"mp10":122,"mp25":58,"no2":569,"rad":658},"timestamp":"2024-03-12T13:48:02.881239788Z"}
app         | 2024/03/12 14:04:04 Inserting log into the MongoDB collection with id: &{ObjectID("65f060d4ea77e1a1709a9dc6")}
app         | 2024/03/12 14:04:04 Message on sensors_log_queue[4]@4746: {"sensor_id":"65ef0e8426d3e11550dcc2e6","data":{"co":8,"co2":500,"mp10":125,"mp25":63,"no2":565,"rad":641},"timestamp":"2024-03-12T13:48:02.879403424Z"}
app         | 2024/03/12 14:04:04 Inserting log into the MongoDB collection with id: &{ObjectID("65f060d4ea77e1a1709a9dc7")}
app         | 2024/03/12 14:04:04 Message on sensors_log_queue[4]@4747: {"sensor_id":"65f052819aaa09e60a915e1d","data":{"co":8,"co2":486,"mp10":125,"mp25":61,"no2":547,"rad":635},"timestamp":"2024-03-12T13:48:02.88726991Z"}
app         | 2024/03/12 14:04:04 Inserting log into the MongoDB collection with id: &{ObjectID("65f060d4ea77e1a1709a9dc8")}
app         | 2024/03/12 14:04:04 Message on sensors_log_queue[4]@4748: {"sensor_id":"65f0575382f1be93d94ae2ca","data":{"co":9,"co2":505,"mp10":128,"mp25":60,"no2":561,"rad":637},"timestamp":"2024-03-12T13:48:02.897612184Z"}
app         | 2024/03/12 14:04:04 Inserting log into the MongoDB collection with id: &{ObjectID("65f060d4ea77e1a1709a9dc9")}
app         | 2024/03/12 14:04:04 Message on sensors_log_queue[4]@4749: {"sensor_id":"65f041fe2cb67aaec8b1ea6a","data":{"co":7,"co2":489,"mp10":118,"mp25":62,"no2":550,"rad":659},"timestamp":"2024-03-12T13:48:02.897691275Z"}
app         | 2024/03/12 14:04:04 Inserting log into the MongoDB collection with id: &{ObjectID("65f060d4ea77e1a1709a9dca")}
app         | 2024/03/12 14:04:04 Message on sensors_log_queue[4]@4750: {"sensor_id":"65f0575382f1be93d94ae2c7","data":{"co":9,"co2":495,"mp10":128,"mp25":59,"no2":573,"rad":635},"timestamp":"2024-03-12T13:48:02.898943864Z"}
app         | 2024/03/12 14:04:04 Inserting log into the MongoDB collection with id: &{ObjectID("65f060d4ea77e1a1709a9dcb")}
app         | 2024/03/12 14:04:04 Message on sensors_log_queue[4]@4751: {"sensor_id":"65f042d4b6b032e6b873d320","data":{"co":9,"co2":502,"mp10":117,"mp25":63,"no2":569,"rad":634},"timestamp":"2024-03-12T13:48:02.89903667Z"}
app         | 2024/03/12 14:04:04 Inserting log into the MongoDB collection with id: &{ObjectID("65f060d4ea77e1a1709a9dcc")}
app         | 2024/03/12 14:04:04 Message on sensors_log_queue[4]@4752: {"sensor_id":"65f0575382f1be93d94ae2c8","data":{"co":8,"co2":514,"mp10":121,"mp25":60,"no2":555,"rad":634},"timestamp":"2024-03-12T13:48:02.901332048Z"}
app         | 2024/03/12 14:04:04 Inserting log into the MongoDB collection with id: &{ObjectID("65f060d4ea77e1a1709a9dcd")}
app         | 2024/03/12 14:04:04 Message on sensors_log_queue[4]@4753: {"sensor_id":"65f052819aaa09e60a915e1f","data":{"co":7,"co2":517,"mp10":119,"mp25":68,"no2":561,"rad":642},"timestamp":"2024-03-12T13:48:02.912210756Z"}
app         | 2024/03/12 14:04:04 Inserting log into the MongoDB collection with id: &{ObjectID("65f060d4ea77e1a1709a9dce")}
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
