package routes

import (
	"encoding/json"
	"io/ioutil"

	"../models"

	"../helpers"
	"github.com/gin-gonic/gin"
)

type UpdateTicketStatusBody struct {
	TicketId     uint
	TicketStatus string
}

func UpdateTicketStatus(ctx *gin.Context) {
	token := ctx.GetHeader("authToken")
	userId, err := helpers.ValidateJwt(token)
	if err != nil || userId <= 0 {
		ctx.JSON(403, gin.H{"msg": "Un-Authorised"})
	} else {
		var payload UpdateTicketStatusBody
		jsonBody, _ := ioutil.ReadAll(ctx.Request.Body)
		json.Unmarshal(jsonBody, &payload)
		if helpers.IsValidTicketStatus(payload.TicketStatus) {
			result := models.DB.Model(&models.Ticket{}).Where("assigned_to = ? AND id = ? AND status = ?", userId, payload.TicketId, helpers.TICKET_STATUS_OPEN).Update("status", payload.TicketStatus)
			if result.RowsAffected > 0 {
				ctx.JSON(200, gin.H{"msg": "Status Updated"})
			} else {
				ctx.JSON(200, gin.H{"msg": "Failed to update status"})
			}
		} else {
			ctx.JSON(400, gin.H{"msg": "Invalid Request"})
		}
	}
}
