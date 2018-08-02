-- truncate all table
SELECT 'TRUNCATE ' || table_name || ' CASCADE;'
FROM information_schema.tables
WHERE table_schema='public' AND table_type='BASE TABLE'
ORDER BY table_name;

-- fixing sequence
SELECT 'SELECT SETVAL(' ||
      quote_literal(quote_ident(PGT.schemaname) || '.' || quote_ident(S.relname)) ||
      ', COALESCE(MAX(' ||quote_ident(C.attname)|| '), 1) ) FROM ' ||
      quote_ident(PGT.schemaname)|| '.'||quote_ident(T.relname)|| ';'
FROM pg_class AS S,
    pg_depend AS D,
    pg_class AS T,
    pg_attribute AS C,
    pg_tables AS PGT
WHERE S.relkind = 'S'
   AND S.oid = D.objid
   AND D.refobjid = T.oid
   AND D.refobjid = C.attrelid
   AND D.refobjsubid = C.attnum
   AND T.relname = PGT.tablename
ORDER BY S.relname;


-- select count(id)
-- from (
-- SELECT DISTINCT ON (id,ts) *
-- FROM "_stats_tutors"
-- where tutor_id=1) a;
--
-- select sum(duration)
-- from (
-- select finish_at_ts-start_at_ts duration,*
-- FROM (
-- SELECT DISTINCT ON (id,ts) *
-- FROM "_stats_tutors"
-- where tutor_id=1) a) b;
--
-- select sum(id)
-- FROM "_stats_tutors"
-- where tutor_id=1 and feedback is not null and btrim(feedback,'') != ''
