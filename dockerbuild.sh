#!/bin/bash

set -e

# Consider the upgrade of image `maching/ubuntu-2004:202010-01` for circleci, we select
# the first in GOPATH in order one by one, and treat it as LOCAL_GOPATH if it exists in the system.
export LOCAL_GOPATH=`echo $GOPATH | awk -F: '{print $1}'`
IFS=':' read -r -a gopath_array <<< "$GOPATH"
for gopath in "${gopath_array[@]}"
do
  [ -d "$gopath/pkg/mod/cache" ] && export LOCAL_GOPATH=$gopath && break
done

export UNAME=$(uname -m)

echo "SSH key path: $AT_SSH_PRIVATE_KEY_PATH"
[ -z "$SSH_PRIVATE_KEY" ] && export SSH_PRIVATE_KEY="$(cat ~/${AT_SSH_PRIVATE_KEY_PATH} | base64)"

trap 'rm -rf .dockerbuild/gopath || true' EXIT

if [ $UNAME = "arm64" ]
then
    echo "Go-Amazing dockerbuild engine haven't support your arm64 base M1 chip yet. Please direclty use makefile commands."
    exit 1
fi

echo "Try initializing network 'rails' or skit it ..."
echo "=> \c"
docker network create --driver bridge rails || true

case "$1" in
  codegen)
    docker-compose run --rm codegen
    ;;
  test)
    docker-compose run --rm test
    ;;
  rpc)
    docker-compose down
    docker-compose run --rm build
    docker-compose build --force-rm bevis-chang-rpc
    docker-compose up -d bevis-chang-rpc
    ;;
  build)
    docker-compose run --rm build
    docker-compose build --force-rm bevis-chang-rpc
    docker-compose build --force-rm bevis-chang-cron
    ;;
  ci-build)
    docker-compose run --rm ci-build
    docker-compose build --force-rm bevis-chang-rpc
    docker-compose build --force-rm bevis-chang-cron
    ;;
esac