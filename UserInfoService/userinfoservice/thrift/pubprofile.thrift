namespace go openstars.pubprofile

struct ProfileData {
  1: string pubkey;
  2: string displayName;
  3: i64 dob;
  4: i64 gender;
  5: string introText;
  6: string avatar;
  7: string imgBackground;
  8: string phone;
  9: string education;
  10: string work;
  11: i64 relationship;
  12: string accommodation;
  13: string linkFB;
  14: string linkGGPlus;
  15: string linkInstagram;
  16: map<string,string> extend;
  17: list<string> image;
  18:i64 lastModified;
}

enum TErrorCode {
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2 ,
    EDataExisted = -3,
    EFailed = -4
}

struct TErrorCodeResult {
    1:i64 ecode
    2:string message
}

struct ResponseProfile {
  1:ProfileData profileData;
  2:TErrorCodeResult errorData;
}

struct ResponseBool {
  1:bool resp;
  2:TErrorCodeResult errorData;
}

service PubProfileService {
  ResponseProfile GetProfileByPubkey(1:string pubkey);
	ResponseProfile GetProfileByUID(1: i64 uid);
  ResponseBool SetProfileByPubkey(1: string pubkey,2: ProfileData profileSet);
  ResponseBool UpdateProfileByUID(1: i64 uid,2: ProfileData profileUpdate);
  ResponseBool UpdateProfileByPubkey(1: string pubkey,2: ProfileData profileUpdate);
}