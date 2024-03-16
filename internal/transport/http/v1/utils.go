package v1

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func parseQuery(urls url.Values, reqForm interface{}) error {

	data, err := json.Marshal(urls)
	if err != nil {
		return err
	}
	fmt.Println(urls)
	if err = json.Unmarshal(data, reqForm); err != nil {
		return err
	}
	return nil
}
