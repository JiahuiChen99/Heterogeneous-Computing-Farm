package Controller

// Cluster returns the cluster schema
func Cluster(c *gin.Context) {
	var response model.Response
	// Get master & service registries and compose response
	response.YakoMasters = zookeeper.MasterRegistry
	response.YakoAgents = zookeeper.ServicesRegistry
	clusterSchema, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, string(clusterSchema))
}