package main

import (
	telenotification "github.com/OpenStars/BackendService/TeleNotification"
)

func main() {
	telenotification.NotifyServiceError("/test", "10.60.68.102", "1805", nil)
	breakChan := make(chan bool)
	<-breakChan
}
