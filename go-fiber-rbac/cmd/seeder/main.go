package main

import (
	"go-fiber-rbac/config"
	"go-fiber-rbac/db"
	"go-fiber-rbac/models"
	"go-fiber-rbac/types"
	"go-fiber-rbac/util"

	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	cfg := config.New(log)

	db := db.Connect(cfg, logrus.New())

	adminRole := models.Role{Name: "admin", Level: int(types.Admin)}
	editorRole := models.Role{Name: "editor", Level: int(types.Editor)}
	userRole := models.Role{Name: "user", Level: int(types.User)}

	viewPost := models.Permission{Name: string(types.ViewPost)}
	editPost := models.Permission{Name: string(types.EditPost)}
	manageUsers := models.Permission{Name: string(types.ManageUsers)}

	db.FirstOrCreate(&adminRole, adminRole)
	db.FirstOrCreate(&editorRole, editorRole)
	db.FirstOrCreate(&userRole, userRole)

	db.FirstOrCreate(&viewPost, viewPost)
	db.FirstOrCreate(&editPost, editPost)
	db.FirstOrCreate(&manageUsers, manageUsers)

	db.Model(&adminRole).Association("Permissions").Append(&viewPost, &editPost, &manageUsers)
	db.Model(&editorRole).Association("Permissions").Append(&viewPost, &editPost)
	db.Model(&userRole).Association("Permissions").Append(&viewPost)

	hash, _ := util.HashPassword("password")

	admin := models.User{Username: "admin", Password: hash, RoleID: adminRole.ID}
	editor := models.User{Username: "editor", Password: hash, RoleID: editorRole.ID}
	user := models.User{Username: "user", Password: hash, RoleID: userRole.ID}

	db.FirstOrCreate(&admin, admin)
	db.FirstOrCreate(&editor, editor)
	db.FirstOrCreate(&user, user)
}
