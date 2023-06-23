package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type UserInterfaceInfo struct {
	Id           	primitive.ObjectID  `bson:"_id,omitempty"`
	UserId    		primitive.ObjectID  `bson:"user_id,omitempty"`
	InterfaceInfoId primitive.ObjectID  `bson:"interface_info_id,omitempty"`
	TotalNum  		uint  				`bson:"total_num"`
	LeftNum     	uint  				`bson:"left_num"`
}

type UserInterfaceInfoAddRequest struct {
	UserId    		string  	`json:"userId,omitempty"`
	InterfaceInfoId string  	`json:"interfaceInfoId,omitempty"`
	TotalNum  		uint  		`json:"totalNum"`
	LeftNum     	uint  		`json:"leftNum"`
}

type InterfaceInfoDTO struct {
	Name 			string  	`json:"name"`
	TotalNum  		uint  		`json:"totalNum"`
}

