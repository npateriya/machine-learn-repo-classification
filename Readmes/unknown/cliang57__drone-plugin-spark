# Spark Notify Plugin for Drone 0.5 Server

### Ref:
1. http://readme.drone.io/usage/variables
2. http://readme.drone.io/0.5/plugin-parameters
3. http://readme.drone.io/0.5/usage/environment-reference
4. https://developer.ciscospark.com
5. https://github.com/hpreston/drone-spark
6. https://github.com/drone-plugins/drone-slack

### in '.../src/github.com/cliang57/drone-plugin-spark' folder

### manage third party packages (ref: https://glide.sh/)
```
# init
rm -rf vendor
mkdir -p vendor
rm -rf glide.lock glide.yaml
export GOPATH=$(cd ../../../..; pwd)
glide init

export GOPATH=$(cd ../../../..; pwd)
glide update
```

### build the go application
```
export GOPATH=$(cd ../../../..; pwd)
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo
```

### example usage (standalone)
```
DRONE_REPO_OWNER=octocat \
    PLUGIN_MESSAGE="Thank you" \
    PLUGIN_AUTH_TOKEN="MmY2****" \
    PLUGIN_ROOMNAME="****" \
    DRONE_SERVER="http://drone.mycompany.com" \
    DRONE_REPO_OWNER=octocat \
    DRONE_REPO_NAME=hello-world \
    DRONE_REPO="octocat/hello-world" \
    DRONE_COMMIT_SHA=7fd1a60b01f91b314f59955a4e4d4e80d8edf11d \
    DRONE_COMMIT_BRANCH=master \
    DRONE_COMMIT_AUTHOR="Octocat A" \
    DRONE_COMMIT_AUTHOR_EMAIL=octocat@github.com \
    DRONE_COMMIT_LINK=https://github.com/drone-plugins/drone-slack/commit/f20bf5b1bec4cc789ac9cbe3ac9a0f68b8640a23 \
    DRONE_COMMIT_MESSAGE="Bug Fix" \
    DRONE_BUILD_NUMBER=1 \
    DRONE_BUILD_STATUS=success \
    DRONE_BUILD_LINK="http://link_to_drone_server_build_page" \
    ./drone-plugin-spark
```

### build and push docker image
```
export GOPATH=$(cd ../../../..; pwd)
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo
docker build -t cliang2/drone-spark .

### docker login
docker push cliang2/drone-spark
```

### example usage (in .drone.yml)
```
  notify:
    image: cliang2/drone-spark
    auth_token: MmY2***
    roomName: "***"
    message: "Great job on your new server build!! (From Drone 0.5)"
    when:
      status: [ success, failure ]
```
