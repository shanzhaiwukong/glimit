## glimit

> One database limit desigin for mysql

### Example
``` golang
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