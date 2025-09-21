package types

type Role string

const (
	User   Role = "admin"
	Editor Role = "editor"
	Admin  Role = "admin"
)

type Level int

const (
	LevelUser   Level = 1
	LevelEditor Level = 5
	LevelAdmin  Level = 10
)

type Permission string

const (
	ViewPost    Permission = "view_post"
	EditPost    Permission = "edit_post"
	ManageUsers Permission = "manage_users"
)
