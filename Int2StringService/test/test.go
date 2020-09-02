package main

import (
	"log"

	"github.com/OpenStars/EtcdBackendService/Int2StringService"
	"github.com/OpenStars/EtcdBackendService/String2Int64Service"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

type Server struct {
	int2string   Int2StringService.Int2StringServiceIf
	string2int   String2Int64Service.String2Int64ServiceIf
	stringbigset StringBigsetService.StringBigsetServiceIf
}

func (s *Server) Run() {
	for {

	}
}

func TestStringBigset() {
	sid := "/trustkeys/tkverifyprofile/stringbigset"
	etcd := []string{"127.0.0.1:2379"}
	defaultEp := GoEndpointBackendManager.EndPoint{
		ServiceID: sid,
		Host:      "127.0.0.1",
		Port:      "8883",
	}
	ai2s := StringBigsetService.NewStringBigsetServiceModel(sid, etcd, defaultEp)
	sv := &Server{
		stringbigset: ai2s,
	}
	sv.Run()
}

func TestI64() {
	sv := Int2StringService.NewInt2StringService("/test/", []string{"10.60.1.20:2379"}, GoEndpointBackendManager.EndPoint{
		Host:      "10.110.1.21",
		Port:      "37183",
		ServiceID: "/test/",
	})
	err := sv.PutData(64, "sonlh")
	data, err := sv.GetData(64)
	if err != nil {
		log.Println("err", err)
	}
	log.Println("data", data)
}

func TestString2Int() {
	sid := "/openstars/services/string2int"
	etcd := []string{"127.0.0.1:2379"}
	defaultEp := GoEndpointBackendManager.EndPoint{
		ServiceID: sid,
		Host:      "127.0.0.1",
		Port:      "8883",
	}
	ai2s := String2Int64Service.NewString2Int64Service(sid, etcd, defaultEp)
	sv := &Server{
		string2int: ai2s,
	}
	sv.Run()
}

func TestInt2String() {
	sid := "/openstars/services/int2string"
	etcd := []string{"127.0.0.1:2379"}
	defaultEp := GoEndpointBackendManager.EndPoint{
		ServiceID: sid,
		Host:      "127.0.0.1",
		Port:      "8883",
	}
	ai2s := Int2StringService.NewInt2StringService(sid, etcd, defaultEp)
	sv := &Server{
		int2string: ai2s,
	}
	sv.Run()
}

func main() {
	// TestInt2String()
	// TestString2Int()

	// TestStringBigset()
	TestI64()
}
