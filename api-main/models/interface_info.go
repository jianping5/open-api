package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type InterfaceInfo struct {
	Id             primitive.ObjectID `bson:"_id,omitempty"`
	UserId         primitive.ObjectID `bson:"user_id"`
	Name           string             `bson:"name,omitempty"`
	Description    string             `bson:"description,omitempty"`
	Url            string             `bson:"url,omitempty"`
	RequestParams  string             `bson:"request_params"`
	RequestHeader  string             `bson:"request_header,omitempty"`
	ResponseHeader string             `bson:"response_header,omitempty"`
	Status         int8               `bson:"status"`
	Method         string             `bson:"method,omitempty"`
	IsDelete       int8               `bson:"is_delete"`
}

// type InterfaceInfoDTO struct {
// 	Id             primitive.ObjectID `json:"id"`
// 	UserId         primitive.ObjectID `json:"userId"`
// 	Name           string             `json:"name"`
// 	Description    string             `json:"description"`
// 	Url            string             `json:"url"`
// 	RequestParams  string             `json:"requestParams"`
// 	RequestHeader  string             `json:"requestHeader"`
// 	ResponseHeader string             `json:"responseHeader"`
// 	Status         int8               `json:"status"`
// 	Method         string             `json:"method"`
// 	IsDelete       int8               `json:"isDelete"`
// }

type InterfaceInfoAddRequest struct {
	Description    string `json:"description"`
	Method         string `json:"method"`
	Name           string `json:"name"`
	RequestParams  string `json:"requestParams"`
	RequestHeader  string `json:"requestHeader"`
	ResponseHeader string `json:"responseHeader"`
	Url            string `json:"url"`
}

type InterfaceInfoUpdateRequest struct {
	Id             string `json:"id"`
	Description    string `json:"description"`
	Method         string `json:"method"`
	Name           string `json:"name"`
	RequestParams  string `json:"requestParams"`
	RequestHeader  string `json:"requestHeader"`
	ResponseHeader string `json:"responseHeader"`
	Url            string `json:"url"`
	Status         int8   `json:"status"`
}

type DeleteRequest struct {
	Id string `json:"id"`
}

type IdRequest struct {
	Id string `json:"id"`
}

type InterfaceInfoInvokeRequest struct {
	Id                string `json:"id"`
	UserRequestParams string `json:"userRequestParams"`
}

type InterfaceInfoQueryRequest struct {
	Id             string `json:"id"`
	UserId         string `json:"userId"`
	Description    string `json:"description"`
	Method         string `json:"method"`
	Name           string `json:"name"`
	RequestHeader  string `json:"requestHeader"`
	ResponseHeader string `json:"responseHeader"`
	Url            string `json:"url"`
	Status         int8   `json:"status"`
	Current        int64  `json:"current"`
	PageSize       int64  `json:"pageSize"`
	SortField      string `json:"sortField"`
	SortOrder      string `json:"sortOrder"`
}

type InterfaceInfoVO struct {
	Id             string `json:"id"`
	UserId         string `json:"userId"`
	Description    string `json:"description"`
	Method         string `json:"method"`
	Name           string `json:"name"`
	RequestParams  string `json:"requestParams"`
	RequestHeader  string `json:"requestHeader"`
	ResponseHeader string `json:"responseHeader"`
	Url            string `json:"url"`
	Status         int8   `json:"status"`
	TotalNum       int8   `json:"totalNum"`
	LeftNum        int8   `json:"leftNum"`
	IsDelete       int8   `json:"isDelete"`
}
