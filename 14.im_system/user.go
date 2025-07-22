package main

import (
	"net"
	"fmt"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	server *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	fmt.Println("New user, conn: ", conn)
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
		server: server,
	}

	go user.ListenMessage()

	return user
}

func (this* User) Online() {
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	//fmt.Println("add user to map.")
	this.server.mapLock.Unlock()

	this.server.BroadCast(this, "online")
}

func (this* User) Offline() {
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	//fmt.Println("add user to map.")
	this.server.mapLock.Unlock()

	this.server.BroadCast(this, "offline")
}

func (this* User) DoMessage(msg string) {
	this.server.BroadCast(this, msg)
}
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		fmt.Println("User recv msg: ", msg)
		this.conn.Write([]byte(msg + "\n"))
	}
}
