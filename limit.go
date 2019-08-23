package glimit

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	// TabRole 角色
	TabRole struct {
		ID   bson.ObjectId `bson:'_id',json:'id'`
		Name string        `json:"name"`
		Tip  string        `json:"tip"`
	}
	// TabUser 用户
	TabUser struct {
		ID         bson.ObjectId `bson:"_id",json:"id"`
		RoleID     string        `json:"role_id"`
		Account    string        `json:"account"`
		Email      string        `json:"email"`
		Phone      string        `json:"phone"`
		Pwd        string        `json:"pwd"`
		Nick       string        `json:"nick"`
		Avatar     string        `json:"avatar"`
		Birthday   int64         `json:"birthday"`
		Status     int           `json:"status"`
		IsDelete   bool          `json:"is_delete"`
		RegistTime time.Time     `json:"regist_time"`
		ActiveTime time.Time     `json:"active_time"`
	}
	// TabMenu 菜单
	TabMenu struct {
		ID       bson.ObjectId `bson:"_id",json:"id"`
		ParentID string        `json:"parent_id"`
		Name     string        `json:"name"`
		Href     string        `json:"href"`
		Icon     string        `json:"icon"`
		Sort     int           `json:"sort"`
	}
	// TabAction 行为
	TabAction struct {
		ID     bson.ObjectId `bson:"_id",json:"id"`
		MenuID string        `json:"menu_id"`
		Name   string        `json:"name"`
		Href   string        `json:"href"`
	}
	// TabRoleMenu 角色菜单
	TabRoleMenu struct {
		ID      bson.ObjectId `bson:"_id",json:"id"`
		RoleID  string        `json:"role_id"`
		MenuIDs []string      `json:"menu_ids"`
	}
	// TabRoleAction 角色行为
	TabRoleAction struct {
		ID        bson.ObjectId `bson:"_id",json:"id"`
		RoleID    string        `json:"role_id"`
		ActionIDs []string      `json:"action_ids"`
	}
)
