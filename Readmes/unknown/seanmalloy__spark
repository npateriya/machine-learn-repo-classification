# Cisco Spark Command Line Client

## How To Use
### Messages
```
$ spark msg -p <person> "hello world" # send message to specific person
$ spark msg -s <space> "hello world"  # send message to specific space
$ spark msg -f <filename> -p <person> # send file attachment to specific person
$ spark msg -f <filename> -s <space>  # send file attachment to specific space
```

### Help
```
$ spark help
$ spark help msg

```
### Config File
The config file is ~/.spark/config.json. It only supports specifying the auth token.
```
{
  "auth": "thisisyourauthtokenforciscospark"
}
```

### Examples
```
# send a message to sparky
$ spark msg -p 722bb271-d7ca-4bce-a9e3-471e4412fa77 "hello world"

# send a message to a space
$ spark msg -s Y2lzY29zcGFyazovL3VzL1JPT00vZGQ3OTljYjAtYTY2Mi0xMWU3LThhNjAtODcyMzRiYzQ2MmQ1 "hello world"
```

## How To Build
```
$ go build
```
