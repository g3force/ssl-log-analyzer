package sslproto

func (l *LogReader) CreateVisionWrapperChannel(channel chan *SSL_WrapperPacket) {
	logMessageChannel := make(chan *LogMessage, 100)
	go l.CreateLogMessageChannel(logMessageChannel)

	for logMessage := range logMessageChannel {
		if logMessage.MessageType == MESSAGE_SSL_VISION_2014 {
			visionMsg := parseVision2014(logMessage)
			channel <- visionMsg
		}
	}
	close(channel)
	return
}

func (l *LogReader) CreateVisionDetectionChannel(channel chan *SSL_DetectionFrame) {
	logMessageChannel := make(chan *LogMessage, 100)
	go l.CreateLogMessageChannel(logMessageChannel)

	for logMessage := range logMessageChannel {
		if logMessage.MessageType == MESSAGE_SSL_VISION_2014 {
			visionMsg := parseVision2014(logMessage)
			if visionMsg.Detection != nil {
				channel <- visionMsg.Detection
			}
		}
	}
	close(channel)
	return
}

func (l *LogReader) CreateRefereeChannel(channel chan *SSL_Referee) {
	logMessageChannel := make(chan *LogMessage, 100)
	go l.CreateLogMessageChannel(logMessageChannel)

	for logMessage := range logMessageChannel {
		if logMessage.MessageType == MESSAGE_SSL_REFBOX_2013 {
			refereeMsg := parseReferee2013(logMessage)
			channel <- refereeMsg
		}
	}
	close(channel)
	return
}

func (l *LogReader) CreateLogMessageChannel(channel chan *LogMessage) (err error) {
	for l.HasMessage() {
		msg, err := l.ReadMessage()
		if err != nil {
			break
		}

		channel <- msg
	}
	close(channel)
	return
}
