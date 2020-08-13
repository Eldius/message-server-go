#!/bin/sh

CURR_DIR=${PWD}

DOCKER_OWNER=eldius
DOCKER_REPO=auth-server
VERSION=0.0.1
CONTAINER_NAME=auth-server
APP_BIN_FILE=auth-server

build_app() {
    if [ "${TEST}" -eq "1" ];then
      echo "######################"
      echo "# testing app code   #"
      echo "######################"
      go test ./... \
        || exit 1
    fi

    echo "######################"
    echo "# building app       #"
    echo "######################"
    yarn install && \
    yarn run build-dev && \
    CGO_ENABLED=0 GOOS=linux go build -v -o bin/$APP_BIN_FILE -a -ldflags '-extldflags "-static"' . \
    || exit 1
}

build_image() {
  echo "##################"
  echo "# building image #"
  echo "##################"
  docker \
      build \
      -t $DOCKER_OWNER/$DOCKER_REPO:$VERSION \
      --no-cache . && \
      docker tag $DOCKER_OWNER/$DOCKER_REPO:$VERSION $DOCKER_OWNER/$DOCKER_REPO:latest \
      || exit 1
}

stop_container() {
  echo "######################"
  echo "# starting container #"
  echo "######################"

  docker stop \
    $CONTAINER_NAME || \
  exit 1

}

start_container() {

  if [ "${START}" -eq "1" ];then
    echo "######################"
    echo "# starting container #"
    echo "######################"

    docker run \
      -p "8000:8000" \
      -p "8000:8000/udp" \
      -it \
      --name $CONTAINER_NAME \
      --rm \
      -d \
      $DOCKER_OWNER/$DOCKER_REPO:$VERSION || \
    exit 1
  fi

}

# TODO create a test script
test_image() {
  if [ "${TEST}" -eq "1" ];then
    echo "######################"
    echo "# testing  container #"
    echo "######################"

    #cd ${CURR_DIR}/container_test
    #py.test . -s || \
    #  exit 1

    echo ""
    echo "Not implemented yet!"
    echo ""
    echo ""
  fi
}

push_image() {
  if [ "${PUSH}" -eq "1" ];then
  echo "##################"
  echo "# pushing image  #"
  echo "##################"
    docker push $DOCKER_OWNER/$DOCKER_REPO:$VERSION && \
    docker push $DOCKER_OWNER/$DOCKER_REPO:latest
  fi
}

TEST="0"
PUSH="0"
START="0"

for var in "$@"
do
  case $var in
    "--test"            )
        echo " -> Test image ON"
        TEST="1"
        ;;
    "--push"           )
        echo " -> Push image ON"
        PUSH="1"
        ;;
    "--start"           )
        echo " -> Start container ON"
        START="1"
        ;;
  esac
done


docker stop $CONTAINER_NAME
build_app
build_image
test_image
push_image
start_container
