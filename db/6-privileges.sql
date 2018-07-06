DROP ROLE IF EXISTS operation;
DROP ROLE IF EXISTS tutor;
CREATE ROLE operation NOLOGIN;
CREATE ROLE tutor NOLOGIN;


GRANT ALL PRIVILEGES
ON ALL SEQUENCES IN SCHEMA PUBLIC
TO operation
, tutor;

GRANT ALL PRIVILEGES
ON ALL TABLES IN SCHEMA PUBLIC
TO operation;

GRANT SELECT
ON ALL TABLES IN SCHEMA PUBLIC
TO tutor;

GRANT INSERT
, DELETE
ON TABLE sessions
, sessions_tutors
, sessions_tutors_students
TO tutor;
