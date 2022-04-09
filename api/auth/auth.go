package auth

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"ewallet/config"
	"ewallet/storage/repo"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Auth struct {
	Cfg  config.Config
	Repo repo.Repo
}

func (auth Auth) Auth(w http.ResponseWriter, r *http.Request) {
	if auth.isPermitted(r.URL.Path) {
		return
	}
	if strings.Contains(r.URL.Path, "/") {
		return
	}
	userId := r.Header.Get("X-Userid")
	hashSum := r.Header.Get("X-Digest")

	// длина hmac-sha1 20 байт, не менее не более
	if len([]byte(hashSum)) != 20 {
		AbortWithStatus(401, w)
		return
	}

	if userId == "" {
		AbortWithStatus(401, w)
		return
	}

	fmt.Println(hashSum)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	h := hmac.New(sha1.New, nil)

	h.Write(body)

	hashBody := hex.EncodeToString(h.Sum(nil))
	if hashBody != hashSum {
		AbortWithStatus(401, w)
	}

	isExists, err := auth.Repo.CheckUserById(userId)
	if err != nil {
		AbortWithStatus(401, w)
		return
	}
	if !isExists {
		AbortWithStatus(401, w)
		return
	}
}

func (auth Auth) HashBody(body []byte) string {
	h := hmac.New(sha1.New, nil)

	if _, err := h.Write(body); err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}

func (a Auth) isPermitted(path string) bool {
	for _, v := range urls {
		if path == v {
			return true
		}
	}
	return false
}

func AbortWithStatus(code int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
}
