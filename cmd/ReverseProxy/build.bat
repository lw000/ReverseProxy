cd ../../../../../
set GOPATH=%cd%
cd src/demo/ReverseProxy/cmd/ReverseProxy
set GOARCH=amd64
set GOOS=windows
go build -v -ldflags="-s -w"