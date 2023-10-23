CREATE table IF NOT EXISTS users
(
    id              serial primary key,
    name            varchar(255) not null,
    username        varchar(255)  not null unique,
    password_hash   varchar(255) not null
    );
CREATE TABLE IF NOT EXISTS todo_items
(
    id              serial primary key,
    title           varchar(255) not null,
    description     varchar(255),
    created_at      timestamp DEFAULT CURRENT_DATE,
    done            boolean not null default false
    );
CREATE TABLE IF NOT EXISTS users_items
(
    id              serial primary key,
    user_id         int references users(id) on delete cascade not null,
    item_id         int references todo_items(id) on delete cascade not null
    );

