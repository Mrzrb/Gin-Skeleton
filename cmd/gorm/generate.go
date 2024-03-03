package main

import (
	"app/app/model/dao"
	"app/helpers"

	"gorm.io/gen"
)

func applyModel() []any {
	m := []any{
		dao.Demo{},
	}
	return m
}

func applyInterface(g *gen.Generator) {
	// g.ApplyInterface(func(repo.ITest) {}, schemas.Test{})
}

func main() {
	helpers.InitMysql()
	// ...
	generate()
}

func generate() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./app/model/query/",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(helpers.MysqlClients[helpers.CommonMysqlClient])

	g.ApplyBasic(applyModel()...)
	applyInterface(g)

	g.Execute()
}
