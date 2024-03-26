-- name: GetUsers :many
select * from users;

-- name: CreateUser :one
insert into users (first_name, last_name, created_at, updated_at)
values ($1, $2, $3, $4)
returning *;

-- name: UpdateUser :one
update users
set first_name = $1, last_name = $2, updated_at = $3
where id = $4
returning *;
