package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type UserInterfaceInfo struct {
	Id           	primitive.ObjectID  `bson:"_id,omitempty"`
	UserId    		primitive.ObjectID  `bson:"user_id,omitempty"`
	InterfaceInfoId primitive.ObjectID  `bson:"interface_info_id,omitempty"`
	TotalNum  		uint  				`bson:"total_num,omitempty"`
	LeftNum     	uint  				`bson:"left_num,omitempty"`
	Status 			uint  				`bson:"status,omitempty"`
	IsDelete     	uint    			`bson:"is_delete,omitempty"`
}