create table sessions (
  id bigserial primary key,
  created_at timestamp with time zone not null default now(),
  class_id bigserial references classes(id),
  tutor_id bigserial references users(id)
);
create table sessions_students (
  id bigserial primary key,
  created_at timestamp with time zone default now(),
  status boolean null,
  feedback text null,
  rating_interaction smallserial,
  rating_cognition smallserial,
  rating_creativity smallserial,
  session_id bigserial references sessions(id),
  student_id bigserial references users(id)
);