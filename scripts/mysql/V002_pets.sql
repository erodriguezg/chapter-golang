CREATE TABLE pets (
	id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	owner_rut INT NOT NULL,
	name VARCHAR(100) NOT NULL,
	race VARCHAR(100) NOT NULL,
	CONSTRAINT pets_pk PRIMARY KEY (id),
	CONSTRAINT pets_fk FOREIGN KEY (owner_rut) REFERENCES persons(rut) ON DELETE CASCADE
);