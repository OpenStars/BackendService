package main

import (
	"fmt"

	"github.com/OpenStars/BackendService/MapNotifyCallService"
)

func main() {
	mappubkey2iostoken := MapNotifyCallService.NewMapNotifyCallService(nil, "/test", "10.60.68.103", "9199")
	token, err := mappubkey2iostoken.GetTokenByPubkey("030ce82a17952be5466f9e64c8eb00467c0063ed6a58dde893b90856e9e6226f95")
	fmt.Println(token, err)
}
