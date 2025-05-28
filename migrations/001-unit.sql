create table vacancies (
                           id serial primary key,
                           email varchar(255) not null,
                           role varchar(255),
                           company varchar(255),
                           salary varchar(255),
                           type varchar(255),
                           location varchar(255)
);

create table users (
                       id serial primary key,
                       name varchar(255),
                       email varchar(255) unique not null,
                       password_hash text not null
);
