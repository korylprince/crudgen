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

var _tmpl_getbykey_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x4f\x4d\x4f\x83\x40\x10\xbd\xf7\x57\x8c\x9c\x16\xd3\xd0\x78\xd5\xf4\x52\x4b\xb5\xb1\x6a\x3f\x4c\x3c\xd3\x76\xa0\xab\xb0\x98\x61\x91\x90\xcd\xfe\x77\x67\x77\x09\x1e\x74\x0e\x90\x7d\x6f\xde\xc7\xcc\x66\x0f\xa8\x8d\x49\x5e\xb2\x0a\xad\x5d\xf4\x4f\xd8\x43\x81\xba\x01\x7d\x41\x18\x71\xc8\xa9\xae\xe0\x7c\x84\x4e\xea\x8b\xa7\x0a\xf9\x8d\x0a\x3e\xb1\x37\x46\xe6\x50\x68\x10\x25\xbf\x13\x96\x37\x31\xdc\x58\xdb\x18\x83\xea\x6c\xed\x24\x6f\xd5\x09\xfe\x64\x08\xf6\x62\x64\xb9\x58\x2b\x8d\x94\x67\x27\x26\xa6\x8c\xac\x78\x7b\x9b\x11\x2f\x32\xbc\x91\x8d\x0e\x96\xd6\xc6\x20\xae\x47\x8b\x29\x20\x51\x4d\x31\x98\x09\xf0\x34\xba\xd2\x1e\x82\xdb\x39\x97\x4c\xb6\x84\x5f\x19\xa1\x88\x0e\xe9\x26\xbd\x7f\x63\xdb\xc3\x6e\xe3\x94\xc1\x70\x25\xb1\x3c\xb3\x25\xac\xf6\xaf\xcf\xae\xc5\xc0\x32\xf2\xfe\x98\xee\xd3\xb0\xff\x6f\x8b\xbb\x28\xf6\x89\x7c\xb2\x8b\xbb\x9a\x83\x92\xe5\xd0\xc2\x0d\xa1\x6e\x49\x39\xd0\xf7\xf1\xb8\xf5\x5f\xaa\x3b\xd7\xce\x55\x4d\x76\x2d\x52\xbf\xaf\x3b\xb1\x5e\x06\xbb\xfa\xf8\xe1\x48\x85\x9d\x18\x6f\x0c\x8c\x4b\x99\x3b\x71\x72\x38\x65\x8a\x59\xf7\xf3\x8d\x22\x16\x45\xbf\xc7\x84\xf5\x21\x9f\xa9\x90\x6f\x27\x3f\x01\x00\x00\xff\xff\x16\x41\xa1\x68\xe0\x01\x00\x00")

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

	info := bindata_file_info{name: "tmpl/GetByKey.tmpl", size: 480, mode: os.FileMode(420), modTime: time.Unix(1432049781, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _tmpl_getlist_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x52\xcd\x6e\x9b\x40\x10\xbe\xf3\x14\x53\xa4\x44\x20\x21\x7c\xa9\x7a\x68\xc5\x25\x89\x53\x59\x75\x92\xc6\xce\xa5\xaa\x7a\x58\x9b\xd9\x6a\x2b\xd8\x45\xb3\xeb\xa6\x08\xf1\xee\xdd\x59\x30\xc6\x56\x73\x29\x92\x0d\x7c\xb3\xf3\xfd\xcc\xb0\x58\x7c\x46\xd7\x75\xf9\xa3\xa8\xb1\xef\xd7\xca\x3a\x20\x74\x07\xd2\x16\x04\x54\xfc\x6a\x24\x4c\x75\x90\x64\x6a\x28\x77\x19\x18\x2a\x91\xb0\x84\x5d\x0b\x42\x83\x69\x9c\x32\x5a\x54\x03\x7c\xd3\x82\x54\x58\x95\xbe\x52\x7a\x8e\x5a\xb9\xe1\x60\x78\x8c\x16\x8b\x95\x9c\xce\x29\xcb\xed\x58\x37\xae\x05\xeb\x48\xe9\x9f\x19\x28\x07\xaf\xaa\xaa\xa0\x44\x29\x0e\x95\x03\x67\x20\xee\xba\xed\xf3\x9a\x3d\x04\x87\xf9\x17\x6c\x6d\xdf\xc7\x91\x3c\xe8\x3d\x5c\x06\x48\xca\x1d\x3b\xbe\xbb\x59\x69\x87\x24\xc5\xde\xe3\xd9\xa4\x78\x54\x09\x66\x40\x69\xf7\xe1\x7d\x0a\x49\x48\xfa\xfd\xc7\xc4\x93\x01\x12\xf1\xcf\x50\x0a\x5d\x04\xfe\x52\x27\xd7\x45\x01\x71\x3c\xc2\x7c\x4d\xf8\x9b\x46\xf9\x54\x1f\x85\x9b\x75\xb5\x1b\xe8\x3f\x16\x7e\x94\xf9\x57\xc2\x46\x10\x26\xb2\x76\xf9\xb6\xf1\xe6\x9c\x4c\xe2\xed\x72\xbd\xbc\x7d\x81\x0b\xb6\x7b\x9e\xaa\xe7\x83\xfb\xcd\xd3\x03\x67\x1c\xab\x1e\x79\xda\xdc\x2d\x37\x70\xf3\x0d\xae\x2c\xac\x57\x0f\xab\x17\xb8\x2a\x3f\xc5\x53\xec\x31\x6f\x9a\x1e\xa3\xb0\xfe\xbb\x02\xb4\xaa\x66\x39\x86\xcd\x33\x18\x0c\xce\x5d\x93\x79\xb5\x93\x6b\x8e\x90\x3f\x1f\x90\xda\xe4\x7f\x08\xf9\xdf\x2f\x17\x09\x78\x7f\x49\x3a\x6b\x40\xa6\x67\xad\xfc\xb6\x32\x16\x47\xfa\xa3\xc4\x51\xe0\xfa\x3a\xc8\x15\x97\x72\x81\x81\x0b\x80\x13\x36\xc8\xf5\x9e\x29\x3c\xfc\x16\x04\x66\xf7\xcb\xce\x97\x1d\x0a\xd2\xd0\x20\xfc\x88\x7f\xdc\x99\xa7\xb1\x05\xce\xcf\x9f\xb4\x42\xd7\x76\x2f\x74\xe2\xd7\xe5\x6f\x61\x57\xb1\xef\x88\x4f\x1b\x3b\xcf\xf1\xaf\x51\xbd\x35\xae\x53\x86\xf0\xa5\xb1\xf5\x02\x44\xd3\xa0\x2e\x13\x7e\xcb\x18\x4b\x67\x93\x1d\x49\x86\x5a\xf0\xb6\x24\xf2\xf1\xfb\xe8\x6f\x00\x00\x00\xff\xff\x6d\xb6\x7a\x1f\xec\x03\x00\x00")

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

	info := bindata_file_info{name: "tmpl/GetList.tmpl", size: 1004, mode: os.FileMode(420), modTime: time.Unix(1432049796, 0)}
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

