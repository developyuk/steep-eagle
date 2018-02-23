create table program_types (
  id bigserial primary key, name text not null
);
insert into program_types (name)
values
  ('regular'),
  ('holiday'),
  ('charity'),
  ('afterschool'),
  ('private');
create table programs (
  id bigserial primary key,
  name text not null,
  type_id bigserial references program_types(id)
);
insert into programs (name, type_id)
values
  ('program1', 1),
  ('program2', 1),
  ('program3', 3),
  ('program4', 4),
  ('program5', 5),
  ('program3', 2);
create table modules (
  id bigserial primary key, name text not null,
  image text, session_total smallserial
);
insert into modules (name, image, session_total)
values
  (
    'web lit', 'https://dl.dropboxusercontent.com/s/hg9zq2gra407viw/web_lit.jpg',
    11
  ),
  (
    'app inv', 'https://dl.dropboxusercontent.com/s/uk3xqoeipn6anm3/app_inv.jpg',
    16
  ),
  (
    'hw1', 'https://dl.dropboxusercontent.com/s/3cdzkqbcxwbrymv/hw1.jpg',
    16
  ),
  (
    'hw2', 'https://dl.dropboxusercontent.com/s/f7jwyg517ymu47x/hw2.jpg',
    16
  );
create table program_modules (
  id bigserial primary key,
  module_id bigserial references modules(id),
  program_id bigserial references programs(id)
);
insert into program_modules (module_id, program_id)
values
  (1, 1),
  (1, 2),
  (2, 2),
  (3, 2),
  (4, 1);
create table branches (
  id bigserial primary key, name text not null,
  image text, address text
);
insert into branches (name, image, address)
values
  ('branch1', '', 'address1'),
  (
    'location2', 'https://www.shareicon.net/data/128x128/2016/05/24/769782_map_512x512.png',
    ''
  ),
  (
    'lokasi3', 'https://www.shareicon.net/data/128x128/2016/05/24/769782_map_512x512.png',
    'adress3'
  );
create table classes (
  id bigserial primary key,
  name text not null,
  day text not null,
  time text not null,
  module_id bigserial references program_modules(id),
  branch_id bigserial references branches(id) check (
    day in (
      'monday', 'tuesday', 'wednesday',
      'thursday', 'friday', 'saturday',
      'sunday'
    )
  )
);
insert into classes (
  name, module_id, branch_id, day, time
)
values
  ('', 1, 2, 'sunday', '12:00'),
  ('', 3, 2, 'sunday', '13:00'),
  ('', 1, 2, 'monday', '12:00'),
  ('', 4, 1, 'thursday', '13:30'),
  ('', 3, 3, 'friday', '9:30'),
  ('', 1, 2, 'tuesday', '19:30'),
  ('', 4, 3, 'wednesday', '8:30'),
  ('', 2, 1, 'wednesday', '12:30'),
  ('', 3, 2, 'saturday', '12:30'),
  ('', 2, 2, 'saturday', '12:30');
create
or replace view classes_ as
select
  d.*
from
  (
    select
      c.*,
      (
        a.dw :: date :: text || ' ' || c.time || ':00'
      ):: timestamp at time zone 'asia/jakarta' ts
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
      inner join classes c on c.day = to_char(dw, 'fmday')
  ) d
where
  d.ts > (now() - interval '1 hour')
  and d.ts < (
    now() - interval '1 hour' + interval '7 days'
  )
order by
  d.id asc;
create type roles as enum (
  'operation', 'tutor', 'student', 'parent',
  'partner'
);
create table users (
  id bigserial primary key, name text not null,
  first_name text, last_name text, email text not null,
  pass text not null, dob text, photo text,
  role roles not null
);
insert into users (name, email, pass, role, photo)
values
  (
    'user1', 'user1@ex.com', 'asdqwe',
    'student', ''
  ),
  (
    'user2', 'user2@ex.com', 'asdqwe',
    'student', 'https://resources.kamernet.nl/content/images/placeholder/no-pic-user.png'
  ),
  (
    'user3', 'user3@ex.com', 'asdqwe',
    'tutor', 'https://resources.kamernet.nl/content/images/placeholder/no-pic-user.png'
  ),
  (
    'user4', 'user4@ex.com', 'asdqwe',
    'tutor', 'https://resources.kamernet.nl/content/images/placeholder/no-pic-user.png'
  ),
  (
    'user5', 'user5@ex.com', 'asdqwe',
    'parent', 'https://resources.kamernet.nl/content/images/placeholder/no-pic-user.png'
  ),
  (
    'user6', 'user6@ex.com', 'asdqwe',
    'parent', 'https://resources.kamernet.nl/content/images/placeholder/no-pic-user.png'
  ),
  (
    'user7', 'user7@ex.com', 'asdqwe',
    'operation', 'https://resources.kamernet.nl/content/images/placeholder/no-pic-user.png'
  ),
  (
    'user8', 'user8@ex.com', 'asdqwe',
    'operation', 'https://resources.kamernet.nl/content/images/placeholder/no-pic-user.png'
  );
create table class_students (
  id bigserial primary key,
  class_id bigserial references classes(id),
  student_id bigserial references users(id)
);
insert into class_students (class_id, student_id)
values
  (1, 1),
  (2, 1),
  (2, 2),
  (3, 2);
