package MapNotifyCallKV

import (
	"bytes"
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"

	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

type TErrorCode int64

const (
	TErrorCode_EGood        TErrorCode = 0
	TErrorCode_ENotFound    TErrorCode = -1
	TErrorCode_EUnknown     TErrorCode = -2
	TErrorCode_EDataExisted TErrorCode = -3
)

func (p TErrorCode) String() string {
	switch p {
	case TErrorCode_EGood:
		return "EGood"
	case TErrorCode_ENotFound:
		return "ENotFound"
	case TErrorCode_EUnknown:
		return "EUnknown"
	case TErrorCode_EDataExisted:
		return "EDataExisted"
	}
	return "<UNSET>"
}

func TErrorCodeFromString(s string) (TErrorCode, error) {
	switch s {
	case "EGood":
		return TErrorCode_EGood, nil
	case "ENotFound":
		return TErrorCode_ENotFound, nil
	case "EUnknown":
		return TErrorCode_EUnknown, nil
	case "EDataExisted":
		return TErrorCode_EDataExisted, nil
	}
	return TErrorCode(0), fmt.Errorf("not a valid TErrorCode string")
}

func TErrorCodePtr(v TErrorCode) *TErrorCode { return &v }

func (p TErrorCode) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *TErrorCode) UnmarshalText(text []byte) error {
	q, err := TErrorCodeFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

func (p *TErrorCode) Scan(value interface{}) error {
	v, ok := value.(int64)
	if !ok {
		return errors.New("Scan value is not int64")
	}
	*p = TErrorCode(v)
	return nil
}

func (p *TErrorCode) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type TKey string

func TKeyPtr(v TKey) *TKey { return &v }

type TData *TStringValue

func TDataPtr(v TData) *TData { return &v }

// Attributes:
//  - Value
type TStringValue struct {
	Value string `thrift:"value,1" db:"value" json:"value"`
}

func NewTStringValue() *TStringValue {
	return &TStringValue{}
}

func (p *TStringValue) GetValue() string {
	return p.Value
}
func (p *TStringValue) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TStringValue) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Value = v
	}
	return nil
}

func (p *TStringValue) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TStringValue"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TStringValue) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("value", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:value: ", p), err)
	}
	if err := oprot.WriteString(string(p.Value)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.value (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:value: ", p), err)
	}
	return err
}

func (p *TStringValue) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TStringValue(%+v)", *p)
}

// Attributes:
//  - ErrorCode
//  - Data
type TDataResult_ struct {
	ErrorCode TErrorCode    `thrift:"errorCode,1" db:"errorCode" json:"errorCode"`
	Data      *TStringValue `thrift:"data,2" db:"data" json:"data,omitempty"`
}

func NewTDataResult_() *TDataResult_ {
	return &TDataResult_{}
}

func (p *TDataResult_) GetErrorCode() TErrorCode {
	return p.ErrorCode
}

var TDataResult__Data_DEFAULT *TStringValue

func (p *TDataResult_) GetData() *TStringValue {
	if !p.IsSetData() {
		return TDataResult__Data_DEFAULT
	}
	return p.Data
}
func (p *TDataResult_) IsSetData() bool {
	return p.Data != nil
}

func (p *TDataResult_) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField1(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField2(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TDataResult_) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := TErrorCode(v)
		p.ErrorCode = temp
	}
	return nil
}

func (p *TDataResult_) ReadField2(iprot thrift.TProtocol) error {
	p.Data = &TStringValue{}
	if err := p.Data.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Data), err)
	}
	return nil
}

func (p *TDataResult_) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TDataResult"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TDataResult_) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("errorCode", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:errorCode: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.ErrorCode)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.errorCode (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:errorCode: ", p), err)
	}
	return err
}

func (p *TDataResult_) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetData() {
		if err := oprot.WriteFieldBegin("data", thrift.STRUCT, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:data: ", p), err)
		}
		if err := p.Data.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Data), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:data: ", p), err)
		}
	}
	return err
}

func (p *TDataResult_) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TDataResult_(%+v)", *p)
}

type TDataServiceR interface {
	// Parameters:
	//  - Key
	GetData(ctx context.Context, key TKey) (r *TDataResult_, err error)
}

type TDataServiceRClient struct {
	c thrift.TClient
}

// Deprecated: Use NewTDataServiceR instead
func NewTDataServiceRClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *TDataServiceRClient {
	return &TDataServiceRClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

// Deprecated: Use NewTDataServiceR instead
func NewTDataServiceRClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *TDataServiceRClient {
	return &TDataServiceRClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewTDataServiceRClient(c thrift.TClient) *TDataServiceRClient {
	return &TDataServiceRClient{
		c: c,
	}
}

// Parameters:
//  - Key
func (p *TDataServiceRClient) GetData(ctx context.Context, key TKey) (r *TDataResult_, err error) {
	var _args0 TDataServiceRGetDataArgs
	_args0.Key = key
	var _result1 TDataServiceRGetDataResult
	if err = p.c.Call(ctx, "getData", &_args0, &_result1); err != nil {
		return
	}
	return _result1.GetSuccess(), nil
}

type TDataServiceRProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      TDataServiceR
}

func (p *TDataServiceRProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *TDataServiceRProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *TDataServiceRProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewTDataServiceRProcessor(handler TDataServiceR) *TDataServiceRProcessor {

	self2 := &TDataServiceRProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self2.processorMap["getData"] = &tDataServiceRProcessorGetData{handler: handler}
	return self2
}

func (p *TDataServiceRProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x3.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x3

}

type tDataServiceRProcessorGetData struct {
	handler TDataServiceR
}

func (p *tDataServiceRProcessorGetData) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := TDataServiceRGetDataArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getData", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	result := TDataServiceRGetDataResult{}
	var retval *TDataResult_
	var err2 error
	if retval, err2 = p.handler.GetData(ctx, args.Key); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getData: "+err2.Error())
		oprot.WriteMessageBegin("getData", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getData", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Key
type TDataServiceRGetDataArgs struct {
	Key TKey `thrift:"key,1" db:"key" json:"key"`
}

func NewTDataServiceRGetDataArgs() *TDataServiceRGetDataArgs {
	return &TDataServiceRGetDataArgs{}
}

func (p *TDataServiceRGetDataArgs) GetKey() TKey {
	return p.Key
}
func (p *TDataServiceRGetDataArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TDataServiceRGetDataArgs) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := TKey(v)
		p.Key = temp
	}
	return nil
}

func (p *TDataServiceRGetDataArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getData_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TDataServiceRGetDataArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("key", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:key: ", p), err)
	}
	if err := oprot.WriteString(string(p.Key)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.key (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:key: ", p), err)
	}
	return err
}

func (p *TDataServiceRGetDataArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TDataServiceRGetDataArgs(%+v)", *p)
}

// Attributes:
//  - Success
type TDataServiceRGetDataResult struct {
	Success *TDataResult_ `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewTDataServiceRGetDataResult() *TDataServiceRGetDataResult {
	return &TDataServiceRGetDataResult{}
}

var TDataServiceRGetDataResult_Success_DEFAULT *TDataResult_

func (p *TDataServiceRGetDataResult) GetSuccess() *TDataResult_ {
	if !p.IsSetSuccess() {
		return TDataServiceRGetDataResult_Success_DEFAULT
	}
	return p.Success
}
func (p *TDataServiceRGetDataResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *TDataServiceRGetDataResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TDataServiceRGetDataResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &TDataResult_{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *TDataServiceRGetDataResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getData_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TDataServiceRGetDataResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *TDataServiceRGetDataResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TDataServiceRGetDataResult(%+v)", *p)
}

type TDataService interface {
	// Parameters:
	//  - Pubkey
	GetTokenByPubkey(ctx context.Context, pubkey string) (r *TDataResult_, err error)
	// Parameters:
	//  - Token
	GetPubkeyByToken(ctx context.Context, token string) (r *TDataResult_, err error)
	// Parameters:
	//  - Pubkey
	//  - Token
	PutData(ctx context.Context, pubkey string, token string) (r TErrorCode, err error)
}

type TDataServiceClient struct {
	c thrift.TClient
}

// Deprecated: Use NewTDataService instead
func NewTDataServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *TDataServiceClient {
	return &TDataServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

// Deprecated: Use NewTDataService instead
func NewTDataServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *TDataServiceClient {
	return &TDataServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewTDataServiceClient(c thrift.TClient) *TDataServiceClient {
	return &TDataServiceClient{
		c: c,
	}
}

// Parameters:
//  - Pubkey
func (p *TDataServiceClient) GetTokenByPubkey(ctx context.Context, pubkey string) (r *TDataResult_, err error) {
	var _args5 TDataServiceGetTokenByPubkeyArgs
	_args5.Pubkey = pubkey
	var _result6 TDataServiceGetTokenByPubkeyResult
	if err = p.c.Call(ctx, "getTokenByPubkey", &_args5, &_result6); err != nil {
		return
	}
	return _result6.GetSuccess(), nil
}

// Parameters:
//  - Token
func (p *TDataServiceClient) GetPubkeyByToken(ctx context.Context, token string) (r *TDataResult_, err error) {
	var _args7 TDataServiceGetPubkeyByTokenArgs
	_args7.Token = token
	var _result8 TDataServiceGetPubkeyByTokenResult
	if err = p.c.Call(ctx, "getPubkeyByToken", &_args7, &_result8); err != nil {
		return
	}
	return _result8.GetSuccess(), nil
}

// Parameters:
//  - Pubkey
//  - Token
func (p *TDataServiceClient) PutData(ctx context.Context, pubkey string, token string) (r TErrorCode, err error) {
	var _args9 TDataServicePutDataArgs
	_args9.Pubkey = pubkey
	_args9.Token = token
	var _result10 TDataServicePutDataResult
	if err = p.c.Call(ctx, "putData", &_args9, &_result10); err != nil {
		return
	}
	return _result10.GetSuccess(), nil
}

type TDataServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      TDataService
}

func (p *TDataServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *TDataServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *TDataServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewTDataServiceProcessor(handler TDataService) *TDataServiceProcessor {

	self11 := &TDataServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self11.processorMap["getTokenByPubkey"] = &tDataServiceProcessorGetTokenByPubkey{handler: handler}
	self11.processorMap["getPubkeyByToken"] = &tDataServiceProcessorGetPubkeyByToken{handler: handler}
	self11.processorMap["putData"] = &tDataServiceProcessorPutData{handler: handler}
	return self11
}

func (p *TDataServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x12 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x12.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x12

}

type tDataServiceProcessorGetTokenByPubkey struct {
	handler TDataService
}

func (p *tDataServiceProcessorGetTokenByPubkey) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := TDataServiceGetTokenByPubkeyArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getTokenByPubkey", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	result := TDataServiceGetTokenByPubkeyResult{}
	var retval *TDataResult_
	var err2 error
	if retval, err2 = p.handler.GetTokenByPubkey(ctx, args.Pubkey); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getTokenByPubkey: "+err2.Error())
		oprot.WriteMessageBegin("getTokenByPubkey", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getTokenByPubkey", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type tDataServiceProcessorGetPubkeyByToken struct {
	handler TDataService
}

func (p *tDataServiceProcessorGetPubkeyByToken) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := TDataServiceGetPubkeyByTokenArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getPubkeyByToken", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	result := TDataServiceGetPubkeyByTokenResult{}
	var retval *TDataResult_
	var err2 error
	if retval, err2 = p.handler.GetPubkeyByToken(ctx, args.Token); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getPubkeyByToken: "+err2.Error())
		oprot.WriteMessageBegin("getPubkeyByToken", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getPubkeyByToken", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type tDataServiceProcessorPutData struct {
	handler TDataService
}

func (p *tDataServiceProcessorPutData) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := TDataServicePutDataArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("putData", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	result := TDataServicePutDataResult{}
	var retval TErrorCode
	var err2 error
	if retval, err2 = p.handler.PutData(ctx, args.Pubkey, args.Token); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing putData: "+err2.Error())
		oprot.WriteMessageBegin("putData", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("putData", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Pubkey
type TDataServiceGetTokenByPubkeyArgs struct {
	Pubkey string `thrift:"pubkey,1" db:"pubkey" json:"pubkey"`
}

func NewTDataServiceGetTokenByPubkeyArgs() *TDataServiceGetTokenByPubkeyArgs {
	return &TDataServiceGetTokenByPubkeyArgs{}
}

func (p *TDataServiceGetTokenByPubkeyArgs) GetPubkey() string {
	return p.Pubkey
}
func (p *TDataServiceGetTokenByPubkeyArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TDataServiceGetTokenByPubkeyArgs) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Pubkey = v
	}
	return nil
}

func (p *TDataServiceGetTokenByPubkeyArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getTokenByPubkey_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TDataServiceGetTokenByPubkeyArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("pubkey", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:pubkey: ", p), err)
	}
	if err := oprot.WriteString(string(p.Pubkey)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.pubkey (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:pubkey: ", p), err)
	}
	return err
}

func (p *TDataServiceGetTokenByPubkeyArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TDataServiceGetTokenByPubkeyArgs(%+v)", *p)
}

// Attributes:
//  - Success
type TDataServiceGetTokenByPubkeyResult struct {
	Success *TDataResult_ `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewTDataServiceGetTokenByPubkeyResult() *TDataServiceGetTokenByPubkeyResult {
	return &TDataServiceGetTokenByPubkeyResult{}
}

var TDataServiceGetTokenByPubkeyResult_Success_DEFAULT *TDataResult_

func (p *TDataServiceGetTokenByPubkeyResult) GetSuccess() *TDataResult_ {
	if !p.IsSetSuccess() {
		return TDataServiceGetTokenByPubkeyResult_Success_DEFAULT
	}
	return p.Success
}
func (p *TDataServiceGetTokenByPubkeyResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *TDataServiceGetTokenByPubkeyResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TDataServiceGetTokenByPubkeyResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &TDataResult_{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *TDataServiceGetTokenByPubkeyResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getTokenByPubkey_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TDataServiceGetTokenByPubkeyResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *TDataServiceGetTokenByPubkeyResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TDataServiceGetTokenByPubkeyResult(%+v)", *p)
}

// Attributes:
//  - Token
type TDataServiceGetPubkeyByTokenArgs struct {
	Token string `thrift:"token,1" db:"token" json:"token"`
}

func NewTDataServiceGetPubkeyByTokenArgs() *TDataServiceGetPubkeyByTokenArgs {
	return &TDataServiceGetPubkeyByTokenArgs{}
}

func (p *TDataServiceGetPubkeyByTokenArgs) GetToken() string {
	return p.Token
}
func (p *TDataServiceGetPubkeyByTokenArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TDataServiceGetPubkeyByTokenArgs) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Token = v
	}
	return nil
}

func (p *TDataServiceGetPubkeyByTokenArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getPubkeyByToken_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TDataServiceGetPubkeyByTokenArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("token", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:token: ", p), err)
	}
	if err := oprot.WriteString(string(p.Token)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.token (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:token: ", p), err)
	}
	return err
}

func (p *TDataServiceGetPubkeyByTokenArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TDataServiceGetPubkeyByTokenArgs(%+v)", *p)
}

// Attributes:
//  - Success
type TDataServiceGetPubkeyByTokenResult struct {
	Success *TDataResult_ `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewTDataServiceGetPubkeyByTokenResult() *TDataServiceGetPubkeyByTokenResult {
	return &TDataServiceGetPubkeyByTokenResult{}
}

var TDataServiceGetPubkeyByTokenResult_Success_DEFAULT *TDataResult_

func (p *TDataServiceGetPubkeyByTokenResult) GetSuccess() *TDataResult_ {
	if !p.IsSetSuccess() {
		return TDataServiceGetPubkeyByTokenResult_Success_DEFAULT
	}
	return p.Success
}
func (p *TDataServiceGetPubkeyByTokenResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *TDataServiceGetPubkeyByTokenResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TDataServiceGetPubkeyByTokenResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &TDataResult_{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *TDataServiceGetPubkeyByTokenResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getPubkeyByToken_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TDataServiceGetPubkeyByTokenResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *TDataServiceGetPubkeyByTokenResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TDataServiceGetPubkeyByTokenResult(%+v)", *p)
}

// Attributes:
//  - Pubkey
//  - Token
type TDataServicePutDataArgs struct {
	Pubkey string `thrift:"pubkey,1" db:"pubkey" json:"pubkey"`
	Token  string `thrift:"token,2" db:"token" json:"token"`
}

func NewTDataServicePutDataArgs() *TDataServicePutDataArgs {
	return &TDataServicePutDataArgs{}
}

func (p *TDataServicePutDataArgs) GetPubkey() string {
	return p.Pubkey
}

func (p *TDataServicePutDataArgs) GetToken() string {
	return p.Token
}
func (p *TDataServicePutDataArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TDataServicePutDataArgs) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Pubkey = v
	}
	return nil
}

func (p *TDataServicePutDataArgs) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Token = v
	}
	return nil
}

func (p *TDataServicePutDataArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("putData_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TDataServicePutDataArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("pubkey", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:pubkey: ", p), err)
	}
	if err := oprot.WriteString(string(p.Pubkey)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.pubkey (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:pubkey: ", p), err)
	}
	return err
}

func (p *TDataServicePutDataArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("token", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:token: ", p), err)
	}
	if err := oprot.WriteString(string(p.Token)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.token (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:token: ", p), err)
	}
	return err
}

func (p *TDataServicePutDataArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TDataServicePutDataArgs(%+v)", *p)
}

// Attributes:
//  - Success
type TDataServicePutDataResult struct {
	Success *TErrorCode `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewTDataServicePutDataResult() *TDataServicePutDataResult {
	return &TDataServicePutDataResult{}
}

var TDataServicePutDataResult_Success_DEFAULT TErrorCode

func (p *TDataServicePutDataResult) GetSuccess() TErrorCode {
	if !p.IsSetSuccess() {
		return TDataServicePutDataResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *TDataServicePutDataResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *TDataServicePutDataResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField0(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TDataServicePutDataResult) ReadField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		temp := TErrorCode(v)
		p.Success = &temp
	}
	return nil
}

func (p *TDataServicePutDataResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("putData_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TDataServicePutDataResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.I32, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *TDataServicePutDataResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TDataServicePutDataResult(%+v)", *p)
}

type TMapNotifyKVService interface {
	TDataService
}

type TMapNotifyKVServiceClient struct {
	c thrift.TClient
	*TDataServiceClient
}

// Deprecated: Use NewTMapNotifyKVService instead
func NewTMapNotifyKVServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *TMapNotifyKVServiceClient {
	return &TMapNotifyKVServiceClient{TDataServiceClient: NewTDataServiceClientFactory(t, f)}
}

// Deprecated: Use NewTMapNotifyKVService instead
func NewTMapNotifyKVServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *TMapNotifyKVServiceClient {
	return &TMapNotifyKVServiceClient{TDataServiceClient: NewTDataServiceClientProtocol(t, iprot, oprot)}
}

func NewTMapNotifyKVServiceClient(c thrift.TClient) *TMapNotifyKVServiceClient {
	return &TMapNotifyKVServiceClient{
		c:                  c,
		TDataServiceClient: NewTDataServiceClient(c),
	}
}

type TMapNotifyKVServiceProcessor struct {
	*TDataServiceProcessor
}

func NewTMapNotifyKVServiceProcessor(handler TMapNotifyKVService) *TMapNotifyKVServiceProcessor {
	self17 := &TMapNotifyKVServiceProcessor{NewTDataServiceProcessor(handler)}
	return self17
}

// HELPER FUNCTIONS AND STRUCTURES
