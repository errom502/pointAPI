create table if not exists token (
    id_user int4 NOT NULL,
    token TEXT NOT NULL,
    date_create TIMESTAMP NOT NULL DEFAULT current_timestamp
    );