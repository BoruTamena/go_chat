create table IF NOT EXISTS groups (
    id UUID primary key default gen_random_uuid(),
    group_name varchar(50) unique not null ,
    owner_id UUID not null ,
    description text not null ,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    constraint owner_constarint foreign key (owner_id) references users(id) on delete cascade
) ;