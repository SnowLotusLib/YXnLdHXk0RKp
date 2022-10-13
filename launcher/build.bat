SET GOOS=linux
go build -ldflags="-w -s" -o  ../build/launcher_linux
SET GOOS=darwin
go build -ldflags="-w -s" -o ../build/launcher_darwin
SET GOOS=windows
go build -ldflags="-w -s" -o ../build/launcher.exe