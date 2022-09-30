create table if not exists Client (
	id TEXT PRIMARY KEY DEFAULT ('c_' || generate_uid(10)),
	login varchar(100) NOT NULL,
	password varchar(100) NOT NULL
);