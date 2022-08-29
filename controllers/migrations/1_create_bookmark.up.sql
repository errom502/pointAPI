create table Bookmarks (
	id serial primary key,
	name text not null,
	address text not null,
	owner text not null,
	info text default '--'
);