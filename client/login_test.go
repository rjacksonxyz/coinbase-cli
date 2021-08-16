package login

import (
	"testing"
)

func TestClientFromJSON(t *testing.T) {
	client := ClientFromJSON("test-credentials.json")
	data := client.Get()
	t.Log(data)
}
