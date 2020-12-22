NAME=`cat package.json | jq .name -r`
VERSION=`cat package.json | jq .version -r`
COMMIT=`git rev-parse --short HEAD`

FULL_VERSION="v$VERSION ($COMMIT)"

sed s/'appVersion = ".*"'/"appVersion = \"$FULL_VERSION\""/g internal/cli/cli.go | gofmt > internal/cli/cli.tmp.go
mv internal/cli/cli.go "/tmp/$NAME-$VERSION-main.go"

go build -o chess cmd/cli/main.go

rm internal/cli/cli.tmp.go
mv "/tmp/$NAME-$VERSION-main.go" internal/cli/cli.go