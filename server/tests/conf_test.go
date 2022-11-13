package tests


import (
	"encoding/json"
	"io"
	"os"
	"testing"
)

func TestConfLoad(t *testing.T) {
	f, err := os.Open("../conf/default.json")
	if err != nil {
		t.Fatal(err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	var conf Conf
	err = json.Unmarshal(data, &conf)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.Server.Host)
}
