CREATE TABLE program_types (
	id bigserial PRIMARY KEY
	,name TEXT NOT NULL UNIQUE
	);

CREATE TABLE programs (
	id bigserial PRIMARY KEY
	,name TEXT NOT NULL
	,type_id INT REFERENCES program_types(id)
	,UNIQUE (
		name
		,type_id
		)
	);

CREATE TABLE modules (
	id bigserial PRIMARY KEY
	,name TEXT NOT NULL UNIQUE
	,image TEXT NULL
	);

CREATE TABLE programs_modules (
	id bigserial PRIMARY KEY
	,module_id INT REFERENCES modules(id)
	,program_id INT REFERENCES programs(id)
	,UNIQUE (
		module_id
		,program_id
		)
	);

CREATE TABLE branches (
	id bigserial PRIMARY KEY
	,name TEXT NOT NULL UNIQUE
	,address TEXT NULL
	);

CREATE type days AS enum (
	'monday'
	,'tuesday'
	,'wednesday'
	,'thursday'
	,'friday'
	,'saturday'
	,'sunday'
	);

CREATE type roles AS enum (
	'operation'
	,'tutor'
	,'student'
	,'parent'
	,'partner'
	);

CREATE TABLE users (
	id bigserial PRIMARY KEY
	,username TEXT NOT NULL
	,email TEXT NULL
	,pass TEXT NULL
	,ROLE roles NOT NULL
	,UNIQUE (
		username
		,email
		)
	);

CREATE TABLE users_profile (
	name TEXT NULL
	,dob TEXT NULL
	,photo TEXT NULL
	,user_id INT REFERENCES users(id)
	,UNIQUE (
		name
		,user_id
		)
	);

CREATE
	OR replace VIEW _users_profile AS

SELECT *
FROM users u
LEFT JOIN users_profile up ON u.id = up.user_id;

CREATE
	OR replace FUNCTION user_by_email_pass (
	_email TEXT
	,_pass TEXT
	)
RETURNS TABLE (
		id BIGINT
		,email TEXT
		,ROLE TEXT
		,name TEXT
		,photo TEXT
		) AS $$

BEGIN
	RETURN query

	SELECT users.id
		,users.email
		,users.ROLE::TEXT
		,users_profile.name
		,users_profile.photo
	FROM users
	LEFT JOIN users_profile ON users.id = users_profile.user_id
	WHERE users.email = user_by_email_pass._email
		AND users.pass = crypt(user_by_email_pass._pass, users.pass);
END;$$

LANGUAGE 'plpgsql';

CREATE TABLE classes (
	id bigserial PRIMARY KEY
	,day days NOT NULL
	,start_at CHAR(5) NOT NULL
	,finish_at CHAR(5) NOT NULL
	,program_module_id INT REFERENCES programs_modules(id)
	,branch_id INT REFERENCES branches(id)
	,tutor_id INT REFERENCES users(id)
	,UNIQUE (
		day
		,start_at
		,finish_at
		,program_module_id
		,branch_id
		)
	);

CREATE
	OR replace VIEW _classes_ts AS

SELECT d.*
FROM (
	SELECT c.*
		,(a.dw::DATE::TEXT || ' ' || c.start_at || ':00')::TIMESTAMP at TIME zone 'asia/jakarta' start_at_ts
		,(a.dw::DATE::TEXT || ' ' || c.finish_at || ':00')::TIMESTAMP at TIME zone 'asia/jakarta' finish_at_ts
	FROM (
		SELECT ts at TIME zone 'asia/jakarta' dw
		FROM generate_series(now(), now() + '7 days', '1 day') g(ts)
		) a
	INNER JOIN classes c ON c.day::TEXT = to_char(dw, 'fmday')
	) d
WHERE d.finish_at_ts + interval '2 hours' > now()
	AND d.finish_at_ts + interval '2 hours' < now() + interval '7 days'
ORDER BY d.id ASC;

CREATE TABLE class_students (
	class_id INT REFERENCES classes(id)
	,student_id INT REFERENCES users(id)
	,UNIQUE (
		class_id
		,student_id
		)
	);

CREATE
	OR replace VIEW _exports AS

SELECT initcap(pt.name) program_name
	,initcap(p.name) AS program
	,upper(m.name) module
	,initcap(b.name) branch
	,initcap(c.day::TEXT)
	,c.start_at
	,c.finish_at
	,initcap(us.name) student_name
	,ut.email tutor_email
	,initcap(ut.username) tutor_username
	,initcap(ut.name) tutor_name
FROM class_students cs
LEFT JOIN classes c ON cs.class_id = c.id
LEFT JOIN branches b ON c.branch_id = b.id
LEFT JOIN programs_modules pm ON c.program_module_id = pm.id
LEFT JOIN modules m ON pm.module_id = m.id
LEFT JOIN programs p ON pm.program_id = p.id
LEFT JOIN program_types pt ON p.type_id = pt.id
LEFT JOIN _users_profile ut ON c.tutor_id = ut.id
	AND ut.ROLE = 'tutor'
LEFT JOIN _users_profile us ON cs.student_id = us.id
	AND us.ROLE = 'student'
ORDER BY cs.class_id
