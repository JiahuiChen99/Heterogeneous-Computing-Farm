package Controller

// UploadApp handles the file that the user wants
// to deploy in the cluster
func UploadApp(c *gin.Context) {
	file, formErr := c.FormFile("app")
	if formErr != nil {
		err := utils.BadRequestError(formErr.Error())
		c.JSON(err.Status, err)
		return
	}

	// Check if YakoMaster's working directory is available
	directory_util.WorkDir("yakomaster")

	// Save the file on the server
	appPath := "/usr/yakomaster/" + file.Filename
	if saveErr := c.SaveUploadedFile(file, appPath); saveErr != nil {
		err := utils.InternalServerError(saveErr.Error())
		c.JSON(err.Status, err)
		return
	}

	// Get the app's resources configuration
	var config model.Config
	if err := c.ShouldBind(&config); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	var iot bool
	if c.Query("iot") == "true" {
		iot = true
	} else {
		iot = false
	}

	// Compute and find the best nodes to deploy the app
	recommendedNodes := findYakoAgents(config, iot)

	// Check if automation is enabled
	if autoDeploy := c.Query("autodeploy"); autoDeploy == "true" {
		// Auto-deploy the app to the best computed node
		yakoAgentID := recommendedNodes[0].ID
		log.Println("Autodeploying application to " + yakoAgentID)
		if !iot {
			// Deploy to a non-IoT agent
			deployStatus := deployApp(&zookeeper.ServicesRegistry[yakoAgentID]
				.GrpcClient, appPath, file.Filename)
			log.Println(deployStatus.Message)
		} else {
			// Deploy to the best IoT agent
			deployAppIoT(yakoAgentID, appPath, file.Filename)
		}
	}

	// File uploaded and stored
	c.JSON(http.StatusOK, recommendedNodes)
}