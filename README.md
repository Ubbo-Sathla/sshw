# sshw

**仅添加支持google authenticator令牌**
> 正版连接 <https://github.com/yinheli/sshw>

## install

use `go install`

```
go install github.com/Ubbo-Sathla/sshw/cmd/sshw@v1.1.1
```

## config example

```yaml
- { name: dev server fully configured, user: appuser, host: 192.168.8.35, port: 22, password: 123456 ,mfa: "otpauth://totp/Ubbo-Sathla?algorithm=SHA1&digits=6&issuer=13123&period=30&secret=123456" }
```
