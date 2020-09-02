package main

import (
	"log"
	"time"

	"github.com/OpenStars/EtcdBackendService/SimpleSessionService"
)

func TestGetSessionKey() {
	sessionsv := SimpleSessionService.NewClient(nil, "/test", "127.0.0.1", "19175")
	usinfo, err := sessionsv.GetSessionV2("391fe23b5d8a769a4acd68acd7641baa387c8b0c6d61dbd7dc82c5bd0ada88cf_1599_9_1597116142")
	log.Println("userInfo", usinfo, "err", err)
	usinfo, err = sessionsv.GetSessionV2("25e05bb1df80e15a15d54030ca99eb332d654f659f58e086039edb07bc02779e_1599_10_1597116142")
	log.Println("userInfo", usinfo, "err", err)
}

func TestCreateSession() {
	sessionsv := SimpleSessionService.NewClient(nil, "/test", "127.0.0.1", "19175")
	sskeysamsung, err := sessionsv.CreateSessionV2(1599, "Samsung", "10.60.68.103", time.Now().Unix()+30)
	log.Println("sessionkey samsung", sskeysamsung, "err", err)
	sskeyiphone, err := sessionsv.CreateSessionV2(1599, "Iphone", "10.60.68.104", time.Now().Unix())
	log.Println("sessionkey iphone", sskeyiphone, "err", err)
}

func main() {
	// TestCreateSession()
	TestGetSessionKey()
}
