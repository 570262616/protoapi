// Code generated by "esc -o generator/data/tpl/tpl.go -modtime 0 -pkg=tpl generator/template"; DO NOT EDIT.

package tpl

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/generator/template/echo_enum.gogo": {
		local:   "generator/template/echo_enum.gogo",
		size:    383,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/3SQMWvDMBCFZ9+veIQOFjTOnpKpaaFLUmjoEjKo9tWYxmchy0MQ99+L7MSZsj2O70kf
b7XCa1cxahb2NnCFnwuc70JnXfOC7R67/QFv249DQeRs+WdrRozF5xRVicLFjaedbVkVjQSispM+IKcs
xiW8lZpRvDd8rnqopuuNjvHpFjfpkW97HnhClmCpVMkQ/Q5SIi+T6Fw1+Aq+kTo36MeASJnYlnusN2it
O87oaQIiZQ987kJrLOa8eL4WJo9MiTLPYfCC8Z9jEjqRPvJLw+YmDZLUrs28kWBG1JDSfwAAAP//H21g
KH8BAAA=
`,
	},

	"/generator/template/echo_service.gogo": {
		local:   "generator/template/echo_service.gogo",
		size:    2577,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/7xWbW/bNhD+LP2Kq5d1pBPLirciSwIj6xK764AkxWLsw8I0o6WzTVQmDYrOnEnsbx9O
cgw7TfcWo18E+l6ee+7hHeF2G05NijBGjVY6TGF4DzNrnJEzdQxnl3BxOYDe2dtBFIYzmXyQY4SiiN7V
R+/DUE1nxjpgYdCwOMbFrBGGQWOs3GQ+jBIzbWdymDuZfGhjMjGNkIdhYnReZfSmUmWQO6v0GLrQeM8Y
Y9ey9efr1m83pRBpef3iKyF2vn75jRBNIXaFaAnR7gpxIsT7298LIUr/8aa8FmJRxPHr2LfodHbQ73s6
9A/jpal/dro0nfUfTP1e39/wXSZE9MWL8iYvGRNi0elwappOcSnEIj7kTfoVp/SRnJ9s+Hb5CWNUOd4n
vPh7+gzpk9AHybg/EmJxMKJeFp39imfnW3K8GtasX6X06wCf10PJhFincrhJpa4xel4NznnzP6pTa8r5
D5+M0rN63R7WJlKrFCIqb8uPz8Jsbo8e50JEfHet4W0Jt03RtijYs8U62aletTtp6U2zi/pV60L9HEbn
89ydmulMZcgqF6fodpve0Qs5Re9B5eAmCEo7tCOZICRGO6l0DjLLKhcZrMkytHno7me4nrzKKsKgKFpg
pR4jROfoJibNwXsyRwPlMvSe0SscnRrtcOH2oFkU0Vs9m7vB/Qy958DIcjl3K1NRqBFohKhnrbFkg0bD
+zp1ZaM41Kn3vOaAOqXCPgyfZjSa6wRuV03c/iR1mqFlub2DothZmjlUbJfOPuUUYWDRza0GgmAJrPfD
gaG1gMSKU2igNBx1QeMf7FGjITlHFApdSKIflU6Z0vy4srzoglZZBRBYzGeEcWqmU6OrhguKrk5H8HJ1
Ls4xz+UYjwiiVoZxEj94YJxEP19dXrDvOvEeECwPg8AvidzJrGctFVI6+lVmKpUOGT9+cPwTpYeUJa06
618VN3P35B0DXXIl5fJqqWRu76K1YUr2QGlColt+GmOls3ncxAatOF4Wq3kFD1PkK46bwR0KNnPHw4BG
bG3eaK1+wbHKHdqN9ZrnmIIzMFQ6BWvmjhapGsJPwhlCsxqqXjIxe1BP5Gog/27HVhJcob1TCdYivLu8
GtRCYPSmN2CN6v+Tm3jf2PvMCvDNLfo89pveCprK/A/sR9v6VwAAAP//A2OiBBEKAAA=
`,
	},

	"/generator/template/echo_struct.gogo": {
		local:   "generator/template/echo_struct.gogo",
		size:    734,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/7SRQW+bQBCFz95fMeVgmShZ90xFpShgCclxWtfNpaqSLQx0W7yQYZFqrfa/V7Muim2p
x1zQ4+3sex/Dcgl3XYXQoEFSFiv4cYCeOtupXn+A7AE2DzvIs2InhehV+Vs1CM7JT0fpvRDLJRt3rRqG
jdqzZQ89XngwWBpLC07MnLsBUqZBkCuNbTWA9+zKnbYtj7I89Kyefw2dSSLn5DElej5eR1PxJS9EPZoS
FnTRFsOjanWlLC5iuJp0TtQREyDRAEkK375fBYJw4Pz/0W5A1yCnnC2+jJowIMx0DSRP4NMUoohLZsgV
qyJfZ0/b/PPXYptn7HJ3Cqrv0VQLfruG+QlFkPwVyVnuNYRj3ksCc/SxmPmzXUyU+PIKuuporyxEuFe6
jSbcd/QnZ0PeK1v+/GJJm2Zx2hW/4hebx9t1kT3l97fF+q3pJ61raNGE9Bg+wvuAQ2hHMjA/+5kuPIcE
eNaH0H9zRrfCi78BAAD//ww5MNzeAgAA
`,
	},

	"/generator/template/php_client.gophp": {
		local:   "generator/template/php_client.gophp",
		size:    3889,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/5RWW2/bNhR+1684MwzYDhwl3R5WxHOGNPG2AklTdM6AoSkMWj6OuUmkIlJNHIH/fSAl
0bpQtqeHICbP5Tsfz+2XX+NN7J2dwXxDBVABBNY0RHhChgmRuILlFuKES05iCsPvmAjKmZ+mb6kf8Ois
vBppG79bJSIvIMvAn9MIQSl9eXMPn+7nMLv5OPc9j5EIRUwChCzzP5EI/9Q/lJp4XioQ/ub8jfPHz9r6
VUwn5vBue31LluJxxtLI/Jl43tnJCdyhEOQJBZycnHlZdgoJYU8IfnmulBeERAjIMklliGAcKmVk6Rqo
+EDfZklSngO+SmQrAaX7x/yeJ7PXAGNJOTOqGArM9a95tFf/mkcRZy4TbKUU0CgOMUImKyoFeC/zAAAq
Uf1GMVwJTaq+0PxjoCnvF0RqDksVY97LBdNlSANYpyzQ3oEyKockScgW+gmKmDOBo1zR/N3rVX90DUMq
BMqh1f/asxh630ajiqXSGl0DPoN/S5YYQu/26sPsdvFl9nl2NZ/d9CrG9deXGypOL61JmILBOxxNanJr
niAJNtCBA4iocNMEVQFGxf3yHwwk+DdEkvk2RmggylFFMUyB4csunUp5pZrYSo3TS8N3BUeX3HcS0hWR
6LZUZ+TrN5gatYkzJp2fzghcVhrZ07JlMql6qrwj/B1NreOxj6S4qVlS7U7KQ7pd9HfE54Dtduyw1mK0
fao8911HTe/A/99C/qGo5EY4rRqWm4S/mJf5XDT+x7zph7avDXsDqz/QE4VxCfhKhexVOO2OLD9zN7tG
vALlwvoaZpkz05RqtH3IMuOvUowNulyv2qiPjhd4qiFq2k1QpglrmZ/UIt9x0TQu+aJogG67+eVxj763
OBtylTyG6aWjaHbAxscUzQF74wNF0T4ZlQzqkZ5f6b1ArwitpcAcdm0EdnJrqUOjN+BMSKgmSZb5f5Ew
dYzgHbDSr91BmmN8I2V8HVJkcuLOscXCeE7SQA77SyLwIaEwhcG7H3/2z/1z/93F+/P354OOpN6ZL1qs
3Tj+sDfD2gs00qr8Btr1Ik3owDxiAWTclpM0Qp5KI/bTeV1g1H7Dkrh+oJemCPXSBBdT8PMlq3gE1WgU
dyg33B43Kas2Cf8ji1OZp7lu1s8tojaErUJMYLozsBsnY+gv6ZtBNS4wsmaXNB0Vo1hud3oj1+ahb+2k
8+9TaZE5x3+CojndOsX27RFlH0pQ1C+V3Wkt/jJYN/5lsRjbIChb4Sv492YUCOgZ3V5HPKV2GZT11ZbN
B4/VOAi7eBj3upfnTC3BTh2LUmXDzc3t329tWO0l5sCGWNcq2XD5dOsW5OxZ4QqGiiaU7bFSawnt2f7A
/mX8hUEODnLy5DbGC+j5+7Jy3/J4lOdPHFZEkiJ5ceXX1omiV7an7K7fnV4GJAyvYqqr53lsxpDueXdy
pVQv//3w5aP5v+wBtisVzCnvvwAAAP//u9r3ZDEPAAA=
`,
	},

	"/generator/template/spring_service.gojava": {
		local:   "generator/template/spring_service.gojava",
		size:    856,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/6yRTY7iMBCF9zlFKauwGHOAbBADQiwgaMgFKkkRLKDssSuNUOS7t8JfsuhuqRFZRX5+
z++rGo/hr6kIamJyKFRBcQHrjBi0OoVZBussh/lsmasoslgesCZoW7W5/YaQRpE+WeMEjKuVt05zvXN4
orNxB3WmQhWaK4XMRlC0YbUgWaG1muv0t9aN8S97/5G3hj1NTXV5wfy/IS83b2Sb4qhLwMKLw1KgPKL3
3VTWeKIQpugJ2ggAoG3/gEOuCdSKZG8qDyE8Fb0DJlBbch+6pPxiCeLFPI8fdyYD3iS+Tl32IcSjmzok
up7ce7WtyhqxjXSJIfTFurhkMmDppCX3NzWP7sW7z5E0jnt7onmUXtUegbjq2v6ItMm2A6Z+/e9AWpAk
b4V4LvXbJ794L71HPZJC9BkAAP//3o5aN1gDAAA=
`,
	},

	"/generator/template/spring_struct.gojava": {
		local:   "generator/template/spring_struct.gojava",
		size:    585,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/5SPwWoCMRCG73mKOdqD8QGWQkFbqBT1sC8w7o7baDYJk1mphLx7ibqtUmhpTjOB+f7v
n81g7luCjhwxCrWwPUFgLx6DqWCxhtW6hufFa62VCtgcsCNISW8uY86VUqYPngUa3+sdRiH+6K3eY3OI
3ml0zguK8U4vo3dzJhTP1b+ONuwDsZy+s/Z4RD2IsfrNRKmUCsPWmgYaizEWv3kZVthTzpAUAEBKU2B0
HYF+MWTbmPP5P7A5ohDsjENbTpd4xPoUzpcJdIFA6TlCyLWQszrvTzedLrSLx73BpKzeCQ+NbJCxz/nh
D6vy5N1EfaMAj78KlfWqldIdE6Zj1y+7m5IdScHWRmzhTka18phkYPczNl9jxuysPgMAAP//ioLHmkkC
AAA=
`,
	},

	"/generator/template/ts/data.gots": {
		local:   "generator/template/ts/data.gots",
		size:    584,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/5yOQWsaQRSA7/MrHl6Epeq90IPFLRSkSJFeSinj7lMH1tntzGypDAOFGBIhBiGag4fk
FBAE9RRIzM9xd/0ZYdVFc8jFOc3wvvneV7IsYkG9zSQ0mYfAJLSQo6AKXWh0IR8IX/k0YPk3mONzRRmX
QD0PVBvBpYqCVCJ0VCgQGsh4C0KJLjC+BQ5WJUGi+MsclMSCwimHWLBZPMS3F+vVYzK6jy+H6+frrJRY
sJtEV+fRcFaufY36g2S22CzPktE07v+PXsbJaJpMevF4GQ/myeomvutF88n6qU+sEtG6AILyFkLR5mFH
gjEE/wW+UIA87IDWxW+0g8aAJgAAWmf4F4aeu+VhN8jAT+n9B/VCNOZD9gm5m6Jmu3D/OF5eoYrWuwEe
BzCuUDSpgydUfEwZJVPlQW6M1qwJ+AeKVdpAD3LV8me7+vu7XbPLdbuSM+bnL62Ru3vfO+GvAQAA//9s
n4weSAIAAA==
`,
	},

	"/generator/template/ts/helper.gots": {
		local:   "generator/template/ts/helper.gots",
		size:    4383,
		modtime: 0,
		compressed: `
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
`,
	},

	"/generator/template/ts/ts_service.gots": {
		local:   "generator/template/ts/ts_service.gots",
		size:    2271,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/9xTXW/cRBe+9684WlXaD23t5pXeqmwIUgQFcgGN1PADJvbZ3SneGXdmnDYyI7WUtFRK
m0iNKiCBFkFRRT8WJChJ2tA/Y+9urvoX0Hjs3Q2iH+KSufLMec45z3nOY6/RcBqw1KUS2jREoBI6yFAQ
hQEsr0I1ElxxEtFqDkOL8jlThDIJbcGZQhbA/OIC+DxAUF2i4AIXn8IFqrqgugghXRZErEKVXKRcVpvm
UWCbC2wCVVUJAs/HVGBgkwuYoUKZVCQMMQDK8lKR4OfQVwWXCdO8NZVwQVClkBn40mqEZ31BoxKdYyLB
V2iAEggsE0l9iCXpILS5sCOQMATCAuiRVWCIAZDgXCxVD5kC4vtcBJR1QHGQEfq0Tf2SUjmERbIAJFUx
UZQzpwHH3/w4DRj17w1uX0ufPhlu3Rl8uZnu3yyX4DTARrL1tWzzQXb9xvBBf/TLleHW/fnFheE3X6RP
fxjevfzi2Xq29yQ9eD7cuj98+DDdvT64vZft38qFffFsHbLtu4NHPx7uXBr9dDl9/u2ofzkPZY++ynbu
p/s3D7/fG24/TncfTRpeXbO1bdUj4o769yzTArrx82BjM/1z29KbX1ywDLM7+4OdhxOGj787/HptcGUt
u/p7ttEfXTmwfAZ39wY3Hmdrf6QHtywP+21Cv36e7t7INtdHl9bTg53s2v5g+7fB1p7T8Bzai7hQkM/R
hATmzcei4D0qEbQxaq901mwJThwAgCQRhHUQjqnVCJtwbJnzEFpzUOugWsiB7xFFzMQS3Pdj5pulyrrW
RfZxmwlaN6F4QhZo7ZRdXS8giky1HRv3ExE2AYXg4kPCgtBYa5LUxTBCUZ11HK+xTKQBQ7bRtyM3PMfn
TCooI3NQ6SoVtTxv5q3/uTMnT7kzM/93T55onTpx6kRl1vE8iCVafRwnZ+2HRMqPSQ/NtO6749txrS3C
6jKZWWub2EPV5YHJqkRcqkqeYSK0DXg+F+4sihXq40cqAPdMZBWDygenlybg6TIdHFex2uWAXBkTz0ue
NjezhqmKWjt4MRe1XZCEJHHNFFrXIiJIT7bMywKLYmVyta63oLDF20ninonVOAKfAcMVFO8UxghRQSzC
FkglzGrmpvdWK4RvQiVJJlJqXclfCg6V+qyT1xKoYsGs/G6SFLNrXYtNiSShbWA4lsTqobWdIEkwlKh1
AvYOulCpnpc2x1VdZDWBEuZK9uXxPOgacyHI2PdRSjBuPAIpyBWyuAIlD1fQlHMNFoiEv0k16azrrk+U
362hEK9qbncpUEacSfyn9kf+A1POLdHT3WadiUmcJPEapbPbqPzuK23737BTIWuAbRKHChbPnF2adlgu
Qy2B3FfWTi2omt+02sw33xq7qP4S17zGMf/WLS93ymtd8oYOOeqOhqf1XwEAAP//PUOqht8IAAA=
`,
	},

	"/generator/template/ts/ts_service.govue": {
		local:   "generator/template/ts/ts_service.govue",
		size:    2033,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/4xUbW8bRRD+fr9iZFXyi+xzxKfKpRWlFIgENKIv3zd3Y3vb8+51dy+tdZzUENJSKWki
NaqABBoERRHNC0hQkrShf+bOdj71L6C99Z3PIBD36Wb3mZlnnpnZZq1m1eBal0poUw+BSuggQ0EUujDf
h7IvuOLEp+UUhgblcKYIZRLagjOFzIWLc7PgcBdBdYmCO1zcgjtUdUF1ETw6L4joQ/lGgJ+i5IFwsFzX
VwLbXGAdqCpLEHg7oAJdE2IKrGlRJhXxPHSBsjSsL/hNdNSY14R1SoNKuCOoUsg0/Frfx6uOoH6GTjG+
4AvURQkE5omkDgSSdBDaXJhyiOcBYS70SB8YogvEvRlI1UOmgDgOFy5lHVAcpI8ObVMno5SVYpDMBUlV
QBTlzKpB4/9/Vg1GB88GTx7EL18MN54OvlyPjx9lDbFqYG6SleVk/XnycHX4/GD0y9JwY+fi3Ozwmy/i
lz8MtxffvFpJjl7EJ6+HGzvD3d348OHgyVFy/Lgg75tXK5Bsbg/2fjzdujf6aTF+/e3oYLEASPa+SrZ2
4uNHp98fDTf348O9SfL7yyaPyTAl9OjgmWE9hq79PFhbj//cNFQvzs0atsnT48HW7oTt/nenXy8PlpaT
+78nawejpRPDarB9NFjdT5b/iE8eGx7mX1/9+nl8uJqsr4zurcQnW8mD48Hmb4ONI6vWtGjP50JBCJd4
z+cMmarDjQAh0rPbg/JCgA1fcB+F6jdcdLggiovyucyxIEPBQ2STmeNCCwAgDAVhHYQzqu9jHc7Mc+5B
6zxUOqhmU+B7RBGtkgT7/YA5eihkNYrG3g3jCVFUh/ERMjeKrIyu3XSJIoW0+eBfF14dUAguPiTM9fRo
Tpy66Pmoq7JuBGgHEiuFuqrnLOudXB4L76aRXWyTwFPgeERKCEP7kv75hPQwigDv6q2XqZKm8mZTvwpS
icBRXKRHBbtSHcP0JwMfRaWa2pGVeb9LJMJ14aW2L+gCUaj3UhfWAqmErug8lLpK+a1m0+MO8bpcqtbZ
mbMzpYl+TkZT6z4hDY2CyKZJkwYU9e+h6nJXO5d8LlVpypG2AW+n3byKYoE6+LFywb7imzZC6YPL16Yd
iuE6OBXNNDYHpq3TuDT8ZW3pOSlEz9G26ULFJ4L0ZEufzDI/UBofRdUWzAneoxI/orfw7TC0rwQqv4XP
II19AUIr74iHCoIpmQtTVVFdKu1xJ+qlMJxIHEWlOpRyQiU9SVlMgSoQ+qWm0j6jm2aH4ViNKKoEeloz
SRnmOhmRosiUFoboSYyi0JjRWLSqrbrIKnkqk07C+QuFKcu+ZhO6eiEQZOA4KCXoDYIGvDUz8w/wmPRY
P1uvubeAFYHSTr1IuglFPatTMaL6lIlCpKTgP1iZvguUPmcS/43R1GJXUAg78/gbgdyqnrMmB/kzYv0V
AAD//yf6ngLxBwAA
`,
	},

	"/": {
		isDir: true,
		local: "",
	},

	"/generator": {
		isDir: true,
		local: "generator",
	},

	"/generator/template": {
		isDir: true,
		local: "generator/template",
	},

	"/generator/template/ts": {
		isDir: true,
		local: "generator/template/ts",
	},
}
