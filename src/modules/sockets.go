package modules

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/javanhut/TheCarrionLanguage/src/object"
)

// Global socket handle registry
var (
	socketHandles = make(map[int64]interface{})
	nextSocketHandle int64 = 1
	socketHandleMutex sync.RWMutex
)

// Global port allocation mutex and tracking
var (
	portAllocationMutex sync.Mutex
	allocatedPorts = make(map[string]bool) // Track allocated ports by "host:port"
)

// Socket types
const (
	SocketTypeTCP = "tcp"
	SocketTypeUDP = "udp"
	SocketTypeWeb = "web"
	SocketTypeUnix = "unix"
)

// Helper functions
func extractSocketString(obj object.Object) (string, bool) {
	switch v := obj.(type) {
	case *object.String:
		return v.Value, true
	case *object.Instance:
		if value, exists := v.Env.Get("value"); exists {
			if strVal, ok := value.(*object.String); ok {
				return strVal.Value, true
			}
		}
		return "", false
	default:
		return "", false
	}
}

func extractSocketInt(obj object.Object) (int64, bool) {
	switch v := obj.(type) {
	case *object.Integer:
		return v.Value, true
	default:
		return 0, false
	}
}

// Socket handle management
func getSocketHandle(handleID int64) (interface{}, bool) {
	socketHandleMutex.RLock()
	defer socketHandleMutex.RUnlock()
	socket, exists := socketHandles[handleID]
	return socket, exists
}

func storeSocketHandle(socket interface{}) int64 {
	socketHandleMutex.Lock()
	defer socketHandleMutex.Unlock()
	handleID := nextSocketHandle
	nextSocketHandle++
	socketHandles[handleID] = socket
	return handleID
}

func removeSocketHandle(handleID int64) {
	socketHandleMutex.Lock()
	defer socketHandleMutex.Unlock()
	delete(socketHandles, handleID)
}

// Port validation function
func isValidPort(port int) bool {
	return port >= 1 && port <= 65535
}

// Port allocation functions
func allocatePort(address string) (string, string) {
	portAllocationMutex.Lock()
	defer portAllocationMutex.Unlock()

	// Parse the address to extract host and port
	host, port, err := net.SplitHostPort(address)
	if err != nil {
		// If parsing fails, return original address
		return address, ""
	}

	// Convert port to integer for incrementing
	portNum, err := strconv.Atoi(port)
	if err != nil {
		// If port is not a number, return original address
		return address, ""
	}

	// Validate port range
	if !isValidPort(portNum) {
		return address, fmt.Sprintf("Port %d is out of valid range (1-65535)", portNum)
	}

	originalPort := portNum
	var message string

	// Try up to 100 ports to find an available one
	for attempts := 0; attempts < 100; attempts++ {
		// Validate current port number is still in range
		if !isValidPort(portNum) {
			message = fmt.Sprintf("Could not find available port starting from %d (reached end of valid port range)", originalPort)
			return address, message
		}

		testAddress := net.JoinHostPort(host, strconv.Itoa(portNum))
		
		// Resolve the address to canonical form for consistent tracking
		canonicalAddress, err := resolveToCanonical(testAddress)
		if err != nil {
			canonicalAddress = testAddress // fallback to original if resolution fails
		}
		
		// Check if port is already tracked as allocated
		if !allocatedPorts[canonicalAddress] {
			// Attempt to bind directly to eliminate TOCTOU race condition
			listener, err := net.Listen("tcp", testAddress)
			if err == nil {
				// Successfully bound - mark as allocated and close listener
				listener.Close()
				allocatedPorts[canonicalAddress] = true
				
				if portNum != originalPort {
					message = fmt.Sprintf("Port %d already allocated, incremented to port %d", originalPort, portNum)
				}
				
				return testAddress, message
			}
			
			// Binding failed, implement exponential backoff for retries
			backoffDuration := time.Duration(1<<uint(attempts%5)) * time.Millisecond
			time.Sleep(backoffDuration)
		}
		
		// Port is in use, try next port
		portNum++
	}

	// If we couldn't find an available port, return original with error message
	message = fmt.Sprintf("Could not find available port starting from %d", originalPort)
	return address, message
}

// Helper function to resolve address to canonical form (e.g., localhost -> 127.0.0.1)
func resolveToCanonical(address string) (string, error) {
	// Use net.ResolveTCPAddr to get the canonical form
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return address, err
	}
	return tcpAddr.String(), nil
}

func releasePort(address string) {
	portAllocationMutex.Lock()
	defer portAllocationMutex.Unlock()
	delete(allocatedPorts, address)
}

// Socket wrapper types
type TCPSocket struct {
	Conn    net.Conn
	Type    string
	Address string
	Timeout time.Duration
}

type TCPListener struct {
	Listener net.Listener
	Type     string
	Address  string
	Timeout  time.Duration
}

type UDPSocket struct {
	Conn    *net.UDPConn
	Type    string
	Address string
	Timeout time.Duration
}

type WebSocket struct {
	Server  *http.Server
	Mux     *http.ServeMux
	Type    string
	Address string
	Timeout time.Duration
}

type UnixSocket struct {
	Conn    net.Conn
	Type    string
	Address string
	Timeout time.Duration
}

var SocketsModule = map[string]*object.Builtin{
	"new_socket": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 {
				return &object.Error{Message: "new_socket requires at least 1 argument: type, [protocol], [port/address], [timeout]"}
			}

			socketType, ok := extractSocketString(args[0])
			if !ok {
				return &object.Error{Message: "new_socket: type must be a string"}
			}

			protocol := "tcp"
			if len(args) > 1 {
				if p, ok := extractSocketString(args[1]); ok {
					protocol = p
				}
			}

			address := "localhost:8080"
			if len(args) > 2 {
				if addr, ok := extractSocketString(args[2]); ok {
					address = addr
				}
			}

			timeout := 30 * time.Second
			if len(args) > 3 {
				if t, ok := extractSocketInt(args[3]); ok {
					if t < 0 {
						// Use default timeout for negative values
						timeout = 30 * time.Second
					} else {
						timeout = time.Duration(t) * time.Second
					}
				}
			}

			switch strings.ToLower(socketType) {
			case "tcp":
				return createTCPSocket(protocol, address, timeout)
			case "udp":
				return createUDPSocket(address, timeout)
			case "web", "http":
				return createWebSocket(address, timeout)
			case "unix":
				return createUnixSocket(address, timeout)
			default:
				return &object.Error{Message: fmt.Sprintf("unsupported socket type: %s", socketType)}
			}
		},
	},

	"client": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 {
				return &object.Error{Message: "client requires at least 2 arguments: type, address, [timeout]"}
			}

			socketType, ok := extractSocketString(args[0])
			if !ok {
				return &object.Error{Message: "client: type must be a string"}
			}

			address, ok := extractSocketString(args[1])
			if !ok {
				return &object.Error{Message: "client: address must be a string"}
			}

			timeout := 30 * time.Second
			if len(args) > 2 {
				if t, ok := extractSocketInt(args[2]); ok {
					timeout = time.Duration(t) * time.Second
				}
			}

			switch strings.ToLower(socketType) {
			case "tcp":
				return connectTCPClient(address, timeout)
			case "udp":
				return connectUDPClient(address, timeout)
			case "unix":
				return connectUnixClient(address, timeout)
			default:
				return &object.Error{Message: fmt.Sprintf("unsupported client type: %s", socketType)}
			}
		},
	},

	"server": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 {
				return &object.Error{Message: "server requires at least 2 arguments: type, port/address, [timeout]"}
			}

			socketType, ok := extractSocketString(args[0])
			if !ok {
				return &object.Error{Message: "server: type must be a string"}
			}

			address, ok := extractSocketString(args[1])
			if !ok {
				return &object.Error{Message: "server: address must be a string"}
			}

			timeout := 30 * time.Second
			if len(args) > 2 {
				if t, ok := extractSocketInt(args[2]); ok {
					timeout = time.Duration(t) * time.Second
				}
			}

			switch strings.ToLower(socketType) {
			case "tcp":
				return startTCPServer(address, timeout)
			case "udp":
				return startUDPServer(address, timeout)
			case "web", "http":
				return startWebServer(address, timeout)
			case "unix":
				return startUnixServer(address, timeout)
			default:
				return &object.Error{Message: fmt.Sprintf("unsupported server type: %s", socketType)}
			}
		},
	},

	"socket_send": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return &object.Error{Message: "socket_send requires 2 arguments: handleID, data"}
			}

			handleID, ok := extractSocketInt(args[0])
			if !ok {
				return &object.Error{Message: "socket_send: handleID must be an integer"}
			}

			data, ok := extractSocketString(args[1])
			if !ok {
				return &object.Error{Message: "socket_send: data must be a string"}
			}

			socket, exists := getSocketHandle(handleID)
			if !exists {
				return &object.Error{Message: "socket_send: invalid socket handle"}
			}

			return sendData(socket, data)
		},
	},

	"socket_receive": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 || len(args) > 2 {
				return &object.Error{Message: "socket_receive requires 1-2 arguments: handleID, [bufferSize]"}
			}

			handleID, ok := extractSocketInt(args[0])
			if !ok {
				return &object.Error{Message: "socket_receive: handleID must be an integer"}
			}

			bufferSize := int64(1024)
			if len(args) > 1 {
				if size, ok := extractSocketInt(args[1]); ok {
					bufferSize = size
				}
			}

			socket, exists := getSocketHandle(handleID)
			if !exists {
				return &object.Error{Message: "socket_receive: invalid socket handle"}
			}

			return receiveData(socket, bufferSize)
		},
	},

	"socket_close": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "socket_close requires 1 argument: handleID"}
			}

			handleID, ok := extractSocketInt(args[0])
			if !ok {
				return &object.Error{Message: "socket_close: handleID must be an integer"}
			}

			socket, exists := getSocketHandle(handleID)
			if !exists {
				return &object.Error{Message: "socket_close: invalid socket handle"}
			}

			err := closeSocket(socket)
			removeSocketHandle(handleID)

			if err != nil {
				return &object.Error{Message: fmt.Sprintf("failed to close socket: %v", err)}
			}

			return &object.None{}
		},
	},

	"socket_listen": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "socket_listen requires 1 argument: handleID"}
			}

			handleID, ok := extractSocketInt(args[0])
			if !ok {
				return &object.Error{Message: "socket_listen: handleID must be an integer"}
			}

			socket, exists := getSocketHandle(handleID)
			if !exists {
				return &object.Error{Message: "socket_listen: invalid socket handle"}
			}

			return listenForConnections(socket)
		},
	},

	"socket_accept": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "socket_accept requires 1 argument: handleID"}
			}

			handleID, ok := extractSocketInt(args[0])
			if !ok {
				return &object.Error{Message: "socket_accept: handleID must be an integer"}
			}

			socket, exists := getSocketHandle(handleID)
			if !exists {
				return &object.Error{Message: "socket_accept: invalid socket handle"}
			}

			return acceptConnection(socket)
		},
	},

	"socket_set_timeout": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return &object.Error{Message: "socket_set_timeout requires 2 arguments: handleID, timeoutSeconds"}
			}

			handleID, ok := extractSocketInt(args[0])
			if !ok {
				return &object.Error{Message: "socket_set_timeout: handleID must be an integer"}
			}

			timeoutSecs, ok := extractSocketInt(args[1])
			if !ok {
				return &object.Error{Message: "socket_set_timeout: timeoutSeconds must be an integer"}
			}

			socket, exists := getSocketHandle(handleID)
			if !exists {
				return &object.Error{Message: "socket_set_timeout: invalid socket handle"}
			}

			return setSocketTimeout(socket, time.Duration(timeoutSecs)*time.Second)
		},
	},

	"socket_get_info": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "socket_get_info requires 1 argument: handleID"}
			}

			handleID, ok := extractSocketInt(args[0])
			if !ok {
				return &object.Error{Message: "socket_get_info: handleID must be an integer"}
			}

			socket, exists := getSocketHandle(handleID)
			if !exists {
				return &object.Error{Message: "socket_get_info: invalid socket handle"}
			}

			return getSocketInfo(socket)
		},
	},

	"socket_send_to": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 3 {
				return &object.Error{Message: "socket_send_to requires 3 arguments: handleID, data, targetAddress"}
			}
			handleID, ok := extractSocketInt(args[0])
			if !ok {
				return &object.Error{Message: "socket_send_to: handleID must be an integer"}
			}
			data, ok := extractSocketString(args[1])
			if !ok {
				return &object.Error{Message: "socket_send_to: data must be a string"}
			}
			targetAddress, ok := extractSocketString(args[2])
			if !ok {
				return &object.Error{Message: "socket_send_to: targetAddress must be a string"}
			}
			socket, exists := getSocketHandle(handleID)
			if !exists {
				return &object.Error{Message: "socket_send_to: invalid socket handle"}
			}
			return sendDataTo(socket, data, targetAddress)
		},
	},

	"socket_receive_from": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return &object.Error{Message: "socket_receive_from requires 2 arguments: handleID, bufferSize"}
			}
			handleID, ok := extractSocketInt(args[0])
			if !ok {
				return &object.Error{Message: "socket_receive_from: handleID must be an integer"}
			}
			bufferSize, ok := extractSocketInt(args[1])
			if !ok {
				return &object.Error{Message: "socket_receive_from: bufferSize must be an integer"}
			}
			socket, exists := getSocketHandle(handleID)
			if !exists {
				return &object.Error{Message: "socket_receive_from: invalid socket handle"}
			}
			return receiveDataFrom(socket, bufferSize)
		},
	},
}

// Implementation functions
func createTCPSocket(protocol, address string, timeout time.Duration) object.Object {
	socket := &TCPSocket{
		Type:    SocketTypeTCP,
		Address: address,
		Timeout: timeout,
	}

	handleID := storeSocketHandle(socket)
	return &object.Integer{Value: handleID}
}

func createUDPSocket(address string, timeout time.Duration) object.Object {
	socket := &UDPSocket{
		Type:    SocketTypeUDP,
		Address: address,
		Timeout: timeout,
	}

	handleID := storeSocketHandle(socket)
	return &object.Integer{Value: handleID}
}

func createWebSocket(address string, timeout time.Duration) object.Object {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:         address,
		Handler:      mux,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}

	socket := &WebSocket{
		Server:  server,
		Mux:     mux,
		Type:    SocketTypeWeb,
		Address: address,
		Timeout: timeout,
	}

	handleID := storeSocketHandle(socket)
	return &object.Integer{Value: handleID}
}

func createUnixSocket(address string, timeout time.Duration) object.Object {
	socket := &UnixSocket{
		Type:    SocketTypeUnix,
		Address: address,
		Timeout: timeout,
	}

	handleID := storeSocketHandle(socket)
	return &object.Integer{Value: handleID}
}

func connectTCPClient(address string, timeout time.Duration) object.Object {
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("failed to connect TCP client: %v", err)}
	}

	socket := &TCPSocket{
		Conn:    conn,
		Type:    SocketTypeTCP,
		Address: address,
		Timeout: timeout,
	}

	handleID := storeSocketHandle(socket)
	return &object.Integer{Value: handleID}
}

func connectUDPClient(address string, timeout time.Duration) object.Object {
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("failed to resolve UDP address: %v", err)}
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("failed to connect UDP client: %v", err)}
	}

	socket := &UDPSocket{
		Conn:    conn,
		Type:    SocketTypeUDP,
		Address: address,
		Timeout: timeout,
	}

	handleID := storeSocketHandle(socket)
	return &object.Integer{Value: handleID}
}

func connectUnixClient(address string, timeout time.Duration) object.Object {
	conn, err := net.DialTimeout("unix", address, timeout)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("failed to connect Unix client: %v", err)}
	}

	socket := &UnixSocket{
		Conn:    conn,
		Type:    SocketTypeUnix,
		Address: address,
		Timeout: timeout,
	}

	handleID := storeSocketHandle(socket)
	return &object.Integer{Value: handleID}
}

func startTCPServer(address string, timeout time.Duration) object.Object {
	// Allocate port with automatic incrementing
	allocatedAddress, message := allocatePort(address)
	
	listener, err := net.Listen("tcp", allocatedAddress)
	if err != nil {
		// Release the allocated port if binding failed
		releasePort(allocatedAddress)
		return &object.Error{Message: fmt.Sprintf("failed to start TCP server: %v", err)}
	}

	// Wrap the listener in TCPListener for consistency with other socket types
	tcpListener := &TCPListener{
		Listener: listener,
		Type:     "tcp_server",
		Address:  allocatedAddress,
		Timeout:  timeout,
	}

	handleID := storeSocketHandle(tcpListener)
	
	// Print message if port was incremented
	if message != "" {
		fmt.Println(message)
	}
	
	return &object.Integer{Value: handleID}
}

func startUDPServer(address string, timeout time.Duration) object.Object {
	// Allocate port with automatic incrementing
	allocatedAddress, message := allocatePort(address)
	
	udpAddr, err := net.ResolveUDPAddr("udp", allocatedAddress)
	if err != nil {
		releasePort(allocatedAddress)
		return &object.Error{Message: fmt.Sprintf("failed to resolve UDP address: %v", err)}
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		releasePort(allocatedAddress)
		return &object.Error{Message: fmt.Sprintf("failed to start UDP server: %v", err)}
	}

	socket := &UDPSocket{
		Conn:    conn,
		Type:    SocketTypeUDP,
		Address: allocatedAddress,
		Timeout: timeout,
	}

	handleID := storeSocketHandle(socket)
	
	// Print message if port was incremented
	if message != "" {
		fmt.Println(message)
	}
	
	return &object.Integer{Value: handleID}
}

func startWebServer(address string, timeout time.Duration) object.Object {
	// Allocate port with automatic incrementing
	allocatedAddress, message := allocatePort(address)
	
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:         allocatedAddress,
		Handler:      mux,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}

	socket := &WebSocket{
		Server:  server,
		Mux:     mux,
		Type:    SocketTypeWeb,
		Address: allocatedAddress,
		Timeout: timeout,
	}

	// Create a channel to synchronize server startup
	startupChan := make(chan error, 1)
	
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			startupChan <- err
		}
	}()

	// Give the server a moment to start and check for immediate errors
	select {
	case err := <-startupChan:
		// Server failed to start
		releasePort(allocatedAddress)
		return &object.Error{Message: fmt.Sprintf("failed to start web server: %v", err)}
	case <-time.After(100 * time.Millisecond):
		// Server appears to have started successfully (no immediate error)
		// The goroutine continues running in the background
	}

	handleID := storeSocketHandle(socket)
	
	// Print message if port was incremented
	if message != "" {
		fmt.Println(message)
	}
	
	return &object.Integer{Value: handleID}
}

func startUnixServer(address string, timeout time.Duration) object.Object {
	// Unix sockets use file paths, not ports, so we don't use port allocation
	// But we can still check if the socket file already exists and increment the filename
	originalAddress := address
	var message string
	
	// Try up to 100 different socket file names
	for attempts := 0; attempts < 100; attempts++ {
		listener, err := net.Listen("unix", address)
		if err != nil {
			// If error is because file exists, try incrementing the filename
			if strings.Contains(err.Error(), "address already in use") || strings.Contains(err.Error(), "bind: address already in use") {
				if attempts == 0 {
					// First increment, add suffix
					address = fmt.Sprintf("%s.%d", originalAddress, attempts+1)
					message = fmt.Sprintf("Unix socket %s already in use, trying %s", originalAddress, address)
				} else {
					// Subsequent increments
					address = fmt.Sprintf("%s.%d", originalAddress, attempts+1)
					message = fmt.Sprintf("Unix socket already in use, incremented to %s", address)
				}
				continue
			} else {
				return &object.Error{Message: fmt.Sprintf("failed to start Unix server: %v", err)}
			}
		}

		// Successfully bound to the socket
		handleID := storeSocketHandle(listener)
		
		// Print message if socket path was incremented
		if message != "" {
			fmt.Println(message)
		}
		
		return &object.Integer{Value: handleID}
	}

	return &object.Error{Message: fmt.Sprintf("could not find available Unix socket name starting from %s", originalAddress)}
}

func sendData(socket interface{}, data string) object.Object {
	switch s := socket.(type) {
	case *TCPSocket:
		if s.Conn == nil {
			return &object.Error{Message: "TCP socket not connected"}
		}
		n, err := s.Conn.Write([]byte(data))
		if err != nil {
			return &object.Error{Message: fmt.Sprintf("failed to send TCP data: %v", err)}
		}
		return &object.Integer{Value: int64(n)}

	case *UDPSocket:
		if s.Conn == nil {
			return &object.Error{Message: "UDP socket not connected"}
		}
		n, err := s.Conn.Write([]byte(data))
		if err != nil {
			return &object.Error{Message: fmt.Sprintf("failed to send UDP data: %v", err)}
		}
		return &object.Integer{Value: int64(n)}

	case *UnixSocket:
		if s.Conn == nil {
			return &object.Error{Message: "Unix socket not connected"}
		}
		n, err := s.Conn.Write([]byte(data))
		if err != nil {
			return &object.Error{Message: fmt.Sprintf("failed to send Unix data: %v", err)}
		}
		return &object.Integer{Value: int64(n)}

	default:
		return &object.Error{Message: "unsupported socket type for send operation"}
	}
}

func receiveData(socket interface{}, bufferSize int64) object.Object {
	switch s := socket.(type) {
	case *TCPSocket:
		if s.Conn == nil {
			return &object.Error{Message: "TCP socket not connected"}
		}
		buffer := make([]byte, bufferSize)
		n, err := s.Conn.Read(buffer)
		if err != nil && err != io.EOF {
			return &object.Error{Message: fmt.Sprintf("failed to receive TCP data: %v", err)}
		}
		return &object.String{Value: string(buffer[:n])}

	case *UDPSocket:
		if s.Conn == nil {
			return &object.Error{Message: "UDP socket not connected"}
		}
		buffer := make([]byte, bufferSize)
		n, err := s.Conn.Read(buffer)
		if err != nil && err != io.EOF {
			return &object.Error{Message: fmt.Sprintf("failed to receive UDP data: %v", err)}
		}
		return &object.String{Value: string(buffer[:n])}

	case *UnixSocket:
		if s.Conn == nil {
			return &object.Error{Message: "Unix socket not connected"}
		}
		buffer := make([]byte, bufferSize)
		n, err := s.Conn.Read(buffer)
		if err != nil && err != io.EOF {
			return &object.Error{Message: fmt.Sprintf("failed to receive Unix data: %v", err)}
		}
		return &object.String{Value: string(buffer[:n])}

	default:
		return &object.Error{Message: "unsupported socket type for receive operation"}
	}
}

func sendDataTo(socket interface{}, data string, targetAddress string) object.Object {
	switch s := socket.(type) {
	case *UDPSocket:
		if s.Conn == nil {
			return &object.Error{Message: "UDP socket not connected"}
		}
		
		// Parse target address
		udpAddr, err := net.ResolveUDPAddr("udp", targetAddress)
		if err != nil {
			return &object.Error{Message: fmt.Sprintf("failed to resolve target address: %v", err)}
		}
		
		// Send data to specific address
		n, err := s.Conn.WriteTo([]byte(data), udpAddr)
		if err != nil {
			return &object.Error{Message: fmt.Sprintf("failed to send UDP data to %s: %v", targetAddress, err)}
		}
		return &object.Integer{Value: int64(n)}
		
	default:
		return &object.Error{Message: "socket_send_to only supports UDP sockets"}
	}
}

func receiveDataFrom(socket interface{}, bufferSize int64) object.Object {
	switch s := socket.(type) {
	case *UDPSocket:
		if s.Conn == nil {
			return &object.Error{Message: "UDP socket not connected"}
		}
		
		buffer := make([]byte, bufferSize)
		n, addr, err := s.Conn.ReadFrom(buffer)
		if err != nil && err != io.EOF {
			return &object.Error{Message: fmt.Sprintf("failed to receive UDP data: %v", err)}
		}
		
		// Return a hash with data and sender address
		result := &object.Hash{
			Pairs: make(map[object.HashKey]object.HashPair),
		}
		
		dataKey := &object.String{Value: "data"}
		result.Pairs[dataKey.HashKey()] = object.HashPair{
			Key:   dataKey,
			Value: &object.String{Value: string(buffer[:n])},
		}
		
		senderKey := &object.String{Value: "sender"}
		result.Pairs[senderKey.HashKey()] = object.HashPair{
			Key:   senderKey,
			Value: &object.String{Value: addr.String()},
		}
		
		return result
		
	default:
		return &object.Error{Message: "socket_receive_from only supports UDP sockets"}
	}
}

func closeSocket(socket interface{}) error {
	// Check if socket is nil
	if socket == nil {
		return nil
	}

	switch s := socket.(type) {
	case *TCPSocket:
		// Check if socket struct is nil
		if s == nil {
			return nil
		}
		// Release port allocation
		if s.Address != "" {
			releasePort(s.Address)
		}
		if s.Conn != nil {
			return s.Conn.Close()
		}
	case *UDPSocket:
		// Check if socket struct is nil
		if s == nil {
			return nil
		}
		// Release port allocation
		if s.Address != "" {
			releasePort(s.Address)
		}
		if s.Conn != nil {
			return s.Conn.Close()
		}
	case *WebSocket:
		// Check if socket struct is nil
		if s == nil {
			return nil
		}
		// Release port allocation
		if s.Address != "" {
			releasePort(s.Address)
		}
		if s.Server != nil {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			return s.Server.Shutdown(ctx)
		}
	case *UnixSocket:
		// Check if socket struct is nil
		if s == nil {
			return nil
		}
		if s.Conn != nil {
			return s.Conn.Close()
		}
	case *TCPListener:
		// Check if socket struct is nil
		if s == nil {
			return nil
		}
		// Release port allocation for TCP listener
		if s.Address != "" {
			releasePort(s.Address)
		}
		if s.Listener != nil {
			return s.Listener.Close()
		}
	case net.Listener:
		// Check if listener interface is nil
		if s == nil {
			return nil
		}
		// For raw listeners (TCP/Unix servers), we need to get the address to release the port
		addr := s.Addr()
		if addr != nil && addr.Network() == "tcp" {
			releasePort(addr.String())
		}
		return s.Close()
	}
	return nil
}

func listenForConnections(socket interface{}) object.Object {
	switch s := socket.(type) {
	case *TCPListener:
		// Return the same handle since it's already a listener wrapper
		handleID := storeSocketHandle(s)
		return &object.Integer{Value: handleID}
	case net.Listener:
		// Return the same handle since it's already a listener
		handleID := storeSocketHandle(s)
		return &object.Integer{Value: handleID}
	case *TCPSocket:
		if listener, ok := s.Conn.(net.Listener); ok {
			// Return listener handle for accepting connections
			return &object.Integer{Value: storeSocketHandle(listener)}
		}
		return &object.Error{Message: "TCP socket is not a listener"}
	case *UnixSocket:
		if listener, ok := s.Conn.(net.Listener); ok {
			return &object.Integer{Value: storeSocketHandle(listener)}
		}
		return &object.Error{Message: "Unix socket is not a listener"}
	default:
		return &object.Error{Message: "unsupported socket type for listen operation"}
	}
}

func acceptConnection(socket interface{}) object.Object {
	switch s := socket.(type) {
	case *TCPListener:
		conn, err := s.Listener.Accept()
		if err != nil {
			return &object.Error{Message: fmt.Sprintf("failed to accept connection: %v", err)}
		}

		newSocket := &TCPSocket{
			Conn:    conn,
			Type:    SocketTypeTCP,
			Address: conn.RemoteAddr().String(),
			Timeout: s.Timeout,
		}

		handleID := storeSocketHandle(newSocket)
		return &object.Integer{Value: handleID}

	case net.Listener:
		conn, err := s.Accept()
		if err != nil {
			return &object.Error{Message: fmt.Sprintf("failed to accept connection: %v", err)}
		}

		var newSocket interface{}
		switch s.Addr().Network() {
		case "tcp":
			newSocket = &TCPSocket{
				Conn:    conn,
				Type:    SocketTypeTCP,
				Address: conn.RemoteAddr().String(),
				Timeout: 30 * time.Second,
			}
		case "unix":
			newSocket = &UnixSocket{
				Conn:    conn,
				Type:    SocketTypeUnix,
				Address: conn.RemoteAddr().String(),
				Timeout: 30 * time.Second,
			}
		default:
			conn.Close()
			return &object.Error{Message: "unsupported listener type"}
		}

		handleID := storeSocketHandle(newSocket)
		return &object.Integer{Value: handleID}

	default:
		return &object.Error{Message: "socket is not a listener"}
	}
}

func setSocketTimeout(socket interface{}, timeout time.Duration) object.Object {
	switch s := socket.(type) {
	case *TCPSocket:
		s.Timeout = timeout
		if s.Conn != nil {
			if tcpConn, ok := s.Conn.(*net.TCPConn); ok {
				tcpConn.SetDeadline(time.Now().Add(timeout))
			}
		}
	case *UDPSocket:
		s.Timeout = timeout
		if s.Conn != nil {
			s.Conn.SetDeadline(time.Now().Add(timeout))
		}
	case *WebSocket:
		s.Timeout = timeout
		s.Server.ReadTimeout = timeout
		s.Server.WriteTimeout = timeout
	case *UnixSocket:
		s.Timeout = timeout
		if s.Conn != nil {
			s.Conn.SetDeadline(time.Now().Add(timeout))
		}
	case *TCPListener:
		s.Timeout = timeout
		// Note: net.Listener doesn't have SetDeadline, timeout is used for accepted connections
	default:
		return &object.Error{Message: "unsupported socket type for timeout operation"}
	}

	return &object.None{}
}

func getSocketInfo(socket interface{}) object.Object {
	result := &object.Hash{
		Pairs: make(map[object.HashKey]object.HashPair),
	}

	var socketType, address string
	var timeout time.Duration

	switch s := socket.(type) {
	case *TCPSocket:
		socketType = s.Type
		address = s.Address
		timeout = s.Timeout
	case *UDPSocket:
		socketType = s.Type
		address = s.Address
		timeout = s.Timeout
	case *WebSocket:
		socketType = s.Type
		address = s.Address
		timeout = s.Timeout
	case *UnixSocket:
		socketType = s.Type
		address = s.Address
		timeout = s.Timeout
	case *TCPListener:
		socketType = s.Type
		address = s.Address
		timeout = s.Timeout
	case net.Listener:
		// Handle listeners created by server functions
		addr := s.Addr()
		switch addr.Network() {
		case "tcp":
			socketType = "tcp_listener"
		case "unix":
			socketType = "unix_listener"
		default:
			socketType = "listener"
		}
		address = addr.String()
		timeout = 30 * time.Second // Default timeout for listeners
	default:
		return &object.Error{Message: "unsupported socket type for info operation"}
	}

	typeKey := &object.String{Value: "type"}
	result.Pairs[typeKey.HashKey()] = object.HashPair{
		Key:   typeKey,
		Value: &object.String{Value: socketType},
	}

	addressKey := &object.String{Value: "address"}
	result.Pairs[addressKey.HashKey()] = object.HashPair{
		Key:   addressKey,
		Value: &object.String{Value: address},
	}

	timeoutKey := &object.String{Value: "timeout"}
	result.Pairs[timeoutKey.HashKey()] = object.HashPair{
		Key:   timeoutKey,
		Value: &object.Integer{Value: int64(timeout.Seconds())},
	}

	return result
}