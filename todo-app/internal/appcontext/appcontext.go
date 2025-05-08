package appcontext

import (
	"gorm.io/gorm"
	"sync"
	"todo_app/internal/task"
	"todo_app/internal/user"
)
import "gorm.io/driver/mysql"

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
		//dsn := "host=localhost user=todo_user password=todo_pass dbname=todo_db port=5432 sslmode=disable"
		dsn := "todo_user:todo_pass@tcp(localhost:3306)/todo_db?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
