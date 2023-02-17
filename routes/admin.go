package routes

import (
	"encoding/json"
	"io/ioutil"

	"../helpers"
	"../models"
	"github.com/gin-gonic/gin"
)

type AssignTicketBody struct {
	TicketId uint
	AssignTo string
}

type UpdateUserTypeBody struct {
	Email    string
	UserType string
}

func AssignTicket(ctx *gin.Context) {
	token := ctx.GetHeader("authToken")
	userId, err := helpers.ValidateJwt(token)
	if err != nil || userId <= 0 {
		ctx.JSON(403, gin.H{"msg": "Un-Authorised"})
	} else {
		var user models.User
		models.DB.Where("id = ?", userId).First(&user)
		if user.UserType == helpers.USER_TYPE_ADMIN {
			var payload AssignTicketBody
			jsonBody, _ := ioutil.ReadAll(ctx.Request.Body)
			json.Unmarshal(jsonBody, &payload)
			var moderator models.User
			models.DB.Model(&models.User{}).Where("email = ? AND user_type = ?", payload.AssignTo, helpers.USER_TYPE_MODERATOR).First(&moderator)
			if moderator.ID > 0 {
				result := models.DB.Model(&models.Ticket{}).Where("assigned_to = ? AND id= ?", 0, payload.TicketId).Update("assigned_to", moderator.ID)
				if result.RowsAffected > 0 {
					ctx.JSON(200, gin.H{"msg": "Ticket Assigned"})
				} else {
					ctx.JSON(503, gin.H{"msg": "Failed to assign ticket"})
				}
			} else {
				ctx.JSON(503, gin.H{"msg": "Looks like you are trying to assign ticket to non-moderator user"})
			}
		} else {
			ctx.JSON(403, gin.H{"msg": "You are not an admin"})
		}
	}
}

func UpdateUserType(ctx *gin.Context) {
	token := ctx.GetHeader("authToken")
	userId, err := helpers.ValidateJwt(token)
	if err != nil || userId <= 0 {
		ctx.JSON(403, gin.H{"msg": "Un-Authorised"})
	} else {
		var user models.User
		models.DB.Where("id = ?", userId).First(&user)
		if user.UserType == helpers.USER_TYPE_ADMIN {
			var payload UpdateUserTypeBody
			jsonBody, _ := ioutil.ReadAll(ctx.Request.Body)
			json.Unmarshal(jsonBody, &payload)
			if helpers.IsValidUserType(payload.UserType) {

				result := models.DB.Model(&models.User{}).Where("email = ? AND user_type = ?", payload.Email, helpers.USER_TYPE_NORMAL).Update("user_type", payload.UserType)
				if result.RowsAffected > 0 {
					ctx.JSON(200, gin.H{"msg": "Update Successfull"})
				} else {
					ctx.JSON(500, gin.H{"msg": "Nothing Updated"})
				}
			} else {
				ctx.JSON(500, gin.H{"msg": "Invalid Type"})
			}
		} else {
			ctx.JSON(403, gin.H{"msg": "You are not an admin"})
		}
	}
}
