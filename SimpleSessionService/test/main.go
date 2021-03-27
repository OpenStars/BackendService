package main

import (
	"log"
	"sync"
	"time"

	"github.com/OpenStars/BackendService/SimpleSessionService"
)

func main() {
	sessionsv := SimpleSessionService.NewClient(nil, "/test", "10.60.68.100", "12195")
	wg := &sync.WaitGroup{}
	wg.Add(1000)
	for i := int64(0); i < 1000; i++ {
		go func(i int64) {
			defer wg.Done()
			sessionkey, err := sessionsv.CreateSession(i, "mobile", "kaka", time.Now().Unix()+i*i)
			if err != nil {
				log.Println("err", err)
			}
			log.Println(i, "session key", sessionkey)
		}(i)
	}

	wg.Wait()
	log.Println("success")
}
