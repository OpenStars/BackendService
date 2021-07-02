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
    TDataResult getTokenByPubkey(1: string pubkey),
    TDataResult getPubkeyByToken(1: string token),
    TErrorCode putData(1: string pubkey, 2: string  token)
}

service TMapNotifyKVService extends TDataService{
    
}


