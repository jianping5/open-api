package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	UserAccount  string  `bson:"user_account,omitempty"`
	UserRole     string  `bson:"user_role,omitempty"`
	UserPassword string  `bson:"user_password,omitempty"`
	AccessKey    string  `bson:"access_key,omitempty"`
	SecretKey    string  `bson:"secret_key,omitempty"`
	IsDelete     bool    `bson:"is_delete,omitempty"`
}

type UserRegister struct {
	UserAccount 	string  `json:"user_account"`
	UserPassword 	string  `json:"user_password"`
	CheckPassword 	string  `json:"user_password"`
}

type UserLogin struct {
	UserAccount 	string  `json:"user_account"`
	UserPassword 	string  `json:"user_password"`
	CheckPassword 	string  `json:"user_password"`
}