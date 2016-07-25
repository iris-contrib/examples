package models

type (
	// User represents the structure of our resource
	User struct {
		Id     string `json:"id" bson:"id"  form:"id"`
		Name   string `json:"name" bson:"name"  form:"name"`
		Email  string `json:"email" bson:"email"  form:"email"`
		Pass   string `json:"pass" bson:"pass" form:"pass"`
		Birth  string `json:"birth" bson:"birth"  form:"birth"`
		Gender bool   `json:"gender" bson:"gender"  form:"gender"`
	}
)
