package mqtt

const (
	CPU          = "cpu"
	GPU          = "gpu"
	Memory       = "memory"
	SysInfo      = "sysinfo"
	// MQTT Topics
	TopicCpu     = "topic/+/" + CPU
	TopicGpu     = "topic/+/" + GPU
	TopicMemory  = "topic/+/" + Memory
	TopicSysInfo = "topic/+/" + SysInfo
)

// messageHandler Callback handler for subscribed events. Processes the event
// according to the message topic
func messageHandler(client mqtt.Client, msg mqtt.Message) {
	// Topic parsing topic/<agent_socket>/<topic_name>
	mqttTopic := strings.Split(msg.Topic(), "/")
	agentSocket := mqttTopic[1]
	switch mqttTopic[2] {
	case CPU:
		var cpu []model.Cpu
		if err := json.Unmarshal(msg.Payload(), &cpu); err != nil {
			log.Println("Err", err)
		}
		updateRegistry(agentSocket, cpu)
	case GPU:
		var gpu []model.Gpu
		if err := json.Unmarshal(msg.Payload(), &gpu); err != nil {
			log.Println("Err", err)
		}
		updateRegistry(agentSocket, gpu)
	case Memory:
		var memory model.Memory
		if err := json.Unmarshal(msg.Payload(), &memory); err != nil {
			log.Println("Err", err)
		}
		updateRegistry(agentSocket, memory)
	case SysInfo:
		var sysinfo model.SysInfo
		if err := json.Unmarshal(msg.Payload(), &sysinfo); err != nil {
			log.Println("Err", err)
		}
		updateRegistry(agentSocket, sysinfo)
	}
}