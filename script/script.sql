CREATE TABLE account
(
    id integer NOT NULL,
    name character varying(50) NOT NULL,
    owner character varying(50) NOT NULL,
    totalstar integer NOT NULL,
    CONSTRAINT id PRIMARY KEY (id)
);