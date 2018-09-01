DROP SCHEMA PUBLIC cascade;

CREATE SCHEMA PUBLIC;

GRANT ALL ON SCHEMA PUBLIC TO postgres;

GRANT ALL ON SCHEMA PUBLIC TO PUBLIC;

DROP TYPE IF EXISTS DAYS cascade;

DROP TYPE IF EXISTS roles cascade;

CREATE TABLE modules
  (
     id       BIGSERIAL PRIMARY KEY,
     name     TEXT NOT NULL UNIQUE,
     image    TEXT,
     _created TIMESTAMP DEFAULT Now(),
     _updated TIMESTAMP DEFAULT Now(),
     _etag    TEXT
  );

CREATE TABLE branches
  (
     id       BIGSERIAL PRIMARY KEY,
     name     TEXT NOT NULL UNIQUE,
     address  TEXT,
     _created TIMESTAMP DEFAULT Now(),
     _updated TIMESTAMP DEFAULT Now(),
     _etag    TEXT
  );

CREATE TYPE DAYS AS enum ( 'monday', 'tuesday', 'wednesday', 'thursday',
'friday', 'saturday', 'sunday' );

CREATE TYPE roles AS enum ( 'operation', 'tutor', 'student', 'parent', 'partner'
);

CREATE TABLE users
  (
     id       BIGSERIAL PRIMARY KEY,
     username TEXT NOT NULL,
     email    TEXT,
     pass     TEXT,
     ROLE     ROLES NOT NULL,
     name     TEXT,
     dob      TEXT,
     photo    TEXT,
     _created TIMESTAMP DEFAULT Now(),
     _updated TIMESTAMP DEFAULT Now(),
     _etag    TEXT,
     UNIQUE ( username, email )
  );

CREATE TABLE classes
  (
     id        BIGSERIAL PRIMARY KEY,
     module_id INT REFERENCES modules (id),
     branch_id INT REFERENCES branches (id),
     tutor_id  INT REFERENCES users (id),
     day       DAYS NOT NULL,
     start_at  CHAR(5) NOT NULL,
     finish_at CHAR(5) NOT NULL,
     _created  TIMESTAMP DEFAULT Now(),
     _updated  TIMESTAMP DEFAULT Now(),
     _etag     TEXT,
     UNIQUE ( day, start_at, finish_at, module_id, branch_id )
  );

CREATE TABLE class_students
  (
     id         BIGSERIAL PRIMARY KEY,
     class_id   INT REFERENCES classes (id),
     student_id INT REFERENCES users (id),
     _created   TIMESTAMP DEFAULT Now(),
     _updated   TIMESTAMP DEFAULT Now(),
     _etag      TEXT,
     UNIQUE ( class_id, student_id )
  );

CREATE TABLE sessions
  (
     id       BIGSERIAL PRIMARY KEY,
     class_id INT REFERENCES classes (id),
     _created TIMESTAMP DEFAULT Now(),
     _updated TIMESTAMP DEFAULT Now(),
     _etag    TEXT
  );

CREATE TABLE sessions_tutors
  (
     id         BIGSERIAL PRIMARY KEY,
     session_id INT REFERENCES sessions (id),
     tutor_id   INT REFERENCES users (id),
     _created   TIMESTAMP DEFAULT Now(),
     _updated   TIMESTAMP DEFAULT Now(),
     _etag      TEXT
  );

CREATE TABLE sessions_students
  (
     id                 BIGSERIAL PRIMARY KEY,
     session_id         INT REFERENCES sessions (id),
     tutor_id           INT REFERENCES users (id),
     student_id         INT REFERENCES users (id),
     status             BOOLEAN,
     feedback           TEXT,
     rating_interaction SMALLINT,
     rating_cognition   SMALLINT,
     rating_creativity  SMALLINT,
     _created           TIMESTAMP DEFAULT Now(),
     _updated           TIMESTAMP DEFAULT Now(),
     _etag              TEXT
  );
