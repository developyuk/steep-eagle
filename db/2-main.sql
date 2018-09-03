DROP
  SCHEMA PUBLIC cascade;
CREATE SCHEMA PUBLIC;
GRANT ALL ON SCHEMA PUBLIC TO postgres;
GRANT ALL ON SCHEMA PUBLIC TO PUBLIC;
CREATE TYPE days_enum AS enum (
  'monday', 'tuesday', 'wednesday',
  'thursday', 'friday', 'saturday',
  'sunday'
);
CREATE TYPE roles_enum AS enum (
  'operation', 'tutor', 'student', 'guardian',
  'partner'
);
CREATE TABLE modules (
  id BIGSERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  image TEXT,
  _created TIMESTAMP DEFAULT Now(),
  _updated TIMESTAMP DEFAULT Now(),
  _etag varchar(40),
  UNIQUE (name)
);
CREATE TABLE branches (
  id BIGSERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  address TEXT,
  _created TIMESTAMP DEFAULT Now(),
  _updated TIMESTAMP DEFAULT Now(),
  _etag varchar(40),
  UNIQUE (name)
);
CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  username TEXT NOT NULL,
  email TEXT,
  contact_no TEXT,
  pass TEXT,
  role roles_enum NOT NULL,

  name TEXT,
  dob TEXT,
  address TEXT,
  school TEXT,
  grade TEXT,
  photo TEXT,
  _created TIMESTAMP DEFAULT Now(),
  _updated TIMESTAMP DEFAULT Now(),
  _deleted BOOLEAN DEFAULT FALSE,
  _etag varchar(40),
  UNIQUE (username, email, contact_no, name)
);
CREATE INDEX users_auth_idx ON users (username);
CREATE TABLE student_guardians (
  id BIGSERIAL PRIMARY KEY,
  student_id INT REFERENCES users (id),
  guardian_id INT REFERENCES users (id),
  _created TIMESTAMP DEFAULT Now(),
  _updated TIMESTAMP DEFAULT Now(),
  _etag varchar(40)
);
CREATE TABLE classes (
  id BIGSERIAL PRIMARY KEY,
  module_id INT REFERENCES modules (id),
  branch_id INT REFERENCES branches (id),
  tutor_id INT REFERENCES users (id),
  day days_enum NOT NULL,
  start_at CHAR(5) NOT NULL,
  finish_at CHAR(5) NOT NULL,
  _created TIMESTAMP DEFAULT Now(),
  _updated TIMESTAMP DEFAULT Now(),
  _etag varchar(40),
  UNIQUE (
    day, start_at, finish_at, module_id,
    branch_id
  )
);
CREATE INDEX classes_tutor_stats_idx ON classes (tutor_id);
CREATE TABLE class_students (
  id BIGSERIAL PRIMARY KEY,
  class_id INT REFERENCES classes (id),
  student_id INT REFERENCES users (id),
  _created TIMESTAMP DEFAULT Now(),
  _updated TIMESTAMP DEFAULT Now(),
  _etag varchar(40),
  UNIQUE (class_id, student_id)
);
CREATE INDEX class_students_class_idx ON class_students (class_id);
CREATE TABLE attendances (
  id BIGSERIAL PRIMARY KEY,
  class_id INT REFERENCES classes (id),
  module_id INT REFERENCES modules (id),
  _created TIMESTAMP DEFAULT Now(),
  _updated TIMESTAMP DEFAULT Now(),
  _etag varchar(40)
);
CREATE INDEX attendances_schedules_idx ON attendances (class_id, _created);
CREATE TABLE attendances_tutors (
  id BIGSERIAL PRIMARY KEY,
  attendance_id INT REFERENCES attendances (id),
  tutor_id INT REFERENCES users (id),
  _created TIMESTAMP DEFAULT Now(),
  _updated TIMESTAMP DEFAULT Now(),
  _etag varchar(40)
);
CREATE INDEX attendances_tutors_students_idx ON attendances_tutors (tutor_id, _created);
CREATE INDEX attendances_tutors_tutor_stats_idx ON attendances_tutors (tutor_id);
CREATE TABLE attendances_students (
  id BIGSERIAL PRIMARY KEY,
  attendance_id INT REFERENCES attendances (id),
  tutor_id INT REFERENCES users (id),
  student_id INT REFERENCES users (id),
  is_presence BOOLEAN,
  feedback TEXT,
  rating_interaction SMALLINT,
  rating_cognition SMALLINT,
  rating_creativity SMALLINT,
  _created TIMESTAMP DEFAULT Now(),
  _updated TIMESTAMP DEFAULT Now(),
  _etag varchar(40)
);
CREATE INDEX attendances_students_students_idx ON attendances_students (
  attendance_id, tutor_id, student_id,
  _created
);
CREATE INDEX attendances_students_tutors_tutor_stats_idx0 ON attendances_students (tutor_id, feedback);
CREATE INDEX attendances_students_tutors_tutor_stats_idx1 ON attendances_students (
  tutor_id, rating_interaction, rating_cognition,
  rating_creativity
);
CREATE TABLE notifications (
  id BIGSERIAL PRIMARY KEY,
  message TEXT,
  is_read BOOLEAN,
  data TEXT,
  _created TIMESTAMP DEFAULT Now(),
  _updated TIMESTAMP DEFAULT Now(),
  _etag varchar(40)
);
