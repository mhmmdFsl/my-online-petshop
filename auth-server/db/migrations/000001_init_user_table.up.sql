create table users(
                      id serial primary key,
                      name text not null ,
                      status text not null,
                      created_at timestamptz not null default current_timestamp,
                      updated_at timestamptz not null
);

create table password_users(
                               id serial primary key ,
                               user_id int not null ,
                               hash_password text not null ,
                               created_at timestamptz not null default current_timestamp,
                               foreign key (user_id) references users(id)
);

create table principal_users(
    id serial primary key ,
    user_id int not null ,
    principal_type text not null ,
    principal_value text not null ,
    status text not null ,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null ,
    foreign key (user_id) references users(id)
);
