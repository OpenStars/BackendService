namespace cpp OpenStars.Common.MapPhoneNumberPubkeyKV
namespace go OpenStars.Common.MapPhoneNumberPubkeyKV
namespace java OpenStars.Common.MapPhoneNumberPubkeyKV

typedef string TKey

enum TErrorCode{
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2 ,
    EDataExisted = -3
}

struct TStringValue{
    1: string value
}

typedef TStringValue TData


struct TDataResult{
    1: TErrorCode errorCode,
    2: optional TStringValue data
    
}

service TDataServiceR{
    TDataResult getData(1: TKey key), 
}

service TDataService{
    TDataResult getPhoneNumberByPubkey(1: string key),
    TDataResult getPubkeyByPhoneNumber(1: string key),
    TErrorCode putData(1: string pubkey, 2: string  phonenumber)
}

service TMapPhoneNumberPubkeyKVService extends TDataService{
    
}


