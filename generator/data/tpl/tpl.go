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
		size:    2781,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/7xWb2/bthN+LX2Kq3/9taRjy4q3IksCI8tiu8uA/MFi7MXCNJOls01EJg2KSp3K7Gcf
TrI9O033L8HeCNTd8bnnHh4PbLXgRCcIY1RoIosJDB9gZrTV0UweQvcCzi8G0OueDgLfn0XxXTRGKIrg
slo65/tyOtPGAvO9msExzmc13/dqY2kn+TCI9bSVRsPMRvFdC+OJrm37HrT+pHVrlXG9GOuaz30/1ior
oXvTSKaQWSPVGDpQ+8AYY9dR89Nx89ebhRDJ4vrV/4R4/f83b4WoC7EjRFOIVkeIIyE+3P5WCLFwn28W
10LMizA8Dl2TVt29ft/Ror8fLk397snS1O2vTP1e393wHSZE8J8n5XW+YEyIebvNqWhahQsh5uE+r9Nf
mNAn4vxoy7fDjxijzOEu4YXf0WdIn5g+SMbdkRDzvRHVMm/vljzb35Dj3bBi/S6hvz18Xg0LJsQmlf1t
KlWO0fNycM7r/1CdSlPOv/+ilZ5V68thbSM1F0IEi9vF52dh1l+OHudCBHxno+CXEu4lRXtBwZ4t1tHr
cqrdR4ZmmplXU60D1dwMzvLMnujpTKbIShen6FaLBu55NEXnQGZgJwhSWTSjKEaItbKRVBlEaVq6yGB0
mqLJfPsww83N612F7xVFE0ykxgjBGdqJTjJwjszBQNoUnWM0roMTrSzObQPqRRGcqlluBw8zdI4DI8tF
btemopAjUAhBzxhtyAa1mnPV1rWN4lAlzvGKA6qEEjvff5rRKFcx3K6LuP0xUkmKhmXmHori9dLMoWS7
dPZpT+F7Bm1uFBAEi2GzHg4MjQEkVpxCPangoAMKP7JHhfrkHFEodCAOfpAqYVLxw9LyqgNKpiWAZzCb
EcaJnk61KgsuKLpcHcCb9bo4wyyLxnhAEJUyjJP43opxHPx0dXHOvm2HDSBY7nueWxK5j9KeMZRIquCX
KJVJZJHxw5XjryittixpVbv+VnKd2yfPGOiQSymXR0spM3MfbDRT3ACpCIlO+WmMtc76cRFbtMJwmazi
5a26yJUct4PbFKxzy32PWmyj3+ha/YxjmVk0W9crzzABq2EoVQJG55YuUtmEX4QzhHrZVL14ohtQdeS6
IQvfa7Ug+yhtPCFAerzEFojX8eUpNQOaBkmRZ/SoIaC3GXRxFOWprdw+CXLbAH1HimJQWQNWZd0K5YcU
RXqtwqDq5j9eVMFW6lKSPxkD61O6QnMvY6zO6fLialCdFQbvewNWK9+CduJcrfGVW8q3L/rXsd/31tCU
5l9gPxoovwcAAP//F5mGs90KAAA=
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

	"/generator/template/go/enum.gogo": {
		local:   "generator/template/go/enum.gogo",
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

	"/generator/template/go/service.gogo": {
		local:   "generator/template/go/service.gogo",
		size:    1979,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/5xVTW/jNhA9i79iagStFHipYNFeHPjQJu5uCjQOGqPXQJbGMhGZVEgq2yzB/14MGcmW
10mLPRgQZh7n472ZcZ7DlaoQapSoC4sVrF+g1cqqohWzWl3C9RJulytYXN+sOGNtUT4WNYJz/C5+es+c
OzMwmwP3noldq7SFlCWTWthtt+al2uVNsTa2KB9zLLdqMva9KPVVqbxPOnzUasIyxvKckt0WO/QehAG7
RRDSot4UJUKppC2ENFA0TXCRQaumQW2YfWnx8PHwyrHEuQ+gC1kj8D/RblVlwHtGdr4StkHv0xKoWn6l
pMV/7BQ0PsG5c/xGtp1dvbTofQapRtMG87Kzg905sQGJwBdaK002mEy8n8JafA2m8GJw0gOUFQFQa/op
nbEktt6XMy7m20L+dw2nU2eRE5QVEeEZO83QppMlPAykPnwuZNWgTo1+BufOXs1ZpO7V+Tu9cSzRaDst
gUIckZtBum+coAmRPZuDxC/pUaeMJYnYBKLmUPLfhKxSjU/ZZTD9MAcpmhAiCQycGf65MFdqt1OSsKFz
78kfpJvNqW7DI+LV6wbkDH4cvh1qHZlLsz5C6Kjkf9wvb9OfP17QlJg2i9mxMXgK98vFRRC6hxH9LEk8
y8/Ziar/LhpRFRb3lYsNPBfNQmuqXuMT7yFpdtl7Dol4p9FR8Nnr4//qzbN93ec5iwlOzhscDv1oyGdz
MPqZH+5bWLFsL++7Wh7oSHCKR/LQGhw3SaMxCnaqNwxq+LEke+C91ULWg3T9GEQ2Ah0f4HT/sZ9h8ce6
jMqg2D1uT3PYyXCbjvAfL/aS0MYerC9dzb+wFsaiHl3PzmAFVsFayAq06izdybDT38BThPOwo4tyq6YQ
F3zYbxfOk/kibLmlgMZqUVqgwn69u6GVQT0lRjojZB2W/ScD17gpusZGNyNeHqagHoN4PFp5GrOOoNkl
oYizHgbxNuz/LPgodaDkrSt/KNY96mdRYpTrbnm/ipIh/7RYpZPwN2e33k+mbxy9bHw33479aTGEpjTf
EfvoPv8bAAD//xZazB67BwAA
`,
	},

	"/generator/template/go/struct.gogo": {
		local:   "generator/template/go/struct.gogo",
		size:    770,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/7SSQW+bQBCFz95fMeVgmShZ90xFpShgCclxWtfNpaqSLQx0W7yQYZFqrfa/V7Muim21
x17Q483svG8Hlku46yqEBg2SsljBtwP01NlO9TppuneQPcDmYQd5VuykEL0qf6oGwTn54Si9F2K5ZOOu
VcOwUXu27KHHCw8GS2NpwYmZczdAyjQIcqWxrQbwnl2507blVpaHntXzj6EzSeScPE6Jno/H0VR8yAvh
nK5BPqpWV8riFl9GTVh5L+rRlLCgC4wYptZFDFeTzok6YjQkGiBJ4cvXq4AWCs7/m/kG/hYfiroGkie3
SlOIIg6ZIUesinydPW3zj5+LbZ6xy9kpqL5HUy347RrmJxRB8i2Ss7nXEMq8sATm6GMx82dLmijx5RV0
1dFeWYhwr3QbTbhv6FfOhrxXtvz+yZI2zeI0K37FLzaPt+sie8rvb4v1/6aftK6hRROmx/Ae3gYcQjuS
gfnZx3ThOSTAvT4M/dNndCu8cA4N/yW/AwAA//8ZKFV/AgMAAA==
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

	"/generator/template/ts/helper.gots": {
		local:   "generator/template/ts/helper.gots",
		size:    3540,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/4xXT2/bRha/81O8FXZDSVYobzYBAjmM4zhO7IVjeWX7soaxGItP1sTUkDszlK1VBGzR
pmmLps0h7aE9tEBRIEUOyakoEgT9Mnacj1G8GZKibDk1D4L5/vze3/kNXa9WnSpsdrmCDg8RuII9FCiZ
xgB2B+DGMtIRi7lrzNBatSOhGRcKuhjGKKGTiLbmkVCgu0zDQZSEAewiJAoD4KIAyWJuIGqQqISF4YDU
rldnMfe0ciGS9HZ4eLiBss/bSMI08hikHQUm0wPJtUZBEJuDGDfaksfaqcLliz9OFd6//Pndt4+P3vx6
8uyHd589PXr9VVazUwWrOf7y0fHTF0e//f/o9ffHj9++++bV8ZOvj7/4xXqcfPfJ8edPTl68fP/q45Nn
zxfWV8aOnz46evPTyY8fHb39/eTZ82KWdWc45B3w7qHwlplajHq9SCxJGcnRyOG9OJIahtBjcUFD/jCC
jox61Kbh0FsMmVJrrIejUXP3gXKd4RBFMBo5NFiowh3scIEBLGsdwyL1rRNJkKjiSCiELhNByMWeA9W6
g4cmKIqkB12tY2M+dAAA7izdXdha3QQfZmtGsNZs3V9YBR+uzKaS2yv//s9Sq9VsgQ9XM+Fi8/795tpY
fiWVr6xtLrXWFlZzzTXyyLKmxG/FTLIeDLNUR+OkdRcBqR25qJh+tovWZDktsIxSVhqwLqMeV3hDYB/l
zbQ43iGtl+P7vg+JCGznKqkRPborowPCnTOikfk1P+qA63YXJnA8pZlOVBGgzajnaWu9vGON3IAeiTqR
IkvVk/gA23oSOWCaVT60P9MDFqcxGTNEncZdkhL8KWs3JYG5C6Sdg1byzbStcyYaOok854ycs6vQZyEw
UFrafYVbFlsBE4CCOCFItZaF2kwQB7E4RhGAjgz3JDKcviwGoNxnYSMFqWR/pPNL67OGW62VxagXRwKF
JqeK54z7EIesjeX6367O1vd4Ddxb7lT1PxasujFdfeVqfa8G7l/P0S5a59o56lnjPDNde+22dd4+R33H
qnfcSmEScDvhYQAMtlqrdDPYvlJ/zHgUddicTBEUhpbIEBpGvkurSK9EQCRQluELximQtZf43wSVhmiX
FskD9PY847uMYRi1Um1+1ZCLNYWoA5qIsmhYXJgG5JsDnSQMYUuG2awvmxD3ljapxH0c1PssTBBixqUy
GHjIenGIDXqhksjXhxKdsUa9HkZtFnYjpRvXZ6/PlsiIyT3wYShYDxtQOkCxd4C8VAPB2/tWoDkTpRHZ
Znn5N2Eq4Dyh+DnIJcLwc4ipi53dmv9KUA62ZHhj82Y5kfmW19KmN2Dz9MLzDpT/YrVFDkvPQSLDueJR
JgaJmdQKfNjemXOslJpZJtU+mpv+LBwp+ywccwmFbdqZm3uYRul1mWoeiHUZxSj1oLyPg0oRhB6iBz8N
sL2Pg50xZJpihm4sfR8Ezf7hQ7MsUQcysZtzv3s6hq19KrApskYgaqyv1ykgV8CkZIPTSXg62jANL1ds
4O10gRfIeudM9H3wTR9nwN3ecQtZAIYKpxufbpEZT5+F07tDBl4nkkus3S33aQ0nQfN6AqZxQmNKOreg
uzzEs/WYiOAD+a1sNDPXuTNGeVgLN6FPi6f42RxtbGt7btB/bjTXPLvvvDMo90+FHU28mcX24kR1y+lF
sV+hMfguzORXR6Uybmml+IVAq6FQchby/2GwbknOT0EfRFyU3UtEtfmpO21crIEYdMYHOsMeFwEeNjtl
d961Db/8d5gHd96FBhAkzJyJW8xrfJKn3rcU6YLEnUroO7RoAURXBbOMlHK7nKVyw5ReL8yuxcg+lAzl
p/86GP1ESB9KG2xgbEoXodt6Ec7LXD9Is9MYtpDjWFhM7E8+OczIoVQv2Xnm9c5AySNZEYou7D8CAAD/
//HGR/3UDQAA
`,
	},

	"/generator/template/ts/objs.gots": {
		local:   "generator/template/ts/objs.gots",
		size:    1210,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/5xSz2sTTRi+z1/xkEOTLl+Te+GDr5+NVShatHopRSabd5Mxm9l1ZtYahgHBihasFGw9
9KAnoVBoexK0/jlN0j9DZtMkG1s89CWE2Xmf55nn/VELAhZgvS00IhEThEaLJCluqIlGD+VUJSbhqSjP
wMJEGi6kBo9jmDahyQ2HNioLTaYIDRKyhUxTE0LmgKmq0dCkXoqQNAuwcJtgAS5Pvw0+v7s4/z7c/zp4
v3fx8+PYKQswyvQ/vO3vHS+t3e/v7A6PTy/P3gz3jwY7r/u/Dob7R8PD7cHB2WD3ZHj+afBlu39yePFj
hwU1xmo1kMy6mlm7AMVli1Ct+ws4x+hVmiiTA2Bt9QHvknOwDAAK+LuC4mZOGCXGwH/9+SmPM3LunwmJ
ZNNDHbP26lirjVpqeinNGFnmhq/7y4IZIQ2piId0C0eLsBZGe8mpuHPWigj0AtVV3qAYpdWl/+urzx7V
1+pL6/XlknMbm7nXidy1IhacYyyXqa6QrN7j+k7S7SayrlSifIVBwOB/+C/lindhi3mE4498e8hfImk8
p9AwBLVx5VEmQyMSiS5PC3RfQ2Wi4EvMPcwYmF9ERRslZOt69nHWyHvs3PxVI6NEoRKTQYd6fqUn4mOA
DxFh+mq1zfXDLbmmkpSU6VU61JvH3NyUudGh3maR7kNvCRO2MULb4uC9xRUyI5eU25wdaTFCrglla5FP
Gc6VF69hfCgymZJ/eALXfikm2+AfuBrwDU81KeJZbP6qX3oiOzLZQu66NIN0jE2P0/+bicXd+h0AAP//
LkxTtboEAAA=
`,
	},

	"/generator/template/ts/service_axios.gots": {
		local:   "generator/template/ts/service_axios.gots",
		size:    1771,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/4xTb08kxRN+P5+isrlk/2SZhV/yu+AiJqin8kIhAT9AM1O72zjbPdfdA0fGTg6ROy+B
g+TIRQU9jJ4hHn800RO4w/syM7vLq/sKpqdn2UWNsV9Ndz311FNP1dQqFacC8y0qoUEDBCqhiQwFUejD
wgoUQ8EVJyEtZjC0KI8zRSiT0BCcKWQ+TM1Og8d9BNUiCpa5+ASWqWqBaiEEdEEQsQJFcodyWayaR4EN
LrAKVBUlCLwdUYG+Tc5hRgplUpEgQB8oy6hCwRfRU7mWgdKsNJWwLKhSyAx8fiXEOU/QsI/OMKHgS9RH
CQQWiKQeRJI0ERpc2BZIEABhPrTJCjBEH4i/GEnVRqaAeB4XPmVNUBxkiB5tUK8vqd+ERTIfJFURUZQz
pwIj//04FeidPO08vp+8eN7dedL5Yjs5f9gfglMBG0k31tPtZ+mDze6zk97Pa92dg6nZ6e7Xnycvvu/u
r75+uZGePU8uXnV3DrqHh8npg87js/T8UWbs65cbkO7ud45+uNy72/txNXn1Te9kNQulR1+mewfJ+cPL
7866u8fJ6dGg4L11y21Zr5nbO3lqlebQrZ86W9vJH7tW3tTstFWYPjnv7B0OFB5/e/nVemdtPb33W7p1
0lu7sHo6+2edzeN0/ffk4pHVYb9N6JfPktPNdHujd3cjudhL7593dn/t7Jw5lZpD2yEXCrI+qhDDlPmY
FbxNJYI2i9rub9ZEHxw7AABxLAhrItxQKyFW4cYC5wHUJ6HURDWdAd8lipiOJbjvRcwzQ5VlrfPsEZsJ
WlfzF2S+1k6/qFuLY/edgEj5EWmj1jMLi8Marrb4YxFUAYXg4gPC/MDs2YCihUGIojjhOEtEmNU1cJiE
QkupsF6rjb3xP3fs5rg7NvZ/9+ZofXx0fLQw4Th4JyvSyFXDHKq3bW4pEkEdpBKUNcu5EwPaSAQTjnZq
NYgkWlMdJ2vV6zdiLBq0BSNaW4Q1c2CU1jaxjarFfZNVCLlUhSzDRGgD8Hbm9hyKJerhh8oHdya0NkPh
/VvzA/AwTROvWKzjGSBz0MQzylvmZmY3xKj133yJY9cOpxQSQdqybl6mWRgpk6t1uQ75Lr0Zx+5MpK4i
8CkwXELxVu5hgAqGrIXJ4fmWcoerUIjjgZVaF7KXXEOhPOFkXAJVJJi1343jvHetzeyqEMe0AQyvLLF+
aG07iGMMJGodg72Dzl0qZ9TmuKqFrCRQwmRfff/UatAyS4ggI89DKcEnilyD5OJyW1yBkgdLaOhcgwUi
4S9WDSrrsusR5bVKKMS/FbezFChDziT+U/lr/4uhGy5iVjjrOtuSPwMAAP//zf5DResGAAA=
`,
	},

	"/generator/template/ts/service_fetch.gots": {
		local:   "generator/template/ts/service_fetch.gots",
		size:    1570,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/4xT3W4bRRi936f4ZFXyj9x1gkQVnCZS+FPNRW0p5gHGu5/tCeuZZWbWwVpGaghuQUpI
JCIkmogWoaKKogYkKCQl8DJe27niFdDOrO2kcIGvdj3nO+fM+c5WSiWnBM0uldCmAQKV0EGGgij0oTWA
fCi44iSkeQNDi/I4U4QyCW3BmULmw0ajBh73EVSXKNjm4gPYpqoLbVRe15y2uYA7zWYDIkk6KDO6hZiZ
phK2BVUKGVAGzUGIm56gocrQBhMK3qc+SiDQIpJ6ltDwGxckCIAwH3pkAAzRB+JvRVL1kCkgnseFT1kH
FAcZokfb1EsZt9BTIPDDiAq0SOaDpCoiinLmlODm//85JZiePhl/9WD08sXk6NH4s8PR+RezHJ0S2JNk
b5gcPks+3588O53+tDs5errRqE0efjp6+d3k8c7ff+wlZy9GF39Njp6aDJOdh2l6FrkguT+0eIu8Ftj0
9IlVz6AHP4wPDkd/HlvJjUbNciWPzscnPy5Un39z+fVwvDtM7v+aHJxOdy8uT+5Nv98ZPz4b7z9Phr+N
Lr68/PZscpw9p0c/fzL6fT853Jve2xtdnCQPzsfHv4yPzpxSxaG9kAsFsQMAEMeCsA7CDTUIsQw3WpwH
UF2DQgdVzQDfJoqkl5DgvhsxL81eFrXOpm/aSdC6nP2DzNfa0WkPe5B3K3HsvhUQKe+SHmpdb23J/Orc
w7xs74ugDCgEF3cI84O0DguKLgYhivyq4/SJSBuWwmENcl2lwmqlsvzGa+7yrRV3efl199ZSdWVpZSm3
6jj4kRFpZ65hE9WbdrYQiaAKUgnKOsUsiQVtJIJVRzuOuZ03856msrgJ3NTacSoViCTaL8qZ66R1v11j
TZNoPVLpw3pBouhTD2eyZeih6nJ/8R4SQXqyCnayWIWG4D0q8XZGAR8Dwz6K9cxwgAqu3APWroZZyK5T
hkx3pldcdcy0QBUJZq2ncZQhnjvKN+qbzXwZWtwfVOG9zfpd12rQ9qBgbRZBF13VRVYQKGFt5ukKc+be
FSh50McU525JzgrFooHqouuRVByF+E+Ca21IUbO5+XJsdRe11NruzEym+zI1fid9MwG69dDWF7T+Vzvi
2LUVLcwWEcdujYU2fK2vLCSO3Xqk5ievLibzb2pwnaMMr4yuF3JxvOiY1rky5OZOcrNSmCubT8v07p8A
AAD//6nV5A0iBgAA
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

	"/generator/template/go": {
		isDir: true,
		local: "generator/template/go",
	},

	"/generator/template/ts": {
		isDir: true,
		local: "generator/template/ts",
	},
}
