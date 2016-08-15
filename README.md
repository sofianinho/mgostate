### The program output/help
```
./mgostate 
mgostate. 
   Helps you test mongodb parameters (like the existance of a database or a collection) from the shell.

Usage:
  mgostate [command]

Available Commands:
  collection  Show the state of your collection in mongodb
  database    Show the state of your database in mongodb
  server      Show the state of your mongodb server/cluster
  state       Show the state in your mongodb

Flags:
  -c, --collection string   Name of the collection to check in the database (default "myTestCollection")
  -d, --database string     Name of the database to check in mongoDB (default "myTestDB")
  -h, --help                help for mgodb
  -u, --url string          Uri to connect to mongoDB (default "mongodb://localhost:27017/")

Use "mgostate [command] --help" for more information about a command.
```
### Normal compilation
The result is a dynamically linked executable. This usually means that a cgo/c code is nested in there.
This could be: "gopkg.in/mgo.v2" or "github.com/spf13/cobra" (probably mgo.v2). It may be, for example, due to the use of DNS resolution (or other).
```
go build ./mgodb.go
```
### Static compilation
To get a statically linked program rather than a dynamically linked one, you need to tell the go compiler not to link to the system's C library.
```
CGO_ENABLED=0 go build ./mgodb.go
```
### Docker
This folder contains the files necessary to create a dockerized image of the software. The image used is a "scratch" as the executable is statically
linked and does not a bash environement to make sense of its libraries. Read further [here](http://blog.oddbit.com/2015/02/05/creating-minimal-docker-images/).
You can also use this version of the software from my dockerhub repo. You can then append you commands to the image name.
```
docker run -t sofianinho/mgostate --help
```


