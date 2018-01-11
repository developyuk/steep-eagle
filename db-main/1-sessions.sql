CREATE TABLE sessions (
  id BIGSERIAL PRIMARY KEY NOT NULL,
  created_at timestamp DEFAULT NOW(),
  class_id BIGSERIAL REFERENCES classes (id)
);
INSERT INTO sessions (class_id)
VALUES
  (1),
  (2),
  (2),
  (3),
  (3),
  (3);
CREATE TABLE sessions_tutors (
  id BIGSERIAL PRIMARY KEY NOT NULL,
  session_id BIGSERIAL REFERENCES sessions (id),
  tutor_id BIGSERIAL REFERENCES users (id),
  unique (session_id, tutor_id)
);
INSERT INTO sessions_tutors (session_id, tutor_id)
VALUES
  (1, 3),
  (2, 3),
  (2, 4),
  (3, 4);
CREATE TABLE sessions_students (
  id BIGSERIAL PRIMARY KEY NOT NULL,
  created_at timestamp DEFAULT NOW(),
  status BOOLEAN DEFAULT FALSE,
  feedback TEXT NULL,
  rating_interaction SMALLSERIAL,
  rating_cognition SMALLSERIAL,
  rating_creativity SMALLSERIAL,
  sessions_tutors_id BIGSERIAL REFERENCES sessions_tutors (id),
  student_id BIGSERIAL REFERENCES users (id)
);
INSERT INTO sessions_students (
  status, feedback, rating_interaction,
  rating_cognition, rating_creativity,
  sessions_tutors_id, student_id
)
VALUES
  (TRUE, 'feedback1', 1, 2, 3, 1, 1),
  (TRUE, 'feedback2', 2, 3, 4, 2, 1),
  (TRUE, 'feedback3', 4, 5, 0, 3, 2),
  (TRUE, 'feedback4', 5, 0, 1, 3, 1),
  (TRUE, 'feedback4', 0, 1, 2, 4, 1);
