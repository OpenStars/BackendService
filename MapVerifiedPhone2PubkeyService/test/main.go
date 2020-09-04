package main

import (
	"log"

	"github.com/OpenStars/EtcdBackendService/MapPhoneNumber2Pubkey"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

func main() {
	mapPub2Phone := MapPhoneNumber2Pubkey.NewMappingPhone2Pubkey("/test/", []string{"10.60.1.20:2379"}, GoEndpointBackendManager.EndPoint{
		Host:      "10.60.68.102",
		Port:      "2383",
		ServiceID: "/test/",
	})
	mapPub2Phone.PutData("0279ac5ffd0de530bc49544e42fe68c725e9a6bbdd37ffacdd983d7d9f6835b1c4", "+84384399862")
	phone, err := mapPub2Phone.GetPhoneNumberByPubkey("0279ac5ffd0de530bc49544e42fe68c725e9a6bbdd37ffacdd983d7d9f6835b1c4")
	log.Println("phone ", phone, "err", err)
	// pubkey, err := mapPub2Phone.GetPhoneNumberByPubkey("0355a44ec09f09b34a2a8394942fcbe04c2b2e63ef0ad75776e0c2dce1a69ce141")
	// log.Println("pubkey", pubkey, "err", err)

	mapPub2Phone.PutData("0355a44ec09f09b34a2a8394942fcbe04c2b2e63ef0ad75776e0c2dce1a69ce141", "+84384399862")
	data, err := mapPub2Phone.GetPhoneNumberByPubkey("0279ac5ffd0de530bc49544e42fe68c725e9a6bbdd37ffacdd983d7d9f6835b1c4")
	// log.Println("data", data, "err", err)
	// mapPub2Phone.PutData("0355a44ec09f09b34a2a8394942fcbe04c2b2e63ef0ad75776e0c2dce1a69ce141", "+84384399862")
	// data, err = mapPub2Phone.GetPhoneNumberByPubkey("0279ac5ffd0de530bc49544e42fe68c725e9a6bbdd37ffacdd983d7d9f6835b1c4")
	log.Println("data", data, "err", err)
}
