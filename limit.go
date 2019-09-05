package glimit

import (
	"time"

	"github.com/go-xorm/xorm"
)

type (
	// TabRole 角色
	TabRole struct {
		ID   int32  `json:"id" xorm:"pk autoincr id" `
		Name string `json:"name" xorm:"varchar(50) default ''"`
		Tip  string `json:"tip" xorm:"varchar(200) default ''"`
	}

	// TabUser 用户
	TabUser struct {
		ID         int32     `json:"id" xorm:"pk autoincr id"`
		RoleID     int32     `json:"role_id" xorm:"int role_id default 0"`
		Account    string    `json:"account" xorm:"varchar(50) account default ''"`
		Email      string    `json:"email" xorm:"varchar(50) default ''"`
		Phone      string    `json:"phone" xorm:"varchar(50) default ''"`
		Pwd        string    `json:"pwd" xorm:"varchar(50) default ''"`
		Nick       string    `json:"nick" xorm:"varchar(50) default ''"`
		Avatar     string    `json:"avatar" xorm:"varchar(100) default ''"`
		Birthday   int64     `json:"birthday" xorm:"default 0"`
		Status     int       `json:"status" xorm:"default 0"`
		IsDelete   bool      `json:"is_delete" xorm:"is_delete default 0"`
		RegistTime time.Time `json:"regist_time" xorm:"timestamp regist_time default CURRENT_TIMESTAMP"`
		ActiveTime time.Time `json:"active_time" xorm:"timestamp active_time default CURRENT_TIMESTAMP"`
	}

	// TabMenu 菜单
	TabMenu struct {
		ID       int32  `json:"id" xorm:"pk autoincr id"`
		ParentID int32  `json:"parent_id" xorm:"int parent_id default 0"`
		Name     string `json:"name" xorm:"varchar(50) default ''"`
		Href     string `json:"href" xorm:"varchar(100) default ''"`
		Icon     string `json:"icon" xorm:"varchar(50) default ''"`
		Sort     int    `json:"sort" xorm:"default 0"`
	}

	// TabAction 行为
	TabAction struct {
		ID     int32  `json:"id" xorm:"pk autoincr id"`
		MenuID int32  `json:"menu_id" xorm:"menu_id default 0"`
		Name   string `json:"name" xorm:"varchar(50) default ''"`
		Href   string `json:"href" xorm:"varchar(100) default ''"`
	}

	// TabRoleMenu 角色菜单
	TabRoleMenu struct {
		ID      int32   `json:"id" xorm:"pk autoincr id"`
		RoleID  int32   `json:"role_id" xorm:"role_id default 0"`
		MenuIDs []int32 `json:"menu_ids" xorm:"menu_ids"`
	}

	// TabRoleAction 角色行为
	TabRoleAction struct {
		ID        int32   `xorm:"id" json:"id"`
		RoleID    int32   `json:"role_id" xorm:"role_id default 0"`
		ActionIDs []int32 `json:"action_ids" xorm:"action_ids"`
	}
)

// GetRoles 获取角色列表
func GetRoles(db *xorm.Engine) (roles []*TabRole, err error) {
	err = db.Find(&roles)
	return
}
