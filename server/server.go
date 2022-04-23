package server

func Start() {
	router := setRouter()

	router.Run("localhost:8080")
}
