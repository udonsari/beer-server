package dto

type KakaoToken struct {
	AccessToken string `json:"access_token"`
}

type KakaoUser struct {
	ID         int64      `json:"id"`
	Properties Properties `json:"properties"`
}

type Properties struct {
	NickName       string `json:"nickname"`
	ProfileImage   string `json:"profile_image"`
	ThumbnailImage string `json:"thumbnail_image"`
}

type GetKakaoUserRequest struct {
	AccessToken string `query:"access_token" validate:"required"`
}
