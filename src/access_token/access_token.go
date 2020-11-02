package access_token

import "time"

const expirationTime = 24

type JWT struct {
	AccessToken string
	UserId      int64
	ClientId    int64
	Expires     int64
}

func GetNewJWT() JWT {
	return JWT{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (jwt JWT) IsExpired() bool {
	return time.Unix(jwt.Expires, 0).Before(time.Now().UTC())
}
