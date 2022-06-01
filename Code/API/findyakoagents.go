// findYakoAgents calculates and finds the top X best & suitable
// yakoagents where the app could be deployed according to the
// requested resources from the client
// Default X = 3
func findYakoAgents(config model.Config, iot bool) []*model.YakoAgent {
	// Priority queue with max heap to rank higher the nodes
	// with more brownie points
	agentsCount := 0
	for agentID, _ := range zookeeper.ServicesRegistry {
		if iot && string(agentID[0]) != "n" {
			agentsCount++
		} else if !iot && string(agentID[0]) == "n" {
			agentsCount++
		}
	}
	pq := make(model.PQNodes, agentsCount)

	var browniePoints uint64
	counter := 0
	// Loop through all the available yakoagents, computes the
	// brownie points and adds it to a priority queue
	for agentID, agentInfo := range zookeeper.ServicesRegistry {
		// Filter out the devices depending on the IoT field
		if iot && string(agentID[0]) != "n" {
			// Set brownie points to 0
			browniePoints = 0
			compliesWithCPUCores(agentInfo.ServiceInfo, config, &browniePoints)
			compliesWithMemory(agentInfo.ServiceInfo, config, &browniePoints)
			pq[counter] = &model.YakoAgent{
				ID:            agentID,
				BrowniePoints: browniePoints,
			}
			counter++
		} else if !iot && string(agentID[0]) == "n" {
			// Set brownie points to 0
			browniePoints = 0
			compliesWithCPUCores(agentInfo.ServiceInfo, config, &browniePoints)
			compliesWithMemory(agentInfo.ServiceInfo, config, &browniePoints)
			pq[counter] = &model.YakoAgent{
				ID:            agentID,
				BrowniePoints: browniePoints,
			}
			counter++
		}
	}
	heap.Init(&pq)

	// Select the top X ones to be recommended
	// X is the number of nodes specified by the user
	x := pq.Len()
	if x > 3 {
		x = 3
	}
	recommendedYakoAgents := make([]*model.YakoAgent, x)
	for i := 0; i < x; i++ {
		recommendedYakoAgents[i] = heap.Pop(&pq).(*model.YakoAgent)
	}

	return recommendedYakoAgents
}