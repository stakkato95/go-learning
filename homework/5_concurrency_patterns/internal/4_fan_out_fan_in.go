package internal

//1 start processing long running data in X goroutines
//2 combine result from X goroutines in on channel (see converter/multiplexer in 1_produce_consumer.go)
