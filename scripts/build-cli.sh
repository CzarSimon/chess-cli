NAME=`cat package.json | jq .name -r`
VERSION=`cat package.json | jq .version -r`
COMMIT=`git rev-parse --short HEAD`

FULL_VERSION="v$VERSION ($COMMIT)"

sed s/'appVersion = ".*"'/"appVersion = \"$FULL_VERSION\""/g cmd/cli/main.go | gofmt > cmd/cli/main.tmp.go

cp main.go /tmp/nlp-main.go
mv cmd/cli/main.go "/tmp/$NAME-$VERSION-main.go"

go build -o chess cmd/cli/main.tmp.go

rm cmd/cli/main.tmp.go
mv "/tmp/$NAME-$VERSION-main.go" cmd/cli/main.go