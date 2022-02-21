--
-- PostgreSQL database dump
--

-- Dumped from database version 11.3
-- Dumped by pg_dump version 13.4

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

ALTER TABLE IF EXISTS ONLY public.tags_for_task DROP CONSTRAINT IF EXISTS tags_for_task_tasks_id_fk;
ALTER TABLE IF EXISTS ONLY public.tags_for_task DROP CONSTRAINT IF EXISTS tags_for_task_entry_tags_id_fk;
ALTER TABLE IF EXISTS ONLY public.resolutions DROP CONSTRAINT IF EXISTS resolutions_resolution_entity_id_fk;
ALTER TABLE IF EXISTS ONLY public.resolution_files DROP CONSTRAINT IF EXISTS resolution_files_resolutions_id_fk;
ALTER TABLE IF EXISTS ONLY public.entry_levels DROP CONSTRAINT IF EXISTS entry_levels_levels_id_fk;
ALTER TABLE IF EXISTS ONLY public.entry_levels DROP CONSTRAINT IF EXISTS entry_levels_entries_id_fk;
ALTER TABLE IF EXISTS ONLY public.tasks DROP CONSTRAINT IF EXISTS completed_tasks_tasks_id_fk;
DROP INDEX IF EXISTS public.test_madrs_id_uindex;
DROP INDEX IF EXISTS public.spawn_tasks_title_uindex;
DROP INDEX IF EXISTS public.resolutions_id_uindex;
DROP INDEX IF EXISTS public.resolution_files_id_uindex;
DROP INDEX IF EXISTS public.resolution_files_filename_uindex;
DROP INDEX IF EXISTS public.resolution_entity_name_uindex;
DROP INDEX IF EXISTS public.resolution_entity_id_uindex;
DROP INDEX IF EXISTS public.levels_shortname_uindex;
DROP INDEX IF EXISTS public.levels_name_uindex;
DROP INDEX IF EXISTS public.entry_tags_shortname_uindex;
DROP INDEX IF EXISTS public.entry_tags_name_uindex;
DROP INDEX IF EXISTS public.entries_fromtime_uindex;
ALTER TABLE IF EXISTS ONLY public.weight DROP CONSTRAINT IF EXISTS weight_pk;
ALTER TABLE IF EXISTS ONLY public.test_madrs DROP CONSTRAINT IF EXISTS test_madrs_pk;
ALTER TABLE IF EXISTS ONLY public.spawn_tasks DROP CONSTRAINT IF EXISTS tasks_pk;
ALTER TABLE IF EXISTS ONLY public.tags_for_task DROP CONSTRAINT IF EXISTS tags_for_task_pk;
ALTER TABLE IF EXISTS ONLY public.tags_for_entry DROP CONSTRAINT IF EXISTS tags_for_entry_pk;
ALTER TABLE IF EXISTS ONLY public.resolutions DROP CONSTRAINT IF EXISTS resolutions_pk;
ALTER TABLE IF EXISTS ONLY public.resolution_files DROP CONSTRAINT IF EXISTS resolution_files_pk;
ALTER TABLE IF EXISTS ONLY public.resolution_entity DROP CONSTRAINT IF EXISTS resolution_entity_pk;
ALTER TABLE IF EXISTS ONLY public.levels DROP CONSTRAINT IF EXISTS levels_pk;
ALTER TABLE IF EXISTS ONLY public.height DROP CONSTRAINT IF EXISTS height_pk;
ALTER TABLE IF EXISTS ONLY public.entry_tags DROP CONSTRAINT IF EXISTS entry_tags_pk;
ALTER TABLE IF EXISTS ONLY public.entry_levels DROP CONSTRAINT IF EXISTS entry_levels_pk;
ALTER TABLE IF EXISTS ONLY public.entries DROP CONSTRAINT IF EXISTS entries_pk;
ALTER TABLE IF EXISTS ONLY public.tasks DROP CONSTRAINT IF EXISTS completed_tasks_pk;
ALTER TABLE IF EXISTS public.weight ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.test_madrs ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.tasks ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.spawn_tasks ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.resolutions ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.resolution_files ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.resolution_entity ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.levels ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.height ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.entry_tags ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.entries ALTER COLUMN id DROP DEFAULT;
DROP SEQUENCE IF EXISTS public.weight_id_seq;
DROP TABLE IF EXISTS public.weight;
DROP TABLE IF EXISTS public.version;
DROP SEQUENCE IF EXISTS public.test_madrs_id_seq;
DROP TABLE IF EXISTS public.test_madrs;
DROP SEQUENCE IF EXISTS public.tasks_id_seq;
DROP TABLE IF EXISTS public.tags_for_task;
DROP TABLE IF EXISTS public.tags_for_entry;
DROP TABLE IF EXISTS public.spawn_tasks;
DROP SEQUENCE IF EXISTS public.resolutions_id_seq;
DROP TABLE IF EXISTS public.resolutions;
DROP SEQUENCE IF EXISTS public.resolution_files_id_seq;
DROP TABLE IF EXISTS public.resolution_files;
DROP SEQUENCE IF EXISTS public.resolution_entity_id_seq;
DROP TABLE IF EXISTS public.resolution_entity;
DROP SEQUENCE IF EXISTS public.levels_id_seq;
DROP TABLE IF EXISTS public.levels;
DROP SEQUENCE IF EXISTS public.height_id_seq;
DROP TABLE IF EXISTS public.height;
DROP SEQUENCE IF EXISTS public.entry_tags_id_seq;
DROP TABLE IF EXISTS public.entry_tags;
DROP TABLE IF EXISTS public.entry_levels;
DROP SEQUENCE IF EXISTS public.entries_id_seq;
DROP TABLE IF EXISTS public.entries;
DROP SEQUENCE IF EXISTS public.completed_tasks_id_seq;
DROP TABLE IF EXISTS public.tasks;
--
-- Name: tasks; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.tasks (
    id integer NOT NULL,
    refid integer,
    completed_at timestamp without time zone,
    title text NOT NULL,
    descr text DEFAULT ''::text NOT NULL,
    added_at timestamp without time zone DEFAULT now() NOT NULL,
    modified_at timestamp without time zone,
    deleted_at timestamp without time zone,
    priority integer DEFAULT 0 NOT NULL
);


--
-- Name: COLUMN tasks.refid; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.tasks.refid IS 'optional';


--
-- Name: completed_tasks_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.completed_tasks_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: completed_tasks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.completed_tasks_id_seq OWNED BY public.tasks.id;


--
-- Name: entries; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.entries (
    id integer NOT NULL,
    added_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    activity_name text DEFAULT ''::text NOT NULL,
    level_achievement integer DEFAULT 0 NOT NULL,
    fromtime timestamp without time zone DEFAULT now() NOT NULL,
    modified_at timestamp without time zone,
    descr text DEFAULT ''::text NOT NULL
);


--
-- Name: COLUMN entries.deleted_at; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.entries.deleted_at IS 'when entry was deleted';


--
-- Name: entries_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.entries_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: entries_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.entries_id_seq OWNED BY public.entries.id;


--
-- Name: entry_levels; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.entry_levels (
    entryid integer NOT NULL,
    levelid integer NOT NULL,
    added_at timestamp without time zone DEFAULT now() NOT NULL,
    level integer DEFAULT 0 NOT NULL
);


--
-- Name: entry_tags; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.entry_tags (
    id integer NOT NULL,
    name text DEFAULT ''::text NOT NULL,
    added_at timestamp without time zone DEFAULT now() NOT NULL,
    shortname text DEFAULT ''::text NOT NULL
);


--
-- Name: entry_tags_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.entry_tags_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: entry_tags_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.entry_tags_id_seq OWNED BY public.entry_tags.id;


--
-- Name: height; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.height (
    id integer NOT NULL,
    added_at timestamp without time zone DEFAULT now() NOT NULL,
    value numeric DEFAULT 0.0 NOT NULL
);


--
-- Name: height_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.height_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: height_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.height_id_seq OWNED BY public.height.id;


--
-- Name: levels; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.levels (
    id integer NOT NULL,
    added_at timestamp without time zone DEFAULT now() NOT NULL,
    name text DEFAULT ''::text NOT NULL,
    default_show boolean DEFAULT true NOT NULL,
    worst_descr text DEFAULT ''::text NOT NULL,
    shortname text DEFAULT ''::text NOT NULL,
    get_previous boolean DEFAULT false NOT NULL
);


--
-- Name: levels_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.levels_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: levels_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.levels_id_seq OWNED BY public.levels.id;


--
-- Name: resolution_entity; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.resolution_entity (
    id integer NOT NULL,
    name text DEFAULT ''::text NOT NULL,
    added_at timestamp without time zone DEFAULT now() NOT NULL
);


--
-- Name: resolution_entity_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.resolution_entity_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: resolution_entity_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.resolution_entity_id_seq OWNED BY public.resolution_entity.id;


--
-- Name: resolution_files; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.resolution_files (
    id integer NOT NULL,
    resolutionid integer NOT NULL,
    filename text DEFAULT ''::text NOT NULL,
    added_at timestamp without time zone DEFAULT now() NOT NULL,
    ctype text DEFAULT ''::text NOT NULL
);


--
-- Name: resolution_files_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.resolution_files_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: resolution_files_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.resolution_files_id_seq OWNED BY public.resolution_files.id;


--
-- Name: resolutions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.resolutions (
    id integer NOT NULL,
    added_at timestamp without time zone DEFAULT now() NOT NULL,
    decisiondate date,
    sentdate date,
    startdate date NOT NULL,
    enddate date,
    entityid integer NOT NULL,
    name text DEFAULT ''::text NOT NULL
);


--
-- Name: resolutions_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.resolutions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: resolutions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.resolutions_id_seq OWNED BY public.resolutions.id;


--
-- Name: spawn_tasks; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.spawn_tasks (
    id integer NOT NULL,
    added_at timestamp without time zone DEFAULT now() NOT NULL,
    descr text DEFAULT ''::text NOT NULL,
    min_respawn_duration_after_done interval DEFAULT '08:00:00'::interval NOT NULL,
    title text DEFAULT ''::text NOT NULL,
    spawn_start_time time without time zone DEFAULT '06:00:00'::time without time zone NOT NULL,
    max_count_daily integer DEFAULT 1 NOT NULL,
    deleted_at timestamp without time zone,
    priority integer DEFAULT 0 NOT NULL
);


--
-- Name: tags_for_entry; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.tags_for_entry (
    entryid integer NOT NULL,
    tagid integer NOT NULL,
    added_at timestamp without time zone DEFAULT now() NOT NULL
);


--
-- Name: tags_for_task; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.tags_for_task (
    taskid integer NOT NULL,
    tagid integer NOT NULL
);


--
-- Name: tasks_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.tasks_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: tasks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.tasks_id_seq OWNED BY public.spawn_tasks.id;


--
-- Name: test_madrs; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.test_madrs (
    id integer NOT NULL,
    added_at timestamp without time zone DEFAULT now() NOT NULL,
    a1 integer NOT NULL,
    a2 integer NOT NULL,
    a3 integer NOT NULL,
    a4 integer NOT NULL,
    a5 integer NOT NULL,
    a6 integer NOT NULL,
    a7 integer NOT NULL,
    a8 integer NOT NULL,
    a9 integer NOT NULL,
    a10 integer NOT NULL,
    score integer NOT NULL
);


--
-- Name: test_madrs_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.test_madrs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: test_madrs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.test_madrs_id_seq OWNED BY public.test_madrs.id;


--
-- Name: version; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.version (
    ver integer DEFAULT 0 NOT NULL
);


--
-- Name: weight; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.weight (
    id integer NOT NULL,
    added_at timestamp without time zone DEFAULT now() NOT NULL,
    value numeric DEFAULT 0.0 NOT NULL
);


--
-- Name: weight_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.weight_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: weight_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.weight_id_seq OWNED BY public.weight.id;


--
-- Name: entries id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.entries ALTER COLUMN id SET DEFAULT nextval('public.entries_id_seq'::regclass);


--
-- Name: entry_tags id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.entry_tags ALTER COLUMN id SET DEFAULT nextval('public.entry_tags_id_seq'::regclass);


--
-- Name: height id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.height ALTER COLUMN id SET DEFAULT nextval('public.height_id_seq'::regclass);


--
-- Name: levels id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.levels ALTER COLUMN id SET DEFAULT nextval('public.levels_id_seq'::regclass);


--
-- Name: resolution_entity id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.resolution_entity ALTER COLUMN id SET DEFAULT nextval('public.resolution_entity_id_seq'::regclass);


--
-- Name: resolution_files id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.resolution_files ALTER COLUMN id SET DEFAULT nextval('public.resolution_files_id_seq'::regclass);


--
-- Name: resolutions id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.resolutions ALTER COLUMN id SET DEFAULT nextval('public.resolutions_id_seq'::regclass);


--
-- Name: spawn_tasks id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.spawn_tasks ALTER COLUMN id SET DEFAULT nextval('public.tasks_id_seq'::regclass);


--
-- Name: tasks id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tasks ALTER COLUMN id SET DEFAULT nextval('public.completed_tasks_id_seq'::regclass);


--
-- Name: test_madrs id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.test_madrs ALTER COLUMN id SET DEFAULT nextval('public.test_madrs_id_seq'::regclass);


--
-- Name: weight id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.weight ALTER COLUMN id SET DEFAULT nextval('public.weight_id_seq'::regclass);


--
-- Name: tasks completed_tasks_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT completed_tasks_pk PRIMARY KEY (id);


--
-- Name: entries entries_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.entries
    ADD CONSTRAINT entries_pk PRIMARY KEY (id);


--
-- Name: entry_levels entry_levels_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.entry_levels
    ADD CONSTRAINT entry_levels_pk PRIMARY KEY (entryid, levelid);


--
-- Name: entry_tags entry_tags_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.entry_tags
    ADD CONSTRAINT entry_tags_pk PRIMARY KEY (id);


--
-- Name: height height_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.height
    ADD CONSTRAINT height_pk PRIMARY KEY (id);


--
-- Name: levels levels_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.levels
    ADD CONSTRAINT levels_pk PRIMARY KEY (id);


--
-- Name: resolution_entity resolution_entity_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.resolution_entity
    ADD CONSTRAINT resolution_entity_pk PRIMARY KEY (id);


--
-- Name: resolution_files resolution_files_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.resolution_files
    ADD CONSTRAINT resolution_files_pk PRIMARY KEY (id);


--
-- Name: resolutions resolutions_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.resolutions
    ADD CONSTRAINT resolutions_pk PRIMARY KEY (id);


--
-- Name: tags_for_entry tags_for_entry_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tags_for_entry
    ADD CONSTRAINT tags_for_entry_pk PRIMARY KEY (entryid, tagid);


--
-- Name: tags_for_task tags_for_task_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tags_for_task
    ADD CONSTRAINT tags_for_task_pk PRIMARY KEY (taskid, tagid);


--
-- Name: spawn_tasks tasks_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.spawn_tasks
    ADD CONSTRAINT tasks_pk PRIMARY KEY (id);


--
-- Name: test_madrs test_madrs_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.test_madrs
    ADD CONSTRAINT test_madrs_pk PRIMARY KEY (id);


--
-- Name: weight weight_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.weight
    ADD CONSTRAINT weight_pk PRIMARY KEY (id);


--
-- Name: entries_fromtime_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX entries_fromtime_uindex ON public.entries USING btree (fromtime);


--
-- Name: entry_tags_name_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX entry_tags_name_uindex ON public.entry_tags USING btree (name);


--
-- Name: entry_tags_shortname_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX entry_tags_shortname_uindex ON public.entry_tags USING btree (shortname);


--
-- Name: levels_name_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX levels_name_uindex ON public.levels USING btree (name);


--
-- Name: levels_shortname_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX levels_shortname_uindex ON public.levels USING btree (shortname);


--
-- Name: resolution_entity_id_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX resolution_entity_id_uindex ON public.resolution_entity USING btree (id);


--
-- Name: resolution_entity_name_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX resolution_entity_name_uindex ON public.resolution_entity USING btree (name);


--
-- Name: resolution_files_filename_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX resolution_files_filename_uindex ON public.resolution_files USING btree (filename);


--
-- Name: resolution_files_id_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX resolution_files_id_uindex ON public.resolution_files USING btree (id);


--
-- Name: resolutions_id_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX resolutions_id_uindex ON public.resolutions USING btree (id);


--
-- Name: spawn_tasks_title_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX spawn_tasks_title_uindex ON public.spawn_tasks USING btree (title);


--
-- Name: test_madrs_id_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX test_madrs_id_uindex ON public.test_madrs USING btree (id);


--
-- Name: tasks completed_tasks_tasks_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT completed_tasks_tasks_id_fk FOREIGN KEY (refid) REFERENCES public.spawn_tasks(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: entry_levels entry_levels_entries_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.entry_levels
    ADD CONSTRAINT entry_levels_entries_id_fk FOREIGN KEY (entryid) REFERENCES public.entries(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: entry_levels entry_levels_levels_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.entry_levels
    ADD CONSTRAINT entry_levels_levels_id_fk FOREIGN KEY (levelid) REFERENCES public.levels(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: resolution_files resolution_files_resolutions_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.resolution_files
    ADD CONSTRAINT resolution_files_resolutions_id_fk FOREIGN KEY (resolutionid) REFERENCES public.resolutions(id);


--
-- Name: resolutions resolutions_resolution_entity_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.resolutions
    ADD CONSTRAINT resolutions_resolution_entity_id_fk FOREIGN KEY (entityid) REFERENCES public.resolution_entity(id);


--
-- Name: tags_for_task tags_for_task_entry_tags_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tags_for_task
    ADD CONSTRAINT tags_for_task_entry_tags_id_fk FOREIGN KEY (tagid) REFERENCES public.entry_tags(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: tags_for_task tags_for_task_tasks_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tags_for_task
    ADD CONSTRAINT tags_for_task_tasks_id_fk FOREIGN KEY (taskid) REFERENCES public.spawn_tasks(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

