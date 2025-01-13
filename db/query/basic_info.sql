-- name: CreateBasicInfo :one
INSERT INTO public.basic_info
(id, first_name, last_name, additional_name, pronouns, head_line)
VALUES(nextval('basic_info_id_seq'::regclass), $1, $2, $3, $4, $5)
    RETURNING *;

-- name: UpdateBasicInfo :exec
UPDATE public.basic_info
SET first_name=$1, last_name=$2, additional_name=$3, pronouns=$4, head_line=$5
WHERE id=$6;