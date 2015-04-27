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

var _tmpl_getlist_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x52\x4d\x6f\x9b\x40\x10\xbd\xf3\x2b\x5e\x91\x12\x41\x85\xf0\xa5\xea\xa1\x15\x97\x24\x4e\x65\xd5\x49\x1a\x3b\x97\xaa\xea\x01\x87\xa1\xda\x0a\x76\xe9\xee\xba\x29\x42\xfc\xf7\xee\x2c\x18\x63\xab\xb9\x14\xc9\x06\x66\x76\xde\xc7\x3c\x16\x8b\x4f\x64\xbb\x2e\xbd\xcf\x6b\xea\xfb\xb5\x30\x16\x9a\xec\x5e\x4b\x83\x1c\x15\xbf\xaa\x12\x53\x1f\xa5\x56\x35\x8a\x5d\x02\xa5\x0b\xd2\x54\x60\xd7\x22\x97\x50\x8d\x15\x4a\xe6\xd5\x50\xbe\x6a\x51\x0a\xaa\x0a\xd7\x29\x1c\x46\x2d\xec\x70\xd0\x3f\x06\x8b\xc5\xaa\x9c\xce\x09\xc3\xe3\x54\x37\xb6\x85\xb1\x5a\xc8\x1f\x09\x84\xc5\x8b\xa8\x2a\x14\x54\xe6\xfb\xca\xc2\x2a\x84\x5d\xb7\x7d\x5c\xb3\x06\xaf\x30\xfd\x4c\xad\xe9\xfb\x30\x28\xf7\xf2\x19\xe7\x06\xa2\x62\x87\xb7\xe6\x57\x95\xde\x5c\x25\x13\xd1\x01\xdc\x6b\x80\x90\xf6\xfd\xbb\x18\x91\x37\xf8\xed\xfb\x34\x9e\x80\xb4\xe6\x9f\xd2\x31\xba\x00\xee\x12\x47\xb1\x59\x86\x30\x1c\xcb\x7c\x4d\xf5\x57\xf5\xf1\xa9\x3e\xf0\x37\x63\x6b\x3b\xc0\x7f\xc8\xdc\x06\xd3\x2f\x9a\x9a\x5c\x53\x54\xd6\x36\xdd\x36\x4e\x9c\x2d\xa3\x70\xbb\x5c\x2f\xaf\x9f\x70\x86\x76\xcb\xcb\x74\x78\xb8\xdd\x3c\xdc\x71\x18\x63\xd7\x55\x1e\x36\x37\xcb\x0d\xae\xbe\xe2\xc2\x60\xbd\xba\x5b\x3d\xe1\xa2\xf8\x18\x4e\xb6\x47\xbf\x71\x7c\xb0\xc2\xfc\x6f\x32\x48\x51\xcd\x7c\x0c\x81\x73\xd1\x0b\x9c\xab\xd6\xea\xc5\x4c\xaa\xd9\x42\xfa\xb8\x27\xdd\x46\xff\x03\xc8\xff\x2e\x53\xd2\xe0\xd8\xa2\x78\x36\x40\x0c\xcf\x5c\xe9\x75\xa5\x0c\x8d\xf0\x07\x8a\x03\xc1\xe5\xa5\xa7\xcb\xce\xe9\x3c\x02\x37\x40\x53\x6d\xa0\xeb\x1d\x92\x7f\xf8\x9d\x6b\xa8\xdd\x4f\x33\x0f\xdb\x37\x4a\xa5\x07\xe2\x7b\xfa\x63\x4f\x34\x8d\x23\x38\x3d\x7f\xe4\xf2\x53\xdb\xe7\x5c\x46\x2e\x2e\x77\xf3\x59\x85\x6e\x22\x3c\x26\x76\xea\xe3\x5f\xab\x7a\x6d\x5d\x47\x0f\xfe\x4b\x63\xe9\x19\xf2\xa6\x21\x59\x44\xfc\x96\x70\x2d\x9e\x6d\x76\x04\x19\x7a\x5e\xdb\x52\x6b\x67\xbf\x0f\xfe\x06\x00\x00\xff\xff\x71\x19\x54\x0b\xe3\x03\x00\x00")

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

	info := bindata_file_info{name: "tmpl/GetList.tmpl", size: 995, mode: os.FileMode(420), modTime: time.Unix(1430158915, 0)}
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

