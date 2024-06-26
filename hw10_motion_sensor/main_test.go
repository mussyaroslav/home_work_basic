package main

import (
	"testing"
	"time"
)

func TestSensorSimulator(t *testing.T) {
	duration := 5 // 5 seconds
	sensorData := SensorSimulator(duration)

	timeout := time.After(time.Duration(duration+1) * time.Second)
	count := 0
loop:
	for {
		select {
		case <-timeout:
			break loop
		case _, ok := <-sensorData:
			if !ok {
				break loop
			}
			count++
		}
	}

	if count < duration-1 {
		t.Errorf("Expected at least %d readings, got %d", duration-1, count)
	}
}

func TestDataProcessor(t *testing.T) {
	sensorData := make(chan int, 20)
	for i := 1; i <= 20; i++ {
		sensorData <- i
	}
	close(sensorData)

	processedData := DataProcessor(sensorData)

	expectedAverages := []float32{5.5, 15.5}
	i := 0
	for avg := range processedData {
		if avg != expectedAverages[i] {
			t.Errorf("Expected average %f, got %f", expectedAverages[i], avg)
		}
		i++
	}
}

func TestDisplayAverages(t *testing.T) {
	averages := make(chan float32, 2)
	averages <- 5.5
	averages <- 15.5
	close(averages)
	DisplayAverages(averages)
}
