# go-pythonic-fileserver
### A Pythonic-style HTTP file server written in Go
By default, the server listens on all IP addresses present on the system ('0.0.0.0') and is bound to TCP port 8000.  The current directory ('.') is shared, relative to the location of 'fileserver.go' or the binary (if compiled.)

#### Usage:

```go run fileserver.go [optional flags]```

#### The flags are:

<b>-p</b> : The TCP port to listen on.  Defaults to 8000

<b>-i</b> : The interface IP address to use.  Defaults to 0.0.0.0

<b>-d</b> : The directory path to share.  Defaults to '.'