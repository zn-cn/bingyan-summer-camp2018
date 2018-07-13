package class
type RbacRoleUser struct {
	RoleId int
	UserId int `orm:"pk"`
}