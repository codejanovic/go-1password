set -ex

mkdir bin
case "$TRAVIS_OS_NAME" in
	"osx")
		GOOS=darwin  GOARCH=amd64 go build -o bin/gordon-darwin-amd64
		cd bin
		tar czf ../gordon-$TRAVIS_BRANCH-darwin-amd64.tar.gz gordon-darwin-amd64
		;;
	"linux")
		GOOS=linux  GOARCH=amd64 go build -o bin/gordon-linux-amd64
		cd bin
		tar czf ../gordon-$TRAVIS_BRANCH-linux-amd64.tar.gz gordon-linux-amd64
		;;
esac