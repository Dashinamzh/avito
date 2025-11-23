package main

import (
	"database/sql"

	pr "github.com/Dashinamzh/avito/internal/repo/PR"
	"github.com/Dashinamzh/avito/internal/repo/team"
	"github.com/Dashinamzh/avito/internal/repo/user"
)

func main() {
	db, err := sql.Open("postgres", "connection_string")
	if err != nil {
		panic(err)
	}

	userRepo := user.NewUserRepo(db)
	teamRepo := team.NewTeamRepo(db)
	PrRepo := pr.NewPrRepo(db)
}
