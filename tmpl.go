package crudgen

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _tmpl_getbykey_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x4f\xcb\x4e\xeb\x30\x10\xdd\xf7\x2b\xe6\x66\xe5\x54\x95\xab\xbb\x05\x75\x53\x9a\x02\xa2\x40\x1f\x48\xac\xd3\x76\xda\x1a\x52\x07\xc6\x2e\x51\x64\xf9\xdf\x19\xdb\x51\x58\xc0\x2c\x12\xf9\x9c\x39\x8f\x19\x8f\x6f\xd1\x3a\x27\x9f\xca\x33\x7a\x3f\x6d\x1f\xb0\x85\x23\x5a\x03\xf6\x84\xd0\xe3\x70\xa0\xfa\x0c\xfb\x2d\x34\xca\x9e\x22\x75\x54\x5f\xa8\xe1\x1d\x5b\xe7\xd4\x01\x8e\x16\x44\xc5\x6f\xc9\x72\x93\xc3\x7f\xef\x8d\x73\xa8\xf7\xde\x0f\x0e\x17\xbd\x83\x5f\x19\x82\xbd\x86\xe6\xb3\x92\xb3\xe9\x88\x63\xe6\xbc\xb4\x2c\x89\x79\x8b\xb4\x50\xc6\x26\x27\xef\x73\x10\xc3\x5e\x39\x02\x24\xaa\x29\x07\x37\x00\x1e\x63\xcf\x36\x42\x70\x35\xe1\x6e\x72\x49\xf8\x51\x12\x8a\x6c\x53\x2c\x8a\x9b\x17\xb6\xdd\xac\x16\x41\x99\x0c\xe7\x0a\xab\x3d\x5b\xc2\x7c\xfd\xfc\x18\x4e\xeb\x58\x46\x5e\xef\x8a\x75\x91\xf6\xff\x6c\x71\x9d\xe5\x31\x91\x2f\x0d\x71\xff\x26\xa0\x55\xd5\xb5\x08\x43\x68\x2f\xa4\x03\x18\xfb\x44\xdc\xc7\x2f\xd5\x4d\x68\x17\xaa\xca\xd5\x05\xa9\x5d\xd7\x8d\xb8\x9f\x25\xbb\x7a\xfb\x16\x48\x8d\x8d\xe8\x6f\x4c\x4c\x48\x99\x04\xb1\xdc\xec\x4a\xcd\x6c\xf8\xc5\x46\x19\x8b\xb2\x9f\x63\xd2\x7a\x97\xcf\x54\xca\xf7\x83\xef\x00\x00\x00\xff\xff\x0a\x99\x8e\xb7\xd7\x01\x00\x00")

func tmpl_getbykey_tmpl_bytes() ([]byte, error) {
	return bindata_read(
		_tmpl_getbykey_tmpl,
		"tmpl/GetByKey.tmpl",
	)
}

func tmpl_getbykey_tmpl() (*asset, error) {
	bytes, err := tmpl_getbykey_tmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "tmpl/GetByKey.tmpl", size: 471, mode: os.FileMode(420), modTime: time.Unix(1429908608, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _tmpl_getlist_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x52\x4d\x6f\x9b\x40\x10\xbd\xf3\x2b\x5e\x91\x12\x41\x85\xf0\xa5\xea\xa1\x15\x97\x24\x4e\x65\xd5\x49\x1a\x3b\x97\xaa\xea\x01\x87\xa1\xda\x6a\x61\xe9\xee\xba\x29\x42\xfc\xf7\xee\x2c\x18\x63\xab\xb9\x14\xc9\xc6\xfb\x66\xe7\x7d\xcc\x78\xb1\xf8\x44\xb6\xeb\xd2\xfb\xbc\xa2\xbe\x5f\x0b\x63\xa1\xc9\xee\x75\x6d\x90\x43\xf2\x51\x95\x98\xea\x28\xb5\xaa\x50\xec\x12\x28\x5d\x90\xa6\x02\xbb\x16\x79\x0d\xd5\x58\xa1\xea\x5c\x0e\xf0\x55\x8b\x52\x90\x2c\x82\xc5\x62\x55\x4e\x90\x30\x7c\x93\xaa\xc6\xb6\x30\x56\x8b\xfa\x47\x02\x61\xf1\x22\xa4\x44\x41\x65\xbe\x97\x16\x56\x21\xec\xba\xed\xe3\x9a\xe5\xbc\x99\xf4\x33\xb5\xa6\xef\xc3\xa0\xdc\xd7\xcf\x38\xf7\x1a\x15\x3b\xbc\x35\xbf\x64\x7a\x73\x95\x4c\x42\x07\x72\x29\x2a\xc7\x2f\x6a\xfb\xfe\x5d\x8c\xc8\x67\xf9\xf6\x7d\x6a\x4f\x40\x5a\xf3\x47\xe9\x18\x5d\x00\xf7\x88\xa3\xd9\x2c\x43\x18\x8e\x30\x3f\x13\xfe\xaa\x3f\xbe\xd5\x07\xfe\x65\x6c\x65\x07\xfa\x0f\x99\x1b\x56\xfa\x45\x53\x93\x6b\x8a\xca\xca\xa6\xdb\xc6\x99\xb3\x65\x14\x6e\x97\xeb\xe5\xf5\x13\xce\xd8\x6e\x79\x6e\x8e\x0f\xb7\x9b\x87\x3b\x9e\xfb\x58\x75\xc8\xc3\xe6\x66\xb9\xc1\xd5\x57\x5c\x18\xac\x57\x77\xab\x27\x5c\x14\x1f\xc3\x29\xf6\x98\x37\x8e\x0f\x51\x58\xff\x4d\x86\x5a\xc8\x59\x8e\x61\xb7\x0c\x7a\x83\x73\xd7\x5a\xbd\x98\xc9\x35\x47\x48\x1f\xf7\xa4\xdb\xe8\x7f\x08\xf9\xdb\xed\x94\x34\x78\x6d\x51\x3c\x6b\x20\xa6\x67\xad\xf4\x5a\x2a\x43\x23\xfd\x41\xe2\x20\x70\x79\xe9\xe5\xb2\x73\x39\xcf\xc0\x05\xd0\x84\x0d\x72\xbd\x63\xf2\x3f\x7e\xe7\x1a\x6a\xf7\xd3\xcc\x97\xed\x0b\xa5\xd2\x83\xf0\x3d\xfd\xb1\x27\x9e\xc6\x16\x9c\xde\x3f\x6a\xf9\xae\xed\x73\x5e\x47\x6e\x5d\xee\xe5\x77\x15\xba\x8e\xf0\xb8\xb1\xd3\x1c\xff\x1a\xd5\x6b\xe3\x3a\x66\xf0\xff\x34\xb6\x9e\x21\x6f\x1a\xaa\x8b\x88\x4f\x09\x63\xf1\x6c\xb2\x23\xc9\x50\xf3\xde\x96\x5a\xbb\xf8\x7d\xf0\x37\x00\x00\xff\xff\x45\xa3\xc6\x71\xce\x03\x00\x00")

func tmpl_getlist_tmpl_bytes() ([]byte, error) {
	return bindata_read(
		_tmpl_getlist_tmpl,
		"tmpl/GetList.tmpl",
	)
}

func tmpl_getlist_tmpl() (*asset, error) {
	bytes, err := tmpl_getlist_tmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "tmpl/GetList.tmpl", size: 974, mode: os.FileMode(420), modTime: time.Unix(1430158270, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"tmpl/GetByKey.tmpl": tmpl_getbykey_tmpl,
	"tmpl/GetList.tmpl": tmpl_getlist_tmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"tmpl": &_bintree_t{nil, map[string]*_bintree_t{
		"GetByKey.tmpl": &_bintree_t{tmpl_getbykey_tmpl, map[string]*_bintree_t{
		}},
		"GetList.tmpl": &_bintree_t{tmpl_getlist_tmpl, map[string]*_bintree_t{
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
        if err != nil {
                return err
        }
        err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
        if err != nil {
                return err
        }
        err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
        if err != nil {
                return err
        }
        return nil
}

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

