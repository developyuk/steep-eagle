CREATE TABLE sessions_logins (
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  status     BOOLEAN     NOT NULL DEFAULT FALSE,
  user_id    INT REFERENCES users (id)
);

CREATE TABLE sessions (
  id         BIGSERIAL PRIMARY KEY,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  class_id   INT REFERENCES classes (id)
);

CREATE TABLE sessions_tutors (
  session_id INT REFERENCES sessions (id),
  tutor_id   INT REFERENCES users (id)
);

CREATE TABLE sessions_students (
  created_at         TIMESTAMPTZ DEFAULT now(),
  status             BOOLEAN  NULL,
  feedback           TEXT     NULL,
  rating_interaction SMALLINT NULL,
  rating_cognition   SMALLINT NULL,
  rating_creativity  SMALLINT NULL,
  session_id         INT REFERENCES sessions (id),
  student_id         INT REFERENCES users (id),
  tutor_id           INT REFERENCES users (id),
  UNIQUE (session_id
    , student_id
    , tutor_id
  )
);
