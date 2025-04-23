package appcontext

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
	"todo-app/internal/user"
)

type AppContext struct {
	UserService user.UserService
}

var (
	appCtx  *AppContext
	once    sync.Once
	initErr error
)

func NewAppContext() (*AppContext, error) {
	once.Do(func() {
		dsn := "host=localhost user=todo_user password=todo_pass dbname=todo_db port=5432 sslmode=disable"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			initErr = err
			return
		}

		userRepo := user.NewRepository(db)

		appCtx = &AppContext{
			UserService: user.NewUserService(userRepo),
		}
	})

	return appCtx, initErr
}
