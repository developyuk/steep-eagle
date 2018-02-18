CREATE ROLE operation WITH LOGIN PASSWORD 'mysecretpassword';
CREATE ROLE tutor WITH LOGIN PASSWORD 'mysecretpassword';
CREATE ROLE student WITH LOGIN PASSWORD 'mysecretpassword';
CREATE ROLE parent WITH LOGIN PASSWORD 'mysecretpassword';
CREATE ROLE partner WITH LOGIN PASSWORD 'mysecretpassword';

REVOKE CONNECT ON DATABASE postgres FROM PUBLIC;
GRANT CONNECT ON DATABASE postgres TO operation, tutor, student, parent,  partner;

REVOKE ALL ON ALL TABLES IN SCHEMA public FROM PUBLIC;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public to operation, tutor;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO operation;
GRANT INSERT ON TABLE sessions,sessions_students TO tutor;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO tutor, student, parent, partner;
