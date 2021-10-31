package firecrackerctl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BootSourceRequest struct {
	KernelImagePath string `json:"kernel_image_path"`
	InitrdPath      string `json:"initrd_path,omitempty"`
	BootArgs        string `json:"boot_args,omitempty"`
}

func (c *Client) BootSource(bsr BootSourceRequest) error {
	reqBody, err := json.Marshal(bsr)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPut, c.url("/boot-source"), bytes.NewReader(reqBody))
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