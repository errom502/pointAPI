create table if not exists bookmark (
	id serial primary key,
	name text not null,
	address text not null,
	owner int not null REFERENCES client(id),
	info text default '-'
);