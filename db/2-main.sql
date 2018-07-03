-- truncate all table
--SELECT 'DROP TABLE IF EXISTS ' || table_name || ' CASCADE;' FROM information_schema.tables WHERE table_schema='public' AND table_type='BASE TABLE';
DROP TABLE IF EXISTS sessions_tutors_students CASCADE;
DROP TABLE IF EXISTS sessions_tutors CASCADE;
DROP TABLE IF EXISTS sessions CASCADE;
DROP TABLE IF EXISTS class_students CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS classes CASCADE;
DROP TABLE IF EXISTS users_profile CASCADE;
DROP TABLE IF EXISTS branches CASCADE;
DROP TABLE IF EXISTS programs_modules CASCADE;
DROP TABLE IF EXISTS programs CASCADE;
DROP TABLE IF EXISTS modules CASCADE;
DROP TABLE IF EXISTS program_types CASCADE;


DROP TYPE IF EXISTS DAYS CASCADE;
DROP TYPE IF EXISTS ROLES CASCADE;

CREATE TABLE program_types (
  id       BIGSERIAL PRIMARY KEY,
  name     TEXT NOT NULL UNIQUE,
  _created TIMESTAMPTZ DEFAULT NOW(),
  _updated TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE programs (
  id       BIGSERIAL PRIMARY KEY,
  name     TEXT NOT NULL,
  type_id  INT REFERENCES program_types (id),
  _created TIMESTAMPTZ DEFAULT NOW(),
  _updated TIMESTAMPTZ DEFAULT NOW(),
  UNIQUE (
    name,
    type_id
  )
);

CREATE TABLE modules (
  id       BIGSERIAL PRIMARY KEY,
  name     TEXT NOT NULL UNIQUE,
  image    TEXT NULL,
  _created TIMESTAMPTZ DEFAULT NOW(),
  _updated TIMESTAMPTZ DEFAULT NOW()
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
  address  TEXT NULL,
  _created TIMESTAMPTZ DEFAULT NOW(),
  _updated TIMESTAMPTZ DEFAULT NOW()
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
  _created TIMESTAMPTZ DEFAULT NOW(),
  _updated TIMESTAMPTZ DEFAULT NOW(),
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
  _created          TIMESTAMPTZ DEFAULT NOW(),
  _updated          TIMESTAMPTZ DEFAULT NOW(),
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

CREATE OR REPLACE VIEW _classes_ts AS
  SELECT d.*
  FROM (
         SELECT
           c.*,
           (a.dw :: DATE :: TEXT || ' ' || c.start_at || ':00') :: TIMESTAMP AT TIME ZONE 'asia/jakarta'  start_at_ts,
           (a.dw :: DATE :: TEXT || ' ' || c.finish_at || ':00') :: TIMESTAMP AT TIME ZONE 'asia/jakarta' finish_at_ts
         FROM (
                SELECT ts AT TIME ZONE 'asia/jakarta' dw
                FROM generate_series(now(), now() + '5 days', '1 day') g(ts)
              ) a
           INNER JOIN classes c ON c.day :: TEXT = to_char(dw, 'fmday')
       ) d
  WHERE d.finish_at_ts + INTERVAL '2 hours' > now()
        AND d.finish_at_ts + INTERVAL '2 hours' < now() + INTERVAL '7 days'
  ORDER BY d.id ASC;


CREATE TABLE sessions (
  id       BIGSERIAL PRIMARY KEY,
  class_id INT REFERENCES classes (id),
  _created TIMESTAMPTZ DEFAULT NOW(),
  _updated TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE sessions_tutors (
  id         BIGSERIAL PRIMARY KEY,
  session_id INT REFERENCES sessions (id),
  tutor_id   INT REFERENCES users (id),
  _created   TIMESTAMPTZ DEFAULT NOW(),
  _updated   TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE sessions_tutors_students (
  id                 BIGSERIAL PRIMARY KEY,
  status             BOOLEAN  NULL,
  feedback           TEXT     NULL,
  rating_interaction SMALLINT NULL,
  rating_cognition   SMALLINT NULL,
  rating_creativity  SMALLINT NULL,
  session_tutor_id   INT REFERENCES sessions_tutors (id),
  student_id         INT REFERENCES users (id),
  _created           TIMESTAMPTZ DEFAULT NOW(),
  _updated           TIMESTAMPTZ DEFAULT NOW()
);

CREATE OR REPLACE VIEW _sessions_tutors AS
  SELECT
    s.id,
    s._created,
    s.class_id,
    st.tutor_id
  FROM sessions s
    LEFT JOIN sessions_tutors st ON s.id = st.session_id;

CREATE OR REPLACE VIEW sessions_tutors_students_status_null AS
  SELECT
    sst.*,
    cs.student_id
  FROM _sessions_tutors sst
    INNER JOIN class_students cs ON sst.class_id = cs.class_id
    INNER JOIN _classes_ts c ON sst.class_id = c.id
  WHERE sst._created > (now() - '12:00:00' :: INTERVAL) AND cs.student_id NOT IN (
    SELECT student_id
    FROM sessions_tutors_students
    WHERE session_tutor_id = sst.id
  );

CREATE OR REPLACE VIEW _stats_tutors AS
  SELECT
    b.*,
    _st.id       _st_id,
    _st._created _st_created_at,
    _st.tutor_id _st_tutor_id,
    ss._created,
    ss.STATUS,
    ss.feedback,
    ss.rating_interaction,
    ss.rating_cognition,
    ss.rating_creativity
  FROM (
         SELECT
           a.ts :: DATE,
           c.id,
           c.tutor_id,
           (a.ts :: DATE :: TEXT || ' ' || c.start_at || ':00') :: TIMESTAMP AT TIME ZONE 'asia/jakarta'  start_at_ts,
           (a.ts :: DATE :: TEXT || ' ' || c.finish_at || ':00') :: TIMESTAMP AT TIME ZONE 'asia/jakarta' finish_at_ts,
           cs.student_id
         FROM (
                SELECT ts
                FROM generate_series((
                                       SELECT min(_created) :: DATE
                                       FROM sessions
                                     ), now() :: DATE, '1 day') g(ts)
              ) a
           INNER JOIN classes c ON c.day :: TEXT = to_char(a.ts, 'fmday')
           INNER JOIN class_students cs ON cs.class_id = c.id
         ORDER BY start_at_ts DESC
       ) b
    LEFT JOIN _sessions_tutors _st ON b.id = _st.class_id
                                      AND _st._created > b.start_at_ts + INTERVAL '5 minutes'
                                      AND _st._created < b.finish_at_ts + INTERVAL '2 hours'
    LEFT JOIN sessions_tutors_students ss ON _st.id = ss.session_tutor_id
                                             AND b.student_id = ss.student_id
  WHERE b.ts >= '2018-03-24' :: DATE
  ORDER BY b.start_at_ts DESC
    , _st.id DESC
