package server

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func readAddress(r *http.Request) (string, error) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024))
	defer r.Body.Close()
	if err != nil {
		return "", err
	}
	dec := json.NewDecoder(bytes.NewReader(body))
	defer r.Body.Close()
	var claimReq ClaimRequest
	if err := dec.Decode(&claimReq); err != nil {
		return "", err
	}
	r.Body = ioutil.NopCloser(bytes.NewReader(body))
	return claimReq.Address, nil
}
