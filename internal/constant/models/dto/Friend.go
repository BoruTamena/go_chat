package dto

const (
	Pending  string = "pending"
	Accepted string = "accepted"
	Blocked  string = "blocked"
)

type FriendRequest struct {
	SenderId string `json:"sender_id"`
	FriendId string `json:"friend_id"`
	Status   string `json:"status"`
}

type FriendUpdate struct {
	UserName string `json:"user_name"`
	Status   string `json:"status"`
}
