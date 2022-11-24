package controllers

import "myGoProjectNew/models"

type Rsp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type LoginInfo struct {
	User  *models.User
	Token string
}

type GetUserListParams struct {
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
	Username  string `json:"username"`
}
