CREATE TABLE foo (barid serial not null);
CREATE TABLE bar (id serial not null, owner text not null);

-- name: JoinWhereClause :many
SELECT foo.*
FROM foo
JOIN bar ON bar.id = barid
WHERE owner = $1;

-- name: JoinParamWhereClause :many
SELECT foo.*
FROM foo
JOIN bar ON bar.id = $2
WHERE owner = $1;