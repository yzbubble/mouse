:: prepare
rd /s /q "dist"
md "dist"

:: build windows
SET GOOS=windows
SET GOARCH=amd64
go build -o dist/mouse-windows.exe .

:: build mac
SET GOOS=darwin
SET GOARCH=amd64
go build  -o dist/mouse-darwin .

:: build linux
SET GOOS=linux
SET GOARCH=amd64
go build -o dist/mouse-linux .
