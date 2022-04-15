# sshw

**仅添加支持google authenticator令牌**
> 正版连接 <https://github.com/yinheli/sshw>

## install sshw command

1. Install command
    ```
    go install github.com/Ubbo-Sathla/sshw/cmd/sshw@latest
    ```
2. Config example

    ```yaml
    - { name: dev server fully configured, user: appuser, host: 192.168.8.35, port: 22, password: 123456 ,mfa: "otpauth://totp/Ubbo-Sathla?algorithm=SHA1&digits=6&issuer=13123&period=30&secret=123456" }
    ```

## otp command

1. Install command

    ```
    go install github.com/Ubbo-Sathla/sshw/cmd/otp@master
    ```

2. Command example

    ```bash
    otp -mfa "otpauth://totp/example.com:ubbo-sathla@example.com?algorithm=SHA1&digits=6&issuer=example.com&period=30&secret=7F2TQQA2OVO47AQJ6V32K2HXFNB5I77F"
    ```
