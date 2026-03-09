package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"sort"
	"strings"

	"bingo-backend/internal/config"
)

func VerifyTelegram(initData string) bool {

	values, _ := url.ParseQuery(initData)

	hash := values.Get("hash")
	values.Del("hash")

	var dataCheck []string

	for k, v := range values {
		dataCheck = append(dataCheck, k+"="+v[0])
	}

	sort.Strings(dataCheck)

	dataCheckString := strings.Join(dataCheck, "\n")

	secret := sha256.Sum256([]byte(config.App.TelegramToken))

	h := hmac.New(sha256.New, secret[:])
	h.Write([]byte(dataCheckString))

	calculated := hex.EncodeToString(h.Sum(nil))

	return calculated == hash
}