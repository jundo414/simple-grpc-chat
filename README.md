# Simple text-based chat using gRPC

## Usage
for starting chat servers
```
$ go run chat-server/server.go
```

for starting chatting
```
$ go run chat-client/client.go {username}
```

## Example
Client#1(Tom):
```
$ go run chat-client/client.go Tom
how r u?
[Tom] how r u?
[Jessie] not good
[Jessie] my dog ran away from home again..
really?! that's too bad..
[Tom] really?! that's too bad..
```

Client#2(Jessie):
```
$ go run chat-client/client.go Jessie
[Tom] how r u?
not good
[Jessie] not good
my dog ran away from home again..          
[Jessie] my dog ran away from home again..
[Tom] really?! that's too bad..
```

## References
- [Magazine] WEB+DB PRESS Vol.110 
