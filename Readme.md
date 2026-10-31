## NET-CAT 🐱

### Authors

- [Houssam Eddine Hamouich (@drawliin)](https://github.com/drawliin)

A concurrent TCP chat server application built with Go, enabling real-time multi-client messaging with message broadcasting.

### Overview

NET-CAT is a pure Go implementation of a TCP-based chat server that demonstrates networking fundamentals and concurrent programming patterns. It allows multiple clients to connect simultaneously and communicate in real-time through a message broadcasting system.

The project showcases essential Go concepts including:
- TCP socket programming
- Goroutine-based concurrency
- Mutex synchronization
- Channel-based communication

### Features
 
- **🔌 TCP Socket Server**: Pure Go TCP listener with configurable port (default: `8989`)
- **👥 Multi-client Support**: Handle up to **10** concurrent client connections
- **📢 Message Broadcasting**: Real-time message delivery to all connected clients
- **🔐 Client Authentication**: Username validation with strict naming rules
- **📜 Message History**: Automatic history playback for newly joined clients (in-memory)
- **⏰ Timestamp Logging**: All messages include date and time stamps
- **✅ Input Validation**: Comprehensive validation for usernames and message content
- **🔄 Graceful Cleanup**: Proper connection termination and resource management
- **🎨 Custom Greeting**: ASCII art welcome message (customizable via `greeting.txt`)

---

### Technical Details

#### Architecture

NET-CAT uses a goroutine-per-client architecture with channel-based communication:

- **TCP Listener** (`main.go`): Accepts incoming connections and spawns goroutines
- **Client Handler** (`helpers/handleClients.go`): Manages individual client lifecycles
- **Message Broadcaster** (`helpers/broadcastMessages.go`): Distributes messages to all clients
- **Synchronization**: Mutex locks protect shared client data structures

#### Concurrency Model

- Each client connection runs in its own goroutine
- Message broadcasting runs in a separate goroutine
- Channels facilitate non-blocking communication between goroutines
- Mutex locks ensure thread-safe access to the clients list

---

### Validation Rules

#### Username Requirements

- Alphanumeric characters only (plus dash `-` and underscore `_`)
- Maximum **15** characters
- Cannot be empty or duplicate
- Case-insensitive duplicate checking

#### Message Requirements

- Printable ASCII characters only (values 32–126)
- No control characters or special encodings
- Non-empty messages only

---

### Prerequisites

- **Go** `1.22.3` or later
- A terminal or command-line interface

---

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/drawliin/NET-CAT.git
   cd NET-CAT
   ```

2. **Build the application:**

   ```bash
   go build -o net-cat
   ```

---

### Usage

#### Starting the Server

- **Run with default port (`8989`):**

  ```bash
  ./net-cat
  ```

- **Run with a custom port:**

  ```bash
  ./net-cat 3000
  ```

#### Connecting a Client

In another terminal, use `nc` (netcat) or `telnet`:

```bash
nc localhost 8989
# or
telnet localhost 8989
```

---

### What You’ll See

#### Server Output

```text
Listening on the port :8989
```

#### Client Interaction

- The server displays the welcome greeting from `greeting.txt`
- You are prompted to enter a valid username
- You receive message history from previous conversations
- You can start chatting with other connected clients
- Messages are timestamped and include your username
- Type messages and press Enter to send
- Disconnect by closing the connection

##### Example Session

```text
Welcome to TCP-Chat!
          _nnnn_
         dGGGGMMb
        ...

[ENTER YOUR NAME]: alice
[2026-01-14 13:42:24][bob]: Hello alice!
[2026-01-14 13:42:24][alice]: Hi bob!
```

---

### Project Structure

```text
NET-CAT/
├── main.go                    # TCP server and connection listener
├── go.mod                     # Go module definition
├── greeting.txt               # ASCII art welcome message
├── helpers/
│   ├── structures.go          # Server, Client, Message types
│   ├── handleClients.go       # Client connection lifecycle
│   ├── broadcastMessages.go   # Message broadcasting system
│   ├── clientWriter.go        # Client output goroutine
│   ├── removeClient.go        # Client disconnection handling
│   ├── isDuplicateName.go     # Username duplicate checking
│   ├── isValidName.go         # Username validation
│   ├── isPrintable.go         # Message content validation
│   └── readGreeting.go        # Greeting file loading
└── Readme.md                  # Project documentation
```

---

### Code Overview

#### Main Server Loop

```go
server := &helpers.Server{
    Broadcast: make(chan helpers.Message),
    Greeting:  helpers.ReadGreeting(),
}

listener, _ := net.Listen("tcp", ":"+port)
go server.BroadcastMessages()

for {
    conn, _ := listener.Accept()
    go server.HandleClients(conn)
}
```

#### Client Handler

The client handler is responsible for:

- Prompting for and validating username
- Loading message history
- Broadcasting join messages
- Handling incoming messages in a loop
- Broadcasting leave messages on disconnect

#### Message Broadcasting

- Non-blocking channel sends prevent server slowdown
- Lagging clients are automatically disconnected
- Client-specific send channels enable flow control
- All clients receive messages except the sender

---

### Customization

#### Custom Greeting

Edit `greeting.txt` to customize the welcome message displayed to new clients:

```bash
echo "Welcome to My Chat Server!" > greeting.txt
```

If `greeting.txt` is missing or unreadable, the default greeting is used.

#### Connection Limit

The maximum client limit is set to **10**. To modify it, edit `helpers/handleClients.go`:

```go
if len(s.clients) >= 10 {  // Change this value
    conn.Write([]byte("Chat is full\n"))
    return
}
```

---

### Limitations

- Maximum 10 concurrent connections
- Messages limited to printable ASCII characters
- Usernames limited to 15 characters
- No authentication beyond username validation
- No encryption (plain TCP)
- Message history not persisted to disk

---

### Error Handling

The server gracefully handles:

- Invalid usernames (format, length, duplicates)
- Connection errors
- Client disconnections
- Invalid message content
- Full server capacity
- Non-printable characters

---

### Learning Outcomes

This project demonstrates:

- TCP socket programming with Go’s `net` package
- Goroutine-based concurrent programming
- Channel-based communication patterns
- Mutex locks for thread-safe operations
- Connection lifecycle management
- Input validation and error handling
- Clean code organization with helper functions
- Message broadcasting patterns

---

### Debugging & Testing

#### Debugging

To debug or understand flow better, you can:

- Enable verbose logging
- Add debug prints to helper functions to trace:
  - Client connections/disconnections
  - Message broadcasting
  - Channel operations

#### Test with Multiple Clients

Open several terminal windows and connect:

```bash
# Terminal 1
./net-cat

# Terminal 2
nc localhost 8989

# Terminal 3
nc localhost 8989

# Terminal 4
nc localhost 8989
```

---

### Performance Considerations

- Goroutines are lightweight; 10 concurrent clients use minimal memory
- Channel operations are efficient for message passing
- Non-blocking sends prevent one slow client from blocking others
- Mutex locks are only held briefly for client list operations

---

### Future Enhancements

- Persistent message history (database)
- User authentication (password-based)
- Private messaging between users
- Message encryption (TLS)
- Admin commands (kick, mute)
- Rate limiting per client
- Web-based client interface
- Configuration file support
- Metrics and monitoring
- Unit tests

---

### License

No license file is included in this repository.

---

### Contributing

Contributions are welcome! Feel free to:

- Report bugs
- Suggest new features
- Submit pull requests
- Improve documentation

Enjoy chatting with **NET-CAT**! 🐱

