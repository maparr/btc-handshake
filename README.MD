# Go Bitcoin Handshake

This project implements a simple Bitcoin handshake protocol in Go. It includes the ability to send and receive version messages and handle verack messages.

## Project Structure

```
btc-handshake/
├── checksum/
│   ├── checksum.go
│   ├── checksum_test.go
├── message/
│   ├── header.go
│   ├── message.go
│   ├── param.go
│   ├── verack.go
│   ├── verack_test.go
│   ├── versionmsg.go
│   ├── versionmsg_test.go
├── go.mod
├── handshake.go
├── README.md
```

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/maparr/btc-handshake.git
   ```
   
2. Navigate to the project directory:
   ```sh
   cd btc-handshake
   ```

## Running the Code

To run the handshake program, use the following command:
```sh
go run handshake.go -node 127.0.0.1 -port 8333
```
Replace `127.0.0.1` with the IP address of the node you want to connect to and `8333` with the port if different.

For example:
```sh
go run handshake.go -node 127.0.0.1 -port 8333
```

## Running btcd

To clone the `btcd` repository and run it with debug level, follow these steps:

1. **Clone the Repository**:
   ```sh
   git clone https://github.com/btcsuite/btcd.git
   ```

2. **Navigate to the `btcd` Directory**:
   ```sh
   cd btcd
   ```

3. **Download and Install Dependencies**:
   ```sh
   go mod tidy
   ```

4. **Run `btcd` with Debug Level**:
   ```sh
   go run . --debuglevel=debug
   ```

## Running the Tests

To run the tests for all packages, use the following command:
```sh
go test ./...
```

To run tests for a specific package, navigate to the package directory and use:
```sh
cd checksum
go test

cd ../message
go test
```

## Result

After running the `handshake.go` program and setting up `btcd` in debug mode, you should see the following line in the `btcd` logs:

```plaintext
2024-05-27 10:50:31.690 [DBG] PEER: Received version (agent /btc-handshake:0.1.0/, pver 70015, block 0) from 127.0.0.1:59994 (inbound)
```

This indicates that the handshake protocol was successfully executed and the `btcd` node received the version message from your Go application.
```