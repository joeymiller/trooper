# **trooper**

## **Description**
CLI returns a set of temporary security credentials (consisting of an access key ID, a secret access key, 
and a security token) that you can use to access AWS resources. trooper removes the dependency of issuing AWS Keys to every engineer on the team. Currently credentials expire 1 hour from time generated.


## **Installation**

Make sure you have a working Go environment. See the [install instructions](https://golang.org/doc/install)

To install trooper, simply run:
```
go get github.com/joeymiller/trooper
```
To compile it from source:
```
cd $GOPATH/src/github.com/joeymiller/trooper
go get -u
go build
```
*Run trooper --help to see a full list of options.*

## **Run**

Generate AWS credentials to console.
```
$ trooper console generate-credentials
```

console generate-credentials results

```
Access Key ID: xxxx
Secret Access Key: xxxx
Session Token: xxxx
Expires: 2017-02-14 19:42:53 +0000 UTC
````

Run trooper in server mode.
```
$ trooper server run
```

Server mode options
```
    --port [PORT | 8080]
    --host [HOST | 127.0.0.1]
```

Server routes
```
REQUEST 
[GET] /

RESPONSE
{
    "status":"ok",
    "code":200,

}
```
```
REQUEST 
[GET] /credentials

RESPONSE
{
    "accessKeyId":"xxxx",
    "secretAccessKey":"xxxx",
    "sessionToken":"xxxx"
}
```


## **TODO**
 - Will work on uploading cloudformation IAM templates for examples. 
 - More command line options.
 - better way to manage app credentials.
 - More stuff here...