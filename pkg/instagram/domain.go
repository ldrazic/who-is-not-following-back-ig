package instagram

type User struct {
	HasAnonymousProfilePicture bool     `json:"has_anonymous_profile_picture"`
	FbidV2                     string   `json:"fbid_v2"`
	TextPostAppJoinerNumber    int      `json:"text_post_app_joiner_number"`
	PK                         string   `json:"pk"`
	PKID                       string   `json:"pk_id"`
	Username                   string   `json:"username"`
	FullName                   string   `json:"full_name"`
	IsPrivate                  bool     `json:"is_private"`
	IsVerified                 bool     `json:"is_verified"`
	ProfilePicID               string   `json:"profile_pic_id"`
	ProfilePicURL              string   `json:"profile_pic_url"`
	AccountBadges              []string `json:"account_badges"`
	IsPossibleScammer          bool     `json:"is_possible_scammer"`
	ThirdPartyDownloadsEnabled int      `json:"third_party_downloads_enabled"`
	IsPossibleBadActor         struct {
		IsPossibleScammer      bool `json:"is_possible_scammer"`
		IsPossibleImpersonator struct {
			IsUnconnectedImpersonator bool `json:"is_unconnected_impersonator"`
		} `json:"is_possible_impersonator"`
	} `json:"is_possible_bad_actor"`
	LatestReelMedia int64 `json:"latest_reel_media"`
	IsFavorite      bool  `json:"is_favorite"`
}

type FollowingResponse struct {
	Users                      []User `json:"users"`
	BigList                    bool   `json:"big_list"`
	PageSize                   int    `json:"page_size"`
	NextMaxID                  string `json:"next_max_id"`
	HasMore                    bool   `json:"has_more"`
	ShouldLimitListOfFollowers bool   `json:"should_limit_list_of_followers"`
	Status                     string `json:"status"`
}
