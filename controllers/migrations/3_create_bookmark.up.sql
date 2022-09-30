create table if not exists Bookmark (
	id serial primary key,
	name text not null,
	latitude float4 not null,
	longitude float4 not null,
	info text not null default '-',
	owner text not null REFERENCES client(id) ON DELETE CASCADE
);