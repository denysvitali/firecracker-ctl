package firecrackerctl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Action string

const (
	ActionInstanceStart  Action = "InstanceStart"
	ActionFlushMetrics   Action = "FlushMetrics"
	ActionSendCtrlAltDel Action = "SendCtrlAltDel"
)

type ActionRequest struct {
	ActionType Action `json:"action_type"`
}

func (c *Client) action(act Action) error {
	aReq := ActionRequest{
		ActionType: act,
	}
	aBytes, err := json.Marshal(aReq)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(
		http.MethodPut,
		c.url("/actions"),
		bytes.NewReader(aBytes),
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