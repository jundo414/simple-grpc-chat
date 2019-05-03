# Simple text-based chat using gRPC

## Usage
### for starting chat servers
```
$ go run chat-server/server.go
```

### for starting chatting
\* RESTRICTION: Don't use duplicated name.
```
$ go run chat-client/client.go
```

## Example
Client#1(Tom):
```
$ go run chat-client/client.go
username: Tom
Jessie> Hi, Tom
Hello, Jessie
How r u doing?
```

Client#2(Jessie):
```
$ go run chat-client/client.go
username: Jessie
Hi, Tom
Tom> Hello, Jessie
Tom> How r u doing?
```

## References
- [Magazine] WEB+DB PRESS Vol.110 
