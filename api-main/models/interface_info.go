package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type InterfaceInfo struct {
	Id           	primitive.ObjectID  `bson:"_id,omitempty"`
	UserId    		primitive.ObjectID  `bson:"user_id,omitempty"`
	Name  			string  			`bson:"name,omitempty"`
	Description     string  			`bson:"description,omitempty"`
	Url 			string  			`bson:"url,omitempty"`
	RequestParams   string  			`bson:"request_params,omitempty"`
	RequestHeader   string  			`bson:"request_header,omitempty"`
	ResponseHeader  string    			`bson:"response_header,omitempty"`
	Status			bool 				`bson:"status,omitemmpty"`
	Method    		string  			`bson:"method,omitempty"`
	IsDelete     	bool    			`bson:"is_delete,omitempty"`
}