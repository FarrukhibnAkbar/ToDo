drop table if exists users cascade;

create table if not exists users (
	id serial not null primary key,
	first_name varchar(50),
	last_name varchar(50)
);