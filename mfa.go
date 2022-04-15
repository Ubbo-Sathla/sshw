package sshw

import (
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

func (n *Node) RegexpMfa(q string) bool {
	mfa, err := regexp.MatchString(`(?i)mfa`, q)
	if err != nil {
		return false
	}
	return mfa
}

func GeneratePassCode(mfa string) string {
	u, _ := url.Parse(mfa)
	q := u.Query()

	secret := q.Get("secret")
	period, _ := strconv.Atoi(q.Get("period"))
	digits, _ := strconv.Atoi(q.Get("digits"))
	algorithm, _ := strconv.Atoi(q.Get("algorithm"))

	passcode, err := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
		Period:    uint(period),
		Skew:      1,
		Digits:    otp.Digits(digits),
		Algorithm: otp.Algorithm(algorithm),
	})
	if err != nil {
		return ""
	}
	return passcode
}
