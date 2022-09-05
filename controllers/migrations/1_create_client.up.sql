create table if not exists Client (
	id SERIAL PRIMARY KEY,
	login varchar(100) NOT NULL,
	password varchar(100) NOT NULL
);