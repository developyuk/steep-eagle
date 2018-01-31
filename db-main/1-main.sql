CREATE TABLE program_types (
  id BIGSERIAL PRIMARY KEY NOT NULL,
  name TEXT NOT NULL
);
INSERT INTO program_types (name)
VALUES
  ('regular'),
  ('holiday'),
  ('charity'),
  ('afterschool'),
  ('private');
CREATE TABLE programs (
  id BIGSERIAL PRIMARY KEY NOT NULL,
  name TEXT NOT NULL,
  type_id BIGSERIAL REFERENCES program_types (id)
);
INSERT INTO programs (name, type_id)
VALUES
  ('program1', 1),
  ('program2', 1),
  ('program3', 3),
  ('program4', 4),
  ('program5', 5),
  ('program3', 2);
CREATE TABLE modules (
  id BIGSERIAL PRIMARY KEY NOT NULL,
  name TEXT NOT NULL,
  image TEXT,
  session_total SMALLSERIAL
);
INSERT INTO modules (name, image, session_total)
VALUES
  (
    'module1', 'http://www.dewsoftoverseas.com/img/e-learning_icon.jpg',
    11
  ),
  ('module2', '', 16),
  (
    'module3', 'http://www.dewsoftoverseas.com/img/e-learning_icon.jpg',
    16
  );
CREATE TABLE program_modules (
  id BIGSERIAL PRIMARY KEY NOT NULL,
  module_id BIGSERIAL REFERENCES modules (id),
  program_id BIGSERIAL REFERENCES programs (id)
);
INSERT INTO program_modules (module_id, program_id)
VALUES
  (1, 1),
  (1, 2),
  (2, 2);
CREATE TABLE branches (
  id BIGSERIAL PRIMARY KEY NOT NULL,
  name TEXT NOT NULL,
  image TEXT,
  address TEXT
);
INSERT INTO branches (name, image, address)
VALUES
  ('branch1', '', 'address1'),
  (
    'branch2', 'https://www.shareicon.net/data/128x128/2016/05/24/769782_map_512x512.png',
    ''
  );
CREATE TABLE classes (
  id BIGSERIAL PRIMARY KEY NOT NULL,
  name TEXT NOT NULL,
  image TEXT,

  day TEXT NOT NULL,
  time TEXT NOT NULL,
  module_id BIGSERIAL REFERENCES program_modules (id),
  branch_id BIGSERIAL REFERENCES branches (id)

  CHECK (day IN ( 'monday', 'tuesday', 'wednesday', 'thursday', 'friday', 'saturday', 'sunday' ))
);
INSERT INTO classes (
  name, module_id, branch_id, day, time,
  image
)
VALUES
  ('', 1, 2, 'monday', '12:00', ''),
  (
    '', 2, 1, 'thursday', '13.30', 'http://www.provideashelter.org/charity/images/group.png'
  ),
  (
    '', 3, 1, 'friday', '9:30', 'https://thumbs.dreamstime.com/z/group-little-kids-early-development-class-happy-drawing-filling-letters-paper-reading-48811717.jpg'
  ),
  (
    '', 1, 2, 'tuesday', '19:30', 'https://thumbs.dreamstime.com/z/group-little-kids-early-development-class-happy-drawing-filling-letters-paper-reading-48811717.jpg'
  ),
  (
    '', 1, 2, 'wednesday', '8:30', 'https://thumbs.dreamstime.com/z/group-little-kids-early-development-class-happy-drawing-filling-letters-paper-reading-48811717.jpg'
  ),
  (
    '', 2, 1, 'wednesday', '12:30', 'https://thumbs.dreamstime.com/z/group-little-kids-early-development-class-happy-drawing-filling-letters-paper-reading-48811717.jpg'
  ),
  (
    '', 3, 1, 'saturday', '12:30', 'https://thumbs.dreamstime.com/z/group-little-kids-early-development-class-happy-drawing-filling-letters-paper-reading-48811717.jpg'
  ),
  (
    '', 1, 2, 'saturday', '12:30', 'https://thumbs.dreamstime.com/z/group-little-kids-early-development-class-happy-drawing-filling-letters-paper-reading-48811717.jpg'
  );
CREATE TYPE roles AS ENUM (
  'operation', 'tutor', 'student', 'parent',
  'partner'
);
CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY NOT NULL,
  name TEXT NOT NULL,
  first_name TEXT,
  last_name TEXT,
  email TEXT NOT NULL,
  pwd TEXT NOT NULL,
  dob TEXT,
  photo TEXT,
  role roles NOT NULL
);
INSERT INTO users (name, email, pwd, role, photo)
VALUES
  (
    'user1', 'user1@ex.com', 'asdqwe', 'student', ''
  ),
  (
    'user2', 'user2@ex.com', 'asdqwe', 'student', 'https://resources.kamernet.nl/Content/images/placeholder/no-pic-user.png'
  ),
  (
    'user3', 'user3@ex.com', 'asdqwe',
    'tutor', 'https://resources.kamernet.nl/Content/images/placeholder/no-pic-user.png'
  ),
  (
    'user4', 'user4@ex.com', 'asdqwe',
    'tutor', 'https://resources.kamernet.nl/Content/images/placeholder/no-pic-user.png'
  ),
  (
    'user5', 'user5@ex.com', 'asdqwe',
    'parent', 'https://resources.kamernet.nl/Content/images/placeholder/no-pic-user.png'
  ),
  (
    'user6', 'user6@ex.com', 'asdqwe',
    'parent', 'https://resources.kamernet.nl/Content/images/placeholder/no-pic-user.png'
  ),
  (
    'user7', 'user7@ex.com', 'asdqwe', 'operation',
    'https://resources.kamernet.nl/Content/images/placeholder/no-pic-user.png'
  ),
  (
    'user8', 'user8@ex.com', 'asdqwe', 'operation',
    'https://resources.kamernet.nl/Content/images/placeholder/no-pic-user.png'
  );
CREATE TABLE class_students (
  id BIGSERIAL PRIMARY KEY NOT NULL,
  class_id BIGSERIAL REFERENCES classes (id),
  student_id BIGSERIAL REFERENCES users (id)
);
INSERT INTO class_students (class_id, student_id)
VALUES
  (1, 1),
  (2, 1),
  (2, 2),
  (3, 2);
