package awair_api

import (
	"fmt"
)

func (c *Client) UserLatestAirData(deviceType string, deviceId int) (*DeviceDataList, error) {
	endpoint := fmt.Sprintf("users/self/devices/%s/%d/air-data/latest", deviceType, deviceId)
	req, err := c.newGetRequest("v1", endpoint)
	if err != nil {
		return nil, err
	}

	if c.UseFarenheit {
		c.appendQueryParam(req, "farenheit", "true")
	}

	data := new(DeviceDataList)
	if err := c.do(req, data); err != nil {
		return nil, err
	}

	return data, nil
}
