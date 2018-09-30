package main

import (
  "net"
  "bufio"
  "fmt"
  "os"
//  "sync"
)

var (
    conn, err = net.Dial("tcp", ":8080")
    //wg sync.WaitGroup
)

func getConn() {
  if err != nil {
    panic(err)
  }

  fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
  go sendMessage()

}

func sendMessage() {
  reader := bufio.NewReader(os.Stdin)
  text, _ := reader.ReadString('\n')
  for {
     if text == "quit" {
       break;
     }
     fmt.Fprintf(conn,text)
  }

}

func main() {
  go getConn()
  select {}
//  wg.Wait()
//  conn.Close()
}
