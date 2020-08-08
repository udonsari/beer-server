package user

type User struct {
	ID         int64
	ExternalID string
	Properties
}

type Properties struct {
	NickName       string
	ProfileImage   string
	ThumbnailImage string
}
