CREATE ROLE anon NOLOGIN;
CREATE ROLE operation NOLOGIN;
CREATE ROLE tutor NOLOGIN;
CREATE ROLE student NOLOGIN;
CREATE ROLE parent NOLOGIN;
CREATE ROLE partner NOLOGIN;

GRANT anon TO postgres;

GRANT USAGE
	ON SCHEMA public
	TO anon;


REVOKE ALL
	ON ALL TABLES IN SCHEMA public
	FROM PUBLIC;

GRANT USAGE
	,SELECT
	ON ALL SEQUENCES IN SCHEMA public
	TO operation
		,tutor;

GRANT SELECT
	,INSERT
	,UPDATE
	,DELETE
	ON ALL TABLES IN SCHEMA public
	TO operation;

GRANT INSERT
	ON TABLE sessions
		,sessions_students
	TO tutor;

GRANT SELECT
	ON ALL TABLES IN SCHEMA public
	TO tutor
		,student
		,parent
		,partner;
