package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

type Run struct {
	RunSave
	RunId string
}

func NewRun(fileContent []byte) (Run, error) {
	var save RunSave
	err := json.Unmarshal(fileContent, &save)
	sum := md5.Sum(fileContent)
	r := Run{RunId: hex.EncodeToString(sum[:]), RunSave: save}
	if err != nil {
		return r, err
	}

	return r, nil
}
