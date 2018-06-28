package v3

import (
	"fmt"
	"net/http"
	"net/url"
	"encoding/json"
	"bytes"
)

func (up *Uptrends) AddMember(group *Group, monitor *Monitor) error {
	return up.addOrRemoveMember(group, monitor, http.MethodPost)
}

func (up *Uptrends) RemoveMember(group *Group, monitor *Monitor) error {
	return up.addOrRemoveMember(group, monitor, http.MethodDelete)
}

func (up *Uptrends) addOrRemoveMember(group *Group, monitor *Monitor, method string) error {
	if group.Guid == "" {
		return fmt.Errorf("group is not persisted")
	}
	if monitor.Guid == "" {
		return fmt.Errorf("monitor is not persisted")
	}

	payload := struct {
		ProbeGuid string
	}{
		ProbeGuid: monitor.Guid,
	}
	body, _ := json.Marshal(payload)

	endpoint := fmt.Sprintf("/probegroups/%v/members", url.QueryEscape(group.Guid))
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(body))
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

func (up *Uptrends) GetMembers(group *Group) ([]Monitor, error) {
	var monitors []Monitor
	if group.Guid == "" {
		return monitors, fmt.Errorf("group is not persisted")
	}

	endpoint := fmt.Sprintf("/probegroups/%v/members", url.QueryEscape(group.Guid))
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
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
