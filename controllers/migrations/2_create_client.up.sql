CREATE TABLE IF NOT EXISTS Client (
	id int4 PRIMARY KEY DEFAULT random_between(1000, 9000),
	login varchar(100) NOT NULL,
	"password" varchar(100) NOT NULL
);