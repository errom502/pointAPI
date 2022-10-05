create table if not exists token (
	id_user TEXT NOT NULL,
	token TEXT, 
	date_create TIMESTAMP NOT NULL DEFAULT current_timestamp
);