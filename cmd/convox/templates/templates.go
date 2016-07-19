// Code generated by go-bindata.
// sources:
// templates/init/rails/.dockerignore
// templates/init/rails/Dockerfile
// templates/init/rails/docker-compose.yml
// templates/init/ruby/.dockerignore
// templates/init/ruby/Dockerfile
// templates/init/ruby/docker-compose.yml
// templates/init/sinatra/.dockerignore
// templates/init/sinatra/Dockerfile
// templates/init/sinatra/docker-compose.yml
// templates/init/unknown/.dockerignore
// templates/init/unknown/Dockerfile
// templates/init/unknown/docker-compose.yml
// templates/templates.go
// DO NOT EDIT!

package templates

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

var _initRailsDockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd7\x4b\x2a\xcd\x4b\xc9\x49\xe5\xd2\xd7\x4b\xce\xcf\x2b\xcb\xaf\x00\x32\x52\xf3\xca\x80\x64\x7a\x66\x09\x97\x7e\x4a\x92\xbe\x96\x5e\x71\x61\x4e\x66\x49\xaa\x31\x2a\x4f\x37\x2b\xbf\xb4\x28\x2f\x31\x87\x4b\x3f\x27\x3f\x5d\x5f\x8b\x4b\xbf\x24\xb7\x80\x0b\x10\x00\x00\xff\xff\xa0\x04\x95\x56\x4e\x00\x00\x00")

func initRailsDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_initRailsDockerignore,
		"init/rails/.dockerignore",
	)
}

func initRailsDockerignore() (*asset, error) {
	bytes, err := initRailsDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/rails/.dockerignore", size: 78, mode: os.FileMode(420), modTime: time.Unix(1464029683, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initRailsDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x50\xbb\x4e\xc4\x30\x10\xec\xfd\x15\x2b\xd1\x5f\x7a\x5a\x24\xa8\xe0\x50\x24\x0a\xba\xf3\x39\xeb\x60\xe2\xd8\x96\x1f\x27\xf2\xf7\x6c\xd6\x09\x21\xa7\xa4\xca\x3c\xec\x99\xf1\x73\x7b\x7e\x05\xe5\xdd\xcd\xff\x34\x51\x1a\x9b\x84\x78\x20\x1c\x26\xf0\xce\x4e\x90\xbf\x10\xb4\xb1\x98\xc0\x21\x76\xd8\x81\xf6\x11\xae\xc5\x75\x16\xc1\xb8\x94\xa5\xb5\xe4\x2f\x4e\xf9\x71\x44\x97\xd9\x7f\x43\xd7\xf9\xd8\x28\xa9\x08\x58\xe3\xc8\xa9\x61\xf2\x05\x2e\xcb\xc1\x20\xd5\x20\x7b\xbc\xcc\x64\x84\x1e\xc7\x24\x9e\xce\xef\x9f\xf0\x82\xe3\x9c\x05\xfc\x35\x32\x84\x66\x61\x76\xf2\xc9\x7a\x35\xec\x64\x66\xa8\x06\xbb\x76\xe9\xec\xfa\xcf\x88\xf6\xe3\xed\xbe\xff\x3a\xf8\xbb\xa4\x7c\x3c\x58\xa6\x84\x39\x3d\x86\x88\xb4\x33\xfc\x15\x6a\xe5\x80\x4b\x61\x0e\x5a\x71\x55\xe9\x51\xb5\xe9\xb7\x2d\x15\x57\x2d\x94\xab\x35\x6a\xd3\x2a\xae\xda\x8c\x6b\x60\xd5\x36\xcc\xe5\x23\x85\x1c\x14\x5a\x47\xcc\xfd\x23\xd2\x10\xaf\xf9\x9f\x4e\xd7\x6b\x4f\x7c\x9b\xf8\x0d\x00\x00\xff\xff\x08\xae\x24\x4e\xf0\x01\x00\x00")

func initRailsDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_initRailsDockerfile,
		"init/rails/Dockerfile",
	)
}

func initRailsDockerfile() (*asset, error) {
	bytes, err := initRailsDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/rails/Dockerfile", size: 496, mode: os.FileMode(420), modTime: time.Unix(1463624980, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initRailsDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x4f\x4d\xb2\xe2\x52\x50\x48\x2a\xcd\xcc\x49\xb1\x52\xd0\x03\x32\x73\x12\x93\x52\x73\x8a\x41\x82\x0a\x0a\xba\x0a\xc9\xf9\x79\x65\xf9\x15\x7a\x05\xf9\x45\x25\x7a\x26\x26\xc6\x7a\x05\x45\xf9\x25\xf9\xc9\xf9\x39\xb6\x25\x39\xc5\xb8\x95\x54\x54\xda\x96\x14\x95\xa6\x02\x15\x80\x44\xe1\x86\x59\x18\x58\x99\x18\x18\x18\x40\x79\x40\xb5\x20\xae\x21\x17\x20\x00\x00\xff\xff\xc0\xe1\x22\xef\x84\x00\x00\x00")

func initRailsDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_initRailsDockerComposeYml,
		"init/rails/docker-compose.yml",
	)
}

func initRailsDockerComposeYml() (*asset, error) {
	bytes, err := initRailsDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/rails/docker-compose.yml", size: 132, mode: os.FileMode(420), modTime: time.Unix(1463624980, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initRubyDockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd7\x4b\x2a\xcd\x4b\xc9\x49\xe5\xd2\xd7\x4b\xce\xcf\x2b\xcb\xaf\x00\x32\x52\xf3\xca\x80\x64\x7a\x66\x09\x17\x20\x00\x00\xff\xff\xc9\x68\x92\x70\x1e\x00\x00\x00")

func initRubyDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_initRubyDockerignore,
		"init/ruby/.dockerignore",
	)
}

func initRubyDockerignore() (*asset, error) {
	bytes, err := initRubyDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/ruby/.dockerignore", size: 30, mode: os.FileMode(420), modTime: time.Unix(1464029683, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initRubyDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x8f\xcd\x4e\xc0\x20\x10\x84\xef\x3c\xc5\x26\xde\xcb\x43\x98\xe8\x49\x6b\x9a\x78\xf0\x56\x0a\x4b\x6d\x0a\xbb\x84\x9f\x46\xde\x5e\x8a\x35\xb1\x72\x62\x67\xbf\x61\x86\xa7\x69\x7c\x01\xcd\x74\xf0\x97\x8c\x65\xa9\x42\x3c\xb4\x31\x54\x60\x72\x15\xf2\x27\x82\xdd\x1c\x26\x20\x44\x83\x06\x2c\x47\x58\x0a\x19\x87\xb0\x51\xca\xca\xb9\xc6\x17\xd2\xec\x3d\x52\xee\xfc\x81\x64\x38\x4a\xad\x74\x1b\xdc\x46\x8d\xb4\x50\xb9\xc0\x7c\x19\x83\xd2\xbb\x5a\x71\x3e\xc5\x08\x2b\xfa\x24\x1e\xc7\xb7\x0f\x78\x46\x7f\x66\x41\x3f\x52\x85\x20\x2f\xe5\xb6\x1e\x1c\xeb\xfd\xb6\xee\x4a\xab\xd1\xa9\x5b\x7a\xa7\xfe\x2a\x62\x7a\x7f\xfd\xdf\xff\xf7\xc3\x67\xf7\x88\x29\x03\xdb\x7e\x6f\xde\x9f\xe0\xa1\xbf\x23\xbe\x03\x00\x00\xff\xff\x8b\xae\xa0\xae\x2a\x01\x00\x00")

func initRubyDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_initRubyDockerfile,
		"init/ruby/Dockerfile",
	)
}

func initRubyDockerfile() (*asset, error) {
	bytes, err := initRubyDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/ruby/Dockerfile", size: 298, mode: os.FileMode(420), modTime: time.Unix(1463624980, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initRubyDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x8c\xc1\x0a\x83\x30\x0c\x86\xef\x3e\x45\xf0\x6e\xe9\x98\x87\x21\xf8\x30\x5a\x03\xca\xa2\x29\x69\x3a\xed\xdb\x2f\x85\x6d\xb7\xdd\xfa\xf5\xff\xf2\x9d\x38\x0f\x0d\xc0\x9c\x37\x5a\x06\x70\xf6\x0c\xbc\xef\xd3\x61\x80\x61\x65\x68\x71\xd9\x14\x16\x0e\x4f\x94\xce\xa6\xc8\x09\x5d\xd9\x09\xce\x4d\x57\x28\x9c\x05\x92\x4e\xa2\x39\x7e\x0f\x5b\x6b\xd0\x34\x23\xa5\x1a\x06\xe8\x6c\x38\x5e\x7c\xb9\xc8\xa2\xae\xef\xef\x2e\x0a\x2b\x07\xa6\x51\x29\xfd\x57\xae\x32\xaa\x64\x34\xa1\xfe\xfe\x62\x0f\x3f\xf4\xde\xfb\x0f\x99\x5b\xf1\xd6\xbc\x03\x00\x00\xff\xff\xae\x01\x4e\xf5\xc8\x00\x00\x00")

func initRubyDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_initRubyDockerComposeYml,
		"init/ruby/docker-compose.yml",
	)
}

func initRubyDockerComposeYml() (*asset, error) {
	bytes, err := initRubyDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/ruby/docker-compose.yml", size: 200, mode: os.FileMode(420), modTime: time.Unix(1463624980, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initSinatraDockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd7\x4b\x2a\xcd\x4b\xc9\x49\xe5\xd2\xd7\x4b\xce\xcf\x2b\xcb\xaf\x00\x32\x52\xf3\xca\x80\x64\x7a\x66\x09\x17\x20\x00\x00\xff\xff\xc9\x68\x92\x70\x1e\x00\x00\x00")

func initSinatraDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_initSinatraDockerignore,
		"init/sinatra/.dockerignore",
	)
}

func initSinatraDockerignore() (*asset, error) {
	bytes, err := initSinatraDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/sinatra/.dockerignore", size: 30, mode: os.FileMode(420), modTime: time.Unix(1464029683, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initSinatraDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x8f\xc1\x6e\x84\x20\x10\x86\xef\x3c\xc5\x24\xbd\xcb\x43\x34\x69\x4f\xad\x8d\x49\x0f\xbd\x49\x61\x70\x89\x30\x43\x00\xcd\xfa\xf6\x8b\xac\x9b\xac\xeb\x49\xbe\xf9\x86\xff\xe7\x63\xe8\xbf\x40\x33\xad\x7c\x95\xd9\x91\x2a\x49\x09\xf1\x56\x49\xdc\x80\xc9\x6f\x50\x2e\x08\xd6\x79\xcc\x40\x88\x06\x0d\x58\x4e\xf0\xbf\x90\xf1\x08\x8e\x72\x51\xde\x57\x7f\x21\xcd\x21\x20\x95\xe6\xaf\x48\x86\x93\xd4\x4a\xd7\x83\x77\x54\x4d\x0b\x1b\x2f\x30\x1e\x8b\x51\xe9\x59\x4d\x38\xee\x30\xc1\x84\x21\x8b\xf7\xfe\xe7\x0f\x3e\x31\xec\x59\xd0\x3e\xa9\x62\x94\x07\x39\x8d\x3b\xcf\x7a\x3e\x8d\x1b\xa9\x35\x9a\x75\x4a\x6f\xd6\x33\x11\xc3\xef\xf7\x6b\xff\xc7\x83\xf7\xee\x09\x73\x01\xb6\xed\xbf\xee\xde\x83\xbb\x76\x8f\xb8\x05\x00\x00\xff\xff\x9c\x51\x49\xbe\x2d\x01\x00\x00")

func initSinatraDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_initSinatraDockerfile,
		"init/sinatra/Dockerfile",
	)
}

func initSinatraDockerfile() (*asset, error) {
	bytes, err := initSinatraDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/sinatra/Dockerfile", size: 301, mode: os.FileMode(420), modTime: time.Unix(1463624980, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initSinatraDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x4f\x4d\xb2\xe2\x52\x50\x48\x2a\xcd\xcc\x49\xb1\x52\xd0\x03\x32\x73\x12\x93\x52\x73\x8a\x41\x82\x0a\x0a\xba\x0a\xc9\xf9\x79\x65\xf9\x15\x7a\x05\xf9\x45\x25\x7a\x26\x26\xc6\x7a\x05\x45\xf9\x25\xf9\xc9\xf9\x39\xb6\x25\x39\xc5\xb8\x95\x54\x54\xda\x96\x14\x95\xa6\x02\x15\x80\x44\xe1\x86\x59\x18\x58\x99\x18\x18\x18\x40\x79\x40\xb5\x20\xae\x21\x17\x20\x00\x00\xff\xff\xc0\xe1\x22\xef\x84\x00\x00\x00")

func initSinatraDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_initSinatraDockerComposeYml,
		"init/sinatra/docker-compose.yml",
	)
}

func initSinatraDockerComposeYml() (*asset, error) {
	bytes, err := initSinatraDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/sinatra/docker-compose.yml", size: 132, mode: os.FileMode(420), modTime: time.Unix(1463624980, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initUnknownDockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\x4b\xcd\x2b\xe3\xd2\x4b\xcf\x2c\xe1\x02\x04\x00\x00\xff\xff\x9c\x10\x28\x7b\x0a\x00\x00\x00")

func initUnknownDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_initUnknownDockerignore,
		"init/unknown/.dockerignore",
	)
}

func initUnknownDockerignore() (*asset, error) {
	bytes, err := initUnknownDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/unknown/.dockerignore", size: 10, mode: os.FileMode(420), modTime: time.Unix(1464952501, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initUnknownDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0b\xf2\xf7\x55\x28\x4d\x2a\xcd\x2b\x29\xb5\x32\x34\xd3\x33\x30\xe1\xe2\x72\xf6\x0f\x88\x54\xd0\x53\xd0\x4f\x2c\x28\xe0\x02\x04\x00\x00\xff\xff\xf1\xa3\x65\xfc\x1f\x00\x00\x00")

func initUnknownDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_initUnknownDockerfile,
		"init/unknown/Dockerfile",
	)
}

func initUnknownDockerfile() (*asset, error) {
	bytes, err := initUnknownDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/unknown/Dockerfile", size: 31, mode: os.FileMode(420), modTime: time.Unix(1464952501, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initUnknownDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xca\x4d\xcc\xcc\xb3\xe2\x52\x50\x48\x2a\xcd\xcc\x49\xb1\x52\xd0\xe3\x02\x04\x00\x00\xff\xff\xa6\xe1\xc1\x85\x11\x00\x00\x00")

func initUnknownDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_initUnknownDockerComposeYml,
		"init/unknown/docker-compose.yml",
	)
}

func initUnknownDockerComposeYml() (*asset, error) {
	bytes, err := initUnknownDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/unknown/docker-compose.yml", size: 17, mode: os.FileMode(420), modTime: time.Unix(1464952501, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x9a\xdb\x6e\x1b\x47\x12\x86\xaf\xc5\xa7\x98\x08\x48\x40\x2d\x14\x6a\xce\x07\x01\xbe\x89\xed\x05\x7c\xb1\x0e\x90\x38\x17\x8b\xd5\x22\x98\x43\x8f\x96\xb0\x44\x6a\x49\x2a\x91\x6c\xf8\xdd\xb7\xbe\xae\x1e\x91\x22\x87\xd4\x19\xf1\xee\xc6\xc0\x98\x33\x3d\xdd\xd5\x55\xdd\x55\x7f\xff\x55\xa3\xa3\x23\xef\xf5\xb4\x31\xde\xa9\x99\x98\x59\xb9\x30\x8d\x57\x5d\x7b\xa7\xd3\xef\xab\xf1\xa4\x29\x17\xe5\x68\x20\x1d\xe6\xd3\xcb\x59\x6d\xe6\xc7\xdc\x2f\xcc\xf9\xc5\x99\xf4\x9b\x1f\x8d\x27\xe3\xc5\xd1\xac\x1c\x9f\xcd\x8f\x46\xcd\xb4\xfe\x68\x66\xe3\xd3\xc9\x74\x66\xb6\xf6\x7a\x63\x3b\xb5\xe3\xb3\xed\x5d\x54\xce\xf7\xf5\xf4\xfc\x62\x3a\x37\xa3\xeb\xf3\xb3\xbe\xae\x97\xd5\xf5\xdd\x53\xd2\x69\xf7\x8c\xf4\xb8\xd7\x84\xf3\xf1\xa4\x5c\xcc\xca\x3b\xe7\xec\xfa\xed\x9c\xb6\xeb\x74\xaf\x99\x2f\x27\x1f\x27\xd3\xdf\x27\x77\xce\xdc\xf5\xdb\x39\x73\xd7\xe9\xae\x99\x6f\xee\x46\xa7\x53\xde\xbc\xf9\xd1\x7b\xff\xe3\x07\xef\xed\x9b\x77\x1f\xbe\x19\x0c\x2e\xca\xfa\x63\x79\x6a\x96\xfd\x07\x83\xb1\x08\x9a\x2d\xbc\xe1\x60\x6f\xbf\xba\x96\x96\x7d\xb9\x41\xfa\xcc\xcc\xe7\x47\xa7\x9f\xc6\x17\x34\xb4\xe7\x0b\x7e\xc6\x53\xfd\xff\x68\x3c\xbd\x5c\x8c\xcf\x78\x98\xda\x01\x17\xe5\xe2\x5f\x47\x68\xce\x0d\x0d\xf3\xc5\x6c\x3c\x39\xb5\xef\x16\xe3\x73\xb3\x3f\x38\x18\x0c\xda\xcb\x49\xed\x39\xd7\xfc\xc9\x94\xcd\x90\x1b\xef\x1f\xff\x64\xda\x43\x6f\x52\x9e\x1b\x4f\x87\x1d\x78\xc3\xae\xd5\xcc\x66\xd3\xd9\x81\xf7\x79\xb0\x77\xfa\xc9\x3e\x79\xc7\xaf\x3c\xb4\x1a\xbd\x37\xbf\x23\xc4\xcc\x86\x56\x6d\x9e\x7f\xb8\x6c\x5b\x79\x46\xec\xc1\xc1\x60\x6f\xdc\xda\x01\xdf\xbc\xf2\x26\xe3\x33\x44\xec\xcd\xcc\xe2\x72\x36\xe1\xf1\xd0\x13\x93\x46\x6f\x91\xde\x0e\xf7\x11\xe4\x7d\xfb\xef\x63\xef\xdb\xdf\xf6\x55\x13\x3b\x97\xc8\xf8\x32\x18\xec\xfd\x56\xce\xbc\xea\xb2\xf5\x74\x1e\x9d\x64\xb0\xf7\xab\xaa\xf3\xca\x1b\x4f\x47\xaf\xa7\x17\xd7\xc3\xef\xa4\xcf\xa1\xe8\x26\xa3\xea\xb3\xb7\x9d\xa6\xa3\xd7\x67\xb2\x4f\x43\x31\xff\x99\xf4\x41\x8c\xca\xdf\x22\x48\x3a\xaa\xde\xae\x51\xd4\x1a\xfd\x80\xea\xc3\x83\x43\x7a\x0c\xe4\xdd\xe2\xfa\xc2\x78\xe5\x7c\x6e\x16\x2c\xf9\x65\xbd\x40\x8a\xb5\xcf\xed\x87\x4c\x33\x69\xa7\x9e\x37\x9d\x8f\xfe\x2a\xdb\xfa\x4e\x1e\x6e\xc6\xb9\x2d\xec\xda\x57\x24\xd8\x3d\x94\x7f\xba\x8d\x83\xbd\xf9\xf8\x93\x7d\x1e\x4f\x16\x69\x3c\xd8\x3b\x07\xab\xbc\x1b\xa1\x7f\x93\x47\xdb\xf8\x41\x3c\xc4\xc3\x4d\x46\xdc\x31\x8f\x75\x95\x61\x3b\x5e\x9f\xeb\xc0\x7b\x2f\x53\x0c\x0f\xdc\x0c\xcc\xe9\xac\x6c\xc7\x23\x66\x97\xc1\xdb\xc7\xfe\x2c\xea\xc8\x58\xab\xcd\xed\xa1\x28\xba\x73\x28\xba\xca\xd0\x15\xcd\x6f\x0b\xc0\xb4\xbb\x04\x60\x9c\xc8\xb8\x31\x74\x43\x82\xb3\x7e\xbb\x90\x77\xf3\x37\xe3\x99\x88\xa8\xa6\xd3\xb3\xd5\xd1\xe5\xd9\xfc\x0e\xcb\xaf\xe7\x6a\xb8\xe0\x4b\x59\x9b\xcf\x5f\x56\x46\x3b\x97\xc0\xcb\x7f\x05\x6a\x7e\x02\xcc\xdf\xac\x40\x96\xf8\xb8\x3a\xc5\x70\xff\xe4\x2a\x68\x4f\xae\xf2\xea\xe4\xca\xcf\xe5\xf2\xdd\x55\x9c\x5c\xa5\x46\xda\x5d\x5b\x2b\x7d\x9a\x50\xae\xec\xe4\x2a\x96\xbe\x61\x79\x72\x55\x37\x7a\x5f\x4b\xdf\x58\x2e\x93\xdc\xee\x53\xcb\xf8\x5a\xc6\x85\xdc\xcb\x55\xb6\x2a\x2b\x92\x3e\x89\x5c\x6d\x24\xed\x22\x27\x97\xb6\x34\x3e\xb9\xca\xe4\x3e\x4d\x75\xee\x42\x64\x64\x32\x3e\x96\xb6\x42\xfa\x56\x72\x5f\xc8\xbb\x44\x7e\xb3\x40\xfa\xc9\x15\x1b\xed\xcf\xdc\xa5\xf4\x8b\x02\xd5\x2b\x96\x79\xa2\x4c\xe7\xad\xe4\xbe\x12\xd9\xa1\xd8\x11\xb6\xda\x27\x77\xfa\x45\xe8\x96\xe9\x6f\x22\xb6\x24\x6e\x1d\x62\x37\x2e\x94\x71\x55\xa6\xfa\xf9\xd2\x16\xf8\xcb\xf5\x61\x3d\xb8\x4a\x9e\xa5\x5f\x21\xb6\x27\xa9\xea\x74\xb3\x86\xfe\x7e\x07\x91\xbd\x9b\xe0\x22\xb8\x0f\x19\xbb\x38\x5f\x41\x56\x81\x84\xfe\xbd\x3c\x94\x37\xfb\xdb\x0e\xff\x7d\x79\x7b\x70\x13\x7e\xbd\xe3\xd1\xe0\x2f\x16\x38\x56\x35\xb0\xc8\x71\x03\xcf\xbb\xf4\xbf\x0b\x05\x6f\xc0\xcb\xc2\x8f\x08\x5b\x73\xe5\xcf\x04\xf9\xb1\xb7\xc3\x04\x8f\x58\x3e\xf6\xb2\xfc\xd0\x23\x28\x8f\x57\x63\x76\x18\x87\xfe\x81\x6d\x27\xd4\x8e\x35\x14\x7f\x99\x8c\xaf\x86\x41\x9c\xc6\x7e\x58\xa4\x79\x74\xe8\xf9\x07\x82\xb2\x25\x93\x7f\x67\x2d\xfd\x6c\xcd\x3b\xf6\x9c\x95\x68\x76\x6c\xff\xff\x72\xb3\xf6\xe5\xe1\xae\x30\xe2\x60\x7c\x54\x10\xa5\xb5\x38\x8a\xdc\x57\x95\x3a\x4b\x2d\xce\x13\xf9\xea\x5c\x46\xde\xb5\xe2\x88\x41\xa2\xce\xdb\x04\xea\x94\x04\x46\x52\xaa\x43\x96\x22\xcb\xf8\x2a\x83\x67\x5f\xda\xab\x52\x83\x29\x22\x08\x65\x5c\x8a\x2c\x02\x31\xd7\xa0\x09\x9c\xa3\xb7\x04\x5d\xa6\x3a\x34\x2e\xd0\x42\x99\xa3\x94\xb6\x32\xd6\x60\x8c\x6a\xd5\xa3\x90\x77\xad\xbc\xcb\x44\x6e\x56\x69\x30\xfa\x89\x0b\xf2\x46\x83\x1f\x7b\x22\x19\x97\x48\xbf\x80\x40\x95\x7e\x39\x81\x4c\xb0\x61\x93\xc8\x09\x65\x9e\xd6\x57\x30\xc0\xde\xc2\xd7\xe0\xc2\x5e\x02\xb3\x92\xb1\x85\xf4\xaf\x7d\xd5\x25\x4b\x55\xef\x5c\xee\x5b\x74\x47\x3f\xd6\x49\xe6\xcd\xe4\x0a\xa4\xad\x96\xb6\x0a\xdb\x58\x0f\x69\xab\xd0\x2b\xd6\x00\x67\x8e\xb6\x50\x90\x88\x63\x05\x92\x86\x67\xde\x45\x0a\x4e\xbc\x67\x0e\x40\xa1\x2c\x74\xcf\x12\xd6\x15\x40\x72\x6d\x00\x0c\xeb\x85\xcd\xa1\x51\x00\x63\x5e\xec\x6a\x12\xfd\x65\x5d\x2a\xb1\xb1\xae\x15\xc0\xb0\x3d\xa0\x6f\xa6\x7b\x53\x18\xb5\xdb\xc8\xdc\x45\xab\xeb\x90\x87\x3a\x4f\x51\xab\xec\x56\x7e\xa3\x44\x41\x91\xf1\x00\x5a\xea\xd6\x81\xf9\x01\x57\xf6\x9f\x3e\xc6\xcd\x43\x1f\xfc\xc0\x4f\xd5\x8f\xe8\x6b\x0a\xe7\x3f\xa5\x82\x2e\x3e\xc7\xfa\x31\x97\x69\x14\x40\xed\x7e\xe1\x2b\xb9\x8e\x63\xcf\xb3\x5a\xf7\x1c\xfb\xcb\x54\x75\xc0\x8f\x22\x19\x93\xa7\x2a\x87\x7d\x8a\x22\xd5\x15\xdf\x4c\x4b\xf5\x03\x80\x11\xb0\x44\x5f\xdf\xa8\x8f\xb2\xee\x49\xa2\xfa\xe0\x17\x95\xbb\x07\x48\x43\xe7\xdb\x89\xbb\xb7\xfb\x57\xa9\x6d\xc8\x6c\x22\x05\x70\xf6\xbb\xe1\x60\xc0\x5e\xe7\xcf\xac\x79\x94\xea\x5a\x33\x77\x28\x7d\xf3\x44\x7d\x2c\x88\x75\xee\x38\xd3\xf7\xc4\x10\xef\x39\x8c\xf0\x33\x0e\x1e\xf6\xbc\x70\xfe\x81\x0f\x20\x97\x83\x02\xfb\xf1\x5b\xd6\xc4\x6f\x36\x01\x1e\xdf\x40\x1f\xf6\xd3\xfa\x16\xef\x83\xbb\x00\x1e\x78\x78\x32\xbc\x23\x64\x1d\xdc\x97\x6f\x76\x22\x3b\x1d\x1e\x81\xeb\x2b\x6a\xbf\x00\xaa\xaf\xea\xee\x20\x3d\x2e\xd2\x87\x62\x7a\x94\x86\x71\x91\xfb\x2f\x80\xe9\xaf\x35\x01\xfb\xfb\xf9\xd9\xa3\x90\xbd\xa3\x1e\x31\xd1\x16\x2a\x02\x43\x73\x40\xea\x38\x5f\x52\x26\x3c\x14\xca\x02\xf2\xf1\xbe\x41\x5e\xa4\x08\x00\xda\x06\xf2\x5b\x44\xfa\x8e\xe7\x1c\xb9\x81\x22\x07\x68\xdf\x21\x3e\xbf\xa0\x83\x45\x3b\xf4\x49\xf4\x9e\x08\x03\x31\x7c\xf7\x1c\x13\x2d\xae\x0d\x34\xe6\xaa\xd3\x65\x9f\xd8\xf5\x0b\xdd\x6f\x27\x13\x14\xa8\x52\x6d\xe7\x1e\x04\xb6\xa8\x4b\x04\xc7\x7a\x11\xa5\xf6\x84\x71\x14\x08\x04\xf1\x43\xd5\xc1\xa2\x83\xb4\x9b\x40\x91\x24\x41\x37\x87\xb4\x85\xbb\xef\xae\xd8\x57\x1b\xf8\xad\x1c\xba\xda\xa8\xe3\x64\xc9\xf4\x79\x3d\x32\x39\x31\x90\x1d\x86\x8a\xb0\x20\xd7\xdd\xd4\x6b\xb9\xc9\x4f\x8e\xcf\xa5\xa8\xf5\x28\xdd\xcc\xe8\x77\x46\xeb\x52\xd0\x23\x62\x76\xc3\xa0\x17\x88\xdc\x3e\x7b\x5c\x04\x07\x51\xf8\x47\x47\xf0\x65\x75\xfd\x5f\x95\xdb\x6c\x75\x68\xf4\xca\x35\xdf\xc9\x7c\xa5\x12\xdb\x1c\x7a\xcd\xe6\x47\xfa\xf2\x9a\x94\x15\x37\xde\x28\xe9\xf5\x38\xf0\xda\xe8\x7b\xfb\x6e\xbf\xee\xcf\xeb\xb6\x3d\xfa\x3b\x87\x8d\xfc\x3f\x3a\x8b\xb8\xb1\xff\xd1\x49\x44\x22\x47\x48\xde\x3a\x2f\x35\x8e\x3c\xbb\x24\x02\x18\x04\x0e\x21\x75\x40\x36\x70\xdf\x38\x02\x0b\x01\x86\x6c\x9b\x5c\xbd\x1a\x12\x54\x94\x4a\xd8\x21\x36\x64\xb0\x1c\x2b\x96\xd4\x35\xee\xa8\x71\xa4\xde\x92\x27\x27\x8b\x2c\x3c\x77\xa4\x94\x63\x2c\x13\x8f\x4d\xb9\x32\x25\xf7\xa9\x83\x7c\x12\x8b\xb4\x50\x92\x05\x69\x42\xdf\x2c\xd6\xb9\x20\xe6\x96\xe0\x25\x4a\xb8\xe3\x50\x75\xb6\x64\x33\xd0\xe3\x05\xe2\x97\xb9\xe3\x04\x32\x4c\x12\xc3\x31\x48\x02\xd2\x25\x17\xa1\x3b\x66\xf2\x48\x09\x71\x58\x2b\x11\xe4\x98\x41\xff\xa0\xd0\x8c\xbf\xf2\x95\x90\x12\x95\x5c\x55\xa1\xc7\x20\xd1\x48\x64\x43\x9e\x21\xbd\x44\xaf\x31\x4a\x76\x49\x42\x88\xde\x28\xd7\x4a\x44\xe9\xf4\x0f\x48\xce\x6a\x5d\x8f\xbc\x51\x9d\x6d\x52\x57\xe8\x5e\x60\xaf\x9d\xd7\x91\x50\xd6\xb0\x23\xc7\x91\x4b\x66\xd8\x63\xd0\xa3\x75\x89\x5d\x97\x38\x40\x94\xfd\x56\xf7\x04\xf9\x24\x2e\xb6\xca\x11\xb8\x4a\x45\xa8\xc4\x17\xbb\x49\x78\x40\x26\x8e\x69\xc6\xb3\x26\x24\x80\x10\x60\xda\x8c\xab\xa0\x94\x2e\x99\x64\x9d\x93\x4a\xd1\x88\xbd\x69\x8d\xa2\x1d\x7b\x47\x1b\x73\x31\xb6\x71\x88\x44\xb2\x58\x47\xba\xaf\xdc\xe3\x83\x61\xa1\x74\xa5\x71\x49\x61\xda\xaa\x4f\xe0\x1f\x24\xa6\x65\xe0\xaa\x27\x91\x26\x1d\xf4\x5d\x47\x3a\x7c\x9c\xe3\xbd\xec\x8e\xf9\x72\x3b\xa9\xbe\x15\x2d\x4f\xc5\xb9\x35\x4a\x7d\xfb\x9b\xc4\x2e\x88\x7b\x10\xa1\xee\x53\xf9\xf9\xe1\xad\x87\x4e\x87\xc5\x83\x4b\x24\x2f\x76\x18\x3f\x91\x4d\x13\x66\xc0\x03\x39\x37\x61\x4c\x28\x51\x27\xf1\x6b\x85\x16\x0b\x71\xc6\xb1\x57\x5f\xc7\x93\x67\x02\x6f\x84\x3b\xcc\x91\x9c\x8e\x31\x84\x9e\xef\x0e\xe4\x32\x54\x17\x06\x96\xc8\x85\xc9\x8d\x71\x65\x42\x89\xb0\x02\xf6\x08\x41\x20\x86\xd0\x6c\x13\xe7\xb6\xb0\xf2\x46\xa1\x80\xf0\x24\x57\x24\xd4\xc9\xe3\x29\x36\xda\x39\x52\x3d\xb8\xc9\xbb\xd1\x93\xda\x00\x7a\x92\xd7\x5a\x48\xf4\xf5\x17\xc8\xe3\xa0\x07\x0e\xa8\x4d\xc0\x9e\xa9\x6b\x90\x43\x93\x41\x90\x17\x43\x30\x60\xd3\xd4\x23\x58\x1f\x0a\x94\xf4\xe5\xde\xd6\x52\x1a\x97\x47\xe7\xcb\x5a\x02\xc4\x81\xe3\x00\x1b\x61\xec\xb6\x66\xd3\x6a\xc8\x03\xf5\x64\x1a\xd4\x6b\x08\x4d\xea\x26\x81\xd3\x99\xe3\x80\x9a\x10\xb6\xd9\x02\x6b\xed\xa0\x2c\x57\x59\xa5\xab\x61\x70\x51\x0b\x61\x3f\x80\x2e\x3f\x53\x1d\x81\x55\xd6\xb4\x75\x3a\xd1\x1f\x22\x44\x51\xd6\xc2\x67\xac\xb0\x00\xdc\x70\x01\x37\xe8\x05\x9c\xb5\xb1\xc2\x47\x5b\x69\x5b\xe1\x20\x8a\xda\x13\x35\x2a\xd6\xb0\x0f\x42\x98\xc3\x77\x45\x60\xf6\xa8\xce\xef\x41\x96\x9e\x4c\xfe\x7b\x24\xad\xc1\xc9\xbd\xa8\x7f\x8f\x98\x87\x83\xcb\x0b\x13\xff\x6d\xc6\x74\x50\xe3\x3f\x98\x47\x3d\x33\xd4\xfc\xac\x1f\x6e\xff\xdf\xa8\x7f\x8f\xd9\x8f\x73\xe6\x1e\x41\x4b\x5f\xee\xfd\xc4\xbe\xe9\xc9\x3d\x32\xee\xeb\xc8\xdb\xed\x78\x56\x3f\xde\x62\xc8\xd7\x92\x0c\xdc\x5a\x85\xa7\xe7\x03\x81\x7b\x17\xaf\xe4\x03\xe9\x5a\x3e\x40\xd1\xbd\x59\xe6\x03\x80\x33\x87\x21\x07\x0f\xfc\x16\x3e\x8b\xdb\x03\xc6\x95\xbb\xb7\x05\x69\x64\x15\x7a\xa0\xc6\x0e\x90\xeb\xee\x90\x4c\x1d\xa7\x73\xc5\xd4\xd2\xf1\x50\x5b\xe6\x72\x05\x79\xf4\x40\x57\x43\x4e\x10\xe9\xa1\x03\x4f\xa4\xe4\x43\x71\xd9\x72\xeb\x5a\xcb\x48\x1c\x74\x45\xf7\x35\xcf\x1d\x7a\x1c\x08\x89\xfb\xda\x07\xf7\x26\xbc\x6c\xd8\xb4\xee\x63\x88\xd1\x75\xe2\xd0\xa0\x94\x64\xcb\x6c\xbe\x5b\xa7\x54\x0f\x6c\x72\x82\xae\x88\xcb\xdc\x7c\x8c\x40\x2e\xe3\x72\xa3\xb9\x06\x87\x99\xcd\x71\x32\xe5\xc3\xa1\xcb\x5b\x42\x57\xfc\x47\x3f\x93\x6a\xd8\x53\xd6\xc3\x2e\xca\x75\x4d\xac\x5c\x1b\x98\x80\xac\x64\x99\xe6\x30\x1c\xd4\xf6\xeb\xa1\x23\x19\x91\x0b\x6f\xd6\x8b\xb9\x6c\xa1\xbc\xd0\x7e\x1c\xc0\x70\x76\xfa\xb0\x97\x76\x1f\x03\xed\xdb\x3a\xbe\x6e\x5c\x7e\x00\xf9\x21\x4f\x22\x6f\x61\xbc\x3d\xdc\x8d\xfb\x78\x11\x2a\x89\xe0\x80\x86\x08\xb0\xa7\x51\xa5\xb0\x65\x3f\x26\x18\x6d\x43\x27\xfa\x93\xc7\xd9\x3c\x30\x51\x1b\x2c\xa7\x4f\x75\x5f\xf8\x00\x43\x0e\x60\x5c\x3e\x00\x91\xb0\x90\xe5\xec\xa4\x0d\x9d\xd9\x23\xca\x95\x1c\xca\x36\xf7\x68\x74\x8d\x79\xc7\x7a\xd2\xb7\xfb\x58\x81\xaf\x52\x4a\x84\xb4\xac\x43\x21\x64\x86\x75\xe9\x7c\x87\x0f\x05\x5b\x72\x83\x8d\xe0\x79\x06\x20\xbc\x9d\x21\x6c\xfe\x05\xd1\x1d\x18\xf8\x90\x3c\x61\x9b\xfa\x2f\x82\x7f\x3d\xd9\x42\xe4\x07\x5f\xd3\x11\xfe\x67\xf9\xfd\x7f\xb5\xfc\xbe\x65\x9b\x9f\x21\x5a\xfb\x68\xf8\xf6\x3f\xe8\xbb\x23\x76\x1f\x4e\xc6\x77\x1b\xf6\x22\x71\xfc\x55\x97\xe2\x7f\xd1\xbf\x68\x7c\x16\x4a\x6e\xa9\x76\xa3\x34\xdb\x44\x2b\x6d\xad\xd6\xf8\xf0\x3d\x7c\xdd\x8f\xfb\xcf\x10\xb8\x0f\xc9\x31\x1f\xe9\x6d\xdc\xf6\xfb\x66\x8f\xca\x8f\xf3\xcb\x1e\x41\x4b\x9f\xec\xfd\xbb\xd1\x4d\x77\xec\x91\x71\x5f\x57\xdc\x6e\xc7\xb3\xba\xe1\x16\x43\x3a\x0f\x7c\x38\x9d\x2e\x92\x30\xe1\x1c\x7a\x01\x07\x7c\x34\x9d\x86\x06\x42\xd9\xba\xbf\x97\xe1\xef\x19\x70\x25\x8e\x96\xee\x18\xc1\x2d\xa9\x79\x54\xee\x6f\x37\xa0\xd1\xd4\x79\x22\x57\xa7\xc2\x3d\x39\x82\x90\x65\xeb\x41\xee\x6f\x32\x2c\x6c\x53\x93\x8a\xf4\x37\x76\xee\x1c\xba\xbf\xeb\xd9\xe6\xd2\xd0\xd4\x32\x72\xc7\x4c\xad\xb6\xdc\xc7\xa5\x1f\x4f\x8b\x36\xc4\x6c\xba\xf3\x2e\x5a\xb4\x31\xfc\x51\x9e\xfc\x52\xb4\xa8\xcf\x82\x8e\x16\x3d\x98\x15\x3d\xa7\x13\xff\x27\x00\x00\xff\xff\xab\xf8\x79\x97\x00\x30\x00\x00")

func templatesGoBytes() ([]byte, error) {
	return bindataRead(
		_templatesGo,
		"templates.go",
	)
}

func templatesGo() (*asset, error) {
	bytes, err := templatesGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates.go", size: 24576, mode: os.FileMode(420), modTime: time.Unix(1468932261, 0)}
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
	"init/rails/.dockerignore": initRailsDockerignore,
	"init/rails/Dockerfile": initRailsDockerfile,
	"init/rails/docker-compose.yml": initRailsDockerComposeYml,
	"init/ruby/.dockerignore": initRubyDockerignore,
	"init/ruby/Dockerfile": initRubyDockerfile,
	"init/ruby/docker-compose.yml": initRubyDockerComposeYml,
	"init/sinatra/.dockerignore": initSinatraDockerignore,
	"init/sinatra/Dockerfile": initSinatraDockerfile,
	"init/sinatra/docker-compose.yml": initSinatraDockerComposeYml,
	"init/unknown/.dockerignore": initUnknownDockerignore,
	"init/unknown/Dockerfile": initUnknownDockerfile,
	"init/unknown/docker-compose.yml": initUnknownDockerComposeYml,
	"templates.go": templatesGo,
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
	"init": &bintree{nil, map[string]*bintree{
		"rails": &bintree{nil, map[string]*bintree{
			".dockerignore": &bintree{initRailsDockerignore, map[string]*bintree{}},
			"Dockerfile": &bintree{initRailsDockerfile, map[string]*bintree{}},
			"docker-compose.yml": &bintree{initRailsDockerComposeYml, map[string]*bintree{}},
		}},
		"ruby": &bintree{nil, map[string]*bintree{
			".dockerignore": &bintree{initRubyDockerignore, map[string]*bintree{}},
			"Dockerfile": &bintree{initRubyDockerfile, map[string]*bintree{}},
			"docker-compose.yml": &bintree{initRubyDockerComposeYml, map[string]*bintree{}},
		}},
		"sinatra": &bintree{nil, map[string]*bintree{
			".dockerignore": &bintree{initSinatraDockerignore, map[string]*bintree{}},
			"Dockerfile": &bintree{initSinatraDockerfile, map[string]*bintree{}},
			"docker-compose.yml": &bintree{initSinatraDockerComposeYml, map[string]*bintree{}},
		}},
		"unknown": &bintree{nil, map[string]*bintree{
			".dockerignore": &bintree{initUnknownDockerignore, map[string]*bintree{}},
			"Dockerfile": &bintree{initUnknownDockerfile, map[string]*bintree{}},
			"docker-compose.yml": &bintree{initUnknownDockerComposeYml, map[string]*bintree{}},
		}},
	}},
	"templates.go": &bintree{templatesGo, map[string]*bintree{}},
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

