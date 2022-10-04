create table if not exists token (
    id_user INT NOT NULL,
    token TEXT DEFAULT (generate_uid(10) || '-' || generate_uid(10) || '-' || generate_uid(10) || '-' || generate_uid(10)),
    date_create TIMESTAMP NOT NULL DEFAULT current_timestamp
    );