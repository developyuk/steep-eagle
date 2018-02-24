create role anon nologin;
create role operation nologin;
create role tutor nologin;
create role student nologin;
create role parent nologin;
create role partner nologin;
grant anon to postgres;
grant usage on schema public to anon;
revoke all on all tables in schema public
from
  public;
grant usage,
select
  on all sequences in schema public to operation,
  tutor,anon;
grant
select
  ,
  insert,
update
  ,
  delete on all tables in schema public to operation;
grant insert on table sessions,
sessions_students to tutor;
grant
select
  on all tables in schema public to tutor,
  student,
  anon,
  parent,
  partner;
