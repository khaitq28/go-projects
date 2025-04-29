package appcontext

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
	"todo-app/internal/task"
	"todo-app/internal/user"
)

type AppContext struct {
	UserService user.UserService
	TaskService task.TaskService
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
		taskRepo := task.NewRepository(db)

		appCtx = &AppContext{
			UserService: user.NewUserService(userRepo),
			TaskService: task.NewTaskService(taskRepo),
		}
	})

	return appCtx, initErr
}
