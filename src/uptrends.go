package main

import (
	"github.com/manGoweb/uptrends/src"
	"strings"
	"fmt"
	"time"
	"math"
)

type Uptrends struct {
	client *v3.Uptrends
}

func (m Monitor) toUptrends() *v3.Monitor {
	upMonitor := v3.MakeMonitor(m.name, m.url)
	upMonitor.Guid = m.id
	upMonitor.CheckFrequency = uint(math.Floor(m.interval.Minutes()))
	return upMonitor
}

func toMonitor(upMonitor v3.Monitor) Monitor {
	return Monitor{
		name: upMonitor.Name,
		url: upMonitor.URL,
		interval: time.Duration(upMonitor.CheckFrequency) * time.Minute,
		id: upMonitor.Guid,
	}
}

func (up *Uptrends) Setup(apiKey, url, alertContacts string) error {
	if alertContacts != "" {
		return fmt.Errorf("alertContacts should be empty, its not used")
	}
	if url != "" {
		return fmt.Errorf("url should be empty, its not used")
	}

	credentials := strings.SplitN(apiKey, ":", 2)
	if len(credentials) != 2 {
		return fmt.Errorf("instead of api key, pass a 'username:password' string")
	}

	up.client = v3.MakeUptrends(credentials[0], credentials[1])
	return nil
}

func (up *Uptrends) GetAll() ([]Monitor, error) {
	var monitors []Monitor
	upMonitors, err := up.client.GetMonitors()
	if err != nil {
		return monitors, err
	}

	for _, upMonitor := range upMonitors {
		monitors = append(monitors, toMonitor(upMonitor))
	}
	return monitors, nil
}

func (up *Uptrends) Add(m Monitor) error {
	return up.client.AddMonitor(m.toUptrends())
}

func (up *Uptrends) Update(m Monitor) error {
	return up.client.EditMonitor(m.toUptrends())
}

func (up *Uptrends) GetByName(name string) (*Monitor, error) {
	monitors, err := up.GetAll()
	if err != nil {
		return nil, err
	}
	for _, monitor := range monitors {
		if monitor.name == name {
			return &monitor, nil
		}
	}
	return nil, fmt.Errorf("monitor with name '%v' not found", name)
}

func (up *Uptrends) Remove(m Monitor) error {
	if m.id != "" {
		return up.client.DeleteMonitor(m.id)
	}

	monitor, err := up.GetByName(m.name)
	if err != nil {
		return err
	}
	return up.client.DeleteMonitor(monitor.id)
}
