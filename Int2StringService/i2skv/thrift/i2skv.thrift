namespace cpp OpenStars.Common.I2SKV
namespace go OpenStars.Common.I2SKV
namespace java OpenStars.Common.I2SKV

typedef i64 TKey

enum TErrorCode{
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2 ,
    EDataExisted = -3
}

struct TStringValue{
    1:string value
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
    TDataResult getData(1: TKey key), 
    TErrorCode putData(1: TKey key, 2: TStringValue data)
}

service TI2StringService extends TDataService{
    
}


