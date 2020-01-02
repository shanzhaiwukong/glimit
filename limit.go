package glimit

import (
	"strings"
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
		ID        int32   `xorm:"id" json:"pk autoincr id"`
		RoleID    int32   `json:"role_id" xorm:"role_id default 0"`
		ActionIDs []int32 `json:"action_ids" xorm:"action_ids"`
	}
)

type (
	// VUserRole 用户角色视图
	VUserRole struct {
		*TabUser `json:"user" xorm:"extends"`
		*TabRole `json:"role" xorm:"extends"`
	}
	// VUserRoleMenu 用户角色菜单视图
	VUserRoleMenu struct {
		*TabUser `json:"user" xorm:"extends" `
		*TabRole `json:"role" xorm:"extends"`
		Menus    []*TabMenu `json:"menus"`
	}
	// VUserRoleAction 用户角色行为视图
	VUserRoleAction struct {
		*TabUser `json:"user" xorm:"extends"`
		*TabRole `json:"role" xorm:"extends"`
		Actions  []*TabAction `json:"actions"`
	}
	//VUserRoleMenuAction 用户角色菜单行为视图
	VUserRoleMenuAction struct {
		*TabUser `json:"user" xorm:"extends"`
		*TabRole `json:"role" xorm:"extends"`
		Menus    []*TabMenu   `json:"menus"`
		Actions  []*TabAction `json:"actions"`
	}
)

// Sync2Default 同步数据库(创建)
func Sync2Default(db *xorm.Engine) error {
	return db.Sync2(&TabRole{}, &TabUser{}, &TabMenu{}, &TabAction{}, &TabRoleMenu{}, &TabRoleAction{})
}

// GetUserByID 根据id获取用户详情
func GetUserByID(db *xorm.Engine, userID int32) (user *TabUser) {
	user = &TabUser{ID: userID}
	if b, err := db.Get(user); !b || err != nil {
		user = nil
	}
	return
}

// GetUserByAccountPwd 根据帐号密码获取用户详情
func GetUserByAccountPwd(db *xorm.Engine, account, pwd string) (user *TabUser) {
	user = &TabUser{Account: account, Pwd: pwd}
	if b, err := db.Get(user); !b || err != nil {
		user = nil
	}
	return
}

// GetUserByEmailPwd 根据邮箱密码获取用户详情
func GetUserByEmailPwd(db *xorm.Engine, email, pwd string) (user *TabUser) {
	user = &TabUser{Email: email, Pwd: pwd}
	if b, err := db.Get(user); !b || err != nil {
		user = nil
	}
	return
}

// GetUserByPhonePwd 根据手机号密码获取用户详情
func GetUserByPhonePwd(db *xorm.Engine, phone, pwd string) (user *TabUser) {
	user = &TabUser{Phone: phone, Pwd: pwd}
	if b, err := db.Get(user); !b || err != nil {
		user = nil
	}
	return
}

// GetRoleByID 根据id获取角色详情
func GetRoleByID(db *xorm.Engine, roleID int32) (role *TabRole) {
	role = &TabRole{ID: roleID}
	if b, err := db.Get(role); !b || err != nil {
		role = nil
	}
	return
}

// GetRoleAll 获取所有角色
func GetRoleAll(db *xorm.Engine) (roles []*TabRole, err error) {
	err = db.Find(&roles)
	return
}

// GetMenuByID 根据id获取菜单详情
func GetMenuByID(db *xorm.Engine, menuID int32) (menu *TabMenu) {
	menu = &TabMenu{ID: menuID}
	if b, err := db.Get(menu); !b || err != nil {
		menu = nil
	}
	return
}

// GetMenuAll 获取所有菜单
func GetMenuAll(db *xorm.Engine) (menus []*TabMenu, err error) {
	err = db.Find(&menus)
	return
}

// GetActionByID 根据id获取行为详情
func GetActionByID(db *xorm.Engine, actionID int32) (action *TabAction) {
	action = &TabAction{ID: actionID}
	if b, err := db.Get(action); !b || err != nil {
		action = nil
	}
	return
}

// GetActionAll 获取所有行为
func GetActionAll(db *xorm.Engine) (actions []*TabAction, err error) {
	err = db.Find(&actions)
	return
}

// GetRoleMenus 获取角色菜单
func GetRoleMenus(db *xorm.Engine, roleID int32) (menus []*TabMenu, err error) {
	err = db.SQL(`select * from tab_menu where FIND_IN_SET(id,(select replace(replace((select menu_ids from tab_role_menu where role_id=? limit 1),'[',''),']','')))`, roleID).Find(&menus)
	return
}

// GetRoleActions 获取角色行为
func GetRoleActions(db *xorm.Engine, roleID int32) (actions []*TabAction, err error) {
	err = db.SQL(`select * from tab_action where FIND_IN_SET(id,(select replace(replace((select action_ids from tab_role_action where role_id=? limit 1),'[',''),']','')))`, roleID).Find(&actions)
	return
}

// CheckURLInActions 检查url是否在行为集合中
func CheckURLInActions(url string, actions []*TabAction) bool {
	for _, a := range actions {
		if a != nil && strings.ToLower(a.Href) == strings.ToLower(url) {
			return true
		}
	}
	return false
}

// GetVUserRoleList 获取用户角色列表详情
func GetVUserRoleList(db *xorm.Engine) (list []*VUserRole, err error) {
	err = db.SQL("select a.*,b.* from tab_user a left join tab_role b on a.role_id=b.id order by a.id desc").Find(&list)
	return
}

// GetVUserRoleByUID 根据用户ID获取用户角色详情
func GetVUserRoleByUID(db *xorm.Engine, uid int32) (one *VUserRole, err error) {
	list := make([]*VUserRole, 0)
	err = db.SQL("select a.*,b.* from tab_user a left join tab_role b on a.role_id=b.id where a.id=?", uid).Find(&list)
	if len(list) > 0 {
		one = list[0]
	}
	return
}

// GetVUserRoleMenuByUID 根据用户ID获取用户角色及菜单列表
func GetVUserRoleMenuByUID(db *xorm.Engine, uid int32) (*VUserRoleMenu, error) {
	res := &VUserRoleMenu{}
	if ur, err := GetVUserRoleByUID(db, uid); err == nil && ur != nil {
		res.TabUser = ur.TabUser
		res.TabRole = ur.TabRole
		menus, _ := GetRoleMenus(db, ur.TabRole.ID)
		res.Menus = menus
		return res, nil
	} else {
		return nil, err
	}
}

// GetVUserRoleActionByUID 根据用户ID获取用户角色及行为列表
func GetVUserRoleActionByUID(db *xorm.Engine, uid int32) (*VUserRoleAction, error) {
	res := &VUserRoleAction{}
	if ur, err := GetVUserRoleByUID(db, uid); err == nil && ur != nil {
		res.TabUser = ur.TabUser
		res.TabRole = ur.TabRole
		actions, _ := GetRoleActions(db, ur.TabRole.ID)
		res.Actions = actions
		return res, nil
	} else {
		return nil, err
	}
}

// GetVUserRoleMenuActionByUID 根据用户ID获取用户角色菜单及行为列表
func GetVUserRoleMenuActionByUID(db *xorm.Engine, uid int32) (*VUserRoleMenuAction, error) {
	res := &VUserRoleMenuAction{}
	if ur, err := GetVUserRoleByUID(db, uid); err == nil && ur != nil {
		res.TabUser = ur.TabUser
		res.TabRole = ur.TabRole
		menus, _ := GetRoleMenus(db, ur.TabRole.ID)
		res.Menus = menus
		actions, _ := GetRoleActions(db, ur.TabRole.ID)
		res.Actions = actions
		return res, nil
	} else {
		return nil, err
	}
}
