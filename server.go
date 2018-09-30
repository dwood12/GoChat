package main

import (
  "net"
  "fmt"
  "bufio"
//  "sync"
)

var (
    ln, listenErr = net.Listen("tcp", ":8080")
    conn, connErr = ln.Accept()
  //  wg sync.WaitGroup
)

func getConn() {

  if listenErr != nil {
    panic(listenErr)
  }

  if connErr != nil {
    panic(connErr)
  }

  for {
    go handleMessage()
  }

}

func handleMessage() {
    message, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
      panic(err)
    }
    fmt.Println(message)
}

func main() {
  go getConn()
  select {}
//  wg.Wait()

//  ln.Close()
//  conn.Close()
}
