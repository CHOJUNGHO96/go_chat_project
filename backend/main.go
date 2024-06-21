package main

import "go_chat_project/network"

func main() {
	n := network.NewServer()
	n.StartServer()
}
