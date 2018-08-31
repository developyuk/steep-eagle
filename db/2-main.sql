
DROP SCHEMA public CASCADE;
CREATE SCHEMA public;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO public;

DROP TYPE IF EXISTS DAYS CASCADE;
DROP TYPE IF EXISTS ROLES CASCADE;

CREATE TABLE modules (
  id       BIGSERIAL PRIMARY KEY,
  name     TEXT NOT NULL UNIQUE,
  image    TEXT,
  _created TIMESTAMP DEFAULT NOW(),
  _updated TIMESTAMP DEFAULT NOW(),
  _etag    TEXT
);


CREATE TABLE branches (
  id       BIGSERIAL PRIMARY KEY,
  name     TEXT NOT NULL UNIQUE,
  address  TEXT,
  _created TIMESTAMP DEFAULT NOW(),
  _updated TIMESTAMP DEFAULT NOW(),
  _etag    TEXT
);

CREATE TYPE DAYS AS ENUM (
  'monday',
  'tuesday',
  'wednesday',
  'thursday',
  'friday',
  'saturday',
  'sunday'
);

CREATE TYPE ROLES AS ENUM (
  'operation',
  'tutor',
  'student',
  'parent',
  'partner'
);

CREATE TABLE users (
  id       BIGSERIAL PRIMARY KEY,
  username TEXT  NOT NULL,
  email    TEXT,
  pass     TEXT,
  role     ROLES NOT NULL,
  name    TEXT,
  dob     TEXT,
  photo   TEXT,
  _created TIMESTAMP DEFAULT NOW(),
  _updated TIMESTAMP DEFAULT NOW(),
  _etag    TEXT,
  UNIQUE (
    username
    , email
  )
);

CREATE TABLE classes (
  id                BIGSERIAL PRIMARY KEY,
  module_id INT REFERENCES modules (id),
  branch_id         INT REFERENCES branches (id),
  tutor_id          INT REFERENCES users (id),
  day               DAYS    NOT NULL,
  start_at          CHAR(5) NOT NULL,
  finish_at         CHAR(5) NOT NULL,
  _created          TIMESTAMP DEFAULT NOW(),
  _updated          TIMESTAMP DEFAULT NOW(),
  _etag    TEXT,
  UNIQUE (
    day,
    start_at,
    finish_at,
    module_id,
    branch_id
  )
);

CREATE TABLE class_students (
  id         BIGSERIAL PRIMARY KEY,
  class_id   INT REFERENCES classes (id),
  student_id INT REFERENCES users (id),
  _created          TIMESTAMP DEFAULT NOW(),
  _updated          TIMESTAMP DEFAULT NOW(),
  _etag    TEXT,
  UNIQUE (
    class_id,
    student_id
  )
);

CREATE TABLE sessions (
  id       BIGSERIAL PRIMARY KEY,
  class_id INT REFERENCES classes (id),
  _created TIMESTAMP DEFAULT NOW(),
  _updated TIMESTAMP DEFAULT NOW(),
  _etag    TEXT
);

CREATE TABLE sessions_tutors (
  id         BIGSERIAL PRIMARY KEY,
  session_id INT REFERENCES sessions (id),
  tutor_id   INT REFERENCES users (id),
  _created   TIMESTAMP DEFAULT NOW(),
  _updated   TIMESTAMP DEFAULT NOW(),
  _etag    TEXT
);

CREATE TABLE sessions_students (
  id                 BIGSERIAL PRIMARY KEY,
  session_id INT REFERENCES sessions (id),
  tutor_id           INT REFERENCES users (id),
  student_id         INT REFERENCES users (id),
  status             BOOLEAN ,
  feedback           TEXT    ,
  rating_interaction SMALLINT,
  rating_cognition   SMALLINT,
  rating_creativity  SMALLINT,
  _created           TIMESTAMP DEFAULT NOW(),
  _updated           TIMESTAMP DEFAULT NOW(),
  _etag    TEXT
);
