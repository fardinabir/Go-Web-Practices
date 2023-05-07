--
-- PostgreSQL database dump
--

-- Dumped from database version 14.7 (Ubuntu 14.7-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.7 (Ubuntu 14.7-0ubuntu0.22.04.1)

-- Started on 2023-03-08 12:28:34 +06

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
-- TOC entry 214 (class 1259 OID 16418)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_name text,
    mobile_no text,
    password text
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 213 (class 1259 OID 16417)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 3368 (class 0 OID 0)
-- Dependencies: 213
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 3218 (class 2604 OID 16421)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 3362 (class 0 OID 16418)
-- Dependencies: 214
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users VALUES (1, '2023-01-19 17:05:54.96073+06', '2023-01-19 17:05:54.96073+06', NULL, 'FardinAbir', '01770045492', '1234');
INSERT INTO public.users VALUES (3, '2023-01-19 17:12:54.028329+06', '2023-01-19 17:12:54.028329+06', NULL, 'FardinAbir', '01770045492', '1234');
INSERT INTO public.users VALUES (4, '2023-01-19 17:14:22.920424+06', '2023-01-19 17:14:22.920424+06', NULL, 'FardinAbir', '01770045492', '1234');
INSERT INTO public.users VALUES (2, '2023-01-19 17:08:46.127323+06', '2023-01-19 17:08:46.127323+06', '2023-01-19 18:19:42.751301+06', 'FardinAbir', '01770045492', '1234');
INSERT INTO public.users VALUES (5, '2023-02-05 09:37:19.156638+06', '2023-02-05 09:37:19.156638+06', '2023-02-05 09:38:34.73931+06', '', '01770045492', '1234');
INSERT INTO public.users VALUES (6, '2023-02-05 09:39:08.085685+06', '2023-02-05 09:39:08.085685+06', NULL, 'fardinNew', '01770045492', '1234');
INSERT INTO public.users VALUES (7, '2023-02-05 12:54:56.899781+06', '2023-02-05 12:54:56.899781+06', NULL, 'fardinNew', '01770045492', '1234');
INSERT INTO public.users VALUES (8, '2023-02-05 12:58:37.526167+06', '2023-02-05 12:58:37.526167+06', NULL, 'fardinNew1', '01770045492', '$argon2id$v=19$m=65536,t=1,p=2$uAcFEsrIQxJb4o29t2tsdw$G4Ml1xGIjfMq5gAt5lEzT0ylNu440gwdEgxzs6xbLUU');


--
-- TOC entry 3369 (class 0 OID 0)
-- Dependencies: 213
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 8, true);


--
-- TOC entry 3221 (class 2606 OID 16425)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 3219 (class 1259 OID 16426)
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


-- Completed on 2023-03-08 12:28:34 +06

--
-- PostgreSQL database dump complete
--

