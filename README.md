#Overview
This app finds the longest compound-word in the list, which is also a concatenation of other sub-words that exist in the list.

##Set GOPATH
Go the the top directory and set the GOPATH using:
```
$ export GOPATH="$GOPATH":"$PWD"
```

##Build and Run the App
Flags:<br>
  * -f filename
  * -b enable benchmarking

```
$ go run main.go -f words.list -b
```

##Test the App
```
$ cd src/lcwsolver
$ go test
```