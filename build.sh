#！/usr/bin/env sh
BINARY=program3
VERSION=$(git describe --abbrev=0 --tags)
FULLVERSION=$(git describe --tags)
BUILD=$(date +%FT%T%z)

go build -v -ldflags "-w -s -X main.Version=${VERSION} -X main.FullVersion=${FULLVERSION} -X main.Build=${BUILD}" ${LDFLAGS} -o ${BINARY}
