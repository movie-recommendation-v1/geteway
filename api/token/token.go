package token

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/movie-recommendation-v1/geteway/api/config"
	pb "github.com/movie-recommendation-v1/geteway/genproto/userservice"
	"github.com/spf13/cast"
	"log"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

type JWTHandler struct {
	Sub        string
	Exp        string
	Iat        string
	Role       string
	SigningKey string
	Token      string
}

//type Tokens struct {
//	AccessToken string
//
//	RefreshToken string
//}

func ValidateToken(tokenStr string) (bool, error) {
	log.Println("token from heaf=der ", tokenStr)
	_, err := ExtractClaim(tokenStr)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractClaim(tokenStr string) (jwt.MapClaims, error) {
	cfg := config.Load()
	var (
		token *jwt.Token
		err   error
	)

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.TOKENKEY), nil
	}

	token, err = jwt.Parse(tokenStr, keyFunc)

	if err != nil {

		return nil, err

	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !(ok && token.Valid) {

		return nil, err

	}

	return claims, nil

}

func (jwtHandler *JWTHandler) ExtractClaims() (jwt.MapClaims, error) {

	token, err := jwt.Parse(jwtHandler.Token, func(t *jwt.Token) (interface{}, error) {

		return []byte(jwtHandler.SigningKey), nil

	})

	if err != nil {

		return nil, err

	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !(ok && token.Valid) {

		slog.Error("invalid jwt token")

		return nil, err

	}

	return claims, nil

}

func GetIdFromToken(r *http.Request) (string, int) {

	var softToken string

	token := r.Header.Get("Authorization")

	if token == "" {

		return "unauthorized", http.StatusUnauthorized

	} else if strings.Contains(token, "Bearer") {

		softToken = strings.TrimPrefix(token, "Bearer ")

	} else {

		softToken = token

	}

	claims, err := ExtractClaim(softToken)

	if err != nil {

		return "unauthorized", http.StatusUnauthorized

	}

	return cast.ToString(claims["username"]), 0

}

type Tokens struct {
	AccessToken  string
	Time         string
	RefreshToken string
}

func GenereteJWTTokenForUser(user *pb.LoginRes) *Tokens {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	cfg := config.Load()
	exp := time.Now().Add(60 * time.Minute)

	claims := accessToken.Claims.(jwt.MapClaims)
	claims["name"] = user.UserRes.Name
	claims["email"] = user.UserRes.Email
	claims["role"] = user.UserRes.Role
	claims["user_id"] = user.UserRes.Id
	claims["iat"] = time.Now().Unix()
	claims["exp"] = exp.Unix()
	access, err := accessToken.SignedString([]byte(cfg.TOKENKEY))
	if err != nil {
		fmt.Println("error while genereting access token : ", err)
	}

	rftclaims := refreshToken.Claims.(jwt.MapClaims)
	claims["name"] = user.UserRes.Name
	claims["email"] = user.UserRes.Email
	claims["role"] = user.UserRes.Role
	claims["user_id"] = user.UserRes.Id
	rftclaims["iat"] = time.Now().Unix()
	rftclaims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	refresh, err := refreshToken.SignedString([]byte(cfg.TOKENKEY))
	if err != nil {
		fmt.Println("error while genereting refresh token : ", err)
	}
	t := fmt.Sprint(exp)
	return &Tokens{
		AccessToken:  access,
		Time:         t[:19],
		RefreshToken: refresh,
	}
}
