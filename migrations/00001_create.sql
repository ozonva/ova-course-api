-- +goose Up
create table courses
(
    id bigserial not null
        constraint courses_pk
            primary key,
    user_id bigint not null,
    name varchar(250) default '' not null,
    description text default '' not null,
    date_start timestamp not null,
    date_finish timestamp not null,
    date_created timestamp default now() not null
);


-- +goose Down
DROP TABLE courses;