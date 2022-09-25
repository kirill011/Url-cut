CREATE DATABASE "Url-cut"
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

CREATE TABLE IF NOT EXISTS public.url
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    link text  NOT NULL,
    short character varying(40) NOT NULL,
    CONSTRAINT url_pkey PRIMARY KEY (id),
    CONSTRAINT unique_short UNIQUE (short)
)
