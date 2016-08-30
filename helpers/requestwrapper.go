package helpers

import (
	"errors"
	"github.com/kataras/iris"
)

type RequestWrapper struct {
	Context       *iris.Context
	postArgsCache map[string][]string
}

func (self *RequestWrapper) PostValuesAll() map[string][]string {
	if self.postArgsCache == nil {
		self.postArgsCache = self.Context.PostValuesAll()
	}

	return self.postArgsCache
}

// PostValues returns the post data values as []string of a single key/name
func (self *RequestWrapper) PostValues(name string) []string {
	values := make([]string, 0)
	if v := self.PostValuesAll(); v != nil && len(v) > 0 {
		values = v[name]
	}
	return values
}

// PostValue returns the post data value of a single key/name
// returns an empty string if nothing found
func (self *RequestWrapper) PostValue(name string) string {
	if v := self.PostValues(name); len(v) > 0 {
		return v[0]
	}
	return ""
}

func NewRequestWrapper(ctx *iris.Context) *RequestWrapper {
	return &RequestWrapper{Context: ctx}
}

func (self *RequestWrapper) Post(name string) (value string, err error) {
	value = self.PostValue(name)
	if value == "" {
		err = errors.New("Field is not provided")
	}

	return
}
