## glimit

> One database limit desigin for mysql

### Example
``` golang
var mysql *xorm.Engine
//删除(drop)已有数据并建表
func dropCreate() {
	_drop := []string{
		"drop table if exists tab_user;",
		"drop table if exists tab_role;",
		"drop table if exists tab_menu;",
		"drop table if exists tab_role_menu;",
		"drop table if exists tab_action;",
		"drop table if exists tab_role_action;",
	}
	sess := mysql.NewSession()
	sess.Begin()
	defer func() {
		if err := recover(); err != nil {
			sess.Rollback()
		}
		sess.Commit()
		sess.Close()
	}()
	for _, d := range _drop {
		_, err := sess.Exec(d)
		if err != nil {
			fmt.Println("drop error", err)
		}
	}
	if err := sess.Sync2(&TabUser{}, &TabRole{}, &TabMenu{}, &TabAction{}, &TabRoleMenu{}, &TabRoleAction{}); err != nil {
		fmt.Println("sync error", err)
	}
}
type dataDefault struct {
	Roles      []*TabRole       `json:"roles"`
	Users      []*TabUser       `json:"users"`
	Menus      []*TabMenu       `json:"menus"`
	Actions    []*TabAction     `json:"actions"`
	RoleMenu   []*TabRoleMenu   `json:"role_menu"`
	RoleAction []*TabRoleAction `json:"role_action"`
}
//初始化数据
func loadData() {
	v := new(dataDefault)
	if err := gutil.LoadJSON2Struct("./assist/dbinit.json", v); err == nil && v != nil {
		mysql.Insert(v.Roles)
		mysql.Insert(v.Users)
		mysql.Insert(v.Menus)
		mysql.Insert(v.Actions)
		mysql.Insert(v.RoleMenu)
		mysql.Insert(v.RoleAction)
	}
}
```