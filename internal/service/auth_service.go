package service

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"github.com/twinj/uuid"

	"github.com/go-redis/redis/v7"
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
	UserId     uint64
}

type AuthService struct {
	redis *redis.Client
}

func NewAuthService(redisClient *redis.Client) *AuthService {
	return &AuthService{redisClient}
}

func (s *AuthService) CreateToken(userid uint64) (*TokenDetails, error) {

	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(viper.GetString("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(viper.GetString("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

// Redis 에 uuid 저장
func (s *AuthService) CreateAuth(userid uint64, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := s.redis.Set(td.AccessUuid, strconv.Itoa(int(userid)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := s.redis.Set(td.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

// Redis 에서 uuid 삭제
func (s *AuthService) DeleteAuth(uuid string) (int64, error) {
	deleted, err := s.redis.Del(uuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

func (s *AuthService) ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")

	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func (s *AuthService) VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := s.ExtractToken(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("ACCESS_SECRET")), nil
	})
	fmt.Println("err:", err)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *AuthService) TokenValid(r *http.Request) error {
	token, err := s.VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func (s *AuthService) ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := s.VerifyToken(r)
	fmt.Println("token:", token)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
		}, nil
	}
	return nil, err
}

func (s *AuthService) FetchAuth(authD *AccessDetails) (uint64, error) {
	userid, err := s.redis.Get(authD.AccessUuid).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	return userID, nil
}
