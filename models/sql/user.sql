create table users(
    id serial primary key,
    email text unique not null,
    passwordhash text not null
);