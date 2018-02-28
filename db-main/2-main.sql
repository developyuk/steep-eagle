create table program_types (
  id bigserial primary key, name text not null
);
create table programs (
  id bigserial primary key,
  name text not null,
  type_id bigserial references program_types(id)
);
create table modules (
  id bigserial primary key, name text not null,
  image text, session_total smallserial
);
create table program_modules (
  id bigserial primary key,
  module_id bigserial references modules(id),
  program_id bigserial references programs(id)
);
create table branches (
  id bigserial primary key, name text not null,
  address text
);
create type days as enum (
  'monday', 'tuesday', 'wednesday',
  'thursday', 'friday', 'saturday',
  'sunday'
);
create table classes (
  id bigserial primary key,
  day days not null,
  start_at text not null,
  finish_at text not null,
  module_id bigserial references program_modules(id),
  branch_id bigserial references branches(id)
);
create
or replace view classes_ts as
select
  d.*
from
  (
    select
      c.*,
      (
        a.dw :: date :: text || ' ' || c.start_at || ':00'
      ):: timestamp at time zone 'asia/jakarta' start_at_ts,
      (
        a.dw :: date :: text || ' ' || c.finish_at || ':00'
      ):: timestamp at time zone 'asia/jakarta' finish_at_ts
    from
      (
        select
          ts at time zone 'asia/jakarta' dw
        from
          generate_series(
            now(), now() at time zone 'asia/jakarta' + '7 days',
            '1 day'
          ) g(ts)
      ) a
      inner join classes c on c.day :: text = to_char(dw, 'fmday')
  ) d
where
  d.start_at_ts > now()
  and d.finish_at_ts < now()  + interval '7 days'
order by
  d.id asc;
create type roles as enum (
  'operation', 'tutor', 'student', 'parent',
  'partner'
);
create table users (
  id bigserial primary key, email text not null unique,
  pass text not null, role roles not null
);
create
or replace function user_by_email_pass(_email text, _pass text) returns TABLE (id bigint, email text, role text) as $$ begin return query
select
  users.id,
  users.email,
  users.role :: text
from
  users
where
  users.email = user_by_email_pass._email
  and users.pass = crypt(
    user_by_email_pass._pass, users.pass
  );
end;
$$ LANGUAGE 'plpgsql';
create table users_profile (
  name text null,
  dob text null,
  photo text null,
  users_id bigserial references users(id)
);
create table class_students (
  id bigserial primary key,
  class_id bigserial references classes(id),
  student_id bigserial references users(id)
);