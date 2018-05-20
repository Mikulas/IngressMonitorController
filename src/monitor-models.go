package main

import "time"

type Monitor struct {
	url  string
	name string
	id   string
	interval time.Duration
}
