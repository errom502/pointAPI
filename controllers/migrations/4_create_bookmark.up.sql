CREATE TABLE IF NOT EXISTS Bookmark (
	id int4 PRIMARY KEY DEFAULT concat(),
	"name" text NOT NULL,
	latitude float4 NOT NULL,
	longitude float4 NOT NULL,
	info text NOT NULL DEFAULT '-'::text,
	"owner" int4 NULL,
	CONSTRAINT bookmark_fk FOREIGN KEY ("owner") REFERENCES Client(id) ON DELETE SET NULL ON UPDATE CASCADE
);