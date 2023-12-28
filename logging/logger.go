package logging

import (
	"fmt"
	"time"
)

type LogInitInterface interface {
	LogInit (filename string)

	Info_Log (messageInput string)
	Warning_Log (messageInput string)
	Error_Log (messageInput string)
}

func Info_Log (messageInput string, a ...interface{}) {
	messageInput = fmt.Sprintf(messageInput, a...)
	messageInput = "[" + time.Now().Format("2006-01-02 15:04:05.000") + "][INFO] " + messageInput
	fmt.Println(messageInput)
}

func Warning_Log (messageInput string, a ...interface{}) {
	messageInput = fmt.Sprintf(messageInput, a...)
	messageInput = "[" + time.Now().Format("2006-01-02 15:04:05.000") + "][WARNING] " + messageInput
	fmt.Println(messageInput)
}

func Error_Log (messageInput string, a ...interface{}) {
	messageInput = fmt.Sprintf(messageInput, a...)
	messageInput = "[" + time.Now().Format("2006-01-02 15:04:05.000") + "][ERROR] " + messageInput
	fmt.Println(messageInput)
}

func Event_Log (messageInput string, a ...interface{}) {
	messageInput = fmt.Sprintf(messageInput, a...)
	messageInput = "[" + time.Now().Format("2006-01-02 15:04:05.000") + "][EVENT] " + messageInput
	fmt.Println(messageInput)
}