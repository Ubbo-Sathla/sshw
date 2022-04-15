package sshw

import (
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"testing"
)

func TestMfa(t *testing.T) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "example.com",
		AccountName: "ubbo-sathla@example.com",
		Period:      30,
		Digits:      otp.DigitsSix,
		Algorithm:   otp.AlgorithmSHA1,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(`mfa url: `, key.URL())
	t.Log(`mfa code: `, GeneratePassCode(key.URL()))
}
