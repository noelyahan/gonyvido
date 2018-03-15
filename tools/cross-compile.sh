HASH=""
VERSION="1.0.0"
RELEASE_PATH="releases"

cd ..
rm $RELEASE_PATH -r
mkdir $RELEASE_PATH

# Windows
echo "Windows 64"
GOOS=windows
GOARCH=amd64
OS_VERSION=win64
go build -o $RELEASE_PATH/gonyvido-$VERSION-$OS_VERSION/bin/gonyvido.exe
cp README.md LICENSE $RELEASE_PATH/gonyvido-$VERSION-$OS_VERSION
cd $RELEASE_PATH
tar -czvf gonyvido-$VERSION-$OS_VERSION.zip  gonyvido-$VERSION-$OS_VERSION
rm gonyvido-$VERSION-$OS_VERSION -r
cd ..

echo "Windows 32"
GOOS=windows
GOARCH=386
OS_VERSION=win32
go build -o $RELEASE_PATH/gonyvido-$VERSION-$OS_VERSION/bin/gonyvido.exe
cp README.md LICENSE $RELEASE_PATH/gonyvido-$VERSION-$OS_VERSION
cd $RELEASE_PATH
tar -czvf gonyvido-$VERSION-$OS_VERSION.zip  gonyvido-$VERSION-$OS_VERSION
rm gonyvido-$VERSION-$OS_VERSION -r
cd ..

# Linux
echo "Linux 64"
GOOS=linux
GOARCH=amd64
OS_VERSION=linux64
go build -o $RELEASE_PATH/gonyvido-$VERSION-$OS_VERSION/bin/gonyvido
cp README.md LICENSE $RELEASE_PATH/gonyvido-$VERSION-$OS_VERSION
cd $RELEASE_PATH
tar -czvf gonyvido-$VERSION-$OS_VERSION.tar.gz  gonyvido-$VERSION-$OS_VERSION
rm gonyvido-$VERSION-$OS_VERSION -r
cd ..


echo "Linux 32"
GOOS=linux
GOARCH=386
OS_VERSION=linux32
go build -o $RELEASE_PATH/gonyvido-$VERSION-$OS_VERSION/bin/gonyvido
cp README.md LICENSE $RELEASE_PATH/gonyvido-$VERSION-$OS_VERSION
cd $RELEASE_PATH
tar -czvf gonyvido-$VERSION-$OS_VERSION.tar.gz  gonyvido-$VERSION-$OS_VERSION
rm gonyvido-$VERSION-$OS_VERSION -r
cd ..


echo "Linux arm"
GOOS=linux
GOARCH=arm
OS_VERSION=linux-arm
go build -o $RELEASE_PATH/gonyvido-$VERSION-$OS_VERSION/bin/gonyvido
cp README.md LICENSE $RELEASE_PATH/gonyvido-$VERSION-$OS_VERSION
cd $RELEASE_PATH
tar -czvf gonyvido-$VERSION-$OS_VERSION.tar.gz  gonyvido-$VERSION-$OS_VERSION
rm gonyvido-$VERSION-$OS_VERSION -r
cd ..

# Mac
echo "OSX 64"
GOOS=darwin
GOARCH=amd64
OS_VERSION=osx
go build -o $RELEASE_PATH/gonyvido-$VERSION-$OS_VERSION/bin/gonyvido
cp README.md LICENSE $RELEASE_PATH/gonyvido-$VERSION-$OS_VERSION
cd $RELEASE_PATH
tar -czvf gonyvido-$VERSION-$OS_VERSION.tar.gz  gonyvido-$VERSION-$OS_VERSION
rm gonyvido-$VERSION-$OS_VERSION -r
cd ..