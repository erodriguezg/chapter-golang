CREATE TABLE persons (
	id bigserial NOT NULL,
	rut int4 NOT NULL,
	first_name varchar NOT NULL,
	last_name varchar NOT NULL,
	birthday bool NULL,
	active bool NULL,
	CONSTRAINT persons_pk PRIMARY KEY (id),
	CONSTRAINT persons_un UNIQUE (rut)
);