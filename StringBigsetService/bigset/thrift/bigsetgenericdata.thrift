namespace cpp OpenStars.Core.BigSet.Generic
namespace java org.openstars.core.bigset.Generic
namespace go openstars.core.bigset.generic

typedef binary TItemKey //key - index of an item, with simple set, itemkey is equivalent to item
typedef binary TItemValue 

struct TItem{
    1: required binary key,
    2: required binary value
}

typedef list<TItem> TItemList

typedef list<TItemKey> TItemKeySet

typedef i64 TKey //Key of bigset service - big set ID 

typedef i64 TContainerKey

typedef TContainerKey TSmallSetIDKey

typedef i16 TLevelType

struct TSmallSet {
	1: TItemList entries,
        2: optional TSmallSetIDKey prev,
        3: optional TSmallSetIDKey nxt
}

struct TItemSet{
    1: list<TItem> items
}

typedef TContainerKey TMetaKey

struct TMetaItem{
    1: TItemKey minItem,
    2: i64 metaID,
    3: i32 count // khi nao split moi can load len
}

struct TNeedSplitInfo{
    1: TMetaKey metaID,
    2: TMetaKey parentID, // parent of metaID
    3: i32 childCount,
    4: bool isSmallSet
}

struct TBigSetGenericData{
    1: TContainerKey parentID = 0, // <=0 means it is top level meta, parent meta ID
    2: TLevelType level = 0, // level 0 means that its children are small-set-ID or it is small set if it is -1    
    3: optional list<TMetaItem> children, // sort by minItem
    4: optional TSmallSet smallset,
    5: optional TNeedSplitInfo splitinfo,
}

struct TSetMetaPathItem{
    1: TItemKey minItem,
    2: i64 metaID,
    3: i8 level,
}


struct TSetMetaPath{
    1: list<TSetMetaPathItem> metaPath,
    2: optional TSetMetaPathItem smallSetInfo,
    3: optional TNeedSplitInfo splitInfo,
}

struct TGetSliceInfo{
    1: list<TMetaKey> smallSetIDs,
    2: i32 firstIdxFrom
}

service MasterMetaService{
    TMetaKey getMetaID(1: TKey key),

    bool setMetaID(1:TKey key, 2: TMetaKey metaID),

}

struct TSmallSetInfo{
    1:i32 count,
    2:TItemKey midItem,
    3:i32 countFromMid,
    4:optional TSmallSetIDKey prev = 0,
    5:optional TSmallSetIDKey nxt = 0
}

enum TErrorCode{
    EGood = 0,
    ENoRootMeta = 1,
    ECouldNotConnectMetadata = 2,
    ECouldNotConnectIDGen = 3,
    ECouldNotConnectSmallSet = 4,
    ECouldNotConnectRootMetaMapping = 5,
    EItemNotExisted = 100,
    EUnknownException = 10,

    EBigSetNotExisted = 101, // not existed bigset with specific name
    EBigSetCreated, // created new bigset ID
    EBigSetAlreadyExisted,
    EBigSetAssigned, //assigned a bigset name to a specific IDs.
}

struct TItemResult{
    1: TErrorCode error, // 0: ok, -1 : error
    2: optional TItem item
}

struct TItemSetResult{
    1: TErrorCode error,
    2: optional TItemSet items,
}

struct TPutItemResult{
    1: TErrorCode error, // 0: ok, -1 : error
    2: bool ok,
    3: optional TItem oldItem,
}

struct TExistedResult{
    1: TErrorCode error, // 0: ok, -1 : error
    2: bool existed
}

struct TMultiPutItemResult{
    1: TErrorCode error,
    2: optional list<TItemKey> added,
    3: optional list<TItem> replaced, // old items were replaced by operation
}


//Small set ID and meta id using the same IDgenerator key
struct TSplitBigSetResult{
    1: TErrorCode error,
    2: bool splited,
    3: TContainerKey newRootID,
    4: TItemKey minItem,
    5: i64 count, // number of item in new set
}

service TBSGenericDataService{

    TPutItemResult bsgPutItem(1:TContainerKey rootID, 2:TItem item),

    /*return true if item is existed in the list otherwise return false*/
    bool bsgRemoveItem(1:TMetaKey key, 2:TItemKey itemKey),

    TExistedResult bsgExisted(1:TContainerKey rootID, 2: TItemKey itemKey),

    TItemResult bsgGetItem(1:TContainerKey rootID, 2: TItemKey itemKey),

    TItemSetResult bsgGetSlice(1:TContainerKey rootID, 2: i32 fromIdx, 3: i32 count)

    TItemSetResult bsgGetSliceFromItem(1:TContainerKey rootID, 2: TItemKey fromKey, 3: i32 count)

    TItemSetResult bsgGetSliceR(1:TContainerKey rootID, 2: i32 fromIdx, 3: i32 count)

    TItemSetResult bsgGetSliceFromItemR(1:TContainerKey rootID, 2: TItemKey fromKey, 3: i32 count)

    TSplitBigSetResult splitBigSet(1:TContainerKey rootID, 2:TContainerKey brotherRootID=0, 3: i64 currentSize),

    TItemSetResult bsgRangeQuery(1:TContainerKey rootID, 2: TItemKey startKey, 3: TItemKey endKey),

    //init big set with multiple items.
    bool bsgBulkLoad(1:TContainerKey rootID, 2: TItemSet setData),

    TMultiPutItemResult bsgMultiPut(1: TContainerKey rootID, 2: TItemSet setData, 3: bool getAddedItems, 4: bool getReplacedItems ),

////////////////////////

    TBigSetGenericData getSetGenData(1:TMetaKey metaID),

    void putSetGenData(1:TMetaKey metaID, 2: TBigSetGenericData metadata),

    i64 getTotalCount(1:TContainerKey rootID),

    i64 removeAll(1:TContainerKey rootID)

}

typedef string TStringKey
struct TStringBigSetInfo{
    1: required TStringKey bigsetName,
    2: required TContainerKey bigsetID,
    3: optional i64 count, // total item    
}

struct TBigSetInfoResult{
    1: required TErrorCode error,
    2: optional TStringBigSetInfo info,
}

struct TCaSItem{
    1: TItemKey itemKey,
    2: TItemValue oldValue,
    3: TItemValue newValue,
}

service TStringBigSetKVService{

    TBigSetInfoResult createStringBigSet(1:TStringKey bsName),

    /*Get BigSet info*/
    TBigSetInfoResult getBigSetInfoByName(1:TStringKey bsName),

    //
    TBigSetInfoResult assignBigSetName(1: TStringKey bsName, 2: TContainerKey bigsetID),


    TPutItemResult bsPutItem(1:TStringKey bsName, 2:TItem item),

    //New function >> too complicated to implement, so I will do it later. You should try S2I64 or S2S later
    //TPutItemResult bsCasItem(1:TStringKey bsName, 2:TCasItem casItem), 
    //New function <<

    /*return true if item is existed in the list otherwise return false*/
    bool bsRemoveItem(1:TStringKey bsName, 2:TItemKey itemKey),

    TExistedResult bsExisted(1:TStringKey bsName, 2: TItemKey itemKey),

    TItemResult bsGetItem(1:TStringKey bsName, 2: TItemKey itemKey),

    TItemSetResult bsGetSlice(1:TStringKey bsName, 2: i32 fromPos, 3: i32 count)

    TItemSetResult bsGetSliceFromItem(1:TStringKey bsName, 2: TItemKey fromKey, 3: i32 count)

    TItemSetResult bsGetSliceR(1:TStringKey bsName, 2: i32 fromPos, 3: i32 count)

    TItemSetResult bsGetSliceFromItemR(1:TStringKey bsName, 2: TItemKey fromKey, 3: i32 count)

    TItemSetResult bsRangeQuery(1:TStringKey bsName, 2: TItemKey startKey, 3: TItemKey endKey),

    //init big set with multiple items.
    bool bsBulkLoad(1:TStringKey bsName, 2: TItemSet setData),

    TMultiPutItemResult bsMultiPut(1: TStringKey bsName, 2: TItemSet setData, 3: bool getAddedItems, 4: bool getReplacedItems ),

    i64 getTotalCount(1:TStringKey bsName),

    i64 removeAll(1:TStringKey bsName)

    i64 totalStringKeyCount(), 

    list<TStringKey> getListKey(1: i64 fromIndex, 2: i32 count),

    list<TStringKey> getListKeyFrom(1: TStringKey keyFrom, 2: i32 count), // keyFrom="" => get from start

}

//String key, big value.
service TBSBigValueService{
    
}


/* 
* BigSet with Int BigSetID key-value items
* This is a interface of a safer big set (a bit slower)
* Non-ProProfessional should use it instead of TBSGenericDataService directly
*/
service TIBSDataService{

    TPutItemResult putItem(1:TKey bigsetID, 2:TItem item),

    /*return true if item is existed in the list otherwise return false*/
    bool removeItem(1:TKey bigsetID, 2:TItemKey itemKey),

    TExistedResult existed(1:TKey bigsetID, 2: TItemKey itemKey),

    TItemResult getItem(1:TKey bigsetID, 2: TItemKey itemKey),

    TItemSetResult getSlice(1:TKey bigsetID, 2: i32 fromPos, 3: i32 count)

    TItemSetResult getSliceFromItem(1:TKey bigsetID, 2: TItemKey fromKey, 3: i32 count)

    TItemSetResult getSliceR(1:TKey bigsetID, 2: i32 fromPos, 3: i32 count)

    TItemSetResult getSliceFromItemR(1:TKey bigsetID, 2: TItemKey fromKey, 3: i32 count)

    TItemSetResult rangeQuery(1:TKey bigsetID, 2: TItemKey startKey, 3: TItemKey endKey),

    //init big set with multiple items.
    bool bulkLoad(1:TKey bigsetID, 2: TItemSet setData),

    TMultiPutItemResult multiPut(1: TKey bigsetID, 2: TItemSet setData, 3: bool getAddedItems, 4: bool getReplacedItems ),
    
    i64 getTotalCount(1:TKey bigsetID),

    i64 removeAll(1:TKey bigsetID)


}

/* Todo: Cassandra interface */

struct SplitJob{
    1: required TContainerKey rootID, 
    2: optional TNeedSplitInfo splitInfo    
}

service BSNotificationPool{
    
    void needSplit(1:TContainerKey rootID, 2: TNeedSplitInfo splitInfo)

    void splitInfoUpdated(1:TContainerKey rootID),

    SplitJob getJob(),

}

service TCluserOrdinatorService{
    oneway void removeCache(1: TContainerKey key),

    i32 put(1: binary key, 2: binary value),
    
}



