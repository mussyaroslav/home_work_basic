package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func SensorSimulator(duration int) chan int {
	c := make(chan int)
	timeout := time.After(time.Duration(duration) * time.Second)
	go func() {
		for {
			s, _ := rand.Int(rand.Reader, big.NewInt(100))
			r := int(s.Int64())
			select {
			case <-timeout:
				close(c)
				return
			case c <- r:
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return c
}

func DataProcessor(sensor <-chan int) chan float32 {
	out := make(chan float32)
	go func() {
		var sum float32
		count := 0
		for v := range sensor {
			sum += float32(v)
			count++
			if count == 10 {
				out <- sum / 10
				sum = 0
				count = 0
			}
		}
		close(out)
	}()
	return out
}

func DisplayAverages(averages <-chan float32) {
	for v := range averages {
		fmt.Println("Average is:", v)
	}
}

func main() {
	sensorData := SensorSimulator(60)
	processedData := DataProcessor(sensorData)
	DisplayAverages(processedData)
}
