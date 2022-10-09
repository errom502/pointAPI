CREATE TABLE IF NOT EXISTS Client (
	id int4 PRIMARY KEY,
	login varchar(100) NOT NULL,
	"password" varchar(100) NOT NULL
);