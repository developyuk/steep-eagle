CREATE TABLE program_types (
	id BIGSERIAL PRIMARY KEY NOT NULL
	,name TEXT NOT NULL
	);

INSERT INTO program_types (name)
VALUES ('regular')
	,('holiday')
	,('charity')
	,('afterschool')
	,('private');

CREATE TABLE programs (
	id BIGSERIAL PRIMARY KEY NOT NULL
	,name TEXT NOT NULL
	,type_id BIGSERIAL REFERENCES program_types(id)
	);

INSERT INTO programs (
	name
	,type_id
	)
VALUES (
	'program1'
	,1
	)
	,(
	'program2'
	,1
	)
	,(
	'program3'
	,3
	)
	,(
	'program4'
	,4
	)
	,(
	'program5'
	,5
	)
	,(
	'program3'
	,2
	);

CREATE TABLE modules (
	id BIGSERIAL PRIMARY KEY NOT NULL
	,name TEXT NOT NULL
	,IMAGE TEXT
	,session_total SMALLSERIAL
	);

INSERT INTO modules (
	name
	,IMAGE
	,session_total
	)
VALUES (
	'Web Lit'
	,'https://dl.dropboxusercontent.com/s/hg9zq2gra407viw/web_lit.jpg'
	,11
	)
	,(
	'App Inv'
	,'https://dl.dropboxusercontent.com/s/uk3xqoeipn6anm3/app_inv.jpg'
	,16
	)
	,(
	'HW1'
	,'https://dl.dropboxusercontent.com/s/3cdzkqbcxwbrymv/hw1.jpg'
	,16
	)
	,(
	'HW2'
	,'https://dl.dropboxusercontent.com/s/f7jwyg517ymu47x/hw2.jpg'
	,16
	);

CREATE TABLE program_modules (
	id BIGSERIAL PRIMARY KEY NOT NULL
	,module_id BIGSERIAL REFERENCES modules(id)
	,program_id BIGSERIAL REFERENCES programs(id)
	);

INSERT INTO program_modules (
	module_id
	,program_id
	)
VALUES (
	1
	,1
	)
	,(
	1
	,2
	)
	,(
	2
	,2
	)
	,(
	3
	,2
	)
	,(
	4
	,1
	);

CREATE TABLE branches (
	id BIGSERIAL PRIMARY KEY NOT NULL
	,name TEXT NOT NULL
	,IMAGE TEXT
	,address TEXT
	);

INSERT INTO branches (
	name
	,IMAGE
	,address
	)
VALUES (
	'branch1'
	,''
	,'address1'
	)
	,(
	'location2'
	,'https://www.shareicon.net/data/128x128/2016/05/24/769782_map_512x512.png'
	,''
	)
	,(
	'lokasi3'
	,'https://www.shareicon.net/data/128x128/2016/05/24/769782_map_512x512.png'
	,'adress3'
	);

CREATE TABLE classes (
	id BIGSERIAL PRIMARY KEY NOT NULL
	,name TEXT NOT NULL
	,IMAGE TEXT
	,day TEXT NOT NULL
	,TIME TEXT NOT NULL
	,module_id BIGSERIAL REFERENCES program_modules(id)
	,branch_id BIGSERIAL REFERENCES branches(id) CHECK (
		day IN (
			'monday'
			,'tuesday'
			,'wednesday'
			,'thursday'
			,'friday'
			,'saturday'
			,'sunday'
			)
		)
	);

INSERT INTO classes (
	name
	,module_id
	,branch_id
	,day
	,TIME
	,IMAGE
	)
VALUES (
	''
	,1
	,2
	,'sunday'
	,'12:00'
	,''
	)
	,(
	''
	,3
	,2
	,'sunday'
	,'13:00'
	,''
	)
	,(
	''
	,1
	,2
	,'monday'
	,'12:00'
	,''
	)
	,(
	''
	,4
	,1
	,'thursday'
	,'13:30'
	,'http://www.provideashelter.org/charity/images/group.png'
	)
	,(
	''
	,3
	,3
	,'friday'
	,'9:30'
	,'https://thumbs.dreamstime.com/z/group-little-kids-early-development-class-happy-drawing-filling-letters-paper-reading-48811717.jpg'
	)
	,(
	''
	,1
	,2
	,'tuesday'
	,'19:30'
	,'https://thumbs.dreamstime.com/z/group-little-kids-early-development-class-happy-drawing-filling-letters-paper-reading-48811717.jpg'
	)
	,(
	''
	,4
	,3
	,'wednesday'
	,'8:30'
	,'https://thumbs.dreamstime.com/z/group-little-kids-early-development-class-happy-drawing-filling-letters-paper-reading-48811717.jpg'
	)
	,(
	''
	,2
	,1
	,'wednesday'
	,'12:30'
	,'https://thumbs.dreamstime.com/z/group-little-kids-early-development-class-happy-drawing-filling-letters-paper-reading-48811717.jpg'
	)
	,(
	''
	,3
	,2
	,'saturday'
	,'12:30'
	,'https://thumbs.dreamstime.com/z/group-little-kids-early-development-class-happy-drawing-filling-letters-paper-reading-48811717.jpg'
	)
	,(
	''
	,2
	,2
	,'saturday'
	,'12:30'
	,'https://thumbs.dreamstime.com/z/group-little-kids-early-development-class-happy-drawing-filling-letters-paper-reading-48811717.jpg'
	);

 CREATE OR replace VIEW classes_ AS
    SELECT d.*
 FROM (
  SELECT c.*
    ,(a.dw::DATE::TEXT || ' ' || c.TIME || ':00')::TIMESTAMP AT TIME ZONE 'Asia/Jakarta' ts
  FROM (
    SELECT ts AT TIME ZONE 'Asia/Jakarta' dw
    FROM GENERATE_SERIES(NOW(), NOW() AT TIME ZONE 'Asia/Jakarta' + '7 DAYS', '1 DAY') g(ts)
    ) a
  INNER JOIN classes c ON c.day = to_char(dw, 'FMday')
  ) d
 WHERE d.ts > (NOW() - INTERVAL '1 hour')
  AND d.ts < (NOW() - INTERVAL '1 hour' + INTERVAL '7 days')
  ORDER by d.id asc;

CREATE TYPE roles AS ENUM (
	'operation'
	,'tutor'
	,'student'
	,'parent'
	,'partner'
	);

CREATE TABLE users (
	id BIGSERIAL PRIMARY KEY NOT NULL
	,name TEXT NOT NULL
	,first_name TEXT
	,last_name TEXT
	,email TEXT NOT NULL
	,pass TEXT NOT NULL
	,dob TEXT
	,photo TEXT
	,ROLE roles NOT NULL
	);

INSERT INTO users (
	name
	,email
	,pass
	,role
	,photo
	)
VALUES (
	'user1'
	,'user1@ex.com'
	,'asdqwe'
	,'student'
	,''
	)
	,(
	'user2'
	,'user2@ex.com'
	,'asdqwe'
	,'student'
	,'https://resources.kamernet.nl/Content/images/placeholder/no-pic-user.png'
	)
	,(
	'user3'
	,'user3@ex.com'
	,'asdqwe'
	,'tutor'
	,'https://resources.kamernet.nl/Content/images/placeholder/no-pic-user.png'
	)
	,(
	'user4'
	,'user4@ex.com'
	,'asdqwe'
	,'tutor'
	,'https://resources.kamernet.nl/Content/images/placeholder/no-pic-user.png'
	)
	,(
	'user5'
	,'user5@ex.com'
	,'asdqwe'
	,'parent'
	,'https://resources.kamernet.nl/Content/images/placeholder/no-pic-user.png'
	)
	,(
	'user6'
	,'user6@ex.com'
	,'asdqwe'
	,'parent'
	,'https://resources.kamernet.nl/Content/images/placeholder/no-pic-user.png'
	)
	,(
	'user7'
	,'user7@ex.com'
	,'asdqwe'
	,'operation'
	,'https://resources.kamernet.nl/Content/images/placeholder/no-pic-user.png'
	)
	,(
	'user8'
	,'user8@ex.com'
	,'asdqwe'
	,'operation'
	,'https://resources.kamernet.nl/Content/images/placeholder/no-pic-user.png'
	);

CREATE TABLE class_students (
	id BIGSERIAL PRIMARY KEY NOT NULL
	,class_id BIGSERIAL REFERENCES classes(id)
	,student_id BIGSERIAL REFERENCES users(id)
	);

INSERT INTO class_students (
	class_id
	,student_id
	)
VALUES (
	1
	,1
	)
	,(
	2
	,1
	)
	,(
	2
	,2
	)
	,(
	3
	,2
	);
