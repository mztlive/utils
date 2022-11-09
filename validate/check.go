package validate

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-playground/validator"
)

func Check[T any](bodyByte io.ReadCloser) (T, error) {
	var body T
	if err := readBody(bodyByte, &body); err != nil {
		return body, err
	}

	if err := validateRequest(&body); err != nil {
		return body, err
	}

	return body, nil
}

// ReadBody 读取Post请求的Body到结构体
func readBody(body io.ReadCloser, structPtr interface{}) error {
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return errors.New("body decode failed")
	}

	err = json.Unmarshal(bodyBytes, &structPtr)
	if err != nil {
		return fmt.Errorf("request body json transfer fail. reason: %w", err)
	}

	return nil
}

// ValidateRequest 对结构体进行规则验证
func validateRequest(structPtr interface{}) error {
	v := validator.New()
	return v.Struct(structPtr)
}
