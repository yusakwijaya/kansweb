PGDMP              
        |            kansweb    16.3    16.3     1           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            2           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            3           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            4           1262    16398    kansweb    DATABASE     i   CREATE DATABASE kansweb WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'C';
    DROP DATABASE kansweb;
                postgres    false                        2615    2200    public    SCHEMA        CREATE SCHEMA public;
    DROP SCHEMA public;
                pg_database_owner    false            5           0    0    SCHEMA public    COMMENT     6   COMMENT ON SCHEMA public IS 'standard public schema';
                   pg_database_owner    false    5            �            1259    16436    users    TABLE        CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    email character varying(150) NOT NULL,
    password character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);
    DROP TABLE public.users;
       public         heap    postgres    false    5            �            1259    16442    users_id_seq    SEQUENCE     �   CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.users_id_seq;
       public          postgres    false    216    5            6           0    0    users_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;
          public          postgres    false    217            �           2604    16443    users id    DEFAULT     d   ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);
 7   ALTER TABLE public.users ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    217    216            -          0    16436    users 
   TABLE DATA           F   COPY public.users (id, name, email, password, created_at) FROM stdin;
    public          postgres    false    216   U       7           0    0    users_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.users_id_seq', 8, true);
          public          postgres    false    217            �           2606    16445    users users_email_key 
   CONSTRAINT     Q   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);
 ?   ALTER TABLE ONLY public.users DROP CONSTRAINT users_email_key;
       public            postgres    false    216            �           2606    16447    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    216            -   
  x�u��r�@�5<En�v7�4��������rQ��HT��c%3�$�����l��� �=�c����{ϧ?�%7�� ���n��D�d�s}����M��	j�/��e��k�V��!	�~V.B���G�����B2�1)�1�F��t���iSF��"8�C���jkgh��Ş�rg���G��N���H�+1fxV�n�I�"+L0�O�m]�i{����|{N�<Uå6|��0͞���m������p��<�$���O�� A`D��)���?AËev��M��Bk��@��*X��vӽjx����y=|���2�=�� Q�@#Y��������yb���QFp6#�7ub*���Loeo��vZٮp�U���f������ (ɘ��u�������<���e�9X7E��#7Y�j�i�C��U�㸭ӗYx���v@*QBc��i�{�����&ޒ���^�����i�]�sۀ�z6�I��a��<#k��5D"*��8�'�y��X�     