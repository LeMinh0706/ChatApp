package server

import "github.com/gin-gonic/gin"

func Static(r *gin.Engine) {
	r.Static("/api/avatar", "./avatar")
}
