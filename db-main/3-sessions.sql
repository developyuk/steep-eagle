CREATE TABLE sessions (
	id bigserial PRIMARY KEY
	,created_at TIMESTAMP WITH TIME zone NOT NULL DEFAULT now()
	,class_id INT REFERENCES classes(id)
	);

CREATE TABLE sessions_tutors (
	session_id INT REFERENCES sessions(id)
	,tutor_id INT REFERENCES users(id)
	);

CREATE
	OR replace VIEW _sessions_tutors AS

SELECT s.id
	,s.created_at
	,s.class_id
	,st.tutor_id
FROM sessions s
LEFT JOIN sessions_tutors st ON s.id = st.session_id;

CREATE TABLE sessions_students (
	created_at TIMESTAMP WITH TIME zone DEFAULT now()
	,STATUS boolean NULL
	,feedback TEXT NULL
	,rating_interaction SMALLINT NULL
	,rating_cognition SMALLINT NULL
	,rating_creativity SMALLINT NULL
	,session_id INT REFERENCES sessions(id)
	,student_id INT REFERENCES users(id)
	,tutor_id INT REFERENCES users(id)
	);

CREATE
	OR replace VIEW _sessions_students_status_null AS

SELECT sst.*
	,cs.student_id
FROM _sessions_tutors sst
INNER JOIN class_students cs ON sst.class_id = cs.class_id
INNER JOIN _classes_ts c ON sst.class_id = c.id
WHERE sst.created_at > (now() - '12:00:00'::interval) AND cs.student_id NOT IN (
		SELECT student_id
		FROM sessions_students
		WHERE session_id = sst.id
		);
