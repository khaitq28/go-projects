package appcontext

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
	"todo-app/internal/user"
)

type AppContext struct {
	UserService user.UserService
}

var (
	appCtx *AppContext
	once   sync.Once
)

func NewAppContext() *AppContext {
	once.Do(func() {
		dsn := "host=localhost user=todo_user password=todo_pass dbname=todo_db port=5432 sslmode=disable"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("failed to connect to db:", err)
		}

		userRepo := user.NewRepository(db)

		appCtx = &AppContext{
			UserService: user.NewUserService(userRepo),
		}
	})

	return appCtx
}
