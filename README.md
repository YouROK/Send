# Send
A small tool for sending and receiving network

## Install
```
git clone "https://github.com/YouROK/Send.git"
cd Send
go install ./src/send
```

## Usage
```
Run program for receiving and sending data
Default port :30123
```

### Receiving data
```
send -h ":30123" -n tcp -r
also
send -n tcp -r
also
send -r
```

### Sending data
```
send -h "127.0.0.1:30123" -n tcp 
also
send -h "127.0.0.1"
```
