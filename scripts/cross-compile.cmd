SET HASH=""
SET VERSION="1.0.0"
SET RELEASE_PATH="releases"

cd ..
rm %RELEASE_PATH% -r
mkdir %RELEASE_PATH%

# Windows
echo "Windows 64"
SET GOOS=windows
SET GOARCH=amd64
SET OS_VERSION=win64
go build -o %RELEASE_PATH%/gonyvido-%VERSION%-%OS_VERSION%/bin/gonyvido.exe
cp README.md LICENSE %RELEASE_PATH%/gonyvido-%VERSION%-%OS_VERSION%
cd %RELEASE_PATH%
rar a -r  gonyvido-%VERSION%-%OS_VERSION%.zip  gonyvido-%VERSION%-%OS_VERSION%
rm gonyvido-%VERSION%-%OS_VERSION% -r
cd ..

echo "Windows 32"
SET GOOS=windows
SET GOARCH=386
SET OS_VERSION=win32
go build -o %RELEASE_PATH%/gonyvido-%VERSION%-%OS_VERSION%/bin/gonyvido.exe
cp README.md LICENSE %RELEASE_PATH%/gonyvido-%VERSION%-%OS_VERSION%
cd %RELEASE_PATH%
rar a -r  gonyvido-%VERSION%-%OS_VERSION%.zip  gonyvido-%VERSION%-%OS_VERSION%
rm gonyvido-%VERSION%-%OS_VERSION% -r
cd ..

# Linux
echo "Linux 64"
SET GOOS=linux
SET GOARCH=amd64
SET OS_VERSION=linux64
go build -o %RELEASE_PATH%/gonyvido-%VERSION%-%OS_VERSION%/bin/gonyvido
cp README.md LICENSE %RELEASE_PATH%/gonyvido-%VERSION%-%OS_VERSION%
cd %RELEASE_PATH%
rar a -r  gonyvido-%VERSION%-%OS_VERSION%.zip  gonyvido-%VERSION%-%OS_VERSION%
rm gonyvido-%VERSION%-%OS_VERSION% -r
cd ..


echo "Linux 32"
SET GOOS=linux
SET GOARCH=386
SET OS_VERSION=linux32
go build -o %RELEASE_PATH%/gonyvido-%VERSION%-%OS_VERSION%/bin/gonyvido
cp README.md LICENSE %RELEASE_PATH%/gonyvido-%VERSION%-%OS_VERSION%
cd %RELEASE_PATH%
rar a -r  gonyvido-%VERSION%-%OS_VERSION%.zip  gonyvido-%VERSION%-%OS_VERSION%
rm gonyvido-%VERSION%-%OS_VERSION% -r
cd ..


echo "Linux arm"
SET GOOS=linux
SET GOARCH=arm
SET OS_VERSION=linux-arm
go build -o %RELEASE_PATH%/gonyvido-%VERSION%-%OS_VERSION%/bin/gonyvido
cp README.md LICENSE %RELEASE_PATH%/gonyvido-%VERSION%-%OS_VERSION%
cd %RELEASE_PATH%
rar a -r  gonyvido-%VERSION%-%OS_VERSION%.zip  gonyvido-%VERSION%-%OS_VERSION%
rm gonyvido-%VERSION%-%OS_VERSION% -r
cd ..

# Mac
echo "OSX 64"
SET GOOS=darwin
SET GOARCH=amd64
SET OS_VERSION=osx
go build -o %RELEASE_PATH%/gonyvido-%VERSION%-%OS_VERSION%/bin/gonyvido
cp README.md LICENSE %RELEASE_PATH%/gonyvido-%VERSION%-%OS_VERSION%
cd %RELEASE_PATH%
rar a -r  gonyvido-%VERSION%-%OS_VERSION%.zip  gonyvido-%VERSION%-%OS_VERSION%
rm gonyvido-%VERSION%-%OS_VERSION% -r
cd ..