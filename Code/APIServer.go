func APIServer() {
	// Default gin router with default middleware:
	// logger & recovery
	router := gin.Default()

	// Add CORS support
	router.Use(cors.Default())

	// Attach routes to gin router
	API.AddRoutes(router)

	// TODO: Use environment variables or secrets managers like Hashicorp Vault
	err := router.Run(addr)
	if err != nil {
		// TODO: Use logger
		panic("API gin Server could not be started!")
	}
}