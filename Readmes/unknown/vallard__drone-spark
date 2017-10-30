# Drone Spark
Drone plugin for sending Cisco Spark notifications.  For the useage information and a listing of the available options please take a look at [the docs.](DOCS.md) 

## Build
Build the binary with the following commands:

```bash
go build
go test
```

## Docker
Build the docker image with the following commands: 

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo
docker build -t vallard/drone-spark .
```
Please note incorrectly building the image for the correct x64 linux and with GCO disabled will result in an error when running the Docker image:

```
docker: Error response from daemon: Container command
'/bin/drone-spark' not found or does not exist..
```

## Usage
Execute from the working directory:
 
```
drone run --rm \ 
  -e SPARK_TOKEN=Xjdre.... \
  -e PLUGIN_ROOM=MySparkRoom \
  -e PLUGIN_TEMPLATE_MD="#Testing \n##Second heading \
  -e DRONE_REPO_OWNER=octocat \
  -e DRONE_REPO_NAME=hello-world \
  -e DRONE_COMMIT_SHA=7fd1a60b01f91b314f59955a4e4d4e80d8edf11d \
  -e DRONE_COMMIT_BRANCH=master \
  -e DRONE_COMMIT_AUTHOR=octocat \
  -e DRONE_BUILD_NUMBER=1 \
  -e DRONE_BUILD_STATUS=success \
  -e DRONE_BUILD_LINK=http://github.com/octocat/hello-world \
  -e DRONE_TAG=1.0.0 \
  vallard/drone-spark

```


## Inspiration 

* [drone-slack](https://github.com/drone-plugins/drone-slack/)
* [drone-spark](https://github.com/hpreston/drone-spark)
