# KARTKA: Plataforma de E-commerce com Microsserviços para Gestão de Pedidos e Estoque 🚀

Kartka é uma aplicação de e-commerce robusta e escalável, construída com uma arquitetura de microsserviços. O backend é desenvolvido em **Go** e utiliza o **Kafka** para gerenciar mensagens de forma assíncrona, proporcionando uma experiência de compra ágil e confiável. A arquitetura Clean Architecture é adotada para garantir a manutenibilidade e escalabilidade da aplicação. O frontend, desenvolvido com **Vue.js** e **TypeScript**, permite uma interação fluida e dinâmica, oferecendo aos usuários a funcionalidade completa de um carrinho de compras.

## Funcionalidades ✨

- **Cadastro e Gerenciamento de Produtos**: API em Go para criar e gerenciar dados de produtos, com processamento assíncrono via Kafka.
- **Carrinho de Compras Interativo**: O frontend em Vue.js permite que usuários adicionem produtos ao carrinho e prossigam para o checkout.
- **Reserva Temporária de Estoque**: Um microsserviço dedicado reserva o produto durante o processo de checkout, liberando-o se o pagamento não for realizado em 10 minutos.
- **Processamento de Mensagens com Kafka**: Assincronicidade e escalabilidade com mensageria para pedidos e reservas de estoque.
- **Clean Architecture**: Código bem organizado, modular e altamente manutenível.

## Tecnologias Utilizadas 🛠

- **Go (Golang)** Backend API.
- **Kafka** Mensageria assíncrona.
- **Vue.js** e **TypeScript** Frontend reativo e seguro.
- **PostgreSQL** Persistência de dados.
- **Docker** Ambiente de desenvolvimento isolado e consistente.

## Arquitetura Limpa + (DI) Injeção de Dependências 🧩
A aplicação segue o padrão Clean Architecture, com a divisão clara de responsabilidades entre camadas. 
Além disso, a aplicação utiliza Injeção de Dependências (DI) para garantir que as dependências entre os componentes sejam gerenciadas de forma flexível e desacoplada. Isso facilita a manutenção, os testes unitários e a escalabilidade do sistema.


## Passo a Passo 🏃‍♂️

1. **Clonar o Repositório em seu ambiente local:**
```bash
  git clone https://github.com/jalvess021/kartka.git
  cd kartka
```
2. **Copie o arquivo `.env.example` para `.env` para configurar as variáveis de ambiente:**
```bash
  cp .env.example .env
```
3. **Se você estiver utilizando o PostgreSQL localmente, acesse o terminal do banco de dados e crie o banco:**
```bash
  psql -U seu_usuario
```
   Depois, crie o banco de dados:
```sql
  CREATE DATABASE kartka;
```
   Agora, crie a tabela `products`:
```sql
  CREATE TABLE products (
    ID UUID PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Price NUMERIC(10, 2) NOT NULL
);
```
4. **Subir os Containers (Kafka, Zookeeper, PostgreSQL, API e Frontend):**
```bash
  docker-compose up -d
```
5. **Acesse o container do KAFKA e crie o Tópico "product":**
```bash
  kafka-topics --bootstrap-server localhost:9092 --create --topic product --if-not-exists
```
6. **Entre no container da API e execute o servidor:**
```bash
  go run cmd/main.go
```

Após feito o passo a passo, o frontend estará acessível LOCALMENTE em: [http://127.0.0.1:5173](http://127.0.0.1:5173).

## Como Funciona? 🔄

1. **Mensagens Kafka**: A API consome mensagens do Kafka associadas ao tópico `product` e processa os dados.
2. **Backend**: A API armazena os dados processados no PostgreSQL.
3. **Frontend**: O Vue.js exibe os produtos de forma interativa, consumindo a API.

### Desenvolvido por [@jalvess021](https://github.com/jalvess021).
