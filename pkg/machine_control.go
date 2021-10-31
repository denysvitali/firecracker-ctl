package firecrackerctl

func (c *Client) Start() error {
	return c.action(ActionInstanceStart)
}

func (c *Client) SendCtrlAltDel() error {
	return c.action(ActionSendCtrlAltDel)
}

func (c *Client) FlushMetrics() error {
	return c.action(ActionFlushMetrics)
}