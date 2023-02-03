package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/darrenli6/hertz-examples/bizdemo/hertz_casbin/biz/model/casbin"
	"github.com/darrenli6/hertz-examples/bizdemo/hertz_casbin/pkg/consts"
	"github.com/golang-jwt/jwt/v4"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func If(condition bool, trueValue, falseValue interface{}) interface{} {
	if condition {
		return trueValue
	}
	return falseValue
}

// RFC3339ToNormalTime RFC3339 日期格式标准化
func RFC3339ToNormalTime(rfc3339 string) string {
	if len(rfc3339) < 19 || rfc3339 == "" || !strings.Contains(rfc3339, "T") {
		return rfc3339
	}
	return strings.Split(rfc3339, "T")[0] + " " + strings.Split(rfc3339, "T")[1][:8]
}

func GenerateToken(id uint, roles []casbin.Role, username string, second int) (string, error) {

	uc := consts.UserClaim{
		Id:       id,
		Username: username,
		Roles:    roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(second))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(consts.JwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func AnalyzeToken(token string) (*consts.UserClaim, error) {
	uc := new(consts.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(consts.JwtKey), nil
	})

	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}

	return uc, err
}

// httpRequest .
func httpRequest(url, method string, data, header []byte) ([]byte, error) {
	var err error
	reader := bytes.NewBuffer(data)
	request, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	// 处理 header
	if len(header) > 0 {
		headerMap := new(map[string]interface{})
		err = json.Unmarshal(header, headerMap)

		if err != nil {
			return nil, err
		}
		for k, v := range *headerMap {
			if k == "" || v == "" {
				continue
			}
			request.Header.Set(k, v.(string))
		}
	}
	request.SetBasicAuth(consts.EmqxKey, consts.EmqxSecret)

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBytes, nil
}

func HttpDelete(url string, data []byte, header ...byte) ([]byte, error) {
	return httpRequest(url, "DELETE", data, header)
}

func HttpPut(url string, data []byte, header ...byte) ([]byte, error) {
	return httpRequest(url, "PUT", data, header)
}

func HttpPost(url string, data []byte, header ...byte) ([]byte, error) {
	return httpRequest(url, "POST", data, header)
}

func HttpGet(url string, header ...byte) ([]byte, error) {
	return httpRequest(url, "GET", []byte{}, header)
}
