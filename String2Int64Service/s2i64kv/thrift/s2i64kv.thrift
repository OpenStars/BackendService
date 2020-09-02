namespace cpp OpenStars.Common.S2I64KV
namespace go OpenStars.Common.S2I64KV
namespace java OpenStars.Common.S2I64KV

typedef string TKey

enum TErrorCode{
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2 ,
    EDataExisted = -3,
    ECasFailed = -4,
}

struct TI64Value{
    1: i64 value = 0
}

struct TCasValue{
    1: i64 oldValue,
    2: i64 newValue
}

typedef TI64Value TData

struct TCasResultType{
    1: TErrorCode err,
    2: i64 oldValue, // old value in store.
}

struct TDataResult{
    1: TErrorCode errorCode,
    2: optional TI64Value data
    
}

service TDataServiceR{
    TDataResult getData(1: TKey key), 
}

service TDataService{
    TDataResult getData(1: TKey key), 
    TErrorCode putData(1: TKey key, 2: TI64Value data),
    
    //Compare and swap
    TCasResultType casData(1: TKey key, 2: TCasValue casVal)
}

service TString2I64KVService extends TDataService{
    
}


