package main

import (
	"log"
)

// Service interface does not allow errors,
// this service handles the errors instead of
// silently discarding them

type UptrendsLogging struct {
	svc *Uptrends
}

func (up *UptrendsLogging) Setup(apiKey, url, alertContacts string) {
	up.svc = &Uptrends{}
	err := up.svc.Setup(apiKey, url, alertContacts)
	if err != nil {
		log.Printf("uptrends setup: %v\n", err)
	}
}

func (up *UptrendsLogging) GetAll() []Monitor {
	monitors, err := up.svc.GetAll()
	if err != nil {
		log.Printf("uptrends get all: %v\n", err)
	}
	return monitors
}

func (up *UptrendsLogging) Add(m Monitor) {
	err := up.svc.Add(m)
	if err != nil {
		log.Printf("uptrends add: %v\n", err)
	}
}

func (up *UptrendsLogging) Update(m Monitor) {
	err := up.svc.Update(m)
	if err != nil {
		log.Printf("uptrends update: %v\n", err)
	}
}

func (up *UptrendsLogging) GetByName(name string) (*Monitor, error) {
	monitor, err := up.svc.GetByName(name)
	if err != nil {
		log.Printf("uptrends get by name: %v\n", err)
	}
	return monitor, err
}

func (up *UptrendsLogging) Remove(m Monitor) {
	err := up.svc.Remove(m)
	if err != nil {
		log.Printf("uptrends remove: %v\n", err)
	}
}
