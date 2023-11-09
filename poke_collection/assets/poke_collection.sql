create database poke_collection;
create table users (id bigserial, name varchar, password varchar, email varchar);
create table pokemon (id bigserial, user_id bigint, pokemon_name varchar);