// Code generated by protoapi:go; DO NOT EDIT.

package apisvr

// Env
type Env struct {
	Env_id   int    `json:"env_id"`
	Env_name string `json:"env_name"`
}

func (r *Env) GetEnv_id() int {
	if r == nil {
		var zeroVal int
		return zeroVal
	}
	return r.Env_id
}

func (r *Env) GetEnv_name() string {
	if r == nil {
		var zeroVal string
		return zeroVal
	}
	return r.Env_name
}