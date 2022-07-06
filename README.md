## TCP v/s UDP server in golang

start the server

```bash
# start a tcp server
go run main.go tcp.go udp.go -p tcp

# start a udp server
go run main.go tcp.go udp.go -p udp
```

## Using Telnet to connect to TCP server

```bash
telnet 127.0.0.1 8080
>> Hi
>> Server
>> TCP
```

## Using Netcat to connect to UDP server

```bash
echo "Hi" | nc -w 1 -u 127.0.0.1 8081
echo "Server" | nc -w 1 -u 127.0.0.1 8081
echo "UDP" | nc -w 1 -u 127.0.0.1 8081
```