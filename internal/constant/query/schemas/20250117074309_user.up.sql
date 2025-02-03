CREATE TABLE IF NOT EXISTS users (
    id UUID primary key default gen_random_uuid(),
    user_name varchar(50) not null ,
    email varchar(50) unique not null ,
    password text not null,
    created_at timestamp  default current_timestamp,
    updated_at timestamp  default current_timestamp
);