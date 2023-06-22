package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	UserAccount  string  `bson:"user_account,omitempty"`
	UserRole     string  `bson:"user_role,omitempty"`
	UserPassword string  `bson:"user_password,omitempty"`
	AccessKey    string  `bson:"access_key,omitempty"`
	SecretKey    string  `bson:"secret_key,omitempty"`
	IsDelete     uint    `bson:"is_delete"`
}

type UserRegister struct {
	UserAccount 	string  `json:"userAccount"`
	UserPassword 	string  `json:"userPassword"`
	CheckPassword 	string  `json:"checkPassword"`
}

type UserLogin struct {
	UserAccount 	string  `json:"userAccount"`
	UserPassword 	string  `json:"userPassword"`
}