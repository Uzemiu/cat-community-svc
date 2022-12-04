// Code generated by goctl. DO NOT EDIT.
package types

type User struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	AvatarUrl      string `json:"avatarUrl"`
	PrivilegeLevel int32  `json:"privilegeLevel"`
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
	UserId    string `json:"userId"`
	Nickname  string `json:"nickname"`
	AvatarUrl string `json:"avatarUrl"`
}

type Cat struct {
	Id           string   `json:"id"`
	CreateAt     string   `json:"createAt"`
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
	Photos       []string `json:"photos"`
}

type CatPreview struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	AvatarUrl   string `json:"avatarUrl"`
	IsCollected bool   `json:"isCollected"`
}

type GetCatPreviewsReq struct {
	CommunityId string `json:"communityId"`
}

type GetCatPreviewsResp struct {
	Status
	Cats []CatPreview `json:"cats"`
}

type GetCatDetailReq struct {
	Id string `json:"campusId"`
}

type GetCatDetailResp struct {
	Status
	Cat Cat `json:"cat"`
}

type Moment struct {
}

type MomentPreview struct {
	Id       string      `json:"id"`
	CreateAt string      `json:"createAt"`
	CatId    string      `json:"catId"`
	CoverUrl string      `json:"coverUrl"`
	Title    string      `json:"title"`
	User     UserPreview `json:"user"`
	Likes    int32       `json:"likes"`
}

type GetMomentPreviewsReq struct {
	CommunityId string `json:"communityId"`
}

type GetMomentPreviewsResp struct {
	Status
	Moments []MomentPreview `json:"moments"`
}

type GetMomentDetailReq struct {
	Id string `json:"id"`
}

type GetMomentDetailResp struct {
	Status
	Moment Moment `json:"moment"`
}

type Comment struct {
	Id       string      `json:"id"`
	CreateAt string      `json:"createAt"`
	Comment  string      `json:"comment"`
	User     UserPreview `json:"user"`
}

type ListCommentReq struct {
	Id string `json:"id"`
}

type ListCommentResp struct {
	Status
	Comments []Comment `json:"comments"`
}

type PostCommentReq struct {
	Comment string `json:"comment"`
}

type PostCommentResp struct {
	Status
}

type News struct {
	Id       string `json:"id"`
	CreateAt string `json:"createAt"`
	ImageUrl string `json:"imageUrl"`
	LinkUrl  string `json:"linkUrl"`
	LinkType int32  `json:"linkType"`
}

type Admin struct {
	Id       string      `json:"id"`
	CreateAt string      `json:"createAt"`
	Name     string      `json:"name"`
	Phone    string      `json:"phone"`
	User     UserPreview `json:"user"`
	Wechat   string      `json:"wechat"`
}

type GetNewsResp struct {
	Status
	News []News `json:"news"`
}

type GetAdminsResp struct {
	Status
	Admins []Admin `json:"admins"`
}