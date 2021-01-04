package app

import ("../controllers/ping"
"../controllers/users")

func mapUrls() {
	router.GET("/ping", ping.Ping)
	
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/internal/users/search", users.Search)
	router.POST("/users", users.CreateUser)
	router.PUT("users/:user_id", users.UpdateUser)
	router.PATCH("users/:user_id", users.UpdateUser)
	router.DELETE("users/:user_id", users.DeleteUser)


}