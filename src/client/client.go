package main

import "log"
import "net/http"
import "io"
import "io/ioutil"
import "os"
import "time"

var ( 
  Info    *log.Logger
  Warning *log.Logger
)

func Init(
  infoHandle io.Writer,
  warningHandle io.Writer ) { 

  Info = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
  Warning = log.New(warningHandle, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile) 
}

func main() {
  Init(os.Stdout, os.Stdout)

  tr := &http.Transport{
    DisableKeepAlives: true, 
  }
  
  client := &http.Client{Transport: tr}

  for {
    time.Sleep(500 * time.Millisecond)
    resp, err := client.Get("http://server:4567/")
    if err != nil {
      Warning.Printf("Failed")
    } 
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    Info.Printf("Received %s", body)
  }
}
