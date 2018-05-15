package gohttp

import (
	"testing"
)

func TestNewRequest(t *testing.T) {
	opt := &Options{
		Json: true,
		Body: map[string]string{},
	}
	r := NewRequest("get", "http://www.baidu.com", opt)
	if _, err := r.Do(); err != nil {
		t.Error(err.Error())
	}
}
