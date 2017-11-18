// Code generated by go-bindata.
// sources:
// pub/index.gohtml
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _pubIndexGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x93\xc1\x6e\xdb\x30\x10\x44\xef\xf9\x8a\xc5\x9e\x2b\xb3\xbe\x15\x05\x29\xa0\x08\xda\x5b\xd0\xa2\x71\x80\x5e\x29\x6a\x6c\xb2\xa1\x48\x95\x5c\x39\xf1\xdf\x17\xb2\xec\xda\x4e\x73\x5c\x72\xf6\x71\x34\x03\x69\x2f\x43\xa4\x68\xd3\xce\x30\x12\xb7\x77\xda\xc3\xf6\xed\x1d\x11\x91\x1e\x20\x96\x9c\xb7\xa5\x42\x0c\x3f\x6d\xbe\x35\x9f\xf8\xfa\x2a\xd9\x01\x86\xf7\x01\x2f\x63\x2e\xc2\xe4\x72\x12\x24\x31\xfc\x12\x7a\xf1\xa6\xc7\x3e\x38\x34\xc7\xe1\x03\x85\x14\x24\xd8\xd8\x54\x67\x23\xcc\x7a\xf5\xf1\x06\xe5\x45\xc6\x06\x7f\xa6\xb0\x37\xfc\xab\x79\xfa\xd2\xdc\xe7\x61\xb4\x12\xba\x88\x2b\x6e\x80\x41\xbf\xc3\x79\x53\x82\x44\xb4\x0f\xc1\x95\x5c\x51\xe6\xc7\x2a\x6d\x72\x8e\x55\xab\xe5\x6a\x91\xc5\x90\x9e\xa9\x20\x1a\xae\x72\x88\xa8\x1e\x10\x26\x5f\xb0\x35\xac\xc6\xa9\x53\xae\x56\x35\xd8\x90\x56\xae\xd6\x39\x02\xb5\x64\xa0\xbb\xdc\x1f\x4e\x8c\x0a\x27\x21\xa7\x65\x3a\x9e\xf8\x75\xbb\x41\x15\x7a\x40\xad\x76\x07\xfa\x89\x5d\xa8\x52\xec\x2c\xd3\xca\xaf\xdb\xbb\x8b\x76\x9b\xcb\x40\x03\xc4\xe7\xde\xf0\x8f\xef\x8f\x1b\x26\x7b\xe4\x19\x56\xce\xc3\x3d\x9f\x20\xd7\x0c\xbe\xbc\x75\x64\xf4\x61\x7f\x7b\xb2\x7c\x9b\xed\x10\x69\x9b\x8b\xe1\x61\x61\x70\x7b\x76\x24\x99\xf0\x6a\x87\x90\xf0\x59\xab\xa3\xf0\x1d\x40\x48\xe3\x24\xa7\x26\xcf\x04\x92\xc3\x08\xc3\x82\x57\x61\x1a\xa3\x75\xf0\x39\xf6\x28\x86\xbf\x26\x41\xa1\x93\x8e\x7e\xd7\x9c\x56\xab\x15\x93\x7a\xe3\x55\xfd\x67\x56\x77\x93\x48\x4e\x27\x72\x9d\xba\x21\x08\xd3\xde\xc6\x09\x86\xef\xe7\x08\xe8\x9f\xff\x9b\x51\xab\x65\xf3\x2a\x79\x35\xc7\x79\x99\x97\x82\xd4\x4d\x43\xef\xf7\xf5\x88\xd4\x5f\xb8\x73\x45\x6f\x56\xb5\x5a\x2a\xd7\x6a\xfe\x2d\xda\xbf\x01\x00\x00\xff\xff\x3b\x5e\x1b\x80\x1d\x03\x00\x00")

func pubIndexGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_pubIndexGohtml,
		"pub/index.gohtml",
	)
}

func pubIndexGohtml() (*asset, error) {
	bytes, err := pubIndexGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pub/index.gohtml", size: 797, mode: os.FileMode(420), modTime: time.Unix(1510915268, 0)}
	a := &asset{bytes: bytes, info: info}
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

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
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
	"pub/index.gohtml": pubIndexGohtml,
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
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"pub": &bintree{nil, map[string]*bintree{
		"index.gohtml": &bintree{pubIndexGohtml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
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

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

