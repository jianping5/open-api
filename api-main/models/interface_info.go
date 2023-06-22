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
	Status			uint 				`bson:"status,omitemmpty"`
	Method    		string  			`bson:"method,omitempty"`
	IsDelete     	uint    			`bson:"is_delete,omitempty"`
}

type InterfaceInfoAddRequest struct {
	Description     string  			`json:"description"`
	Method     		string  			`json:"method"`
	Name     		string  			`json:"name"`
	RequestParams   string  			`json:"requestParams"`
	RequestHeader   string  			`json:"requestHeader"`
	ResponseHeader  string  			`json:"responseHeader"`
	Url     		string  			`json:"url"`
}

type InterfaceInfoUpdateRequest struct {
	Id              string              `json:"id"`
	Description     string  			`json:"description"`
	Method     		string  			`json:"method"`
	Name     		string  			`json:"name"`
	RequestParams   string  			`json:"requestParams"`
	RequestHeader   string  			`json:"requestHeader"`
	ResponseHeader  string  			`json:"responseHeader"`
	Url     		string  			`json:"url"`
	Status          uint                `json:"status"`
}

type DeleteRequest struct {
	Id              string              `json:"id"`
}

type IdRequest struct {
	Id              string              `json:"id"`
}

type InterfaceInfoInvokeRequest struct {
	Id              	string			`json:"id"`
	UserRequestParams   string			`json:"userRequestParams"`
}

type InterfaceInfoQueryRequest struct {
	Id              string              `json:"id"`
	UserId          string              `json:"userId"`
	Description     string  			`json:"description"`
	Method     		string  			`json:"method"`
	Name     		string  			`json:"name"`
	RequestHeader   string  			`json:"requestHeader"`
	ResponseHeader  string  			`json:"responseHeader"`
	Url     		string  			`json:"url"`
	Status          uint                `json:"status"`
	Current         uint                `json:"current"`
	PageSize        uint                `json:"pageSize"`
	SortField       uint                `json:"sortField"`
	SortOrder       uint                `json:"sortOrder"`
}

type InterfaceInfoVO struct {
	Id              string              `json:"id"`
	UserId          string              `json:"userId"`
	Description     string  			`json:"description"`
	Method     		string  			`json:"method"`
	Name     		string  			`json:"name"`
	RequestParams   string  			`json:"requestParams"`
	RequestHeader   string  			`json:"requestHeader"`
	ResponseHeader  string  			`json:"responseHeader"`
	Url     		string  			`json:"url"`
	Status          uint                `json:"status"`
	TotalNum        uint                `json:"totalNum"`
	LeftNum         uint                `json:"leftNum"`
	IsDelete        uint                `json:"isDelete"`
}