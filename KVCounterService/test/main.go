package main

import (
	"TrustKeys/SocialNetworks/Centerhub/model/share"
	"log"

	"github.com/OpenStars/EtcdBackendService/KVCounterService"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

func main() {
	kvcountersv := KVCounterService.NewKVCounterServiceModel("/aa/bb/",
		[]string{"127.0.0.1:2379"},
		GoEndpointBackendManager.EndPoint{
			Host:      "10.60.68.103",
			Port:      "7974",
			ServiceID: "/aa/bbb",
		})
	v, err := kvcountersv.GetCurrentValue(share.GenIDPost)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(v)
}
