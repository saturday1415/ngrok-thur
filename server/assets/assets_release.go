// Code generated by go-bindata.
// sources:
// assets/server/tls/snakeoil.crt
// assets/server/tls/snakeoil.key
// DO NOT EDIT!

// +build release

package assets

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

var _assetsServerTlsSnakeoilCrt = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x93\xbb\x16\xaa\x30\x02\x45\x7b\xbe\x62\x7a\xd6\x2c\x41\x10\xa5\x4c\xc2\x43\xe4\xa1\x01\x04\xa4\x13\x51\x20\x08\x21\x80\xbc\xbe\x7e\xd6\xbd\xc5\x34\xf7\x94\xfb\x14\xbb\xda\xff\xfd\x33\xa8\x9b\x96\xf7\x1f\xa4\xfb\xa1\x65\x58\x08\x84\xfa\x5f\xca\xb9\x96\x85\x28\x41\x08\x3c\x28\x42\x18\x7d\x23\x69\xec\x4d\x4f\x83\x1f\x02\x3c\x58\xd4\xac\xac\x2b\x53\x9d\x05\x08\xf0\xdd\x00\x1a\x08\x5d\x5f\x9f\xb5\xf9\xa1\x45\x18\x6b\x1a\x28\x05\xee\x79\x8e\x56\xa7\xf1\xa6\x2c\x04\x6f\x63\x16\x96\xab\x06\x64\x97\xe8\x8b\x47\xf4\xcd\x25\xde\xd3\x98\x85\xcd\x25\x40\x70\x49\xfd\x7f\xe6\x42\x77\xd1\x43\x70\x83\x85\x17\x71\x10\xb8\x2e\x3a\xfb\x34\x4f\xac\xdf\x63\xaf\x8e\xae\x65\x41\xeb\x1f\xbb\x6e\x00\x70\x45\x00\x9f\xc0\x9f\x1f\x15\x36\x02\x58\x07\x75\x7d\xd1\x39\x18\x8d\xd7\xc3\xc5\xac\x7e\xb5\x26\x0b\xae\x66\x7d\x47\x61\x64\x0d\xe5\x9b\xf8\xdb\x4b\x25\xa3\x19\x9f\xc8\x45\x15\x7f\x90\xf9\x26\xaf\x30\xba\xbe\xe4\x0b\xfb\xa5\xe1\x7a\xe4\xbb\xe7\x44\x7c\x89\x33\xd7\xdb\xa0\xa6\xde\xaa\xfa\x38\x8e\x3f\xc7\x69\x73\x03\x3b\xbe\xe9\xdf\x21\xb2\xab\xac\x5f\x89\xcc\xda\x3b\x4c\x6e\x3b\x07\xc0\x23\xe3\x15\xf3\x13\x11\xab\x53\xc3\xeb\x94\xe7\x92\xd5\x54\x25\x27\x5e\x5b\xd5\x17\x71\x0e\xe7\x1e\x4f\xcf\xe7\x81\x88\xdf\x3d\x25\x17\xd6\xb2\xd3\xc0\x9f\x0e\xbd\xd9\x38\x32\x79\x1b\x94\xf0\x24\x3c\x3b\x6a\x75\xd1\x0b\x47\x06\xb7\x63\xc9\x72\x31\xbe\x7c\x3a\xce\x96\xc5\xe1\x33\x9f\x95\x0e\xd9\x0d\x73\x8a\xa0\xba\x33\x30\xdb\xd3\x16\x80\x74\x3f\x4d\xcb\x7e\xda\x8c\x32\x18\x8f\x41\x2d\x6f\x0c\xab\x25\x02\x8d\x8d\x05\x74\xa1\xd2\xae\x53\xae\x3b\xf8\xe5\xf4\xbd\x6f\xee\x76\x3a\xbf\xb6\x32\xd0\x5e\x95\x4c\xfb\xb8\x7c\x86\xe2\x89\x57\x78\x1f\x2c\xf1\x2e\xfd\x15\xa6\x0d\x23\xf7\x25\xb7\xbc\x26\xce\x69\x6d\x55\x07\xe1\x2b\x98\x6f\x32\xe1\x86\xfa\x0e\x17\xdd\x98\xf8\xd6\x5a\xf6\xdb\xd2\x35\x2b\x72\x2d\xc7\x96\x06\x30\x80\x2e\x10\x4c\x14\x30\x33\xb0\x32\x49\xc3\x3a\x84\xf8\x0e\x80\x6c\x41\x80\x11\x65\x89\x8f\xc4\x6b\x3e\x74\x7d\xe9\x72\x12\x8a\x9f\xe2\x69\x2d\x6f\xcb\x05\x45\x0a\xf3\x4b\xd6\x93\x31\x6d\xfc\xc4\x48\x56\x25\x90\x60\xb6\x9d\xc5\xb6\xba\x3b\xc9\xfb\xcb\xdf\x69\xce\x87\x20\x85\x46\x73\x14\x2b\xb7\x8f\x32\xc5\xf7\xb8\xa7\x35\xe7\x36\x79\xbc\xfa\x61\xbb\x0f\x3f\xaa\xc8\xf2\x72\x4b\xbb\x31\xa1\x96\xa2\xbd\x1e\x89\xd7\x87\xdf\x6a\x7f\x98\x4f\xb5\x31\xad\x8f\x7b\x5e\xa6\xbd\xc8\x3b\x81\x58\x1f\xe2\xc2\xc8\x7f\x75\xc6\x8d\xdb\x71\x31\xe6\xbe\x40\xfe\x87\x0d\xa7\xa4\x0a\x54\xfb\x18\xd7\x07\x5a\x3b\x83\x71\x3e\xbc\x50\x56\x14\xb2\xbd\x6f\x77\xb9\x3e\x7b\x1f\xe1\x44\xf1\x70\x24\x2f\x6a\xe1\x65\xf1\x36\x5a\xe8\x31\x17\x8f\x03\x93\x3b\xcd\xf0\x5b\x53\x96\x86\xae\xcb\xf6\xd0\x16\x2a\x92\x75\x23\x71\xc6\x5f\x88\x78\x33\x51\xc0\x9d\x8e\xfe\x59\x7d\xc8\xab\x7a\x41\x73\x20\x7d\x9b\xb5\x3d\xe7\x42\xa9\xcc\x16\xe1\xaa\xbd\xed\x5c\xa0\x2e\x54\xa1\x71\x0b\x0f\x04\x3c\xce\x24\xa0\xe6\x5d\x09\x7b\x73\xb1\x7b\xf3\x27\xa1\xdb\xf9\x38\x67\x8f\x96\xaf\x5e\x7e\x55\x16\x18\x29\xd7\xa4\x97\x17\x97\x44\x49\xda\xee\x19\xe7\x81\x19\x51\xed\x14\x72\x7f\xe3\xd4\x3d\xed\xdf\x60\xff\x17\x00\x00\xff\xff\xc5\xba\x13\xa4\xcd\x03\x00\x00"

func assetsServerTlsSnakeoilCrtBytes() ([]byte, error) {
	return bindataRead(
		_assetsServerTlsSnakeoilCrt,
		"assets/server/tls/snakeoil.crt",
	)
}

func assetsServerTlsSnakeoilCrt() (*asset, error) {
	bytes, err := assetsServerTlsSnakeoilCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/server/tls/snakeoil.crt", size: 973, mode: os.FileMode(420), modTime: time.Unix(1534868018, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsServerTlsSnakeoilKey = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xd5\x37\x12\xab\x5a\x02\x04\xd0\x9c\x55\xbc\x9c\x9a\xc2\x0b\x08\x7e\x70\xb9\x20\xbc\xf0\x20\x91\xe1\xbd\x17\x20\x58\xfd\xd4\x7f\x93\x4e\xa7\x9d\x74\xd4\xe7\x3f\xff\x46\x90\x64\xf5\xf5\xc7\xf5\xc0\x1f\xdb\x55\x43\xe0\x4b\x7f\x74\xe9\xf3\xb7\x41\x4c\x55\x95\x66\xa0\x0a\x00\xe8\x10\x38\x12\xe8\x3a\x4d\x12\xc2\xaf\xc5\x68\x72\xb3\x77\x22\x8d\x9b\xa2\xda\x7f\xf1\xef\x32\x4c\xe8\x10\xf5\x2b\x55\x2f\x53\x8a\xbe\xe9\xaa\x89\x4a\x28\x17\x2d\x92\xf9\xa1\x95\xd1\xda\xb2\xc7\xfe\xc5\xa2\x73\x72\xb4\x2e\x25\x5f\xf6\xc6\xc7\xaf\x8b\x77\x9d\x28\x2a\xd9\xe3\x36\x3d\x3d\xb2\xa5\x7e\x0b\xf5\x26\x5d\xaf\x96\x5e\xc6\x40\x78\xdb\x98\x01\x04\x16\x59\xd0\x87\x5c\x86\xad\x3a\xf3\xbe\x75\xe4\x39\xa5\x0e\x4d\x4d\x58\x23\xef\x12\x4e\x2e\x9c\xab\x73\x24\x09\xd3\x12\x3d\x39\xb5\xda\x32\x2e\xdc\x86\x72\xcc\x2a\x0f\x06\xdd\x16\xcf\xa9\x45\x5b\x1f\x51\x0c\xbe\xd1\xa4\xca\xa0\x81\xcd\xd6\x4b\x4e\x44\x5a\x39\xeb\x34\xb1\x95\xa7\xf2\x98\xa1\x3e\x2c\x46\xe5\x35\xc1\x02\x4e\xfd\xb8\x3d\x10\x93\xc7\xf1\x23\x8f\xfb\x59\x7b\x5f\xd6\xeb\xe8\x7b\x41\x1c\xbe\x86\x60\xd0\x1d\x1c\x6a\x13\x85\xcd\x0f\x0b\x13\x7a\x89\x74\x65\x0c\x93\xd0\x6b\xa4\x81\x98\x35\xf4\xb4\x46\x75\xe2\x13\x1c\xfa\x40\x5d\xf0\x8b\xb0\x78\xaf\x64\x5d\x08\xcd\x8c\x1e\x51\x44\x24\xce\xb8\x53\x1b\x06\xef\x71\xb9\x68\x0f\x67\x98\x5c\x23\xb4\x17\xa2\x10\xc7\x65\xbf\xe3\x2b\xad\x72\x31\x77\x54\x11\x38\x40\x00\x93\x2a\x00\x07\x6a\x33\x1e\x63\x17\x0c\xda\x44\xd3\x00\x62\x03\xf0\xe3\x1c\x34\x7e\x65\xa5\x8b\xfa\x09\xdb\x4c\xef\xc0\xc6\x45\x68\x41\x06\xfc\xe8\x81\x15\xfb\x82\x9b\x79\xe0\x87\x16\x83\xd5\x79\x7f\xff\xea\x7b\x9d\xc2\xd7\x66\x6f\x0f\xa6\x5c\xd5\x08\xf1\xde\x89\xcf\x5d\xef\xac\x6b\xa6\x31\x4f\xd2\x5e\x9a\x54\x60\xb9\x9d\xb0\xce\x29\xa5\xec\xf3\xab\x9d\x8d\x37\x4c\xf4\x4f\xa9\x3b\xc4\x10\xec\xb6\x7f\x55\xb4\x13\xd8\xe4\xf6\xca\xb4\x1c\x2b\xbf\x08\x9f\xa7\x4b\x8e\x12\xb7\xd6\xf3\xfa\x7b\xd9\xcc\x42\x83\x19\x59\xfa\xf5\xc0\x3e\x42\xc7\x92\x1f\x24\xa3\xfe\x8c\x2d\xde\x9f\x1e\x21\xfa\x86\xfd\xb1\x0e\x32\xb2\xde\x1e\x65\x1c\x32\xc3\x2e\xbe\x85\xcc\x03\xf6\xe4\x78\xc3\xef\x02\x25\xdb\x7b\xf3\x16\x67\x56\xf1\x67\x1d\x6e\xee\xf9\xfa\x12\x0f\x00\x93\xfa\x91\xcf\x38\xde\x2f\xba\xee\x26\x21\xa3\x64\x0a\x19\xfe\x86\x16\x8d\x02\xa9\x40\xe9\x11\x11\x7c\xc1\x49\xba\x3c\x18\x56\x66\xab\xab\x85\x0a\xab\xf3\xd7\x6f\x5c\xd7\xf7\x34\x1e\x0d\x38\xea\xaa\x04\xd8\x0d\x98\xec\xa0\x81\x3c\xed\xda\xd7\xeb\xba\x82\xfd\x66\x1f\x8c\x82\x41\x54\xae\x10\x27\xc0\x54\x22\x48\x6a\x30\xc9\x02\x30\x8d\x64\x99\xb7\x6f\x3a\xd7\x85\x7c\xc5\x31\xc5\x82\x31\x1c\x16\x81\x21\x97\x00\x92\x7d\x1f\xbd\xfc\xad\x79\x0b\xa1\xf1\x35\x4f\x1e\xd2\xe8\x8f\x32\x33\x84\xd2\x58\x8f\xac\xdb\xf3\x0d\x07\x67\x4f\x2b\xef\x39\xa1\x93\xc3\x28\xd7\x7e\x17\xc9\x1e\x9d\x46\x99\x66\x74\x86\x5a\xb0\x09\x0b\x93\x1e\x4d\x16\x5e\x3f\xe9\x37\x7d\xdf\xb5\xe1\xa9\x33\xb9\x89\x08\x03\x2a\x3b\xb1\xcf\x4f\x42\x4a\xe7\x14\x17\x28\x07\xf2\x72\xbe\x0e\xc2\xeb\x4b\x9b\xda\xcb\xa8\x3c\xa0\xab\x70\x46\xab\x96\x79\xb9\x47\x80\x73\x51\xe8\xe3\x54\xcb\xff\x6f\x31\x18\x91\x93\x64\xe7\xec\x69\xbe\x9a\xe3\x84\x4e\x80\xcb\xce\x85\x33\x0e\x7f\x5b\x63\x7b\x9e\xf9\x85\xd6\x83\x1e\x56\x89\x72\x92\x0d\xa4\xe2\xc7\xb7\x40\x5b\x19\x83\xab\x21\x52\x24\xfc\x61\x2f\x6f\x23\x90\x7e\x71\x5f\xc6\x79\x38\xf2\xd4\xa4\x36\x61\x71\x31\x1a\x61\x29\xeb\x47\x34\x2d\x4b\xbf\xa5\x19\xf0\x31\x52\x25\x1e\x94\xee\xf4\x1d\x2b\x72\x7d\x1f\xd7\x86\x87\x11\xc3\x11\x75\x96\x14\x72\xea\x21\x52\x71\x25\xc3\xfd\x90\xac\xe6\x23\x49\x79\x6a\x1d\x04\x1d\x45\x15\x53\xec\x5d\x23\xa8\x30\x40\x45\x2e\xcd\x23\x77\x8c\xc1\x24\x83\x22\xc1\xc2\x81\x7f\xb5\x63\xe5\x8c\xa6\x7c\x9e\x5a\xa9\xad\x88\x16\x36\xac\x68\x3b\xe3\x2a\xcd\x68\x4b\x2d\x13\xd9\x33\x1f\x9b\x3b\xae\x6d\x60\xe1\x47\x7a\xca\xc2\x30\x8e\xcd\xb0\xbb\xea\xee\x94\x4f\x36\xc0\xd5\xcc\xf4\x53\x56\x95\xd1\xd0\x50\xac\x58\xbb\x91\xee\xd9\xba\xc2\xee\x14\x53\xc4\x3c\x4a\xff\x84\xef\xce\x0c\xad\x1b\x98\xb8\xa8\x2a\x11\x4a\x5e\xf2\x37\x91\xc8\x62\x1d\x34\xf9\x4c\xca\x69\x2c\x47\x63\xc3\x2e\xc2\x91\x8a\x58\xa3\xce\xf9\xe9\x22\xbc\xb7\x6c\xcd\xd5\xbd\xd6\x80\x7b\xf4\x29\x67\xcf\x4f\xf8\xfd\xa8\x12\xac\x3e\xc0\x26\xd3\xf2\xed\x49\x8c\xb6\xd0\x8c\x1b\x93\x1f\xaa\xeb\xf5\xcd\xe9\x4d\xe9\xc8\x98\x7b\xd5\x6c\xd5\x07\xf4\x8c\xc4\x8d\x40\xfb\x13\x2f\xe8\xf6\xa3\xb3\x82\xa8\x8a\x46\x58\x05\xca\x93\x6c\x1a\xf4\xd1\xaf\x9b\x19\x4f\x03\xe0\x4a\xab\x78\x7b\x40\xd6\x5b\xe6\xf8\xdd\xa2\x7b\x00\x90\x79\x76\xe4\x69\x28\x9c\x90\x34\xc8\x7a\x25\xe3\x43\x46\xa1\xa5\x55\x7c\xe9\xc7\x4f\xe8\xc1\xb6\xed\xa6\x61\x68\xa4\x42\xa3\x02\x5c\x23\x8d\xf8\x64\x83\x4e\x7c\xa0\x39\x76\xd1\x51\x0c\x27\x47\xe6\xe8\xbb\x35\xa2\xa4\xa3\x90\x4a\x7c\x60\x81\xa3\x0b\x95\x03\x7d\x0a\x55\x04\x1c\x18\xe3\xfd\x61\x06\x70\x97\xa7\xa5\x68\x98\x44\x5c\xbf\x0b\x9c\x89\x22\x84\xda\x04\x84\xb7\xfd\xe4\x13\xf2\xb3\x4d\x7d\xc4\x5f\x53\x78\x12\xc8\x5d\xbc\x42\x6d\x13\xaa\x9c\x70\xd6\xf2\xc9\x2b\x63\x25\x07\x54\xb5\xc8\x5e\xdc\xaf\xaf\x29\x72\x1d\x73\x34\x9f\xfc\xda\xb1\x95\xd9\x6e\x2a\xe9\x4e\xbe\x38\x2b\x16\x47\xe2\xff\xde\xc7\x40\x3b\x88\xb6\x9f\x95\xf8\xcb\xeb\xec\xc9\xd5\xe4\x1d\xf5\x6b\x9b\x92\xdf\xf0\xfb\x39\xec\x82\x74\x87\x7a\xe6\xbe\xe5\x95\xc9\x1f\x11\x2c\xdd\x3d\x5f\x0e\x4d\x87\x5c\xa3\x83\x7f\xfe\x41\xfe\xb2\x22\xbd\xc4\xff\xcf\xcd\x7f\x03\x00\x00\xff\xff\x7c\x5a\x1f\xf4\x8f\x06\x00\x00"

func assetsServerTlsSnakeoilKeyBytes() ([]byte, error) {
	return bindataRead(
		_assetsServerTlsSnakeoilKey,
		"assets/server/tls/snakeoil.key",
	)
}

func assetsServerTlsSnakeoilKey() (*asset, error) {
	bytes, err := assetsServerTlsSnakeoilKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/server/tls/snakeoil.key", size: 1679, mode: os.FileMode(420), modTime: time.Unix(1534868018, 0)}
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
	"assets/server/tls/snakeoil.crt": assetsServerTlsSnakeoilCrt,
	"assets/server/tls/snakeoil.key": assetsServerTlsSnakeoilKey,
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
	"assets": &bintree{nil, map[string]*bintree{
		"server": &bintree{nil, map[string]*bintree{
			"tls": &bintree{nil, map[string]*bintree{
				"snakeoil.crt": &bintree{assetsServerTlsSnakeoilCrt, map[string]*bintree{}},
				"snakeoil.key": &bintree{assetsServerTlsSnakeoilKey, map[string]*bintree{}},
			}},
		}},
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

