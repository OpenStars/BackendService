/*
author: trungthanh
*/

namespace cpp core.auth.simplesession
namespace java com.auth.simplesession

typedef i64 TUID
typedef i64 TInternalSessionID
typedef string TUserRights //permission data of a user (in JSON)

typedef string TSessionKey

enum TSessionCode{
    ENoSession = 0,
    EFullRight = 1,
    EPartialRight,
    ERightExtend100=100,
    ERightExtend101,
    ERightExtend102,
    ERightExtend103,
    ERightExtend104,
    ERightExtend105,
    ERightExtend106,
    ERightExtend107,
    ERightExtend108,

}

struct TUserSessionInfo{
    1: TSessionCode code
    2: TUID uid,
    3: TUserRights permissions,
    4: string deviceInfo,
    5: string data,
    6: i64 expiredTime, //
    7: i32 version
}


enum TErrorCode{
    ESuccess = 0,
    EFailed = 1,
}

struct TSessionKeyResult{
    1: TErrorCode error,
    2: optional TSessionKey session
}


struct TUserResult{
    1: required TErrorCode error,
    2: optional TUserSessionInfo userInfo
}
    

service TSimpleSessionService {
    TUserResult getSession(1: TSessionKey sessionKey), 
    
}

service TSimpleSessionService_W extends TSimpleSessionService{
    TSessionKeyResult createSession(1: TUserSessionInfo userInfo),
    bool removeSession(1: TSessionKey sessionKey), 
}
