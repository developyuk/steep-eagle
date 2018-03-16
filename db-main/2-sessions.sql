create table sessions (
  id bigserial primary key,
  created_at timestamp with time zone not null default now(),
  class_id integer references classes(id)
);
create table sessions_tutors (
  id bigserial primary key,
  session_id integer references sessions(id),
  tutor_id integer references users(id)
);

create
or replace view sessions__tutors as
select s.id,s.created_at,s.class_id,st.tutor_id
from sessions s
left join sessions_tutors st on s.id = st.session_id;
create table sessions_students (
  id bigserial primary key,
  created_at timestamp with time zone default now(),
  status boolean null,
  feedback text not null,
  rating_interaction smallint,
  rating_cognition smallint,
  rating_creativity smallint,
  session_id integer references sessions(id),
  student_id integer references users(id),
  tutor_id integer references users(id)
);
