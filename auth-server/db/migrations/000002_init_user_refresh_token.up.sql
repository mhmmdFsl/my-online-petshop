create table if not exists user_refresh_tokens(
    id serial primary key ,
    user_id int not null ,
    refresh_token text not null ,
    created_at timestamptz not null default current_timestamp ,
    expired_at timestamptz not null ,
    foreign key (user_id) references users(id)
);