create role anon nologin;
create role operation nologin;
create role tutor nologin;
grant all privileges on sequence users_id_seq to anon;
grant all privileges on function user_by_email_pass to anon;
grant select on table users,users_profile to anon;
grant all privileges on all sequences in schema public to operation, tutor;
grant all privileges on all tables in schema public to operation;
grant select on classes,classes_ts,branches,modules,users,users_profile to tutor;
grant insert on table sessions,
sessions_students to tutor;
