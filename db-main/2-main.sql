CREATE TABLE program_types (
  id   BIGSERIAL PRIMARY KEY,
  name TEXT NOT NULL UNIQUE
);

CREATE TABLE programs (
  id      BIGSERIAL PRIMARY KEY,
  name    TEXT NOT NULL,
  type_id INT REFERENCES program_types (id),
  UNIQUE (
    name,
    type_id
  )
);

CREATE TABLE modules (
  id    BIGSERIAL PRIMARY KEY,
  name  TEXT NOT NULL UNIQUE,
  image TEXT NULL
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
  id      BIGSERIAL PRIMARY KEY,
  name    TEXT NOT NULL UNIQUE,
  address TEXT NULL
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
  email    TEXT  NULL,
  pass     TEXT  NULL,
  ROLE     ROLES NOT NULL,
  UNIQUE (
    username
    , email
  )
);

CREATE TABLE users_profile (
  name    TEXT NULL,
  dob     TEXT NULL,
  photo   TEXT NULL,
  user_id INT REFERENCES users (id),
  UNIQUE (
    name
    , user_id
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
  UNIQUE (
    day,
    start_at,
    finish_at,
    program_module_id,
    branch_id
  )
);

CREATE TABLE class_students (
  class_id   INT REFERENCES classes (id),
  student_id INT REFERENCES users (id),
  UNIQUE (
    class_id,
    student_id
  )
);
