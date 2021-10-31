package firecrackerctl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type DrivesRequest struct {
	DriveId      string `json:"drive_id"`
	IsReadOnly   bool   `json:"is_read_only"`
	IsRootDevice bool   `json:"is_root_device"`
	PathOnHost   string `json:"path_on_host"`
	PartUuid     string `json:"part_uuid,omitempty"`
	RateLimiter  string `json:"rate_limiter,omitempty"`
}

func (c *Client) Drives(dr DrivesRequest) error {
	reqBody, err := json.Marshal(dr)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(
		http.MethodPut,
		c.url(
			fmt.Sprintf("/drives/%s", url.PathEscape(dr.DriveId)),
		),
		bytes.NewReader(reqBody),
	)
	if err != nil {
		return err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusNoContent {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("unexpected %s, additionally: %v", res.Status, err)
		}
		return fmt.Errorf("unexpected %s: %s", res.Status, bodyBytes)
	}
	return nil
}
