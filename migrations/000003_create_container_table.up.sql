create table if not exists containers (
    id bigserial primary key ,
    name varchar,
    image_id bigint references images("id"),
    command varchar,
    created timestamp,
    status int8
)