package utils

import (
	"crypto/rand"
	"fmt"
	"log"
	"time"
	"github.com/dgrijalva/jwt-go"
)

func CreateToken(uid, secret string) (string, error) {
        at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
                "uid":  uid,
                "exp":  time.Now().Add(time.Minute * 15).Unix(),
        })
        token, err := at.SignedString([]byte(secret))
        if err != nil {
                return "", err
        }
        return token, nil
}

func ParseToken(token string, secret string) (string, error) {
        claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
                return []byte(secret), nil
        })
        if err != nil {
                return "", err
        }
        return claim.Claims.(jwt.MapClaims)["uid"].(string), nil
}
// CreateUUID 生成UUID
func CreateUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}
	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}
