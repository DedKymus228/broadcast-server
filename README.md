# Broadcast Server

A simple broadcast chat server and client application built in Go using WebSockets. 

## Overview

This project consists of two main parts:
1.  **WebSocket Server**: Handles incoming connections, manages clients, and broadcasts messages to all connected clients (except the sender).
2.  **CLI Client**: A command-line application that allows users to connect to the server, send messages, and receive broadcasts in real-time.

## Prerequisites

- Go (version 1.18 or higher recommended)

## Installation

Clone the repository:

```bash
git clone https://github.com/DedKymus228/broadcast-server.git
cd broadcast-server
```

Download dependencies:

```bash
go mod tidy
```

## Usage

You can run the server and the client in separate terminal windows.

### Starting the Server

The server runs on port `8080` by default.

```bash
go run main.go start
```

You can optionally specify a different port (assuming you have a flag for it, e.g., `--port`):
```bash
go run main.go start --port 8080
```

### Connecting with the Client

Open a new terminal window and connect to the server. By default, it connects to `localhost:8080`.

```bash
go run main.go connect
```

You can also specify the address and port (if your command supports it):
```bash
go run main.go connect -a localhost -p 8080
```

### Chatting

Once connected, simply type your message and press `Enter` to broadcast it to all other connected clients. Press `Ctrl+C` to gracefully exit the application.

## Features

- Real-time communication using WebSockets (`github.com/gorilla/websocket`).
- Graceful shutdown of connections.
- Prevents echo of messages (the sender doesn't receive their own message back).
- Command-line interface built with Cobra (`github.com/spf13/cobra`).

## Structure

- `cmd/`: Command-line setup using Cobra (e.g., `start` and `connect` commands).
- `internal/client/`: Logic for the WebSocket client and the Hub (which manages connections).
- `internal/server/`: The HTTP/WebSocket server logic.
- `main.go`: Application entry point.
