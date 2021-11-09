package models

type (
	GetUserResponse struct {
		user User `json:"user"`
	}

	User struct {
		ID        string `bson:"_id"`
		Firstname string `json:"firstname" bson:"firstname"`
		Lastname  string `json:"lastname" bson:"lastname"`
		Email     string `json:"email" bson:"email"`
		Password  string `json:"password" bson:"hashedPassword"`
		Created   int64  `bson:"created"`
	}

	UserRequest struct {
		ID        string `bson:"_id"`
		Firstname string `json:"firstname" bson:"firstname"`
		Lastname  string `json:"lastname" bson:"lastname"`
		Email     string `json:"email" bson:"email"`
		Password  string `json:"password" bson:"hashedPassword"`
		Created   int64  `bson:"created"`
	}
)
