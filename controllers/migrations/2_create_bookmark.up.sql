drop table if exists Bookmarks;
create table Bookmarks (
	id serial primary key,
	name text not null,
	address text not null,
	owner int not null REFERENCES client(id),
	info text default '--'
);