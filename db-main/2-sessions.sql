create table sessions (
  id bigserial primary key,
  created_at timestamp with time zone not null default now(),
  class_id integer references classes(id),
  tutor_id integer references users(id)
);
create table sessions_students (
  id bigserial primary key,
  created_at timestamp with time zone default now(),
  status boolean null,
  feedback text not null,
  rating_interaction smallserial,
  rating_cognition smallserial,
  rating_creativity smallserial,
  session_id integer references sessions(id),
  student_id integer references users(id),
  tutor_id integer references users(id)
);
