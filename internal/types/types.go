// Code generated by goctl. DO NOT EDIT.
package types

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatarUrl"`
}

type SignInReq struct {
	AuthType string   `json:"authType,options=phone|email|wechat"`
	AuthId   string   `json:"authId"`
	Password string   `json:"password,optional"`
	Params   []string `json:"params,optional"`
}

type SignInResp struct {
	Status
	UserId       string `json:"userId"`
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
}

type SetPasswordReq struct {
	Password string `json:"password"`
}

type SetPasswordResp struct {
	Status
}

type SendVerifyCodeReq struct {
	AuthType string `json:"authType,options=phone|email"`
	AuthId   string `json:"authId"`
}

type SendVerifyCodeResp struct {
	Status
}

type Status struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type UserPreview struct {
	Id        string `json:"id"`
	Nickname  string `json:"nickname"`
	AvatarUrl string `json:"avatarUrl"`
}

type Post struct {
	Id          string      `json:"id"`
	CreateAt    int64       `json:"createAt"`
	IsAnonymous bool        `json:"isAnonymous"`
	Title       string      `json:"title"`
	Text        string      `json:"text"`
	CoverUrl    string      `json:"coverUrl"`
	Tags        []string    `json:"tags"`
	Likes       int64       `json:"likes"`
	Comments    int64       `json:"comments"`
	User        UserPreview `json:"user"`
}

type GetPostPreviewsReq struct {
	Page int64 `form:"page"`
}

type GetPostPreviewsResp struct {
	Status
	Posts []Post `json:"posts"`
	Total int64  `json:"total"`
}

type GetPostDetailReq struct {
	PostId string `form:"postId"`
}

type GetPostDetailResp struct {
	Post Post `json:"post"`
	Status
}

type NewPostReq struct {
	Id          string   `json:"id,optional"`
	IsAnonymous bool     `json:"isAnonymous"`
	Title       string   `json:"title"`
	Text        string   `json:"text"`
	CoverUrl    string   `json:"coverUrl,optional"`
	Tags        []string `json:"tags"`
}

type NewPostResp struct {
	PostId string `json:"postId"`
	Status
}

type DeletePostReq struct {
	Id string `json:"id"`
}

type DeletePostResp struct {
	Status
}

type SearchPostReq struct {
	Keyword string `form:"keyword"`
	Page    int64  `form:"page"`
}

type SearchPostResp struct {
	Status
	Posts []Post `json:"posts"`
	Total int64  `json:"total"`
}

type Cat struct {
	Id           string   `json:"id"`
	CreateAt     int64    `json:"createAt"`
	Age          string   `json:"age"`
	CommunityId  string   `json:"communityId"`
	Color        string   `json:"color"`
	Details      string   `json:"details"`
	Name         string   `json:"name"`
	Popularity   int64    `json:"popularity"`
	Sex          string   `json:"sex"`
	Status       int32    `json:"status"`
	Area         string   `json:"area"`
	IsSnipped    bool     `json:"isSnipped"`
	IsSterilized bool     `json:"isSterilized"`
	Avatars      []string `json:"avatars"`
}

type CatPreview struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Area        string `json:"area"`
	Color       string `json:"color"`
	AvatarUrl   string `json:"avatarUrl"`
	IsCollected bool   `json:"isCollected"`
}

type GetCatPreviewsReq struct {
	CommunityId string `form:"communityId"`
	Page        int64  `form:"page"`
}

type GetCatPreviewsResp struct {
	Status
	Cats  []CatPreview `json:"cats"`
	Total int64        `json:"total"`
}

type GetCatDetailReq struct {
	CatId string `form:"catId"`
}

type GetCatDetailResp struct {
	Status
	Cat Cat `json:"cat"`
}

type DeleteCatReq struct {
	CatId string `json:"catId"`
}

type DeleteCatResp struct {
	Status
}

type NewCatReq struct {
	Id           string   `json:"id,optional"`
	Age          string   `json:"age"`
	CommunityId  string   `json:"communityId"`
	Color        string   `json:"color"`
	Details      string   `json:"details"`
	Name         string   `json:"name"`
	Sex          string   `json:"sex"`
	Area         string   `json:"area"`
	IsSnipped    bool     `json:"isSnipped"`
	IsSterilized bool     `json:"isSterilized"`
	Avatars      []string `json:"avatars"`
}

type NewCatResp struct {
	Status
	CatId string `json:"catId"`
}

type SearchCatReq struct {
	CommunityId string `form:"communityId"`
	Keyword     string `form:"keyword"`
	Page        int64  `form:"page"`
}

type SearchCatResp struct {
	Status
	Cats  []CatPreview `json:"cats"`
	Total int64        `json:"total"`
}

type Moment struct {
	Id          string      `json:"id"`
	CreateAt    int64       `json:"createAt"`
	CatId       string      `json:"catId,optional"`
	Photos      []string    `json:"photos"`
	Title       string      `json:"title"`
	Text        string      `json:"text"`
	User        UserPreview `json:"user"`
	CommunityId string      `json:"communityId"`
}

type GetMomentPreviewsReq struct {
	CommunityId string `form:"communityId"`
	Page        int64  `form:"page"`
}

type GetMomentPreviewsResp struct {
	Status
	Moments []Moment `json:"moments"`
	Total   int64    `json:"total"`
}

type GetMomentDetailReq struct {
	MomentId string `form:"momentId"`
}

type GetMomentDetailResp struct {
	Status
	Moment Moment `json:"moment"`
}

type DeleteMomentReq struct {
	MomentId string `json:"momentId"`
}

type DeleteMomentResp struct {
	Status
}

type NewMomentReq struct {
	Id          string   `json:"id,optional"`
	Title       string   `json:"title"`
	CatId       string   `json:"catId,optional"`
	Text        string   `json:"text"`
	Photos      []string `json:"photos"`
	CommunityId string   `json:"communityId"`
}

type NewMomentResp struct {
	MomentId string `json:"momentId"`
	Status
}

type SearchMomentReq struct {
	CommunityId string `form:"communityId"`
	Keyword     string `form:"keyword"`
	Page        int64  `form:"page"`
}

type SearchMomentResp struct {
	Status
	Moments []Moment `json:"moments"`
	Total   int64    `json:"total"`
}

type Comment struct {
	Id        string      `json:"id"`
	CreateAt  int64       `json:"createAt"`
	Text      string      `json:"text"`
	User      UserPreview `json:"user"`
	Comments  int64       `json:"comments"`
	ReplyName string      `json:"replyName,optional"`
}

type NewCommentReq struct {
	Text  string `json:"text"`
	Id    string `json:"id,optional"`
	Scope string `json:"scope"`
}

type NewCommentResp struct {
	Status
}

type GetCommentsReq struct {
	Scope string `form:"scope"`
	Page  int64  `form:"page"`
	Id    string `form:"id"`
}

type GetCommentsResp struct {
	Status
	Comments []Comment `json:"comments"`
	Total    int64     `json:"total"`
}

type News struct {
	Id          string `json:"id"`
	CreateAt    int64  `json:"createAt"`
	CommunityId string `json:"communityId"`
	ImageUrl    string `json:"imageUrl"`
	LinkUrl     string `json:"linkUrl"`
	Type        string `json:"type"`
}

type Admin struct {
	Id          string      `json:"id"`
	CreateAt    int64       `json:"createAt"`
	CommunityId string      `json:"communityId"`
	Name        string      `json:"name"`
	Phone       string      `json:"phone"`
	User        UserPreview `json:"user"`
	Wechat      string      `json:"wechat"`
}

type Notice struct {
	Id       string `json:"id"`
	Text     string `json:"text"`
	CreateAt int64  `json:"createAt"`
}

type GetNewsReq struct {
	CommunityId string `form:"communityId"`
}

type GetNewsResp struct {
	Status
	News []News `json:"news"`
}

type GetAdminsReq struct {
	CommunityId string `form:"communityId"`
}

type GetAdminsResp struct {
	Status
	Admins []Admin `json:"admins"`
}

type GetNoticesReq struct {
	CommunityId string `form:"communityId"`
}

type GetNoticesResp struct {
	Status
	Notices []Notice `json:"notices"`
}

type NewNoticeReq struct {
	Id          string `json:"id,optional"`
	CommunityId string `json:"communityId,optional"`
	Text        string `json:"text"`
}

type NewNoticeResp struct {
	Status
	NoticeId string `json:"noticeId"`
}

type NewNewsReq struct {
	Id          string `json:"id,optional"`
	CommunityId string `json:"communityId,optional"`
	ImageUrl    string `json:"imageUrl"`
	LinkUrl     string `json:"linkUrl"`
	Type        string `json:"type"`
}

type NewNewsResp struct {
	Status
	NewId string `json:"newId"`
}

type DeleteNoticeReq struct {
	Id string `json:"id"`
}

type DeleteNoticeResp struct {
	Status
}

type DeleteNewsReq struct {
	Id string `json:"id"`
}

type DeleteNewsResp struct {
	Status
}

type Community struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ParentId string `json:"parentId"`
}

type ListCommunityReq struct {
}

type ListCommunityResp struct {
	Communities []Community `json:"communities"`
	Status
}

type NewCommunityReq struct {
	Id       string `json:"id"`
	Name     string `json:"name""`
	ParentId string `json:"parentId"`
}

type NewCommunityResp struct {
	NewId string `json:"newId"`
	Status
}

type DeleteCommunityReq struct {
	Id string `json:"id"`
}

type DeleteCommunityResp struct {
	Status
}

type DoLikeReq struct {
	TargetId   string `json:"targetId"`
	TargetType int64  `json:"targetType"`
}

type DoLikeResp struct {
	Status
}

type GetUserLikedReq struct {
	TargetId   string `form:"targetId"`
	TargetType int64  `form:"targetType"`
}

type GetUserLikedResp struct {
	Status
	Liked bool `json:"liked"`
}

type GetLikedCountReq struct {
	TargetId   string `form:"targetId"`
	TargetType int64  `form:"targetType"`
}

type GetLikedCountResp struct {
	Status
	Count int64 `json:"count"`
}

type GetUserInfoReq struct {
}

type GetUserInfoResp struct {
	Status
	User UserPreview `json:"user"`
}

type UpdateUserInfoReq struct {
	AvatarUrl string `json:"avatarUrl,optional"`
	Nickname  string `json:"nickname,optional"`
}

type UpdateUserInfoResp struct {
	Status
}

type ApplySignedUrlReq struct {
	Prefix string `json:"prefix,optional"`
	Suffix string `json:"suffix,optional"`
}

type ApplySignedUrlResp struct {
	Status
	Url          string `json:"url"`
	SessionToken string `json:"sessionToken"`
}
