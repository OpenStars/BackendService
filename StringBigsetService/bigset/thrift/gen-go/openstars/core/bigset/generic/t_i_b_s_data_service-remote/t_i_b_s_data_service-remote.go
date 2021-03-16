// Autogenerated by Thrift Compiler (0.13.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"context"
	"flag"
	"fmt"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"github.com/apache/thrift/lib/go/thrift"
	"openstars/core/bigset/generic"
)

var _ = generic.GoUnusedProtection__

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  TPutItemResult putItem(TKey bigsetID, TItem item)")
  fmt.Fprintln(os.Stderr, "  bool removeItem(TKey bigsetID, TItemKey itemKey)")
  fmt.Fprintln(os.Stderr, "  TExistedResult existed(TKey bigsetID, TItemKey itemKey)")
  fmt.Fprintln(os.Stderr, "  TItemResult getItem(TKey bigsetID, TItemKey itemKey)")
  fmt.Fprintln(os.Stderr, "  TItemSetResult getSlice(TKey bigsetID, i32 fromPos, i32 count)")
  fmt.Fprintln(os.Stderr, "  TItemSetResult getSliceFromItem(TKey bigsetID, TItemKey fromKey, i32 count)")
  fmt.Fprintln(os.Stderr, "  TItemSetResult getSliceR(TKey bigsetID, i32 fromPos, i32 count)")
  fmt.Fprintln(os.Stderr, "  TItemSetResult getSliceFromItemR(TKey bigsetID, TItemKey fromKey, i32 count)")
  fmt.Fprintln(os.Stderr, "  TItemSetResult rangeQuery(TKey bigsetID, TItemKey startKey, TItemKey endKey)")
  fmt.Fprintln(os.Stderr, "  bool bulkLoad(TKey bigsetID, TItemSet setData)")
  fmt.Fprintln(os.Stderr, "  TMultiPutItemResult multiPut(TKey bigsetID, TItemSet setData, bool getAddedItems, bool getReplacedItems)")
  fmt.Fprintln(os.Stderr, "  i64 getTotalCount(TKey bigsetID)")
  fmt.Fprintln(os.Stderr, "  i64 removeAll(TKey bigsetID)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

type httpHeaders map[string]string

func (h httpHeaders) String() string {
  var m map[string]string = h
  return fmt.Sprintf("%s", m)
}

func (h httpHeaders) Set(value string) error {
  parts := strings.Split(value, ": ")
  if len(parts) != 2 {
    return fmt.Errorf("header should be of format 'Key: Value'")
  }
  h[parts[0]] = parts[1]
  return nil
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  headers := make(httpHeaders)
  var parsedUrl *url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Var(headers, "H", "Headers to set on the http(s) request (e.g. -H \"Key: Value\")")
  flag.Parse()
  
  if len(urlString) > 0 {
    var err error
    parsedUrl, err = url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
    if len(headers) > 0 {
      httptrans := trans.(*thrift.THttpClient)
      for key, value := range headers {
        httptrans.SetHeader(key, value)
      }
    }
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  iprot := protocolFactory.GetProtocol(trans)
  oprot := protocolFactory.GetProtocol(trans)
  client := generic.NewTIBSDataServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "putItem":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "PutItem requires 2 args")
      flag.Usage()
    }
    argvalue0, err253 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err253 != nil {
      Usage()
      return
    }
    value0 := generic.TKey(argvalue0)
    arg254 := flag.Arg(2)
    mbTrans255 := thrift.NewTMemoryBufferLen(len(arg254))
    defer mbTrans255.Close()
    _, err256 := mbTrans255.WriteString(arg254)
    if err256 != nil {
      Usage()
      return
    }
    factory257 := thrift.NewTJSONProtocolFactory()
    jsProt258 := factory257.GetProtocol(mbTrans255)
    argvalue1 := generic.NewTItem()
    err259 := argvalue1.Read(jsProt258)
    if err259 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.PutItem(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "removeItem":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "RemoveItem requires 2 args")
      flag.Usage()
    }
    argvalue0, err260 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err260 != nil {
      Usage()
      return
    }
    value0 := generic.TKey(argvalue0)
    argvalue1 := []byte(flag.Arg(2))
    value1 := generic.TItemKey(argvalue1)
    fmt.Print(client.RemoveItem(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "existed":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "Existed requires 2 args")
      flag.Usage()
    }
    argvalue0, err262 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err262 != nil {
      Usage()
      return
    }
    value0 := generic.TKey(argvalue0)
    argvalue1 := []byte(flag.Arg(2))
    value1 := generic.TItemKey(argvalue1)
    fmt.Print(client.Existed(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "getItem":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "GetItem requires 2 args")
      flag.Usage()
    }
    argvalue0, err264 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err264 != nil {
      Usage()
      return
    }
    value0 := generic.TKey(argvalue0)
    argvalue1 := []byte(flag.Arg(2))
    value1 := generic.TItemKey(argvalue1)
    fmt.Print(client.GetItem(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "getSlice":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "GetSlice requires 3 args")
      flag.Usage()
    }
    argvalue0, err266 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err266 != nil {
      Usage()
      return
    }
    value0 := generic.TKey(argvalue0)
    tmp1, err267 := (strconv.Atoi(flag.Arg(2)))
    if err267 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    tmp2, err268 := (strconv.Atoi(flag.Arg(3)))
    if err268 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    fmt.Print(client.GetSlice(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "getSliceFromItem":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "GetSliceFromItem requires 3 args")
      flag.Usage()
    }
    argvalue0, err269 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err269 != nil {
      Usage()
      return
    }
    value0 := generic.TKey(argvalue0)
    argvalue1 := []byte(flag.Arg(2))
    value1 := generic.TItemKey(argvalue1)
    tmp2, err271 := (strconv.Atoi(flag.Arg(3)))
    if err271 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    fmt.Print(client.GetSliceFromItem(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "getSliceR":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "GetSliceR requires 3 args")
      flag.Usage()
    }
    argvalue0, err272 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err272 != nil {
      Usage()
      return
    }
    value0 := generic.TKey(argvalue0)
    tmp1, err273 := (strconv.Atoi(flag.Arg(2)))
    if err273 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    tmp2, err274 := (strconv.Atoi(flag.Arg(3)))
    if err274 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    fmt.Print(client.GetSliceR(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "getSliceFromItemR":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "GetSliceFromItemR requires 3 args")
      flag.Usage()
    }
    argvalue0, err275 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err275 != nil {
      Usage()
      return
    }
    value0 := generic.TKey(argvalue0)
    argvalue1 := []byte(flag.Arg(2))
    value1 := generic.TItemKey(argvalue1)
    tmp2, err277 := (strconv.Atoi(flag.Arg(3)))
    if err277 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    fmt.Print(client.GetSliceFromItemR(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "rangeQuery":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "RangeQuery requires 3 args")
      flag.Usage()
    }
    argvalue0, err278 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err278 != nil {
      Usage()
      return
    }
    value0 := generic.TKey(argvalue0)
    argvalue1 := []byte(flag.Arg(2))
    value1 := generic.TItemKey(argvalue1)
    argvalue2 := []byte(flag.Arg(3))
    value2 := generic.TItemKey(argvalue2)
    fmt.Print(client.RangeQuery(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "bulkLoad":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "BulkLoad requires 2 args")
      flag.Usage()
    }
    argvalue0, err281 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err281 != nil {
      Usage()
      return
    }
    value0 := generic.TKey(argvalue0)
    arg282 := flag.Arg(2)
    mbTrans283 := thrift.NewTMemoryBufferLen(len(arg282))
    defer mbTrans283.Close()
    _, err284 := mbTrans283.WriteString(arg282)
    if err284 != nil {
      Usage()
      return
    }
    factory285 := thrift.NewTJSONProtocolFactory()
    jsProt286 := factory285.GetProtocol(mbTrans283)
    argvalue1 := generic.NewTItemSet()
    err287 := argvalue1.Read(jsProt286)
    if err287 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.BulkLoad(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "multiPut":
    if flag.NArg() - 1 != 4 {
      fmt.Fprintln(os.Stderr, "MultiPut requires 4 args")
      flag.Usage()
    }
    argvalue0, err288 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err288 != nil {
      Usage()
      return
    }
    value0 := generic.TKey(argvalue0)
    arg289 := flag.Arg(2)
    mbTrans290 := thrift.NewTMemoryBufferLen(len(arg289))
    defer mbTrans290.Close()
    _, err291 := mbTrans290.WriteString(arg289)
    if err291 != nil {
      Usage()
      return
    }
    factory292 := thrift.NewTJSONProtocolFactory()
    jsProt293 := factory292.GetProtocol(mbTrans290)
    argvalue1 := generic.NewTItemSet()
    err294 := argvalue1.Read(jsProt293)
    if err294 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    argvalue2 := flag.Arg(3) == "true"
    value2 := argvalue2
    argvalue3 := flag.Arg(4) == "true"
    value3 := argvalue3
    fmt.Print(client.MultiPut(context.Background(), value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "getTotalCount":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetTotalCount requires 1 args")
      flag.Usage()
    }
    argvalue0, err297 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err297 != nil {
      Usage()
      return
    }
    value0 := generic.TKey(argvalue0)
    fmt.Print(client.GetTotalCount(context.Background(), value0))
    fmt.Print("\n")
    break
  case "removeAll":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "RemoveAll requires 1 args")
      flag.Usage()
    }
    argvalue0, err298 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err298 != nil {
      Usage()
      return
    }
    value0 := generic.TKey(argvalue0)
    fmt.Print(client.RemoveAll(context.Background(), value0))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
