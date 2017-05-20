package models

type Member struct {
	MemberID  int64  `json:"member_id"`
	UserID    int64  `json:"user_id"`
    ImageID   *int64 `json:"image_id"`
	AddressID *int64 `json:"address_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

    User      *User   `json:"user"`
}
