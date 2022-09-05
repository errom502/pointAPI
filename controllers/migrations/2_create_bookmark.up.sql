create table if not exists Bookmark (
	id serial primary key,
	name text not null,
	address text not null,
	info text not null default '-',
	owner int not null REFERENCES client(id) ON DELETE CASCADE
);