
drop table if exists Bookmarks;
drop table if exists  client;
create table client (
    id SERIAL PRIMARY KEY,
    login varchar(100) NOT NULL,
    password varchar(100) NOT NULL
);