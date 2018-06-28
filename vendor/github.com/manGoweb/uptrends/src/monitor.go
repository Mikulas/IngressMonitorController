package v3

import (
	"net/http"
	"encoding/json"
	"bytes"
	"net/url"
	"strconv"
	"fmt"
)

type Monitor struct {
	Guid string `json:"Guid,omitempty"`
	Name string
	URL string
	Port uint
	CheckFrequency uint // minutes
	ProbeType string
	IsActive bool
	GenerateAlert bool
	Notes string
	PerformanceLimit1 uint
	PerformanceLimit2 uint
	ErrorOnLimit1 bool
	ErrorOnLimit2 bool
	MinBytes uint
	ErrorOnMinBytes bool
	Timeout uint // milliseconds
	TcpConnectTimeout uint // milliseconds
	DnsLookupMode string
	UserAgent string
	UserName string
	Password string
	IsCompetitor bool
	Checkpoints string
	UseOnlyPrimaryCheckpoints bool
	Action string

	DNSQueryType string `json:"DNSQueryType,omitempty"`
	DNSTestValue string `json:"DNSTestValue,omitempty"`
	DNSExpectedResult string `json:"DNSExpectedResult,omitempty"`

	DatabaseName string `json:"DatabaseName,omitempty"`

	HttpMethod string `json:"HttpMethod,omitempty"`
	PostData string `json:"PostData,omitempty"`

	ConnectMethod string `json:"ConnectMethod,omitempty"`
}

func MakeMonitor(name, urlStr string) *Monitor {
	urlObj, _ := url.Parse(urlStr)

	portStr := urlObj.Port()
	port := uint64(80)
	if urlObj.Scheme == "https" {
		port = 443
	}
	if portStr != "" {
		port, _ = strconv.ParseUint(portStr, 10, 32)
	}

	probeType := "Http"
	if urlObj.Scheme == "https" {
		probeType = "Https"
	}

	return &Monitor{
		URL: urlStr,
		ProbeType: probeType,
		Port: uint(port),
		CheckFrequency: 1,
		PerformanceLimit1: 2000,
		ErrorOnLimit1: false,
		PerformanceLimit2: 5000,
		ErrorOnLimit2: false,
		GenerateAlert: true,
		HttpMethod: "Get",
		MinBytes: 0,
		IsActive: true,
		Name: name,
		Timeout: 15000,
		TcpConnectTimeout: 5000,
	}
}

func (up *Uptrends) GetMonitors() ([]Monitor, error) {
	var monitors []Monitor
	req, err := http.NewRequest(http.MethodGet, "/probes", nil)
	if err != nil {
		return monitors, err
	}
	resp, err := up.request(req)
	if err != nil {
		return monitors, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&monitors)
	return monitors, err
}

func (up *Uptrends) GetMonitor(guid string) (Monitor, error) {
	monitor := Monitor{}
	endpoint := fmt.Sprintf("/probes/%v", url.QueryEscape(guid))
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return monitor, err
	}
	resp, err := up.request(req)
	if err != nil {
		return monitor, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&monitor)
	return monitor, err
}

func (up *Uptrends) AddMonitor(monitor *Monitor) error {
	body, err := json.Marshal(monitor)
	req, err := http.NewRequest(http.MethodPost, "/probes", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	resp, err := up.request(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respMonitor := Monitor{}
	err = json.NewDecoder(resp.Body).Decode(&respMonitor)
	if err != nil {
		return err
	}
	*monitor = respMonitor
	return nil
}

func (up *Uptrends) EditMonitor(monitor *Monitor) error {
	if monitor.Guid == "" {
		return fmt.Errorf("unset guid, add the monitor instead")
	}

	body, err := json.Marshal(monitor)
	endpoint := fmt.Sprintf("/probes/%v", url.QueryEscape(monitor.Guid))
	req, err := http.NewRequest(http.MethodPut, endpoint, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	resp, err := up.request(req)
	defer resp.Body.Close()
	return err
}

func (up *Uptrends) DeleteMonitor(guid string) error {
	endpoint := fmt.Sprintf("/probes/%v", url.QueryEscape(guid))
	req, err := http.NewRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return err
	}
	resp, err := up.request(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
