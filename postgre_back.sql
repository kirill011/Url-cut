PGDMP                         {            Url-cut    14.5    14.5 
    �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    16493    Url-cut    DATABASE     f   CREATE DATABASE "Url-cut" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'Russian_Russia.1251';
    DROP DATABASE "Url-cut";
                postgres    false            �            1255    16524    TF_del_dt()    FUNCTION     �   CREATE FUNCTION public."TF_del_dt"() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
begin
delete from public.url
where dt < CURRENT_TIMESTAMP;
RETURN NULL;
end
$$;
 $   DROP FUNCTION public."TF_del_dt"();
       public          postgres    false            �            1259    16495    url    TABLE     �   CREATE TABLE public.url (
    id integer NOT NULL,
    link text NOT NULL,
    short character varying(40) NOT NULL,
    dt timestamp with time zone
);
    DROP TABLE public.url;
       public         heap    postgres    false            �            1259    16494 
   url_id_seq    SEQUENCE     �   ALTER TABLE public.url ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.url_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    210            ^           2606    16503    url unique_short 
   CONSTRAINT     L   ALTER TABLE ONLY public.url
    ADD CONSTRAINT unique_short UNIQUE (short);
 :   ALTER TABLE ONLY public.url DROP CONSTRAINT unique_short;
       public            postgres    false    210            `           2606    16501    url url_pkey 
   CONSTRAINT     J   ALTER TABLE ONLY public.url
    ADD CONSTRAINT url_pkey PRIMARY KEY (id);
 6   ALTER TABLE ONLY public.url DROP CONSTRAINT url_pkey;
       public            postgres    false    210            a           2620    16525    url tr_del_dt    TRIGGER     o   CREATE TRIGGER tr_del_dt BEFORE INSERT ON public.url FOR EACH STATEMENT EXECUTE FUNCTION public."TF_del_dt"();
 &   DROP TRIGGER tr_del_dt ON public.url;
       public          postgres    false    211    210           