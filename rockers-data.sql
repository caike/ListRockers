CREATE TABLE people (
    id integer NOT NULL,
    name character varying(255)
);
--
-- Name: people_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--
CREATE SEQUENCE people_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
--
-- Name: people_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--
ALTER SEQUENCE people_id_seq OWNED BY people.id;
--
-- Name: id; Type: DEFAULT; Schema: public; Owner: -
--
ALTER TABLE ONLY people ALTER COLUMN id SET DEFAULT nextval('people_id_seq'::regclass);
ALTER TABLE ONLY people
    ADD CONSTRAINT people_pkey PRIMARY KEY (id);
SET search_path TO "$user",public;

INSERT INTO people (name) VALUES ('James Hetfield'), ('Angela Gossow'),
  ('Michael Amott'), ('Björn Gelotte');
