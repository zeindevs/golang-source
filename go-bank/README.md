# gobank

<div align="center">
	
### A complete JSON API project in Golang where we are building a bank API

<p>
`make run` (for now ^^)`
</p>
</div>

#### Login

`POST /login`

```sh
http POST localhost:3000/login number:=767441 password=12345678 --json
```

#### Get List Account

`GET /account`

```sh
http GET localhost:3000/account
```

#### Create Account

`POST /account`

```sh
http POST localhost:3000/account firstName=zee lastName=dev password=12345678 --json
```

#### Get Account

`GET /account/2`

```sh
http GET localhost:3000/account/2 x-jwt-token:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50TnVtYmVyIjo3Njc0NDEsImV4cGlyZXNBdCI6MTUwMDB9.V4axK5hOWBHR1hNltDau3mYHMkM0oqad7Ut3WGVo6KM"
```

## License

Copyright (c) 2024 zeindevs. Licensed under the MIT License (MIT).
