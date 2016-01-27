#Overview
This app finds the longest compound word from a sortest list of words

##Set GOPATH
Go the the top directory and set the GOPATH using:
```
$ export GOPATH="$GOPATH":"$PWD"
```

##Build and Run the longest compound word solver App
Flags:<br>
  * -f filename
  * -b enable benchmarking

```
$ go run main.go -f words.list -b
```

##Test the longest compound word solver App
```
$ cd src/lcwsolver
$ go test
```