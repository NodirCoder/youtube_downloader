@echo off

if /I not "%GOPATH%" == "" (
    echo Initilization...
    go mod init main
    go get github.com/kkdai/youtube/v2
    echo Done!
) else (
    echo Golang isn't installed!!
    echo Set GOPATH in your global environment!
    pause
)
