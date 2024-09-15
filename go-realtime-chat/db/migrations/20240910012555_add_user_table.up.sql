CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "username" VARCHAR UNIQUE NOT NULL,
  "email" VARCHAR UNIQUE NOT NULL,
  "password" VARCHAR NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT NOW()
);
