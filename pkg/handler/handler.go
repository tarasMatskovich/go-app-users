package handler;

import (
    "github.com/gin-gonic/gin"
)

type Handler struct {

}

func (h *Handler) InitRoutes() *gin.Engine {
    router := gin.New()

    router.GET("/users", h.getUsers)

    return router
}
