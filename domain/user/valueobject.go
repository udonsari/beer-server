package user

const KakaoOauthURL = "https://kauth.kakao.com/oauth/authorize"
const KakaoTokenURL = "https://kauth.kakao.com/oauth/token"
const KakaoUserURL = "https://kapi.kakao.com/v2/user/me"

// TODO Remove
// TODO 카카오 인증서버 Mock 만들어서 안귀찮게 하기
// TODO 인증 방식이 늘어날 경우, accessToken에 kakao 인지, 그 외 다른 방식인지 판단 가능하게 처리
const KakaoAppKey = "FILLIT"

type Token struct {
	AccessToken string `json:"access_token"`
}

type KakaoUser struct {
	ID         int64      `json:"id"`
	Properties Properties `json:"properties"`
}
