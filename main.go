package main

import (
	"./models"
	"./routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDB() // new

	auth := r.Group("/auth")
	{
		auth.POST("/signup", routes.Signup)
		auth.POST("/login", routes.Login)
	}

	ticket := r.Group("/ticket")
	{
		ticket.POST("/create", routes.CreateTicket)
		ticket.GET("/list", routes.ListTickets)
		ticket.GET("/:ticketid", routes.GetTicket)
		ticket.DELETE("/:ticketid", routes.DeleteTicket)
	}

	comment := r.Group("/comment")
	{
		comment.POST("/create", routes.AddComment)
		comment.DELETE("/:commentid", routes.DeleteComment)
	}

	system := r.Group("/system")
	{
		system.GET("/user/types", routes.AllUserTypes)
		system.GET("/ticket/status/types", routes.AllTicketStatusTypes)
	}

	admin := r.Group("/admin")
	{
		admin.PUT("/assign/ticket", routes.AssignTicket)
		admin.PATCH("/update/user/type", routes.UpdateUserType)
	}

	moderator := r.Group("/moderator")
	{
		moderator.PATCH("/update/ticket/status", routes.UpdateTicketStatus)
	}

	r.Run()

}
