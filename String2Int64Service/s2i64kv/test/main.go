package main

import (
	"context"
	"fmt"
	"log"

	"github.com/OpenStars/EtcdBackendService/String2Int64Service"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"

	"github.com/OpenStars/EtcdBackendService/s2skv/thrift/gen-go/OpenStars/Common/S2SKV"
	"github.com/OpenStars/EtcdBackendService/s2skv/transports"
)

func TestPutPubkey2Uid() {
	client := transports.GetS2SCompactClient("127.0.0.1", "8883")
	if client == nil || client.Client == nil {
		log.Println("Model cannot connect to backend")
		return
	}
	defer client.BackToPool()
	rs, err := client.Client.(*S2SKV.TString2StringServiceClient).PutPubkey2Uid(context.Background(), "kaka")
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println("rs:", rs)
}

func TestGetPubkey2Uid() {
	client := transports.GetS2SCompactClient("10.110.1.21", "37173")
	if client == nil || client.Client == nil {
		log.Println("Model cannot connect to backend")
		return
	}
	defer client.BackToPool()
	rs, err := client.Client.(*S2SKV.TString2StringServiceClient).GetPubkey2Uid(context.Background(), "sonlh")
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println("rs:", rs)
}

func TestGetUid2Pubkey() {
	client := transports.GetS2SCompactClient("127.0.0.1", "8883")
	if client == nil || client.Client == nil {
		log.Println("Model cannot connect to backend")
		return
	}
	defer client.BackToPool()
	rs, err := client.Client.(*S2SKV.TString2StringServiceClient).GetUid2Pubkey(context.Background(), 2)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println("rs:", rs)
}

func TestPutAddress2Uid() {
	client := transports.GetS2SCompactClient("127.0.0.1", "8883")
	if client == nil || client.Client == nil {
		log.Println("Model cannot connect to backend")
		return
	}
	defer client.BackToPool()
	rs, err := client.Client.(*S2SKV.TString2StringServiceClient).PutAddress2Uid(context.Background(), "0x123", 1)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println("rs:", rs)
}
func TestGetAddress2Pubkey() {
	client := transports.GetS2SCompactClient("127.0.0.1", "8883")
	if client == nil || client.Client == nil {
		log.Println("Model cannot connect to backend")
		return
	}
	defer client.BackToPool()
	rs, err := client.Client.(*S2SKV.TString2StringServiceClient).GetAddress2Uid(context.Background(), "0x123")
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println("rs:", rs)
}
func TestService() {
	testservice := String2Int64Service.NewString2Int64Service("/test/", []string{"10.60.1.20:2379"}, GoEndpointBackendManager.EndPoint{
		Host:      "10.110.1.21",
		Port:      "37173",
		ServiceID: "/test",
	})
	err := testservice.PutData("sonlh", 64)
	uid, err := testservice.GetData("sonlh")
	if err != nil {
		log.Println("err", err)
	}
	log.Println("uid", uid)
}
func main() {
	// TestGetPubkey2Uid()
	// TestPutPubkey2Uid()
	// TestGetUid2Pubkey()
	//TestPutAddress2Uid()
	// TestGetAddress2Pubkey()
	TestService()
}
