// Basic timer

package main

import "time"
import "fmt"

type TimerCallback func()

type Timer interface {
        // Start the timer
        Start()

        // Stop the timer
        Stop()

        // Add callback
        AddCallback(callback TimerCallback)
}

type BasicTimer struct {
        callbacks []TimerCallback
        stopTimer chan int
        length int
}

func NewBasicTimer(length int) *BasicTimer {
        this := new(BasicTimer)
        this.length = length
        return this
}

func (this *BasicTimer) Length() int {
        return this.length
}

func (this *BasicTimer) runTimer() {
        for {
                select {
                case <-this.stopTimer:
                        return
                case <-time.After(time.Duration(this.Length()) * time.Millisecond):
                        for cb := range(this.callbacks) {
                                this.callbacks[cb]()
                        }
                }
        }
}

func (this *BasicTimer) Start() {
        go this.runTimer()
}

func (this *BasicTimer) Stop() {
        this.stopTimer <- 1
}

func (this *BasicTimer) AddCallback(callback TimerCallback) {
        this.callbacks = append(this.callbacks, callback)
}

func main() {
        t1 := NewBasicTimer(1000)
        start := time.Now().Unix()
        cb1 := func() {
                diff := time.Now().Unix() - start
                fmt.Printf("Done %d seconds later\n", diff)
        }
        t1.AddCallback(cb1)
        t1.Start()
        time.Sleep(time.Duration(1200) * time.Millisecond);
}
