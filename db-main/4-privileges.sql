CREATE ROLE anon nologin;

CREATE ROLE operation nologin;

CREATE ROLE tutor nologin;

GRANT ALL PRIVILEGES
	ON sequence users_id_seq
	TO anon;

GRANT ALL PRIVILEGES
	ON FUNCTION user_by_email_pass
	TO anon;

GRANT SELECT
	ON _users_profile
	TO anon;

GRANT ALL PRIVILEGES
	ON ALL sequences IN SCHEMA PUBLIC
	TO operation
		,tutor;

GRANT ALL PRIVILEGES
	ON ALL tables IN SCHEMA PUBLIC
	TO operation;

GRANT SELECT
	ON ALL tables IN SCHEMA PUBLIC
	TO tutor;

GRANT INSERT
	ON TABLE sessions
		,sessions_tutors
		,sessions_students
	TO tutor;
