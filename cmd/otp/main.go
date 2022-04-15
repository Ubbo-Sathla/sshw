package main

import (
	"flag"
	"fmt"
	"github.com/Ubbo-Sathla/sshw"
)

var mfa string

func init() {
	flag.StringVar(&mfa, "mfa", "", "otpauth://totp/example.com:ubbo-sathla@example.com?algorithm=SHA1&digits=6&issuer=example.com&period=30&secret=YPHWG4MTELOL4EOVWH3EBWL3XAG4XAKY")
	flag.Parse()
}
func main() {
	fmt.Print(sshw.GeneratePassCode(mfa))
}
