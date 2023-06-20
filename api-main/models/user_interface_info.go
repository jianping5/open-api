package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type InterfaceInfo struct {
	Id           	primitive.ObjectID  `bson:"_id,omitempty"`
	UserId    		primitive.ObjectID  `bson:"user_id,omitempty"`
	InterfaceInfoId primitive.ObjectID  `bson:"interface_info_id,omitempty"`
	TotalNum  		uint  				`bson:"total_num,omitempty"`
	LeftNum     	uint  				`bson:"left_num,omitempty"`
	Status 			bool  				`bson:"status,omitempty"`
	IsDelete     	bool    			`bson:"is_delete,omitempty"`
}