package main

import (
	"fmt"

	"github.com/OpenStars/BackendService/PublicProfileService"
)

func main() {
	profileService := PublicProfileService.NewClient("10.60.68.102", "1805")
	profile, err := profileService.GetProfileByPubkey("0267bf7f75c27c5a2b0cc1257c561e9bb3be39aeee829c008f71175c682f4279ba")
	fmt.Println(profile, err)
	wait := make(chan bool)
	<-wait
}
