CREATE TABLE IF NOT EXISTS user_account (
	id serial,
	email varchar(200) unique,
	password varchar(255),
	address text,
	PRIMARY KEY (id)
);