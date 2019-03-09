cd ../../../../../
set GOPATH=%cd%
cd src/demo/ReverseProxy/cmd/ginsrv
set GOARCH=amd64
set GOOS=windows
go build