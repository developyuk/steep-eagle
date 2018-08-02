DROP TABLE IF EXISTS branches CASCADE;
DROP TABLE IF EXISTS classes CASCADE;
DROP TABLE IF EXISTS class_students CASCADE;
DROP TABLE IF EXISTS modules CASCADE;
DROP TABLE IF EXISTS programs CASCADE;
DROP TABLE IF EXISTS programs_modules CASCADE;
DROP TABLE IF EXISTS program_types CASCADE;
DROP TABLE IF EXISTS sessions CASCADE;
DROP TABLE IF EXISTS sessions_tutors CASCADE;
DROP TABLE IF EXISTS sessions_tutors_students CASCADE;
DROP TABLE IF EXISTS users CASCADE;


DROP TYPE IF EXISTS DAYS CASCADE;
DROP TYPE IF EXISTS ROLES CASCADE;

CREATE TABLE program_types (
  id       BIGSERIAL PRIMARY KEY,
  name     TEXT NOT NULL UNIQUE,
  _created TIMESTAMP DEFAULT NOW(),
  _updated TIMESTAMP DEFAULT NOW(),
  _etag    TEXT
);

CREATE TABLE programs (
  id       BIGSERIAL PRIMARY KEY,
  name     TEXT NOT NULL,
  type_id  INT REFERENCES program_types (id),
  _created TIMESTAMP DEFAULT NOW(),
  _updated TIMESTAMP DEFAULT NOW(),
  _etag    TEXT,
  UNIQUE (
    name,
    type_id
  )
);

CREATE TABLE modules (
  id       BIGSERIAL PRIMARY KEY,
  name     TEXT NOT NULL UNIQUE,
  image    TEXT,
  _created TIMESTAMP DEFAULT NOW(),
  _updated TIMESTAMP DEFAULT NOW(),
  _etag    TEXT
);

CREATE TABLE programs_modules (
  id         BIGSERIAL PRIMARY KEY,
  module_id  INT REFERENCES modules (id),
  program_id INT REFERENCES programs (id),
  UNIQUE (
    module_id,
    program_id
  )
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
  user_id INT REFERENCES users (id),
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
  day               DAYS    NOT NULL,
  start_at          CHAR(5) NOT NULL,
  finish_at         CHAR(5) NOT NULL,
  program_module_id INT REFERENCES programs_modules (id),
  branch_id         INT REFERENCES branches (id),
  tutor_id          INT REFERENCES users (id),
  _created          TIMESTAMP DEFAULT NOW(),
  _updated          TIMESTAMP DEFAULT NOW(),
  _etag    TEXT,
  UNIQUE (
    day,
    start_at,
    finish_at,
    program_module_id,
    branch_id
  )
);

CREATE TABLE class_students (
  id         BIGSERIAL PRIMARY KEY,
  class_id   INT REFERENCES classes (id),
  student_id INT REFERENCES users (id),
  UNIQUE (
    class_id,
    student_id
  )
);

CREATE OR REPLACE FUNCTION user_by_email_pass(
  _email TEXT,
  _pass  TEXT
)
  RETURNS TABLE(
    id    BIGINT,
    email TEXT,
    ROLE  TEXT,
    name  TEXT,
    photo TEXT
  ) AS $$

BEGIN
  RETURN QUERY

  SELECT
    users.id,
    users.email,
    users.ROLE :: TEXT,
    users_profile.name,
    users_profile.photo
  FROM users
    LEFT JOIN users_profile ON users.id = users_profile.user_id
  WHERE users.email = user_by_email_pass._email
        AND users.pass = crypt(user_by_email_pass._pass, users.pass);
END;$$

LANGUAGE 'plpgsql';

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

CREATE TABLE sessions_tutors_students (
  id                 BIGSERIAL PRIMARY KEY,
  status             BOOLEAN ,
  feedback           TEXT    ,
  rating_interaction SMALLINT,
  rating_cognition   SMALLINT,
  rating_creativity  SMALLINT,
  session_tutor_id   INT REFERENCES sessions_tutors (id),
  student_id         INT REFERENCES users (id),
  _created           TIMESTAMP DEFAULT NOW(),
  _updated           TIMESTAMP DEFAULT NOW(),
  _etag    TEXT
);

SELECT 'DROP TABLE IF EXISTS ' || table_name || ' CASCADE;'
FROM information_schema.tables
WHERE table_schema = 'public' AND table_type = 'BASE TABLE'
ORDER BY table_name;
