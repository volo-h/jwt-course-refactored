create database jwtexamples;

CREATE TABLE IF NOT EXISTS books (
  id bigserial PRIMARY KEY,
  title varchar(250) NOT NULL,
  author varchar(250) NOT NULL,
  year varchar(250) NOT NULL
);

create table users (
  id serial primary key,
  email text not null unique,
  password text not null
);
