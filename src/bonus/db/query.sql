-- name: GetPrivilege :one
SELECT * FROM privilege
WHERE username=$1 LIMIT 1;

-- name: GetPrivilegeById :one
SELECT * FROM privilege
WHERE id=$1 LIMIT 1;

-- name: ListPrivilegeHistories :many
SELECT * FROM privilege_history
WHERE privilege_id=$1;

-- name: CreatePrivilege :one
INSERT INTO privilege(username, balance)
VALUES ($1, 0)
RETURNING *;

-- name: UpdatePrivilegeBalance :exec
UPDATE privilege
SET balance=$2
WHERE username=$1;

-- name: CreatePrivilegeHistory :one
INSERT INTO privilege_history(privilege_id, ticket_uid, datetime, balance_diff, operation_type)
VALUES ($1, $2, NOW(), $3, $4)
RETURNING *;

-- name: GetPrivilegeHistory :one
SELECT * FROM privilege_history
WHERE ticket_uid=$1 LIMIT 1;

-- name: DeletePrivilegeHistory :exec
DELETE FROM privilege_history
WHERE ticket_uid=$1;