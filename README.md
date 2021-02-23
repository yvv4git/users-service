# users-service
A simple example of a microservice architecture.

## BUILD
```
make build
```

## SERVER
```
./server.bin
```

## CLIENT
```
client.bin --host 127.0.0.1 --port 1234 create Vladimir yvv4test@gmail.com 32
client.bin --host 127.0.0.1 --port 1234 read 2
client.bin --host 127.0.0.1 --port 1234 update 2 Anonymous anonym@mail.ru 100
client.bin --host 127.0.0.1 --port 1234 delete 3
```