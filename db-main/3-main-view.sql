CREATE OR REPLACE VIEW _users_profile AS
  SELECT *
  FROM users u
    LEFT JOIN users_profile up ON u.id = up.user_id;

CREATE OR REPLACE VIEW _classes_ts AS
  SELECT d.*
  FROM (
         SELECT
           c.*,
           (a.dw :: DATE :: TEXT || ' ' || c.start_at || ':00') :: TIMESTAMP AT TIME ZONE 'asia/jakarta'  start_at_ts,
           (a.dw :: DATE :: TEXT || ' ' || c.finish_at || ':00') :: TIMESTAMP AT TIME ZONE 'asia/jakarta' finish_at_ts
         FROM (
                SELECT ts AT TIME ZONE 'asia/jakarta' dw
                FROM generate_series(now(), now() + '5 days', '1 day') g(ts)
              ) a
           INNER JOIN classes c ON c.day :: TEXT = to_char(dw, 'fmday')
       ) d
  WHERE d.finish_at_ts + INTERVAL '2 hours' > now()
        AND d.finish_at_ts + INTERVAL '2 hours' < now() + INTERVAL '7 days'
  ORDER BY d.id ASC;

CREATE OR REPLACE VIEW _exports AS
  SELECT
    cs.class_id               id,
    initcap(pt.name)          program_name,
    initcap(p.name)        AS program,
    upper(m.name)             module,
    initcap(b.name)           branch,
    initcap(c.day :: TEXT) AS day,
    c.start_at,
    c.finish_at,
    initcap(us.name)          student_name,
    ut.email                  tutor_email,
    initcap(ut.username)      tutor_username,
    initcap(ut.name)          tutor_name
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
  ORDER BY cs.class_id;

CREATE OR REPLACE VIEW _classes_ts_search AS
  SELECT
    ct.*,
    a.q
  FROM _classes_ts ct
    JOIN (
           SELECT DISTINCT ON (id)
             id,
             lower(CONCAT(
                     module
                     , ', '
                     , branch
                     , ', '
                     , day
                     , ', '
                     , start_at
                     , ', '
                     , finish_at
                     , ', '
                     , tutor_username
                     , ', '
                     , tutor_name
                   )) q
           FROM _exports
         ) a ON ct.id = a.id;
