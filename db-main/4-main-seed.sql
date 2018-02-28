insert into program_types (name)
values
  ('regular'),
  ('holiday'),
  ('charity'),
  ('afterschool'),
  ('private');
insert into programs (name, type_id)
values
  ('program1', 1),
  ('program2', 3),
  ('program3', 2),
  ('program4', 4),
  ('program5', 5),
  ('program6', 1);
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
insert into program_modules (module_id, program_id)
values
  (1, 1),
  (3, 2),
  (2, 2),
  (4, 3),
  (1, 1);
insert into branches (name, address)
values
  (
    'pondok indah', 'pondok indah address'
  ),
  ('kemang', 'kemang address'),
  ('pondok', 'pondok address');
insert into classes (module_id, branch_id, day, start_at, finish_at)
values
  (4, 1, 'saturday', '13:00', '15:00'),
  (3, 2, 'saturday', '13:00', '15:00'),
  (1, 3, 'saturday', '15:00', '17:00'),
  (1, 1, 'saturday', '20:00', '20:00');
insert into users (email, pass, role)
values
  (
    'john.snow@ex.com',
    crypt(
      'asdqwe',
      gen_salt('bf')
    ),
    'student'
  ),
  (
    'daenarys.targaryen@ex.com',
    crypt(
      'asdqwe',
      gen_salt('bf')
    ),
    'student'
  ),
  (
    'tyrion.lannister@ex.com',
    crypt(
      'asdqwe',
      gen_salt('bf')
    ),
    'student'
  ),
  (
    'jamie.lannister@ex.com',
    crypt(
      'asdqwe',
      gen_salt('bf')
    ),
    'student'
  ),
  (
    'aristotle@ex.com',
    crypt(
      'asdqwe',
      gen_salt('bf')
    ),
    'student'
  ),
  (
    'plato@ex.com',
    crypt(
      'asdqwe',
      gen_salt('bf')
    ),
    'student'
  ),
  (
    'pythagoras@ex.com',
    crypt(
      'asdqwe',
      gen_salt('bf')
    ),
    'student'
  ),
  (
    'socrates@ex.com',
    crypt(
      'asdqwe',
      gen_salt('bf')
    ),
    'student'
  ),
  (
    'parminedes@ex.com',
    crypt(
      'asdqwe',
      gen_salt('bf')
    ),
    'student'
  ),
  (
    'tony.stark@ex.com',
    crypt(
      'asdqwe',
      gen_salt('bf')
    ),
    'student'
  ),
  (
    'steve.rogers@ex.com',
    crypt(
      'asdqwe',
      gen_salt('bf')
    ),
    'student'
  ),
  (
    'bruce.banner@ex.com',
    crypt(
      'asdqwe',
      gen_salt('bf')
    ),
    'student'
  ),
  (
    'natasha.romanoff@ex.com',
    crypt(
      'asdqwe',
      gen_salt('bf')
    ),
    'student'
  ),
  (
    'peter.parker@ex.com',
    crypt(
      'asdqwe',
      gen_salt('bf')
    ),
    'student'
  ),
  (
    'scott.lang@ex.com',
    crypt(
      'asdqwe',
      gen_salt('bf')
    ),
    'student'
  ),
  (
    'stephen.strange@ex.com',
    crypt(
      'asdqwe',
      gen_salt('bf')
    ),
    'student'
  ),
  (
    'james.rhodes@ex.com',
    crypt(
      'asdqwe',
      gen_salt('bf')
    ),
    'student'
  ),
  (
    'maria.hill@ex.com',
    crypt(
      'asdqwe',
      gen_salt('bf')
    ),
    'student'
  ),
  (
    'thomas.alva.edison@ex.com',
    crypt(
      'asdqwe',
      gen_salt('bf')
    ),
    'student'
  ),
  (
    'tesla@ex.com',
    crypt(
      'asdqwe',
      gen_salt('bf')
    ),
    'student'
  ),
  (
    'isaac.newton@ex.com',
    crypt(
      'asdqwe',
      gen_salt('bf')
    ),
    'student'
  ),
  (
    'michael.faraday@ex.com',
    crypt(
      'asdqwe',
      gen_salt('bf')
    ),
    'student'
  ),
  (
    'user3@ex.com',
    crypt(
      'asdqwe',
      gen_salt('bf')
    ),
    'tutor'
  ),
  (
    'user4@ex.com',
    crypt(
      'asdqwe',
      gen_salt('bf')
    ),
    'tutor'
  ),
  (
    'user5@ex.com',
    crypt(
      'asdqwe',
      gen_salt('bf')
    ),
    'parent'
  ),
  (
    'user6@ex.com',
    crypt(
      'asdqwe',
      gen_salt('bf')
    ),
    'parent'
  ),
  (
    'user7@ex.com',
    crypt(
      'asdqwe',
      gen_salt('bf')
    ),
    'operation'
  ),
  (
    'user8@ex.com',
    crypt(
      'asdqwe',
      gen_salt('bf')
    ),
    'operation'
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
