package awair_api

import (
	"fmt"
)

func (c *Client) DeviceAPIUsage(deviceType string, deviceId int) (*DeviceUsage, error) {
	req, err := c.newGetRequest("v1", fmt.Sprintf("users/self/devices/%s/%d/api-usages", deviceType, deviceId))
	if err != nil {
		return nil, err
	}

	data := new(DeviceUsage)
	if err := c.do(req, data); err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) Devices() (*DeviceList, error) {
	req, err := c.newGetRequest("v1", "users/self/devices")
	if err != nil {
		return nil, err
	}

	data := new(DeviceList)
	if err := c.do(req, data); err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) UserInfo() (*User, error) {
	req, err := c.newGetRequest("v1", "users/self")
	if err != nil {
		return nil, err
	}

	data := new(User)
	if err := c.do(req, data); err != nil {
		return nil, err
	}

	return data, nil
}
