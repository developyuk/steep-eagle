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
	,UNIQUE (session_id
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

SELECT b.*
	,_st.id _st_id
	,_st.created_at _st_created_at
	,_st.tutor_id _st_tutor_id
	,ss.created_at
	,ss.STATUS
	,ss.feedback
	,ss.rating_interaction
	,ss.rating_cognition
	,ss.rating_creativity
FROM (
	SELECT a.ts::DATE
		,c.id
		,c.tutor_id
		,(a.ts::DATE::TEXT || ' ' || c.start_at || ':00')::TIMESTAMP at TIME zone 'asia/jakarta' start_at_ts
		,(a.ts::DATE::TEXT || ' ' || c.finish_at || ':00')::TIMESTAMP at TIME zone 'asia/jakarta' finish_at_ts
		,cs.student_id
	FROM (
		SELECT ts
		FROM generate_series((
					SELECT min(created_at)::DATE
					FROM sessions
					), now()::DATE, '1 day') g(ts)
		) a
	INNER JOIN classes c ON c.day::TEXT = to_char(a.ts, 'fmday')
	INNER JOIN class_students cs ON cs.class_id = c.id
	ORDER BY start_at_ts DESC
	) b
LEFT JOIN _sessions_tutors _st ON b.id = _st.class_id
	AND _st.created_at > b.start_at_ts + interval '5 minutes'
	AND _st.created_at < b.finish_at_ts + interval '2 hours'
LEFT JOIN sessions_students ss ON _st.id = ss.session_id
	AND b.student_id = ss.student_id
	AND _st.tutor_id = ss.tutor_id
WHERE b.ts >= '2018-03-24'::DATE
ORDER BY b.start_at_ts DESC
	,_st.id DESC
