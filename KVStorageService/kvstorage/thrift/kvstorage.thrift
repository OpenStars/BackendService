namespace cpp OpenStars.Platform.KVStorage
namespace go OpenStars.Platform.KVStorage
namespace java OpenStars.Platform.KVStorage

typedef string TKey

enum TErrorCode{
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2 ,
    EDataExisted = -3,
      EIterExceed = -4
}

struct KVItem{
    1: string key
    2: string value,
}

typedef KVItem TData


struct TDataResult{
    1: TErrorCode errorCode,
    2: optional KVItem data
    
}

struct TListDataResult{
    1: TErrorCode errorCode,
    2: optional list<KVItem> data
    3: optional list<string> missingkeys
}

service TDataServiceR{
    TDataResult getData(1:string key), 
}

service TDataService{
    TDataResult getData(1:string key), 
    TErrorCode putData(1:string key, 2: KVItem data)
    TListDataResult getListData(1:list<string> lskeys)
    TErrorCode removeData(1:string key) 
    TErrorCode putMultiData(1:list<KVItem> listData)

  i64 openIterate()
    TDataResult nextItem(1:i64 sessionkey);
    TErrorCode closeIterate(1:i64 sessionkey);
        TListDataResult nexListItems(1:i64 sessionkey,2:i64 count);
}

service KVStorageService extends TDataService{
    
}


