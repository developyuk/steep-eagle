create table sessions (
  id bigserial primary key,
  created_at timestamp with time zone not null default now(),
  class_id bigserial references classes(id),
  tutor_id bigserial references users(id)
);
--insert into sessions (class_id, tutor_id)
--values
--  (1, 4),
--  (2, 3),
--  (2, 4),
--  (3, 3),
--  (3, 4),
--  (3, 3);
create table sessions_students (
  id bigserial primary key,
  created_at timestamp with time zone default now(),
  status boolean default false,
  feedback text null,
  rating_interaction smallserial,
  rating_cognition smallserial,
  rating_creativity smallserial,
  session_id bigserial references sessions(id),
  student_id bigserial references users(id)
);
--insert into sessions_students (
--  status, feedback, rating_interaction,
--  rating_cognition, rating_creativity,
--  session_id, student_id
--)
--values
--  (true, 'feedback1', 1, 2, 3, 1, 1),
--  (true, 'feedback2', 2, 3, 4, 2, 1),
--  (true, 'feedback3', 4, 5, 0, 3, 2),
--  (true, 'feedback4', 5, 0, 1, 3, 1),
--  (true, 'feedback4', 0, 1, 2, 4, 1);
