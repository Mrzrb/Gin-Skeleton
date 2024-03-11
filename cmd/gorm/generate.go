package main

import (
	"app/app/model/dao"
	"app/config"
	"app/helpers"

	"gorm.io/gen"
)

func applyModel() []any {
	m := []any{
		&dao.Demo{},
		&dao.User{},
	}
	return m
}

func applyInterface(g *gen.Generator) {
	// g.ApplyInterface(func(repo.ITest) {}, schemas.Test{})
}

func main() {
	config.InitConf("dev")
	helpers.InitMysql()
	// ...
	generate()
}

func generate() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./app/model/query/",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	db := helpers.MysqlClients[helpers.CommonMysqlClient]
	g.UseDB(db)
	db.AutoMigrate(&dao.Demo{})

	g.ApplyBasic(applyModel()...)
	applyInterface(g)

	g.Execute()
}
