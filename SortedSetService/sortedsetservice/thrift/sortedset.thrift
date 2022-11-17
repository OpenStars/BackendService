namespace cpp Database.SortedSet
namespace go Database.SortedSet
namespace java Database.SortedSet

enum TErrorCode{
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2 ,
    EDataExisted = -3
    EIterExceed = -4
}

struct TItem {
    1:string key
    2:binary value
    3:i64 score
    4:i64 index
}

struct TItemSet {
    1:string set_id
    2:string key
    3:binary value
    4:i64 score
}

struct TSortedSet{
    1: string id
    2: list<TItem> sets // max_length = 200 item
    3: map<string,TItem> dict
}


struct TListItemResult {
    1:i16 code
    2:string message
    3:list<TItem> lsItems
    4:i64 total // total item in sets
}


struct TBoolResult {
    1:i16 code
    2:string message
    3:bool result
}


service TDataService{
    
}

service SortedSetService extends TDataService{
    TBoolResult AddItemToSet(1:string set_id,2:TItem item)
    TBoolResult AddListItem(1:list<TItemSet> lsItems)
    TBoolResult RemoveItemInSet(1:string set_id,2:string item_id)
    TBoolResult RemoveListItem(1:list<TItemSet> lsItems)
    TListItemResult GetListItem(1:string set_id,2:i16 offset,3:i16 limit,4:bool is_desc) //asc || desc
}


