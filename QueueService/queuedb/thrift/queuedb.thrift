namespace cpp Database.QueueDb
namespace go Database.QueueDb
namespace java Database.QueueDb

enum TErrorCode{
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2 ,
    EDataExisted = -3
    EIterExceed = -4
    EEmpty = -5
}


struct TItem{
    1:string key
    2:binary value
}

struct TItemQueue{
    1:string queue_id
    2:string key
    3:string value
}

struct TQueue{
    1:list<TItem> listItems
    2:i64 max_item
}

struct TItemResult {
    1:TErrorCode code
    2:TItem data

}

struct TListItemResult {
    1:TErrorCode code
    2:list<TItem> data
    3:i64 total
}

struct TListItemQueueResult {
    1:TErrorCode code
    2:list<TItemQueue> data
    3:i64 total
}

struct TBoolResult{
    1:TErrorCode code
    2:bool data
}



# service TDataServiceR{
#     TDataResult getData(1: i64 key),
# }

service TDataService{
    
}

service QueueDbService extends TDataService{
    TBoolResult addItem(1:string queue_id,2:TItem item,3:i64 max_item);
    TBoolResult addListItems(1:list<TItemQueue> items,2:i64 max_item);
    TBoolResult removeItem(1:string queue_id,2:string item_key)
    TListItemQueueResult removeListItems(1:list<TItemQueue> items)
    TListItemResult listItems(1:string queue_id,2:i32 offset,3:i32 limit,4:bool is_desc)
    TItemResult getItem(1:string queue_id,2:string item_key)
}


