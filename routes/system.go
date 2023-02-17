package routes

import (
	"../helpers"

	"github.com/gin-gonic/gin"
)

func AllUserTypes(ctx *gin.Context) {

	ctx.JSON(200, gin.H{"data": gin.H{"user_types": helpers.AllUserTypes()}})
}

func AllTicketStatusTypes(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"data": gin.H{"ticket_status_types": helpers.TicketStatuses()}})
}
