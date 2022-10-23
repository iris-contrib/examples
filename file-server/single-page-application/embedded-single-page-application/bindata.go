// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package main generated by go-bindata.// sources:
// data/public/app.js
// data/public/app2/index.html
// data/public/css/main.css
// data/views/index.html
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

type assetFile struct {
	*bytes.Reader
	name            string
	childInfos      []os.FileInfo
	childInfoOffset int
}

type assetOperator struct{}

// Open implement http.FileSystem interface
func (f *assetOperator) Open(name string) (http.File, error) {
	var err error
	if len(name) > 0 && name[0] == '/' {
		name = name[1:]
	}
	content, err := Asset(name)
	if err == nil {
		return &assetFile{name: name, Reader: bytes.NewReader(content)}, nil
	}
	children, err := AssetDir(name)
	if err == nil {
		childInfos := make([]os.FileInfo, 0, len(children))
		for _, child := range children {
			childPath := filepath.Join(name, child)
			info, errInfo := AssetInfo(filepath.Join(name, child))
			if errInfo == nil {
				childInfos = append(childInfos, info)
			} else {
				childInfos = append(childInfos, newDirFileInfo(childPath))
			}
		}
		return &assetFile{name: name, childInfos: childInfos}, nil
	} else {
		// If the error is not found, return an error that will
		// result in a 404 error. Otherwise the server returns
		// a 500 error for files not found.
		if strings.Contains(err.Error(), "not found") {
			return nil, os.ErrNotExist
		}
		return nil, err
	}
}

// Close no need do anything
func (f *assetFile) Close() error {
	return nil
}

// Readdir read dir's children file info
func (f *assetFile) Readdir(count int) ([]os.FileInfo, error) {
	if len(f.childInfos) == 0 {
		return nil, os.ErrNotExist
	}
	if count <= 0 {
		return f.childInfos, nil
	}
	if f.childInfoOffset+count > len(f.childInfos) {
		count = len(f.childInfos) - f.childInfoOffset
	}
	offset := f.childInfoOffset
	f.childInfoOffset += count
	return f.childInfos[offset : offset+count], nil
}

// Stat read file info from asset item
func (f *assetFile) Stat() (os.FileInfo, error) {
	if len(f.childInfos) != 0 {
		return newDirFileInfo(f.name), nil
	}
	return AssetInfo(f.name)
}

// newDirFileInfo return default dir file info
func newDirFileInfo(name string) os.FileInfo {
	return &bindataFileInfo{
		name:    name,
		size:    0,
		mode:    os.FileMode(2147484068), // equal os.FileMode(0644)|os.ModeDir
		modTime: time.Time{}}
}

// AssetFile return a http.FileSystem instance that data backend by asset
func AssetFile() http.FileSystem {
	return &assetOperator{}
}

var _publicAppJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xcf\xcc\x4b\xc9\x2f\xd7\x4b\xcc\x49\x2d\x2a\xd1\x50\x4a\x2c\x28\xd0\xcb\x2a\x56\xc8\xc9\x4f\x4c\x49\x4d\x51\x48\x2b\xca\xcf\x55\x88\x51\xd2\x57\xd2\xb4\x06\x04\x00\x00\xff\xff\xa9\x06\xf7\xa3\x27\x00\x00\x00")

func publicAppJsBytes() ([]byte, error) {
	return bindataRead(
		_publicAppJs,
		"public/app.js",
	)
}

func publicAppJs() (*asset, error) {
	bytes, err := publicAppJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "public/app.js", size: 39, mode: os.FileMode(420), modTime: time.Unix(1663416115, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _publicApp2IndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xcd\x31\x0a\x02\x31\x10\x85\xe1\x7e\x4e\xf1\x4a\x6d\x0c\xbb\xf5\x10\xb0\xdb\x42\x41\xd0\x0b\x44\x33\x9a\x40\xd6\x0c\x32\x85\xde\x5e\x86\x6c\xf9\xe0\xf1\xfd\x5c\x6c\x6d\x91\x88\x8b\xa4\x1c\x09\x00\xd8\xaa\x35\x89\x47\x55\xcc\x1c\xc6\x20\x0e\xe3\x40\x7c\xef\xf9\xb7\x1d\xcb\x14\xb1\xbb\x5a\xb2\xfa\xc0\x72\x3b\x9f\x70\x49\x2f\xd9\x63\x91\xd6\x3a\x9e\x9f\xbe\x22\xa9\xce\xa1\xbe\xb3\x7c\x0f\x1e\x02\x87\x32\x39\x36\x10\x57\x3d\xff\x0f\x00\x00\xff\xff\xdd\xbe\x30\x69\x85\x00\x00\x00")

func publicApp2IndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_publicApp2IndexHtml,
		"public/app2/index.html",
	)
}

func publicApp2IndexHtml() (*asset, error) {
	bytes, err := publicApp2IndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "public/app2/index.html", size: 133, mode: os.FileMode(420), modTime: time.Unix(1663416115, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _publicCssMainCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\xca\x4f\xa9\x54\xa8\xe6\x52\x50\x50\x50\x48\x4a\x4c\xce\x4e\x2f\xca\x2f\xcd\x4b\xd1\x4d\xce\xcf\xc9\x2f\xb2\x52\x48\xca\x49\x4c\xce\xb6\xe6\xaa\xe5\x02\x04\x00\x00\xff\xff\x96\x97\xac\xb1\x26\x00\x00\x00")

func publicCssMainCssBytes() ([]byte, error) {
	return bindataRead(
		_publicCssMainCss,
		"public/css/main.css",
	)
}

func publicCssMainCss() (*asset, error) {
	bytes, err := publicCssMainCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "public/css/main.css", size: 38, mode: os.FileMode(420), modTime: time.Unix(1663416115, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xce\xbd\x0e\xc2\x30\x0c\x04\xe0\xdd\x4f\x71\xea\x04\x4b\xa2\xee\xc6\x33\x23\x43\x5f\x20\xb4\x86\x04\xa5\x34\x6a\x22\x7e\x54\xf5\xdd\x51\x14\x46\xeb\x7c\xfa\x8e\x7d\x99\xa3\x10\xb1\x57\x37\x09\x01\x00\x97\x50\xa2\xca\xb6\xc1\x5c\xdc\x5d\xcd\x50\x4f\xec\x3b\xdb\x16\x10\xdb\xf6\x4c\x7c\x5d\xa6\xef\xbf\xe4\x7b\xc1\x61\xd0\x39\x45\x57\xf4\x88\xb3\xc6\xb8\xe0\xb6\x2e\x33\x5e\x41\xdf\xd9\x86\xe7\xa4\x1f\x53\x35\xb0\xf5\xbd\x10\xb5\x5e\x1e\xd7\x90\x0a\xf2\x3a\x9e\x3a\xeb\x52\x32\x8f\xdc\x09\xc0\xb6\x05\x55\x6b\x4a\x65\xeb\xd6\x5f\x00\x00\x00\xff\xff\xd6\xa4\xa5\x16\xb2\x00\x00\x00")

func viewsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_viewsIndexHtml,
		"views/index.html",
	)
}

func viewsIndexHtml() (*asset, error) {
	bytes, err := viewsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/index.html", size: 178, mode: os.FileMode(420), modTime: time.Unix(1663416115, 0)}
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
	"public/app.js":          publicAppJs,
	"public/app2/index.html": publicApp2IndexHtml,
	"public/css/main.css":    publicCssMainCss,
	"views/index.html":       viewsIndexHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
	"public": {nil, map[string]*bintree{
		"app.js": {publicAppJs, map[string]*bintree{}},
		"app2": {nil, map[string]*bintree{
			"index.html": {publicApp2IndexHtml, map[string]*bintree{}},
		}},
		"css": {nil, map[string]*bintree{
			"main.css": {publicCssMainCss, map[string]*bintree{}},
		}},
	}},
	"views": {nil, map[string]*bintree{
		"index.html": {viewsIndexHtml, map[string]*bintree{}},
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
