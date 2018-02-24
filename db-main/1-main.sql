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
  ('program2', 3),
  ('program3', 2),
  ('program4', 4),
  ('program5', 5),
  ('program6', 1);
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
    'hw 1', 'https://dl.dropboxusercontent.com/s/3cdzkqbcxwbrymv/hw1.jpg',
    16
  ),
  (
    'hw 2', 'https://dl.dropboxusercontent.com/s/f7jwyg517ymu47x/hw2.jpg',
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
  (3, 2),
  (2, 2),
  (4, 3),
  (1, 1);
create table branches (
  id bigserial primary key, name text not null, address text
);
insert into branches (name, address)
values
  ('pondok indah', 'pondok indah address'),
  (
    'kemang', 'kemang address'
  ),
  (
    'pondok', 'pondok address'
  );
create table classes (
  id bigserial primary key,
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
  module_id, branch_id, day, time
)
values
  (4, 1, 'saturday', '13:00'),
  (3, 2, 'saturday', '13:00'),
  (1, 3, 'saturday', '15:00'),
  (1, 1, 'saturday', '20:00');
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
    'john snow', 'john.snow@ex.com', 'asdqwe',
    'student', 'https://resources.kamernet.nl/content/images/placeholder/no-pic-user.png'
  ),
  (
    'daenarys targaryen', 'daenarys.targaryen@ex.com', 'asdqwe',
    'student', 'https://resources.kamernet.nl/content/images/placeholder/no-pic-user.png'
  ),
  (
    'tyrion lannister', 'tyrion.lannister@ex.com', 'asdqwe',
    'student', 'https://resources.kamernet.nl/content/images/placeholder/no-pic-user.png'
  ),
  (
    'jamie lannister', 'jamie.lannister@ex.com', 'asdqwe',
    'student', 'https://resources.kamernet.nl/content/images/placeholder/no-pic-user.png'
  ),


  (
    'aristotle', 'aristotle@ex.com', 'asdqwe',
    'student', 'https://resources.kamernet.nl/content/images/placeholder/no-pic-user.png'
  ),
  (
    'plato', 'plato@ex.com', 'asdqwe',
    'student', 'https://resources.kamernet.nl/content/images/placeholder/no-pic-user.png'
  ),
  (
    'pythagoras', 'pythagoras@ex.com', 'asdqwe',
    'student', 'https://resources.kamernet.nl/content/images/placeholder/no-pic-user.png'
  ),
  (
    'socrates', 'socrates@ex.com', 'asdqwe',
    'student', 'https://resources.kamernet.nl/content/images/placeholder/no-pic-user.png'
  ),
  (
    'parminedes', 'parminedes@ex.com', 'asdqwe',
    'student', 'https://resources.kamernet.nl/content/images/placeholder/no-pic-user.png'
  ),

  (
    'tony stark', 'tony.stark@ex.com', 'asdqwe',
    'student', 'https://resources.kamernet.nl/content/images/placeholder/no-pic-user.png'
  ),
  (
    'steve rogers', 'steve.rogers@ex.com', 'asdqwe',
    'student', 'https://resources.kamernet.nl/content/images/placeholder/no-pic-user.png'
  ),
  (
    'bruce banner', 'bruce.banner@ex.com', 'asdqwe',
    'student', 'https://resources.kamernet.nl/content/images/placeholder/no-pic-user.png'
  ),
  (
    'natasha romanoff', 'natasha.romanoff@ex.com', 'asdqwe',
    'student', 'https://resources.kamernet.nl/content/images/placeholder/no-pic-user.png'
  ),
  (
    'peter parker', 'peter.parker@ex.com', 'asdqwe',
    'student', 'https://resources.kamernet.nl/content/images/placeholder/no-pic-user.png'
  ),
  (
    'scott lang', 'scott.lang@ex.com', 'asdqwe',
    'student', 'https://resources.kamernet.nl/content/images/placeholder/no-pic-user.png'
  ),
  (
    'stephen strange', 'stephen.strange@ex.com', 'asdqwe',
    'student', 'https://resources.kamernet.nl/content/images/placeholder/no-pic-user.png'
  ),
  (
    'james rhodes', 'james.rhodes@ex.com', 'asdqwe',
    'student', 'https://resources.kamernet.nl/content/images/placeholder/no-pic-user.png'
  ),
  (
    'maria hill', 'maria.hill@ex.com', 'asdqwe',
    'student', 'https://resources.kamernet.nl/content/images/placeholder/no-pic-user.png'
  ),

  (
    'thomas alva edison', 'thomas.alva.edison@ex.com', 'asdqwe',
    'student', 'https://resources.kamernet.nl/content/images/placeholder/no-pic-user.png'
  ),
  (
    'tesla', 'tesla@ex.com', 'asdqwe',
    'student', 'https://resources.kamernet.nl/content/images/placeholder/no-pic-user.png'
  ),
  (
    'isaac newton', 'isaac.newton@ex.com', 'asdqwe',
    'student', 'https://resources.kamernet.nl/content/images/placeholder/no-pic-user.png'
  ),
  (
    'michael faraday', 'michael.faraday@ex.com', 'asdqwe',
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
  (1, 2),
  (1, 3),
  (1, 4),
  (2, 5),
  (2, 6),
  (2, 7),
  (2, 8),
  (2, 9),
  (3, 10),
  (3, 11),
  (3, 12),
  (3, 13),
  (3, 14),
  (3, 15),
  (3, 16),
  (3, 17),
  (3, 18),
  (4, 19),
  (4, 20),
  (4, 21),
  (4, 22);
