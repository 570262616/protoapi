// Code generated by protoapi:go; DO NOT EDIT.

package echosvr

// Msg
type Msg struct {
	Msg string `json:"msg"`
}

func (r *Msg) GetMsg() string {
	if r == nil {
		var zeroVal string
		return zeroVal
	}
	return r.Msg
}
