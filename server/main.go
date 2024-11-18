package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"plugin"
	"shared/interfaces"
)

func main() {
	content, err := os.ReadFile("./players/greedy.go")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	ReceivePlayerImpl(string(content), "NewGreedyPlayer")
}

func Run() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error creating listener:", err)
		os.Exit(1)
	}
	defer listener.Close()

	go handleConnection(listener)

	fmt.Println("Server is listening on port 8080")

}

func handleConnection(listener net.Listener) {
	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		defer conn.Close()

		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}

		receivedString := string(buffer[:n])
		fmt.Println("Received string:", receivedString)
		ReceivePlayerImpl(receivedString, "NewGreedyPlayer")

		break
	}
}

func ReceivePlayerImpl(code string, constructor string) {
	file, _ := os.Create("player_impl.go")
	file.WriteString(code)
	file.Close()

	// TODO: add unique identifier to players uploaded
	_, err := exec.Command("go", "build", "-buildmode=plugin", "-o", "impl.so", "player_impl.go").Output()
	if err != nil {
		fmt.Printf("error in exec: %s\n", err)
	}
	plug, err := plugin.Open("impl.so")
	if err != nil {
		fmt.Printf("error plugin: %s\n", err)
	}
	playerImpl, err := plug.Lookup(constructor)
	if err != nil {
		fmt.Printf("error in lookup: %s\n", err)
	}
	// TODO: create shared interface for player definition
	loadedPlayer := playerImpl.(func(int) interfaces.Player)(1)

	loadedPlayer.Move()
}