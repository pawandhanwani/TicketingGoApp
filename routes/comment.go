package routes

import (
	"encoding/json"
	"io/ioutil"

	"../models"

	"../helpers"
	"github.com/gin-gonic/gin"
)

type CreateCommentBody struct {
	Comment  string
	TickedId uint
}

func AddComment(ctx *gin.Context) {
	token := ctx.GetHeader("authToken")
	userId, err := helpers.ValidateJwt(token)
	if err != nil || userId <= 0 {
		ctx.JSON(403, gin.H{"msg": "Un-Authorised"})
	} else {
		var payload CreateCommentBody
		jsonBody, _ := ioutil.ReadAll(ctx.Request.Body)
		json.Unmarshal(jsonBody, &payload)
		comment := models.Comment{TicketId: payload.TickedId, CreatedBy: userId, Comment: payload.Comment}
		result := models.DB.Create(&comment)
		if result.RowsAffected > 0 {
			ctx.JSON(201, gin.H{"msg": "Comment Posted!"})
		} else {
			ctx.JSON(500, gin.H{"msg": "Unable to post comment!"})
		}
	}
}

func DeleteComment(ctx *gin.Context) {
	token := ctx.GetHeader("authToken")
	userId, err := helpers.ValidateJwt(token)
	commentid := ctx.Param("commentid")
	if err != nil || userId <= 0 {
		ctx.JSON(403, gin.H{"msg": "Un-Authorised"})
	} else {
		result := models.DB.Where("created_by = ? AND id = ?", userId, commentid).Delete(&models.Comment{})
		if result.RowsAffected > 0 {
			ctx.JSON(200, gin.H{"msg": "Deleted"})
		} else {
			ctx.JSON(200, gin.H{"msg": "Unable to delete"})
		}
	}
}
