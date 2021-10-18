# Takeout

O objetivo deste projeto é usar a seguinte stack:

- GoLang
- MongoDB
- Docker

## Primeiros passos

Certifique-se de ter as ferramentas instaladas:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker-compose](https://docs.docker.com/compose/install/)

O próximo passo é executar a aplicação. Para isso, na raiz do projeto, execute:

```shell
docker-compose up -d
```

Por fim, para executar a api:

```shell
docker exec takeout_api go run .
```

Então, após, as seguintes urls estarão disponíveis para você visualizar os acontecimentos:

- Base URL : localhost:8080

|Verbo|Rota|
|--- |--- |
|GET| [/post](http://localhost:8080/post) |
|GET| [/post/{ id }](http://localhost:8080/post/1)|
|DELETE| [/post/{ id }](http://localhost:8080/post/1)|
|PUT| [/post/{ id }](http://localhost:8080/post/1)|

Exemplos de Corpo de Requisicao

```
{
	"name_user": "Jhon Doe",
	"type_post": "Saude",
	"visibility": "all",
	"case_status" : "closed",
	"image": "https://images.pexels.com/photos/9496721/pexels-photo-9496721.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500",
	"description": "Lorem Ipsum..."
}
```

Campos Requiridos e Validação para Criar

```
	name_user    string `validate:"required,maxlength:30"`
	type_post    string `validate:"required,maxlength:10"`
	Visibility   string `validate:"required,maxlength:10"`
	case_status  string `validate:"required,maxlength:10"`
	image        string `validate:"required,maxlength:255"`
	description  string `validate:"required"`
```

- Erro Recebido caso infligir a Validação

```
{
  "error": "{field} e Requirido"
}

ou

{
  "error": {field} tem que ser menor que {maxlength}
}
```