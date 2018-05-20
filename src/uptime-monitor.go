package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/url"
)

type UpTimeMonitorService struct {
	apiKey        string
	url           string
	alertContacts string
}

func (monitor *UpTimeMonitorService) Setup(apiKey string, url string, alertContacts string) {
	monitor.apiKey = apiKey
	monitor.url = url
	monitor.alertContacts = alertContacts
}

func (monitor *UpTimeMonitorService) GetByName(name string) (*Monitor, error) {
	action := "getMonitors"

	client := createHttpClient(monitor.url + action)

	body := "api_key=" + monitor.apiKey + "&format=json&logs=1" + "&search=" + name

	response, err := client.postUrlEncodedFormBody(body)
	if err != nil {
		return nil, err
	}

	if response.statusCode == 200 {

		var f UptimeMonitorGetMonitorsResponse
		json.Unmarshal(response.bytes, &f)

		if f.Monitors != nil && len(f.Monitors) > 0 {
			return UptimeMonitorMonitorToBaseMonitorMapper(f.Monitors[0]), nil
		}
		return nil, nil
	}

	errorString := "GetByName Request failed for name: " + name

	log.Println(errorString)
	return nil, errors.New(errorString)
}

func (monitor *UpTimeMonitorService) GetAll() []Monitor {

	action := "getMonitors"

	client := createHttpClient(monitor.url + action)

	body := "api_key=" + monitor.apiKey + "&format=json&logs=1"

	response, err := client.postUrlEncodedFormBody(body)
	if err != nil {
		log.Println(err)
		return nil
	}

	if response.statusCode == 200 {

		var f UptimeMonitorGetMonitorsResponse
		json.Unmarshal(response.bytes, &f)

		return UptimeMonitorMonitorsToBaseMonitorsMapper(f.Monitors)

	}

	log.Println("GetAllMonitors Request failed. Status Code: " + string(response.statusCode))
	return nil

}

func (monitor *UpTimeMonitorService) Add(m Monitor) {
	action := "newMonitor"

	client := createHttpClient(monitor.url + action)

	body := "api_key=" + monitor.apiKey + "&format=json&type=1&url=" + url.QueryEscape(m.url) + "&friendly_name=" + url.QueryEscape(m.name) + "&alert_contacts=" + monitor.alertContacts

	response, err := client.postUrlEncodedFormBody(body)
	if err != nil {
		log.Println("failed to add monitor: ", err)
		return
	}

	if response.statusCode == 200 {
		var f UptimeMonitorNewMonitorResponse
		json.Unmarshal(response.bytes, &f)

		if f.Stat == "ok" {
			log.Println("Monitor Added: " + m.name)
			log.Println(string(response.bytes))

		} else {
			log.Println("Monitor couldn't be added: " + m.name)
			log.Println(string(body))
			log.Println(string(response.bytes))
		}
	} else {
		log.Printf("AddMonitor Request failed. Status Code: " + string(response.statusCode))
	}
}

func (monitor *UpTimeMonitorService) Update(m Monitor) {
	action := "editMonitor"

	client := createHttpClient(monitor.url + action)

	body := "api_key=" + monitor.apiKey + "&format=json&id=" + m.id + "&friendly_name=" + m.name + "&url=" + m.url

	response, err := client.postUrlEncodedFormBody(body)
	if err != nil {
		log.Println("failed to update monitor", err)
		return
	}

	if response.statusCode == 200 {
		var f UptimeMonitorStatusMonitorResponse
		json.Unmarshal(response.bytes, &f)

		if f.Stat == "ok" {
			log.Println("Monitor Updated: " + m.name)
		} else {
			log.Println("Monitor couldn't be updated: " + m.name)
			log.Println(string(body))
			log.Println(string(response.bytes))
		}
	} else {
		log.Println("UpdateMonitor Request failed. Status Code: " + string(response.statusCode))
	}
}

func (monitor *UpTimeMonitorService) Remove(m Monitor) {
	action := "deleteMonitor"

	client := createHttpClient(monitor.url + action)

	body := "api_key=" + monitor.apiKey + "&format=json&id=" + m.id

	response, err := client.postUrlEncodedFormBody(body)
	if err != nil {
		log.Println("failde to remove monitor", err)
		return
	}

	if response.statusCode == 200 {
		var f UptimeMonitorStatusMonitorResponse
		json.Unmarshal(response.bytes, &f)

		if f.Stat == "ok" {
			log.Println("Monitor Removed: " + m.name)
		} else {
			log.Println("Monitor couldn't be removed: " + m.name)
			log.Println(string(body))
			log.Println(string(response.bytes))
		}
	} else {
		log.Println("RemoveMonitor Request failed. Status Code: " + string(response.statusCode))
	}
}
