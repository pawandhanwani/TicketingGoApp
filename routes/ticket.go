package routes

import (
	"encoding/json"
	"io/ioutil"

	"../models"

	"../helpers"
	"github.com/gin-gonic/gin"
)

type CreateTicketBody struct {
	Title       string
	Description string
}

func CreateTicket(ctx *gin.Context) {
	token := ctx.GetHeader("authToken")
	userId, err := helpers.ValidateJwt(token)

	if err != nil || userId <= 0 {
		ctx.JSON(403, gin.H{"msg": "Un-Authorised"})
	} else {
		var payload CreateTicketBody
		jsonBody, _ := ioutil.ReadAll(ctx.Request.Body)
		json.Unmarshal(jsonBody, &payload)
		ticket := models.Ticket{CreatedBy: userId, Title: payload.Title, Description: payload.Description, Status: "Unassigned"}
		result := models.DB.Create(&ticket)
		if result.RowsAffected > 0 {
			ctx.JSON(200, gin.H{"msg": "Ticket Created", "TicketData": ticket})
		} else {
			ctx.JSON(200, gin.H{"msg": "Ticket Created"})
		}
	}
}

func ListTickets(ctx *gin.Context) {
	token := ctx.GetHeader("authToken")
	userId, err := helpers.ValidateJwt(token)
	if err != nil || userId <= 0 {
		ctx.JSON(403, gin.H{"msg": "Un-Authorised"})
	} else {
		var tickets []models.Ticket
		models.DB.Where("created_by = ?", userId).Find(&tickets)
		ctx.JSON(200, gin.H{"data": tickets})
	}
}

func GetTicket(ctx *gin.Context) {
	token := ctx.GetHeader("authToken")
	ticketid := ctx.Param("ticketid")
	userId, err := helpers.ValidateJwt(token)

	if err != nil || userId <= 0 {
		ctx.JSON(403, gin.H{"msg": "Un-Authorised"})
	} else {
		var ticket models.Ticket
		models.DB.Preload("Comment").Where("created_by = ? AND id = ?", userId, ticketid).Find(&ticket)
		ctx.JSON(200, gin.H{"data": ticket})
	}
}

func DeleteTicket(ctx *gin.Context) {
	token := ctx.GetHeader("authToken")
	ticketid := ctx.Param("ticketid")
	userId, err := helpers.ValidateJwt(token)

	if err != nil || userId <= 0 {
		ctx.JSON(403, gin.H{"msg": "Un-Authorised"})
	} else {
		result := models.DB.Where("created_by = ? AND id = ?", userId, ticketid).Delete(&models.Ticket{})
		if result.RowsAffected > 0 {
			ctx.JSON(200, gin.H{"msg": "Deleted"})
		} else {
			ctx.JSON(200, gin.H{"msg": "Unable to delete"})
		}
	}
}
