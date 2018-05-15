package gohttp

import (
	"net/http"
	"io"
	"io/ioutil"
	"encoding/json"
)

type Response struct {
	Ok         bool
	Body       io.ReadCloser
	StatusCode int
}

func WrapResponse(req *Request, resp *http.Response) (*Response) {
	res := new(Response)
	if resp.StatusCode == 200 {
		res.Ok = true
	}
	res.StatusCode = resp.StatusCode
	res.Body = resp.Body
	return res
}

func (r *Response) ParseJSONBody() map[string]interface{} {
	if bs, err := ioutil.ReadAll(r.Body); err != nil {
		return nil
	} else {
		ret := make(map[string]interface{})
		if err = json.Unmarshal(bs, &ret); err != nil {
			println(err.Error())
			return nil
		}
		return ret
	}
}

func (r *Response) ReadBody() ([]byte, error) {
	return ioutil.ReadAll(r.Body)
}

func (r *Response) Close() {
	r.Body.Close()
}
