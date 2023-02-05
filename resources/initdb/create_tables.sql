CREATE DATABASE jwtexamples;

CREATE TABLE IF NOT EXISTS users (
  id bigserial PRIMARY KEY,
  email varchar(250) NOT NULL unique,
  password varchar(250) NOT NULL
);
