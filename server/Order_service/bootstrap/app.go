package bootstrap

import "gorm.io/gorm"

type Application struct {
	DB *gorm.DB
}

func App() Application {
	Envinitializer()
	app := &Application{}

	app.DB = CreatedatabaseInstance()
	return *app
}
