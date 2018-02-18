CREATE TABLE sessions (
  id BIGSERIAL PRIMARY KEY NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  class_id BIGSERIAL REFERENCES classes (id),
  tutor_id BIGSERIAL REFERENCES users (id)
);
INSERT INTO sessions (class_id,tutor_id)
VALUES
  (1,4),
  (2,3),
  (2,4),
  (3,3),
  (3,4),
  (3,3);
CREATE TABLE sessions_students (
  id BIGSERIAL PRIMARY KEY NOT NULL,
  created_at timestamp DEFAULT NOW(),
  status BOOLEAN DEFAULT FALSE,
  feedback TEXT NULL,
  rating_interaction SMALLSERIAL,
  rating_cognition SMALLSERIAL,
  rating_creativity SMALLSERIAL,
  session_id BIGSERIAL REFERENCES sessions (id),
  student_id BIGSERIAL REFERENCES users (id)
);
INSERT INTO sessions_students (
  status, feedback, rating_interaction,
  rating_cognition, rating_creativity,
  session_id, student_id
)
VALUES
  (TRUE, 'feedback1', 1, 2, 3, 1, 1),
  (TRUE, 'feedback2', 2, 3, 4, 2, 1),
  (TRUE, 'feedback3', 4, 5, 0, 3, 2),
  (TRUE, 'feedback4', 5, 0, 1, 3, 1),
  (TRUE, 'feedback4', 0, 1, 2, 4, 1);
