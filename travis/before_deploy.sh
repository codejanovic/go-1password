set -ex

mkdir bin
case "$TRAVIS_OS_NAME" in
	"osx")
		GOOS=darwin  GOARCH=amd64 go build -o bin/go-1password-darwin-amd64
		cd bin
		tar czf ../go-1password-$TRAVIS_BRANCH-darwin-amd64.tar.gz go-1password-darwin-amd64
		;;
	"linux")
		GOOS=linux  GOARCH=amd64 go build -o bin/go-1password-linux-amd64
		cd bin
		tar czf ../go-1password-$TRAVIS_BRANCH-linux-amd64.tar.gz go-1password-linux-amd64
		;;
esac