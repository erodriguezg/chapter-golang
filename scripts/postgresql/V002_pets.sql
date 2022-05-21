CREATE TABLE pets (
	id bigserial NOT NULL,
	owner_rut int4 NOT NULL,
	"name" varchar NOT NULL,
	race varchar NOT NULL,
	CONSTRAINT pets_pk PRIMARY KEY (id),
	CONSTRAINT pets_fk FOREIGN KEY (owner_rut) REFERENCES persons(rut) ON DELETE CASCADE
);