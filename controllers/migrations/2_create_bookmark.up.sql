create table if not exists bookmark (
	id serial primary key,
	name text not null,
	address text not null,
    info text default '-',
	owner int not null REFERENCES client(id) ON DELETE CASCADE
);