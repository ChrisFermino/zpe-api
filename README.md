# Zpe API

## English

The version of Golang used in this project was [1.20.4](https://go.dev/dl/).

### Database Creation
The database used was Postgres, so I used Docker to generate a container for this image. The database configurations are in the `docker-compose.yaml` file.

Links to install Docker on [Windows](https://docs.docker.com/desktop/install/windows-install/), [Linux](https://docs.docker.com/desktop/install/linux-install/), or [Mac](https://docs.docker.com/desktop/install/mac-install/).

To start the Postgres database in Docker, navigate to the project's root directory, open the `CMD` or `Bash`, and execute the following command:

`docker compose up -d`

### Running the Project
With the database set up, to run the application, navigate to the root folder of the project, open a CMD or Bash, and execute the command `go get ./...` to install all project dependencies.
Then you can execute the command `go build main.go`, which will generate an executable of the program. Just run it, and the program will be running on port :8080.

### Using Swagger Documentation

To access the documentation of the project, Swagger is configured in the application, the link below will open the respective documentation.
http://localhost:8080/swagger/index.html

To use the Swagger documentation, access the link above and remember to log in and enter the generated token in the "Authorize" button along with the word "Bearer " (with a space) in front of the token.

## Português

A versão do Golang utilizada nesse projeto foi a [1.20.4](https://go.dev/dl/)

### Criação do Banco de dados

O banco de dados utilizado foi o Postgres, então utilizei o docker para gerar um container para esta imagem.
As configurações do banco de dados estão no arquivo `docker-compose.yaml`.

Links para instalar o docker no [Windows](https://docs.docker.com/desktop/install/windows-install/) ou [Linux](https://docs.docker.com/desktop/install/linux-install/) ou [Mac](https://docs.docker.com/desktop/install/mac-install/)

Para subir a base do Postres no docker acesse a raiz do projeto abra o `CMD` ou o `Bash` e execute o comando abaixo

`docker compose up -d`

### Rodando o Projeto

Com o banco de dados pronto, para rodar a aplicação acesse a pasta raiz do projeto abra um `CMD` ou `Bash` e execute o comando `go get ./...` para instalar todas as dependências do projeto.
Em seguida pode executar o comando `go build main.go`, isso irá gerar um executável do programa, só executar ele e o programa está rodando na porta `:8080`

### Utilização da documentação com Swagger

Para acessar a documentação do projeto está configurado o Swagger na aplicação, o link abaixo irá abrir a respectiva documentação.

http://localhost:8080/swagger/index.html

para usar a documentação do Swagger acesse o link acima e lembresse de fazer o login e colocar o token gerado no botão "Authorize" juntamente com a palavra "Bearer "(com espaço) na frente do token