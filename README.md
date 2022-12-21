# golang-concurrency
Golang Concurrency

## 1. ping-pong

### Naive
```sh
$ go test -v 1.ping-pong/naive_test.go
```

### Referree
```sh
$ go test -v 1.ping-pong/referree_test.go
```

## 2. golang-concurrency-patterns

### Patterns 1 (pipeline)
```sh
$ go test -v 2.patterns/patterns1_test.go
```

### Patterns 2 (fan in/fan out)
```sh
$ go test -v 2.patterns/patterns2_test.go
```

### Patterns 3 (fan in/fan out)
```sh
$ go test -v 2.patterns/patterns3_test.go
```

### Or, all of them Patterns...
```sh
$ go run ./2.patterns
```

```sh
$ go run ./2.patterns/main.go
```

## 3. bank

### Mutex
```sh
$ go test -v 3.bank/bank_account_mutex_test.go
```

### Channel
```sh
$ go test -v 3.bank/bank_account_channel_test.go
```

## 4. get-url

### get url
```sh
$ go test -v 4.get-url/get_url_test.go
```

### stopper
#### http wait
```sh
$ go run 4.get-url/wait/main.go
```

#### testing stopper
```sh
$ go test -v 4.get-url/get_url_stopper_test.go
```

## 5. multi

### multi
```sh
$ go test -v 5.multi/multi_test.go
```