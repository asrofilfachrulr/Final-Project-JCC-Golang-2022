--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2 (Ubuntu 14.2-1.pgdg20.04+1)
-- Dumped by pg_dump version 14.2 (Ubuntu 14.2-1.pgdg20.04+1)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: cart_items; Type: TABLE; Schema: public; Owner: anya
--

CREATE TABLE public.cart_items (
    cart_id bigint NOT NULL,
    product_id bigint NOT NULL,
    price bigint NOT NULL,
    qty bigint NOT NULL,
    sub_total bigint NOT NULL,
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.cart_items OWNER TO anya;

--
-- Name: cart_items_id_seq; Type: SEQUENCE; Schema: public; Owner: anya
--

CREATE SEQUENCE public.cart_items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.cart_items_id_seq OWNER TO anya;

--
-- Name: cart_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: anya
--

ALTER SEQUENCE public.cart_items_id_seq OWNED BY public.cart_items.id;


--
-- Name: carts; Type: TABLE; Schema: public; Owner: anya
--

CREATE TABLE public.carts (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id bigint NOT NULL,
    merchant_id bigint NOT NULL,
    total bigint NOT NULL
);


ALTER TABLE public.carts OWNER TO anya;

--
-- Name: carts_id_seq; Type: SEQUENCE; Schema: public; Owner: anya
--

CREATE SEQUENCE public.carts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.carts_id_seq OWNER TO anya;

--
-- Name: carts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: anya
--

ALTER SEQUENCE public.carts_id_seq OWNED BY public.carts.id;


--
-- Name: categories; Type: TABLE; Schema: public; Owner: anya
--

CREATE TABLE public.categories (
    id bigint NOT NULL,
    name text NOT NULL
);


ALTER TABLE public.categories OWNER TO anya;

--
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: anya
--

CREATE SEQUENCE public.categories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.categories_id_seq OWNER TO anya;

--
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: anya
--

ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;


--
-- Name: countries; Type: TABLE; Schema: public; Owner: anya
--

CREATE TABLE public.countries (
    id bigint NOT NULL,
    name text NOT NULL
);


ALTER TABLE public.countries OWNER TO anya;

--
-- Name: countries_id_seq; Type: SEQUENCE; Schema: public; Owner: anya
--

CREATE SEQUENCE public.countries_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.countries_id_seq OWNER TO anya;

--
-- Name: countries_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: anya
--

ALTER SEQUENCE public.countries_id_seq OWNED BY public.countries.id;


--
-- Name: merchant_addresses; Type: TABLE; Schema: public; Owner: anya
--

CREATE TABLE public.merchant_addresses (
    merchant_id bigint,
    city text NOT NULL,
    offline_store_address text,
    country_id bigint
);


ALTER TABLE public.merchant_addresses OWNER TO anya;

--
-- Name: merchants; Type: TABLE; Schema: public; Owner: anya
--

CREATE TABLE public.merchants (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text NOT NULL,
    rating smallint,
    admin_id bigint NOT NULL
);


ALTER TABLE public.merchants OWNER TO anya;

--
-- Name: merchants_id_seq; Type: SEQUENCE; Schema: public; Owner: anya
--

CREATE SEQUENCE public.merchants_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.merchants_id_seq OWNER TO anya;

--
-- Name: merchants_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: anya
--

ALTER SEQUENCE public.merchants_id_seq OWNED BY public.merchants.id;


--
-- Name: payments; Type: TABLE; Schema: public; Owner: anya
--

CREATE TABLE public.payments (
    id bigint NOT NULL,
    name text NOT NULL,
    method text NOT NULL
);


ALTER TABLE public.payments OWNER TO anya;

--
-- Name: payments_id_seq; Type: SEQUENCE; Schema: public; Owner: anya
--

CREATE SEQUENCE public.payments_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.payments_id_seq OWNER TO anya;

--
-- Name: payments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: anya
--

ALTER SEQUENCE public.payments_id_seq OWNED BY public.payments.id;


--
-- Name: products; Type: TABLE; Schema: public; Owner: anya
--

CREATE TABLE public.products (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text NOT NULL,
    merchant_id bigint NOT NULL,
    price bigint NOT NULL,
    category_id bigint NOT NULL
);


ALTER TABLE public.products OWNER TO anya;

--
-- Name: products_id_seq; Type: SEQUENCE; Schema: public; Owner: anya
--

CREATE SEQUENCE public.products_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.products_id_seq OWNER TO anya;

--
-- Name: products_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: anya
--

ALTER SEQUENCE public.products_id_seq OWNED BY public.products.id;


--
-- Name: prohibit_categories; Type: TABLE; Schema: public; Owner: anya
--

CREATE TABLE public.prohibit_categories (
    category_id bigint NOT NULL,
    country_id bigint NOT NULL
);


ALTER TABLE public.prohibit_categories OWNER TO anya;

--
-- Name: roles; Type: TABLE; Schema: public; Owner: anya
--

CREATE TABLE public.roles (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id bigint NOT NULL,
    name text NOT NULL
);


ALTER TABLE public.roles OWNER TO anya;

--
-- Name: roles_id_seq; Type: SEQUENCE; Schema: public; Owner: anya
--

CREATE SEQUENCE public.roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.roles_id_seq OWNER TO anya;

--
-- Name: roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: anya
--

ALTER SEQUENCE public.roles_id_seq OWNED BY public.roles.id;


--
-- Name: shipments; Type: TABLE; Schema: public; Owner: anya
--

CREATE TABLE public.shipments (
    id bigint NOT NULL,
    name text NOT NULL,
    method text NOT NULL
);


ALTER TABLE public.shipments OWNER TO anya;

--
-- Name: shipments_id_seq; Type: SEQUENCE; Schema: public; Owner: anya
--

CREATE SEQUENCE public.shipments_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.shipments_id_seq OWNER TO anya;

--
-- Name: shipments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: anya
--

ALTER SEQUENCE public.shipments_id_seq OWNED BY public.shipments.id;


--
-- Name: transaction_items; Type: TABLE; Schema: public; Owner: anya
--

CREATE TABLE public.transaction_items (
    transaction_id bigint NOT NULL,
    product_id bigint NOT NULL
);


ALTER TABLE public.transaction_items OWNER TO anya;

--
-- Name: transactions; Type: TABLE; Schema: public; Owner: anya
--

CREATE TABLE public.transactions (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id bigint,
    merchant_id bigint,
    payment_id bigint,
    shipment_id bigint,
    username text NOT NULL,
    merchant_name text NOT NULL,
    payment_info text NOT NULL,
    shipment_info text NOT NULL,
    shipping_address text NOT NULL,
    total bigint NOT NULL,
    status text NOT NULL
);


ALTER TABLE public.transactions OWNER TO anya;

--
-- Name: transactions_id_seq; Type: SEQUENCE; Schema: public; Owner: anya
--

CREATE SEQUENCE public.transactions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transactions_id_seq OWNER TO anya;

--
-- Name: transactions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: anya
--

ALTER SEQUENCE public.transactions_id_seq OWNED BY public.transactions.id;


--
-- Name: user_addresses; Type: TABLE; Schema: public; Owner: anya
--

CREATE TABLE public.user_addresses (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id bigint NOT NULL,
    address_line text NOT NULL,
    city text NOT NULL,
    country_id bigint NOT NULL,
    phone_number bigint NOT NULL,
    postal_code bigint
);


ALTER TABLE public.user_addresses OWNER TO anya;

--
-- Name: user_addresses_id_seq; Type: SEQUENCE; Schema: public; Owner: anya
--

CREATE SEQUENCE public.user_addresses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_addresses_id_seq OWNER TO anya;

--
-- Name: user_addresses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: anya
--

ALTER SEQUENCE public.user_addresses_id_seq OWNED BY public.user_addresses.id;


--
-- Name: user_credentials; Type: TABLE; Schema: public; Owner: anya
--

CREATE TABLE public.user_credentials (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id bigint NOT NULL,
    username text NOT NULL,
    password text NOT NULL
);


ALTER TABLE public.user_credentials OWNER TO anya;

--
-- Name: user_credentials_id_seq; Type: SEQUENCE; Schema: public; Owner: anya
--

CREATE SEQUENCE public.user_credentials_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_credentials_id_seq OWNER TO anya;

--
-- Name: user_credentials_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: anya
--

ALTER SEQUENCE public.user_credentials_id_seq OWNED BY public.user_credentials.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: anya
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    full_name text NOT NULL,
    username text NOT NULL,
    email text NOT NULL
);


ALTER TABLE public.users OWNER TO anya;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: anya
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO anya;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: anya
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: cart_items id; Type: DEFAULT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.cart_items ALTER COLUMN id SET DEFAULT nextval('public.cart_items_id_seq'::regclass);


--
-- Name: carts id; Type: DEFAULT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.carts ALTER COLUMN id SET DEFAULT nextval('public.carts_id_seq'::regclass);


--
-- Name: categories id; Type: DEFAULT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);


--
-- Name: countries id; Type: DEFAULT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.countries ALTER COLUMN id SET DEFAULT nextval('public.countries_id_seq'::regclass);


--
-- Name: merchants id; Type: DEFAULT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.merchants ALTER COLUMN id SET DEFAULT nextval('public.merchants_id_seq'::regclass);


--
-- Name: payments id; Type: DEFAULT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.payments ALTER COLUMN id SET DEFAULT nextval('public.payments_id_seq'::regclass);


--
-- Name: products id; Type: DEFAULT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_id_seq'::regclass);


--
-- Name: roles id; Type: DEFAULT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.roles ALTER COLUMN id SET DEFAULT nextval('public.roles_id_seq'::regclass);


--
-- Name: shipments id; Type: DEFAULT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.shipments ALTER COLUMN id SET DEFAULT nextval('public.shipments_id_seq'::regclass);


--
-- Name: transactions id; Type: DEFAULT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.transactions ALTER COLUMN id SET DEFAULT nextval('public.transactions_id_seq'::regclass);


--
-- Name: user_addresses id; Type: DEFAULT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.user_addresses ALTER COLUMN id SET DEFAULT nextval('public.user_addresses_id_seq'::regclass);


--
-- Name: user_credentials id; Type: DEFAULT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.user_credentials ALTER COLUMN id SET DEFAULT nextval('public.user_credentials_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: cart_items; Type: TABLE DATA; Schema: public; Owner: anya
--

COPY public.cart_items (cart_id, product_id, price, qty, sub_total, id, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: carts; Type: TABLE DATA; Schema: public; Owner: anya
--

COPY public.carts (id, created_at, updated_at, deleted_at, user_id, merchant_id, total) FROM stdin;
\.


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: anya
--

COPY public.categories (id, name) FROM stdin;
1	Foodie
2	Beer
3	Beverage
4	Gadget
5	Laptop
6	Electronics
7	Men
8	Women
9	Outdoors
10	Health
11	Household
12	Books
13	Tools
\.


--
-- Data for Name: countries; Type: TABLE DATA; Schema: public; Owner: anya
--

COPY public.countries (id, name) FROM stdin;
1	Indonesia
2	Singapore
3	Malaysia
4	Thailand
5	Brunei
6	Philipines
7	Laos
8	Vietnam
9	Cambodia
10	Myanmar
\.


--
-- Data for Name: merchant_addresses; Type: TABLE DATA; Schema: public; Owner: anya
--

COPY public.merchant_addresses (merchant_id, city, offline_store_address, country_id) FROM stdin;
\.


--
-- Data for Name: merchants; Type: TABLE DATA; Schema: public; Owner: anya
--

COPY public.merchants (id, created_at, updated_at, deleted_at, name, rating, admin_id) FROM stdin;
\.


--
-- Data for Name: payments; Type: TABLE DATA; Schema: public; Owner: anya
--

COPY public.payments (id, name, method) FROM stdin;
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: anya
--

COPY public.products (id, created_at, updated_at, deleted_at, name, merchant_id, price, category_id) FROM stdin;
\.


--
-- Data for Name: prohibit_categories; Type: TABLE DATA; Schema: public; Owner: anya
--

COPY public.prohibit_categories (category_id, country_id) FROM stdin;
\.


--
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: anya
--

COPY public.roles (id, created_at, updated_at, deleted_at, user_id, name) FROM stdin;
1	2022-04-03 09:50:39.827923+07	2022-04-03 09:50:39.827923+07	\N	1	customer
2	2022-04-03 09:50:39.894991+07	2022-04-03 09:50:39.894991+07	\N	2	customer
3	2022-04-03 09:50:39.960284+07	2022-04-03 09:50:39.960284+07	\N	3	customer
4	2022-04-03 09:50:40.027112+07	2022-04-03 09:50:40.027112+07	\N	4	customer
5	2022-04-03 09:50:40.092376+07	2022-04-03 09:50:40.092376+07	\N	5	customer
6	2022-04-03 09:50:40.158731+07	2022-04-03 09:50:40.158731+07	\N	6	customer
7	2022-04-03 09:50:40.225303+07	2022-04-03 09:50:40.225303+07	\N	7	customer
8	2022-04-03 09:50:40.291018+07	2022-04-03 09:50:40.291018+07	\N	8	customer
9	2022-04-03 09:50:40.36335+07	2022-04-03 09:50:40.36335+07	\N	9	dev
\.


--
-- Data for Name: shipments; Type: TABLE DATA; Schema: public; Owner: anya
--

COPY public.shipments (id, name, method) FROM stdin;
\.


--
-- Data for Name: transaction_items; Type: TABLE DATA; Schema: public; Owner: anya
--

COPY public.transaction_items (transaction_id, product_id) FROM stdin;
\.


--
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: anya
--

COPY public.transactions (id, created_at, updated_at, deleted_at, user_id, merchant_id, payment_id, shipment_id, username, merchant_name, payment_info, shipment_info, shipping_address, total, status) FROM stdin;
\.


--
-- Data for Name: user_addresses; Type: TABLE DATA; Schema: public; Owner: anya
--

COPY public.user_addresses (id, created_at, updated_at, deleted_at, user_id, address_line, city, country_id, phone_number, postal_code) FROM stdin;
\.


--
-- Data for Name: user_credentials; Type: TABLE DATA; Schema: public; Owner: anya
--

COPY public.user_credentials (id, created_at, updated_at, deleted_at, user_id, username, password) FROM stdin;
1	2022-04-03 09:50:39.821457+07	2022-04-03 09:50:39.821457+07	\N	1	john	$2a$10$1uKDG.gNrfSuLoT06Iq1Xe02Kj2KehNOHMgfd/H3selRAoPlnWDF6
2	2022-04-03 09:50:39.890091+07	2022-04-03 09:50:39.890091+07	\N	2	mary	$2a$10$9Ph56vQTajfYAAEL1G8p6u0J5s1L1n7WUTNiJYjwka9oFtgyLz.MG
3	2022-04-03 09:50:39.955609+07	2022-04-03 09:50:39.955609+07	\N	3	xi	$2a$10$dptFe8lTD2uErpp110sUCeUk.b9yoqiBT.Hp5qON58oMD/23ZCQES
4	2022-04-03 09:50:40.022047+07	2022-04-03 09:50:40.022047+07	\N	4	mark	$2a$10$9qiuw3LRVvaCX45NMJ36quFE/2lM65ucgWSNzR35fHJY9ePr62cvm
5	2022-04-03 09:50:40.087513+07	2022-04-03 09:50:40.087513+07	\N	5	ng	$2a$10$O0wMGYv1VTYj6VP9l4YHVeaaiWm9QLuQyApKUS94DCZOTRkxNN8zm
6	2022-04-03 09:50:40.153851+07	2022-04-03 09:50:40.153851+07	\N	6	jack	$2a$10$cy5B7B7GWxPVYkJ.Gh0VPed0Q6qXFQdPxdFLi/61XpMEXUgV3n1Se
7	2022-04-03 09:50:40.220324+07	2022-04-03 09:50:40.220324+07	\N	7	tony	$2a$10$O/EJ6YkqK62e/BLndhqOVuvA0iJ7Af/rP7AeUyOpo9jpau6SR/WUK
8	2022-04-03 09:50:40.286222+07	2022-04-03 09:50:40.286222+07	\N	8	andy	$2a$10$uvP8AyFdtg4X1EZFEfKDo.TF8a8pqXxSPY9m9V/mT4Oue89ZkboKy
9	2022-04-03 09:50:40.358267+07	2022-04-03 09:50:40.358267+07	\N	9	anya	$2a$10$2/EWcaQMl7C9oGJmNWBR/ODo3u38DzT7z/e1SnLQexGf913MNwYB6
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: anya
--

COPY public.users (id, created_at, updated_at, deleted_at, full_name, username, email) FROM stdin;
1	2022-04-03 09:50:39.758057+07	2022-04-03 09:50:39.758057+07	\N	John Doe	john	john@mail.com
2	2022-04-03 09:50:39.758057+07	2022-04-03 09:50:39.758057+07	\N	Mary Sue	mary	mary@mail.com
3	2022-04-03 09:50:39.758057+07	2022-04-03 09:50:39.758057+07	\N	Xi Ng	xi	nihaoma@mail.com
4	2022-04-03 09:50:39.758057+07	2022-04-03 09:50:39.758057+07	\N	Mark Bob	mark	mark@mail.com
5	2022-04-03 09:50:39.758057+07	2022-04-03 09:50:39.758057+07	\N	Patricia Ng	ng	ng@mail.com
6	2022-04-03 09:50:39.758057+07	2022-04-03 09:50:39.758057+07	\N	Jack Tyler	jack	jack@mail.com
7	2022-04-03 09:50:39.758057+07	2022-04-03 09:50:39.758057+07	\N	Tony Like	tony	tony@mail.com
8	2022-04-03 09:50:39.758057+07	2022-04-03 09:50:39.758057+07	\N	Andy Lim	andy	andy@mail.com
9	2022-04-03 09:50:40.295949+07	2022-04-03 09:50:40.295949+07	\N	anya	anya	riidloa@gmail.com
\.


--
-- Name: cart_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: anya
--

SELECT pg_catalog.setval('public.cart_items_id_seq', 1, false);


--
-- Name: carts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: anya
--

SELECT pg_catalog.setval('public.carts_id_seq', 1, false);


--
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: anya
--

SELECT pg_catalog.setval('public.categories_id_seq', 13, true);


--
-- Name: countries_id_seq; Type: SEQUENCE SET; Schema: public; Owner: anya
--

SELECT pg_catalog.setval('public.countries_id_seq', 10, true);


--
-- Name: merchants_id_seq; Type: SEQUENCE SET; Schema: public; Owner: anya
--

SELECT pg_catalog.setval('public.merchants_id_seq', 1, false);


--
-- Name: payments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: anya
--

SELECT pg_catalog.setval('public.payments_id_seq', 1, false);


--
-- Name: products_id_seq; Type: SEQUENCE SET; Schema: public; Owner: anya
--

SELECT pg_catalog.setval('public.products_id_seq', 1, false);


--
-- Name: roles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: anya
--

SELECT pg_catalog.setval('public.roles_id_seq', 9, true);


--
-- Name: shipments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: anya
--

SELECT pg_catalog.setval('public.shipments_id_seq', 1, false);


--
-- Name: transactions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: anya
--

SELECT pg_catalog.setval('public.transactions_id_seq', 1, false);


--
-- Name: user_addresses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: anya
--

SELECT pg_catalog.setval('public.user_addresses_id_seq', 1, false);


--
-- Name: user_credentials_id_seq; Type: SEQUENCE SET; Schema: public; Owner: anya
--

SELECT pg_catalog.setval('public.user_credentials_id_seq', 9, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: anya
--

SELECT pg_catalog.setval('public.users_id_seq', 9, true);


--
-- Name: carts carts_pkey; Type: CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.carts
    ADD CONSTRAINT carts_pkey PRIMARY KEY (id);


--
-- Name: categories categories_name_key; Type: CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_name_key UNIQUE (name);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: countries countries_name_key; Type: CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.countries
    ADD CONSTRAINT countries_name_key UNIQUE (name);


--
-- Name: countries countries_pkey; Type: CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.countries
    ADD CONSTRAINT countries_pkey PRIMARY KEY (id);


--
-- Name: merchants merchants_pkey; Type: CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.merchants
    ADD CONSTRAINT merchants_pkey PRIMARY KEY (id);


--
-- Name: payments payments_pkey; Type: CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.payments
    ADD CONSTRAINT payments_pkey PRIMARY KEY (id);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- Name: prohibit_categories prohibit_categories_pkey; Type: CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.prohibit_categories
    ADD CONSTRAINT prohibit_categories_pkey PRIMARY KEY (category_id, country_id);


--
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


--
-- Name: shipments shipments_pkey; Type: CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.shipments
    ADD CONSTRAINT shipments_pkey PRIMARY KEY (id);


--
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- Name: user_addresses user_addresses_pkey; Type: CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.user_addresses
    ADD CONSTRAINT user_addresses_pkey PRIMARY KEY (id);


--
-- Name: user_credentials user_credentials_pkey; Type: CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.user_credentials
    ADD CONSTRAINT user_credentials_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- Name: idx_cart_items_deleted_at; Type: INDEX; Schema: public; Owner: anya
--

CREATE INDEX idx_cart_items_deleted_at ON public.cart_items USING btree (deleted_at);


--
-- Name: idx_carts_deleted_at; Type: INDEX; Schema: public; Owner: anya
--

CREATE INDEX idx_carts_deleted_at ON public.carts USING btree (deleted_at);


--
-- Name: idx_merchants_deleted_at; Type: INDEX; Schema: public; Owner: anya
--

CREATE INDEX idx_merchants_deleted_at ON public.merchants USING btree (deleted_at);


--
-- Name: idx_products_deleted_at; Type: INDEX; Schema: public; Owner: anya
--

CREATE INDEX idx_products_deleted_at ON public.products USING btree (deleted_at);


--
-- Name: idx_roles_deleted_at; Type: INDEX; Schema: public; Owner: anya
--

CREATE INDEX idx_roles_deleted_at ON public.roles USING btree (deleted_at);


--
-- Name: idx_transactions_deleted_at; Type: INDEX; Schema: public; Owner: anya
--

CREATE INDEX idx_transactions_deleted_at ON public.transactions USING btree (deleted_at);


--
-- Name: idx_user_addresses_deleted_at; Type: INDEX; Schema: public; Owner: anya
--

CREATE INDEX idx_user_addresses_deleted_at ON public.user_addresses USING btree (deleted_at);


--
-- Name: idx_user_credentials_deleted_at; Type: INDEX; Schema: public; Owner: anya
--

CREATE INDEX idx_user_credentials_deleted_at ON public.user_credentials USING btree (deleted_at);


--
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: anya
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


--
-- Name: cart_items fk_cart_items_product; Type: FK CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.cart_items
    ADD CONSTRAINT fk_cart_items_product FOREIGN KEY (product_id) REFERENCES public.products(id) ON DELETE CASCADE;


--
-- Name: cart_items fk_carts_cart_items; Type: FK CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.cart_items
    ADD CONSTRAINT fk_carts_cart_items FOREIGN KEY (cart_id) REFERENCES public.carts(id);


--
-- Name: carts fk_carts_merchant; Type: FK CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.carts
    ADD CONSTRAINT fk_carts_merchant FOREIGN KEY (merchant_id) REFERENCES public.merchants(id) ON DELETE CASCADE;


--
-- Name: carts fk_carts_user; Type: FK CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.carts
    ADD CONSTRAINT fk_carts_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: merchant_addresses fk_merchant_addresses_country; Type: FK CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.merchant_addresses
    ADD CONSTRAINT fk_merchant_addresses_country FOREIGN KEY (country_id) REFERENCES public.countries(id);


--
-- Name: merchant_addresses fk_merchant_addresses_merchant; Type: FK CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.merchant_addresses
    ADD CONSTRAINT fk_merchant_addresses_merchant FOREIGN KEY (merchant_id) REFERENCES public.merchants(id) ON DELETE CASCADE;


--
-- Name: merchants fk_merchants_user; Type: FK CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.merchants
    ADD CONSTRAINT fk_merchants_user FOREIGN KEY (admin_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: products fk_products_category; Type: FK CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT fk_products_category FOREIGN KEY (category_id) REFERENCES public.categories(id);


--
-- Name: products fk_products_merchant; Type: FK CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT fk_products_merchant FOREIGN KEY (merchant_id) REFERENCES public.merchants(id) ON DELETE CASCADE;


--
-- Name: prohibit_categories fk_prohibit_categories_category; Type: FK CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.prohibit_categories
    ADD CONSTRAINT fk_prohibit_categories_category FOREIGN KEY (category_id) REFERENCES public.categories(id);


--
-- Name: prohibit_categories fk_prohibit_categories_country; Type: FK CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.prohibit_categories
    ADD CONSTRAINT fk_prohibit_categories_country FOREIGN KEY (country_id) REFERENCES public.countries(id);


--
-- Name: roles fk_roles_user; Type: FK CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT fk_roles_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: transactions fk_transactions_merchant; Type: FK CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT fk_transactions_merchant FOREIGN KEY (merchant_id) REFERENCES public.merchants(id);


--
-- Name: transactions fk_transactions_payment; Type: FK CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT fk_transactions_payment FOREIGN KEY (payment_id) REFERENCES public.payments(id);


--
-- Name: transactions fk_transactions_shipment; Type: FK CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT fk_transactions_shipment FOREIGN KEY (shipment_id) REFERENCES public.shipments(id);


--
-- Name: transaction_items fk_transactions_transaction_items; Type: FK CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.transaction_items
    ADD CONSTRAINT fk_transactions_transaction_items FOREIGN KEY (transaction_id) REFERENCES public.transactions(id);


--
-- Name: transactions fk_transactions_user; Type: FK CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT fk_transactions_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: user_addresses fk_user_addresses_country; Type: FK CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.user_addresses
    ADD CONSTRAINT fk_user_addresses_country FOREIGN KEY (country_id) REFERENCES public.countries(id);


--
-- Name: user_addresses fk_user_addresses_user; Type: FK CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.user_addresses
    ADD CONSTRAINT fk_user_addresses_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: user_credentials fk_user_credentials_user; Type: FK CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.user_credentials
    ADD CONSTRAINT fk_user_credentials_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

