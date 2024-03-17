package models

type User struct {
	GUID string
}

type UserData struct {
	GUID string  
	RefreshToken string
}

type Token struct{
	AccessToken string
	RefreshToken string
}

func (token *Token)NewToken(accessToken, refreshToken string) *Token{
	return &Token{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}
}