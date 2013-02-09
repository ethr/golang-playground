// This program shows how channels can be used to communicate between different
// go routines.
// Written in an object oriented way, which may not be the best way to do this,
// with 2 interfaces Handler & LoadBalance and concrete implementations below.
// Golang doesn't work well with this approach atall well and the interface
// business is pretty screwed up.
//
// Basic idea:
// Load balance manages a set of Handlers, forwarding requests to them as
// required and returning the result down the channel. Each Handler runs in its
// own go routine (thread), so does the LoadBalancer.
package main

import "fmt"
import "strconv"

type Handler interface {
        Run()
}

type LoadBalancer interface {
        Register(c chan string, handler Handler) bool
        Run()
}

/// A basic round-robin load balancer
type LoadBalancerImpl struct {
        next int
        input chan string
        handleChans []chan string
}

// Implement the Register function for LoadBalancerImpl to implement
// LoadBalancer interface
func (this *LoadBalancerImpl) Register(c chan string) bool {
        this.handleChans = append(this.handleChans, c)
        return true
}

func (this *LoadBalancerImpl) run() {
        /// Blocks until something is sent to the channel the iterates
        for value := range this.input {
                c := this.handleChans[this.next]
                c <- value
                v := <- c
                this.input <- v
                this.next++
                if this.next >= len(this.handleChans) {
                        this.next = 0
                }
        }
}

func (this *LoadBalancerImpl) Run() {
        go this.run()
}

func NewLoadBalancerImpl(input chan string) *LoadBalancerImpl {
        this := new(LoadBalancerImpl)
        this.input = input
        this.next = 0
        return this
}

// Echo's what is sent down the channel
type EchoHandler struct {
        input chan string
}

func NewEchoHandler(input chan string) *EchoHandler {
        this := new(EchoHandler)
        this.input = input
        return this
}

func (this *EchoHandler) run() {
        for value := range this.input {
                this.input <- value
        }
}

func (this *EchoHandler) Run() {
        go this.run()
}

// Handler which returns the length of any string sent down the channel
type LengthHandler struct {
        input chan string
}

func NewLengthHandler(input chan string) *LengthHandler {
        this := new(LengthHandler)
        this.input = input
        return this
}

func (this *LengthHandler) run() {
        for value := range this.input {
                this.input <- strconv.Itoa(len(value))
        }
}

func (this *LengthHandler) Run() {
        go this.run()
}

// Handler which reverses a string
type ReverseHandler struct {
        input chan string
}

func NewReverseHandler(input chan string) *ReverseHandler {
        this := new(ReverseHandler)
        this.input = input
        return this
}

func (this *ReverseHandler) run() {
        for value := range this.input {
                runes := []rune(value)
                for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
                        runes[i], runes[j] = runes[j], runes[i]
                }
                this.input <- string(runes)
        }
}

func (this *ReverseHandler) Run() {
        go this.run()
}

func main() {
        mainChan := make(chan string)
        h1chan := make(chan string)
        h2chan := make(chan string)
        h3chan := make(chan string)

        h1 := NewEchoHandler(h1chan)
        h2 := NewLengthHandler(h2chan)
        h3 := NewReverseHandler(h3chan)
        lb := NewLoadBalancerImpl(mainChan)

        lb.Register(h1chan)
        lb.Register(h2chan)
        lb.Register(h3chan)

        h1.Run()
        h2.Run()
        h3.Run()
        lb.Run()

        for i := 0; i < 3; i++ {
                mainChan <- "foo"
                ret := <- mainChan
                fmt.Println(ret)
        }
}
