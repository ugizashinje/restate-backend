--
-- PostgreSQL database dump
--

-- Dumped from database version 16.2 (Ubuntu 16.2-1.pgdg22.04+1)
-- Dumped by pg_dump version 16.2 (Ubuntu 16.2-1.pgdg22.04+1)

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
-- Name: emaildeliverytype; Type: TYPE; Schema: public; Owner: admin
--

CREATE TYPE public.emaildeliverytype AS ENUM (
    'attachment',
    'inline'
);


ALTER TYPE public.emaildeliverytype OWNER TO admin;

--
-- Name: objecttype; Type: TYPE; Schema: public; Owner: admin
--

CREATE TYPE public.objecttype AS ENUM (
    'query',
    'chart',
    'dashboard',
    'dataset'
);


ALTER TYPE public.objecttype OWNER TO admin;

--
-- Name: sliceemailreportformat; Type: TYPE; Schema: public; Owner: admin
--

CREATE TYPE public.sliceemailreportformat AS ENUM (
    'visualization',
    'data'
);


ALTER TYPE public.sliceemailreportformat OWNER TO admin;

--
-- Name: tagtype; Type: TYPE; Schema: public; Owner: admin
--

CREATE TYPE public.tagtype AS ENUM (
    'custom',
    'type',
    'owner',
    'favorited_by'
);


ALTER TYPE public.tagtype OWNER TO admin;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: ab_permission; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.ab_permission (
    id integer NOT NULL,
    name character varying(100) NOT NULL
);


ALTER TABLE public.ab_permission OWNER TO admin;

--
-- Name: ab_permission_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.ab_permission_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.ab_permission_id_seq OWNER TO admin;

--
-- Name: ab_permission_view; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.ab_permission_view (
    id integer NOT NULL,
    permission_id integer,
    view_menu_id integer
);


ALTER TABLE public.ab_permission_view OWNER TO admin;

--
-- Name: ab_permission_view_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.ab_permission_view_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.ab_permission_view_id_seq OWNER TO admin;

--
-- Name: ab_permission_view_role; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.ab_permission_view_role (
    id integer NOT NULL,
    permission_view_id integer,
    role_id integer
);


ALTER TABLE public.ab_permission_view_role OWNER TO admin;

--
-- Name: ab_permission_view_role_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.ab_permission_view_role_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.ab_permission_view_role_id_seq OWNER TO admin;

--
-- Name: ab_register_user; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.ab_register_user (
    id integer NOT NULL,
    first_name character varying(64) NOT NULL,
    last_name character varying(64) NOT NULL,
    username character varying(64) NOT NULL,
    password character varying(256),
    email character varying(64) NOT NULL,
    registration_date timestamp without time zone,
    registration_hash character varying(256)
);


ALTER TABLE public.ab_register_user OWNER TO admin;

--
-- Name: ab_register_user_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.ab_register_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.ab_register_user_id_seq OWNER TO admin;

--
-- Name: ab_role; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.ab_role (
    id integer NOT NULL,
    name character varying(64) NOT NULL
);


ALTER TABLE public.ab_role OWNER TO admin;

--
-- Name: ab_role_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.ab_role_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.ab_role_id_seq OWNER TO admin;

--
-- Name: ab_user; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.ab_user (
    id integer NOT NULL,
    first_name character varying(64) NOT NULL,
    last_name character varying(64) NOT NULL,
    username character varying(64) NOT NULL,
    password character varying(256),
    active boolean,
    email character varying(320) NOT NULL,
    last_login timestamp without time zone,
    login_count integer,
    fail_login_count integer,
    created_on timestamp without time zone,
    changed_on timestamp without time zone,
    created_by_fk integer,
    changed_by_fk integer
);


ALTER TABLE public.ab_user OWNER TO admin;

--
-- Name: ab_user_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.ab_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.ab_user_id_seq OWNER TO admin;

--
-- Name: ab_user_role; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.ab_user_role (
    id integer NOT NULL,
    user_id integer,
    role_id integer
);


ALTER TABLE public.ab_user_role OWNER TO admin;

--
-- Name: ab_user_role_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.ab_user_role_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.ab_user_role_id_seq OWNER TO admin;

--
-- Name: ab_view_menu; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.ab_view_menu (
    id integer NOT NULL,
    name character varying(255) NOT NULL
);


ALTER TABLE public.ab_view_menu OWNER TO admin;

--
-- Name: ab_view_menu_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.ab_view_menu_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.ab_view_menu_id_seq OWNER TO admin;

--
-- Name: alembic_version; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.alembic_version (
    version_num character varying(32) NOT NULL
);


ALTER TABLE public.alembic_version OWNER TO admin;

--
-- Name: annotation; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.annotation (
    created_on timestamp without time zone,
    changed_on timestamp without time zone,
    id integer NOT NULL,
    start_dttm timestamp without time zone,
    end_dttm timestamp without time zone,
    layer_id integer,
    short_descr character varying(500),
    long_descr text,
    changed_by_fk integer,
    created_by_fk integer,
    json_metadata text
);


ALTER TABLE public.annotation OWNER TO admin;

--
-- Name: annotation_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.annotation_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.annotation_id_seq OWNER TO admin;

--
-- Name: annotation_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.annotation_id_seq OWNED BY public.annotation.id;


--
-- Name: annotation_layer; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.annotation_layer (
    created_on timestamp without time zone,
    changed_on timestamp without time zone,
    id integer NOT NULL,
    name character varying(250),
    descr text,
    changed_by_fk integer,
    created_by_fk integer
);


ALTER TABLE public.annotation_layer OWNER TO admin;

--
-- Name: annotation_layer_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.annotation_layer_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.annotation_layer_id_seq OWNER TO admin;

--
-- Name: annotation_layer_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.annotation_layer_id_seq OWNED BY public.annotation_layer.id;


--
-- Name: cache_keys; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.cache_keys (
    id integer NOT NULL,
    cache_key character varying(256) NOT NULL,
    cache_timeout integer,
    datasource_uid character varying(64) NOT NULL,
    created_on timestamp without time zone
);


ALTER TABLE public.cache_keys OWNER TO admin;

--
-- Name: cache_keys_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.cache_keys_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.cache_keys_id_seq OWNER TO admin;

--
-- Name: cache_keys_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.cache_keys_id_seq OWNED BY public.cache_keys.id;


--
-- Name: css_templates; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.css_templates (
    created_on timestamp without time zone,
    changed_on timestamp without time zone,
    id integer NOT NULL,
    template_name character varying(250),
    css text,
    changed_by_fk integer,
    created_by_fk integer
);


ALTER TABLE public.css_templates OWNER TO admin;

--
-- Name: css_templates_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.css_templates_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.css_templates_id_seq OWNER TO admin;

--
-- Name: css_templates_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.css_templates_id_seq OWNED BY public.css_templates.id;


--
-- Name: dashboard_roles; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.dashboard_roles (
    id integer NOT NULL,
    role_id integer NOT NULL,
    dashboard_id integer
);


ALTER TABLE public.dashboard_roles OWNER TO admin;

--
-- Name: dashboard_roles_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.dashboard_roles_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.dashboard_roles_id_seq OWNER TO admin;

--
-- Name: dashboard_roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.dashboard_roles_id_seq OWNED BY public.dashboard_roles.id;


--
-- Name: dashboard_slices; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.dashboard_slices (
    id integer NOT NULL,
    dashboard_id integer,
    slice_id integer
);


ALTER TABLE public.dashboard_slices OWNER TO admin;

--
-- Name: dashboard_slices_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.dashboard_slices_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.dashboard_slices_id_seq OWNER TO admin;

--
-- Name: dashboard_slices_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.dashboard_slices_id_seq OWNED BY public.dashboard_slices.id;


--
-- Name: dashboard_user; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.dashboard_user (
    id integer NOT NULL,
    user_id integer,
    dashboard_id integer
);


ALTER TABLE public.dashboard_user OWNER TO admin;

--
-- Name: dashboard_user_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.dashboard_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.dashboard_user_id_seq OWNER TO admin;

--
-- Name: dashboard_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.dashboard_user_id_seq OWNED BY public.dashboard_user.id;


--
-- Name: dashboards; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.dashboards (
    created_on timestamp without time zone,
    changed_on timestamp without time zone,
    id integer NOT NULL,
    dashboard_title character varying(500),
    position_json text,
    created_by_fk integer,
    changed_by_fk integer,
    css text,
    description text,
    slug character varying(255),
    json_metadata text,
    published boolean,
    uuid uuid,
    certified_by text,
    certification_details text,
    is_managed_externally boolean DEFAULT false NOT NULL,
    external_url text
);


ALTER TABLE public.dashboards OWNER TO admin;

--
-- Name: dashboards_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.dashboards_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.dashboards_id_seq OWNER TO admin;

--
-- Name: dashboards_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.dashboards_id_seq OWNED BY public.dashboards.id;


--
-- Name: dbs; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.dbs (
    created_on timestamp without time zone,
    changed_on timestamp without time zone,
    id integer NOT NULL,
    database_name character varying(250) NOT NULL,
    sqlalchemy_uri character varying(1024) NOT NULL,
    created_by_fk integer,
    changed_by_fk integer,
    password bytea,
    cache_timeout integer,
    extra text,
    select_as_create_table_as boolean,
    allow_ctas boolean,
    expose_in_sqllab boolean,
    force_ctas_schema character varying(250),
    allow_run_async boolean,
    allow_dml boolean,
    verbose_name character varying(250),
    impersonate_user boolean,
    allow_file_upload boolean DEFAULT true NOT NULL,
    encrypted_extra bytea,
    server_cert bytea,
    allow_cvas boolean,
    uuid uuid,
    configuration_method character varying(255) DEFAULT 'sqlalchemy_form'::character varying,
    is_managed_externally boolean DEFAULT false NOT NULL,
    external_url text
);


ALTER TABLE public.dbs OWNER TO admin;

--
-- Name: dbs_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.dbs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.dbs_id_seq OWNER TO admin;

--
-- Name: dbs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.dbs_id_seq OWNED BY public.dbs.id;


--
-- Name: dynamic_plugin; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.dynamic_plugin (
    created_on timestamp without time zone,
    changed_on timestamp without time zone,
    id integer NOT NULL,
    name character varying(50) NOT NULL,
    key character varying(50) NOT NULL,
    bundle_url character varying(1000) NOT NULL,
    created_by_fk integer,
    changed_by_fk integer
);


ALTER TABLE public.dynamic_plugin OWNER TO admin;

--
-- Name: dynamic_plugin_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.dynamic_plugin_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.dynamic_plugin_id_seq OWNER TO admin;

--
-- Name: dynamic_plugin_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.dynamic_plugin_id_seq OWNED BY public.dynamic_plugin.id;


--
-- Name: embedded_dashboards; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.embedded_dashboards (
    created_on timestamp without time zone,
    changed_on timestamp without time zone,
    allow_domain_list text,
    uuid uuid,
    dashboard_id integer NOT NULL,
    changed_by_fk integer,
    created_by_fk integer
);


ALTER TABLE public.embedded_dashboards OWNER TO admin;

--
-- Name: favstar; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.favstar (
    id integer NOT NULL,
    user_id integer,
    class_name character varying(50),
    obj_id integer,
    dttm timestamp without time zone
);


ALTER TABLE public.favstar OWNER TO admin;

--
-- Name: favstar_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.favstar_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.favstar_id_seq OWNER TO admin;

--
-- Name: favstar_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.favstar_id_seq OWNED BY public.favstar.id;


--
-- Name: filter_sets; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.filter_sets (
    created_on timestamp without time zone,
    changed_on timestamp without time zone,
    id integer NOT NULL,
    name character varying(500) NOT NULL,
    description text,
    json_metadata text NOT NULL,
    owner_id integer NOT NULL,
    owner_type character varying(255) NOT NULL,
    dashboard_id integer NOT NULL,
    created_by_fk integer,
    changed_by_fk integer
);


ALTER TABLE public.filter_sets OWNER TO admin;

--
-- Name: filter_sets_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.filter_sets_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.filter_sets_id_seq OWNER TO admin;

--
-- Name: filter_sets_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.filter_sets_id_seq OWNED BY public.filter_sets.id;


--
-- Name: key_value; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.key_value (
    id integer NOT NULL,
    resource character varying(32) NOT NULL,
    value bytea NOT NULL,
    uuid uuid,
    created_on timestamp without time zone,
    created_by_fk integer,
    changed_on timestamp without time zone,
    changed_by_fk integer,
    expires_on timestamp without time zone
);


ALTER TABLE public.key_value OWNER TO admin;

--
-- Name: key_value_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.key_value_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.key_value_id_seq OWNER TO admin;

--
-- Name: key_value_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.key_value_id_seq OWNED BY public.key_value.id;


--
-- Name: keyvalue; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.keyvalue (
    id integer NOT NULL,
    value text NOT NULL
);


ALTER TABLE public.keyvalue OWNER TO admin;

--
-- Name: keyvalue_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.keyvalue_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.keyvalue_id_seq OWNER TO admin;

--
-- Name: keyvalue_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.keyvalue_id_seq OWNED BY public.keyvalue.id;


--
-- Name: logs; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.logs (
    id integer NOT NULL,
    action character varying(512),
    user_id integer,
    json text,
    dttm timestamp without time zone,
    dashboard_id integer,
    slice_id integer,
    duration_ms integer,
    referrer character varying(1024)
);


ALTER TABLE public.logs OWNER TO admin;

--
-- Name: logs_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.logs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.logs_id_seq OWNER TO admin;

--
-- Name: logs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.logs_id_seq OWNED BY public.logs.id;


--
-- Name: query; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.query (
    id integer NOT NULL,
    client_id character varying(11) NOT NULL,
    database_id integer NOT NULL,
    tmp_table_name character varying(256),
    tab_name character varying(256),
    sql_editor_id character varying(256),
    user_id integer,
    status character varying(16),
    schema character varying(256),
    sql text,
    select_sql text,
    executed_sql text,
    "limit" integer,
    select_as_cta boolean,
    select_as_cta_used boolean,
    progress integer,
    rows integer,
    error_message text,
    start_time numeric(20,6),
    changed_on timestamp without time zone,
    end_time numeric(20,6),
    results_key character varying(64),
    start_running_time numeric(20,6),
    end_result_backend_time numeric(20,6),
    tracking_url text,
    extra_json text,
    tmp_schema_name character varying(256),
    ctas_method character varying(16),
    limiting_factor character varying(255) DEFAULT 'UNKNOWN'::character varying
);


ALTER TABLE public.query OWNER TO admin;

--
-- Name: query_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.query_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.query_id_seq OWNER TO admin;

--
-- Name: query_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.query_id_seq OWNED BY public.query.id;


--
-- Name: report_execution_log; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.report_execution_log (
    id integer NOT NULL,
    scheduled_dttm timestamp without time zone NOT NULL,
    start_dttm timestamp without time zone,
    end_dttm timestamp without time zone,
    value double precision,
    value_row_json text,
    state character varying(50) NOT NULL,
    error_message text,
    report_schedule_id integer NOT NULL,
    uuid uuid
);


ALTER TABLE public.report_execution_log OWNER TO admin;

--
-- Name: report_execution_log_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.report_execution_log_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.report_execution_log_id_seq OWNER TO admin;

--
-- Name: report_execution_log_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.report_execution_log_id_seq OWNED BY public.report_execution_log.id;


--
-- Name: report_recipient; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.report_recipient (
    id integer NOT NULL,
    type character varying(50) NOT NULL,
    recipient_config_json text,
    report_schedule_id integer NOT NULL,
    created_on timestamp without time zone,
    changed_on timestamp without time zone,
    created_by_fk integer,
    changed_by_fk integer
);


ALTER TABLE public.report_recipient OWNER TO admin;

--
-- Name: report_recipient_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.report_recipient_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.report_recipient_id_seq OWNER TO admin;

--
-- Name: report_recipient_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.report_recipient_id_seq OWNED BY public.report_recipient.id;


--
-- Name: report_schedule; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.report_schedule (
    id integer NOT NULL,
    type character varying(50) NOT NULL,
    name character varying(150) NOT NULL,
    description text,
    context_markdown text,
    active boolean,
    crontab character varying(1000) NOT NULL,
    sql text,
    chart_id integer,
    dashboard_id integer,
    database_id integer,
    last_eval_dttm timestamp without time zone,
    last_state character varying(50),
    last_value double precision,
    last_value_row_json text,
    validator_type character varying(100),
    validator_config_json text,
    log_retention integer,
    grace_period integer,
    created_on timestamp without time zone,
    changed_on timestamp without time zone,
    created_by_fk integer,
    changed_by_fk integer,
    working_timeout integer,
    report_format character varying(50) DEFAULT 'PNG'::character varying,
    creation_method character varying(255) DEFAULT 'alerts_reports'::character varying,
    timezone character varying(100) DEFAULT 'UTC'::character varying NOT NULL,
    extra_json text NOT NULL,
    force_screenshot boolean,
    custom_width integer,
    custom_height integer
);


ALTER TABLE public.report_schedule OWNER TO admin;

--
-- Name: report_schedule_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.report_schedule_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.report_schedule_id_seq OWNER TO admin;

--
-- Name: report_schedule_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.report_schedule_id_seq OWNED BY public.report_schedule.id;


--
-- Name: report_schedule_user; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.report_schedule_user (
    id integer NOT NULL,
    user_id integer NOT NULL,
    report_schedule_id integer NOT NULL
);


ALTER TABLE public.report_schedule_user OWNER TO admin;

--
-- Name: report_schedule_user_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.report_schedule_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.report_schedule_user_id_seq OWNER TO admin;

--
-- Name: report_schedule_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.report_schedule_user_id_seq OWNED BY public.report_schedule_user.id;


--
-- Name: rls_filter_roles; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.rls_filter_roles (
    id integer NOT NULL,
    role_id integer NOT NULL,
    rls_filter_id integer
);


ALTER TABLE public.rls_filter_roles OWNER TO admin;

--
-- Name: rls_filter_roles_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.rls_filter_roles_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.rls_filter_roles_id_seq OWNER TO admin;

--
-- Name: rls_filter_roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.rls_filter_roles_id_seq OWNED BY public.rls_filter_roles.id;


--
-- Name: rls_filter_tables; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.rls_filter_tables (
    id integer NOT NULL,
    table_id integer,
    rls_filter_id integer
);


ALTER TABLE public.rls_filter_tables OWNER TO admin;

--
-- Name: rls_filter_tables_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.rls_filter_tables_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.rls_filter_tables_id_seq OWNER TO admin;

--
-- Name: rls_filter_tables_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.rls_filter_tables_id_seq OWNED BY public.rls_filter_tables.id;


--
-- Name: row_level_security_filters; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.row_level_security_filters (
    created_on timestamp without time zone,
    changed_on timestamp without time zone,
    id integer NOT NULL,
    clause text NOT NULL,
    created_by_fk integer,
    changed_by_fk integer,
    filter_type character varying(255),
    group_key character varying(255),
    name character varying(255) NOT NULL,
    description text
);


ALTER TABLE public.row_level_security_filters OWNER TO admin;

--
-- Name: row_level_security_filters_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.row_level_security_filters_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.row_level_security_filters_id_seq OWNER TO admin;

--
-- Name: row_level_security_filters_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.row_level_security_filters_id_seq OWNED BY public.row_level_security_filters.id;


--
-- Name: saved_query; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.saved_query (
    created_on timestamp without time zone,
    changed_on timestamp without time zone,
    id integer NOT NULL,
    user_id integer,
    db_id integer,
    label character varying(256),
    schema character varying(128),
    sql text,
    description text,
    changed_by_fk integer,
    created_by_fk integer,
    extra_json text,
    last_run timestamp without time zone,
    rows integer,
    uuid uuid,
    template_parameters text
);


ALTER TABLE public.saved_query OWNER TO admin;

--
-- Name: saved_query_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.saved_query_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.saved_query_id_seq OWNER TO admin;

--
-- Name: saved_query_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.saved_query_id_seq OWNED BY public.saved_query.id;


--
-- Name: sl_columns; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.sl_columns (
    uuid uuid,
    created_on timestamp without time zone,
    changed_on timestamp without time zone,
    id integer NOT NULL,
    is_aggregation boolean NOT NULL,
    is_additive boolean NOT NULL,
    is_dimensional boolean NOT NULL,
    is_filterable boolean NOT NULL,
    is_increase_desired boolean NOT NULL,
    is_managed_externally boolean NOT NULL,
    is_partition boolean NOT NULL,
    is_physical boolean NOT NULL,
    is_temporal boolean NOT NULL,
    is_spatial boolean NOT NULL,
    name text,
    type text,
    unit text,
    expression text,
    description text,
    warning_text text,
    external_url text,
    extra_json text,
    created_by_fk integer,
    changed_by_fk integer,
    advanced_data_type text
);


ALTER TABLE public.sl_columns OWNER TO admin;

--
-- Name: sl_columns_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.sl_columns_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sl_columns_id_seq OWNER TO admin;

--
-- Name: sl_columns_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.sl_columns_id_seq OWNED BY public.sl_columns.id;


--
-- Name: sl_dataset_columns; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.sl_dataset_columns (
    dataset_id integer NOT NULL,
    column_id integer NOT NULL
);


ALTER TABLE public.sl_dataset_columns OWNER TO admin;

--
-- Name: sl_dataset_tables; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.sl_dataset_tables (
    dataset_id integer NOT NULL,
    table_id integer NOT NULL
);


ALTER TABLE public.sl_dataset_tables OWNER TO admin;

--
-- Name: sl_dataset_users; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.sl_dataset_users (
    dataset_id integer NOT NULL,
    user_id integer NOT NULL
);


ALTER TABLE public.sl_dataset_users OWNER TO admin;

--
-- Name: sl_datasets; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.sl_datasets (
    uuid uuid,
    created_on timestamp without time zone,
    changed_on timestamp without time zone,
    id integer NOT NULL,
    database_id integer NOT NULL,
    is_physical boolean,
    is_managed_externally boolean NOT NULL,
    name text,
    expression text,
    external_url text,
    extra_json text,
    created_by_fk integer,
    changed_by_fk integer
);


ALTER TABLE public.sl_datasets OWNER TO admin;

--
-- Name: sl_datasets_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.sl_datasets_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sl_datasets_id_seq OWNER TO admin;

--
-- Name: sl_datasets_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.sl_datasets_id_seq OWNED BY public.sl_datasets.id;


--
-- Name: sl_table_columns; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.sl_table_columns (
    table_id integer NOT NULL,
    column_id integer NOT NULL
);


ALTER TABLE public.sl_table_columns OWNER TO admin;

--
-- Name: sl_tables; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.sl_tables (
    uuid uuid,
    created_on timestamp without time zone,
    changed_on timestamp without time zone,
    id integer NOT NULL,
    database_id integer NOT NULL,
    is_managed_externally boolean NOT NULL,
    catalog text,
    schema text,
    name text,
    external_url text,
    extra_json text,
    created_by_fk integer,
    changed_by_fk integer
);


ALTER TABLE public.sl_tables OWNER TO admin;

--
-- Name: sl_tables_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.sl_tables_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sl_tables_id_seq OWNER TO admin;

--
-- Name: sl_tables_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.sl_tables_id_seq OWNED BY public.sl_tables.id;


--
-- Name: slice_user; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.slice_user (
    id integer NOT NULL,
    user_id integer,
    slice_id integer
);


ALTER TABLE public.slice_user OWNER TO admin;

--
-- Name: slice_user_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.slice_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.slice_user_id_seq OWNER TO admin;

--
-- Name: slice_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.slice_user_id_seq OWNED BY public.slice_user.id;


--
-- Name: slices; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.slices (
    created_on timestamp without time zone,
    changed_on timestamp without time zone,
    id integer NOT NULL,
    slice_name character varying(250),
    datasource_type character varying(200),
    datasource_name character varying(2000),
    viz_type character varying(250),
    params text,
    created_by_fk integer,
    changed_by_fk integer,
    description text,
    cache_timeout integer,
    perm character varying(2000),
    datasource_id integer,
    schema_perm character varying(1000),
    uuid uuid,
    query_context text,
    last_saved_at timestamp without time zone,
    last_saved_by_fk integer,
    certified_by text,
    certification_details text,
    is_managed_externally boolean DEFAULT false NOT NULL,
    external_url text,
    CONSTRAINT ck_chart_datasource CHECK (((datasource_type)::text = 'table'::text))
);


ALTER TABLE public.slices OWNER TO admin;

--
-- Name: slices_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.slices_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.slices_id_seq OWNER TO admin;

--
-- Name: slices_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.slices_id_seq OWNED BY public.slices.id;


--
-- Name: sql_metrics; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.sql_metrics (
    created_on timestamp without time zone,
    changed_on timestamp without time zone,
    id integer NOT NULL,
    metric_name character varying(255) NOT NULL,
    verbose_name character varying(1024),
    metric_type character varying(32),
    table_id integer,
    expression text NOT NULL,
    description text,
    created_by_fk integer,
    changed_by_fk integer,
    d3format character varying(128),
    warning_text text,
    extra text,
    uuid uuid,
    currency character varying(128)
);


ALTER TABLE public.sql_metrics OWNER TO admin;

--
-- Name: sql_metrics_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.sql_metrics_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sql_metrics_id_seq OWNER TO admin;

--
-- Name: sql_metrics_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.sql_metrics_id_seq OWNED BY public.sql_metrics.id;


--
-- Name: sqlatable_user; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.sqlatable_user (
    id integer NOT NULL,
    user_id integer,
    table_id integer
);


ALTER TABLE public.sqlatable_user OWNER TO admin;

--
-- Name: sqlatable_user_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.sqlatable_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sqlatable_user_id_seq OWNER TO admin;

--
-- Name: sqlatable_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.sqlatable_user_id_seq OWNED BY public.sqlatable_user.id;


--
-- Name: ssh_tunnels; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.ssh_tunnels (
    created_on timestamp without time zone,
    changed_on timestamp without time zone,
    created_by_fk integer,
    changed_by_fk integer,
    extra_json text,
    uuid uuid,
    id integer NOT NULL,
    database_id integer,
    server_address character varying(256),
    server_port integer,
    username bytea,
    password bytea,
    private_key bytea,
    private_key_password bytea
);


ALTER TABLE public.ssh_tunnels OWNER TO admin;

--
-- Name: ssh_tunnels_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.ssh_tunnels_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.ssh_tunnels_id_seq OWNER TO admin;

--
-- Name: ssh_tunnels_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.ssh_tunnels_id_seq OWNED BY public.ssh_tunnels.id;


--
-- Name: tab_state; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.tab_state (
    created_on timestamp without time zone,
    changed_on timestamp without time zone,
    extra_json text,
    id integer NOT NULL,
    user_id integer,
    label character varying(256),
    active boolean,
    database_id integer,
    schema character varying(256),
    sql text,
    query_limit integer,
    latest_query_id character varying(11),
    autorun boolean NOT NULL,
    template_params text,
    created_by_fk integer,
    changed_by_fk integer,
    hide_left_bar boolean DEFAULT false NOT NULL,
    saved_query_id integer
);


ALTER TABLE public.tab_state OWNER TO admin;

--
-- Name: tab_state_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.tab_state_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.tab_state_id_seq OWNER TO admin;

--
-- Name: tab_state_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.tab_state_id_seq OWNED BY public.tab_state.id;


--
-- Name: table_columns; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.table_columns (
    created_on timestamp without time zone,
    changed_on timestamp without time zone,
    id integer NOT NULL,
    table_id integer,
    column_name character varying(255) NOT NULL,
    is_dttm boolean,
    is_active boolean,
    type text,
    groupby boolean,
    filterable boolean,
    description text,
    created_by_fk integer,
    changed_by_fk integer,
    expression text,
    verbose_name character varying(1024),
    python_date_format character varying(255),
    uuid uuid,
    extra text,
    advanced_data_type character varying(255)
);


ALTER TABLE public.table_columns OWNER TO admin;

--
-- Name: table_columns_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.table_columns_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.table_columns_id_seq OWNER TO admin;

--
-- Name: table_columns_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.table_columns_id_seq OWNED BY public.table_columns.id;


--
-- Name: table_schema; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.table_schema (
    created_on timestamp without time zone,
    changed_on timestamp without time zone,
    extra_json text,
    id integer NOT NULL,
    tab_state_id integer,
    database_id integer NOT NULL,
    schema character varying(256),
    "table" character varying(256),
    description text,
    expanded boolean,
    created_by_fk integer,
    changed_by_fk integer
);


ALTER TABLE public.table_schema OWNER TO admin;

--
-- Name: table_schema_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.table_schema_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.table_schema_id_seq OWNER TO admin;

--
-- Name: table_schema_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.table_schema_id_seq OWNED BY public.table_schema.id;


--
-- Name: tables; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.tables (
    created_on timestamp without time zone,
    changed_on timestamp without time zone,
    id integer NOT NULL,
    table_name character varying(250) NOT NULL,
    main_dttm_col character varying(250),
    default_endpoint text,
    database_id integer NOT NULL,
    created_by_fk integer,
    changed_by_fk integer,
    "offset" integer,
    description text,
    is_featured boolean,
    cache_timeout integer,
    schema character varying(255),
    sql text,
    params text,
    perm character varying(1000),
    filter_select_enabled boolean,
    fetch_values_predicate text,
    is_sqllab_view boolean DEFAULT false,
    template_params text,
    schema_perm character varying(1000),
    extra text,
    uuid uuid,
    is_managed_externally boolean DEFAULT false NOT NULL,
    external_url text,
    normalize_columns boolean DEFAULT false,
    always_filter_main_dttm boolean DEFAULT false
);


ALTER TABLE public.tables OWNER TO admin;

--
-- Name: tables_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.tables_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.tables_id_seq OWNER TO admin;

--
-- Name: tables_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.tables_id_seq OWNED BY public.tables.id;


--
-- Name: tag; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.tag (
    created_on timestamp without time zone,
    changed_on timestamp without time zone,
    id integer NOT NULL,
    name character varying(250),
    type character varying,
    created_by_fk integer,
    changed_by_fk integer,
    description text
);


ALTER TABLE public.tag OWNER TO admin;

--
-- Name: tag_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.tag_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.tag_id_seq OWNER TO admin;

--
-- Name: tag_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.tag_id_seq OWNED BY public.tag.id;


--
-- Name: tagged_object; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.tagged_object (
    created_on timestamp without time zone,
    changed_on timestamp without time zone,
    id integer NOT NULL,
    tag_id integer,
    object_id integer,
    object_type character varying,
    created_by_fk integer,
    changed_by_fk integer
);


ALTER TABLE public.tagged_object OWNER TO admin;

--
-- Name: tagged_object_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.tagged_object_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.tagged_object_id_seq OWNER TO admin;

--
-- Name: tagged_object_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.tagged_object_id_seq OWNED BY public.tagged_object.id;


--
-- Name: url; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.url (
    created_on timestamp without time zone,
    changed_on timestamp without time zone,
    id integer NOT NULL,
    url text,
    created_by_fk integer,
    changed_by_fk integer
);


ALTER TABLE public.url OWNER TO admin;

--
-- Name: url_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.url_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.url_id_seq OWNER TO admin;

--
-- Name: url_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.url_id_seq OWNED BY public.url.id;


--
-- Name: user_attribute; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.user_attribute (
    created_on timestamp without time zone,
    changed_on timestamp without time zone,
    id integer NOT NULL,
    user_id integer,
    welcome_dashboard_id integer,
    created_by_fk integer,
    changed_by_fk integer
);


ALTER TABLE public.user_attribute OWNER TO admin;

--
-- Name: user_attribute_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.user_attribute_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.user_attribute_id_seq OWNER TO admin;

--
-- Name: user_attribute_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.user_attribute_id_seq OWNED BY public.user_attribute.id;


--
-- Name: user_favorite_tag; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.user_favorite_tag (
    user_id integer NOT NULL,
    tag_id integer NOT NULL
);


ALTER TABLE public.user_favorite_tag OWNER TO admin;

--
-- Name: annotation id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.annotation ALTER COLUMN id SET DEFAULT nextval('public.annotation_id_seq'::regclass);


--
-- Name: annotation_layer id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.annotation_layer ALTER COLUMN id SET DEFAULT nextval('public.annotation_layer_id_seq'::regclass);


--
-- Name: cache_keys id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.cache_keys ALTER COLUMN id SET DEFAULT nextval('public.cache_keys_id_seq'::regclass);


--
-- Name: css_templates id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.css_templates ALTER COLUMN id SET DEFAULT nextval('public.css_templates_id_seq'::regclass);


--
-- Name: dashboard_roles id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dashboard_roles ALTER COLUMN id SET DEFAULT nextval('public.dashboard_roles_id_seq'::regclass);


--
-- Name: dashboard_slices id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dashboard_slices ALTER COLUMN id SET DEFAULT nextval('public.dashboard_slices_id_seq'::regclass);


--
-- Name: dashboard_user id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dashboard_user ALTER COLUMN id SET DEFAULT nextval('public.dashboard_user_id_seq'::regclass);


--
-- Name: dashboards id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dashboards ALTER COLUMN id SET DEFAULT nextval('public.dashboards_id_seq'::regclass);


--
-- Name: dbs id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dbs ALTER COLUMN id SET DEFAULT nextval('public.dbs_id_seq'::regclass);


--
-- Name: dynamic_plugin id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dynamic_plugin ALTER COLUMN id SET DEFAULT nextval('public.dynamic_plugin_id_seq'::regclass);


--
-- Name: favstar id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.favstar ALTER COLUMN id SET DEFAULT nextval('public.favstar_id_seq'::regclass);


--
-- Name: filter_sets id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.filter_sets ALTER COLUMN id SET DEFAULT nextval('public.filter_sets_id_seq'::regclass);


--
-- Name: key_value id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.key_value ALTER COLUMN id SET DEFAULT nextval('public.key_value_id_seq'::regclass);


--
-- Name: keyvalue id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.keyvalue ALTER COLUMN id SET DEFAULT nextval('public.keyvalue_id_seq'::regclass);


--
-- Name: logs id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.logs ALTER COLUMN id SET DEFAULT nextval('public.logs_id_seq'::regclass);


--
-- Name: query id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.query ALTER COLUMN id SET DEFAULT nextval('public.query_id_seq'::regclass);


--
-- Name: report_execution_log id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.report_execution_log ALTER COLUMN id SET DEFAULT nextval('public.report_execution_log_id_seq'::regclass);


--
-- Name: report_recipient id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.report_recipient ALTER COLUMN id SET DEFAULT nextval('public.report_recipient_id_seq'::regclass);


--
-- Name: report_schedule id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.report_schedule ALTER COLUMN id SET DEFAULT nextval('public.report_schedule_id_seq'::regclass);


--
-- Name: report_schedule_user id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.report_schedule_user ALTER COLUMN id SET DEFAULT nextval('public.report_schedule_user_id_seq'::regclass);


--
-- Name: rls_filter_roles id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.rls_filter_roles ALTER COLUMN id SET DEFAULT nextval('public.rls_filter_roles_id_seq'::regclass);


--
-- Name: rls_filter_tables id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.rls_filter_tables ALTER COLUMN id SET DEFAULT nextval('public.rls_filter_tables_id_seq'::regclass);


--
-- Name: row_level_security_filters id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.row_level_security_filters ALTER COLUMN id SET DEFAULT nextval('public.row_level_security_filters_id_seq'::regclass);


--
-- Name: saved_query id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.saved_query ALTER COLUMN id SET DEFAULT nextval('public.saved_query_id_seq'::regclass);


--
-- Name: sl_columns id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sl_columns ALTER COLUMN id SET DEFAULT nextval('public.sl_columns_id_seq'::regclass);


--
-- Name: sl_datasets id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sl_datasets ALTER COLUMN id SET DEFAULT nextval('public.sl_datasets_id_seq'::regclass);


--
-- Name: sl_tables id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sl_tables ALTER COLUMN id SET DEFAULT nextval('public.sl_tables_id_seq'::regclass);


--
-- Name: slice_user id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.slice_user ALTER COLUMN id SET DEFAULT nextval('public.slice_user_id_seq'::regclass);


--
-- Name: slices id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.slices ALTER COLUMN id SET DEFAULT nextval('public.slices_id_seq'::regclass);


--
-- Name: sql_metrics id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sql_metrics ALTER COLUMN id SET DEFAULT nextval('public.sql_metrics_id_seq'::regclass);


--
-- Name: sqlatable_user id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sqlatable_user ALTER COLUMN id SET DEFAULT nextval('public.sqlatable_user_id_seq'::regclass);


--
-- Name: ssh_tunnels id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ssh_tunnels ALTER COLUMN id SET DEFAULT nextval('public.ssh_tunnels_id_seq'::regclass);


--
-- Name: tab_state id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tab_state ALTER COLUMN id SET DEFAULT nextval('public.tab_state_id_seq'::regclass);


--
-- Name: table_columns id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.table_columns ALTER COLUMN id SET DEFAULT nextval('public.table_columns_id_seq'::regclass);


--
-- Name: table_schema id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.table_schema ALTER COLUMN id SET DEFAULT nextval('public.table_schema_id_seq'::regclass);


--
-- Name: tables id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tables ALTER COLUMN id SET DEFAULT nextval('public.tables_id_seq'::regclass);


--
-- Name: tag id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tag ALTER COLUMN id SET DEFAULT nextval('public.tag_id_seq'::regclass);


--
-- Name: tagged_object id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tagged_object ALTER COLUMN id SET DEFAULT nextval('public.tagged_object_id_seq'::regclass);


--
-- Name: url id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.url ALTER COLUMN id SET DEFAULT nextval('public.url_id_seq'::regclass);


--
-- Name: user_attribute id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.user_attribute ALTER COLUMN id SET DEFAULT nextval('public.user_attribute_id_seq'::regclass);


--
-- Data for Name: ab_permission; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.ab_permission (id, name) FROM stdin;
1	can_read
2	can_write
3	can_this_form_post
4	can_this_form_get
5	can_delete
6	can_show
7	can_userinfo
8	can_edit
9	can_list
10	can_add
11	resetmypassword
12	resetpasswords
13	userinfoedit
14	copyrole
15	can_get
16	can_invalidate
17	can_warm_up_cache
18	can_export
19	can_get_embedded
20	can_delete_embedded
21	can_set_embedded
22	can_duplicate
23	can_get_or_create_dataset
24	can_get_column_values
25	can_import_
26	can_bulk_create
27	can_execute_sql_query
28	can_format_sql
29	can_export_csv
30	can_get_results
31	can_estimate_query_cost
32	can_download
33	can_time_range
34	can_query_form_data
35	can_query
36	can_save
37	can_samples
38	can_external_metadata
39	can_external_metadata_by_name
40	can_store
41	can_get_value
42	can_my_queries
43	can_sqllab
44	can_fetch_datasource_metadata
45	can_dashboard
46	can_import_dashboards
47	can_sqllab_history
48	can_explore
49	can_explore_json
50	can_slice
51	can_dashboard_permalink
52	can_profile
53	can_log
54	can_post
55	can_expanded
56	can_put
57	can_activate
58	can_delete_query
59	can_migrate_query
60	can_tags
61	can_recent_activity
62	can_grant_guest_token
63	menu_access
64	all_datasource_access
65	all_database_access
66	all_query_access
67	can_csv
68	can_share_dashboard
69	can_share_chart
70	database_access
71	schema_access
72	datasource_access
\.


--
-- Data for Name: ab_permission_view; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.ab_permission_view (id, permission_id, view_menu_id) FROM stdin;
1	1	1
2	2	1
3	1	2
4	2	2
5	1	3
6	2	3
7	1	4
8	2	4
9	1	5
10	2	5
11	1	6
12	2	6
13	1	7
14	2	7
15	1	8
16	2	8
17	1	9
18	2	9
19	1	10
20	3	15
21	4	15
22	3	16
23	4	16
24	3	17
25	4	17
26	5	19
27	6	19
28	7	19
29	8	19
30	9	19
31	10	19
32	11	19
33	12	19
34	13	19
35	5	20
36	6	20
37	8	20
38	9	20
39	10	20
40	14	20
41	15	21
42	6	22
43	15	23
44	9	24
45	1	25
46	1	26
47	16	27
48	17	4
49	18	4
50	1	29
51	2	29
52	1	30
53	2	30
54	19	8
55	20	8
56	21	8
57	18	8
58	18	9
59	22	6
60	23	6
61	17	6
62	18	6
63	24	31
64	1	32
65	1	33
66	1	34
67	2	34
68	1	35
69	2	35
70	9	36
71	10	36
72	8	36
73	5	36
74	25	37
75	18	37
76	1	38
77	2	38
78	18	1
79	1	39
80	2	39
81	26	39
82	27	40
83	28	40
84	29	40
85	30	40
86	1	40
87	31	40
88	2	41
89	5	41
90	6	41
91	32	41
92	8	41
93	9	41
94	10	41
95	33	42
96	34	42
97	35	42
98	3	43
99	4	43
100	3	44
101	4	44
102	3	45
103	4	45
104	36	31
105	37	31
106	38	31
107	39	31
108	15	31
109	40	47
110	41	47
111	1	49
112	9	1
113	42	50
114	43	51
115	44	51
116	45	51
117	46	51
118	47	51
119	48	51
120	17	51
121	49	51
122	50	51
123	51	51
124	52	51
125	53	51
126	54	52
127	55	52
128	5	52
129	54	53
130	5	53
131	56	53
132	15	53
133	57	53
134	58	53
135	59	53
136	5	55
137	6	55
138	32	55
139	8	55
140	9	55
141	10	55
142	60	56
143	61	7
144	1	57
145	62	57
146	1	58
147	63	59
148	63	60
149	63	61
150	63	62
151	63	38
152	63	63
153	63	64
154	63	65
155	63	66
156	63	67
157	63	68
158	63	69
159	63	70
160	63	71
161	63	72
162	63	55
163	63	73
164	63	74
165	63	75
166	63	76
167	63	77
168	63	78
169	64	79
170	65	80
171	66	81
172	67	51
173	68	51
174	69	51
175	70	82
176	71	83
177	71	84
178	72	85
\.


--
-- Data for Name: ab_permission_view_role; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.ab_permission_view_role (id, permission_view_id, role_id) FROM stdin;
1	20	1
2	21	1
3	22	1
4	23	1
5	24	1
6	25	1
7	26	1
8	27	1
9	28	1
10	29	1
11	30	1
12	31	1
13	32	1
14	33	1
15	34	1
16	35	1
17	36	1
18	37	1
19	38	1
20	39	1
21	40	1
22	41	1
23	42	1
24	43	1
25	9	1
26	10	1
27	44	1
28	45	1
29	46	1
30	47	1
31	48	1
32	49	1
33	7	1
34	8	1
35	3	1
36	4	1
37	50	1
38	51	1
39	52	1
40	53	1
41	54	1
42	55	1
43	56	1
44	57	1
45	15	1
46	16	1
47	58	1
48	17	1
49	18	1
50	59	1
51	60	1
52	61	1
53	62	1
54	11	1
55	12	1
56	63	1
57	64	1
58	65	1
59	66	1
60	67	1
61	68	1
62	69	1
63	70	1
64	71	1
65	72	1
66	73	1
67	74	1
68	75	1
69	19	1
70	5	1
71	6	1
72	76	1
73	77	1
74	78	1
75	1	1
76	2	1
77	79	1
78	80	1
79	81	1
80	82	1
81	83	1
82	84	1
83	85	1
84	86	1
85	87	1
86	88	1
87	89	1
88	90	1
89	91	1
90	92	1
91	93	1
92	94	1
93	95	1
94	96	1
95	97	1
96	98	1
97	99	1
98	100	1
99	101	1
100	102	1
101	103	1
102	104	1
103	105	1
104	106	1
105	107	1
106	108	1
107	109	1
108	110	1
109	111	1
110	112	1
111	113	1
112	114	1
113	115	1
114	116	1
115	117	1
116	118	1
117	119	1
118	120	1
119	121	1
120	122	1
121	123	1
122	124	1
123	125	1
124	126	1
125	127	1
126	128	1
127	129	1
128	130	1
129	131	1
130	132	1
131	133	1
132	134	1
133	135	1
134	136	1
135	137	1
136	138	1
137	139	1
138	140	1
139	141	1
140	142	1
141	143	1
142	13	1
143	14	1
144	144	1
145	145	1
146	146	1
147	147	1
148	148	1
149	149	1
150	150	1
151	151	1
152	152	1
153	153	1
154	154	1
155	155	1
156	156	1
157	157	1
158	158	1
159	159	1
160	160	1
161	161	1
162	162	1
163	163	1
164	164	1
165	165	1
166	166	1
167	167	1
168	168	1
169	169	1
170	170	1
171	171	1
172	172	1
173	173	1
174	174	1
175	3	3
176	4	3
177	5	3
178	6	3
179	7	3
180	8	3
181	9	3
182	10	3
183	11	3
184	12	3
185	15	3
186	16	3
187	17	3
188	22	3
189	23	3
190	28	3
191	32	3
192	41	3
193	42	3
194	43	3
195	44	3
196	45	3
197	46	3
198	47	3
199	49	3
200	50	3
201	51	3
202	52	3
203	53	3
204	54	3
205	55	3
206	57	3
207	59	3
208	60	3
209	62	3
210	63	3
211	64	3
212	65	3
213	66	3
214	67	3
215	68	3
216	69	3
217	70	3
218	71	3
219	72	3
220	73	3
221	74	3
222	75	3
223	79	3
224	80	3
225	81	3
226	83	3
227	87	3
228	90	3
229	93	3
230	95	3
231	96	3
232	97	3
233	98	3
234	99	3
235	100	3
236	101	3
237	102	3
238	103	3
239	104	3
240	105	3
241	106	3
242	107	3
243	108	3
244	109	3
245	110	3
246	111	3
247	112	3
248	115	3
249	116	3
250	117	3
251	119	3
252	121	3
253	122	3
254	123	3
255	124	3
256	125	3
257	126	3
258	127	3
259	128	3
260	136	3
261	137	3
262	138	3
263	139	3
264	140	3
265	141	3
266	142	3
267	143	3
268	144	3
269	146	3
270	152	3
271	153	3
272	154	3
273	155	3
274	156	3
275	157	3
276	158	3
277	159	3
278	160	3
279	161	3
280	162	3
281	163	3
282	164	3
283	169	3
284	170	3
285	172	3
286	173	3
287	174	3
288	7	4
289	8	4
290	11	4
291	15	4
292	16	4
293	17	4
294	22	4
295	23	4
296	28	4
297	32	4
298	41	4
299	42	4
300	43	4
301	44	4
302	45	4
303	46	4
304	47	4
305	49	4
306	50	4
307	51	4
308	52	4
309	53	4
310	54	4
311	55	4
312	57	4
313	64	4
314	65	4
315	66	4
316	67	4
317	68	4
318	69	4
319	70	4
320	71	4
321	72	4
322	73	4
323	79	4
324	80	4
325	81	4
326	83	4
327	87	4
328	90	4
329	93	4
330	95	4
331	96	4
332	97	4
333	106	4
334	107	4
335	108	4
336	109	4
337	110	4
338	111	4
339	112	4
340	115	4
341	116	4
342	117	4
343	119	4
344	121	4
345	122	4
346	123	4
347	124	4
348	125	4
349	136	4
350	137	4
351	138	4
352	139	4
353	140	4
354	141	4
355	142	4
356	143	4
357	144	4
358	146	4
359	152	4
360	153	4
361	154	4
362	155	4
363	156	4
364	157	4
365	159	4
366	161	4
367	162	4
368	172	4
369	173	4
370	174	4
371	1	5
372	2	5
373	17	5
374	19	5
375	78	5
376	82	5
377	84	5
378	85	5
379	86	5
380	113	5
381	114	5
382	118	5
383	129	5
384	130	5
385	131	5
386	132	5
387	133	5
388	134	5
389	135	5
390	165	5
391	166	5
392	167	5
393	168	5
394	172	5
395	1	6
396	2	6
397	3	6
398	4	6
399	5	6
400	6	6
401	7	6
402	8	6
403	9	6
404	10	6
405	12	6
\.


--
-- Data for Name: ab_register_user; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.ab_register_user (id, first_name, last_name, username, password, email, registration_date, registration_hash) FROM stdin;
\.


--
-- Data for Name: ab_role; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.ab_role (id, name) FROM stdin;
1	Admin
2	Public
3	Alpha
4	Gamma
5	sql_lab
6	embed_role
\.


--
-- Data for Name: ab_user; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.ab_user (id, first_name, last_name, username, password, active, email, last_login, login_count, fail_login_count, created_on, changed_on, created_by_fk, changed_by_fk) FROM stdin;
2	first	second	emed_role	pbkdf2:sha256:600000$OLYtVl0pZUVy5vqY$8cba41024e58b47a2a5155c589765d031c0fec19fe1d58b1f0275b4ffbbb21f5	t	embed@embed.com	\N	\N	\N	2024-04-10 14:51:23.138802	2024-04-10 14:51:23.138811	1	1
1	Superset	Admin	admin	pbkdf2:sha256:600000$FPglVHle9z52WbzJ$54d830f634c3574440373e33470c14d9696852bf4385b0dc5a0384cea8469672	t	admin@superset.com	2024-04-11 02:34:43.202525	3	0	2024-04-10 14:48:31.566428	2024-04-10 14:48:31.566438	\N	\N
\.


--
-- Data for Name: ab_user_role; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.ab_user_role (id, user_id, role_id) FROM stdin;
1	1	1
2	2	6
\.


--
-- Data for Name: ab_view_menu; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.ab_view_menu (id, name) FROM stdin;
1	SavedQuery
2	CssTemplate
3	ReportSchedule
4	Chart
5	Annotation
6	Dataset
7	Log
8	Dashboard
9	Database
10	Query
11	SupersetIndexView
12	UtilView
13	LocaleView
14	SecurityApi
15	ResetPasswordView
16	ResetMyPasswordView
17	UserInfoEditView
18	AuthDBView
19	UserDBModelView
20	RoleModelView
21	OpenApi
22	SwaggerView
23	MenuApi
24	AsyncEventsRestApi
25	AdvancedDataType
26	AvailableDomains
27	CacheRestApi
28	CurrentUserRestApi
29	DashboardFilterStateRestApi
30	DashboardPermalinkRestApi
31	Datasource
32	EmbeddedDashboard
33	Explore
34	ExploreFormDataRestApi
35	ExplorePermalinkRestApi
36	FilterSets
37	ImportExportRestApi
38	Row Level Security
39	Tag
40	SQLLab
41	DynamicPlugin
42	Api
43	CsvToDatabaseView
44	ExcelToDatabaseView
45	ColumnarToDatabaseView
46	EmbeddedView
47	KV
48	R
49	Profile
50	SqlLab
51	Superset
52	TableSchemaView
53	TabStateView
54	TaggedObjectView
55	Tags
56	TagView
57	SecurityRestApi
58	RowLevelSecurity
59	Security
60	List Users
61	List Roles
62	Action Log
63	Home
64	Data
65	Databases
66	Dashboards
67	Charts
68	Datasets
69	Manage
70	Plugins
71	CSS Templates
72	Import Dashboards
73	Alerts & Report
74	Annotation Layers
75	SQL Lab
76	SQL Editor
77	Saved Queries
78	Query Search
79	all_datasource_access
80	all_database_access
81	all_query_access
82	[PostgreSQL].(id:1)
83	[PostgreSQL].[information_schema]
84	[PostgreSQL].[public]
85	[PostgreSQL].[warrants](id:1)
\.


--
-- Data for Name: alembic_version; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.alembic_version (version_num) FROM stdin;
b7851ee5522f
\.


--
-- Data for Name: annotation; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.annotation (created_on, changed_on, id, start_dttm, end_dttm, layer_id, short_descr, long_descr, changed_by_fk, created_by_fk, json_metadata) FROM stdin;
\.


--
-- Data for Name: annotation_layer; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.annotation_layer (created_on, changed_on, id, name, descr, changed_by_fk, created_by_fk) FROM stdin;
\.


--
-- Data for Name: cache_keys; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.cache_keys (id, cache_key, cache_timeout, datasource_uid, created_on) FROM stdin;
\.


--
-- Data for Name: css_templates; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.css_templates (created_on, changed_on, id, template_name, css, changed_by_fk, created_by_fk) FROM stdin;
\.


--
-- Data for Name: dashboard_roles; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.dashboard_roles (id, role_id, dashboard_id) FROM stdin;
\.


--
-- Data for Name: dashboard_slices; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.dashboard_slices (id, dashboard_id, slice_id) FROM stdin;
1	1	1
\.


--
-- Data for Name: dashboard_user; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.dashboard_user (id, user_id, dashboard_id) FROM stdin;
1	1	1
3	2	1
\.


--
-- Data for Name: dashboards; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.dashboards (created_on, changed_on, id, dashboard_title, position_json, created_by_fk, changed_by_fk, css, description, slug, json_metadata, published, uuid, certified_by, certification_details, is_managed_externally, external_url) FROM stdin;
2024-04-10 14:52:48.203502	2024-04-11 02:38:09.503615	1	putni	{"CHART-cVPGoqyVSl":{"children":[],"id":"CHART-cVPGoqyVSl","meta":{"chartId":1,"height":50,"sliceName":"putni","uuid":"7461d16f-603c-4e01-8d62-214b777cb5ef","width":4},"parents":["ROOT_ID","GRID_ID","ROW-A_rCqD9UVR"],"type":"CHART"},"DASHBOARD_VERSION_KEY":"v2","GRID_ID":{"children":["ROW-A_rCqD9UVR"],"id":"GRID_ID","parents":["ROOT_ID"],"type":"GRID"},"HEADER_ID":{"id":"HEADER_ID","meta":{"text":"putni"},"type":"HEADER"},"ROOT_ID":{"children":["GRID_ID"],"id":"ROOT_ID","type":"ROOT"},"ROW-A_rCqD9UVR":{"children":["CHART-cVPGoqyVSl"],"id":"ROW-A_rCqD9UVR","meta":{"background":"BACKGROUND_TRANSPARENT"},"parents":["ROOT_ID","GRID_ID"],"type":"ROW"}}	1	1		\N	\N	{"chart_configuration": {}, "global_chart_configuration": {"scope": {"rootPath": ["ROOT_ID"], "excluded": []}, "chartsInScope": [1]}, "color_scheme": "", "refresh_frequency": 0, "expanded_slices": {}, "label_colors": {}, "timed_refresh_immune_slices": [], "cross_filters_enabled": true, "default_filters": "{}", "shared_label_colors": {}, "color_scheme_domain": []}	f	806e4f3c-04de-4523-afc7-a8086665a195			f	\N
\.


--
-- Data for Name: dbs; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.dbs (created_on, changed_on, id, database_name, sqlalchemy_uri, created_by_fk, changed_by_fk, password, cache_timeout, extra, select_as_create_table_as, allow_ctas, expose_in_sqllab, force_ctas_schema, allow_run_async, allow_dml, verbose_name, impersonate_user, allow_file_upload, encrypted_extra, server_cert, allow_cvas, uuid, configuration_method, is_managed_externally, external_url) FROM stdin;
2024-04-10 14:51:50.50732	2024-04-10 14:51:52.237203	1	PostgreSQL	postgresql+psycopg2://admin:XXXXXXXXXX@postgres:5432/dev	1	1	\\x7a6f4135305343556a2b5a2b6a4b326f3442574f4d673d3d	\N	{"allows_virtual_table_explore":true}	f	f	t	\N	f	f	\N	f	f	\\x75567445473164326d49757a5233592b676871372b413d3d	\N	f	b0dbda74-a533-417d-a399-c48590bb224d	dynamic_form	f	\N
\.


--
-- Data for Name: dynamic_plugin; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.dynamic_plugin (created_on, changed_on, id, name, key, bundle_url, created_by_fk, changed_by_fk) FROM stdin;
\.


--
-- Data for Name: embedded_dashboards; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.embedded_dashboards (created_on, changed_on, allow_domain_list, uuid, dashboard_id, changed_by_fk, created_by_fk) FROM stdin;
\.


--
-- Data for Name: favstar; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.favstar (id, user_id, class_name, obj_id, dttm) FROM stdin;
\.


--
-- Data for Name: filter_sets; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.filter_sets (created_on, changed_on, id, name, description, json_metadata, owner_id, owner_type, dashboard_id, created_by_fk, changed_by_fk) FROM stdin;
\.


--
-- Data for Name: key_value; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.key_value (id, resource, value, uuid, created_on, created_by_fk, changed_on, changed_by_fk, expires_on) FROM stdin;
2	superset_metastore_cache	\\x22386a583839745462326856774d4175594974744856624271704c306a6f6c75484f4b74484c3947796d686173584b37514a61723248667461465f33784b57355522	8df384b0-bcab-3842-b6f4-4d4f8fe093b5	2024-04-10 14:52:18.127307	1	\N	\N	2024-04-17 14:52:18.125314
1	superset_metastore_cache	\\x7b226f776e6572223a20312c202264617461736f757263655f6964223a20312c202264617461736f757263655f74797065223a20227461626c65222c202263686172745f6964223a206e756c6c2c2022666f726d5f64617461223a20227b5c2264617461736f757263655c223a5c22315f5f7461626c655c222c5c2276697a5f747970655c223a5c226269675f6e756d6265725f746f74616c5c222c5c226d65747269635c223a5c22636f756e745c222c5c226164686f635f66696c746572735c223a5b7b5c22636c617573655c223a5c2257484552455c222c5c227375626a6563745c223a5c22637265617465645f61745c222c5c226f70657261746f725c223a5c2254454d504f52414c5f52414e47455c222c5c22636f6d70617261746f725c223a5c224e6f2066696c7465725c222c5c2265787072657373696f6e547970655c223a5c2253494d504c455c227d5d2c5c226865616465725f666f6e745f73697a655c223a302e342c5c227375626865616465725f666f6e745f73697a655c223a302e31352c5c22795f617869735f666f726d61745c223a5c22534d4152545f4e554d4245525c222c5c2274696d655f666f726d61745c223a5c22736d6172745f646174655c222c5c2265787472615f666f726d5f646174615c223a7b7d7d227d	4ed79c77-ba5a-3d08-a4ea-fe77e2173654	2024-04-10 14:52:18.113159	1	2024-04-10 14:52:29.122122	1	2024-04-17 14:52:29.119633
3	superset_metastore_cache	\\x7b226f776e6572223a20312c202264617461736f757263655f6964223a20312c202264617461736f757263655f74797065223a20227461626c65222c202263686172745f6964223a20312c2022666f726d5f64617461223a20227b5c2264617461736f757263655c223a5c22315f5f7461626c655c222c5c2276697a5f747970655c223a5c226269675f6e756d6265725f746f74616c5c222c5c22736c6963655f69645c223a312c5c226d65747269635c223a5c22636f756e745c222c5c226164686f635f66696c746572735c223a5b7b5c22636c617573655c223a5c2257484552455c222c5c22636f6d70617261746f725c223a5c224e6f2066696c7465725c222c5c2265787072657373696f6e547970655c223a5c2253494d504c455c222c5c226f70657261746f725c223a5c2254454d504f52414c5f52414e47455c222c5c227375626a6563745c223a5c22637265617465645f61745c227d5d2c5c226865616465725f666f6e745f73697a655c223a302e342c5c227375626865616465725f666f6e745f73697a655c223a302e31352c5c22795f617869735f666f726d61745c223a5c22534d4152545f4e554d4245525c222c5c2274696d655f666f726d61745c223a5c22736d6172745f646174655c222c5c2265787472615f666f726d5f646174615c223a7b7d7d227d	1c51f9ff-6655-319b-92b4-6f7628aeb395	2024-04-10 14:52:45.992643	1	\N	\N	2024-04-17 14:52:45.991698
4	superset_metastore_cache	\\x224a5037334a5575726e3463423442415950454f436577707564754b755f6b67553855715f35457833524142317159412d345f51625251336959536c596f446d5322	8db27ecf-e73d-3f45-8ec0-0cef23bb66cd	2024-04-10 14:52:46.01336	1	\N	\N	2024-04-17 14:52:46.012282
5	superset_metastore_cache	\\x7b226f776e6572223a20312c202276616c7565223a20227b7d227d	e24ac0c6-df62-3fee-91c3-16b174cd8c7c	2024-04-10 14:53:07.01222	1	2024-04-10 14:53:57.313124	1	2024-07-09 14:53:57.310342
6	superset_metastore_cache	\\x2278395176553257563445744f775a5757474f656368626d306b377a6f5a6830744447356959663039314f7064313245776c4b31525a5561356f6262626c57385122	e2cfdfc5-409d-361a-be65-26f162afb5bd	2024-04-10 14:53:07.03385	1	2024-04-10 14:53:57.32706	1	2024-07-09 14:53:57.324777
7	superset_metastore_cache	\\x7b226f776e6572223a20312c202276616c7565223a20227b7d227d	64930e0e-b114-3b6a-b351-67e726ce2f8e	2024-04-11 02:27:46.366661	1	2024-04-11 02:38:10.128503	1	2024-07-10 02:38:10.125516
8	superset_metastore_cache	\\x22636c5f75506b4f506e49683951464963354d4667444a6f74734d4878434c327339756a4c46387562415841314f563652535153686f754977684847315678505922	f0ab75f2-945d-34d9-957c-fd967c514ad2	2024-04-11 02:27:46.389721	1	2024-04-11 02:38:10.150208	1	2024-07-10 02:38:10.147237
9	app	\\x22714b764d3948774f56643876385f656a53354c543849777377635475306c747a54626d39796a4c73524938556e384c444c536259384e64535050767041725f3222	ca4e4125-7261-3c4b-ad75-9da40b3d6dda	2024-04-11 02:38:58.398694	1	\N	\N	\N
10	dashboard_permalink	\\x7b227374617465223a207b2261637469766554616273223a205b5d2c2022646174614d61736b223a207b7d2c202275726c506172616d73223a205b5d7d2c202264617368626f6172644964223a202238303665346633632d303464652d343532332d616663372d613830383636363561313935227d	7de6621e-c603-3ebd-992d-e05ccfdb50b8	2024-04-11 02:38:58.406743	1	\N	\N	\N
\.


--
-- Data for Name: keyvalue; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.keyvalue (id, value) FROM stdin;
\.


--
-- Data for Name: logs; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.logs (id, action, user_id, json, dttm, dashboard_id, slice_id, duration_ms, referrer) FROM stdin;
1	welcome	1	{"path": "/superset/welcome/", "object_ref": "Superset.welcome"}	2024-04-10 14:49:01.873974	\N	0	322	http://localhost:8088/login/?next=http://localhost:8088/roles/list/
2	LogRestApi.recent_activity	1	{"path": "/api/v1/log/recent_activity/", "q": "(page_size:6)", "object_ref": "LogRestApi.recent_activity", "rison": {"page_size": 6}}	2024-04-10 14:49:02.522305	\N	0	29	http://localhost:8088/superset/welcome/
3	DatabaseRestApi.get_list	1	{"path": "/api/v1/database/", "q": "(filters:!((col:database_name,opr:neq,value:examples)))", "rison": {"filters": [{"col": "database_name", "opr": "neq", "value": "examples"}]}}	2024-04-10 14:49:02.57652	\N	0	109	http://localhost:8088/superset/welcome/
4	DatabaseRestApi.get_list	1	{"path": "/api/v1/database/", "q": "(filters:!((col:allow_file_upload,opr:upload_is_enabled,value:!t)))", "rison": {"filters": [{"col": "allow_file_upload", "opr": "upload_is_enabled", "value": true}]}}	2024-04-10 14:49:02.608325	\N	0	149	http://localhost:8088/superset/welcome/
5	DashboardRestApi.get_list	1	{"path": "/api/v1/dashboard/", "q": "(filters:!(),order_column:changed_on_delta_humanized,order_direction:desc,page:0,page_size:5)", "rison": {"filters": [], "order_column": "changed_on_delta_humanized", "order_direction": "desc", "page": 0, "page_size": 5}}	2024-04-10 14:49:02.848968	\N	0	88	http://localhost:8088/superset/welcome/
6	ChartRestApi.get_list	1	{"path": "/api/v1/chart/", "q": "(filters:!(),order_column:changed_on_delta_humanized,order_direction:desc,page:0,page_size:5)", "rison": {"filters": [], "order_column": "changed_on_delta_humanized", "order_direction": "desc", "page": 0, "page_size": 5}}	2024-04-10 14:49:02.893132	\N	0	264	http://localhost:8088/superset/welcome/
7	log	1	{"source": "sqlLab", "ts": 1712760542356, "event_name": "spa_navigation", "path": "/superset/welcome/", "event_type": "user", "event_id": "ZvvS4iq5Lh", "visibility": "visible"}	2024-04-10 14:49:03.384	\N	0	0	http://localhost:8088/superset/welcome/
8	ChartRestApi.get_list	1	{"path": "/api/v1/chart/", "q": "(filters:!((col:owners,opr:rel_m_m,value:'1')),order_column:changed_on_delta_humanized,order_direction:desc,page:0,page_size:5)", "rison": {"filters": [{"col": "owners", "opr": "rel_m_m", "value": "1"}], "order_column": "changed_on_delta_humanized", "order_direction": "desc", "page": 0, "page_size": 5}}	2024-04-10 14:49:04.608688	\N	0	108	http://localhost:8088/superset/welcome/
9	DashboardRestApi.get_list	1	{"path": "/api/v1/dashboard/", "q": "(filters:!((col:owners,opr:rel_m_m,value:'1')),order_column:changed_on_delta_humanized,order_direction:desc,page:0,page_size:5)", "rison": {"filters": [{"col": "owners", "opr": "rel_m_m", "value": "1"}], "order_column": "changed_on_delta_humanized", "order_direction": "desc", "page": 0, "page_size": 5}}	2024-04-10 14:49:04.635827	\N	0	97	http://localhost:8088/superset/welcome/
10	SavedQueryRestApi.get_list	1	{"path": "/api/v1/saved_query/", "q": "(filters:!((col:created_by,opr:rel_o_m,value:'1')),order_column:changed_on_delta_humanized,order_direction:desc,page:0,page_size:5)", "rison": {"filters": [{"col": "created_by", "opr": "rel_o_m", "value": "1"}], "order_column": "changed_on_delta_humanized", "order_direction": "desc", "page": 0, "page_size": 5}}	2024-04-10 14:49:04.744041	\N	0	97	http://localhost:8088/superset/welcome/
11	ChartRestApi.info	1	{"path": "/api/v1/chart/_info", "q": "(keys:!(permissions))", "rison": {"keys": ["permissions"]}}	2024-04-10 14:49:04.745792	\N	0	24	http://localhost:8088/superset/welcome/
12	DashboardRestApi.info	1	{"path": "/api/v1/dashboard/_info", "q": "(keys:!(permissions))", "rison": {"keys": ["permissions"]}}	2024-04-10 14:49:04.771223	\N	0	15	http://localhost:8088/superset/welcome/
13	DatabaseRestApi.get_list	1	{"path": "/api/v1/database/", "q": "(filters:!((col:database_name,opr:neq,value:examples)))", "rison": {"filters": [{"col": "database_name", "opr": "neq", "value": "examples"}]}}	2024-04-10 14:51:31.790803	\N	0	29	http://localhost:8088/dashboard/list/
14	DatabaseRestApi.get_list	1	{"path": "/api/v1/database/", "q": "(filters:!((col:allow_file_upload,opr:upload_is_enabled,value:!t)))", "rison": {"filters": [{"col": "allow_file_upload", "opr": "upload_is_enabled", "value": true}]}}	2024-04-10 14:51:31.793787	\N	0	46	http://localhost:8088/dashboard/list/
15	DashboardRestApi.get_list	1	{"path": "/api/v1/dashboard/", "q": "(order_column:changed_on_delta_humanized,order_direction:desc,page:0,page_size:25)", "rison": {"order_column": "changed_on_delta_humanized", "order_direction": "desc", "page": 0, "page_size": 25}}	2024-04-10 14:51:31.852179	\N	0	17	http://localhost:8088/dashboard/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc&viewMode=table
16	DashboardRestApi.info	1	{"path": "/api/v1/dashboard/_info", "q": "(keys:!(permissions))", "rison": {"keys": ["permissions"]}}	2024-04-10 14:51:31.859142	\N	0	26	http://localhost:8088/dashboard/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc&viewMode=table
17	log	1	{"source": "sqlLab", "ts": 1712760691710, "event_name": "spa_navigation", "path": "/dashboard/list/", "event_type": "user", "event_id": "hWjjGQ79zN", "visibility": "visible"}	2024-04-10 14:51:32.753095	\N	0	0	http://localhost:8088/dashboard/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc&viewMode=table
18	DatabaseRestApi.info	1	{"path": "/api/v1/database/_info", "q": "(keys:!(permissions))", "rison": {"keys": ["permissions"]}}	2024-04-10 14:51:34.596431	\N	0	25	http://localhost:8088/databaseview/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc
19	DatabaseRestApi.get_list	1	{"path": "/api/v1/database/", "q": "(order_column:changed_on_delta_humanized,order_direction:desc,page:0,page_size:25)", "rison": {"order_column": "changed_on_delta_humanized", "order_direction": "desc", "page": 0, "page_size": 25}}	2024-04-10 14:51:34.654379	\N	0	58	http://localhost:8088/databaseview/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc
20	DatabaseRestApi.get_list	1	{"path": "/api/v1/database/", "q": "(filters:!((col:allow_file_upload,opr:upload_is_enabled,value:!t)))", "rison": {"filters": [{"col": "allow_file_upload", "opr": "upload_is_enabled", "value": true}]}}	2024-04-10 14:51:34.656582	\N	0	66	http://localhost:8088/databaseview/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc
21	log	1	{"source": "sqlLab", "ts": 1712760694491, "event_name": "spa_navigation", "path": "/databaseview/list/", "event_type": "user", "event_id": "57cVt9yiY", "visibility": "visible"}	2024-04-10 14:51:35.502904	\N	0	0	http://localhost:8088/databaseview/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc
22	DatabaseRestApi.get_list	1	{"path": "/api/v1/database/", "q": "(filters:!((col:allow_file_upload,opr:upload_is_enabled,value:!t)))", "rison": {"filters": [{"col": "allow_file_upload", "opr": "upload_is_enabled", "value": true}]}}	2024-04-10 14:51:38.228174	\N	0	30	http://localhost:8088/databaseview/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc
23	DatabaseRestApi.available	1	{"path": "/api/v1/database/available/", "object_ref": "DatabaseRestApi.available"}	2024-04-10 14:51:38.263429	\N	0	71	http://localhost:8088/databaseview/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc
24	validation_error	1	{"path": "/api/v1/database/validate_parameters/", "engine": "postgresql"}	2024-04-10 14:51:42.720087	\N	0	\N	http://localhost:8088/databaseview/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc
29	test_connection_attempt	1	{"path": "/api/v1/database/", "engine": "PostgresEngineSpec"}	2024-04-10 14:51:50.483409	\N	0	\N	http://localhost:8088/databaseview/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc
37	DatasetRestApi.info	1	{"path": "/api/v1/dataset/_info", "q": "(keys:!(permissions))", "rison": {"keys": ["permissions"]}}	2024-04-10 14:51:57.416718	\N	0	35	http://localhost:8088/tablemodelview/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc
42	DatabaseRestApi.schemas	1	{"path": "/api/v1/database/1/schemas/", "q": "(force:!f)", "url_rule": "/api/v1/database/<int:pk>/schemas/", "object_ref": "DatabaseRestApi.schemas", "pk": 1, "rison": {"force": false}}	2024-04-10 14:52:01.113462	\N	0	43	http://localhost:8088/dataset/add/
46	DatasetRestApi.post	1	{"path": "/api/v1/dataset/", "object_ref": "DatasetRestApi.post"}	2024-04-10 14:52:09.988976	\N	0	216	http://localhost:8088/dataset/add/
51	log	1	{"source": "sqlLab", "ts": 1712760736733, "event_name": "spa_navigation", "path": "/explore/", "event_type": "user", "event_id": "q-Cew982q", "visibility": "visible"}	2024-04-10 14:52:18.068557	\N	0	0	http://localhost:8088/explore/?viz_type=big_number_total&datasource=1__table
52	log	1	{"source": "sqlLab", "ts": 1712760737024, "event_name": "mount_explorer", "event_type": "user", "event_id": "-TZlFqich2", "visibility": "visible"}	2024-04-10 14:52:18.068602	\N	0	0	http://localhost:8088/explore/?viz_type=big_number_total&datasource=1__table
55	ChartDataRestApi.data	1	{"path": "/api/v1/chart/data", "object_ref": "ChartDataRestApi.data"}	2024-04-10 14:52:28.197428	\N	0	94	http://localhost:8088/explore/?form_data_key=8jX89tTb2hVwMAuYIttHVbBqpL0joluHOKtHL9GymhasXK7QJar2HftaF_3xKW5U&viz_type=big_number_total&datasource=1__table&datasource_id=1&datasource_type=table
60	DashboardRestApi.get_list	1	{"path": "/api/v1/dashboard/", "q": "(columns:!(id,dashboard_title),filters:!((col:dashboard_title,opr:ct,value:''),(col:owners,opr:rel_m_m,value:1)),order_column:dashboard_title,page:0,page_size:100)", "rison": {"columns": ["id", "dashboard_title"], "filters": [{"col": "dashboard_title", "opr": "ct", "value": ""}, {"col": "owners", "opr": "rel_m_m", "value": 1}], "order_column": "dashboard_title", "page": 0, "page_size": 100}}	2024-04-10 14:52:42.178911	\N	0	51	http://localhost:8088/explore/?form_data_key=8jX89tTb2hVwMAuYIttHVbBqpL0joluHOKtHL9GymhasXK7QJar2HftaF_3xKW5U&viz_type=big_number_total&datasource=1__table&datasource_id=1&datasource_type=table
65	ChartDataRestApi.data	1	{"path": "/api/v1/chart/data", "form_data": {"filters": [{"col": "created_at", "op": "TEMPORAL_RANGE", "val": "No filter"}], "extras": {"having": "", "where": ""}, "applied_time_extras": {}, "columns": [], "metrics": ["count"], "annotation_layers": [], "series_limit": 0, "order_desc": true, "url_params": {"datasource_id": "1", "datasource_type": "table", "save_action": "saveas", "slice_id": "1", "viz_type": "big_number_total"}, "custom_params": {}, "custom_form_data": {}, "slice_id": 1}, "object_ref": "ChartDataRestApi.data"}	2024-04-10 14:52:45.182004	\N	1	163	http://localhost:8088/explore/?viz_type=big_number_total&datasource=1__table&datasource_id=1&datasource_type=table&save_action=saveas&slice_id=1
68	ExploreFormDataRestApi.post	1	{"path": "/api/v1/explore/form_data", "tab_id": "7", "object_ref": "ExploreFormDataRestApi.post"}	2024-04-10 14:52:46.022455	\N	0	36	http://localhost:8088/explore/?viz_type=big_number_total&datasource=1__table&datasource_id=1&datasource_type=table&save_action=saveas&slice_id=1
78	DashboardRestApi.get_charts	1	{"path": "/api/v1/dashboard/1/charts", "url_rule": "/api/v1/dashboard/<id_or_slug>/charts", "object_ref": "DashboardRestApi.get_charts", "id_or_slug": "1"}	2024-04-10 14:52:49.038157	\N	0	9	http://localhost:8088/superset/dashboard/1/?edit=true
83	log	1	{"source": "sqlLab", "ts": 1712760768925, "event_name": "spa_navigation", "path": "/superset/dashboard/1/", "event_type": "user", "event_id": "JM2Txb9ivX", "visibility": "visible"}	2024-04-10 14:52:50.097469	\N	0	0	http://localhost:8088/superset/dashboard/1/?edit=true
84	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "tSg4Vm5xD", "version": "v2", "ts": 1712760769064, "event_name": "mount_dashboard", "is_soft_navigation": false, "is_edit_mode": true, "mount_duration": 902, "is_empty": true, "is_published": false, "event_type": "user", "event_id": "IvRLdD1Ifq", "visibility": "visible"}	2024-04-10 14:52:50.097489	\N	0	0	http://localhost:8088/superset/dashboard/1/?edit=true
25	validation_error	1	{"path": "/api/v1/database/validate_parameters/", "engine": "postgresql"}	2024-04-10 14:51:45.078324	\N	0	\N	http://localhost:8088/databaseview/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc
30	test_connection_success	1	{"path": "/api/v1/database/", "engine": "PostgresEngineSpec"}	2024-04-10 14:51:50.501294	\N	0	\N	http://localhost:8088/databaseview/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc
34	DatabaseRestApi.put	1	{"path": "/api/v1/database/1", "url_rule": "/api/v1/database/<int:pk>", "object_ref": "DatabaseRestApi.put", "pk": 1}	2024-04-10 14:51:52.279628	\N	0	69	http://localhost:8088/databaseview/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc
39	log	1	{"source": "sqlLab", "ts": 1712760717264, "event_name": "spa_navigation", "path": "/tablemodelview/list/", "event_type": "user", "event_id": "rvcRocEU6", "visibility": "visible"}	2024-04-10 14:51:58.299499	\N	0	0	http://localhost:8088/tablemodelview/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc
43	DatasetRestApi.get_list	1	{"path": "/api/v1/dataset/", "q": "(filters:!((col:database,opr:rel_o_m,value:1),(col:schema,opr:eq,value:public),(col:sql,opr:dataset_is_null_or_empty,value:!t)),page:0)", "rison": {"filters": [{"col": "database", "opr": "rel_o_m", "value": 1}, {"col": "schema", "opr": "eq", "value": "public"}, {"col": "sql", "opr": "dataset_is_null_or_empty", "value": true}], "page": 0}}	2024-04-10 14:52:03.437487	\N	0	87	http://localhost:8088/dataset/add/
47	DatasetRestApi.get_list	1	{"path": "/api/v1/dataset/", "q": "(columns:!(id,table_name,datasource_type,database.database_name,schema),filters:!((col:table_name,opr:ct,value:warrants)),order_column:table_name,order_direction:asc,page:0,page_size:1)", "rison": {"columns": ["id", "table_name", "datasource_type", "database.database_name", "schema"], "filters": [{"col": "table_name", "opr": "ct", "value": "warrants"}], "order_column": "table_name", "order_direction": "asc", "page": 0, "page_size": 1}}	2024-04-10 14:52:10.184774	\N	0	56	http://localhost:8088/chart/add/?dataset=warrants
56	ChartDataRestApi.data	1	{"path": "/api/v1/chart/data", "object_ref": "ChartDataRestApi.data"}	2024-04-10 14:52:28.363916	\N	0	85	http://localhost:8088/explore/?form_data_key=8jX89tTb2hVwMAuYIttHVbBqpL0joluHOKtHL9GymhasXK7QJar2HftaF_3xKW5U&viz_type=big_number_total&datasource=1__table&datasource_id=1&datasource_type=table
66	ChartDataRestApi.data	1	{"path": "/api/v1/chart/data", "form_data": {"filters": [{"col": "created_at", "op": "TEMPORAL_RANGE", "val": "No filter"}], "extras": {"having": "", "where": ""}, "applied_time_extras": {}, "columns": [], "metrics": ["count"], "annotation_layers": [], "series_limit": 0, "order_desc": true, "url_params": {"datasource_id": "1", "datasource_type": "table", "save_action": "saveas", "slice_id": "1", "viz_type": "big_number_total"}, "custom_params": {}, "custom_form_data": {}, "slice_id": 1}, "object_ref": "ChartDataRestApi.data"}	2024-04-10 14:52:45.19051	\N	1	143	http://localhost:8088/explore/?viz_type=big_number_total&datasource=1__table&datasource_id=1&datasource_type=table&save_action=saveas&slice_id=1
69	log	1	{"source": "explore", "source_id": 1, "impression_id": "D6HflmryO", "version": "v2", "ts": 1712760764937, "event_name": "change_explore_controls", "event_type": "user", "event_id": "NdLCKI8Hml", "visibility": "visible"}	2024-04-10 14:52:46.324623	\N	0	0	http://localhost:8088/explore/?form_data_key=JP73JUurn4cB4BAYPEOCewpuduKu_kgU8Uq_5Ex3RAB1qYA-4_QbRQ3iYSlYoDmS&viz_type=big_number_total&datasource=1__table&datasource_id=1&datasource_type=table&save_action=saveas&slice_id=1
70	log	1	{"source": "explore", "source_id": 1, "impression_id": "D6HflmryO", "version": "v2", "ts": 1712760765262, "event_name": "load_chart", "slice_id": 1, "applied_filters": [{"column": "created_at"}], "is_cached": true, "force_refresh": false, "row_count": 1, "datasource": "1__table", "start_offset": 28246, "duration": 283, "viz_type": "big_number_total", "data_age": 17262, "event_type": "timing", "trigger_event": "NdLCKI8Hml"}	2024-04-10 14:52:46.324639	\N	0	0	http://localhost:8088/explore/?form_data_key=JP73JUurn4cB4BAYPEOCewpuduKu_kgU8Uq_5Ex3RAB1qYA-4_QbRQ3iYSlYoDmS&viz_type=big_number_total&datasource=1__table&datasource_id=1&datasource_type=table&save_action=saveas&slice_id=1
71	log	1	{"source": "explore", "source_id": 1, "impression_id": "D6HflmryO", "version": "v2", "ts": 1712760765289, "event_name": "render_chart", "slice_id": 1, "viz_type": "big_number_total", "start_offset": 28539, "duration": 16, "event_type": "timing", "trigger_event": "NdLCKI8Hml"}	2024-04-10 14:52:46.324645	\N	0	0	http://localhost:8088/explore/?form_data_key=JP73JUurn4cB4BAYPEOCewpuduKu_kgU8Uq_5Ex3RAB1qYA-4_QbRQ3iYSlYoDmS&viz_type=big_number_total&datasource=1__table&datasource_id=1&datasource_type=table&save_action=saveas&slice_id=1
79	DashboardRestApi.get_datasets	1	{"path": "/api/v1/dashboard/1/datasets", "url_rule": "/api/v1/dashboard/<id_or_slug>/datasets", "object_ref": "DashboardRestApi.get_datasets", "id_or_slug": "1"}	2024-04-10 14:52:49.040176	\N	0	8	http://localhost:8088/superset/dashboard/1/?edit=true
88	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "tSg4Vm5xD", "version": "v2", "ts": 1712760772583, "event_name": "hide_browser_tab", "start_offset": 4420, "duration": 3926, "event_type": "timing", "trigger_event": "IvRLdD1Ifq"}	2024-04-10 14:52:57.531166	1	0	0	http://localhost:8088/superset/dashboard/1/?edit=true
26	validation_error	1	{"path": "/api/v1/database/validate_parameters/", "engine": "postgresql"}	2024-04-10 14:51:46.620975	\N	0	\N	http://localhost:8088/databaseview/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc
31	DatabaseRestApi.post	1	{"path": "/api/v1/database/", "object_ref": "DatabaseRestApi.post"}	2024-04-10 14:51:50.570606	\N	0	109	http://localhost:8088/databaseview/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc
35	DatabaseRestApi.get_list	1	{"path": "/api/v1/database/", "q": "(order_column:changed_on_delta_humanized,order_direction:desc,page:0,page_size:25)", "rison": {"order_column": "changed_on_delta_humanized", "order_direction": "desc", "page": 0, "page_size": 25}}	2024-04-10 14:51:52.388395	\N	0	35	http://localhost:8088/databaseview/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc
40	DatabaseRestApi.get_list	1	{"path": "/api/v1/database/", "q": "(filters:!((col:database_name,opr:ct,value:'')),order_columns:database_name,order_direction:asc,page:0,page_size:100)", "rison": {"filters": [{"col": "database_name", "opr": "ct", "value": ""}], "order_columns": "database_name", "order_direction": "asc", "page": 0, "page_size": 100}}	2024-04-10 14:51:58.728768	\N	0	24	http://localhost:8088/dataset/add/
45	DatabaseRestApi.table_metadata	1	{"path": "/api/v1/database/1/table/warrants/public/", "url_rule": "/api/v1/database/<int:pk>/table/<path:table_name>/<schema_name>/", "object_ref": "DatabaseRestApi.table_metadata"}	2024-04-10 14:52:07.490966	\N	0	105	http://localhost:8088/dataset/add/
49	ExploreRestApi.get	1	{"path": "/api/v1/explore/", "viz_type": "big_number_total", "datasource_id": "1", "datasource_type": "table", "object_ref": "ExploreRestApi.get"}	2024-04-10 14:52:16.87822	\N	0	29	http://localhost:8088/explore/?viz_type=big_number_total&datasource=1__table
53	ExploreFormDataRestApi.post	1	{"path": "/api/v1/explore/form_data", "tab_id": "7", "object_ref": "ExploreFormDataRestApi.post"}	2024-04-10 14:52:18.139527	\N	0	44	http://localhost:8088/explore/?viz_type=big_number_total&datasource=1__table
57	ExploreFormDataRestApi.put	1	{"path": "/api/v1/explore/form_data/8jX89tTb2hVwMAuYIttHVbBqpL0joluHOKtHL9GymhasXK7QJar2HftaF_3xKW5U", "tab_id": "7", "url_rule": "/api/v1/explore/form_data/<string:key>", "object_ref": "ExploreFormDataRestApi.put", "key": "8jX89tTb2hVwMAuYIttHVbBqpL0joluHOKtHL9GymhasXK7QJar2HftaF_3xKW5U"}	2024-04-10 14:52:29.147268	\N	0	40	http://localhost:8088/explore/?form_data_key=8jX89tTb2hVwMAuYIttHVbBqpL0joluHOKtHL9GymhasXK7QJar2HftaF_3xKW5U&viz_type=big_number_total&datasource=1__table&datasource_id=1&datasource_type=table
62	ExploreRestApi.get	1	{"path": "/api/v1/explore/", "viz_type": "big_number_total", "datasource_id": "1", "datasource_type": "table", "save_action": "saveas", "slice_id": "1", "object_ref": "ExploreRestApi.get"}	2024-04-10 14:52:44.912562	\N	1	51	http://localhost:8088/explore/?viz_type=big_number_total&datasource=1__table&datasource_id=1&datasource_type=table&save_action=saveas&slice_id=1
67	ChartDataRestApi.data	1	{"path": "/api/v1/chart/data", "form_data": {"filters": [{"col": "created_at", "op": "TEMPORAL_RANGE", "val": "No filter"}], "extras": {"having": "", "where": ""}, "applied_time_extras": {}, "columns": [], "metrics": ["count"], "annotation_layers": [], "series_limit": 0, "order_desc": true, "url_params": {"datasource_id": "1", "datasource_type": "table", "save_action": "saveas", "slice_id": "1", "viz_type": "big_number_total"}, "custom_params": {}, "custom_form_data": {}, "slice_id": 1}, "object_ref": "ChartDataRestApi.data"}	2024-04-10 14:52:45.224008	\N	1	75	http://localhost:8088/explore/?viz_type=big_number_total&datasource=1__table&datasource_id=1&datasource_type=table&save_action=saveas&slice_id=1
73	DashboardRestApi.info	1	{"path": "/api/v1/dashboard/_info", "q": "(keys:!(permissions))", "rison": {"keys": ["permissions"]}}	2024-04-10 14:52:46.724343	\N	0	33	http://localhost:8088/dashboard/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc&viewMode=table
76	DatabaseRestApi.get_list	1	{"path": "/api/v1/database/", "q": "(filters:!((col:database_name,opr:neq,value:examples)))", "rison": {"filters": [{"col": "database_name", "opr": "neq", "value": "examples"}]}}	2024-04-10 14:52:49.001174	\N	0	43	http://localhost:8088/superset/dashboard/1/?edit=true
81	DashboardRestApi.favorite_status	1	{"path": "/api/v1/dashboard/favorite_status/", "q": "!(1)", "object_ref": "DashboardRestApi.favorite_status", "rison": [1]}	2024-04-10 14:52:49.299539	\N	0	13	http://localhost:8088/superset/dashboard/1/?edit=true
85	DatabaseRestApi.get_list	1	{"path": "/api/v1/database/", "q": "(filters:!((col:database_name,opr:neq,value:examples)))", "rison": {"filters": [{"col": "database_name", "opr": "neq", "value": "examples"}]}}	2024-04-10 14:52:52.636083	\N	0	33	http://localhost:8088/chart/add?dashboard_id=1
90	ChartDataRestApi.data	1	{"path": "/api/v1/chart/data", "form_data": {"filters": [{"col": "created_at", "op": "TEMPORAL_RANGE", "val": "No filter"}], "extras": {"having": "", "where": ""}, "applied_time_extras": {}, "columns": [], "metrics": ["count"], "annotation_layers": [], "series_limit": 0, "order_desc": true, "url_params": {}, "custom_params": {}, "custom_form_data": {}, "slice_id": 1}, "dashboard_id": "1", "object_ref": "ChartDataRestApi.data"}	2024-04-10 14:52:58.443626	1	1	41	http://localhost:8088/superset/dashboard/1/?edit=true
27	DatabaseRestApi.validate_parameters	1	{"path": "/api/v1/database/validate_parameters/", "object_ref": "DatabaseRestApi.validate_parameters"}	2024-04-10 14:51:50.391303	\N	0	78	http://localhost:8088/databaseview/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc
32	DatabaseRestApi.get_list	1	{"path": "/api/v1/database/", "q": "(order_column:changed_on_delta_humanized,order_direction:desc,page:0,page_size:25)", "rison": {"order_column": "changed_on_delta_humanized", "order_direction": "desc", "page": 0, "page_size": 25}}	2024-04-10 14:51:50.655164	\N	0	18	http://localhost:8088/databaseview/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc
36	DatabaseRestApi.get_list	1	{"path": "/api/v1/database/", "q": "(filters:!((col:allow_file_upload,opr:upload_is_enabled,value:!t)))", "rison": {"filters": [{"col": "allow_file_upload", "opr": "upload_is_enabled", "value": true}]}}	2024-04-10 14:51:52.394605	\N	0	31	http://localhost:8088/databaseview/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc
41	log	1	{"source": "sqlLab", "ts": 1712760718645, "event_name": "spa_navigation", "path": "/dataset/add/", "event_type": "user", "event_id": "SQQEf3QtF", "visibility": "visible"}	2024-04-10 14:51:59.655984	\N	0	0	http://localhost:8088/dataset/add/
50	DatabaseRestApi.table_extra_metadata	1	{"path": "/api/v1/database/1/table_extra/warrants/public/", "url_rule": "/api/v1/database/<int:pk>/table_extra/<path:table_name>/<schema_name>/", "object_ref": "DatabaseRestApi.table_extra_metadata"}	2024-04-10 14:52:17.055231	\N	0	0	http://localhost:8088/explore/?viz_type=big_number_total&datasource=1__table
54	log	1	{"source": "sqlLab", "ts": 1712760746279, "event_name": "change_explore_controls", "event_type": "user", "event_id": "KTp_tqhT4", "visibility": "visible"}	2024-04-10 14:52:27.292548	\N	0	0	http://localhost:8088/explore/?form_data_key=8jX89tTb2hVwMAuYIttHVbBqpL0joluHOKtHL9GymhasXK7QJar2HftaF_3xKW5U&viz_type=big_number_total&datasource=1__table&datasource_id=1&datasource_type=table
58	log	1	{"source": "sqlLab", "ts": 1712760748208, "event_name": "load_chart", "slice_id": 0, "applied_filters": [{"column": "created_at"}], "is_cached": null, "force_refresh": false, "row_count": 1, "datasource": "1__table", "start_offset": 11349, "duration": 126, "viz_type": "big_number_total", "data_age": null, "event_type": "timing", "trigger_event": "KTp_tqhT4"}	2024-04-10 14:52:29.273754	\N	0	0	http://localhost:8088/explore/?form_data_key=8jX89tTb2hVwMAuYIttHVbBqpL0joluHOKtHL9GymhasXK7QJar2HftaF_3xKW5U&viz_type=big_number_total&datasource=1__table&datasource_id=1&datasource_type=table
59	log	1	{"source": "sqlLab", "ts": 1712760748246, "event_name": "render_chart", "slice_id": 0, "viz_type": "big_number_total", "start_offset": 11488, "duration": 25, "event_type": "timing", "trigger_event": "KTp_tqhT4"}	2024-04-10 14:52:29.273762	\N	0	0	http://localhost:8088/explore/?form_data_key=8jX89tTb2hVwMAuYIttHVbBqpL0joluHOKtHL9GymhasXK7QJar2HftaF_3xKW5U&viz_type=big_number_total&datasource=1__table&datasource_id=1&datasource_type=table
63	DatabaseRestApi.table_extra_metadata	1	{"path": "/api/v1/database/1/table_extra/warrants/public/", "url_rule": "/api/v1/database/<int:pk>/table_extra/<path:table_name>/<schema_name>/", "object_ref": "DatabaseRestApi.table_extra_metadata"}	2024-04-10 14:52:45.032566	\N	0	0	http://localhost:8088/explore/?viz_type=big_number_total&datasource=1__table&datasource_id=1&datasource_type=table&save_action=saveas&slice_id=1
72	DashboardRestApi.get_list	1	{"path": "/api/v1/dashboard/", "q": "(order_column:changed_on_delta_humanized,order_direction:desc,page:0,page_size:25)", "rison": {"order_column": "changed_on_delta_humanized", "order_direction": "desc", "page": 0, "page_size": 25}}	2024-04-10 14:52:46.715734	\N	0	19	http://localhost:8088/dashboard/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc&viewMode=table
75	dashboard	1	{"path": "/superset/dashboard/1/", "edit": "true", "url_rule": "/superset/dashboard/<dashboard_id_or_slug>/", "object_ref": "Superset.dashboard", "dashboard_id_or_slug": "1", "dashboard_id": 1, "dashboard_version": "v2", "dash_edit_perm": true, "edit_mode": true}	2024-04-10 14:52:48.448071	1	0	165	http://localhost:8088/dashboard/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc&viewMode=table
80	DashboardRestApi.get	1	{"path": "/api/v1/dashboard/1", "url_rule": "/api/v1/dashboard/<id_or_slug>", "object_ref": "DashboardRestApi.get", "dashboard_id": 1}	2024-04-10 14:52:49.046909	1	0	17	http://localhost:8088/superset/dashboard/1/?edit=true
86	DatabaseRestApi.get_list	1	{"path": "/api/v1/database/", "q": "(filters:!((col:allow_file_upload,opr:upload_is_enabled,value:!t)))", "rison": {"filters": [{"col": "allow_file_upload", "opr": "upload_is_enabled", "value": true}]}}	2024-04-10 14:52:52.641473	\N	0	41	http://localhost:8088/chart/add?dashboard_id=1
89	fetch_datasource_metadata	1	{"path": "/superset/fetch_datasource_metadata", "datasourceKey": "1__table", "object_ref": "Superset.fetch_datasource_metadata"}	2024-04-10 14:52:58.438471	\N	0	37	http://localhost:8088/superset/dashboard/1/?edit=true
28	DatabaseRestApi.validate_parameters	1	{"path": "/api/v1/database/validate_parameters/", "object_ref": "DatabaseRestApi.validate_parameters"}	2024-04-10 14:51:50.398839	\N	0	34	http://localhost:8088/databaseview/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc
33	DatabaseRestApi.validate_parameters	1	{"path": "/api/v1/database/validate_parameters/", "object_ref": "DatabaseRestApi.validate_parameters"}	2024-04-10 14:51:52.151401	\N	0	52	http://localhost:8088/databaseview/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc
38	DatasetRestApi.get_list	1	{"path": "/api/v1/dataset/", "q": "(order_column:changed_on_delta_humanized,order_direction:desc,page:0,page_size:25)", "rison": {"order_column": "changed_on_delta_humanized", "order_direction": "desc", "page": 0, "page_size": 25}}	2024-04-10 14:51:57.446905	\N	0	63	http://localhost:8088/tablemodelview/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc
44	DatabaseRestApi.tables	1	{"path": "/api/v1/database/1/tables/", "q": "(force:!f,schema_name:public)", "url_rule": "/api/v1/database/<int:pk>/tables/", "object_ref": "DatabaseRestApi.tables", "pk": 1, "rison": {"force": false, "schema_name": "public"}}	2024-04-10 14:52:03.487214	\N	0	142	http://localhost:8088/dataset/add/
48	log	1	{"source": "sqlLab", "ts": 1712760730012, "event_name": "spa_navigation", "path": "/chart/add/", "event_type": "user", "event_id": "zNfyEtCtr", "visibility": "visible"}	2024-04-10 14:52:11.054234	\N	0	0	http://localhost:8088/chart/add/?dataset=warrants
61	ChartRestApi.post	1	{"path": "/api/v1/chart/", "object_ref": "ChartRestApi.post"}	2024-04-10 14:52:44.782746	\N	0	52	http://localhost:8088/explore/?form_data_key=8jX89tTb2hVwMAuYIttHVbBqpL0joluHOKtHL9GymhasXK7QJar2HftaF_3xKW5U&viz_type=big_number_total&datasource=1__table&datasource_id=1&datasource_type=table
64	ChartRestApi.favorite_status	1	{"path": "/api/v1/chart/favorite_status/", "q": "!(1)", "object_ref": "ChartRestApi.favorite_status", "rison": [1]}	2024-04-10 14:52:45.168749	\N	0	54	http://localhost:8088/explore/?viz_type=big_number_total&datasource=1__table&datasource_id=1&datasource_type=table&save_action=saveas&slice_id=1
74	log	1	{"source": "explore", "source_id": 1, "impression_id": "D6HflmryO", "version": "v2", "ts": 1712760766641, "event_name": "spa_navigation", "path": "/dashboard/list/", "event_type": "user", "event_id": "LLqIcuUIh", "visibility": "visible"}	2024-04-10 14:52:47.673443	\N	0	0	http://localhost:8088/dashboard/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc&viewMode=table
77	DatabaseRestApi.get_list	1	{"path": "/api/v1/database/", "q": "(filters:!((col:allow_file_upload,opr:upload_is_enabled,value:!t)))", "rison": {"filters": [{"col": "allow_file_upload", "opr": "upload_is_enabled", "value": true}]}}	2024-04-10 14:52:49.014701	\N	0	45	http://localhost:8088/superset/dashboard/1/?edit=true
82	ChartRestApi.get_list	1	{"path": "/api/v1/chart/", "q": "(columns:!(changed_on_delta_humanized,changed_on_utc,datasource_id,datasource_type,datasource_url,datasource_name_text,description_markeddown,description,id,params,slice_name,thumbnail_url,url,viz_type,owners.id,created_by.id),filters:!((col:viz_type,opr:neq,value:filter_box),(col:owners,opr:rel_m_m,value:1)),order_column:changed_on_delta_humanized,order_direction:desc,page_size:200)", "rison": {"columns": ["changed_on_delta_humanized", "changed_on_utc", "datasource_id", "datasource_type", "datasource_url", "datasource_name_text", "description_markeddown", "description", "id", "params", "slice_name", "thumbnail_url", "url", "viz_type", "owners.id", "created_by.id"], "filters": [{"col": "viz_type", "opr": "neq", "value": "filter_box"}, {"col": "owners", "opr": "rel_m_m", "value": 1}], "order_column": "changed_on_delta_humanized", "order_direction": "desc", "page_size": 200}}	2024-04-10 14:52:49.369528	\N	0	70	http://localhost:8088/superset/dashboard/1/?edit=true
87	log	1	{"source": "sqlLab", "ts": 1712760772555, "event_name": "spa_navigation", "path": "/chart/add", "event_type": "user", "event_id": "AncMEhYuW", "visibility": "visible"}	2024-04-10 14:52:53.573674	\N	0	0	http://localhost:8088/chart/add?dashboard_id=1
91	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "tSg4Vm5xD", "version": "v2", "ts": 1712760778494, "event_name": "load_chart", "slice_id": 1, "applied_filters": [{"column": "created_at"}], "is_cached": true, "force_refresh": false, "row_count": 1, "datasource": "1__table", "start_offset": 10186, "duration": 144, "has_extra_filters": false, "viz_type": "big_number_total", "data_age": 30494, "event_type": "timing", "trigger_event": "IvRLdD1Ifq"}	2024-04-10 14:52:59.575304	1	0	0	http://localhost:8088/superset/dashboard/1/?edit=true
92	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "tSg4Vm5xD", "version": "v2", "ts": 1712760778541, "event_name": "render_chart", "slice_id": 1, "viz_type": "big_number_total", "start_offset": 10340, "duration": 38, "event_type": "timing", "trigger_event": "IvRLdD1Ifq"}	2024-04-10 14:52:59.575318	1	0	0	http://localhost:8088/superset/dashboard/1/?edit=true
93	DashboardRestApi.put	1	{"path": "/api/v1/dashboard/1", "url_rule": "/api/v1/dashboard/<pk>", "object_ref": "DashboardRestApi.put", "pk": "1"}	2024-04-10 14:53:06.408075	\N	0	22	http://localhost:8088/superset/dashboard/1/?edit=true
94	DashboardFilterStateRestApi.post	1	{"path": "/api/v1/dashboard/1/filter_state", "tab_id": "7", "url_rule": "/api/v1/dashboard/<int:pk>/filter_state", "object_ref": "DashboardFilterStateRestApi.post", "pk": 1}	2024-04-10 14:53:07.048675	\N	0	45	http://localhost:8088/superset/dashboard/1/
95	dashboard	1	{"path": "/superset/dashboard/1/", "edit": "true", "native_filters_key": "x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q", "url_rule": "/superset/dashboard/<dashboard_id_or_slug>/", "object_ref": "Superset.dashboard", "dashboard_id_or_slug": "1", "dashboard_id": 1, "dashboard_version": "v2", "dash_edit_perm": true, "edit_mode": true}	2024-04-10 14:53:30.473469	1	0	18	\N
96	DatabaseRestApi.get_list	1	{"path": "/api/v1/database/", "q": "(filters:!((col:database_name,opr:neq,value:examples)))", "rison": {"filters": [{"col": "database_name", "opr": "neq", "value": "examples"}]}}	2024-04-10 14:53:30.984742	\N	0	67	http://localhost:8088/superset/dashboard/1/?edit=true&native_filters_key=x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q
97	DatabaseRestApi.get_list	1	{"path": "/api/v1/database/", "q": "(filters:!((col:allow_file_upload,opr:upload_is_enabled,value:!t)))", "rison": {"filters": [{"col": "allow_file_upload", "opr": "upload_is_enabled", "value": true}]}}	2024-04-10 14:53:30.987091	\N	0	107	http://localhost:8088/superset/dashboard/1/?edit=true&native_filters_key=x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q
98	DashboardRestApi.get	1	{"path": "/api/v1/dashboard/1", "url_rule": "/api/v1/dashboard/<id_or_slug>", "object_ref": "DashboardRestApi.get", "dashboard_id": 1}	2024-04-10 14:53:30.998512	1	0	35	http://localhost:8088/superset/dashboard/1/?edit=true&native_filters_key=x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q
181	DashboardFilterStateRestApi.post	1	{"path": "/api/v1/dashboard/1/filter_state", "tab_id": "6", "url_rule": "/api/v1/dashboard/<int:pk>/filter_state", "object_ref": "DashboardFilterStateRestApi.post", "pk": 1}	2024-04-11 02:29:15.873066	\N	0	43	http://localhost:8088/superset/dashboard/1/
99	DashboardRestApi.get_charts	1	{"path": "/api/v1/dashboard/1/charts", "url_rule": "/api/v1/dashboard/<id_or_slug>/charts", "object_ref": "DashboardRestApi.get_charts", "id_or_slug": "1"}	2024-04-10 14:53:31.007882	\N	0	17	http://localhost:8088/superset/dashboard/1/?edit=true&native_filters_key=x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q
103	ChartRestApi.get_list	1	{"path": "/api/v1/chart/", "q": "(columns:!(changed_on_delta_humanized,changed_on_utc,datasource_id,datasource_type,datasource_url,datasource_name_text,description_markeddown,description,id,params,slice_name,thumbnail_url,url,viz_type,owners.id,created_by.id),filters:!((col:viz_type,opr:neq,value:filter_box),(col:owners,opr:rel_m_m,value:1)),order_column:changed_on_delta_humanized,order_direction:desc,page_size:200)", "rison": {"columns": ["changed_on_delta_humanized", "changed_on_utc", "datasource_id", "datasource_type", "datasource_url", "datasource_name_text", "description_markeddown", "description", "id", "params", "slice_name", "thumbnail_url", "url", "viz_type", "owners.id", "created_by.id"], "filters": [{"col": "viz_type", "opr": "neq", "value": "filter_box"}, {"col": "owners", "opr": "rel_m_m", "value": 1}], "order_column": "changed_on_delta_humanized", "order_direction": "desc", "page_size": 200}}	2024-04-10 14:53:31.380191	\N	0	43	http://localhost:8088/superset/dashboard/1/?edit=true&native_filters_key=x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q
110	DashboardRestApi.info	1	{"path": "/api/v1/dashboard/_info", "q": "(keys:!(permissions))", "rison": {"keys": ["permissions"]}}	2024-04-10 14:53:49.88412	\N	0	33	http://localhost:8088/dashboard/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc&viewMode=table
115	DashboardRestApi.get_datasets	1	{"path": "/api/v1/dashboard/1/datasets", "url_rule": "/api/v1/dashboard/<id_or_slug>/datasets", "object_ref": "DashboardRestApi.get_datasets", "id_or_slug": "1"}	2024-04-10 14:53:56.726912	\N	0	259	http://localhost:8088/superset/dashboard/1/
128	csrf_token	1	{"path": "/api/v1/security/csrf_token/", "object_ref": "SecurityRestApi.csrf_token"}	2024-04-10 14:54:20.688962	\N	0	13	\N
130	guest_token	1	{"path": "/api/v1/security/guest_token/", "object_ref": "SecurityRestApi.guest_token"}	2024-04-10 14:57:46.395154	\N	0	27	http://localhost:8088/api/v1/security/guest_token
100	DashboardRestApi.get_datasets	1	{"path": "/api/v1/dashboard/1/datasets", "url_rule": "/api/v1/dashboard/<id_or_slug>/datasets", "object_ref": "DashboardRestApi.get_datasets", "id_or_slug": "1"}	2024-04-10 14:53:31.037346	\N	0	54	http://localhost:8088/superset/dashboard/1/?edit=true&native_filters_key=x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q
104	ChartDataRestApi.data	1	{"path": "/api/v1/chart/data", "form_data": {"filters": [{"col": "created_at", "op": "TEMPORAL_RANGE", "val": "No filter"}], "extras": {"having": "", "where": ""}, "applied_time_extras": {}, "columns": [], "metrics": ["count"], "annotation_layers": [], "series_limit": 0, "order_desc": true, "url_params": {"native_filters_key": "x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q"}, "custom_params": {}, "custom_form_data": {}, "slice_id": 1}, "dashboard_id": "1", "object_ref": "ChartDataRestApi.data"}	2024-04-10 14:53:31.420993	1	1	22	http://localhost:8088/superset/dashboard/1/?edit=true&native_filters_key=x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q
111	DashboardRestApi.favorite_status	1	{"path": "/api/v1/dashboard/favorite_status/", "q": "!(1)", "object_ref": "DashboardRestApi.favorite_status", "rison": [1]}	2024-04-10 14:53:49.972443	\N	0	3	http://localhost:8088/dashboard/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc&viewMode=table
116	DashboardRestApi.favorite_status	1	{"path": "/api/v1/dashboard/favorite_status/", "q": "!(1)", "object_ref": "DashboardRestApi.favorite_status", "rison": [1]}	2024-04-10 14:53:56.818071	\N	0	18	http://localhost:8088/superset/dashboard/1/
119	DashboardFilterStateRestApi.post	1	{"path": "/api/v1/dashboard/1/filter_state", "tab_id": "7", "url_rule": "/api/v1/dashboard/<int:pk>/filter_state", "object_ref": "DashboardFilterStateRestApi.post", "pk": 1}	2024-04-10 14:53:57.341071	\N	0	35	http://localhost:8088/superset/dashboard/1/
101	DashboardFilterStateRestApi.get	1	{"path": "/api/v1/dashboard/1/filter_state/x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q", "url_rule": "/api/v1/dashboard/<int:pk>/filter_state/<string:key>", "object_ref": "DashboardFilterStateRestApi.get", "pk": 1, "key": "x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q"}	2024-04-10 14:53:31.057752	\N	0	10	http://localhost:8088/superset/dashboard/1/?edit=true&native_filters_key=x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q
113	DashboardRestApi.get	1	{"path": "/api/v1/dashboard/1", "url_rule": "/api/v1/dashboard/<id_or_slug>", "object_ref": "DashboardRestApi.get", "dashboard_id": 1}	2024-04-10 14:53:56.682615	1	0	218	http://localhost:8088/superset/dashboard/1/
125	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "g8yZI7qhm", "version": "v2", "ts": 1712760845469, "event_name": "force_refresh_chart", "slice_id": 1, "is_cached": false, "event_type": "user", "event_id": "kKV0BO3x3", "visibility": "visible"}	2024-04-10 14:54:06.63644	1	0	0	http://localhost:8088/superset/dashboard/1/?native_filters_key=x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q
126	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "g8yZI7qhm", "version": "v2", "ts": 1712760845593, "event_name": "load_chart", "slice_id": 1, "applied_filters": [{"column": "created_at"}], "is_cached": null, "force_refresh": true, "row_count": 1, "datasource": "1__table", "start_offset": 9044, "duration": 124, "has_extra_filters": false, "viz_type": "big_number_total", "data_age": null, "event_type": "timing", "trigger_event": "kKV0BO3x3"}	2024-04-10 14:54:06.636451	1	0	0	http://localhost:8088/superset/dashboard/1/?native_filters_key=x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q
127	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "g8yZI7qhm", "version": "v2", "ts": 1712760845609, "event_name": "render_chart", "slice_id": 1, "viz_type": "big_number_total", "start_offset": 9179, "duration": 4, "event_type": "timing", "trigger_event": "kKV0BO3x3"}	2024-04-10 14:54:06.636456	1	0	0	http://localhost:8088/superset/dashboard/1/?native_filters_key=x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q
129	csrf_token	1	{"path": "/api/v1/security/csrf_token/", "object_ref": "SecurityRestApi.csrf_token"}	2024-04-10 14:57:46.343416	\N	0	20	\N
102	DashboardRestApi.favorite_status	1	{"path": "/api/v1/dashboard/favorite_status/", "q": "!(1)", "object_ref": "DashboardRestApi.favorite_status", "rison": [1]}	2024-04-10 14:53:31.337495	\N	0	10	http://localhost:8088/superset/dashboard/1/?edit=true&native_filters_key=x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q
109	DashboardRestApi.get_list	1	{"path": "/api/v1/dashboard/", "q": "(order_column:changed_on_delta_humanized,order_direction:desc,page:0,page_size:25)", "rison": {"order_column": "changed_on_delta_humanized", "order_direction": "desc", "page": 0, "page_size": 25}}	2024-04-10 14:53:49.879555	\N	0	29	http://localhost:8088/dashboard/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc&viewMode=table
114	DashboardRestApi.get_charts	1	{"path": "/api/v1/dashboard/1/charts", "url_rule": "/api/v1/dashboard/<id_or_slug>/charts", "object_ref": "DashboardRestApi.get_charts", "id_or_slug": "1"}	2024-04-10 14:53:56.684632	\N	0	216	http://localhost:8088/superset/dashboard/1/
118	ChartDataRestApi.data	1	{"path": "/api/v1/chart/data", "form_data": {"filters": [{"col": "created_at", "op": "TEMPORAL_RANGE", "val": "No filter"}], "extras": {"having": "", "where": ""}, "applied_time_extras": {}, "columns": [], "metrics": ["count"], "annotation_layers": [], "series_limit": 0, "order_desc": true, "url_params": {}, "custom_params": {}, "custom_form_data": {}, "slice_id": 1}, "dashboard_id": "1", "object_ref": "ChartDataRestApi.data"}	2024-04-10 14:53:56.883183	1	1	22	http://localhost:8088/superset/dashboard/1/
124	ChartDataRestApi.data	1	{"path": "/api/v1/chart/data", "form_data": {"filters": [{"col": "created_at", "op": "TEMPORAL_RANGE", "val": "No filter"}], "extras": {"having": "", "where": ""}, "applied_time_extras": {}, "columns": [], "metrics": ["count"], "annotation_layers": [], "series_limit": 0, "order_desc": true, "url_params": {}, "custom_params": {}, "custom_form_data": {}, "slice_id": 1}, "dashboard_id": "1", "force": "true", "object_ref": "ChartDataRestApi.data"}	2024-04-10 14:54:05.587835	1	1	66	http://localhost:8088/superset/dashboard/1/?native_filters_key=x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q
105	log	1	{"source": "sqlLab", "ts": 1712760810830, "event_name": "spa_navigation", "path": "/superset/dashboard/1/", "event_type": "user", "event_id": "jRyREAUYhg", "visibility": "visible"}	2024-04-10 14:53:32.524881	\N	0	0	http://localhost:8088/superset/dashboard/1/?edit=true&native_filters_key=x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q
106	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "g8yZI7qhm", "version": "v2", "ts": 1712760811074, "event_name": "mount_dashboard", "is_soft_navigation": false, "is_edit_mode": true, "mount_duration": 641, "is_empty": false, "is_published": false, "event_type": "user", "event_id": "dtAcbCfrlD", "visibility": "visible"}	2024-04-10 14:53:32.524895	\N	0	0	http://localhost:8088/superset/dashboard/1/?edit=true&native_filters_key=x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q
107	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "g8yZI7qhm", "version": "v2", "ts": 1712760811457, "event_name": "load_chart", "slice_id": 1, "applied_filters": [{"column": "created_at"}], "is_cached": true, "force_refresh": false, "row_count": 1, "datasource": "1__table", "start_offset": 919, "duration": 105, "has_extra_filters": false, "viz_type": "big_number_total", "data_age": 63458, "event_type": "timing", "trigger_event": "dtAcbCfrlD"}	2024-04-10 14:53:32.524901	\N	0	0	http://localhost:8088/superset/dashboard/1/?edit=true&native_filters_key=x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q
108	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "g8yZI7qhm", "version": "v2", "ts": 1712760811492, "event_name": "render_chart", "slice_id": 1, "viz_type": "big_number_total", "start_offset": 1032, "duration": 27, "event_type": "timing", "trigger_event": "dtAcbCfrlD"}	2024-04-10 14:53:32.524906	\N	0	0	http://localhost:8088/superset/dashboard/1/?edit=true&native_filters_key=x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q
112	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "g8yZI7qhm", "version": "v2", "ts": 1712760829753, "event_name": "spa_navigation", "path": "/dashboard/list/", "event_type": "user", "event_id": "UxnhaCgvs", "visibility": "visible"}	2024-04-10 14:53:50.785998	1	0	0	http://localhost:8088/dashboard/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc&viewMode=table
117	ChartRestApi.get_list	1	{"path": "/api/v1/chart/", "q": "(columns:!(changed_on_delta_humanized,changed_on_utc,datasource_id,datasource_type,datasource_url,datasource_name_text,description_markeddown,description,id,params,slice_name,thumbnail_url,url,viz_type,owners.id,created_by.id),filters:!((col:viz_type,opr:neq,value:filter_box),(col:owners,opr:rel_m_m,value:1)),order_column:changed_on_delta_humanized,order_direction:desc,page_size:200)", "rison": {"columns": ["changed_on_delta_humanized", "changed_on_utc", "datasource_id", "datasource_type", "datasource_url", "datasource_name_text", "description_markeddown", "description", "id", "params", "slice_name", "thumbnail_url", "url", "viz_type", "owners.id", "created_by.id"], "filters": [{"col": "viz_type", "opr": "neq", "value": "filter_box"}, {"col": "owners", "opr": "rel_m_m", "value": 1}], "order_column": "changed_on_delta_humanized", "order_direction": "desc", "page_size": 200}}	2024-04-10 14:53:56.856728	\N	0	44	http://localhost:8088/superset/dashboard/1/
120	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "g8yZI7qhm", "version": "v2", "ts": 1712760836426, "event_name": "spa_navigation", "path": "/superset/dashboard/1/", "event_type": "user", "event_id": "zVYwakCwmF", "visibility": "visible"}	2024-04-10 14:53:57.921803	1	0	0	http://localhost:8088/superset/dashboard/1/?native_filters_key=x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q
121	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "g8yZI7qhm", "version": "v2", "ts": 1712760836741, "event_name": "mount_dashboard", "is_soft_navigation": true, "is_edit_mode": true, "mount_duration": 316, "is_empty": false, "is_published": false, "event_type": "user", "event_id": "8jNmPGwoP", "visibility": "visible"}	2024-04-10 14:53:57.921808	1	0	0	http://localhost:8088/superset/dashboard/1/?native_filters_key=x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q
122	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "g8yZI7qhm", "version": "v2", "ts": 1712760836890, "event_name": "load_chart", "slice_id": 1, "applied_filters": [{"column": "created_at"}], "is_cached": true, "force_refresh": false, "row_count": 1, "datasource": "1__table", "start_offset": 373, "duration": 92, "has_extra_filters": false, "viz_type": "big_number_total", "data_age": 88890, "event_type": "timing", "trigger_event": "8jNmPGwoP"}	2024-04-10 14:53:57.921811	1	0	0	http://localhost:8088/superset/dashboard/1/?native_filters_key=x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q
123	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "g8yZI7qhm", "version": "v2", "ts": 1712760836905, "event_name": "render_chart", "slice_id": 1, "viz_type": "big_number_total", "start_offset": 476, "duration": 3, "event_type": "timing", "trigger_event": "8jNmPGwoP"}	2024-04-10 14:53:57.921813	1	0	0	http://localhost:8088/superset/dashboard/1/?native_filters_key=x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q
131	DashboardRestApi.get_charts	1	{"path": "/api/v1/dashboard/1/charts", "url_rule": "/api/v1/dashboard/<id_or_slug>/charts", "object_ref": "DashboardRestApi.get_charts", "id_or_slug": "1"}	2024-04-11 02:27:45.75467	\N	0	35	http://localhost:8088/superset/dashboard/1/
132	DashboardRestApi.get	1	{"path": "/api/v1/dashboard/1", "url_rule": "/api/v1/dashboard/<id_or_slug>", "object_ref": "DashboardRestApi.get", "dashboard_id": 1}	2024-04-11 02:27:45.762477	1	0	57	http://localhost:8088/superset/dashboard/1/
133	DashboardRestApi.get_datasets	1	{"path": "/api/v1/dashboard/1/datasets", "url_rule": "/api/v1/dashboard/<id_or_slug>/datasets", "object_ref": "DashboardRestApi.get_datasets", "id_or_slug": "1"}	2024-04-11 02:27:45.864435	\N	0	133	http://localhost:8088/superset/dashboard/1/
134	DashboardRestApi.favorite_status	1	{"path": "/api/v1/dashboard/favorite_status/", "q": "!(1)", "object_ref": "DashboardRestApi.favorite_status", "rison": [1]}	2024-04-11 02:27:45.877242	\N	0	7	http://localhost:8088/superset/dashboard/1/
135	ChartDataRestApi.data	1	{"path": "/api/v1/chart/data", "form_data": {"filters": [{"col": "created_at", "op": "TEMPORAL_RANGE", "val": "No filter"}], "extras": {"having": "", "where": ""}, "applied_time_extras": {}, "columns": [], "metrics": ["count"], "annotation_layers": [], "series_limit": 0, "order_desc": true, "url_params": {}, "custom_params": {}, "custom_form_data": {}, "slice_id": 1}, "dashboard_id": "1", "object_ref": "ChartDataRestApi.data"}	2024-04-11 02:27:45.985712	1	1	78	http://localhost:8088/superset/dashboard/1/
136	DashboardFilterStateRestApi.post	1	{"path": "/api/v1/dashboard/1/filter_state", "tab_id": "6", "url_rule": "/api/v1/dashboard/<int:pk>/filter_state", "object_ref": "DashboardFilterStateRestApi.post", "pk": 1}	2024-04-11 02:27:46.405676	\N	0	44	http://localhost:8088/superset/dashboard/1/
137	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "0FoKHdlmc", "version": "v2", "ts": 1712802466298, "event_name": "spa_navigation", "path": "/superset/dashboard/1/", "event_type": "user", "event_id": "VdY-z1jATL", "visibility": "visible"}	2024-04-11 02:27:47.710286	1	0	0	http://localhost:8088/superset/dashboard/1/?native_filters_key=cl_uPkOPnIh9QFIc5MFgDJotsMHxCL2s9ujLF8ubAXA1OV6RSQShouIwhHG1VxPY
138	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "0FoKHdlmc", "version": "v2", "ts": 1712802466478, "event_name": "mount_dashboard", "is_soft_navigation": true, "is_edit_mode": false, "mount_duration": 179, "is_empty": false, "is_published": false, "event_type": "user", "event_id": "KMHjcNaCj6", "visibility": "visible"}	2024-04-11 02:27:47.710296	1	0	0	http://localhost:8088/superset/dashboard/1/?native_filters_key=cl_uPkOPnIh9QFIc5MFgDJotsMHxCL2s9ujLF8ubAXA1OV6RSQShouIwhHG1VxPY
139	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "0FoKHdlmc", "version": "v2", "ts": 1712802466661, "event_name": "load_chart", "slice_id": 1, "applied_filters": [{"column": "created_at"}], "is_cached": null, "force_refresh": false, "row_count": 1, "datasource": "1__table", "start_offset": 223, "duration": 140, "has_extra_filters": false, "viz_type": "big_number_total", "data_age": null, "event_type": "timing", "trigger_event": "KMHjcNaCj6"}	2024-04-11 02:27:47.710301	1	0	0	http://localhost:8088/superset/dashboard/1/?native_filters_key=cl_uPkOPnIh9QFIc5MFgDJotsMHxCL2s9ujLF8ubAXA1OV6RSQShouIwhHG1VxPY
140	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "0FoKHdlmc", "version": "v2", "ts": 1712802466683, "event_name": "render_chart", "slice_id": 1, "viz_type": "big_number_total", "start_offset": 378, "duration": 7, "event_type": "timing", "trigger_event": "KMHjcNaCj6"}	2024-04-11 02:27:47.710305	1	0	0	http://localhost:8088/superset/dashboard/1/?native_filters_key=cl_uPkOPnIh9QFIc5MFgDJotsMHxCL2s9ujLF8ubAXA1OV6RSQShouIwhHG1VxPY
142	ReportScheduleRestApi.get_list	1	{"path": "/api/v1/report/", "q": "(filters:!((col:dashboard_id,opr:eq,value:1),(col:creation_method,opr:eq,value:dashboards),(col:created_by,opr:rel_o_m,value:1)))", "rison": {"filters": [{"col": "dashboard_id", "opr": "eq", "value": 1}, {"col": "creation_method", "opr": "eq", "value": "dashboards"}, {"col": "created_by", "opr": "rel_o_m", "value": 1}]}}	2024-04-11 02:28:08.662301	\N	0	28	http://localhost:8088/superset/dashboard/1/?native_filters_key=cl_uPkOPnIh9QFIc5MFgDJotsMHxCL2s9ujLF8ubAXA1OV6RSQShouIwhHG1VxPY
145	DashboardRestApi.favorite_status	1	{"path": "/api/v1/dashboard/favorite_status/", "q": "!(1)", "object_ref": "DashboardRestApi.favorite_status", "rison": [1]}	2024-04-11 02:28:34.472544	\N	0	2	http://localhost:8088/dashboard/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc&viewMode=table
149	DashboardRestApi.get_datasets	1	{"path": "/api/v1/dashboard/1/datasets", "url_rule": "/api/v1/dashboard/<id_or_slug>/datasets", "object_ref": "DashboardRestApi.get_datasets", "id_or_slug": "1"}	2024-04-11 02:28:46.381138	\N	0	37	http://localhost:8088/superset/dashboard/1/
153	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "0FoKHdlmc", "version": "v2", "ts": 1712802526305, "event_name": "spa_navigation", "path": "/superset/dashboard/1/", "event_type": "user", "event_id": "vXWP1CL0LS", "visibility": "visible"}	2024-04-11 02:28:47.570243	1	0	0	http://localhost:8088/superset/dashboard/1/?native_filters_key=cl_uPkOPnIh9QFIc5MFgDJotsMHxCL2s9ujLF8ubAXA1OV6RSQShouIwhHG1VxPY
154	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "0FoKHdlmc", "version": "v2", "ts": 1712802526406, "event_name": "mount_dashboard", "is_soft_navigation": true, "is_edit_mode": false, "mount_duration": 101, "is_empty": false, "is_published": false, "event_type": "user", "event_id": "_Vh3X8j-Dp", "visibility": "visible"}	2024-04-11 02:28:47.570262	1	0	0	http://localhost:8088/superset/dashboard/1/?native_filters_key=cl_uPkOPnIh9QFIc5MFgDJotsMHxCL2s9ujLF8ubAXA1OV6RSQShouIwhHG1VxPY
155	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "0FoKHdlmc", "version": "v2", "ts": 1712802526520, "event_name": "load_chart", "slice_id": 1, "applied_filters": [{"column": "created_at"}], "is_cached": true, "force_refresh": false, "row_count": 1, "datasource": "1__table", "start_offset": 133, "duration": 82, "has_extra_filters": false, "viz_type": "big_number_total", "data_age": 61520, "event_type": "timing", "trigger_event": "_Vh3X8j-Dp"}	2024-04-11 02:28:47.570271	1	0	0	http://localhost:8088/superset/dashboard/1/?native_filters_key=cl_uPkOPnIh9QFIc5MFgDJotsMHxCL2s9ujLF8ubAXA1OV6RSQShouIwhHG1VxPY
156	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "0FoKHdlmc", "version": "v2", "ts": 1712802526533, "event_name": "render_chart", "slice_id": 1, "viz_type": "big_number_total", "start_offset": 223, "duration": 5, "event_type": "timing", "trigger_event": "_Vh3X8j-Dp"}	2024-04-11 02:28:47.570278	1	0	0	http://localhost:8088/superset/dashboard/1/?native_filters_key=cl_uPkOPnIh9QFIc5MFgDJotsMHxCL2s9ujLF8ubAXA1OV6RSQShouIwhHG1VxPY
157	dashboard	1	{"path": "/superset/dashboard/2/", "edit": "true", "url_rule": "/superset/dashboard/<dashboard_id_or_slug>/", "object_ref": "Superset.dashboard", "dashboard_id_or_slug": "2", "dashboard_id": 2, "dashboard_version": "v2", "dash_edit_perm": true, "edit_mode": true}	2024-04-11 02:28:52.843791	2	0	549	http://localhost:8088/superset/dashboard/1/?native_filters_key=cl_uPkOPnIh9QFIc5MFgDJotsMHxCL2s9ujLF8ubAXA1OV6RSQShouIwhHG1VxPY
160	DashboardRestApi.get	1	{"path": "/api/v1/dashboard/2", "url_rule": "/api/v1/dashboard/<id_or_slug>", "object_ref": "DashboardRestApi.get", "dashboard_id": 2}	2024-04-11 02:28:53.507399	2	0	16	http://localhost:8088/superset/dashboard/2/?edit=true
165	log	1	{"source": "sqlLab", "ts": 1712802533335, "event_name": "spa_navigation", "path": "/superset/dashboard/2/", "event_type": "user", "event_id": "W6XP_-Y1j8", "visibility": "visible"}	2024-04-11 02:28:54.567752	\N	0	0	http://localhost:8088/superset/dashboard/2/?edit=true
166	log	1	{"source": "dashboard", "source_id": 2, "impression_id": "1ICOUe_oW", "version": "v2", "ts": 1712802533538, "event_name": "mount_dashboard", "is_soft_navigation": false, "is_edit_mode": true, "mount_duration": 1320, "is_empty": true, "is_published": false, "event_type": "user", "event_id": "dLuBzzArZ", "visibility": "visible"}	2024-04-11 02:28:54.567762	\N	0	0	http://localhost:8088/superset/dashboard/2/?edit=true
141	ReportScheduleRestApi.get_list	1	{"path": "/api/v1/report/", "q": "(filters:!((col:dashboard_id,opr:eq,value:1),(col:creation_method,opr:eq,value:dashboards),(col:created_by,opr:rel_o_m,value:1)))", "rison": {"filters": [{"col": "dashboard_id", "opr": "eq", "value": 1}, {"col": "creation_method", "opr": "eq", "value": "dashboards"}, {"col": "created_by", "opr": "rel_o_m", "value": 1}]}}	2024-04-11 02:27:49.613333	\N	0	37	http://localhost:8088/superset/dashboard/1/?native_filters_key=cl_uPkOPnIh9QFIc5MFgDJotsMHxCL2s9ujLF8ubAXA1OV6RSQShouIwhHG1VxPY
143	DashboardRestApi.info	1	{"path": "/api/v1/dashboard/_info", "q": "(keys:!(permissions))", "rison": {"keys": ["permissions"]}}	2024-04-11 02:28:34.35292	\N	0	38	http://localhost:8088/dashboard/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc&viewMode=table
146	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "0FoKHdlmc", "version": "v2", "ts": 1712802514256, "event_name": "spa_navigation", "path": "/dashboard/list/", "event_type": "user", "event_id": "W9gMlbR1w", "visibility": "visible"}	2024-04-11 02:28:35.280101	1	0	0	http://localhost:8088/dashboard/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc&viewMode=table
147	DashboardRestApi.get_charts	1	{"path": "/api/v1/dashboard/1/charts", "url_rule": "/api/v1/dashboard/<id_or_slug>/charts", "object_ref": "DashboardRestApi.get_charts", "id_or_slug": "1"}	2024-04-11 02:28:46.359217	\N	0	14	http://localhost:8088/superset/dashboard/1/
150	DashboardRestApi.favorite_status	1	{"path": "/api/v1/dashboard/favorite_status/", "q": "!(1)", "object_ref": "DashboardRestApi.favorite_status", "rison": [1]}	2024-04-11 02:28:46.444232	\N	0	3	http://localhost:8088/superset/dashboard/1/
158	DatabaseRestApi.get_list	1	{"path": "/api/v1/database/", "q": "(filters:!((col:allow_file_upload,opr:upload_is_enabled,value:!t)))", "rison": {"filters": [{"col": "allow_file_upload", "opr": "upload_is_enabled", "value": true}]}}	2024-04-11 02:28:53.441374	\N	0	66	http://localhost:8088/superset/dashboard/2/?edit=true
164	ChartRestApi.get_list	1	{"path": "/api/v1/chart/", "q": "(columns:!(changed_on_delta_humanized,changed_on_utc,datasource_id,datasource_type,datasource_url,datasource_name_text,description_markeddown,description,id,params,slice_name,thumbnail_url,url,viz_type,owners.id,created_by.id),filters:!((col:viz_type,opr:neq,value:filter_box),(col:owners,opr:rel_m_m,value:1)),order_column:changed_on_delta_humanized,order_direction:desc,page_size:200)", "rison": {"columns": ["changed_on_delta_humanized", "changed_on_utc", "datasource_id", "datasource_type", "datasource_url", "datasource_name_text", "description_markeddown", "description", "id", "params", "slice_name", "thumbnail_url", "url", "viz_type", "owners.id", "created_by.id"], "filters": [{"col": "viz_type", "opr": "neq", "value": "filter_box"}, {"col": "owners", "opr": "rel_m_m", "value": 1}], "order_column": "changed_on_delta_humanized", "order_direction": "desc", "page_size": 200}}	2024-04-11 02:28:53.833961	\N	0	71	http://localhost:8088/superset/dashboard/2/?edit=true
144	DashboardRestApi.get_list	1	{"path": "/api/v1/dashboard/", "q": "(order_column:changed_on_delta_humanized,order_direction:desc,page:0,page_size:25)", "rison": {"order_column": "changed_on_delta_humanized", "order_direction": "desc", "page": 0, "page_size": 25}}	2024-04-11 02:28:34.391653	\N	0	71	http://localhost:8088/dashboard/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc&viewMode=table
148	DashboardRestApi.get	1	{"path": "/api/v1/dashboard/1", "url_rule": "/api/v1/dashboard/<id_or_slug>", "object_ref": "DashboardRestApi.get", "dashboard_id": 1}	2024-04-11 02:28:46.364243	1	0	18	http://localhost:8088/superset/dashboard/1/
151	ChartDataRestApi.data	1	{"path": "/api/v1/chart/data", "form_data": {"filters": [{"col": "created_at", "op": "TEMPORAL_RANGE", "val": "No filter"}], "extras": {"having": "", "where": ""}, "applied_time_extras": {}, "columns": [], "metrics": ["count"], "annotation_layers": [], "series_limit": 0, "order_desc": true, "url_params": {}, "custom_params": {}, "custom_form_data": {}, "slice_id": 1}, "dashboard_id": "1", "object_ref": "ChartDataRestApi.data"}	2024-04-11 02:28:46.500051	1	1	31	http://localhost:8088/superset/dashboard/1/
152	DashboardFilterStateRestApi.post	1	{"path": "/api/v1/dashboard/1/filter_state", "tab_id": "6", "url_rule": "/api/v1/dashboard/<int:pk>/filter_state", "object_ref": "DashboardFilterStateRestApi.post", "pk": 1}	2024-04-11 02:28:46.993225	\N	0	52	http://localhost:8088/superset/dashboard/1/
159	DatabaseRestApi.get_list	1	{"path": "/api/v1/database/", "q": "(filters:!((col:database_name,opr:neq,value:examples)))", "rison": {"filters": [{"col": "database_name", "opr": "neq", "value": "examples"}]}}	2024-04-11 02:28:53.479535	\N	0	82	http://localhost:8088/superset/dashboard/2/?edit=true
161	DashboardRestApi.get_datasets	1	{"path": "/api/v1/dashboard/2/datasets", "url_rule": "/api/v1/dashboard/<id_or_slug>/datasets", "object_ref": "DashboardRestApi.get_datasets", "id_or_slug": "2"}	2024-04-11 02:28:53.51294	\N	0	6	http://localhost:8088/superset/dashboard/2/?edit=true
162	DashboardRestApi.get_charts	1	{"path": "/api/v1/dashboard/2/charts", "url_rule": "/api/v1/dashboard/<id_or_slug>/charts", "object_ref": "DashboardRestApi.get_charts", "id_or_slug": "2"}	2024-04-11 02:28:53.517155	\N	0	7	http://localhost:8088/superset/dashboard/2/?edit=true
163	DashboardRestApi.favorite_status	1	{"path": "/api/v1/dashboard/favorite_status/", "q": "!(2)", "object_ref": "DashboardRestApi.favorite_status", "rison": [2]}	2024-04-11 02:28:53.776139	\N	0	10	http://localhost:8088/superset/dashboard/2/?edit=true
167	DashboardRestApi.info	1	{"path": "/api/v1/dashboard/_info", "q": "(keys:!(permissions))", "rison": {"keys": ["permissions"]}}	2024-04-11 02:28:58.665267	\N	0	31	http://localhost:8088/dashboard/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc&viewMode=table
168	DashboardRestApi.get_list	1	{"path": "/api/v1/dashboard/", "q": "(order_column:changed_on_delta_humanized,order_direction:desc,page:0,page_size:25)", "rison": {"order_column": "changed_on_delta_humanized", "order_direction": "desc", "page": 0, "page_size": 25}}	2024-04-11 02:28:58.675162	\N	0	38	http://localhost:8088/dashboard/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc&viewMode=table
169	DashboardRestApi.favorite_status	1	{"path": "/api/v1/dashboard/favorite_status/", "q": "!(2,1)", "object_ref": "DashboardRestApi.favorite_status", "rison": [2, 1]}	2024-04-11 02:28:58.797579	\N	0	1	http://localhost:8088/dashboard/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc&viewMode=table
170	log	1	{"source": "dashboard", "source_id": 2, "impression_id": "1ICOUe_oW", "version": "v2", "ts": 1712802538509, "event_name": "spa_navigation", "path": "/dashboard/list/", "event_type": "user", "event_id": "mFa9ZCyMs", "visibility": "visible"}	2024-04-11 02:28:59.521982	2	0	0	http://localhost:8088/dashboard/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc&viewMode=table
171	DashboardRestApi.delete	1	{"path": "/api/v1/dashboard/2", "url_rule": "/api/v1/dashboard/<pk>", "object_ref": "DashboardRestApi.delete", "pk": "2"}	2024-04-11 02:29:12.338499	\N	0	41	http://localhost:8088/dashboard/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc&viewMode=table
172	DashboardRestApi.get_list	1	{"path": "/api/v1/dashboard/", "q": "(order_column:changed_on_delta_humanized,order_direction:desc,page:0,page_size:25)", "rison": {"order_column": "changed_on_delta_humanized", "order_direction": "desc", "page": 0, "page_size": 25}}	2024-04-11 02:29:12.455381	\N	0	28	http://localhost:8088/dashboard/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc&viewMode=table
173	DashboardRestApi.favorite_status	1	{"path": "/api/v1/dashboard/favorite_status/", "q": "!(1)", "object_ref": "DashboardRestApi.favorite_status", "rison": [1]}	2024-04-11 02:29:12.59146	\N	0	2	http://localhost:8088/dashboard/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc&viewMode=table
174	DashboardRestApi.get_charts	1	{"path": "/api/v1/dashboard/1/charts", "url_rule": "/api/v1/dashboard/<id_or_slug>/charts", "object_ref": "DashboardRestApi.get_charts", "id_or_slug": "1"}	2024-04-11 02:29:15.174009	\N	0	28	http://localhost:8088/superset/dashboard/1/
175	DashboardRestApi.get	1	{"path": "/api/v1/dashboard/1", "url_rule": "/api/v1/dashboard/<id_or_slug>", "object_ref": "DashboardRestApi.get", "dashboard_id": 1}	2024-04-11 02:29:15.180467	1	0	45	http://localhost:8088/superset/dashboard/1/
176	DashboardRestApi.get_datasets	1	{"path": "/api/v1/dashboard/1/datasets", "url_rule": "/api/v1/dashboard/<id_or_slug>/datasets", "object_ref": "DashboardRestApi.get_datasets", "id_or_slug": "1"}	2024-04-11 02:29:15.212275	\N	0	72	http://localhost:8088/superset/dashboard/1/
177	DashboardRestApi.favorite_status	1	{"path": "/api/v1/dashboard/favorite_status/", "q": "!(2)", "object_ref": "DashboardRestApi.favorite_status", "rison": [2]}	2024-04-11 02:29:15.328954	\N	0	5	http://localhost:8088/superset/dashboard/1/
178	DashboardRestApi.favorite_status	1	{"path": "/api/v1/dashboard/favorite_status/", "q": "!(1)", "object_ref": "DashboardRestApi.favorite_status", "rison": [1]}	2024-04-11 02:29:15.361385	\N	0	13	http://localhost:8088/superset/dashboard/1/
179	ChartRestApi.get_list	1	{"path": "/api/v1/chart/", "q": "(columns:!(changed_on_delta_humanized,changed_on_utc,datasource_id,datasource_type,datasource_url,datasource_name_text,description_markeddown,description,id,params,slice_name,thumbnail_url,url,viz_type,owners.id,created_by.id),filters:!((col:viz_type,opr:neq,value:filter_box),(col:owners,opr:rel_m_m,value:1)),order_column:changed_on_delta_humanized,order_direction:desc,page_size:200)", "rison": {"columns": ["changed_on_delta_humanized", "changed_on_utc", "datasource_id", "datasource_type", "datasource_url", "datasource_name_text", "description_markeddown", "description", "id", "params", "slice_name", "thumbnail_url", "url", "viz_type", "owners.id", "created_by.id"], "filters": [{"col": "viz_type", "opr": "neq", "value": "filter_box"}, {"col": "owners", "opr": "rel_m_m", "value": 1}], "order_column": "changed_on_delta_humanized", "order_direction": "desc", "page_size": 200}}	2024-04-11 02:29:15.391608	\N	0	52	http://localhost:8088/superset/dashboard/1/
180	ChartDataRestApi.data	1	{"path": "/api/v1/chart/data", "form_data": {"filters": [{"col": "created_at", "op": "TEMPORAL_RANGE", "val": "No filter"}], "extras": {"having": "", "where": ""}, "applied_time_extras": {}, "columns": [], "metrics": ["count"], "annotation_layers": [], "series_limit": 0, "order_desc": true, "url_params": {}, "custom_params": {}, "custom_form_data": {}, "slice_id": 1}, "dashboard_id": "1", "object_ref": "ChartDataRestApi.data"}	2024-04-11 02:29:15.472919	1	1	25	http://localhost:8088/superset/dashboard/1/
189	DashboardRestApi.info	1	{"path": "/api/v1/dashboard/_info", "q": "(keys:!(permissions))", "rison": {"keys": ["permissions"]}}	2024-04-11 02:30:06.766206	\N	0	35	http://localhost:8088/dashboard/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc&viewMode=table
193	DashboardRestApi.get_charts	1	{"path": "/api/v1/dashboard/1/charts", "url_rule": "/api/v1/dashboard/<id_or_slug>/charts", "object_ref": "DashboardRestApi.get_charts", "id_or_slug": "1"}	2024-04-11 02:30:23.93716	\N	0	14	http://localhost:8088/superset/dashboard/1/
198	ChartDataRestApi.data	1	{"path": "/api/v1/chart/data", "form_data": {"filters": [{"col": "created_at", "op": "TEMPORAL_RANGE", "val": "No filter"}], "extras": {"having": "", "where": ""}, "applied_time_extras": {}, "columns": [], "metrics": ["count"], "annotation_layers": [], "series_limit": 0, "order_desc": true, "url_params": {}, "custom_params": {}, "custom_form_data": {}, "slice_id": 1}, "dashboard_id": "1", "object_ref": "ChartDataRestApi.data"}	2024-04-11 02:30:24.142542	1	1	29	http://localhost:8088/superset/dashboard/1/
210	ChartRestApi.get_list	1	{"path": "/api/v1/chart/", "q": "(columns:!(changed_on_delta_humanized,changed_on_utc,datasource_id,datasource_type,datasource_url,datasource_name_text,description_markeddown,description,id,params,slice_name,thumbnail_url,url,viz_type,owners.id,created_by.id),filters:!((col:viz_type,opr:neq,value:filter_box),(col:owners,opr:rel_m_m,value:1)),order_column:changed_on_delta_humanized,order_direction:desc,page_size:200)", "rison": {"columns": ["changed_on_delta_humanized", "changed_on_utc", "datasource_id", "datasource_type", "datasource_url", "datasource_name_text", "description_markeddown", "description", "id", "params", "slice_name", "thumbnail_url", "url", "viz_type", "owners.id", "created_by.id"], "filters": [{"col": "viz_type", "opr": "neq", "value": "filter_box"}, {"col": "owners", "opr": "rel_m_m", "value": 1}], "order_column": "changed_on_delta_humanized", "order_direction": "desc", "page_size": 200}}	2024-04-11 02:36:58.700723	\N	0	39	http://localhost:8088/superset/dashboard/1/?native_filters_key=cl_uPkOPnIh9QFIc5MFgDJotsMHxCL2s9ujLF8ubAXA1OV6RSQShouIwhHG1VxPY
215	ReportScheduleRestApi.get_list	1	{"path": "/api/v1/report/", "q": "(filters:!((col:dashboard_id,opr:eq,value:1),(col:creation_method,opr:eq,value:dashboards),(col:created_by,opr:rel_o_m,value:1)))", "rison": {"filters": [{"col": "dashboard_id", "opr": "eq", "value": 1}, {"col": "creation_method", "opr": "eq", "value": "dashboards"}, {"col": "created_by", "opr": "rel_o_m", "value": 1}]}}	2024-04-11 02:38:09.663127	\N	0	26	http://localhost:8088/superset/dashboard/1/
221	ReportScheduleRestApi.get_list	1	{"path": "/api/v1/report/", "q": "(filters:!((col:dashboard_id,opr:eq,value:1),(col:creation_method,opr:eq,value:dashboards),(col:created_by,opr:rel_o_m,value:1)))", "rison": {"filters": [{"col": "dashboard_id", "opr": "eq", "value": 1}, {"col": "creation_method", "opr": "eq", "value": "dashboards"}, {"col": "created_by", "opr": "rel_o_m", "value": 1}]}}	2024-04-11 02:38:49.888454	\N	0	31	http://localhost:8088/superset/dashboard/1/?native_filters_key=x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q
223	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "g8yZI7qhm", "version": "v2", "ts": 1712803139216, "event_name": "hide_browser_tab", "start_offset": 1558523, "duration": 3738, "event_type": "timing", "trigger_event": "kKV0BO3x3"}	2024-04-11 02:39:03.992269	1	0	0	http://localhost:8088/superset/dashboard/1/?native_filters_key=x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q
229	DashboardRestApi.get	1	{"path": "/api/v1/dashboard/1", "url_rule": "/api/v1/dashboard/<id_or_slug>", "object_ref": "DashboardRestApi.get", "dashboard_id": 1}	2024-04-11 03:55:42.445007	1	0	37	http://localhost:8088/superset/dashboard/1/?edit=true
234	DashboardRestApi.favorite_status	1	{"path": "/api/v1/dashboard/favorite_status/", "q": "!(1)", "object_ref": "DashboardRestApi.favorite_status", "rison": [1]}	2024-04-11 03:55:42.899619	\N	0	10	http://localhost:8088/superset/dashboard/1/?edit=true
240	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "60B7TEFNo", "version": "v2", "ts": 1712807744419, "event_name": "load_chart", "slice_id": 1, "applied_filters": [{"column": "created_at"}], "is_cached": null, "force_refresh": false, "row_count": 1, "datasource": "1__table", "start_offset": 16212, "duration": 112, "has_extra_filters": false, "viz_type": "big_number_total", "data_age": null, "event_type": "timing", "trigger_event": "3E-VaL_2p"}	2024-04-11 03:55:46.357182	1	0	0	http://localhost:8088/superset/dashboard/1/?edit=true
241	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "60B7TEFNo", "version": "v2", "ts": 1712807744459, "event_name": "render_chart", "slice_id": 1, "viz_type": "big_number_total", "start_offset": 16333, "duration": 32, "event_type": "timing", "trigger_event": "3E-VaL_2p"}	2024-04-11 03:55:46.35719	1	0	0	http://localhost:8088/superset/dashboard/1/?edit=true
182	log	1	{"source": "dashboard", "source_id": 2, "impression_id": "1ICOUe_oW", "version": "v2", "ts": 1712802555074, "event_name": "spa_navigation", "path": "/superset/dashboard/1/", "event_type": "user", "event_id": "oMR-jTnMZM", "visibility": "visible"}	2024-04-11 02:29:16.574193	2	0	0	http://localhost:8088/superset/dashboard/1/?native_filters_key=cl_uPkOPnIh9QFIc5MFgDJotsMHxCL2s9ujLF8ubAXA1OV6RSQShouIwhHG1VxPY
183	log	1	{"source": "dashboard", "source_id": 2, "impression_id": "1ICOUe_oW", "version": "v2", "ts": 1712802555240, "event_name": "mount_dashboard", "is_soft_navigation": true, "is_edit_mode": true, "mount_duration": 167, "is_empty": true, "is_published": false, "event_type": "user", "event_id": "YMyY_OPFl_", "visibility": "visible"}	2024-04-11 02:29:16.574211	2	0	0	http://localhost:8088/superset/dashboard/1/?native_filters_key=cl_uPkOPnIh9QFIc5MFgDJotsMHxCL2s9ujLF8ubAXA1OV6RSQShouIwhHG1VxPY
184	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "1ICOUe_oW", "version": "v2", "ts": 1712802555481, "event_name": "load_chart", "slice_id": 1, "applied_filters": [{"column": "created_at"}], "is_cached": true, "force_refresh": false, "row_count": 1, "datasource": "1__table", "start_offset": 341, "duration": 67, "has_extra_filters": false, "viz_type": "big_number_total", "data_age": 90481, "event_type": "timing", "trigger_event": "YMyY_OPFl_"}	2024-04-11 02:29:16.574218	2	0	0	http://localhost:8088/superset/dashboard/1/?native_filters_key=cl_uPkOPnIh9QFIc5MFgDJotsMHxCL2s9ujLF8ubAXA1OV6RSQShouIwhHG1VxPY
185	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "1ICOUe_oW", "version": "v2", "ts": 1712802555533, "event_name": "render_chart", "slice_id": 1, "viz_type": "big_number_total", "start_offset": 426, "duration": 33, "event_type": "timing", "trigger_event": "YMyY_OPFl_"}	2024-04-11 02:29:16.574225	2	0	0	http://localhost:8088/superset/dashboard/1/?native_filters_key=cl_uPkOPnIh9QFIc5MFgDJotsMHxCL2s9ujLF8ubAXA1OV6RSQShouIwhHG1VxPY
190	DashboardRestApi.get_list	1	{"path": "/api/v1/dashboard/", "q": "(order_column:changed_on_delta_humanized,order_direction:desc,page:0,page_size:25)", "rison": {"order_column": "changed_on_delta_humanized", "order_direction": "desc", "page": 0, "page_size": 25}}	2024-04-11 02:30:06.768717	\N	0	39	http://localhost:8088/dashboard/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc&viewMode=table
195	DashboardRestApi.get_datasets	1	{"path": "/api/v1/dashboard/1/datasets", "url_rule": "/api/v1/dashboard/<id_or_slug>/datasets", "object_ref": "DashboardRestApi.get_datasets", "id_or_slug": "1"}	2024-04-11 02:30:23.966933	\N	0	42	http://localhost:8088/superset/dashboard/1/
204	csrf_token	1	{"path": "/api/v1/security/csrf_token/", "object_ref": "SecurityRestApi.csrf_token"}	2024-04-11 02:32:30.222102	\N	0	23	\N
207	guest_token	1	{"path": "/api/v1/security/guest_token/", "object_ref": "SecurityRestApi.guest_token"}	2024-04-11 02:34:49.724098	\N	0	21	http://localhost:8088/api/v1/security/guest_token
212	DatabaseRestApi.get_list	1	{"path": "/api/v1/database/", "q": "(filters:!((col:database_name,opr:neq,value:examples)))", "rison": {"filters": [{"col": "database_name", "opr": "neq", "value": "examples"}]}}	2024-04-11 02:37:57.685893	\N	0	55	http://localhost:8088/superset/dashboard/1/?native_filters_key=cl_uPkOPnIh9QFIc5MFgDJotsMHxCL2s9ujLF8ubAXA1OV6RSQShouIwhHG1VxPY
226	dashboard	1	{"path": "/superset/dashboard/1/", "edit": "true", "url_rule": "/superset/dashboard/<dashboard_id_or_slug>/", "object_ref": "Superset.dashboard", "dashboard_id_or_slug": "1", "dashboard_id": 1, "dashboard_version": "v2", "dash_edit_perm": true, "edit_mode": true}	2024-04-11 03:55:28.445363	1	0	175	http://localhost:8088/dashboard/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc&viewMode=table
230	DashboardRestApi.get_charts	1	{"path": "/api/v1/dashboard/1/charts", "url_rule": "/api/v1/dashboard/<id_or_slug>/charts", "object_ref": "DashboardRestApi.get_charts", "id_or_slug": "1"}	2024-04-11 03:55:42.486901	\N	0	37	http://localhost:8088/superset/dashboard/1/?edit=true
235	ChartRestApi.get_list	1	{"path": "/api/v1/chart/", "q": "(columns:!(changed_on_delta_humanized,changed_on_utc,datasource_id,datasource_type,datasource_url,datasource_name_text,description_markeddown,description,id,params,slice_name,thumbnail_url,url,viz_type,owners.id,created_by.id),filters:!((col:viz_type,opr:neq,value:filter_box),(col:owners,opr:rel_m_m,value:1)),order_column:changed_on_delta_humanized,order_direction:desc,page_size:200)", "rison": {"columns": ["changed_on_delta_humanized", "changed_on_utc", "datasource_id", "datasource_type", "datasource_url", "datasource_name_text", "description_markeddown", "description", "id", "params", "slice_name", "thumbnail_url", "url", "viz_type", "owners.id", "created_by.id"], "filters": [{"col": "viz_type", "opr": "neq", "value": "filter_box"}, {"col": "owners", "opr": "rel_m_m", "value": 1}], "order_column": "changed_on_delta_humanized", "order_direction": "desc", "page_size": 200}}	2024-04-11 03:55:42.968524	\N	0	60	http://localhost:8088/superset/dashboard/1/?edit=true
186	ReportScheduleRestApi.get_list	1	{"path": "/api/v1/report/", "q": "(filters:!((col:dashboard_id,opr:eq,value:1),(col:creation_method,opr:eq,value:dashboards),(col:created_by,opr:rel_o_m,value:1)))", "rison": {"filters": [{"col": "dashboard_id", "opr": "eq", "value": 1}, {"col": "creation_method", "opr": "eq", "value": "dashboards"}, {"col": "created_by", "opr": "rel_o_m", "value": 1}]}}	2024-04-11 02:29:23.545684	\N	0	21	http://localhost:8088/superset/dashboard/1/?native_filters_key=cl_uPkOPnIh9QFIc5MFgDJotsMHxCL2s9ujLF8ubAXA1OV6RSQShouIwhHG1VxPY
191	DashboardRestApi.favorite_status	1	{"path": "/api/v1/dashboard/favorite_status/", "q": "!(1)", "object_ref": "DashboardRestApi.favorite_status", "rison": [1]}	2024-04-11 02:30:06.847879	\N	0	3	http://localhost:8088/dashboard/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc&viewMode=table
197	ChartRestApi.get_list	1	{"path": "/api/v1/chart/", "q": "(columns:!(changed_on_delta_humanized,changed_on_utc,datasource_id,datasource_type,datasource_url,datasource_name_text,description_markeddown,description,id,params,slice_name,thumbnail_url,url,viz_type,owners.id,created_by.id),filters:!((col:viz_type,opr:neq,value:filter_box),(col:owners,opr:rel_m_m,value:1)),order_column:changed_on_delta_humanized,order_direction:desc,page_size:200)", "rison": {"columns": ["changed_on_delta_humanized", "changed_on_utc", "datasource_id", "datasource_type", "datasource_url", "datasource_name_text", "description_markeddown", "description", "id", "params", "slice_name", "thumbnail_url", "url", "viz_type", "owners.id", "created_by.id"], "filters": [{"col": "viz_type", "opr": "neq", "value": "filter_box"}, {"col": "owners", "opr": "rel_m_m", "value": 1}], "order_column": "changed_on_delta_humanized", "order_direction": "desc", "page_size": 200}}	2024-04-11 02:30:24.125751	\N	0	71	http://localhost:8088/superset/dashboard/1/
200	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "1ICOUe_oW", "version": "v2", "ts": 1712802623887, "event_name": "spa_navigation", "path": "/superset/dashboard/1/", "event_type": "user", "event_id": "2qwXKFTKc5", "visibility": "visible"}	2024-04-11 02:30:25.215085	1	0	0	http://localhost:8088/superset/dashboard/1/?native_filters_key=cl_uPkOPnIh9QFIc5MFgDJotsMHxCL2s9ujLF8ubAXA1OV6RSQShouIwhHG1VxPY
201	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "1ICOUe_oW", "version": "v2", "ts": 1712802623987, "event_name": "mount_dashboard", "is_soft_navigation": true, "is_edit_mode": true, "mount_duration": 99, "is_empty": false, "is_published": false, "event_type": "user", "event_id": "GVgIOlIgMV", "visibility": "visible"}	2024-04-11 02:30:25.215094	1	0	0	http://localhost:8088/superset/dashboard/1/?native_filters_key=cl_uPkOPnIh9QFIc5MFgDJotsMHxCL2s9ujLF8ubAXA1OV6RSQShouIwhHG1VxPY
202	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "1ICOUe_oW", "version": "v2", "ts": 1712802624173, "event_name": "load_chart", "slice_id": 1, "applied_filters": [{"column": "created_at"}], "is_cached": true, "force_refresh": false, "row_count": 1, "datasource": "1__table", "start_offset": 154, "duration": 132, "has_extra_filters": false, "viz_type": "big_number_total", "data_age": 159173, "event_type": "timing", "trigger_event": "GVgIOlIgMV"}	2024-04-11 02:30:25.215097	1	0	0	http://localhost:8088/superset/dashboard/1/?native_filters_key=cl_uPkOPnIh9QFIc5MFgDJotsMHxCL2s9ujLF8ubAXA1OV6RSQShouIwhHG1VxPY
203	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "1ICOUe_oW", "version": "v2", "ts": 1712802624192, "event_name": "render_chart", "slice_id": 1, "viz_type": "big_number_total", "start_offset": 299, "duration": 6, "event_type": "timing", "trigger_event": "GVgIOlIgMV"}	2024-04-11 02:30:25.2151	1	0	0	http://localhost:8088/superset/dashboard/1/?native_filters_key=cl_uPkOPnIh9QFIc5MFgDJotsMHxCL2s9ujLF8ubAXA1OV6RSQShouIwhHG1VxPY
209	guest_token	1	{"path": "/api/v1/security/guest_token/", "object_ref": "SecurityRestApi.guest_token"}	2024-04-11 02:35:30.130754	\N	0	24	http://localhost:8088/api/v1/security/guest_token
214	DashboardRestApi.put	1	{"path": "/api/v1/dashboard/1", "url_rule": "/api/v1/dashboard/<pk>", "object_ref": "DashboardRestApi.put", "pk": "1"}	2024-04-11 02:38:09.522442	\N	0	37	http://localhost:8088/superset/dashboard/1/?native_filters_key=cl_uPkOPnIh9QFIc5MFgDJotsMHxCL2s9ujLF8ubAXA1OV6RSQShouIwhHG1VxPY
217	ReportScheduleRestApi.get_list	1	{"path": "/api/v1/report/", "q": "(filters:!((col:dashboard_id,opr:eq,value:1),(col:creation_method,opr:eq,value:dashboards),(col:created_by,opr:rel_o_m,value:1)))", "rison": {"filters": [{"col": "dashboard_id", "opr": "eq", "value": 1}, {"col": "creation_method", "opr": "eq", "value": "dashboards"}, {"col": "created_by", "opr": "rel_o_m", "value": 1}]}}	2024-04-11 02:38:17.893517	\N	0	28	http://localhost:8088/superset/dashboard/1/?native_filters_key=cl_uPkOPnIh9QFIc5MFgDJotsMHxCL2s9ujLF8ubAXA1OV6RSQShouIwhHG1VxPY
220	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "1ICOUe_oW", "version": "v2", "ts": 1712803106757, "event_name": "hide_browser_tab", "start_offset": 482870, "duration": 20313, "event_type": "timing", "trigger_event": "D36Vp9WMc"}	2024-04-11 02:38:48.719221	1	0	0	http://localhost:8088/superset/dashboard/1/?native_filters_key=cl_uPkOPnIh9QFIc5MFgDJotsMHxCL2s9ujLF8ubAXA1OV6RSQShouIwhHG1VxPY
222	DashboardPermalinkRestApi.post	1	{"path": "/api/v1/dashboard/1/permalink", "url_rule": "/api/v1/dashboard/<pk>/permalink", "object_ref": "DashboardPermalinkRestApi.post", "pk": "1"}	2024-04-11 02:38:58.415452	\N	0	18	http://localhost:8088/superset/dashboard/1/?native_filters_key=x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q
225	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "5a9QGcRun", "version": "v2", "ts": 1712737304996, "event_name": "hide_browser_tab", "start_offset": 51002061, "duration": 27374244, "event_type": "timing", "trigger_event": "V4soE0yb9b"}	2024-04-11 03:55:24.972354	1	0	0	http://localhost:8088/superset/dashboard/1/?edit=true
231	DatabaseRestApi.get_list	1	{"path": "/api/v1/database/", "q": "(filters:!((col:database_name,opr:ct,value:'')),order_columns:database_name,order_direction:asc,page:0,page_size:100)", "rison": {"filters": [{"col": "database_name", "opr": "ct", "value": ""}], "order_columns": "database_name", "order_direction": "asc", "page": 0, "page_size": 100}}	2024-04-11 03:55:42.491548	\N	0	100	http://localhost:8088/dataset/add/
236	log	1	{"source": "sqlLab", "ts": 1712807742218, "event_name": "spa_navigation", "path": "/dataset/add/", "event_type": "user", "event_id": "Qd6KcG4M1", "visibility": "visible"}	2024-04-11 03:55:43.258245	\N	0	0	http://localhost:8088/dataset/add/
187	ChartRestApi.get_list	1	{"path": "/api/v1/chart/", "q": "(columns:!(changed_on_delta_humanized,changed_on_utc,datasource_id,datasource_type,datasource_url,datasource_name_text,description_markeddown,description,id,params,slice_name,thumbnail_url,url,viz_type,owners.id,created_by.id),filters:!((col:viz_type,opr:neq,value:filter_box),(col:owners,opr:rel_m_m,value:1)),order_column:changed_on_delta_humanized,order_direction:desc,page_size:200)", "rison": {"columns": ["changed_on_delta_humanized", "changed_on_utc", "datasource_id", "datasource_type", "datasource_url", "datasource_name_text", "description_markeddown", "description", "id", "params", "slice_name", "thumbnail_url", "url", "viz_type", "owners.id", "created_by.id"], "filters": [{"col": "viz_type", "opr": "neq", "value": "filter_box"}, {"col": "owners", "opr": "rel_m_m", "value": 1}], "order_column": "changed_on_delta_humanized", "order_direction": "desc", "page_size": 200}}	2024-04-11 02:29:25.241256	\N	0	38	http://localhost:8088/superset/dashboard/1/?native_filters_key=cl_uPkOPnIh9QFIc5MFgDJotsMHxCL2s9ujLF8ubAXA1OV6RSQShouIwhHG1VxPY
192	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "1ICOUe_oW", "version": "v2", "ts": 1712802606659, "event_name": "spa_navigation", "path": "/dashboard/list/", "event_type": "user", "event_id": "BcsPn972p", "visibility": "visible"}	2024-04-11 02:30:07.683602	1	0	0	http://localhost:8088/dashboard/list/?pageIndex=0&sortColumn=changed_on_delta_humanized&sortOrder=desc&viewMode=table
196	DashboardRestApi.favorite_status	1	{"path": "/api/v1/dashboard/favorite_status/", "q": "!(1)", "object_ref": "DashboardRestApi.favorite_status", "rison": [1]}	2024-04-11 02:30:24.06902	\N	0	19	http://localhost:8088/superset/dashboard/1/
199	DashboardFilterStateRestApi.post	1	{"path": "/api/v1/dashboard/1/filter_state", "tab_id": "6", "url_rule": "/api/v1/dashboard/<int:pk>/filter_state", "object_ref": "DashboardFilterStateRestApi.post", "pk": 1}	2024-04-11 02:30:24.595936	\N	0	40	http://localhost:8088/superset/dashboard/1/
205	guest_token	1	{"path": "/api/v1/security/guest_token/", "object_ref": "SecurityRestApi.guest_token"}	2024-04-11 02:32:30.258172	\N	0	17	http://localhost:8088/api/v1/security/guest_token
208	csrf_token	1	{"path": "/api/v1/security/csrf_token/", "object_ref": "SecurityRestApi.csrf_token"}	2024-04-11 02:35:30.085099	\N	0	14	\N
216	DashboardFilterStateRestApi.post	1	{"path": "/api/v1/dashboard/1/filter_state", "tab_id": "6", "url_rule": "/api/v1/dashboard/<int:pk>/filter_state", "object_ref": "DashboardFilterStateRestApi.post", "pk": 1}	2024-04-11 02:38:10.162692	\N	0	42	http://localhost:8088/superset/dashboard/1/
218	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "g8yZI7qhm", "version": "v2", "ts": 1712760884033, "event_name": "hide_browser_tab", "start_offset": 47607, "duration": 1498140, "event_type": "timing", "trigger_event": "kKV0BO3x3"}	2024-04-11 02:38:48.417247	1	0	0	http://localhost:8088/superset/dashboard/1/?native_filters_key=x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q
219	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "g8yZI7qhm", "version": "v2", "ts": 1712803127371, "event_name": "hide_browser_tab", "start_offset": 1546678, "duration": 34, "event_type": "timing", "trigger_event": "kKV0BO3x3"}	2024-04-11 02:38:48.417253	1	0	0	http://localhost:8088/superset/dashboard/1/?native_filters_key=x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q
227	DatabaseRestApi.get_list	1	{"path": "/api/v1/database/", "q": "(filters:!((col:database_name,opr:neq,value:examples)))", "rison": {"filters": [{"col": "database_name", "opr": "neq", "value": "examples"}]}}	2024-04-11 03:55:42.3288	\N	0	59	http://localhost:8088/dataset/add/
233	DashboardRestApi.get_datasets	1	{"path": "/api/v1/dashboard/1/datasets", "url_rule": "/api/v1/dashboard/<id_or_slug>/datasets", "object_ref": "DashboardRestApi.get_datasets", "id_or_slug": "1"}	2024-04-11 03:55:42.556033	\N	0	126	http://localhost:8088/superset/dashboard/1/?edit=true
237	log	1	{"source": "sqlLab", "ts": 1712807742232, "event_name": "spa_navigation", "path": "/superset/dashboard/1/", "event_type": "user", "event_id": "M1oXuXX04", "visibility": "visible"}	2024-04-11 03:55:44.350841	\N	0	0	http://localhost:8088/superset/dashboard/1/?edit=true
238	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "60B7TEFNo", "version": "v2", "ts": 1712807742580, "event_name": "mount_dashboard", "is_soft_navigation": false, "is_edit_mode": true, "mount_duration": 14485, "is_empty": false, "is_published": false, "event_type": "user", "event_id": "3E-VaL_2p", "visibility": "hidden"}	2024-04-11 03:55:44.350847	\N	0	0	http://localhost:8088/superset/dashboard/1/?edit=true
188	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "1ICOUe_oW", "version": "v2", "ts": 1712802565129, "event_name": "toggle_edit_dashboard", "edit_mode": true, "event_type": "user", "event_id": "jRfkvYe0-", "visibility": "visible"}	2024-04-11 02:29:26.153511	1	0	0	http://localhost:8088/superset/dashboard/1/?native_filters_key=cl_uPkOPnIh9QFIc5MFgDJotsMHxCL2s9ujLF8ubAXA1OV6RSQShouIwhHG1VxPY
194	DashboardRestApi.get	1	{"path": "/api/v1/dashboard/1", "url_rule": "/api/v1/dashboard/<id_or_slug>", "object_ref": "DashboardRestApi.get", "dashboard_id": 1}	2024-04-11 02:30:23.939726	1	0	17	http://localhost:8088/superset/dashboard/1/
206	csrf_token	1	{"path": "/api/v1/security/csrf_token/", "object_ref": "SecurityRestApi.csrf_token"}	2024-04-11 02:34:49.677492	\N	0	10	\N
211	log	1	{"source": "dashboard", "source_id": 1, "impression_id": "1ICOUe_oW", "version": "v2", "ts": 1712803018581, "event_name": "toggle_edit_dashboard", "edit_mode": true, "event_type": "user", "event_id": "D36Vp9WMc", "visibility": "visible"}	2024-04-11 02:36:59.611476	1	0	0	http://localhost:8088/superset/dashboard/1/?native_filters_key=cl_uPkOPnIh9QFIc5MFgDJotsMHxCL2s9ujLF8ubAXA1OV6RSQShouIwhHG1VxPY
213	DatabaseRestApi.get_list	1	{"path": "/api/v1/database/", "q": "(filters:!((col:allow_file_upload,opr:upload_is_enabled,value:!t)))", "rison": {"filters": [{"col": "allow_file_upload", "opr": "upload_is_enabled", "value": true}]}}	2024-04-11 02:37:57.697323	\N	0	67	http://localhost:8088/superset/dashboard/1/?native_filters_key=cl_uPkOPnIh9QFIc5MFgDJotsMHxCL2s9ujLF8ubAXA1OV6RSQShouIwhHG1VxPY
224	ReportScheduleRestApi.get_list	1	{"path": "/api/v1/report/", "q": "(filters:!((col:dashboard_id,opr:eq,value:1),(col:creation_method,opr:eq,value:dashboards),(col:created_by,opr:rel_o_m,value:1)))", "rison": {"filters": [{"col": "dashboard_id", "opr": "eq", "value": 1}, {"col": "creation_method", "opr": "eq", "value": "dashboards"}, {"col": "created_by", "opr": "rel_o_m", "value": 1}]}}	2024-04-11 02:39:11.13406	\N	0	23	http://localhost:8088/superset/dashboard/1/?native_filters_key=x9QvU2WV4EtOwZWWGOechbm0k7zoZh0tDG5iYf091Opd12EwlK1RZUa5obbblW8Q
228	DatabaseRestApi.get_list	1	{"path": "/api/v1/database/", "q": "(filters:!((col:allow_file_upload,opr:upload_is_enabled,value:!t)))", "rison": {"filters": [{"col": "allow_file_upload", "opr": "upload_is_enabled", "value": true}]}}	2024-04-11 03:55:42.341231	\N	0	54	http://localhost:8088/dataset/add/
232	DatabaseRestApi.get_list	1	{"path": "/api/v1/database/", "q": "(filters:!((col:allow_file_upload,opr:upload_is_enabled,value:!t)))", "rison": {"filters": [{"col": "allow_file_upload", "opr": "upload_is_enabled", "value": true}]}}	2024-04-11 03:55:42.550297	\N	0	75	http://localhost:8088/superset/dashboard/1/?edit=true
239	ChartDataRestApi.data	1	{"path": "/api/v1/chart/data", "form_data": {"filters": [{"col": "created_at", "op": "TEMPORAL_RANGE", "val": "No filter"}], "extras": {"having": "", "where": ""}, "applied_time_extras": {}, "columns": [], "metrics": ["count"], "annotation_layers": [], "series_limit": 0, "order_desc": true, "url_params": {}, "custom_params": {}, "custom_form_data": {}, "slice_id": 1}, "dashboard_id": "1", "object_ref": "ChartDataRestApi.data"}	2024-04-11 03:55:44.412301	1	1	57	http://localhost:8088/superset/dashboard/1/?edit=true
\.


--
-- Data for Name: query; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.query (id, client_id, database_id, tmp_table_name, tab_name, sql_editor_id, user_id, status, schema, sql, select_sql, executed_sql, "limit", select_as_cta, select_as_cta_used, progress, rows, error_message, start_time, changed_on, end_time, results_key, start_running_time, end_result_backend_time, tracking_url, extra_json, tmp_schema_name, ctas_method, limiting_factor) FROM stdin;
\.


--
-- Data for Name: report_execution_log; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.report_execution_log (id, scheduled_dttm, start_dttm, end_dttm, value, value_row_json, state, error_message, report_schedule_id, uuid) FROM stdin;
\.


--
-- Data for Name: report_recipient; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.report_recipient (id, type, recipient_config_json, report_schedule_id, created_on, changed_on, created_by_fk, changed_by_fk) FROM stdin;
\.


--
-- Data for Name: report_schedule; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.report_schedule (id, type, name, description, context_markdown, active, crontab, sql, chart_id, dashboard_id, database_id, last_eval_dttm, last_state, last_value, last_value_row_json, validator_type, validator_config_json, log_retention, grace_period, created_on, changed_on, created_by_fk, changed_by_fk, working_timeout, report_format, creation_method, timezone, extra_json, force_screenshot, custom_width, custom_height) FROM stdin;
\.


--
-- Data for Name: report_schedule_user; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.report_schedule_user (id, user_id, report_schedule_id) FROM stdin;
\.


--
-- Data for Name: rls_filter_roles; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.rls_filter_roles (id, role_id, rls_filter_id) FROM stdin;
\.


--
-- Data for Name: rls_filter_tables; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.rls_filter_tables (id, table_id, rls_filter_id) FROM stdin;
\.


--
-- Data for Name: row_level_security_filters; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.row_level_security_filters (created_on, changed_on, id, clause, created_by_fk, changed_by_fk, filter_type, group_key, name, description) FROM stdin;
\.


--
-- Data for Name: saved_query; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.saved_query (created_on, changed_on, id, user_id, db_id, label, schema, sql, description, changed_by_fk, created_by_fk, extra_json, last_run, rows, uuid, template_parameters) FROM stdin;
\.


--
-- Data for Name: sl_columns; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.sl_columns (uuid, created_on, changed_on, id, is_aggregation, is_additive, is_dimensional, is_filterable, is_increase_desired, is_managed_externally, is_partition, is_physical, is_temporal, is_spatial, name, type, unit, expression, description, warning_text, external_url, extra_json, created_by_fk, changed_by_fk, advanced_data_type) FROM stdin;
\.


--
-- Data for Name: sl_dataset_columns; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.sl_dataset_columns (dataset_id, column_id) FROM stdin;
\.


--
-- Data for Name: sl_dataset_tables; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.sl_dataset_tables (dataset_id, table_id) FROM stdin;
\.


--
-- Data for Name: sl_dataset_users; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.sl_dataset_users (dataset_id, user_id) FROM stdin;
\.


--
-- Data for Name: sl_datasets; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.sl_datasets (uuid, created_on, changed_on, id, database_id, is_physical, is_managed_externally, name, expression, external_url, extra_json, created_by_fk, changed_by_fk) FROM stdin;
\.


--
-- Data for Name: sl_table_columns; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.sl_table_columns (table_id, column_id) FROM stdin;
\.


--
-- Data for Name: sl_tables; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.sl_tables (uuid, created_on, changed_on, id, database_id, is_managed_externally, catalog, schema, name, external_url, extra_json, created_by_fk, changed_by_fk) FROM stdin;
\.


--
-- Data for Name: slice_user; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.slice_user (id, user_id, slice_id) FROM stdin;
1	1	1
\.


--
-- Data for Name: slices; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.slices (created_on, changed_on, id, slice_name, datasource_type, datasource_name, viz_type, params, created_by_fk, changed_by_fk, description, cache_timeout, perm, datasource_id, schema_perm, uuid, query_context, last_saved_at, last_saved_by_fk, certified_by, certification_details, is_managed_externally, external_url) FROM stdin;
2024-04-10 14:52:44.743922	2024-04-10 14:52:44.743933	1	putni	table	public.warrants	big_number_total	{"datasource":"1__table","viz_type":"big_number_total","metric":"count","adhoc_filters":[{"clause":"WHERE","subject":"created_at","operator":"TEMPORAL_RANGE","comparator":"No filter","expressionType":"SIMPLE"}],"header_font_size":0.4,"subheader_font_size":0.15,"y_axis_format":"SMART_NUMBER","time_format":"smart_date","extra_form_data":{},"dashboards":[]}	1	1	\N	\N	[PostgreSQL].[warrants](id:1)	1	[PostgreSQL].[public]	7461d16f-603c-4e01-8d62-214b777cb5ef	{"datasource":{"id":1,"type":"table"},"force":false,"queries":[{"filters":[{"col":"created_at","op":"TEMPORAL_RANGE","val":"No filter"}],"extras":{"having":"","where":""},"applied_time_extras":{},"columns":[],"metrics":["count"],"annotation_layers":[],"series_limit":0,"order_desc":true,"url_params":{},"custom_params":{},"custom_form_data":{}}],"form_data":{"datasource":"1__table","viz_type":"big_number_total","metric":"count","adhoc_filters":[{"clause":"WHERE","subject":"created_at","operator":"TEMPORAL_RANGE","comparator":"No filter","expressionType":"SIMPLE"}],"header_font_size":0.4,"subheader_font_size":0.15,"y_axis_format":"SMART_NUMBER","time_format":"smart_date","extra_form_data":{},"dashboards":[],"force":false,"result_format":"json","result_type":"full"},"result_format":"json","result_type":"full"}	2024-04-10 14:52:44.73462	1	\N	\N	f	\N
\.


--
-- Data for Name: sql_metrics; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.sql_metrics (created_on, changed_on, id, metric_name, verbose_name, metric_type, table_id, expression, description, created_by_fk, changed_by_fk, d3format, warning_text, extra, uuid, currency) FROM stdin;
2024-04-10 14:52:09.945705	2024-04-10 14:52:09.945711	1	count	COUNT(*)	count	1	COUNT(*)	\N	1	1	\N	\N	\N	33d2a276-3587-4983-ad58-3d337b0482f6	\N
\.


--
-- Data for Name: sqlatable_user; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.sqlatable_user (id, user_id, table_id) FROM stdin;
1	1	1
\.


--
-- Data for Name: ssh_tunnels; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.ssh_tunnels (created_on, changed_on, created_by_fk, changed_by_fk, extra_json, uuid, id, database_id, server_address, server_port, username, password, private_key, private_key_password) FROM stdin;
\.


--
-- Data for Name: tab_state; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.tab_state (created_on, changed_on, extra_json, id, user_id, label, active, database_id, schema, sql, query_limit, latest_query_id, autorun, template_params, created_by_fk, changed_by_fk, hide_left_bar, saved_query_id) FROM stdin;
\.


--
-- Data for Name: table_columns; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.table_columns (created_on, changed_on, id, table_id, column_name, is_dttm, is_active, type, groupby, filterable, description, created_by_fk, changed_by_fk, expression, verbose_name, python_date_format, uuid, extra, advanced_data_type) FROM stdin;
2024-04-10 14:52:09.933584	2024-04-10 14:52:09.933588	1	1	id	f	t	TEXT	t	t	\N	1	1	\N	\N	\N	e282c5b5-0664-4db1-ad10-3c72c3e92a47	\N	\N
2024-04-10 14:52:09.933624	2024-04-10 14:52:09.933625	2	1	is_deleted	f	t	BOOLEAN	t	t	\N	1	1	\N	\N	\N	dcde4767-c7b7-4c91-b651-2ed64a4c510d	\N	\N
2024-04-10 14:52:09.933645	2024-04-10 14:52:09.933646	3	1	is_active	f	t	BOOLEAN	t	t	\N	1	1	\N	\N	\N	fa5c198c-2ea4-4a38-a0dd-73c05f06d903	\N	\N
2024-04-10 14:52:09.933663	2024-04-10 14:52:09.933664	4	1	created_at	t	t	TIMESTAMP WITH TIME ZONE	t	t	\N	1	1	\N	\N	\N	bb1b7b56-7426-42af-b810-daf6898d3b1d	\N	\N
2024-04-10 14:52:09.93368	2024-04-10 14:52:09.933681	5	1	updated_at	t	t	TIMESTAMP WITH TIME ZONE	t	t	\N	1	1	\N	\N	\N	7423fd51-475b-490e-a3df-88ce8010ac0a	\N	\N
2024-04-10 14:52:09.933697	2024-04-10 14:52:09.933699	6	1	deleted_at	t	t	TIMESTAMP WITH TIME ZONE	t	t	\N	1	1	\N	\N	\N	5a4434f5-9d9a-477e-bfd4-0bdfca922656	\N	\N
2024-04-10 14:52:09.933715	2024-04-10 14:52:09.933716	7	1	issue_date	t	t	TIMESTAMP WITH TIME ZONE	t	t	\N	1	1	\N	\N	\N	d329ce87-455d-4c6d-b323-880ccee8d381	\N	\N
2024-04-10 14:52:09.933733	2024-04-10 14:52:09.933734	8	1	expected_start	t	t	TIMESTAMP WITH TIME ZONE	t	t	\N	1	1	\N	\N	\N	18c82ab9-cd93-4d4f-96bb-5d22210c3b2d	\N	\N
2024-04-10 14:52:09.933751	2024-04-10 14:52:09.933752	9	1	closing_date	t	t	TIMESTAMP WITH TIME ZONE	t	t	\N	1	1	\N	\N	\N	61540658-b5ab-4c07-bd87-aa80fe1a7230	\N	\N
2024-04-10 14:52:09.933768	2024-04-10 14:52:09.933769	10	1	driver_id	f	t	TEXT	t	t	\N	1	1	\N	\N	\N	4c5aaf6e-2169-4e7e-8704-4a8fdb7cd0f8	\N	\N
2024-04-10 14:52:09.933786	2024-04-10 14:52:09.933787	11	1	passengers	f	t	TEXT[]	t	t	\N	1	1	\N	\N	\N	5c222c46-8e38-4ba5-bde3-255e7eded5c9	\N	\N
2024-04-10 14:52:09.933803	2024-04-10 14:52:09.933804	12	1	vehicle_id	f	t	TEXT	t	t	\N	1	1	\N	\N	\N	fe995f2f-ddff-44da-a6ea-a37fcddb462a	\N	\N
2024-04-10 14:52:09.933821	2024-04-10 14:52:09.933822	13	1	trailer_id	f	t	TEXT	t	t	\N	1	1	\N	\N	\N	3fe06e52-9c56-4b8a-b7c2-6ba12196d605	\N	\N
2024-04-10 14:52:09.933838	2024-04-10 14:52:09.933839	14	1	company_id	f	t	TEXT	t	t	\N	1	1	\N	\N	\N	a357c9ad-00ec-4e0d-b4af-7c7de61939fe	\N	\N
2024-04-10 14:52:09.933856	2024-04-10 14:52:09.933857	15	1	dispatcher_id	f	t	TEXT	t	t	\N	1	1	\N	\N	\N	0bcbfcb3-6137-4bb9-ac23-4b10998b2a33	\N	\N
2024-04-10 14:52:09.933873	2024-04-10 14:52:09.933874	16	1	technical_correctness	f	t	TEXT	t	t	\N	1	1	\N	\N	\N	91c8e2d3-de9e-4163-be67-bae951d6744b	\N	\N
2024-04-10 14:52:09.933891	2024-04-10 14:52:09.933892	17	1	status	f	t	TEXT	t	t	\N	1	1	\N	\N	\N	3c4e0f74-8c4c-4c4a-8635-2351eb8c8090	\N	\N
2024-04-10 14:52:09.933908	2024-04-10 14:52:09.933909	18	1	name	f	t	TEXT	t	t	\N	1	1	\N	\N	\N	80b846b1-b963-49aa-b9fb-63ca7d6d4da1	\N	\N
2024-04-10 14:52:09.933926	2024-04-10 14:52:09.933927	19	1	note	f	t	TEXT	t	t	\N	1	1	\N	\N	\N	72b21b67-cd72-42c4-ac6f-1bd738086cc4	\N	\N
\.


--
-- Data for Name: table_schema; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.table_schema (created_on, changed_on, extra_json, id, tab_state_id, database_id, schema, "table", description, expanded, created_by_fk, changed_by_fk) FROM stdin;
\.


--
-- Data for Name: tables; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.tables (created_on, changed_on, id, table_name, main_dttm_col, default_endpoint, database_id, created_by_fk, changed_by_fk, "offset", description, is_featured, cache_timeout, schema, sql, params, perm, filter_select_enabled, fetch_values_predicate, is_sqllab_view, template_params, schema_perm, extra, uuid, is_managed_externally, external_url, normalize_columns, always_filter_main_dttm) FROM stdin;
2024-04-10 14:52:09.871912	2024-04-10 14:52:09.938716	1	warrants	created_at	\N	1	1	1	0	\N	f	\N	public	\N	\N	[PostgreSQL].[warrants](id:1)	t	\N	f	\N	[PostgreSQL].[public]	\N	cb417661-2689-4582-8036-b1121fd5ecab	f	\N	f	f
\.


--
-- Data for Name: tag; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.tag (created_on, changed_on, id, name, type, created_by_fk, changed_by_fk, description) FROM stdin;
\.


--
-- Data for Name: tagged_object; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.tagged_object (created_on, changed_on, id, tag_id, object_id, object_type, created_by_fk, changed_by_fk) FROM stdin;
\.


--
-- Data for Name: url; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.url (created_on, changed_on, id, url, created_by_fk, changed_by_fk) FROM stdin;
\.


--
-- Data for Name: user_attribute; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.user_attribute (created_on, changed_on, id, user_id, welcome_dashboard_id, created_by_fk, changed_by_fk) FROM stdin;
\.


--
-- Data for Name: user_favorite_tag; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.user_favorite_tag (user_id, tag_id) FROM stdin;
\.


--
-- Name: ab_permission_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.ab_permission_id_seq', 72, true);


--
-- Name: ab_permission_view_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.ab_permission_view_id_seq', 178, true);


--
-- Name: ab_permission_view_role_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.ab_permission_view_role_id_seq', 405, true);


--
-- Name: ab_register_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.ab_register_user_id_seq', 1, false);


--
-- Name: ab_role_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.ab_role_id_seq', 6, true);


--
-- Name: ab_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.ab_user_id_seq', 2, true);


--
-- Name: ab_user_role_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.ab_user_role_id_seq', 2, true);


--
-- Name: ab_view_menu_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.ab_view_menu_id_seq', 85, true);


--
-- Name: annotation_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.annotation_id_seq', 1, false);


--
-- Name: annotation_layer_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.annotation_layer_id_seq', 1, false);


--
-- Name: cache_keys_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.cache_keys_id_seq', 1, false);


--
-- Name: css_templates_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.css_templates_id_seq', 1, false);


--
-- Name: dashboard_roles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.dashboard_roles_id_seq', 1, false);


--
-- Name: dashboard_slices_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.dashboard_slices_id_seq', 1, true);


--
-- Name: dashboard_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.dashboard_user_id_seq', 3, true);


--
-- Name: dashboards_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.dashboards_id_seq', 2, true);


--
-- Name: dbs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.dbs_id_seq', 1, true);


--
-- Name: dynamic_plugin_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.dynamic_plugin_id_seq', 1, false);


--
-- Name: favstar_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.favstar_id_seq', 1, false);


--
-- Name: filter_sets_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.filter_sets_id_seq', 1, false);


--
-- Name: key_value_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.key_value_id_seq', 10, true);


--
-- Name: keyvalue_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.keyvalue_id_seq', 1, false);


--
-- Name: logs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.logs_id_seq', 241, true);


--
-- Name: query_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.query_id_seq', 1, false);


--
-- Name: report_execution_log_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.report_execution_log_id_seq', 1, false);


--
-- Name: report_recipient_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.report_recipient_id_seq', 1, false);


--
-- Name: report_schedule_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.report_schedule_id_seq', 1, false);


--
-- Name: report_schedule_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.report_schedule_user_id_seq', 1, false);


--
-- Name: rls_filter_roles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.rls_filter_roles_id_seq', 1, false);


--
-- Name: rls_filter_tables_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.rls_filter_tables_id_seq', 1, false);


--
-- Name: row_level_security_filters_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.row_level_security_filters_id_seq', 1, false);


--
-- Name: saved_query_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.saved_query_id_seq', 1, false);


--
-- Name: sl_columns_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.sl_columns_id_seq', 1, false);


--
-- Name: sl_datasets_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.sl_datasets_id_seq', 1, false);


--
-- Name: sl_tables_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.sl_tables_id_seq', 1, false);


--
-- Name: slice_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.slice_user_id_seq', 1, true);


--
-- Name: slices_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.slices_id_seq', 1, true);


--
-- Name: sql_metrics_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.sql_metrics_id_seq', 1, true);


--
-- Name: sqlatable_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.sqlatable_user_id_seq', 1, true);


--
-- Name: ssh_tunnels_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.ssh_tunnels_id_seq', 1, false);


--
-- Name: tab_state_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.tab_state_id_seq', 1, false);


--
-- Name: table_columns_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.table_columns_id_seq', 19, true);


--
-- Name: table_schema_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.table_schema_id_seq', 1, false);


--
-- Name: tables_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.tables_id_seq', 1, true);


--
-- Name: tag_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.tag_id_seq', 1, false);


--
-- Name: tagged_object_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.tagged_object_id_seq', 1, false);


--
-- Name: url_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.url_id_seq', 1, false);


--
-- Name: user_attribute_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.user_attribute_id_seq', 1, false);


--
-- Name: tables _customer_location_uc; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tables
    ADD CONSTRAINT _customer_location_uc UNIQUE (database_id, schema, table_name);


--
-- Name: ab_permission ab_permission_name_key; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ab_permission
    ADD CONSTRAINT ab_permission_name_key UNIQUE (name);


--
-- Name: ab_permission ab_permission_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ab_permission
    ADD CONSTRAINT ab_permission_pkey PRIMARY KEY (id);


--
-- Name: ab_permission_view ab_permission_view_permission_id_view_menu_id_key; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ab_permission_view
    ADD CONSTRAINT ab_permission_view_permission_id_view_menu_id_key UNIQUE (permission_id, view_menu_id);


--
-- Name: ab_permission_view ab_permission_view_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ab_permission_view
    ADD CONSTRAINT ab_permission_view_pkey PRIMARY KEY (id);


--
-- Name: ab_permission_view_role ab_permission_view_role_permission_view_id_role_id_key; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ab_permission_view_role
    ADD CONSTRAINT ab_permission_view_role_permission_view_id_role_id_key UNIQUE (permission_view_id, role_id);


--
-- Name: ab_permission_view_role ab_permission_view_role_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ab_permission_view_role
    ADD CONSTRAINT ab_permission_view_role_pkey PRIMARY KEY (id);


--
-- Name: ab_register_user ab_register_user_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ab_register_user
    ADD CONSTRAINT ab_register_user_pkey PRIMARY KEY (id);


--
-- Name: ab_register_user ab_register_user_username_key; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ab_register_user
    ADD CONSTRAINT ab_register_user_username_key UNIQUE (username);


--
-- Name: ab_role ab_role_name_key; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ab_role
    ADD CONSTRAINT ab_role_name_key UNIQUE (name);


--
-- Name: ab_role ab_role_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ab_role
    ADD CONSTRAINT ab_role_pkey PRIMARY KEY (id);


--
-- Name: ab_user ab_user_email_key; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ab_user
    ADD CONSTRAINT ab_user_email_key UNIQUE (email);


--
-- Name: ab_user ab_user_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ab_user
    ADD CONSTRAINT ab_user_pkey PRIMARY KEY (id);


--
-- Name: ab_user_role ab_user_role_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ab_user_role
    ADD CONSTRAINT ab_user_role_pkey PRIMARY KEY (id);


--
-- Name: ab_user_role ab_user_role_user_id_role_id_key; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ab_user_role
    ADD CONSTRAINT ab_user_role_user_id_role_id_key UNIQUE (user_id, role_id);


--
-- Name: ab_user ab_user_username_key; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ab_user
    ADD CONSTRAINT ab_user_username_key UNIQUE (username);


--
-- Name: ab_view_menu ab_view_menu_name_key; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ab_view_menu
    ADD CONSTRAINT ab_view_menu_name_key UNIQUE (name);


--
-- Name: ab_view_menu ab_view_menu_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ab_view_menu
    ADD CONSTRAINT ab_view_menu_pkey PRIMARY KEY (id);


--
-- Name: alembic_version alembic_version_pkc; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.alembic_version
    ADD CONSTRAINT alembic_version_pkc PRIMARY KEY (version_num);


--
-- Name: annotation_layer annotation_layer_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.annotation_layer
    ADD CONSTRAINT annotation_layer_pkey PRIMARY KEY (id);


--
-- Name: annotation annotation_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.annotation
    ADD CONSTRAINT annotation_pkey PRIMARY KEY (id);


--
-- Name: cache_keys cache_keys_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.cache_keys
    ADD CONSTRAINT cache_keys_pkey PRIMARY KEY (id);


--
-- Name: query client_id; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.query
    ADD CONSTRAINT client_id UNIQUE (client_id);


--
-- Name: css_templates css_templates_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.css_templates
    ADD CONSTRAINT css_templates_pkey PRIMARY KEY (id);


--
-- Name: dashboard_roles dashboard_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dashboard_roles
    ADD CONSTRAINT dashboard_roles_pkey PRIMARY KEY (id);


--
-- Name: dashboard_slices dashboard_slices_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dashboard_slices
    ADD CONSTRAINT dashboard_slices_pkey PRIMARY KEY (id);


--
-- Name: dashboard_user dashboard_user_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dashboard_user
    ADD CONSTRAINT dashboard_user_pkey PRIMARY KEY (id);


--
-- Name: dashboards dashboards_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dashboards
    ADD CONSTRAINT dashboards_pkey PRIMARY KEY (id);


--
-- Name: dbs dbs_database_name_key; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dbs
    ADD CONSTRAINT dbs_database_name_key UNIQUE (database_name);


--
-- Name: dbs dbs_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dbs
    ADD CONSTRAINT dbs_pkey PRIMARY KEY (id);


--
-- Name: dbs dbs_verbose_name_key; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dbs
    ADD CONSTRAINT dbs_verbose_name_key UNIQUE (verbose_name);


--
-- Name: dynamic_plugin dynamic_plugin_key_key; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dynamic_plugin
    ADD CONSTRAINT dynamic_plugin_key_key UNIQUE (key);


--
-- Name: dynamic_plugin dynamic_plugin_name_key; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dynamic_plugin
    ADD CONSTRAINT dynamic_plugin_name_key UNIQUE (name);


--
-- Name: dynamic_plugin dynamic_plugin_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dynamic_plugin
    ADD CONSTRAINT dynamic_plugin_pkey PRIMARY KEY (id);


--
-- Name: favstar favstar_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.favstar
    ADD CONSTRAINT favstar_pkey PRIMARY KEY (id);


--
-- Name: filter_sets filter_sets_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.filter_sets
    ADD CONSTRAINT filter_sets_pkey PRIMARY KEY (id);


--
-- Name: dashboards idx_unique_slug; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dashboards
    ADD CONSTRAINT idx_unique_slug UNIQUE (slug);


--
-- Name: key_value key_value_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.key_value
    ADD CONSTRAINT key_value_pkey PRIMARY KEY (id);


--
-- Name: keyvalue keyvalue_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.keyvalue
    ADD CONSTRAINT keyvalue_pkey PRIMARY KEY (id);


--
-- Name: logs logs_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.logs
    ADD CONSTRAINT logs_pkey PRIMARY KEY (id);


--
-- Name: query query_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.query
    ADD CONSTRAINT query_pkey PRIMARY KEY (id);


--
-- Name: report_execution_log report_execution_log_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.report_execution_log
    ADD CONSTRAINT report_execution_log_pkey PRIMARY KEY (id);


--
-- Name: report_recipient report_recipient_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.report_recipient
    ADD CONSTRAINT report_recipient_pkey PRIMARY KEY (id);


--
-- Name: report_schedule report_schedule_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.report_schedule
    ADD CONSTRAINT report_schedule_pkey PRIMARY KEY (id);


--
-- Name: report_schedule_user report_schedule_user_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.report_schedule_user
    ADD CONSTRAINT report_schedule_user_pkey PRIMARY KEY (id);


--
-- Name: rls_filter_roles rls_filter_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.rls_filter_roles
    ADD CONSTRAINT rls_filter_roles_pkey PRIMARY KEY (id);


--
-- Name: rls_filter_tables rls_filter_tables_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.rls_filter_tables
    ADD CONSTRAINT rls_filter_tables_pkey PRIMARY KEY (id);


--
-- Name: row_level_security_filters row_level_security_filters_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.row_level_security_filters
    ADD CONSTRAINT row_level_security_filters_pkey PRIMARY KEY (id);


--
-- Name: saved_query saved_query_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.saved_query
    ADD CONSTRAINT saved_query_pkey PRIMARY KEY (id);


--
-- Name: sl_columns sl_columns_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sl_columns
    ADD CONSTRAINT sl_columns_pkey PRIMARY KEY (id);


--
-- Name: sl_columns sl_columns_uuid_key; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sl_columns
    ADD CONSTRAINT sl_columns_uuid_key UNIQUE (uuid);


--
-- Name: sl_dataset_columns sl_dataset_columns_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sl_dataset_columns
    ADD CONSTRAINT sl_dataset_columns_pkey PRIMARY KEY (dataset_id, column_id);


--
-- Name: sl_dataset_tables sl_dataset_tables_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sl_dataset_tables
    ADD CONSTRAINT sl_dataset_tables_pkey PRIMARY KEY (dataset_id, table_id);


--
-- Name: sl_dataset_users sl_dataset_users_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sl_dataset_users
    ADD CONSTRAINT sl_dataset_users_pkey PRIMARY KEY (dataset_id, user_id);


--
-- Name: sl_datasets sl_datasets_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sl_datasets
    ADD CONSTRAINT sl_datasets_pkey PRIMARY KEY (id);


--
-- Name: sl_datasets sl_datasets_uuid_key; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sl_datasets
    ADD CONSTRAINT sl_datasets_uuid_key UNIQUE (uuid);


--
-- Name: sl_table_columns sl_table_columns_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sl_table_columns
    ADD CONSTRAINT sl_table_columns_pkey PRIMARY KEY (table_id, column_id);


--
-- Name: sl_tables sl_tables_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sl_tables
    ADD CONSTRAINT sl_tables_pkey PRIMARY KEY (id);


--
-- Name: sl_tables sl_tables_uuid_key; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sl_tables
    ADD CONSTRAINT sl_tables_uuid_key UNIQUE (uuid);


--
-- Name: slice_user slice_user_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.slice_user
    ADD CONSTRAINT slice_user_pkey PRIMARY KEY (id);


--
-- Name: slices slices_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.slices
    ADD CONSTRAINT slices_pkey PRIMARY KEY (id);


--
-- Name: sql_metrics sql_metrics_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sql_metrics
    ADD CONSTRAINT sql_metrics_pkey PRIMARY KEY (id);


--
-- Name: sqlatable_user sqlatable_user_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sqlatable_user
    ADD CONSTRAINT sqlatable_user_pkey PRIMARY KEY (id);


--
-- Name: ssh_tunnels ssh_tunnels_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ssh_tunnels
    ADD CONSTRAINT ssh_tunnels_pkey PRIMARY KEY (id);


--
-- Name: tab_state tab_state_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tab_state
    ADD CONSTRAINT tab_state_pkey PRIMARY KEY (id);


--
-- Name: table_columns table_columns_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.table_columns
    ADD CONSTRAINT table_columns_pkey PRIMARY KEY (id);


--
-- Name: table_schema table_schema_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.table_schema
    ADD CONSTRAINT table_schema_pkey PRIMARY KEY (id);


--
-- Name: tables tables_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tables
    ADD CONSTRAINT tables_pkey PRIMARY KEY (id);


--
-- Name: tag tag_name_key; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tag
    ADD CONSTRAINT tag_name_key UNIQUE (name);


--
-- Name: tag tag_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tag
    ADD CONSTRAINT tag_pkey PRIMARY KEY (id);


--
-- Name: tagged_object tagged_object_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tagged_object
    ADD CONSTRAINT tagged_object_pkey PRIMARY KEY (id);


--
-- Name: dashboard_slices uq_dashboard_slice; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dashboard_slices
    ADD CONSTRAINT uq_dashboard_slice UNIQUE (dashboard_id, slice_id);


--
-- Name: dashboards uq_dashboards_uuid; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dashboards
    ADD CONSTRAINT uq_dashboards_uuid UNIQUE (uuid);


--
-- Name: dbs uq_dbs_uuid; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dbs
    ADD CONSTRAINT uq_dbs_uuid UNIQUE (uuid);


--
-- Name: report_schedule uq_report_schedule_name_type; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.report_schedule
    ADD CONSTRAINT uq_report_schedule_name_type UNIQUE (name, type);


--
-- Name: row_level_security_filters uq_rls_name; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.row_level_security_filters
    ADD CONSTRAINT uq_rls_name UNIQUE (name);


--
-- Name: saved_query uq_saved_query_uuid; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.saved_query
    ADD CONSTRAINT uq_saved_query_uuid UNIQUE (uuid);


--
-- Name: slices uq_slices_uuid; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.slices
    ADD CONSTRAINT uq_slices_uuid UNIQUE (uuid);


--
-- Name: sql_metrics uq_sql_metrics_metric_name; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sql_metrics
    ADD CONSTRAINT uq_sql_metrics_metric_name UNIQUE (metric_name, table_id);


--
-- Name: sql_metrics uq_sql_metrics_uuid; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sql_metrics
    ADD CONSTRAINT uq_sql_metrics_uuid UNIQUE (uuid);


--
-- Name: table_columns uq_table_columns_column_name; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.table_columns
    ADD CONSTRAINT uq_table_columns_column_name UNIQUE (column_name, table_id);


--
-- Name: table_columns uq_table_columns_uuid; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.table_columns
    ADD CONSTRAINT uq_table_columns_uuid UNIQUE (uuid);


--
-- Name: tables uq_tables_uuid; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tables
    ADD CONSTRAINT uq_tables_uuid UNIQUE (uuid);


--
-- Name: url url_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.url
    ADD CONSTRAINT url_pkey PRIMARY KEY (id);


--
-- Name: user_attribute user_attribute_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.user_attribute
    ADD CONSTRAINT user_attribute_pkey PRIMARY KEY (id);


--
-- Name: ix_cache_keys_datasource_uid; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX ix_cache_keys_datasource_uid ON public.cache_keys USING btree (datasource_uid);


--
-- Name: ix_creation_method; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX ix_creation_method ON public.report_schedule USING btree (creation_method);


--
-- Name: ix_key_value_expires_on; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX ix_key_value_expires_on ON public.key_value USING btree (expires_on);


--
-- Name: ix_key_value_uuid; Type: INDEX; Schema: public; Owner: admin
--

CREATE UNIQUE INDEX ix_key_value_uuid ON public.key_value USING btree (uuid);


--
-- Name: ix_logs_user_id_dttm; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX ix_logs_user_id_dttm ON public.logs USING btree (user_id, dttm);


--
-- Name: ix_query_results_key; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX ix_query_results_key ON public.query USING btree (results_key);


--
-- Name: ix_report_schedule_active; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX ix_report_schedule_active ON public.report_schedule USING btree (active);


--
-- Name: ix_row_level_security_filters_filter_type; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX ix_row_level_security_filters_filter_type ON public.row_level_security_filters USING btree (filter_type);


--
-- Name: ix_ssh_tunnels_database_id; Type: INDEX; Schema: public; Owner: admin
--

CREATE UNIQUE INDEX ix_ssh_tunnels_database_id ON public.ssh_tunnels USING btree (database_id);


--
-- Name: ix_ssh_tunnels_uuid; Type: INDEX; Schema: public; Owner: admin
--

CREATE UNIQUE INDEX ix_ssh_tunnels_uuid ON public.ssh_tunnels USING btree (uuid);


--
-- Name: ix_tab_state_id; Type: INDEX; Schema: public; Owner: admin
--

CREATE UNIQUE INDEX ix_tab_state_id ON public.tab_state USING btree (id);


--
-- Name: ix_table_schema_id; Type: INDEX; Schema: public; Owner: admin
--

CREATE UNIQUE INDEX ix_table_schema_id ON public.table_schema USING btree (id);


--
-- Name: ix_tagged_object_object_id; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX ix_tagged_object_object_id ON public.tagged_object USING btree (object_id);


--
-- Name: ti_dag_state; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX ti_dag_state ON public.annotation USING btree (layer_id, start_dttm, end_dttm);


--
-- Name: ti_user_id_changed_on; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX ti_user_id_changed_on ON public.query USING btree (user_id, changed_on);


--
-- Name: ab_permission_view ab_permission_view_permission_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ab_permission_view
    ADD CONSTRAINT ab_permission_view_permission_id_fkey FOREIGN KEY (permission_id) REFERENCES public.ab_permission(id);


--
-- Name: ab_permission_view_role ab_permission_view_role_permission_view_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ab_permission_view_role
    ADD CONSTRAINT ab_permission_view_role_permission_view_id_fkey FOREIGN KEY (permission_view_id) REFERENCES public.ab_permission_view(id);


--
-- Name: ab_permission_view_role ab_permission_view_role_role_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ab_permission_view_role
    ADD CONSTRAINT ab_permission_view_role_role_id_fkey FOREIGN KEY (role_id) REFERENCES public.ab_role(id);


--
-- Name: ab_permission_view ab_permission_view_view_menu_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ab_permission_view
    ADD CONSTRAINT ab_permission_view_view_menu_id_fkey FOREIGN KEY (view_menu_id) REFERENCES public.ab_view_menu(id);


--
-- Name: ab_user ab_user_changed_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ab_user
    ADD CONSTRAINT ab_user_changed_by_fk_fkey FOREIGN KEY (changed_by_fk) REFERENCES public.ab_user(id);


--
-- Name: ab_user ab_user_created_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ab_user
    ADD CONSTRAINT ab_user_created_by_fk_fkey FOREIGN KEY (created_by_fk) REFERENCES public.ab_user(id);


--
-- Name: ab_user_role ab_user_role_role_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ab_user_role
    ADD CONSTRAINT ab_user_role_role_id_fkey FOREIGN KEY (role_id) REFERENCES public.ab_role(id);


--
-- Name: ab_user_role ab_user_role_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ab_user_role
    ADD CONSTRAINT ab_user_role_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.ab_user(id);


--
-- Name: annotation annotation_changed_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.annotation
    ADD CONSTRAINT annotation_changed_by_fk_fkey FOREIGN KEY (changed_by_fk) REFERENCES public.ab_user(id);


--
-- Name: annotation annotation_created_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.annotation
    ADD CONSTRAINT annotation_created_by_fk_fkey FOREIGN KEY (created_by_fk) REFERENCES public.ab_user(id);


--
-- Name: annotation_layer annotation_layer_changed_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.annotation_layer
    ADD CONSTRAINT annotation_layer_changed_by_fk_fkey FOREIGN KEY (changed_by_fk) REFERENCES public.ab_user(id);


--
-- Name: annotation_layer annotation_layer_created_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.annotation_layer
    ADD CONSTRAINT annotation_layer_created_by_fk_fkey FOREIGN KEY (created_by_fk) REFERENCES public.ab_user(id);


--
-- Name: annotation annotation_layer_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.annotation
    ADD CONSTRAINT annotation_layer_id_fkey FOREIGN KEY (layer_id) REFERENCES public.annotation_layer(id);


--
-- Name: css_templates css_templates_changed_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.css_templates
    ADD CONSTRAINT css_templates_changed_by_fk_fkey FOREIGN KEY (changed_by_fk) REFERENCES public.ab_user(id);


--
-- Name: css_templates css_templates_created_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.css_templates
    ADD CONSTRAINT css_templates_created_by_fk_fkey FOREIGN KEY (created_by_fk) REFERENCES public.ab_user(id);


--
-- Name: dashboards dashboards_changed_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dashboards
    ADD CONSTRAINT dashboards_changed_by_fk_fkey FOREIGN KEY (changed_by_fk) REFERENCES public.ab_user(id);


--
-- Name: dashboards dashboards_created_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dashboards
    ADD CONSTRAINT dashboards_created_by_fk_fkey FOREIGN KEY (created_by_fk) REFERENCES public.ab_user(id);


--
-- Name: dbs dbs_changed_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dbs
    ADD CONSTRAINT dbs_changed_by_fk_fkey FOREIGN KEY (changed_by_fk) REFERENCES public.ab_user(id);


--
-- Name: dbs dbs_created_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dbs
    ADD CONSTRAINT dbs_created_by_fk_fkey FOREIGN KEY (created_by_fk) REFERENCES public.ab_user(id);


--
-- Name: dynamic_plugin dynamic_plugin_changed_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dynamic_plugin
    ADD CONSTRAINT dynamic_plugin_changed_by_fk_fkey FOREIGN KEY (changed_by_fk) REFERENCES public.ab_user(id);


--
-- Name: dynamic_plugin dynamic_plugin_created_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dynamic_plugin
    ADD CONSTRAINT dynamic_plugin_created_by_fk_fkey FOREIGN KEY (created_by_fk) REFERENCES public.ab_user(id);


--
-- Name: favstar favstar_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.favstar
    ADD CONSTRAINT favstar_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.ab_user(id);


--
-- Name: filter_sets filter_sets_changed_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.filter_sets
    ADD CONSTRAINT filter_sets_changed_by_fk_fkey FOREIGN KEY (changed_by_fk) REFERENCES public.ab_user(id);


--
-- Name: filter_sets filter_sets_created_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.filter_sets
    ADD CONSTRAINT filter_sets_created_by_fk_fkey FOREIGN KEY (created_by_fk) REFERENCES public.ab_user(id);


--
-- Name: filter_sets filter_sets_dashboard_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.filter_sets
    ADD CONSTRAINT filter_sets_dashboard_id_fkey FOREIGN KEY (dashboard_id) REFERENCES public.dashboards(id);


--
-- Name: dashboard_roles fk_dashboard_roles_dashboard_id_dashboards; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dashboard_roles
    ADD CONSTRAINT fk_dashboard_roles_dashboard_id_dashboards FOREIGN KEY (dashboard_id) REFERENCES public.dashboards(id) ON DELETE CASCADE;


--
-- Name: dashboard_roles fk_dashboard_roles_role_id_ab_role; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dashboard_roles
    ADD CONSTRAINT fk_dashboard_roles_role_id_ab_role FOREIGN KEY (role_id) REFERENCES public.ab_role(id) ON DELETE CASCADE;


--
-- Name: dashboard_slices fk_dashboard_slices_dashboard_id_dashboards; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dashboard_slices
    ADD CONSTRAINT fk_dashboard_slices_dashboard_id_dashboards FOREIGN KEY (dashboard_id) REFERENCES public.dashboards(id) ON DELETE CASCADE;


--
-- Name: dashboard_slices fk_dashboard_slices_slice_id_slices; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dashboard_slices
    ADD CONSTRAINT fk_dashboard_slices_slice_id_slices FOREIGN KEY (slice_id) REFERENCES public.slices(id) ON DELETE CASCADE;


--
-- Name: dashboard_user fk_dashboard_user_dashboard_id_dashboards; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dashboard_user
    ADD CONSTRAINT fk_dashboard_user_dashboard_id_dashboards FOREIGN KEY (dashboard_id) REFERENCES public.dashboards(id) ON DELETE CASCADE;


--
-- Name: dashboard_user fk_dashboard_user_user_id_ab_user; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.dashboard_user
    ADD CONSTRAINT fk_dashboard_user_user_id_ab_user FOREIGN KEY (user_id) REFERENCES public.ab_user(id) ON DELETE CASCADE;


--
-- Name: embedded_dashboards fk_embedded_dashboards_dashboard_id_dashboards; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.embedded_dashboards
    ADD CONSTRAINT fk_embedded_dashboards_dashboard_id_dashboards FOREIGN KEY (dashboard_id) REFERENCES public.dashboards(id) ON DELETE CASCADE;


--
-- Name: report_schedule_user fk_report_schedule_user_report_schedule_id_report_schedule; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.report_schedule_user
    ADD CONSTRAINT fk_report_schedule_user_report_schedule_id_report_schedule FOREIGN KEY (report_schedule_id) REFERENCES public.report_schedule(id) ON DELETE CASCADE;


--
-- Name: report_schedule_user fk_report_schedule_user_user_id_ab_user; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.report_schedule_user
    ADD CONSTRAINT fk_report_schedule_user_user_id_ab_user FOREIGN KEY (user_id) REFERENCES public.ab_user(id) ON DELETE CASCADE;


--
-- Name: slice_user fk_slice_user_slice_id_slices; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.slice_user
    ADD CONSTRAINT fk_slice_user_slice_id_slices FOREIGN KEY (slice_id) REFERENCES public.slices(id) ON DELETE CASCADE;


--
-- Name: slice_user fk_slice_user_user_id_ab_user; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.slice_user
    ADD CONSTRAINT fk_slice_user_user_id_ab_user FOREIGN KEY (user_id) REFERENCES public.ab_user(id) ON DELETE CASCADE;


--
-- Name: sql_metrics fk_sql_metrics_table_id_tables; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sql_metrics
    ADD CONSTRAINT fk_sql_metrics_table_id_tables FOREIGN KEY (table_id) REFERENCES public.tables(id) ON DELETE CASCADE;


--
-- Name: sqlatable_user fk_sqlatable_user_table_id_tables; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sqlatable_user
    ADD CONSTRAINT fk_sqlatable_user_table_id_tables FOREIGN KEY (table_id) REFERENCES public.tables(id) ON DELETE CASCADE;


--
-- Name: sqlatable_user fk_sqlatable_user_user_id_ab_user; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sqlatable_user
    ADD CONSTRAINT fk_sqlatable_user_user_id_ab_user FOREIGN KEY (user_id) REFERENCES public.ab_user(id) ON DELETE CASCADE;


--
-- Name: table_columns fk_table_columns_table_id_tables; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.table_columns
    ADD CONSTRAINT fk_table_columns_table_id_tables FOREIGN KEY (table_id) REFERENCES public.tables(id) ON DELETE CASCADE;


--
-- Name: key_value key_value_changed_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.key_value
    ADD CONSTRAINT key_value_changed_by_fk_fkey FOREIGN KEY (changed_by_fk) REFERENCES public.ab_user(id);


--
-- Name: key_value key_value_created_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.key_value
    ADD CONSTRAINT key_value_created_by_fk_fkey FOREIGN KEY (created_by_fk) REFERENCES public.ab_user(id);


--
-- Name: logs logs_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.logs
    ADD CONSTRAINT logs_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.ab_user(id);


--
-- Name: query query_database_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.query
    ADD CONSTRAINT query_database_id_fkey FOREIGN KEY (database_id) REFERENCES public.dbs(id);


--
-- Name: query query_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.query
    ADD CONSTRAINT query_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.ab_user(id);


--
-- Name: report_execution_log report_execution_log_report_schedule_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.report_execution_log
    ADD CONSTRAINT report_execution_log_report_schedule_id_fkey FOREIGN KEY (report_schedule_id) REFERENCES public.report_schedule(id);


--
-- Name: report_recipient report_recipient_changed_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.report_recipient
    ADD CONSTRAINT report_recipient_changed_by_fk_fkey FOREIGN KEY (changed_by_fk) REFERENCES public.ab_user(id);


--
-- Name: report_recipient report_recipient_created_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.report_recipient
    ADD CONSTRAINT report_recipient_created_by_fk_fkey FOREIGN KEY (created_by_fk) REFERENCES public.ab_user(id);


--
-- Name: report_recipient report_recipient_report_schedule_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.report_recipient
    ADD CONSTRAINT report_recipient_report_schedule_id_fkey FOREIGN KEY (report_schedule_id) REFERENCES public.report_schedule(id);


--
-- Name: report_schedule report_schedule_changed_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.report_schedule
    ADD CONSTRAINT report_schedule_changed_by_fk_fkey FOREIGN KEY (changed_by_fk) REFERENCES public.ab_user(id);


--
-- Name: report_schedule report_schedule_chart_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.report_schedule
    ADD CONSTRAINT report_schedule_chart_id_fkey FOREIGN KEY (chart_id) REFERENCES public.slices(id);


--
-- Name: report_schedule report_schedule_created_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.report_schedule
    ADD CONSTRAINT report_schedule_created_by_fk_fkey FOREIGN KEY (created_by_fk) REFERENCES public.ab_user(id);


--
-- Name: report_schedule report_schedule_dashboard_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.report_schedule
    ADD CONSTRAINT report_schedule_dashboard_id_fkey FOREIGN KEY (dashboard_id) REFERENCES public.dashboards(id);


--
-- Name: report_schedule report_schedule_database_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.report_schedule
    ADD CONSTRAINT report_schedule_database_id_fkey FOREIGN KEY (database_id) REFERENCES public.dbs(id);


--
-- Name: rls_filter_roles rls_filter_roles_rls_filter_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.rls_filter_roles
    ADD CONSTRAINT rls_filter_roles_rls_filter_id_fkey FOREIGN KEY (rls_filter_id) REFERENCES public.row_level_security_filters(id);


--
-- Name: rls_filter_roles rls_filter_roles_role_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.rls_filter_roles
    ADD CONSTRAINT rls_filter_roles_role_id_fkey FOREIGN KEY (role_id) REFERENCES public.ab_role(id);


--
-- Name: rls_filter_tables rls_filter_tables_rls_filter_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.rls_filter_tables
    ADD CONSTRAINT rls_filter_tables_rls_filter_id_fkey FOREIGN KEY (rls_filter_id) REFERENCES public.row_level_security_filters(id);


--
-- Name: rls_filter_tables rls_filter_tables_table_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.rls_filter_tables
    ADD CONSTRAINT rls_filter_tables_table_id_fkey FOREIGN KEY (table_id) REFERENCES public.tables(id);


--
-- Name: row_level_security_filters row_level_security_filters_changed_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.row_level_security_filters
    ADD CONSTRAINT row_level_security_filters_changed_by_fk_fkey FOREIGN KEY (changed_by_fk) REFERENCES public.ab_user(id);


--
-- Name: row_level_security_filters row_level_security_filters_created_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.row_level_security_filters
    ADD CONSTRAINT row_level_security_filters_created_by_fk_fkey FOREIGN KEY (created_by_fk) REFERENCES public.ab_user(id);


--
-- Name: saved_query saved_query_changed_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.saved_query
    ADD CONSTRAINT saved_query_changed_by_fk_fkey FOREIGN KEY (changed_by_fk) REFERENCES public.ab_user(id);


--
-- Name: saved_query saved_query_created_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.saved_query
    ADD CONSTRAINT saved_query_created_by_fk_fkey FOREIGN KEY (created_by_fk) REFERENCES public.ab_user(id);


--
-- Name: saved_query saved_query_db_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.saved_query
    ADD CONSTRAINT saved_query_db_id_fkey FOREIGN KEY (db_id) REFERENCES public.dbs(id);


--
-- Name: tab_state saved_query_id; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tab_state
    ADD CONSTRAINT saved_query_id FOREIGN KEY (saved_query_id) REFERENCES public.saved_query(id) ON DELETE SET NULL;


--
-- Name: saved_query saved_query_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.saved_query
    ADD CONSTRAINT saved_query_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.ab_user(id);


--
-- Name: sl_columns sl_columns_changed_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sl_columns
    ADD CONSTRAINT sl_columns_changed_by_fk_fkey FOREIGN KEY (changed_by_fk) REFERENCES public.ab_user(id);


--
-- Name: sl_columns sl_columns_created_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sl_columns
    ADD CONSTRAINT sl_columns_created_by_fk_fkey FOREIGN KEY (created_by_fk) REFERENCES public.ab_user(id);


--
-- Name: sl_dataset_columns sl_dataset_columns_column_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sl_dataset_columns
    ADD CONSTRAINT sl_dataset_columns_column_id_fkey FOREIGN KEY (column_id) REFERENCES public.sl_columns(id);


--
-- Name: sl_dataset_columns sl_dataset_columns_dataset_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sl_dataset_columns
    ADD CONSTRAINT sl_dataset_columns_dataset_id_fkey FOREIGN KEY (dataset_id) REFERENCES public.sl_datasets(id);


--
-- Name: sl_dataset_tables sl_dataset_tables_dataset_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sl_dataset_tables
    ADD CONSTRAINT sl_dataset_tables_dataset_id_fkey FOREIGN KEY (dataset_id) REFERENCES public.sl_datasets(id);


--
-- Name: sl_dataset_tables sl_dataset_tables_table_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sl_dataset_tables
    ADD CONSTRAINT sl_dataset_tables_table_id_fkey FOREIGN KEY (table_id) REFERENCES public.sl_tables(id);


--
-- Name: sl_dataset_users sl_dataset_users_dataset_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sl_dataset_users
    ADD CONSTRAINT sl_dataset_users_dataset_id_fkey FOREIGN KEY (dataset_id) REFERENCES public.sl_datasets(id);


--
-- Name: sl_dataset_users sl_dataset_users_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sl_dataset_users
    ADD CONSTRAINT sl_dataset_users_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.ab_user(id);


--
-- Name: sl_datasets sl_datasets_changed_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sl_datasets
    ADD CONSTRAINT sl_datasets_changed_by_fk_fkey FOREIGN KEY (changed_by_fk) REFERENCES public.ab_user(id);


--
-- Name: sl_datasets sl_datasets_created_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sl_datasets
    ADD CONSTRAINT sl_datasets_created_by_fk_fkey FOREIGN KEY (created_by_fk) REFERENCES public.ab_user(id);


--
-- Name: sl_datasets sl_datasets_database_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sl_datasets
    ADD CONSTRAINT sl_datasets_database_id_fkey FOREIGN KEY (database_id) REFERENCES public.dbs(id);


--
-- Name: sl_table_columns sl_table_columns_column_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sl_table_columns
    ADD CONSTRAINT sl_table_columns_column_id_fkey FOREIGN KEY (column_id) REFERENCES public.sl_columns(id);


--
-- Name: sl_table_columns sl_table_columns_table_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sl_table_columns
    ADD CONSTRAINT sl_table_columns_table_id_fkey FOREIGN KEY (table_id) REFERENCES public.sl_tables(id);


--
-- Name: sl_tables sl_tables_changed_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sl_tables
    ADD CONSTRAINT sl_tables_changed_by_fk_fkey FOREIGN KEY (changed_by_fk) REFERENCES public.ab_user(id);


--
-- Name: sl_tables sl_tables_created_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sl_tables
    ADD CONSTRAINT sl_tables_created_by_fk_fkey FOREIGN KEY (created_by_fk) REFERENCES public.ab_user(id);


--
-- Name: sl_tables sl_tables_database_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sl_tables
    ADD CONSTRAINT sl_tables_database_id_fkey FOREIGN KEY (database_id) REFERENCES public.dbs(id);


--
-- Name: slices slices_changed_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.slices
    ADD CONSTRAINT slices_changed_by_fk_fkey FOREIGN KEY (changed_by_fk) REFERENCES public.ab_user(id);


--
-- Name: slices slices_created_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.slices
    ADD CONSTRAINT slices_created_by_fk_fkey FOREIGN KEY (created_by_fk) REFERENCES public.ab_user(id);


--
-- Name: slices slices_last_saved_by_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.slices
    ADD CONSTRAINT slices_last_saved_by_fk FOREIGN KEY (last_saved_by_fk) REFERENCES public.ab_user(id);


--
-- Name: sql_metrics sql_metrics_changed_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sql_metrics
    ADD CONSTRAINT sql_metrics_changed_by_fk_fkey FOREIGN KEY (changed_by_fk) REFERENCES public.ab_user(id);


--
-- Name: sql_metrics sql_metrics_created_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sql_metrics
    ADD CONSTRAINT sql_metrics_created_by_fk_fkey FOREIGN KEY (created_by_fk) REFERENCES public.ab_user(id);


--
-- Name: ssh_tunnels ssh_tunnels_database_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.ssh_tunnels
    ADD CONSTRAINT ssh_tunnels_database_id_fkey FOREIGN KEY (database_id) REFERENCES public.dbs(id);


--
-- Name: tab_state tab_state_changed_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tab_state
    ADD CONSTRAINT tab_state_changed_by_fk_fkey FOREIGN KEY (changed_by_fk) REFERENCES public.ab_user(id);


--
-- Name: tab_state tab_state_created_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tab_state
    ADD CONSTRAINT tab_state_created_by_fk_fkey FOREIGN KEY (created_by_fk) REFERENCES public.ab_user(id);


--
-- Name: tab_state tab_state_database_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tab_state
    ADD CONSTRAINT tab_state_database_id_fkey FOREIGN KEY (database_id) REFERENCES public.dbs(id) ON DELETE CASCADE;


--
-- Name: tab_state tab_state_latest_query_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tab_state
    ADD CONSTRAINT tab_state_latest_query_id_fkey FOREIGN KEY (latest_query_id) REFERENCES public.query(client_id) ON DELETE SET NULL;


--
-- Name: tab_state tab_state_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tab_state
    ADD CONSTRAINT tab_state_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.ab_user(id);


--
-- Name: table_columns table_columns_changed_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.table_columns
    ADD CONSTRAINT table_columns_changed_by_fk_fkey FOREIGN KEY (changed_by_fk) REFERENCES public.ab_user(id);


--
-- Name: table_columns table_columns_created_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.table_columns
    ADD CONSTRAINT table_columns_created_by_fk_fkey FOREIGN KEY (created_by_fk) REFERENCES public.ab_user(id);


--
-- Name: table_schema table_schema_changed_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.table_schema
    ADD CONSTRAINT table_schema_changed_by_fk_fkey FOREIGN KEY (changed_by_fk) REFERENCES public.ab_user(id);


--
-- Name: table_schema table_schema_created_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.table_schema
    ADD CONSTRAINT table_schema_created_by_fk_fkey FOREIGN KEY (created_by_fk) REFERENCES public.ab_user(id);


--
-- Name: table_schema table_schema_database_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.table_schema
    ADD CONSTRAINT table_schema_database_id_fkey FOREIGN KEY (database_id) REFERENCES public.dbs(id) ON DELETE CASCADE;


--
-- Name: table_schema table_schema_tab_state_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.table_schema
    ADD CONSTRAINT table_schema_tab_state_id_fkey FOREIGN KEY (tab_state_id) REFERENCES public.tab_state(id) ON DELETE CASCADE;


--
-- Name: tables tables_changed_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tables
    ADD CONSTRAINT tables_changed_by_fk_fkey FOREIGN KEY (changed_by_fk) REFERENCES public.ab_user(id);


--
-- Name: tables tables_created_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tables
    ADD CONSTRAINT tables_created_by_fk_fkey FOREIGN KEY (created_by_fk) REFERENCES public.ab_user(id);


--
-- Name: tables tables_database_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tables
    ADD CONSTRAINT tables_database_id_fkey FOREIGN KEY (database_id) REFERENCES public.dbs(id);


--
-- Name: tag tag_changed_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tag
    ADD CONSTRAINT tag_changed_by_fk_fkey FOREIGN KEY (changed_by_fk) REFERENCES public.ab_user(id);


--
-- Name: tag tag_created_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tag
    ADD CONSTRAINT tag_created_by_fk_fkey FOREIGN KEY (created_by_fk) REFERENCES public.ab_user(id);


--
-- Name: tagged_object tagged_object_changed_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tagged_object
    ADD CONSTRAINT tagged_object_changed_by_fk_fkey FOREIGN KEY (changed_by_fk) REFERENCES public.ab_user(id);


--
-- Name: tagged_object tagged_object_created_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tagged_object
    ADD CONSTRAINT tagged_object_created_by_fk_fkey FOREIGN KEY (created_by_fk) REFERENCES public.ab_user(id);


--
-- Name: tagged_object tagged_object_tag_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tagged_object
    ADD CONSTRAINT tagged_object_tag_id_fkey FOREIGN KEY (tag_id) REFERENCES public.tag(id);


--
-- Name: url url_changed_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.url
    ADD CONSTRAINT url_changed_by_fk_fkey FOREIGN KEY (changed_by_fk) REFERENCES public.ab_user(id);


--
-- Name: url url_created_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.url
    ADD CONSTRAINT url_created_by_fk_fkey FOREIGN KEY (created_by_fk) REFERENCES public.ab_user(id);


--
-- Name: user_attribute user_attribute_changed_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.user_attribute
    ADD CONSTRAINT user_attribute_changed_by_fk_fkey FOREIGN KEY (changed_by_fk) REFERENCES public.ab_user(id);


--
-- Name: user_attribute user_attribute_created_by_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.user_attribute
    ADD CONSTRAINT user_attribute_created_by_fk_fkey FOREIGN KEY (created_by_fk) REFERENCES public.ab_user(id);


--
-- Name: user_attribute user_attribute_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.user_attribute
    ADD CONSTRAINT user_attribute_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.ab_user(id);


--
-- Name: user_attribute user_attribute_welcome_dashboard_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.user_attribute
    ADD CONSTRAINT user_attribute_welcome_dashboard_id_fkey FOREIGN KEY (welcome_dashboard_id) REFERENCES public.dashboards(id);


--
-- Name: user_favorite_tag user_favorite_tag_tag_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.user_favorite_tag
    ADD CONSTRAINT user_favorite_tag_tag_id_fkey FOREIGN KEY (tag_id) REFERENCES public.tag(id);


--
-- Name: user_favorite_tag user_favorite_tag_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.user_favorite_tag
    ADD CONSTRAINT user_favorite_tag_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.ab_user(id);


--
-- PostgreSQL database dump complete
--

