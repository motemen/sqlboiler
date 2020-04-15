// Code generated by go-bindata. DO NOT EDIT.
// sources:
// override/templates/17_upsert.go.tpl (5.893kB)
// override/templates/singleton/mssql_upsert.go.tpl (1.357kB)
// override/templates_test/singleton/mssql_main_test.go.tpl (3.945kB)
// override/templates_test/singleton/mssql_suites_test.go.tpl (255B)
// override/templates_test/upsert.go.tpl (1.723kB)

package driver

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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
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

var _templates17_upsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\xd1\x6f\xdb\xb6\x13\x7e\x96\xfe\x8a\x6b\xf0\x43\x23\xfd\xe6\x28\xdb\x6b\x06\x3f\x24\x4d\xdb\x05\x6d\x32\xb7\x6e\x16\x60\x41\x10\xd0\xd2\xc9\x26\x42\x93\x2a\x45\xd9\xf1\x3c\xff\xef\xc3\x91\x94\x25\x3b\x76\xe2\xb4\xcd\xb0\x87\x22\x96\x78\xba\xfb\xf8\x7d\xc7\x3b\x5e\xe7\xf3\x03\xf8\x1f\x13\x9c\x95\x70\xd4\x85\xe4\x98\x7e\x61\x99\x7c\x61\x03\x81\xe0\xfe\x24\x17\x6c\x8c\x8b\x45\x68\x4d\xcb\x74\x84\x63\xe6\x96\xe9\x83\xc6\x02\xfe\x86\xa4\xdf\xac\xda\x0f\x78\x0e\xc9\x71\x96\xbd\x17\x6a\xc0\x04\x1c\x2c\x16\xe1\xe1\x21\x5c\x16\x25\x6a\xf3\x1e\x98\x31\x38\x2e\x4c\x09\x4c\x02\x97\xf4\xae\x03\x4c\x66\x90\x29\xb4\xef\xaa\x22\x63\x06\x41\x69\xe0\x43\xa9\x34\x82\x92\x90\x2a\x99\x0b\x9e\x9a\x24\xcc\x2b\x99\x42\xa4\xe0\xff\xf3\xb9\xc3\x9f\x5c\x16\x7d\x2e\x87\x95\x60\x7a\xb1\x88\xeb\x28\x91\x05\x21\x95\x81\xe4\x42\xbd\x51\xd2\xe0\xbd\x59\x2c\x52\x73\x4f\xae\xe8\x21\xf1\x2f\x3b\x30\x9f\xa3\xcc\x08\xa4\x8f\xfc\x46\x89\x6a\x2c\xcb\x8e\x07\xe7\x1f\x61\xa0\xb8\x48\xfc\x43\x0c\xa8\xb5\xd2\x30\x0f\x03\x8d\xa6\xd2\x12\x54\xe2\x02\xbb\xb8\xed\x98\xf6\xbb\xf7\x68\x4e\x4f\xa2\x78\x3e\x47\x51\xa2\xc5\xd1\x81\x7a\xc1\x5b\xfa\x75\x99\x2d\x16\x9d\x47\x91\xc4\xe1\x22\x0c\x97\xa0\x43\x47\x37\x11\xd8\xa2\x9c\x7e\xf6\x98\xe4\xe9\x1a\xf9\xbd\xef\x63\x1f\xac\xcf\x92\xde\x59\x02\x76\x96\xa3\xf7\xd2\x7a\xcc\xc3\x80\xe7\x04\x8a\xb2\xf3\xdf\x14\xe3\x57\x1b\xf4\x55\x17\x24\x17\x84\x22\x28\x88\xa2\xc8\xfa\xbb\xd2\xac\x78\xab\x75\x84\x5a\xc7\x71\x18\x2c\x36\x09\xb7\x45\xa9\x4d\x42\x41\x55\x72\x39\xa4\x67\xbc\xc7\xb4\x32\x4a\x3f\xe7\xe0\xb4\x5c\x17\xdf\xa6\x62\xef\x21\x9f\x04\xc4\x71\xf7\xd6\x43\x6a\xb1\xfa\x50\xda\xc6\xdc\xbf\x6a\x7d\xf5\x34\xd7\xbb\x4b\xbe\x21\xcf\xda\x79\x45\x30\x5e\x4e\xd6\x25\xd1\x3f\x5c\xc2\xdd\x64\xfa\x6f\xa9\xb4\x2c\x94\x3c\x07\x05\xdd\x86\x50\x5f\x38\xed\x7a\x99\x5c\xe0\x34\xda\x9b\xcf\x93\xde\xdd\xd0\xb5\x9d\x23\x90\x0a\xe6\xf3\x95\x56\x04\x85\x56\x13\x9e\x61\x06\xb9\xd2\x50\xd9\xdd\xee\x59\x05\xc2\x80\xba\x14\xb1\x2d\x88\xbf\x3d\xc3\xc7\x58\x1a\x36\x2e\x6e\x9d\xd5\xed\x08\x45\x81\x7a\x0f\x12\x58\x38\xeb\x26\x4b\x7e\x53\xea\xae\xb4\xd2\xad\xe4\x53\xa6\x4e\x30\x57\x1a\x1d\xa9\xd6\x68\xe7\xe4\x7a\x98\x3e\xcd\x6e\x09\xae\x45\x6b\xb9\x0c\xc3\x40\xfe\x75\x8a\x39\xab\x84\xb1\xad\xf8\x6b\x85\x9a\x63\x99\x5c\x28\xf9\x27\x6a\xe5\x97\xfa\x48\xb2\x7a\xd1\x4f\xd5\x54\x36\xb2\x7b\xa6\xaf\xb8\x19\x79\xe3\x0e\xa8\x38\x0c\x83\xc3\x43\x38\xa9\xb8\xc8\x20\x65\xe9\x08\xe1\x0e\x67\xc0\xe5\x81\xe0\x12\xa1\x1a\x0a\x2e\x66\x70\x00\xe3\x59\xf9\x55\xc0\xa4\x84\x82\xfe\x16\x5a\x0d\x04\x8e\xcb\x30\x18\x54\x39\x81\x29\x8d\x1e\x33\x39\x14\x48\xa5\xf1\xa4\xca\x73\xd4\x51\x6c\x57\x93\x2b\xcd\x0d\xf6\x8d\xe6\x72\x18\x95\x46\xa7\x4a\x4e\x92\x33\xa3\x58\xb4\x92\x1b\xc9\x07\x2e\x33\x3a\x24\x24\xd8\x6d\x07\x52\xf2\xaa\x99\x1c\xe2\x6a\x0e\x51\xbe\x94\x96\xa8\x75\xdf\xa9\xd5\xb7\x79\x7d\x32\x33\x18\xed\x27\xfb\x4f\xc1\x58\xc9\xc9\x47\x60\xac\xda\x7d\x0b\x8c\x87\x3e\x5b\x8a\x3e\xe2\x8b\x04\x39\xea\x02\xad\xfa\x85\x38\x0c\x1a\xc6\x7b\x55\xcd\xf8\xa0\xca\x63\x9b\xb3\x1b\xf5\x77\xf9\xf9\x86\x34\x3e\xaf\x4c\xf2\xf9\xa3\x4a\xef\xc8\x93\x55\xbd\xe3\xc4\xcf\x28\xd0\xd3\xdf\x5f\xdf\xe1\xec\x66\xe7\x40\x97\x52\xb8\x50\x61\x30\x61\xda\x26\xbc\x3d\xcc\xa1\x3d\x47\xaf\x7c\x60\x22\xa0\xbe\x67\x68\x34\x04\x64\x95\xf2\xb3\xd6\x13\xa5\x79\x18\x04\xdb\x10\x1c\x0b\x51\xd7\x9c\x47\xac\x36\x1c\x88\xdd\xac\x55\x65\xda\x1f\x34\x2a\xd2\x63\xbc\xdc\x07\xb4\xcf\x45\x9f\xae\x0c\xe3\x42\xe0\x18\xa5\x89\xea\x8d\x3e\x1d\xeb\xb8\x32\x8a\x5c\x52\xf2\xf0\x0e\x4c\xd6\x13\xd2\xf2\x46\x3c\x36\xa1\xa8\xe0\x30\x2e\xcb\x63\x39\xdb\x56\x0b\x7a\x9a\x8f\x99\x9e\x7d\xc0\xd9\xb2\x36\x4f\x62\x78\xfd\xfa\x79\x5e\x36\x55\x94\x49\xec\x10\x35\x1c\xb0\xa2\x40\x99\xf9\x2d\x5f\x1f\xf1\x9b\xba\x0f\x5c\xf3\x9f\x7e\x39\xba\x49\x92\x84\xf6\x47\x89\x6e\xff\xf1\x1c\x04\x4a\x6f\x1e\x53\x23\xf8\xd9\x79\x7c\xb2\x0f\x54\xd2\x4e\x1d\x46\xf9\x8a\xbf\xde\x15\x3a\x90\xaa\x4a\x64\xb6\x2e\x0f\x6c\xc1\xf3\x18\x53\xbb\x0f\x10\xbc\xb4\x5d\xc2\xb6\x09\x0a\xb7\x2e\xe0\x39\xea\x21\x46\x1a\x9f\x25\xdc\xf7\xfa\xf1\xcc\xd2\xe9\x09\x7c\xd7\x3f\xea\xae\x15\xc5\xcb\xd6\xd3\x0f\x39\x1a\x0f\xf3\xc3\x67\xb6\x47\xb0\x3d\xb3\x9d\xc1\xee\x04\x35\x8a\xbb\x2f\x5f\x58\x71\x8f\x7f\xa3\xe2\xb6\x10\x25\xd4\x57\x67\xd0\x75\xf6\xae\x94\x7d\xa2\x57\xe7\xfd\xfe\xa7\x8f\x51\xc6\x99\xc0\xd4\x74\x60\x6f\x2d\xd4\xde\xd6\x2d\x6f\x38\x6b\x35\x49\xad\x7a\x67\x99\x98\x8e\xb8\x41\x02\x45\x12\x8f\xd9\x1d\x46\xd7\x37\xa5\x2d\xf9\x1d\x4b\xd1\xae\x11\xa8\x83\x05\xa9\x2a\x66\xd1\xd2\xe3\xee\xf0\xe2\x15\x20\xcb\xf3\xdb\xf2\xe4\xe0\xfb\x83\xfb\xb8\xa9\xdb\xa1\x35\x5d\x32\x3c\x61\xa2\xc2\x73\x56\x14\x76\x5f\xd4\x0e\x9a\xdb\xcc\x09\x97\x99\x5f\xda\xb6\xdb\x2f\xb3\x62\x7b\x7e\x2d\xdd\x2e\x31\xc4\x2e\xc3\xd6\xae\x59\x2b\xf7\xac\x76\xdd\x21\x29\xc8\xd0\xa7\xa0\x43\xac\xd1\xbc\x34\x5e\x9b\x02\xc1\x46\xa8\xab\x58\xeb\x42\xb9\xb0\xed\x54\x54\xb6\x1c\x68\xcc\x29\x2d\x93\x33\x99\x71\x8d\xa9\x89\xea\x17\x7f\x90\xc5\xef\x79\xa4\x28\x25\x26\x4c\xac\x5c\x1d\xed\x62\xf9\x4e\xab\x71\xbd\x05\xeb\xd0\xdf\x05\x56\x74\x8a\x5d\xef\x76\x48\x4a\xb8\xbe\xe1\xd2\xa0\xce\x59\x8a\x73\x77\x1d\x26\xee\xd6\xc9\x6a\x11\x59\x7f\xd8\x04\xef\x19\xbd\x3d\x74\xcb\x47\x7d\x6b\x5f\x19\x55\x96\xb7\x70\x3b\x43\x9c\xe2\xa0\x1a\x9e\xab\x0c\x6d\xa8\x7c\x6c\x92\x77\x85\xe6\xd2\x08\x19\x35\xeb\xf6\x62\xa5\xeb\x00\xf6\x94\xc7\x4f\x5b\x13\x65\xb1\xbf\x89\xd3\x24\xb4\x1a\xf8\xac\xb4\xc6\x51\x6a\xee\x5d\xdb\x9b\xda\xcf\xec\x6d\x6d\xcd\x15\x6d\xd5\xda\xad\xc7\x9c\xee\x80\x6b\xba\x09\x4d\x3d\x46\xee\xc0\xfe\x46\xf6\x02\x97\xc9\x34\x86\x24\xb6\xc2\x7d\x56\xd3\xa8\x85\xc2\x85\xa3\xa3\x9b\xf4\x53\x66\x4f\x06\x49\xe8\x8f\x7d\x9b\x8e\x4d\x9e\x7c\xa8\xc8\x0e\x3d\xcf\xf1\xea\xb7\xb5\x3c\x09\xdd\x2e\x94\x5f\x45\xf2\x56\xeb\x0b\xf5\x59\x4d\xdd\x4d\xd9\x47\xa4\x23\x72\x78\x08\xb6\x34\xdb\xd1\x58\xee\x1b\x9f\xa3\xc0\xe4\xcc\x8c\x68\x86\x9e\x8e\x50\x82\x19\xa1\xc6\xfd\x92\x66\x45\x57\xbd\xfc\x21\x02\xbb\x8b\xed\x1c\xdd\xd6\x07\xde\x6e\x8e\xe6\xdb\xcd\x14\xad\x33\xf2\xf0\xbb\xa7\x09\x59\xdd\x7f\x33\x65\x6e\x9c\x0e\xa9\x23\x5e\x69\x56\x44\xa8\xb5\xeb\x46\xcf\xe8\x8b\xf5\x2c\xbc\x76\xfd\xde\xed\x3e\x5f\xcf\x0d\x3b\x98\xdb\x39\x01\xba\x6e\xbb\x3b\x07\x58\xce\x0b\xc1\x23\x13\xf8\xf2\xbf\x53\x33\x75\x9c\x1b\xd4\xdf\x34\x7d\xfb\xf9\x7a\x29\x9b\x77\x2a\xb9\x68\x4f\xde\x8b\xf0\x9f\x00\x00\x00\xff\xff\x11\x7a\xc8\x7b\x05\x17\x00\x00")

func templates17_upsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates17_upsertGoTpl,
		"templates/17_upsert.go.tpl",
	)
}

func templates17_upsertGoTpl() (*asset, error) {
	bytes, err := templates17_upsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/17_upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf6, 0x71, 0xfe, 0x9a, 0xba, 0x83, 0x8e, 0xe5, 0x7c, 0x3, 0x59, 0x5e, 0x4d, 0xd3, 0x20, 0x8, 0x2f, 0xc9, 0x7e, 0x68, 0xc6, 0x33, 0x7f, 0xfd, 0x58, 0xe6, 0xaf, 0x87, 0xfb, 0x33, 0xb8, 0x56}}
	return a, nil
}

var _templatesSingletonMssql_upsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x53\x4d\x6f\xda\x40\x10\x3d\x7b\x7f\xc5\xd4\x52\x24\xaf\xb2\x72\x9a\x6b\x23\x2a\xd1\xe0\x26\x54\xc4\x40\x6c\xda\x03\xe1\xb0\xe0\x31\x59\xc9\x2c\x68\x3f\x50\xa3\x2a\xff\xbd\x1a\xdb\x94\xcf\x4a\x55\x2f\xe0\x9d\x8f\xa7\x37\xef\xcd\xdc\xdc\xc0\xdc\xab\xaa\x98\x6c\x2c\x1a\x37\xf6\x68\xde\x9e\xb2\x6c\x3c\x68\xa2\x16\x24\xd0\xc3\x3a\xe9\x70\x85\xda\x81\x75\x46\xe9\x25\x78\x4b\xbf\xee\x15\xc1\xd7\x8d\x3d\xe9\x24\x6c\xcc\x7a\xab\x0a\x2c\x62\x56\x7a\xbd\xb8\x8c\x1b\x15\x4a\x42\x61\xd4\x16\x8d\x8d\x7b\x4a\x56\xb8\x70\x02\x9c\x9c\x57\x98\xca\x15\xb6\xf8\x02\x36\x46\xad\xa4\x79\x13\xe0\x37\x85\x74\x28\x40\x69\x02\x82\xe9\x6c\x57\xb1\xf6\x6e\xe3\xf7\x01\xbe\xa3\xf6\x8b\x05\x6d\x6d\x87\x42\x2b\xa9\x97\x15\xc6\xfd\x02\xb5\x1b\xfb\xb5\xc3\xac\x52\x0b\x24\x1a\xf1\x60\x2c\x80\xfe\x9f\xc7\x3b\x78\xce\x58\x30\xf7\x25\x7c\x3a\x6c\x7d\x40\xf7\xc5\x97\x25\x9a\x88\xb3\xa0\xc0\x12\xcd\x41\x72\xe4\x77\xc9\xb9\x2f\xa9\xdd\x3a\x69\x5c\x5f\x17\xf8\x93\x50\x6e\x19\x0b\xca\x95\x8b\xbf\x6e\x8c\xd2\xae\xa4\x22\x01\xe1\x53\xf2\xfc\x90\x40\x3f\xcd\x87\x70\x65\x41\x5a\x98\xba\xd9\x8b\x0e\x0f\x74\xe0\x97\xda\x26\x59\x3f\x7d\x80\x28\x4b\x06\xc9\x7d\x0e\x57\x96\xd7\xad\x76\x06\xd1\xf4\xca\xce\x38\x21\xb0\x20\x38\xe0\x56\xc9\x05\xbe\xae\xab\x02\x8d\xad\x07\x9e\x58\xac\x99\x1d\x26\x04\x54\xa8\xa3\x56\x6e\x2e\x60\xcf\x5f\xc0\x2d\x6f\x01\x95\x5e\xda\xf8\xdb\x5a\xfd\x29\x14\xad\xda\x51\xa3\x1f\xbf\x0e\x45\x78\x7d\x10\x1a\x8c\x39\x3f\x9a\xa1\x1d\x61\x98\x42\x14\x52\x62\x6d\x40\x09\xd8\x92\x46\x46\xea\x25\xee\x0c\x27\xfb\x02\x55\x82\x82\x0f\x1d\xf8\x58\xbf\xce\x51\xa0\x9b\xf6\x80\x60\x82\x77\x16\x5c\x10\x6a\x6a\x67\x31\x49\x02\x1d\x52\xb6\xfe\x0c\x05\x6c\x05\x6c\x39\xa3\x96\x33\x40\xd2\xee\xc4\xbc\xeb\xce\x91\x30\xec\x42\xd7\x8f\xc7\x24\x85\xa7\x6e\x7e\xff\x98\xf4\x20\xa7\x47\x78\xd9\xb7\x51\xaf\x9b\x27\x90\x25\x64\x5a\xed\xf3\xde\xa3\x0c\xdd\x48\x1a\xb9\x22\xd3\x6d\x74\xac\xe0\xa9\xc8\xc7\xe6\x34\x87\xc1\x2f\xd3\x6e\x93\x7f\x65\x9d\x0e\xf3\x7f\x61\xde\x4f\xb3\xe4\x39\x87\x88\x76\xed\x7b\x77\x30\x49\xb2\xfa\x3b\x3c\x5b\x8b\xe6\x7c\x04\x84\x24\xe6\x7f\x6f\x61\x7b\x84\xa7\x4b\x48\x63\xa8\xb2\xae\x68\x8e\x9e\xc3\xe7\x76\x37\xce\x29\xbf\xe8\xe1\x24\x1f\x4d\x72\x68\xb8\x27\xbd\xda\xfe\xbb\x70\x27\x66\x4b\xb8\x01\x12\x10\xce\xc4\xbe\x30\xa4\x9d\x7d\x07\xac\x2c\x9e\xa0\xb7\xe0\x77\x61\xbd\x40\x2c\x30\xe8\xbc\xd1\x30\xf7\x65\x9c\x35\x1e\x71\xf6\xce\x7e\x07\x00\x00\xff\xff\xa0\xc3\x9b\xd6\x4d\x05\x00\x00")

func templatesSingletonMssql_upsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSingletonMssql_upsertGoTpl,
		"templates/singleton/mssql_upsert.go.tpl",
	)
}

func templatesSingletonMssql_upsertGoTpl() (*asset, error) {
	bytes, err := templatesSingletonMssql_upsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/singleton/mssql_upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa6, 0x69, 0x72, 0x38, 0xff, 0x12, 0xac, 0xfa, 0x17, 0xbb, 0xc6, 0xa8, 0xc0, 0x42, 0x97, 0xf5, 0x69, 0xf1, 0x9d, 0x12, 0xd9, 0x70, 0x6f, 0x7a, 0xe3, 0x4, 0xf9, 0x1a, 0xc7, 0xd0, 0xf4, 0x4}}
	return a, nil
}

var _templates_testSingletonMssql_main_testGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x57\xed\x53\xdb\x36\x18\xff\x6c\xff\x15\x4f\x73\xd7\xd6\x66\x9e\x68\xd7\xdd\x3e\xd0\xcb\xf5\xf2\x62\x5a\xae\x24\x81\x38\x5b\xb7\xa3\x0c\x94\x58\x06\x1d\xb6\x64\x24\x19\x9a\x31\xfe\xf7\x9d\x24\xdb\xb1\xd3\x24\x85\x4f\xe3\x4b\xd0\xa3\xe7\xf5\xf7\xbc\xc9\x77\x58\x80\xb8\xfa\x36\x8a\xa2\xd3\xe3\x1b\xb2\x84\x2e\x08\x72\x45\xbe\xe5\x68\x54\x48\x35\xe0\x59\x4e\x53\xe2\x5d\x7a\x1f\x32\xff\xef\xde\xf1\x2c\x9c\xc2\xac\xd7\x3f\x0e\x01\xed\xf5\x86\xc3\xaf\xf2\xa7\xc1\x64\x1c\xcd\xa6\xbd\xa3\xf1\x0c\xd0\x1e\x1c\x4e\xa6\xe1\xd1\xc7\x31\x7c\x0e\xff\x42\x7b\x1f\xd0\xde\x57\xf6\x61\x1a\x1e\x86\xd3\x70\x3c\x08\x23\xb4\x77\xe9\xbb\xae\x5a\xe6\x04\x32\x29\x6f\xd3\x19\x91\x8a\x08\x90\x4a\x14\x0b\x05\x0f\xae\x13\xcf\x07\x9c\x31\xd0\x7f\x7b\xf2\x36\x45\xc3\xbe\xa6\x8d\x71\x46\x0c\x4d\x2a\x41\xd9\x95\xeb\x5c\x73\xa9\x00\x5a\xa4\x42\x12\xb1\x46\xca\xb1\x94\x6b\x24\x29\xd3\x8c\xc7\xa4\xc5\xc5\x45\xa5\x8b\x32\xe5\x3a\x8a\x48\x35\xec\x1b\x93\xb5\xd4\x0d\xcd\xa3\xd3\xe3\x41\x16\xc3\x9c\xf3\xd4\x7d\x74\xdd\xa4\x60\x0b\xa0\x8c\x2a\xcf\xb7\x7e\x8f\x30\x65\xd0\x85\x57\x8d\xb8\x1e\x1e\x6b\x4e\x2f\x83\xbd\xc6\x8d\x0f\x92\xa8\x22\xf7\x7c\x20\x42\x70\xa1\x35\xe8\x24\x10\x21\x2c\xc1\x75\x9d\x3b\x9a\x13\x81\x22\xa2\x86\x24\xc1\x45\xaa\xbc\x8e\x91\x47\x72\x71\x4d\x32\xdc\x09\xa0\x13\xcf\x79\xc7\xdf\xc1\x68\x43\xd5\x9c\x4a\x14\x64\x17\xab\x86\xa0\x13\xc0\xdb\x5f\xdf\xbd\xf3\x5d\xd7\xc9\x50\x09\x79\x17\xac\xc4\x47\xa2\x22\x03\x45\x25\x10\xcf\x19\xce\x8c\xca\x0c\x99\x5c\x6c\xe5\xd4\xb7\x96\xcf\x24\x68\x2b\x9f\xbe\xb5\x7c\x26\x6b\x5b\xf9\xf4\x6d\xc9\xa7\xf3\xd6\xe0\x3b\x62\xed\x78\x0c\x53\x95\xef\xad\xfa\x2a\x94\x0c\x77\x23\xf5\x5b\x05\x34\x4f\x33\xfc\x46\x6d\x34\x64\xfa\x9c\xa7\xb5\x89\x1b\x9a\xcb\xdb\x74\x91\xc5\x1d\x8d\xae\x4e\x72\x17\xee\x70\x8a\x51\x9f\x5c\x51\xf6\x07\x4e\x69\x8c\x15\xe5\xcc\xf3\x51\x79\x20\x9e\xeb\x38\x86\xc5\x1a\x1f\x73\x15\x66\xb9\x5a\x7a\x3b\xd1\x0b\xa0\x7d\x7c\x9e\x0e\x9b\xa9\x5a\x47\x79\xac\x74\x8c\xb9\xf2\xcc\x3f\xe1\x6d\x81\x53\xe9\x6d\x87\x3d\x80\x37\xb5\x12\x4b\x79\xae\x27\x15\xbc\xb5\x9a\x9a\xf0\x3c\x3d\x75\x6e\x6b\x45\x2b\x8a\xeb\xf8\x68\x70\x4d\x16\x37\x9e\xce\x09\x4d\x4c\xef\xbd\xe8\x02\xa3\xa9\xee\x46\x47\x10\x55\x08\xa6\xa9\xae\xf3\xe8\xba\xce\xfe\x3e\x0c\x04\xc1\x8a\x00\x06\x81\x59\xcc\x33\xfa\x0f\x89\x21\x9e\x83\x76\x0d\x19\x15\x29\x61\x5e\xb3\x88\x7c\xe8\x76\xe1\x8d\x51\xb7\x56\x5b\xb5\x06\x14\x29\x3c\x4f\x89\xbd\xf0\xaa\xc6\xf3\xad\x4d\x9a\xc0\x8b\x56\x81\x69\x4d\xa5\xab\x5d\xc8\x50\x2c\x78\x3e\x33\x6a\x3d\xff\xfd\x7a\x00\xad\x08\x9c\xc7\xb6\xe4\xc2\x84\xf2\x64\x59\xd7\x71\xac\x84\x76\xe2\xa0\x0b\xe4\x1b\x59\xa0\x01\xcf\x32\xcc\x62\xaf\x53\xd6\x76\x00\x9d\x9f\xa3\x4e\x00\x76\x22\xe8\xd3\xef\xe6\xa4\x8b\x51\x9f\x4e\xcc\x49\xf7\xaf\x3e\xc5\xe6\xd4\xc0\x4a\x1b\x49\x02\xe3\xc9\x41\x17\xb8\x44\x93\x9c\x30\xaf\x63\xe0\x91\x17\x76\xea\x21\x79\x9b\xea\xae\xdb\x90\xaf\x86\xcb\x5c\x48\xf4\x45\xe0\xdc\x23\x42\x1b\x4e\x30\x4d\x49\x0c\x8a\x03\xcf\x09\x83\xef\x14\x42\x42\x53\xd3\xcb\x36\xd0\x98\x24\x44\x80\x1e\xda\x7a\xb2\xc3\x05\x74\x21\x41\x83\x94\x4b\xe2\xf9\xf0\x68\xaa\xc5\x91\x2a\x2e\xfd\x7c\x35\x5f\x2a\x22\x51\xbf\x48\x12\x33\xef\x1b\x40\xa1\x48\xc5\x66\x25\x30\x72\x7f\xf8\x99\x2c\x87\x44\x2a\xc1\x97\x44\x78\x8d\x5d\x1b\x40\xe2\xaf\x0b\xd9\x24\x59\x1b\x6e\x33\x6f\x4d\x2e\x2c\xd4\xee\xc4\x6d\x45\x41\x6a\x59\xb0\x49\x83\x85\x4d\xe2\x2a\xfc\x0d\xc6\xbe\x60\xba\xd1\x56\x92\x29\x74\x22\x28\x53\x29\xd3\x46\xfc\x75\x9a\x8d\xa0\xec\x55\xcf\xf7\x9f\xe8\xdf\x3d\xa6\x0a\x12\x2e\x36\xbb\x68\xbc\x2c\xb5\x30\x9a\xee\x58\xb0\x32\x1d\xf1\x98\x78\x66\xfc\xdb\x45\xee\x97\xbf\xda\x7d\x79\x4f\xd5\xe2\x1a\xcc\xed\x83\xeb\x2c\xb0\x24\xe5\x9e\x3c\x58\x75\xbf\x25\x54\xb7\x09\x4e\x65\xfb\xda\x52\x5c\x5d\x33\x7a\x9d\x36\xaf\x62\x2a\x75\xa1\x75\xb4\xc3\x5b\x7d\x6c\xb7\xe1\xea\x2d\xa0\xab\xf2\xa0\x0b\x1a\xcc\x28\xd7\x68\x26\xde\xa5\xeb\x0c\xa6\x61\x6f\x16\xc2\xb0\x37\xeb\xf5\x7b\x51\x08\x2f\xe5\x7b\xd7\xf9\x38\x71\x1d\xfb\x28\x5b\xd1\xcf\xde\x9e\x4b\xd7\x89\xc2\x19\x4c\xc3\xde\xf0\x62\x30\x19\x8d\x8e\x66\xb3\x70\x78\x11\x8d\x7b\x27\xd1\xa7\xc9\x0c\x26\x63\x23\x7a\xb9\xde\x83\x95\xfb\x19\x12\x05\x1b\x64\xb1\x27\x6f\xd3\x00\x9e\xdf\xe1\xfe\xf6\x98\x9b\x43\x6b\x15\xf1\xfe\x3e\x44\x94\x2d\x08\x8c\x22\x88\x4e\x8f\xe1\x97\x37\x6f\x7f\x03\xaa\x60\x81\x19\xcc\x09\xc4\x9c\x11\xb8\xa7\xea\xda\x70\x0e\xa7\x93\x93\x55\xb8\x67\x70\x74\x08\xe1\x9f\x47\xd1\x2c\x82\x73\x78\x80\x18\x2b\x3c\xc7\x92\x5c\xe8\xc1\x0c\xff\xae\xce\x92\xe1\x5c\x5e\x73\x65\x2f\x1e\xe1\x0c\x02\x84\x10\x83\x73\x38\x7b\x7f\xbe\x0d\xf4\x5a\xb7\x17\x85\xc7\xe1\x60\x66\xc6\x3d\x1c\x4e\x27\x23\x90\x4b\x89\x2a\xe5\x12\x5c\xc7\xf9\xf2\x29\x9c\x86\x96\xa1\x0b\xaf\x5f\xca\xd7\xba\x64\xdb\xce\xbe\x94\x1b\x70\xff\x1f\xb2\xa0\x08\x16\x31\xbf\x67\xcd\x1c\xd0\x44\xef\x14\xfb\x00\x6f\xf4\x79\x45\xab\x86\xe0\x8f\x77\xd3\xc1\xf3\x97\xd3\x53\xbb\xba\x02\x44\x8f\xd6\xa0\x1a\x0d\x65\x5b\x07\x80\xc5\x95\x04\x84\x50\xd5\xee\x75\x68\x8b\x0d\x7b\xab\x14\xb6\x52\x08\x21\xdf\xb0\xd5\x53\xdb\xea\x90\x68\x4c\xee\xa7\x04\xc7\x44\x58\xa3\x7a\xfe\x4b\x15\xf3\x42\x6d\x1c\xff\x3b\x36\x43\xa9\x5c\x4b\x9a\xe9\xce\x0b\x55\x13\x5b\x23\xbf\x01\xa3\xbe\x9e\x16\x6c\x03\x82\xcd\x41\x5b\x0d\x4f\x51\x30\x46\xd9\xd5\x41\xa7\x46\xc6\x06\xe7\xbb\xdf\x0d\x66\x5e\xa8\xd6\x60\xfe\xc1\xdc\x5e\x7f\x0d\x3d\x25\x55\x0b\xce\x74\x79\x79\xe5\x77\x5c\x60\xb3\xe1\xef\xa8\xb4\xba\xec\xed\x55\x60\xf4\x1b\x7b\xed\x8f\x23\x67\xc5\x51\x02\x77\x9b\x96\xcf\x05\xe3\x41\x27\x80\x58\xd0\x3b\x22\x90\x59\xb3\xfd\x82\xa6\xf1\x69\x41\xc4\xb2\x0c\xa9\xea\x95\xea\x35\xb2\xde\x8b\xb6\xaf\xec\x17\x86\xfe\x2d\x5f\x8d\x1a\x89\xad\x0f\x45\x46\xd3\xe0\x3b\x7c\xda\x91\x3c\xba\xff\x05\x00\x00\xff\xff\x68\x86\x33\x94\x69\x0f\x00\x00")

func templates_testSingletonMssql_main_testGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonMssql_main_testGoTpl,
		"templates_test/singleton/mssql_main_test.go.tpl",
	)
}

func templates_testSingletonMssql_main_testGoTpl() (*asset, error) {
	bytes, err := templates_testSingletonMssql_main_testGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/mssql_main_test.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe3, 0x48, 0x72, 0x18, 0x21, 0xc4, 0x79, 0x4b, 0x1c, 0xde, 0x6, 0x5c, 0x37, 0xfb, 0x48, 0xf1, 0x61, 0xfd, 0x28, 0x9e, 0x5e, 0x56, 0x88, 0xbf, 0x7b, 0xc0, 0xb4, 0x66, 0x57, 0x69, 0xe0, 0x87}}
	return a, nil
}

var _templates_testSingletonMssql_suites_testGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xc1\x0a\xc2\x30\x10\x44\xef\xfd\x8a\xa5\xe4\xd0\x4a\x9b\x0f\x10\x3c\x78\xd4\x83\x88\xb4\x1f\x10\xed\xb6\x04\xe2\x5a\xba\x5b\x10\x42\xfe\x5d\xd2\x46\xe9\xc1\xdb\x0c\x6f\x32\x99\xed\x67\x7a\x40\x83\x2c\xed\xc8\x38\x49\x21\xb0\x13\x64\xb1\x34\xe8\xa6\x04\x9f\x01\x78\x5f\xc3\x64\x68\x40\x50\x96\x3a\x7c\x57\xa0\xc4\xdc\x1d\xc2\xfe\x00\xba\x89\x8a\x43\x48\x39\xdb\x27\xa8\x4f\x7c\x7e\x59\x5a\x30\xd4\x3f\x8e\x8e\xb7\x56\x19\x67\x0d\xc7\x22\xa5\x8f\x51\x22\xaf\x8d\xdf\x96\x8b\x79\xe2\x92\x16\x7d\x9b\xa9\xc8\xbd\x5f\x9f\xe8\x76\xbc\xba\x79\x32\x2e\x84\xbc\x82\x38\xf8\x0f\x59\x2f\x2a\x97\xbf\x90\xba\xed\x8c\xe4\x42\xf6\x09\x00\x00\xff\xff\x11\x5d\x4c\xce\xff\x00\x00\x00")

func templates_testSingletonMssql_suites_testGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonMssql_suites_testGoTpl,
		"templates_test/singleton/mssql_suites_test.go.tpl",
	)
}

func templates_testSingletonMssql_suites_testGoTpl() (*asset, error) {
	bytes, err := templates_testSingletonMssql_suites_testGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/mssql_suites_test.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x10, 0xc4, 0x71, 0xaf, 0xd9, 0x16, 0x41, 0x8b, 0x4b, 0xfc, 0xe8, 0xba, 0xfd, 0xfa, 0x4d, 0x2c, 0x1, 0xd1, 0x0, 0xe1, 0xb0, 0x78, 0xee, 0x7f, 0xd0, 0x65, 0xf3, 0xa1, 0x43, 0xba, 0x3c, 0xe7}}
	return a, nil
}

var _templates_testUpsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x54\x4d\x6f\xdb\x30\x0c\x3d\x5b\xbf\x82\x0b\xb6\x41\x1e\x5c\x15\xbb\x76\xc8\x21\xfd\x38\x14\xc3\x82\xa0\x71\xce\x83\x6a\xd3\xa9\x10\x45\x32\x24\x7a\x49\x66\xe8\xbf\x0f\x92\xd3\x36\x6d\xda\xa1\x87\xed\xd0\x43\x62\x4b\x78\x7c\x8f\x7c\x24\xdd\xf7\x27\xf0\x51\x6a\x25\x3d\x9c\x8d\x41\x4c\xe2\x1b\x7a\x51\xca\x5b\x8d\x30\x3c\xc4\x54\xae\x31\x04\xd6\x74\xa6\x02\x42\x4f\x7d\x3f\x44\x88\x45\x3b\xd3\x9d\x93\x3a\x84\x45\xeb\xd1\x11\x27\xf8\x12\x01\xca\x2c\x45\x99\x43\xcf\x32\x12\x33\xe9\xa4\xd6\xa8\x79\xce\x58\xa6\x1a\xd0\x68\xf8\x03\xc1\xa5\xdd\x98\xb9\x32\xcb\x4e\x4b\x17\xc2\x44\xeb\x0b\xab\xbb\xb5\xf1\x39\x8c\xc7\x7f\x43\xce\x9c\x5a\x4b\xb7\xfb\x8e\xbb\x87\x80\x9e\x65\x19\x89\xf9\x4a\xb5\x7c\x14\xff\x5b\x65\x96\x40\xa9\x8c\x8d\xa2\x3b\xb0\x46\xef\xa0\x1d\xe2\x60\x85\x3b\xa8\x86\xc8\x51\xce\xb2\xc0\x58\xe6\x11\xeb\x68\x81\x93\xa6\xb6\x6b\xf5\x1b\xc5\x14\x37\x73\xc4\x9a\xe7\x2c\xfb\x25\x1d\xa0\x4b\x3f\xeb\x58\x76\x7a\x0a\x13\x22\x5c\xb7\x04\x74\x87\x70\x3d\x9d\x5f\xdd\x94\xe0\x55\x8d\x60\x1b\x90\x06\x16\xb3\x78\xc3\x32\x1b\x19\x0f\xec\x7a\xac\xa0\x0f\xc9\x8d\x48\x7a\xa8\x39\x27\xd7\x55\xc4\x63\x32\x05\x7c\xb6\x05\xbc\x62\xc0\xe5\x79\xb9\x6b\xd1\x17\x40\xae\xc3\xfc\x5b\xe2\xf9\x30\x06\xa3\xf4\xde\x88\xab\x98\x69\xc3\x47\x0b\x93\x2c\x20\xfb\x28\xf2\x72\x42\xe0\x93\xf4\x19\x7c\xf2\xa3\x22\xf2\xed\x7d\xe9\x7b\xd5\x80\xb1\x04\x62\x6a\x2f\xac\x21\xdc\x52\x08\x15\x6d\x63\x65\xd5\x70\x16\xe7\xb2\x5a\x2d\x9d\xed\x4c\xcd\xf3\xbe\x47\x53\x87\xc0\xb2\x01\xf2\xa3\xf3\x54\x6e\x79\x62\x39\x64\x38\xba\xb8\xb5\x4a\x8b\x73\x5c\x2a\x93\x38\xb4\xc7\xc3\xbb\x72\xcb\x2b\xda\x16\xb1\xc0\x7b\x85\x37\x81\x72\x96\xd5\xd8\xa0\x83\x38\xbc\x3c\x87\x1e\x7e\xc2\x18\x68\x2b\x6e\xac\xd6\xb7\xb2\x5a\xf1\x1c\x42\xec\xf0\x43\x2f\xac\xd8\xcf\xf2\x6b\x85\xc7\x9e\xa0\xa9\xe1\x24\x04\x88\xa7\xa4\x7f\x6d\x1a\x74\x3c\x7f\x7a\x7a\x5b\x5f\xba\x24\xf7\x72\x53\x8e\xba\x51\xd9\xce\x50\xba\x78\x36\x59\xf7\x8b\xc8\x73\x71\x11\x31\x6f\x4c\xff\xb1\xf2\xe3\x2c\xf9\xbd\x6c\x84\x24\xe1\x08\xfa\xfa\x04\x32\xda\x48\x43\x60\x0d\x82\xc3\xca\xba\xba\x80\xa5\xa5\xb3\x51\x31\xe0\xf7\x49\x3f\x5b\x97\xc5\xec\x72\x52\x5e\xbd\xb4\x2e\xff\x62\x21\x1a\xa9\x3d\xbe\x0a\x3b\xfa\x70\x08\x21\xfe\xeb\xfa\xbc\xbf\xb9\x7a\x27\x63\x15\xd8\x9f\x00\x00\x00\xff\xff\xf6\x71\x76\xb4\xbb\x06\x00\x00")

func templates_testUpsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testUpsertGoTpl,
		"templates_test/upsert.go.tpl",
	)
}

func templates_testUpsertGoTpl() (*asset, error) {
	bytes, err := templates_testUpsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xea, 0xe3, 0xe8, 0x1e, 0xc3, 0xef, 0x74, 0x5d, 0xf, 0xf2, 0x30, 0xa1, 0x6, 0x82, 0x8f, 0x70, 0xe4, 0xb, 0xca, 0x51, 0x16, 0xa2, 0x3d, 0x8f, 0x40, 0x59, 0xec, 0xfe, 0x96, 0x39, 0x2, 0x36}}
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
	"templates/17_upsert.go.tpl":                        templates17_upsertGoTpl,
	"templates/singleton/mssql_upsert.go.tpl":           templatesSingletonMssql_upsertGoTpl,
	"templates_test/singleton/mssql_main_test.go.tpl":   templates_testSingletonMssql_main_testGoTpl,
	"templates_test/singleton/mssql_suites_test.go.tpl": templates_testSingletonMssql_suites_testGoTpl,
	"templates_test/upsert.go.tpl":                      templates_testUpsertGoTpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"17_upsert.go.tpl": &bintree{templates17_upsertGoTpl, map[string]*bintree{}},
		"singleton": &bintree{nil, map[string]*bintree{
			"mssql_upsert.go.tpl": &bintree{templatesSingletonMssql_upsertGoTpl, map[string]*bintree{}},
		}},
	}},
	"templates_test": &bintree{nil, map[string]*bintree{
		"singleton": &bintree{nil, map[string]*bintree{
			"mssql_main_test.go.tpl":   &bintree{templates_testSingletonMssql_main_testGoTpl, map[string]*bintree{}},
			"mssql_suites_test.go.tpl": &bintree{templates_testSingletonMssql_suites_testGoTpl, map[string]*bintree{}},
		}},
		"upsert.go.tpl": &bintree{templates_testUpsertGoTpl, map[string]*bintree{}},
	}},
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
