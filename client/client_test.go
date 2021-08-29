package client

import (
	"testing"
)

func TestClientFromJSON(t *testing.T) {
	client, err := ClientFromJSON("test-credentials.json")
	if err != nil {
		t.Error(err)
	}
	t.Log(client)
}
