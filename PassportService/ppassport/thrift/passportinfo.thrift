namespace cpp OpenStars.Platform.Passport
namespace go OpenStars.Platform.Passport
namespace java OpenStars.Platform.Passport

typedef i64 TKey

enum TErrorCode{
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2 ,
    EDataExisted = -3
}

struct TPassportInfo{
    1: string sha2Pwd,
    7: map<string, string> ExtData,

}

typedef TPassportInfo TData


struct TDataResult{
    1: TErrorCode errorCode,
    2: optional TPassportInfo data
    
}

service TDataServiceR{
    TDataResult getData(1: TKey key), 
}

service TDataService{
    TDataResult getData(1: TKey key), 
    TErrorCode putData(1: TKey key , 2: TPassportInfo data)
}

service TPassportService extends TDataService{
    
}


