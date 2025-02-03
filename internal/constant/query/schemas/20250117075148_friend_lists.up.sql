CREATE TABLE IF NOT EXISTS friend_lists (
    id UUID primary key default gen_random_uuid(),
    user_id UUID not null ,
    friend_id  UUID not null ,
    status text check(status in('pending','accepted','blocked')) default 'pending',
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    constraint unique_friendship unique(user_id,friend_id),
    constraint user_1 foreign key (user_id) references users(id) on delete cascade,
    constraint user_2 foreign key (friend_id) references users(id) on delete cascade
    
);