namespace cpp OpenStars.Counters.KVStepCounter
namespace go OpenStars.Counters.KVStepCounter
namespace java OpenStars.Counters.KVStepCounter

typedef string TKey

enum TErrorCode{
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2 ,
    EDataExisted = -3
}

struct TKVCounter{
    1: TKey cid
    2: i64 value=1,
}

typedef TKVCounter TData


struct TDataResult{
    1: TErrorCode errorCode,
    2: optional TKVCounter data
    
}

service TDataServiceR{
    TDataResult getData(1: TKey key), 
}

service TDataService{
    TDataResult getData(1: TKey key), 
    TErrorCode putData(1: TKVCounter data)
}

exception InvalidOperation {
	1: i32 errorCode;
	2: string message;
}

struct TKVCounterItem{
    1:string key
    2:i64 value
}
struct ListTKVcounterItemResult{
    1:TErrorCode errorCode,
    2:list<TKVCounterItem> listItems,
}
service KVStepCounterService extends TDataService{
    i32 createGenerator(1: string genName),
    i32 removeGenerator(1: string genName),
    i64 getCurrentValue(1: string genName),
    i64 getValue(1: string genName) ,
    i64 getStepValue(1: string genName, 2: i64 step),
    i64 setValue(1:string genName,2:i64 value),
    i64 decrement(1:string genName,2:i64 value)

    ListTKVcounterItemResult getMultiCurrentValue(1:list<string> listKeys)
    ListTKVcounterItemResult getMultiValue(1:list<string> listKeys)
    ListTKVcounterItemResult getMultiStepValue(1:list<string> listKeys,2:i64 step)
}



