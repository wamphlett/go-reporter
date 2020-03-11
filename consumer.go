package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Consumer struct {
	Queue   chan *Report
	Reports map[string][]*Report
	stats   map[string]*Stats
}

type Stats struct {
	Valid, Invalid int
}

func NewConsumer() *Consumer {

	c := &Consumer{Queue: make(chan *Report, 100), Reports: make(map[string][]*Report, 1000), stats: make(map[string]*Stats)}

	go func() {
		for r := range c.Queue {
			c.processReport(r)
		}
	}()

	return c
}

func (c *Consumer) consumeReport(report *Report) {
	fmt.Println("Adding new report to the queue")
	c.Queue <- report
}

func (c *Consumer) processReport(report *Report) {
	// Simulate some sort of connection to an external resource
	time.Sleep(time.Duration(rand.Intn(1000-100)+100) * time.Millisecond)

	fmt.Println("Reading new report from the queue")

	if _, ok := c.stats[report.Identifier]; !ok {
		fmt.Println("NEW ID")
		c.stats[report.Identifier] = &Stats{}
	}

	if len(report.Identifier) == 0 {
		c.stats[report.Identifier].Invalid++
		PrintError("Report is missing an identifier")
		return
	}

	if len(report.Message) == 0 {
		c.stats[report.Identifier].Invalid++
		PrintError("Report must have a message")
		return
	}

	if len(report.Type) == 0 {
		report.Type = "error"
	} else {
		if !IsInSlice(report.Type, []string{"error", "warn", "info"}) {
			c.stats[report.Identifier].Invalid++
			PrintError("Unknown report type, expected: error, warn, info")
			return
		}
	}

	fmt.Println("Accepted new report")
	c.stats[report.Identifier].Valid++

	if _, ok := c.Reports[report.Identifier]; !ok {
		c.Reports[report.Identifier] = []*Report{}
	}

	c.Reports[report.Identifier] = append(c.Reports[report.Identifier], report)
}
