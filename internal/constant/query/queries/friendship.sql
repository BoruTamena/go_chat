

-- name: GetFiendByUserName :one 
select * from users where email = $1 ;


-- name: GetFriendList :many 

select users.id,user_name,email,password from users join friend_lists
on users.id=friend_lists.user_id ;


-- name: AddFriend :one 
insert into friend_lists(
    user_id,
    friend_id,
    status
) values ( $1,$2,$3) returning *;


-- name: UpdateFriendStatus :one 
update friend_lists 
set status=$3 where 
user_id=$1 and friend_id=$2 
returning *;
