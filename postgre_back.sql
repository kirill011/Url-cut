toc.dat                                                                                             0000600 0004000 0002000 00000010010 14406524751 0014437 0                                                                                                    ustar 00postgres                        postgres                        0000000 0000000                                                                                                                                                                        PGDMP            
    
            {            Url-cut    14.5    14.5     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false         �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false         �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false         �           1262    16493    Url-cut    DATABASE     f   CREATE DATABASE "Url-cut" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'Russian_Russia.1251';
    DROP DATABASE "Url-cut";
                postgres    false         �            1255    16524    TF_del_dt()    FUNCTION     �   CREATE FUNCTION public."TF_del_dt"() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
begin
delete from public.url
where dt < CURRENT_TIMESTAMP;
RETURN NULL;
end
$$;
 $   DROP FUNCTION public."TF_del_dt"();
       public          postgres    false         �            1255    16570    TF_del_short()    FUNCTION     �   CREATE FUNCTION public."TF_del_short"() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
Begin
insert into deleted (short) 
values
(OLD.short);
return OLD;
End
$$;
 '   DROP FUNCTION public."TF_del_short"();
       public          postgres    false         �            1255    16572    TF_isert_short()    FUNCTION     �   CREATE FUNCTION public."TF_isert_short"() RETURNS trigger
    LANGUAGE plpgsql
    AS $$begin

delete from deleted 
where short = new.short;
return new;
end$$;
 )   DROP FUNCTION public."TF_isert_short"();
       public          postgres    false         �            1259    16556    deleted    TABLE     =   CREATE TABLE public.deleted (
    short character varying
);
    DROP TABLE public.deleted;
       public         heap    postgres    false         �            1259    16495    url    TABLE     �   CREATE TABLE public.url (
    id integer NOT NULL,
    link text NOT NULL,
    short character varying NOT NULL,
    dt timestamp with time zone
);
    DROP TABLE public.url;
       public         heap    postgres    false         �            1259    16494 
   url_id_seq    SEQUENCE     �   ALTER TABLE public.url ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.url_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    210         d           2606    16560    url unique_short 
   CONSTRAINT     L   ALTER TABLE ONLY public.url
    ADD CONSTRAINT unique_short UNIQUE (short);
 :   ALTER TABLE ONLY public.url DROP CONSTRAINT unique_short;
       public            postgres    false    210         f           2606    16501    url url_pkey 
   CONSTRAINT     J   ALTER TABLE ONLY public.url
    ADD CONSTRAINT url_pkey PRIMARY KEY (id);
 6   ALTER TABLE ONLY public.url DROP CONSTRAINT url_pkey;
       public            postgres    false    210         i           2620    16525    url tr_del_dt    TRIGGER     o   CREATE TRIGGER tr_del_dt BEFORE INSERT ON public.url FOR EACH STATEMENT EXECUTE FUNCTION public."TF_del_dt"();
 &   DROP TRIGGER tr_del_dt ON public.url;
       public          postgres    false    212    210         h           2620    16571    url tr_del_sh    TRIGGER     l   CREATE TRIGGER tr_del_sh BEFORE DELETE ON public.url FOR EACH ROW EXECUTE FUNCTION public."TF_del_short"();
 &   DROP TRIGGER tr_del_sh ON public.url;
       public          postgres    false    213    210         g           2620    16573    url tr_insert_short    TRIGGER     t   CREATE TRIGGER tr_insert_short BEFORE INSERT ON public.url FOR EACH ROW EXECUTE FUNCTION public."TF_isert_short"();
 ,   DROP TRIGGER tr_insert_short ON public.url;
       public          postgres    false    214    210                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                restore.sql                                                                                         0000600 0004000 0002000 00000007430 14406524751 0015400 0                                                                                                    ustar 00postgres                        postgres                        0000000 0000000                                                                                                                                                                        --
-- NOTE:
--
-- File paths need to be edited. Search for $$PATH$$ and
-- replace it with the path to the directory containing
-- the extracted data files.
--
--
-- PostgreSQL database dump
--

-- Dumped from database version 14.5
-- Dumped by pg_dump version 14.5

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

DROP DATABASE "Url-cut";
--
-- Name: Url-cut; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE "Url-cut" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'Russian_Russia.1251';


ALTER DATABASE "Url-cut" OWNER TO postgres;

\connect -reuse-previous=on "dbname='Url-cut'"

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: TF_del_dt(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public."TF_del_dt"() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
begin
delete from public.url
where dt < CURRENT_TIMESTAMP;
RETURN NULL;
end
$$;


ALTER FUNCTION public."TF_del_dt"() OWNER TO postgres;

--
-- Name: TF_del_short(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public."TF_del_short"() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
Begin
insert into deleted (short) 
values
(OLD.short);
return OLD;
End
$$;


ALTER FUNCTION public."TF_del_short"() OWNER TO postgres;

--
-- Name: TF_isert_short(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public."TF_isert_short"() RETURNS trigger
    LANGUAGE plpgsql
    AS $$begin

delete from deleted 
where short = new.short;
return new;
end$$;


ALTER FUNCTION public."TF_isert_short"() OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: deleted; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.deleted (
    short character varying
);


ALTER TABLE public.deleted OWNER TO postgres;

--
-- Name: url; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.url (
    id integer NOT NULL,
    link text NOT NULL,
    short character varying NOT NULL,
    dt timestamp with time zone
);


ALTER TABLE public.url OWNER TO postgres;

--
-- Name: url_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.url ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.url_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: url unique_short; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.url
    ADD CONSTRAINT unique_short UNIQUE (short);


--
-- Name: url url_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.url
    ADD CONSTRAINT url_pkey PRIMARY KEY (id);


--
-- Name: url tr_del_dt; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER tr_del_dt BEFORE INSERT ON public.url FOR EACH STATEMENT EXECUTE FUNCTION public."TF_del_dt"();


--
-- Name: url tr_del_sh; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER tr_del_sh BEFORE DELETE ON public.url FOR EACH ROW EXECUTE FUNCTION public."TF_del_short"();


--
-- Name: url tr_insert_short; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER tr_insert_short BEFORE INSERT ON public.url FOR EACH ROW EXECUTE FUNCTION public."TF_isert_short"();


--
-- PostgreSQL database dump complete
--

                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        