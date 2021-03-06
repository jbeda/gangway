// Code generated by go-bindata.
// sources:
// templates/commandline.tmpl
// templates/home.tmpl
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

var _templatesCommandlineTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x58\x6d\x57\xdb\xb8\x12\xfe\xce\xaf\x98\xf5\xf6\x9c\x6d\x0f\x55\x14\x28\x50\xda\x1b\xe7\x9e\x34\x40\x9b\x16\x1a\x4a\x42\x29\x9c\xfd\xb0\xb2\x3c\xb1\x45\x64\xc9\x48\x72\x5e\xe8\xe5\xbf\xdf\x23\x27\x06\x12\x52\xca\xbd\xdb\xfd\x12\xa3\x19\xcd\x3c\xf3\xcc\x8b\x25\xd3\xf8\x6d\xaf\xdb\xee\x9f\x1f\xef\x43\xea\x32\xd9\x5c\x6b\xf8\x07\x48\xa6\x92\x30\x40\x15\x78\x01\xb2\xb8\xb9\x06\xd0\xc8\xd0\x31\x48\x9d\xcb\x09\x5e\x15\x62\x14\x06\x6d\xad\x1c\x2a\x47\xfa\xd3\x1c\x03\xe0\xb3\x55\x18\x38\x9c\x38\xea\xdd\xfc\x0b\x78\xca\x8c\x45\x17\x9e\xf6\x0f\xc8\x6e\x40\xef\xdc\x28\x96\x61\x18\x8c\x04\x8e\x73\x6d\xdc\x3d\xe3\xb1\x88\x5d\x1a\xc6\x38\x12\x1c\x49\xb9\x78\x09\x42\x09\x27\x98\x24\x96\x33\x89\xe1\xc6\x4b\xc8\xd8\x44\x64\x45\x56\x09\x6a\xf5\xb9\x6b\x27\x9c\xc4\xe6\x7b\xa6\x92\x31\x9b\x36\xe8\x6c\xb9\xe6\x35\xbf\x11\x02\xed\x5e\x0f\x80\x90\x72\xa7\x14\x6a\x08\xa9\xc1\x41\x18\x78\x46\xf6\x2d\xa5\x03\xad\x9c\xad\x25\x5a\x27\x12\x59\x2e\x6c\x8d\xeb\x8c\x0a\xae\xd5\xbf\x07\x2c\x13\x72\x1a\x1e\x31\x87\x46\x30\xb9\xde\xe1\x5a\xd9\x00\x0c\xca\x30\xb0\x6e\x2a\xd1\xa6\x88\x2e\xb8\x73\xbc\xac\x59\x42\xe2\xb1\xba\xb4\x35\x2e\x75\x11\x0f\x24\x33\x58\x22\xb1\x4b\x36\xa1\x52\x44\x96\x66\x73\x1c\x71\x8d\xb4\x5e\xdb\xa8\xd7\x6b\x9b\x94\xdb\x05\x79\x2d\x13\xaa\xc6\xad\x0d\xe6\xec\x80\x10\x38\x36\xc2\x66\x15\x3f\xcb\x8d\xc8\x1d\x58\xc3\x9f\x0c\x9b\x7b\x7b\xba\x51\xdb\xd8\xaa\xd5\x67\x8b\x12\xe5\xd2\x06\x20\x94\xc3\xc4\x08\x37\x0d\x03\x9b\xb2\xcd\xed\x1d\x72\xd9\x7f\x7f\x7d\xd8\xba\xea\xb6\x78\xf7\x70\xbd\x75\xb8\xb7\x39\xd8\x74\xa3\x83\xf3\xd7\x03\xbb\x13\x8f\x87\x5d\xd4\xbb\xbb\x13\x51\x7c\xfd\xc0\x4e\x86\x61\x00\xdc\x68\x6b\xb5\x11\x89\x50\x61\xc0\x94\x56\xd3\x4c\x17\x36\x68\x36\xe8\x2c\xce\x5f\x12\x32\xd7\x59\xae\x15\x2a\x37\x97\x93\x88\xd9\xf4\x11\x0a\xed\x74\xdb\x8c\xb6\x5d\xd2\x8a\xbf\x1e\xb5\xd3\xd7\x67\x57\x51\xfd\xb4\xcb\xbf\x0d\xbf\xb4\x3e\x7c\x3e\xe8\xe1\xb5\x58\xaf\x8f\x68\xfd\x7a\x67\x32\x18\x3f\x95\xc2\x2f\x28\xfe\x02\x25\x97\x62\x86\x15\x1d\xa7\x33\x6d\x8c\x1e\xdf\xd6\x7e\x05\xa7\xad\xde\x9b\x62\x70\x62\x36\xf6\xaf\xd8\xe9\xc1\x01\x1e\xbd\xa1\xdb\x9b\xef\x3f\xec\xec\x7e\xb0\x1b\xbd\x68\x34\xd9\xc5\x83\x6f\xef\xce\x72\xb9\x7b\x7d\xdc\xf9\x31\x27\xa0\xff\x08\x99\x5c\x16\x89\x50\x96\x3a\xad\x65\xc4\xcc\x2d\xab\x72\xf5\x18\xa9\xc9\xf9\x6b\xda\x3b\x65\xaf\x77\xde\x98\xfa\xf1\xc5\xc6\xd4\x5d\x1c\x1f\x5c\x6e\x9e\x5d\xa9\xee\xc5\x39\x3b\xef\x7d\x8a\xbe\x6d\xa4\x5f\xfb\xe2\x8b\xe4\xad\x9f\x92\xfa\xdb\x73\xf1\x53\x12\xab\x9b\xad\x3b\xfa\x74\xfe\xf1\x90\x9f\x9f\x1c\xbf\xba\xe8\x1c\x0b\xd9\xaf\xbf\x32\x53\x35\xbd\x18\x0c\xe3\xf7\x07\xe3\xeb\xf6\xf8\x74\x77\xf3\xf3\xbb\x2d\xda\xea\x6f\x3d\xa5\xd9\x00\xfe\x4f\x32\x5c\x8a\x3c\xd2\xcc\xc4\xb5\x4b\x4b\x37\x6b\x75\x3f\x37\xb7\xa2\x79\xf8\xbf\x76\x2c\xab\x8c\x71\x9d\x4f\x89\xd3\xe4\x16\x6e\x9e\xbb\x07\xf2\x47\xb2\x68\xd7\x2f\x36\xec\xbb\xd3\x2f\x07\x8c\x8d\x5f\x4f\xb0\xa5\xce\x22\xaa\x7b\xbb\xc9\xbb\xa3\xad\xa3\xfd\x4f\x62\xff\xec\xe8\xe4\x63\x3d\x5f\xa7\x93\x88\x3f\x69\x64\x1b\xb4\x3a\xdc\x00\x1a\x91\x8e\xa7\xb3\x3f\xcb\x9f\x52\xa6\xd8\x08\xb8\x64\xd6\x86\x81\x14\x49\xea\x48\x24\x0b\x04\xff\x13\x80\xd1\x12\xc3\x40\xb1\x91\x48\x98\x13\x5a\x05\xcd\x5b\xbb\xd2\x36\x16\xb7\xb6\x8a\x8d\xc8\xd8\xb0\x3c\x47\x53\x9e\x74\x4c\x28\x34\x41\xb3\xc1\x40\xc4\x61\x20\x75\xa2\xc9\x9d\x78\x3e\x62\xbf\x07\x95\x75\x64\x98\x8a\x89\xdf\x15\x34\x93\xea\x6c\x63\x4b\x68\x85\xac\xb6\x1b\x1f\x28\xa4\x22\x46\xa2\x15\xc9\x30\x26\xde\x3c\xd6\xe3\xe5\x08\x4b\x3b\x29\x7c\x18\x33\x48\xea\x31\x0a\x17\x34\x0f\xcb\xa7\x07\x69\x50\x29\x96\x90\x68\x21\xcb\x53\x67\x11\xdc\x13\xf1\x34\x33\x1d\x09\x89\xb7\xb1\x5b\x1f\x86\x62\xa3\x9f\x41\xff\x1e\x34\xf7\x90\xeb\x18\xe1\xe3\x59\xff\x31\xe0\x05\xc9\x9d\x35\xc4\xcc\x31\xc2\xb8\x13\x23\xe6\xd0\xae\x8c\x25\x2a\x9c\xd3\x8a\x70\x2d\x25\xcb\x2d\x06\xcd\x86\xa8\x54\xd5\xd1\x4a\x44\x79\xb4\x37\x33\x54\x45\x83\x8a\xe6\xc3\x3c\xd3\x58\x8c\xee\x44\x0d\xaa\xd8\xbd\xe5\xda\xaa\xe2\xdf\x2b\xf8\xa2\xab\x74\xab\xda\xe2\xbb\xd0\xb7\x06\x2a\x87\x06\x62\x66\x86\xa8\xc8\xab\x15\x39\x3b\x43\xc9\x75\x86\xf0\xfd\x3b\xd4\x4e\x2d\x1a\x7f\x8d\x82\x9b\x9b\xda\x52\x8c\xe9\xd6\x32\xd4\xf6\x43\x5f\x1d\x05\xda\x78\x58\xa7\x21\x41\x07\x5c\x67\x59\xd9\x68\x42\x21\x30\xce\xd1\x5a\xaf\x72\xe9\x0c\xae\x2d\x0b\xeb\xd0\x7c\x9e\x21\xc2\xa7\x22\x42\xa3\xd0\xa1\x05\x3e\xd3\xbc\x84\xa9\x2e\x60\x2c\xa4\x04\x85\x18\x7b\x5b\xae\xd5\x40\x24\x85\x41\xe8\xe6\xa8\x3a\x7b\xd0\xd6\x4a\x21\x77\xf0\xbc\xdb\xd9\x6b\xbf\x00\x56\xb8\x14\x95\xe0\xe5\xfc\xc0\x40\x1b\xef\xc2\x00\x97\x02\x95\x7b\x40\x6a\xbb\xb9\x4c\x61\x71\x47\x64\x96\x58\xe7\x0f\x49\xf7\x53\x5c\x08\xfd\x3e\xe7\xc2\x09\x29\xdc\xf4\x25\x0c\x8b\x08\xb9\x93\xfe\x9e\x39\x85\x08\x41\x28\xeb\x98\x94\x18\x83\x14\x43\x04\xab\xdf\x2e\x85\x76\x0b\xb4\x30\x17\x8d\xdc\xe0\x52\x04\x8d\xb2\xc5\xab\x37\x0a\x53\x49\xc1\x12\x2c\x2f\x2a\x41\x73\xed\x19\xf0\xc2\x48\x20\x87\x5d\xa8\x5e\xb4\xd6\x69\xc3\x12\x5c\xbe\x95\x0e\x6f\x19\x10\x83\x12\x99\x45\x5a\x3d\xff\x9a\xf9\xb0\x7f\xc3\x85\x75\x2c\x92\x58\x73\x13\xf7\x17\x8d\x84\xa2\xcf\x9e\x17\x65\x9f\xfd\x07\xd8\x78\x08\x7f\x7c\xcf\x8d\x50\x0e\x9c\x96\x7a\x8c\xe6\xf9\xb3\xfa\x8b\x9b\x3f\x5e\x50\x96\xc5\x3b\x5b\x74\x9e\x39\x4f\x25\xcd\x74\x0c\xeb\x13\xa8\xdd\x13\xda\x22\xd6\x90\x8d\xee\x64\x40\x0b\x6b\xa8\xd4\x9c\xc9\x12\xaa\x12\xc3\x52\xda\xa8\xcf\xdb\x42\x2e\x1b\xf4\x41\x76\x57\xd5\xbb\xab\x38\x56\xf5\x04\x61\xef\x4a\x39\xeb\x56\x5f\x60\x9c\x20\x2f\x1c\x96\x8d\x3e\xd0\x52\xea\xb1\x50\xc9\xdb\x06\x8d\x96\x27\x3f\x5f\x86\x7b\x50\xde\x9f\x14\x18\x79\xaa\x21\xb8\x37\x4c\xed\x16\xdc\xdc\x04\xf0\x27\x34\x81\x33\xf2\x70\xca\x6a\x39\x66\x6b\x55\xf4\xb3\x61\x02\x8b\x8e\xcc\x27\x6e\xd5\x5c\x12\x62\xd1\x8c\xd0\x84\x5e\xd7\x3a\xee\xf4\xca\xd5\xe9\xc9\xe1\x4c\xc9\xd1\x38\x31\xf0\x13\x87\xc4\x4f\x9f\x2e\xcf\xd7\xff\x0d\xdc\x60\x8c\xca\x7f\x90\xd9\x32\x80\xfd\x8c\x09\xe9\xbd\xc3\x9f\x65\x36\x48\xe9\x98\xe4\x46\x8f\x44\x8c\x26\xd4\x22\xe6\x3f\xd0\x11\x66\x92\x50\xc4\x39\x11\xd6\x16\x68\x48\x61\x64\x19\x77\xa7\x5c\xce\x83\x7e\xc4\x74\xf6\xa2\x20\x22\x0e\x67\xe1\xfb\x55\x67\xef\x69\x46\x16\xb9\x41\x77\xcf\xb0\x57\x0a\xbc\xf1\x8f\x6d\x0d\x0e\x0c\xda\x94\x38\x3d\x44\x55\xda\x9e\xcc\x24\x7d\x2f\x78\xdc\x56\xc4\xf7\xcc\x3a\x7b\x95\xc5\xca\x14\xfb\xaf\xe1\x89\x5b\x5d\xdf\x79\xf1\xc3\x95\xca\xc2\xce\x35\x55\x55\x96\xdd\x17\x16\x1f\x71\xbf\xd4\xce\x2b\xe6\x6e\x79\xf0\xee\x1d\x88\x0d\x3a\xbb\x42\x35\xe8\xec\x5f\x09\xff\x0d\x00\x00\xff\xff\xf5\x13\x6f\xc8\x5b\x10\x00\x00")

func templatesCommandlineTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesCommandlineTmpl,
		"templates/commandline.tmpl",
	)
}

func templatesCommandlineTmpl() (*asset, error) {
	bytes, err := templatesCommandlineTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/commandline.tmpl", size: 4187, mode: os.FileMode(420), modTime: time.Unix(1524675103, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHomeTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x56\x51\x6f\xe3\x36\x0c\x7e\xef\xaf\xe0\xe9\x5e\x36\x2c\xb2\x9b\xbb\x15\xd8\x72\x76\x86\xad\xdd\x80\x62\x77\xeb\x80\x76\x07\xec\x91\x96\x19\x9b\xad\x2c\xf9\x24\x39\x6e\x6e\xd8\x7f\x1f\x24\xc7\x49\x9a\x1b\xb0\x05\x48\x2c\x52\x14\xf9\x91\x1f\x69\xa5\x78\x75\x73\x77\xfd\xf0\xe7\xef\x3f\x43\x1b\x3a\xbd\xbe\x28\xe2\x03\x34\x9a\xa6\x14\x64\x44\x54\x10\xd6\xeb\x0b\x80\xa2\xa3\x80\xd0\x86\xd0\x4b\xfa\x34\xf0\xb6\x14\xd7\xd6\x04\x32\x41\x3e\xec\x7a\x12\xa0\x26\xa9\x14\x81\x9e\x43\x1e\xdd\xbc\x03\xd5\xa2\xf3\x14\xca\x3f\x1e\x7e\x91\xdf\x89\xfc\xe8\xc6\x60\x47\xa5\xd8\x32\x8d\xbd\x75\xe1\xe4\xf0\xc8\x75\x68\xcb\x9a\xb6\xac\x48\x26\x61\x01\x6c\x38\x30\x6a\xe9\x15\x6a\x2a\x97\x0b\xe8\xf0\x99\xbb\xa1\x9b\x15\xd9\xe5\xde\x75\xe0\xa0\x69\xdd\xa0\x69\x46\xdc\x15\xf9\x24\x5e\xc4\x9d\x57\x52\xc2\xf5\xfd\x3d\x80\x94\xc9\x52\xb3\x79\x82\xd6\xd1\xa6\x14\x31\x23\xbf\xca\xf3\x8d\x35\xc1\x67\x8d\xb5\x8d\x26\xec\xd9\x67\xca\x76\x39\x2b\x6b\x7e\xd8\x60\xc7\x7a\x57\x7e\xc0\x40\x8e\x51\x7f\x73\xab\xac\xf1\x02\x1c\xe9\x52\xf8\xb0\xd3\xe4\x5b\xa2\x20\x8e\x8e\xcf\x77\xce\x22\xa9\xda\x3c\xfa\x4c\x69\x3b\xd4\x1b\x8d\x8e\x52\x24\x7c\xc4\xe7\x5c\x73\xe5\xf3\x6e\x1f\x87\x3f\x53\x7e\x99\x2d\x2f\x2f\xb3\x37\xb9\xf2\x2f\xf4\x59\xc7\x26\x53\xde\x4f\x31\x53\xa4\xb8\x02\xe8\x1d\x65\x15\xfa\x16\xfe\x4a\x62\xfc\x54\xa8\x9e\x1a\x67\x07\x53\x4b\x65\xb5\x75\x2b\xa8\x34\xaa\xa7\x77\x07\x83\xbd\xf6\xf5\xb7\xdf\x6f\xaa\xb7\x57\x47\x7d\x2c\x88\xf4\xfc\x99\x56\xe0\x3b\xd4\xfa\x6c\x67\xaa\xca\x0a\xae\xad\xf1\x56\xa3\x5f\x7c\xb0\x06\x95\x5d\xbc\x1f\x14\xd7\xb8\x57\xd3\xe2\x3d\x57\xe4\x30\xb0\x35\xf0\xc1\x1a\xbb\xb8\xa1\x47\xfc\x38\xc0\x3d\x1a\x3f\x29\x7e\xe2\xe0\x83\x23\xec\xe0\x23\x39\x3c\xd9\xb8\xb6\x83\x63\x72\xf0\x1b\x8d\x0b\xe8\xac\xb1\xbe\x47\x45\x47\x14\x76\x4b\x6e\xa3\xed\xb8\x02\x1c\x82\x3d\xea\x47\xeb\x6a\x39\x3a\xec\x57\x60\xac\xeb\xf0\x04\xf8\xd8\x72\x20\x99\xfc\xac\x62\xad\xa6\x9d\xbf\x2f\xd2\x23\x8b\x64\xcb\x4a\x5b\xf5\xb4\x2f\x5f\x8f\x75\xcd\xa6\x59\xc1\x25\x2c\xaf\xfa\xe7\xd9\xfa\xdc\x38\x9b\x99\x91\x51\xe9\x4f\x6a\x7f\x52\x42\x36\x2d\x39\x0e\xb3\x8f\x22\xdf\xb3\x56\xe4\xd3\x80\x15\x95\xad\x77\x89\x4e\x83\x5b\x50\x1a\xbd\x2f\x85\xe6\xa6\x0d\xb2\xd2\x03\x41\xfc\x11\xe0\xac\xa6\x52\x18\xdc\x72\x93\x6a\x2a\x26\xda\x8b\x9a\x0f\x67\x0c\x6e\x53\xf6\x3d\xb9\x34\x56\xc8\x86\x9c\x58\x17\x08\x5c\x97\x42\xdb\xc6\xca\xa3\x7a\xdf\x99\xaf\xc5\x7c\xba\x72\x68\x6a\x19\xad\xc4\x71\x90\x70\xbd\xcf\xa8\x18\xf4\x6c\xe8\x22\x34\x68\xb9\x26\x69\x8d\xec\xa8\x96\xf1\x60\x6d\xc7\x88\x69\x36\xcf\x07\x7d\x14\x06\x9d\x10\x44\x7c\x9d\xad\x58\xd3\x21\xa8\x8f\x5e\x0c\x6e\xcf\x4f\xee\xd7\x78\x44\x59\x63\x40\x89\x2a\xf0\x16\x03\xf9\x7f\x75\x56\x0d\x21\x58\x13\x7b\x5d\x63\xef\x49\xac\x0b\x9e\xb7\x5e\xf2\x24\xd6\x1d\x99\xa1\xc8\x79\x7d\xc8\xb0\xc8\x6b\xde\x26\x12\x72\x83\xd3\xe2\xa4\xb2\x9e\x54\xea\x63\x63\x65\x8f\xb5\xac\x6c\x10\x29\x23\x36\x35\x3d\xcb\x0a\x4d\x2a\xf4\x17\x84\x9c\x90\x30\x67\x54\xb9\x75\xfc\xce\x62\xbb\x9c\x6d\x63\x2f\x44\xde\xc8\x04\x72\x50\xa3\x7b\x22\x23\xdf\x1e\xa8\x80\x5f\x87\x8a\x9c\xa1\x40\x1e\x7e\x1c\x42\x4b\x26\xb0\x4a\x8d\x50\xe4\xed\xf2\xe0\xef\x24\xb8\xb3\xe3\xde\xdb\x21\x7a\x0c\x78\x75\x1e\xd0\x6a\xf0\xcb\x37\x90\x1a\x4e\xac\x1f\x5a\xf6\x30\x04\xd6\x1c\x76\x30\xb2\xd6\xd0\x92\xee\x61\x67\x87\x38\x6b\x73\x58\x82\x91\x43\x1b\xb5\xee\x14\x98\xd2\x83\x8f\xe8\xd3\x26\x1a\xb8\xeb\xc9\xdc\xde\xc4\x17\x82\x21\x15\xe0\xab\xbb\xdb\x9b\xeb\xaf\x21\x4e\x6e\x06\xf7\xdc\x18\x60\x03\xc1\x42\x43\x01\x7c\x40\x17\xa8\xce\x8a\xbc\xbd\x3a\x64\x33\x73\xf2\xff\x52\x9b\x9b\x25\xd7\xb6\x61\x23\x26\x86\x62\x5b\x6a\x1b\x39\x4b\xcd\x71\xec\x95\x60\xa4\x46\xd7\x10\x8c\xb8\x25\x2f\x69\xb3\x89\x10\x27\x21\xd5\x62\x9a\xbb\x75\xc2\x79\x6b\x4e\x47\xe1\x05\xac\x99\xd0\xf3\x2e\x4a\x8b\xa8\x3d\xdc\x42\x70\xaf\x1c\xf7\xc1\xef\x2f\x22\x9f\x24\xf0\x4e\x9d\xdc\x0f\xb6\xa6\xec\xf1\xd3\x40\x6e\x97\xae\x86\x69\x29\xdf\x64\xcb\x6c\x99\xde\xfc\x8f\x5e\xac\x8b\x7c\x3a\xfa\x85\x97\xc7\x97\x37\xc5\x7f\xd9\xc6\xcb\xf5\xcc\x28\x21\x9f\x5e\x47\x45\x3e\xfd\x2d\xb8\xf8\x27\x00\x00\xff\xff\x21\x03\x34\xb3\x28\x08\x00\x00")

func templatesHomeTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesHomeTmpl,
		"templates/home.tmpl",
	)
}

func templatesHomeTmpl() (*asset, error) {
	bytes, err := templatesHomeTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/home.tmpl", size: 2088, mode: os.FileMode(420), modTime: time.Unix(1524578480, 0)}
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
	"templates/commandline.tmpl": templatesCommandlineTmpl,
	"templates/home.tmpl": templatesHomeTmpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"commandline.tmpl": &bintree{templatesCommandlineTmpl, map[string]*bintree{}},
		"home.tmpl": &bintree{templatesHomeTmpl, map[string]*bintree{}},
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

