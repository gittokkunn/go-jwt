package jwt

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type UserInfo struct {
	ID string `json: "id"`
	Name string `json:"name"`
	Pass string `json:"pass"`
	jwt.StandardClaims
}

// ID発行済みかを判定してHOME画面に飛ばす
func LoginMyPage(c *gin.Context) {
	// Cookieに入っているJWTの取得
	JWT, err := c.Cookie("JWT")
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/signup")
		return
	}
	if JWT != "" {
		token, err := decodeJWT(JWT)
		if err != nil {
			panic(err)
		}
		userID := validToken(*token, "tokkunn")
		fmt.Println(userID)
		c.HTML(http.StatusOK, "mypage.html", nil)
		return
	}else {
		return
	}
}

func SignUp(c *gin.Context) {
	// Cookieに入っているJWTの取得
	JWT, err := c.Cookie("JWT")
	if err != nil {
		c.HTML(http.StatusOK, "signup.html", nil)
		return
	}
	// トークンが発行されていないか判定
	if JWT != "" {
		c.Redirect(http.StatusMovedPermanently, "/mypage")
		return
	} else {
		return
	}
}

func AddUser(c *gin.Context) {
	// Cookieに入っているJWTの取得
	JWT, err := c.Cookie("JWT")
	fmt.Println(JWT)
	if err != nil {
		var user = &UserInfo{
			ID:   generateRandomID(),
			//Name: "tokkunn",
			Name: "tokkun",
			Pass: "password",
		}
		newJWT := GenerateJWT(user)
		c.SetCookie("JWT", newJWT,1000 * 60 * 60 * 60 * 24 * 7, "/", "localhost", false, true)
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}
	// トークンが発行されていない時の処理
	if JWT != "" {
		c.Redirect(http.StatusMovedPermanently, "/mypage")
		return
	} else {
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}
}

func GenerateJWT(info *UserInfo) string {
	jwtSecret := "dfasdfslkjlds" // Secret
	// ユーザ情報からJWTを生成
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), info)
	tokenstring, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Fatalln(err)
	}
	return tokenstring
}


// ランダム文字列生成
func generateRandomID() string {
	var n uint64
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return strconv.FormatUint(n, 36)
}

func decodeJWT(jwtk string) (*jwt.Token, error) {
	jwtSecret := "dfasdfslkjlds"
	token, err := jwt.Parse(jwtk, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func validToken(token jwt.Token, userid string) interface{} {
	claims := token.Claims.(jwt.MapClaims)
	if claims["name"] != "tokkunn" {
		panic("ユーザが違います")
	}
	return claims["name"]
}