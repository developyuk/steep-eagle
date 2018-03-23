CREATE TABLE sessions_logins (
	created_at TIMESTAMP WITH TIME zone NOT NULL DEFAULT now()
	,status boolean NOT NULL DEFAULT false
	,user_id INT REFERENCES users(id)
	);

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
	,status boolean NULL
	,feedback TEXT NULL
	,rating_interaction SMALLINT NULL
	,rating_cognition SMALLINT NULL
	,rating_creativity SMALLINT NULL
	,session_id INT REFERENCES sessions(id)
	,student_id INT REFERENCES users(id)
	,tutor_id INT REFERENCES users(id)
	,UNIQUE (
		status
		,session_id
		,student_id
		,tutor_id
		)
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

CREATE
	OR replace VIEW _stats_tutors AS

SELECT a.ts
	,c.id
	,(a.ts::DATE::TEXT || ' ' || c.start_at || ':00')::TIMESTAMP at TIME zone 'asia/jakarta' start_at_ts
	,(a.ts::DATE::TEXT || ' ' || c.finish_at || ':00')::TIMESTAMP at TIME zone 'asia/jakarta' finish_at_ts
	,cs.class_id cs_class_id
	,cs.student_id cs_student_id
	,_st.id s_id
	,_st.created_at s_created_at
	,_st.class_id s_class_id
	,_st.tutor_id s_tutor_id
	,ss.*
FROM (
	SELECT ts
	FROM generate_series((
				SELECT min(created_at)::DATE
				FROM sessions
				), now()::DATE, '1 day') g(ts)
	) a
INNER JOIN classes c ON c.day::TEXT = to_char(a.ts, 'fmday')
INNER JOIN class_students cs ON cs.class_id = c.id
LEFT JOIN _sessions_tutors _st ON c.id = _st.class_id
	AND a.ts::DATE = _st.created_at::DATE
LEFT JOIN sessions_students ss ON _st.tutor_id = ss.tutor_id
	AND _st.id = ss.session_id
	AND cs.student_id = ss.student_id
ORDER BY start_at_ts ASC
