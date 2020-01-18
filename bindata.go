// Code generated by go-bindata. DO NOT EDIT.
// sources:
// amanar_config_schema.json (14.643kB)

package amanar

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _amanar_config_schemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x5f\x6f\xdc\x36\x12\x7f\xdf\x4f\x31\xd8\x3b\xe0\xae\x07\x67\x17\xed\xa3\xdf\x0c\xa7\xe9\xa5\x70\x2f\x6e\x12\x17\x28\x0e\x3e\x97\x2b\x8e\x56\x8c\x29\x52\x21\xa9\xb5\x17\x41\xbe\xfb\x61\x86\x92\x56\xff\xf6\x8f\x9d\x38\x75\x9c\xea\xc1\xde\x95\x86\xe4\xe8\x37\x33\xbf\x19\x72\xc9\x0f\x13\x80\x69\x58\x17\x38\x3d\x86\xa9\x5d\xbc\xc3\x24\x4c\x8f\xe8\x9e\xc3\xf7\xa5\x72\x28\xa7\xc7\xf0\xdf\x09\x00\xc0\x54\xe4\xc2\x08\x77\x95\x58\x93\xaa\x65\xe9\x44\x50\xd6\x4c\x27\x00\x97\x2c\x5f\x88\x10\xd0\x99\x73\x67\x0b\x74\x41\xa1\x9f\x1e\xc3\x87\xd8\xf0\x7f\xb7\xcf\x66\xff\x6a\xbe\x02\x4c\x25\xfa\xc4\xa9\x82\x3b\x38\x86\xe9\xa6\x0d\x14\x0e\x53\x75\x8b\x12\x6e\x54\xc8\xe0\xf6\x19\xdc\x28\xad\x61\x81\x20\xb4\xb6\x37\x28\x41\x19\x08\x19\x42\xb0\x05\x68\x5c\xa1\x06\x9b\x42\xc8\x94\x07\x9f\x64\x98\x8b\x23\x58\x94\x21\x36\x32\x36\x50\xc3\xd2\xa3\x04\xeb\x20\xc9\x30\xb9\x46\x39\x83\xb7\x24\xcd\xdd\x79\x48\xad\x03\x61\x92\xcc\x3a\x0f\xc1\x92\xb8\x32\x1e\x5d\x40\xc9\x8f\x68\x24\x87\xa5\x47\x1a\xa5\x10\x2e\xf8\x38\x1c\xc2\xef\x27\xbf\x9c\x41\xaa\x34\x56\xfd\x51\x97\x46\x68\xbb\xb4\x25\xf7\x24\x8c\x84\x85\xe0\xa1\x0d\x3c\xb7\xc9\x35\x3a\x38\xb5\x79\x61\x3d\xfe\xc3\xc3\x8f\xb7\x01\x8d\x57\xd6\xc0\x0b\x85\x5a\x7a\xf8\x67\x16\x42\xe1\x8f\xe7\x73\x69\x13\x3f\x93\x2c\x3f\x4b\x6c\x3e\x4f\x62\x9b\xfa\xff\x33\x1a\x72\xfe\x37\xac\xdb\x3f\x4b\xb9\xfd\x77\xb3\x29\x43\xfb\x71\x02\xf0\x91\xad\x21\xa4\x54\x04\xaf\xd0\x1d\x83\xa4\x42\x7b\x8c\xe6\x1a\xb1\xd3\xa8\x81\x5b\x66\xab\xdd\x44\x38\x27\xd6\xec\x25\x7c\x5b\x05\xcc\x7d\x4b\x6e\xdc\xa1\xaa\x27\xd6\xe0\xab\xb4\xf1\xa9\x78\x7d\x68\x7d\xee\x79\xde\x74\x25\x4a\x1d\xae\x84\x94\x0e\xbd\x9f\x1e\x41\x75\xa3\xab\xe3\x65\xab\x83\x8f\x47\x07\xf6\x9c\x58\xe3\x83\x30\xa1\xdb\xba\xf9\x7c\xd9\x52\x7a\x04\xac\xea\x49\xd3\xc9\x71\x7f\xac\x9e\x8f\x9f\x30\xb6\xec\x27\x0b\x4d\xfe\x0b\x0e\x8d\x44\x07\x75\x0f\xfe\x08\x6e\x32\x95\x64\x20\x1c\x42\xa1\x30\x41\x76\x36\x65\x52\xeb\x72\x7e\x4d\x08\x99\x08\x20\x2d\x7b\xb6\xc4\x02\x0d\x7b\xd7\x6f\x04\x48\xd7\xad\xab\xc1\xa2\x47\x73\x00\x04\x0b\x35\x64\x08\xda\x26\x42\x83\x0f\x22\xa8\x04\x12\x87\x12\x4d\x50\x42\x7b\x0a\xae\xda\x71\xa8\x05\xf7\xfc\x4c\xa2\x53\x2b\x76\x64\xf4\xb3\xe9\x51\xf7\x2d\xb7\x1a\x7a\xab\xb1\x87\x66\x19\x98\x66\xf0\x94\x06\xc2\xbc\xd0\x22\xe0\x55\x21\x42\x36\x1d\x48\x5c\xf6\xee\x7c\x3c\xfa\x7c\x43\x1e\x30\x5a\xe7\xfb\x65\x0f\x85\x7d\xb1\xd8\x12\xdd\xea\x68\x5d\x85\x86\xcf\x46\x1c\xae\xf1\x2c\xf8\xc9\x42\xdd\x14\x7c\x70\xca\x2c\xa3\x2b\xd5\xdc\x1a\x3d\x11\x65\xdf\xba\xd0\xb6\x70\x6c\xd8\xc7\x62\x80\x73\xcf\x50\x87\x69\x4a\xa2\xcc\x9a\x1d\x55\x89\xeb\x3e\x9f\xa2\x93\x6d\xdf\x3a\xaf\xd0\xe3\x9b\x7d\x51\xfd\x36\x43\xa8\x64\xa3\xfe\x94\x24\x54\x52\x6a\xe1\x60\x15\xe3\x92\x83\xa8\x16\x42\x0f\xa5\x2f\x85\xd6\x6b\x90\x2a\x4d\xd1\x71\x9a\x89\x1f\xd1\x04\x40\xb3\x52\xce\x9a\x1c\x4d\xf0\x33\x78\x61\x1d\xe0\xad\xc8\x0b\x8d\x47\x70\x83\x90\x8b\x35\x64\x62\x85\x14\x8a\xb1\xfb\x66\x70\xea\xa6\x70\x56\x96\x09\x07\x2f\xa5\x1f\x61\x6c\xc8\xaa\x11\x7c\x10\x4b\x65\x96\x5b\xc3\xb7\xc2\xac\xf7\x34\x32\x0f\x3d\x2f\x9d\xea\x3f\xcc\x95\x39\x43\xb3\x64\x1b\x7f\xbf\x07\xce\x6d\x19\xa5\xa7\x45\x37\xaf\x54\x0f\x87\xd9\xa5\xd7\x6a\x94\x7a\xee\x14\x76\x63\x2e\xa9\x95\x0f\xc4\xbf\x11\x66\x67\x35\x7a\x46\x95\x5c\x35\x7e\xea\xbc\x14\x58\x6e\x1b\x0d\x61\xcb\x50\x94\x81\x1c\x42\x8a\x20\xc0\xdb\xd2\x11\x9b\x53\x49\x43\x0c\x3b\x70\x92\xb6\xd9\x87\xae\xbd\x9b\xa7\xa6\x8d\x1a\x0b\x8d\x7e\x2c\x2e\xa2\x09\x38\x1a\xb7\x3e\xa5\xf7\xeb\x47\xcc\xe5\x40\x91\x9d\xec\xd4\x1d\x69\xec\xf9\x78\xf8\x70\xec\x3b\x2c\x1c\x7a\x4a\x43\xcc\x4d\xc8\xc0\xf9\x60\x1d\xd6\x05\x5f\x3b\xcb\x29\x0f\x04\xc9\x4a\x68\x8a\x99\x60\xe1\xef\xbf\x9d\x5c\x9c\xbd\xbd\x3a\x3f\x79\xfb\xef\x5a\xfe\xf4\xec\x25\x24\x36\xcf\xc9\x52\x7f\x54\x46\x44\x21\xdb\xa2\x73\x4a\x7d\x7e\x5e\xdd\x79\xfd\xea\xec\xc7\x3f\xc6\x88\x65\x5f\x98\x54\x32\xdb\xe2\x21\x5e\x03\x9e\xec\x42\x7f\x38\x5c\x24\x3e\x84\xab\x40\x97\x2b\xef\xd9\x03\x99\x31\xa9\x88\xf0\xb6\x5c\x66\x8c\x4f\x03\xe0\x06\xd7\x3d\x50\x12\x1a\x5f\x11\x94\xdd\x20\xd8\x82\xe6\x1e\xc2\x80\xbb\x92\x06\x1c\x10\x13\x2c\xa3\x4c\x40\xad\xd5\xbb\x2b\x46\x3f\xb2\xc1\x56\xe9\x3d\x7c\xd8\x12\xeb\x73\x56\xac\xfd\x5e\xc6\xc1\x7e\x86\xcd\x60\x54\x02\x3a\x23\xf2\x86\xc2\xbc\xbf\xb1\x4e\xd6\x53\x9e\x24\x13\x66\x49\x33\xa3\x5f\xac\x0f\x24\x9b\x96\x9a\x79\xec\xb9\x08\xe2\x27\xa7\x0a\x6e\x45\xdd\xd1\x8c\xa6\xe1\xb1\x66\xa0\x0b\x1d\x54\x2e\x02\x8e\x5b\x1c\x76\x91\xf8\xdd\xac\x53\x49\xde\xd5\x46\x55\xb3\x3d\xe5\xde\x06\xd6\xea\x45\xaf\xca\x52\xc9\x1d\x7a\x54\xa2\x11\xe2\x2b\x2a\x56\xb6\xd4\xa6\xf5\x35\x20\xd4\x56\x4f\x07\xb8\xd1\xae\x31\x77\xb7\xd9\x41\xbc\xc1\x76\x1d\xa6\x49\x56\xf5\xdc\xb6\x2a\xc3\x94\x87\xb0\x2e\x54\xc2\xa5\x0b\xfd\xc5\xe8\x11\x6f\xa2\xf8\x8c\xa7\x14\xb3\xdb\x5c\x6f\x77\x82\x4a\x93\x03\x08\xa0\x25\xbd\x9b\x0a\x36\xd7\x08\x29\xf4\x30\xdb\x98\xf4\x3e\x60\x6d\x9c\xfd\xe2\xe5\xf3\x66\x6d\xa0\xee\x17\xd6\xb6\x84\x1b\x11\x19\xb4\x2c\x24\x05\x03\xfc\x6e\x4b\x48\x84\x81\x54\x19\x19\x97\x28\x16\x6b\x2e\xe7\x94\x69\x67\xb9\x01\x82\x11\xfa\x3f\x09\xc6\xad\xcf\xb6\x3d\x19\xbf\xbf\xc5\x1c\x1b\x26\x74\xa5\xe9\x56\x85\x0f\x47\x88\x91\xde\x7c\x5d\x8e\xb1\xbd\x44\x92\x50\xd5\xdc\x9e\xfb\xf6\x59\xcd\x95\xa6\x5b\xe2\x0d\x26\xbf\xed\xf7\xfa\x8a\xf8\xad\x55\x6c\x5e\xad\x84\x53\x94\x33\x77\xd3\xdc\xd0\x58\x57\xa9\xd5\x12\xdd\xb6\xaa\xb2\xd5\xb6\x09\xbd\xcc\xfa\xf0\xd0\xdc\xb8\x47\xd1\xbb\x06\xfe\x09\x48\xe5\x30\x09\xd6\xad\xc9\x15\x82\x88\x91\x2b\xb4\xde\xe5\x26\x55\x56\x8d\x91\x4e\x69\xf5\xa2\x9a\xf1\x51\x84\x87\xb8\x84\x39\x53\x12\xc5\xdc\x95\xe6\xb4\xeb\x61\xf0\x7a\xd8\x1d\xcd\xfa\x0c\xc6\x35\x9c\x05\x82\xcf\x84\x43\x09\x0b\x4c\xa9\x3a\x5e\x60\x62\x73\x52\x6a\xa5\xbc\x5a\xe8\xaa\x5c\x56\x34\x0f\xa1\xf7\x7e\x94\x6c\x3c\xea\x80\xf7\x21\xe5\x56\x47\x50\x77\x54\x17\xad\x03\xbb\x40\xc9\xcb\x6d\x71\x8d\xad\x43\xde\x89\x35\x06\x79\xde\xfc\x18\xd1\xea\x06\xd0\x7d\x60\xaa\xab\xbf\x4e\xf1\xd7\xe4\xb0\x8b\xd7\x67\x71\x7d\xc5\x1a\xbd\xe6\xf5\x42\xce\x5f\x12\x54\x5c\xe9\xa6\x61\x69\x22\x4c\x72\x15\xb6\xa3\xb8\xe7\x22\x24\x19\xd1\x2c\x2f\xc6\x33\x40\x4f\x34\x89\xbd\x2f\xd1\x29\x5b\xfa\x1f\xbe\x48\x3d\x7f\x60\xfa\xe2\x49\x1d\x73\x8b\x80\x5f\x2b\x05\xe1\x07\x78\xf3\xeb\x99\x0a\x1b\x5f\x7f\x22\x29\x6c\x63\x01\xff\x5e\xab\x7a\xad\xf1\x0e\x25\xe0\x03\xe7\xa1\x71\xfd\xee\x1a\xbb\xe7\x55\x81\x4e\x21\xd7\xb3\x23\xd9\x39\x32\x59\xcb\xd6\xec\x02\x1e\x54\xf0\x2c\xb7\xa7\x7a\x3f\x6d\x58\xcf\xcf\xa2\x96\x8f\x32\x59\x7c\x7a\xe9\x5e\x1a\xf5\xbe\x44\x50\x1c\x2a\xa9\xaa\x56\x44\x09\xd4\x06\xbb\x06\xd6\x56\xf9\x7e\x2a\x0c\xb1\x61\x6a\x4b\x23\xa9\x72\xd7\xd6\x5e\x53\xa2\xad\x38\xf0\xe0\xc0\x7a\x58\xec\x1e\x98\xea\x3c\xbe\x2f\x51\x5f\x15\xce\x3e\x26\xae\x23\xfb\xbd\x61\xcd\xe0\xdc\x59\x28\xb4\xf2\xe1\xa9\x54\xe7\x2d\xc4\xf9\xbd\x1e\x1b\xb5\x8d\xeb\xf7\x29\xd4\xc6\xdd\x6c\x08\xad\x65\xd8\xbb\x11\xda\x0b\xb1\xb2\x4e\x05\xf4\x33\xee\xf1\x5b\x23\xb3\x16\x6e\xf7\xa0\x33\x86\xec\x89\x92\x58\x61\x7d\x50\xc9\xa3\x62\xb0\x76\xb5\x76\x1e\xd5\x7b\xa2\xa5\x5a\x0d\xfe\xe3\x2c\xd4\xc6\xb4\x7b\x90\x32\xad\x36\xf2\x7d\x6b\xb4\x0d\xb9\xc9\xc5\xb7\xc6\x6c\x35\x76\x7f\x55\x69\xdb\xaa\xb4\x0c\xb5\x3e\x94\xde\xc6\x79\x4b\x54\xbb\x2d\x78\x89\x69\x89\x06\x5d\xbd\x54\x25\x80\xbb\x87\xd8\x2a\xfe\xba\x58\x2d\x86\x79\xc0\xdb\xc2\x56\xbb\xe2\xc6\x16\x05\x7c\x7b\xd9\x8c\xec\x61\xf0\xa6\xcd\x83\x3b\x38\xee\x40\x16\xfe\x9a\xa8\xb0\x5e\x87\x39\x70\xd5\xb5\x5e\xa9\x39\x50\x9c\xec\xf7\x25\x7e\x7f\x6a\xc6\xb9\xff\x6f\x4e\x19\x76\x7d\xca\x67\xb6\xd4\xb2\xeb\x79\xc1\x3e\xca\xe9\xe8\xd0\x86\xf7\x81\x81\x57\xe3\xaa\x9d\xa4\xa3\x8b\x69\x1c\x65\x15\x2c\x55\x08\xb1\x70\x3d\xfc\x63\x84\x66\xe8\xaf\x5f\x14\x9a\x7a\xf8\xa7\x49\xf1\xef\xbc\x35\x9f\xca\xf0\x3f\xbf\x79\xf5\x9f\x71\x9a\x6f\xb1\xf4\x96\xad\x02\xdf\x16\x53\x6f\x8a\x90\xa7\xc4\xb9\x1b\xfb\x7f\x45\x84\xdb\x32\xc5\xa7\xd1\x89\xf2\xad\xbd\x7c\xec\xdc\xf1\x2b\xc3\x53\x08\xd5\xd9\x2e\x2d\xc8\xfd\x87\xb5\xa8\xf2\x11\xc4\xe8\xa0\x4f\xb4\x9a\x6c\x36\xf0\xde\x8f\x6e\x5e\x28\xad\x7d\x6f\x57\x79\xdc\x1a\xeb\xec\x4a\x49\x94\x63\xdb\x91\x09\xdd\xce\xb6\x5f\x26\x1f\xa7\x4c\xf0\xd5\x21\x0c\x4f\xb6\x0a\x34\x81\x92\xb6\x0c\x71\xde\x74\xda\x1a\x24\x9a\x84\x66\x51\x1e\x03\xd9\x8f\x7f\xe2\xb3\xb1\x56\xc5\xdb\xc0\x3d\xd2\xb4\x2b\xad\x4e\x5c\xcc\x2e\xda\x3f\x8b\xcd\xce\x2b\xa6\xfb\x8e\xb7\xce\x89\x95\x50\x9a\xd2\xcd\x0c\x2e\x8c\x56\xd7\xc8\xdd\xc5\xfd\xb4\x2d\x58\x8e\x1a\xad\x3d\x37\xdb\x44\x93\xa0\x4a\x17\x45\x92\x41\x50\xd5\x10\xd5\xc6\x7d\x11\xa8\x0d\xbf\x0b\xef\x0c\x29\x1c\xae\x94\x2d\xbd\x5e\x03\xde\x2a\xcf\x5b\xfa\xae\x71\xcd\x4b\x08\xb9\x95\x2a\x5d\xcf\xda\x88\xc5\x81\xa4\x4d\x4a\x4a\x8b\x28\x21\x43\x87\xc7\x50\x9f\x1f\x59\x5a\x2d\xcc\x72\x66\xdd\x72\x5e\x5c\x2f\xe7\xf4\xe6\xf3\xba\xe9\xfc\x49\x70\xf8\xb6\x33\x05\xed\x6b\x0f\x4f\x1c\x98\x06\x00\xf6\x9f\x3b\x68\x5f\xfd\x53\x01\xed\x6b\x27\xc9\x3d\x80\xba\x9f\xa0\xe9\xbd\x72\x59\x2e\x6e\x3b\x56\xfc\xfe\xb3\xe4\xbd\xfd\x67\x09\x3a\xe2\x3b\x36\xb9\x51\xf8\xee\x3e\x5b\x90\x2a\x5e\x72\x51\x26\x1e\x3d\x3b\x68\x96\x5a\xab\xf9\x27\xe4\xc5\x9d\x67\x41\x76\xa1\x72\xb2\xf7\x38\xc8\xe7\x43\xe2\x4b\x64\xb1\x91\xbb\xc3\x7b\xfd\x3b\x3b\x4e\x83\x4c\xfa\x9f\xe2\x7f\x3e\x5a\x37\xf9\x38\x99\x4c\x26\xff\x0f\x00\x00\xff\xff\x41\x6b\x73\xc6\x33\x39\x00\x00")

func amanar_config_schemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_amanar_config_schemaJson,
		"amanar_config_schema.json",
	)
}

func amanar_config_schemaJson() (*asset, error) {
	bytes, err := amanar_config_schemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "amanar_config_schema.json", size: 14643, mode: os.FileMode(0644), modTime: time.Unix(1579364484, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe8, 0xab, 0x2b, 0xae, 0x6f, 0xb, 0xc9, 0xc6, 0x5c, 0xb3, 0x44, 0xe8, 0x5, 0xe2, 0x29, 0xbc, 0x3, 0x61, 0x1c, 0x80, 0xb8, 0xaf, 0x1f, 0xca, 0xc7, 0xb, 0x7b, 0x26, 0x48, 0x2, 0x51, 0x7f}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"amanar_config_schema.json": amanar_config_schemaJson,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"amanar_config_schema.json": &bintree{amanar_config_schemaJson, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
