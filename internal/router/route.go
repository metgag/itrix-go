package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoute(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	InitTodoRoute(r, db)

	return r
}
