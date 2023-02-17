package routes

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"

	"../helpers"
	"../models"
	"github.com/gin-gonic/gin"
)

type SignUpBody struct {
	Name     string
	Password string
	Email    string
}

type LoginBody struct {
	Email    string
	Password string
}

func Signup(ctx *gin.Context) {
	var payload SignUpBody
	jsonBody, _ := ioutil.ReadAll(ctx.Request.Body)
	json.Unmarshal(jsonBody, &payload)
	pHash := md5.Sum([]byte(payload.Password))
	pToken := hex.EncodeToString(pHash[:])
	user := models.User{Email: payload.Email, Name: payload.Name, Password: pToken, UserType: helpers.USER_TYPE_NORMAL}
	result := models.DB.Create(&user)
	token, _ := helpers.GenerateJwt(user.ID)
	if result.RowsAffected > 0 {
		ctx.JSON(200, gin.H{"data": gin.H{"token": token}, "msg": "Signup Successfull"})
	} else {
		ctx.JSON(200, gin.H{"msg": "Signup Failed"})
	}
}

func Login(ctx *gin.Context) {
	var payload LoginBody
	jsonBody, _ := ioutil.ReadAll(ctx.Request.Body)
	json.Unmarshal(jsonBody, &payload)
	var user models.User
	pHash := md5.Sum([]byte(payload.Password))
	pToken := hex.EncodeToString(pHash[:])
	models.DB.Select("email,id").Where("email = ? AND password = ?", payload.Email, pToken).Find(&user)
	if user.ID > 0 {
		token, _ := helpers.GenerateJwt(user.ID)
		ctx.JSON(200, gin.H{"data": gin.H{"token": token}})
	} else {
		ctx.JSON(403, gin.H{"data": gin.H{"msg": "Login Failed"}})
	}

}
