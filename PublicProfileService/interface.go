package PubProfileClient

import "github.com/OpenStars/BackendService/PublicProfileService/tpubprofileservice/thrift/gen-go/openstars/pubprofile"

type Client interface {
	GetProfileByUID(uid int64) (r *pubprofile.ProfileData, err error)
	GetProfileByPubkey(pubkey string) (r *pubprofile.ProfileData, err error)
	SetProfileByPubkey(pubkey string, profileUpdate *pubprofile.ProfileData) (r bool, err error)
	UpdateProfileByPubkey(pubkey string, profileUpdate *pubprofile.ProfileData) (r bool, err error)
	UpdateProfileByUID(uid int64, profileUpdate *pubprofile.ProfileData) (r bool, err error)
}
