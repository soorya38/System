package models

type Profile struct {
	ID           int    `bson:"id" json:"id"`
	Name         string `bson:"name" json:"name"`
	UserName     string `bson:"username" json:"username"`
	Password     string `bson:"-" json:"-"`
	ProfileImage string `bson:"profileimage" json:"profileimage"`
	Email        string `bson:"email" json:"email"`
	Phone        string `bson:"phone" json:"phone"`
	Bio          string `bson:"bio" json:"bio"`
	Role         string `bson:"role" json:"role"`
}
