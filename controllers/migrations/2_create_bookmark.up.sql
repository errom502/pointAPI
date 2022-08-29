drop table if exists bookmarks;
drop table if exists  client;
create table bookmarks (
	id serial primary key,
	name text not null,
	address text not null,
	owner int not null REFERENCES client(id),
	info text default '--'
);