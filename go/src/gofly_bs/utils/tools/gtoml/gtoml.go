package gtoml

import (
	"bytes"

	"github.com/BurntSushi/toml"

	"gofly/utils/tools/gerror"
	"gofly/utils/tools/json"
)

func Encode(v interface{}) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	if err := toml.NewEncoder(buffer).Encode(v); err != nil {
		err = gerror.Wrap(err, `toml.Encoder.Encode failed`)
		return nil, err
	}
	return buffer.Bytes(), nil
}

func Decode(v []byte) (interface{}, error) {
	var result interface{}
	if err := toml.Unmarshal(v, &result); err != nil {
		err = gerror.Wrap(err, `toml.Unmarshal failed`)
		return nil, err
	}
	return result, nil
}

func DecodeTo(v []byte, result interface{}) (err error) {
	err = toml.Unmarshal(v, result)
	if err != nil {
		err = gerror.Wrap(err, `toml.Unmarshal failed`)
	}
	return err
}

func ToJson(v []byte) ([]byte, error) {
	if r, err := Decode(v); err != nil {
		return nil, err
	} else {
		return json.Marshal(r)
	}
}
