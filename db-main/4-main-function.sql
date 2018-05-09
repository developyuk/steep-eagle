CREATE OR REPLACE FUNCTION user_by_email_pass(
  _email TEXT,
  _pass  TEXT
)
  RETURNS TABLE(
    id    BIGINT,
    email TEXT,
    ROLE  TEXT,
    name  TEXT,
    photo TEXT
  ) AS $$

BEGIN
  RETURN QUERY

  SELECT
    users.id,
    users.email,
    users.ROLE :: TEXT,
    users_profile.name,
    users_profile.photo
  FROM users
    LEFT JOIN users_profile ON users.id = users_profile.user_id
  WHERE users.email = user_by_email_pass._email
        AND users.pass = crypt(user_by_email_pass._pass, users.pass);
END;$$

LANGUAGE 'plpgsql';

