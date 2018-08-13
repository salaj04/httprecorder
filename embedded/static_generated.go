// Code generated by "esc -o embedded/static_generated.go -pkg embedded assets"; DO NOT EDIT.

package embedded

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

	"/assets/index.html": {
		local:   "assets/index.html",
		size:    11858,
		modtime: 1534117361,
		compressed: `
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
`,
	},

	"/": {
		isDir: true,
		local: "",
	},

	"/assets": {
		isDir: true,
		local: "assets",
	},
}