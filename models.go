package authority

// Permission represents the database model of permissions
type Permission struct {
	ID   uint
	Name string
}

// TableName sets the table name
func (p Permission) TableName() string {
	return tablePrefix + "permissions"
}

// RolePermission stores the relationship between roles and permissions
type RolePermission struct {
	ID           uint
	RoleID       uint
	PermissionID uint
}

// TableName sets the table name
func (r RolePermission) TableName() string {
	return tablePrefix + "role_permissions"
}

// Role represents the database model of roles
type Role struct {
	ID   uint
	Name string
}

// TableName sets the table name
func (r Role) TableName() string {
	return tablePrefix + "roles"
}

// UserRole represents the relationship between users and roles
type UserRole struct {
	ID     uint
	UserID uint
	RoleID uint
}

// TableName sets the table name
func (u UserRole) TableName() string {
	return tablePrefix + "user_roles"
}
