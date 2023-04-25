package gkit

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type AccessDetails struct {
	AccessUuid string
	UserUuid   string
}

func CreateToken(userUuid string, atSecret, rtSecret string, atExpires, rtExpires int64) (TokenDetails, error) {
	td := TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * time.Duration(atExpires)).Unix()
	td.AccessUuid = uuid.New().String()

	td.RtExpires = time.Now().Add(time.Minute * time.Duration(rtExpires)).Unix()
	td.RefreshUuid = uuid.New().String()

	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_uuid"] = userUuid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(atSecret))
	if err != nil {
		return td, err
	}

	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_uuid"] = userUuid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(rtSecret))
	if err != nil {
		return td, err
	}

	return td, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(tokenString string, secret string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(tokenString string, secret string) error {
	token, err := VerifyToken(tokenString, secret)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return err
	}

	return nil
}

func ExtractTokenMetadata(tokenString string, secret string) (AccessDetails, error) {
	td := AccessDetails{}

	token, err := VerifyToken(tokenString, secret)
	if err != nil {
		return td, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return td, err
		}

		userUuid := claims["user_uuid"].(string)
		td.AccessUuid = accessUuid
		td.UserUuid = userUuid

		return td, nil
	}

	return td, err
}

func ExtractRefreshTokenMetadata(tokenString string, secret string) (AccessDetails, error) {
	td := AccessDetails{}

	token, err := VerifyToken(tokenString, secret)
	if err != nil {
		return td, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string)
		if !ok {
			return td, err
		}

		userUuid := claims["user_uuid"].(string)
		td.AccessUuid = refreshUuid
		td.UserUuid = userUuid

		return td, nil
	}

	return td, err
}

func TokenAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := ExtractToken(c.Request)
		err := TokenValid(tokenString, secret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}
