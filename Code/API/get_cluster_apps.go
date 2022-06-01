package Controller

// GetClusterApps returns a list of applications that have
// been uploaded to the system
func GetClusterApps(c *gin.Context) {
	// Read from working directory
	apps, err := ioutil.ReadDir("/usr/yakomaster/")
	if err != nil {
		log.Println("Could not read from working directory /usr/yakomaster")
	}
	var appsNames []string
	for _, app := range apps {
		appsNames = append(appsNames, app.Name())
	}

	c.JSON(http.StatusOK, appsNames)
}