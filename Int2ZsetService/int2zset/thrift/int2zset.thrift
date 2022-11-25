namespace cpp Database.Int2Zset
namespace go Database.Int2Zset
namespace java Database.Int2Zset

enum TErrorCode{
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2 ,
    EDataExisted = -3
    EIterExceed = -4
    EEmpty = -5
    
}

struct TItemSet {
    1:i64 set_id
    2:string key
    3:binary value
    4:i64 score
}

struct TItem{
    1:string key
    2:binary value
    3:i64 score
}

struct TZset{
    1:list<TItem> data
}


struct TBoolResult{
    1:TErrorCode code
    2:bool data
}

struct TListItemSetResult {
    1:TErrorCode code
    2:list<TItemSet> data
    3:i64 total
}

struct TListItemResult {
    1:TErrorCode code
    2:list<TItem> data
    3:i64 total
}

service TDataService{
    
}

service Int2ZsetService extends TDataService{
    TBoolResult addItem(1:i64 set_id,2:TItem item,3:i64 max_item);
    TBoolResult addListItems(1:list<TItemSet> items,2:i64 max_item);
    TBoolResult removeItem(1:i64 set_id,2:string item_key)
    TListItemSetResult removeListItems(1:list<TItemSet> items)
    TListItemResult listItems(1:i64 set_id,2:i32 offset,3:i32 limit,4:bool is_desc)
}


