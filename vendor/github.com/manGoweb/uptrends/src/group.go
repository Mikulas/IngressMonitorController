package v3

import (
	"net/http"
	"encoding/json"
	"bytes"
	"net/url"
	"fmt"
)

type Group struct {
	Guid string `json:"Guid,omitempty"`
	Name string
}

func MakeGroup(name string) *Group {
	return &Group{Name: name}
}

func (up *Uptrends) GetGroups() ([]Group, error) {
	groups := []Group{}
	req, err := http.NewRequest(http.MethodGet, "/probegroups", nil)
	if err != nil {
		return groups, err
	}
	resp, err := up.request(req)
	if err != nil {
		return groups, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&groups)
	return groups, err
}

func (up *Uptrends) GetGroup(guid string) (Group, error) {
	group := Group{}
	endpoint := fmt.Sprintf("/probegroups/%v", url.QueryEscape(guid))
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return group, err
	}
	resp, err := up.request(req)
	if err != nil {
		return group, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&group)
	return group, err
}

func (up *Uptrends) AddGroup(group *Group) error {
	body, err := json.Marshal(group)
	req, err := http.NewRequest(http.MethodPost, "/probegroups", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	resp, err := up.request(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respGroup := Group{}
	err = json.NewDecoder(resp.Body).Decode(&respGroup)
	if err != nil {
		return err
	}
	*group = respGroup
	return nil
}

func (up *Uptrends) EditGroup(group *Group) error {
	if group.Guid == "" {
		return fmt.Errorf("unset guid, add the group instead")
	}

	body, err := json.Marshal(group)
	endpoint := fmt.Sprintf("/probegroups/%v", url.QueryEscape(group.Guid))
	req, err := http.NewRequest(http.MethodPut, endpoint, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	resp, err := up.request(req)
	defer resp.Body.Close()
	return err
}

func (up *Uptrends) DeleteGroup(guid string) error {
	endpoint := fmt.Sprintf("/probegroups/%v", url.QueryEscape(guid))
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
