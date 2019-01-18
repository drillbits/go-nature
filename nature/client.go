package nature

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Client represents a client for Nature Remo Device HTTP API.
type Client struct {
	c    *http.Client
	addr string
}

// Local is a client for Local API
type Local struct {
	*Client
}

// NewLocal creates a new Local client.
func NewLocal(addr string) *Local {
	return &Local{
		Client: &Client{
			c:    http.DefaultClient,
			addr: addr,
		},
	}
}

// FetchNewestSignal fetches the newest received IR signal.
func (l *Local) FetchNewestSignal() (*IRSignal, error) {
	u := fmt.Sprintf("http://%s/messages", l.addr)

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Requested-With", "curl")

	resp, err := l.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d %s: something went wrong", resp.StatusCode, resp.Status)
	}

	var sig IRSignal
	err = json.NewDecoder(resp.Body).Decode(&sig)
	if err != nil {
		// nodata => {"format":"us","freq":38,"data":[
		if err == io.ErrUnexpectedEOF {
			return nil, nil
		}
		return nil, err
	}

	return &sig, nil
}

// EmitSignal emits IR signals provided by request body.
func (l *Local) EmitSignal(sig *IRSignal) error {
	u := fmt.Sprintf("http://%s/messages", l.addr)
	b, err := json.Marshal(sig)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, u, bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	req.Header.Set("X-Requested-With", "curl")
	req.Header.Set("Content-Type", "application/json")

	resp, err := l.c.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%d %s: something went wrong", resp.StatusCode, resp.Status)
	}

	return nil
}
