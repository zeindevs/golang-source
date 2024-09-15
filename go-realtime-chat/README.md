# go-realtime-chat

## Aproach

- postgres image from docker hub
    - create the go-chat database
- database setup
    - make a connection to the db
    - add a db migration file to create the `users` table
- `/signup` endpoint to create a new user
    - repository <- service <- handler (dependencies)
- `/login` & `/logout` endpoints
    - jwt with http-only cookies
        - prone to csrf attacks
    - vs. token-based authentication
        - prone to xss attacks
    - best: short-lived access token + refresh token

## Docker postgres

```sh
docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine
```

