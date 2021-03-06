// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/neuron-device-plugin.yaml (2.258kB)
// assets/nvidia-device-plugin.yaml (2.361kB)
// assets/vpc-admission-webhook-config.yaml (524B)
// assets/vpc-admission-webhook-csr.yaml (234B)
// assets/vpc-admission-webhook-dep.yaml (1.106kB)
// assets/vpc-admission-webhook.yaml (231B)
// assets/vpc-resource-controller-dep.yaml (1.11kB)
// assets/vpc-resource-controller.yaml (565B)

package addons

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

var _neuronDevicePluginYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x56\xc1\x6e\xe3\x36\x10\xbd\xfb\x2b\x06\x5e\xe4\x56\x49\xcd\xb6\x28\x0a\xf5\x14\xec\xe6\x10\x20\x9b\x1a\x75\xdb\x4b\xd1\xc3\x98\x1c\x4b\x84\x29\x0e\xcb\x19\x39\x76\xbf\xbe\xa0\xac\x78\x25\xc7\x48\xd3\x73\x7d\x31\x41\xbe\x37\x6f\xf8\x38\x33\xd0\x07\x68\x55\xa3\xd4\x55\xb5\xeb\x37\x94\x02\x29\x49\xe9\xb8\xb2\x6c\xa4\x32\x1c\x0c\x45\x95\x8a\x0e\x4a\xc1\x16\x5f\x21\x95\xe1\x2e\xf6\x4a\x85\x28\x27\x6c\xa8\x08\xa4\x95\xa5\xbd\x33\x54\x44\xdf\x37\x2e\x48\xb5\xc0\xe8\x7e\xa7\x24\x8e\x43\x0d\x18\xa3\x54\xfb\xdb\xc5\xce\x05\x5b\xc3\x67\xa4\x8e\xc3\x9a\x74\xd1\x91\xa2\x45\xc5\x7a\x01\x10\xb0\xa3\x1a\x02\xf5\x89\x43\x31\x0b\x56\xd8\x81\x20\xa4\x23\x4c\x22\x1a\xaa\x21\x27\x54\xc8\x51\x94\xba\x85\x44\x32\x39\x8a\x90\x27\xa3\x9c\xf2\x1a\xa0\x43\x35\xed\x23\x6e\xc8\xcb\x69\xe3\x6d\x19\x59\x00\xf4\xd1\xa2\xd2\x5a\x13\x2a\x35\xc7\x13\x4b\x8f\x91\x6a\xf8\x85\xbd\x77\xa1\xf9\x6d\x00\x2c\x00\x94\xba\xe8\x51\x69\x94\x9a\x5c\x25\xff\x30\x04\x56\x54\xc7\xe1\x2c\x0d\x20\xa6\x25\xdb\x7b\x4a\x25\xfa\xd8\x62\x39\x77\xdd\x24\xa7\xce\xa0\x2f\x22\xdb\x1a\x96\xcb\x91\xe6\x67\xf9\xff\xfb\x0d\x00\x5e\xcc\x18\x72\x67\x4f\xe9\x32\x8f\x02\x76\x74\xac\xe1\xd3\x28\x78\x67\x2d\x07\xf9\x39\xf8\xe3\x19\x01\xc0\x31\xf3\x38\xd5\x70\x7f\x70\xa2\x72\x49\xc6\x67\x29\xb1\xc3\xbf\x39\x94\x86\xbb\xea\x94\xcf\x7b\xf8\x00\xb4\xdd\x92\xd1\x1a\x9e\x78\x3d\x1a\x32\x1e\x7e\x80\x2f\x98\x76\xa0\xad\x13\x88\x6c\x01\x05\x10\x5e\x6c\x01\xb4\xb6\xe0\xf0\x13\x3c\xb7\x14\x80\x02\x6e\x3c\xd9\x6f\x40\x5b\xba\x84\x9c\xa3\x9d\xfd\x86\x44\x42\x69\x4f\x92\x17\xdc\x27\x43\x02\x5b\x4e\x97\xc4\x2c\x2a\x20\x0c\xda\xa2\xe6\xc8\x47\x30\xf8\x35\xdc\x86\x32\x7d\x8c\x69\x01\xb7\x4a\x09\x10\xb6\xe8\x7c\x9f\xa8\x3c\xe3\xd6\x44\x6f\x35\x96\xa2\xec\xa4\x42\xdb\xb9\xe0\x44\x29\x15\xc6\xf7\xf9\xbf\x6a\x7a\x4c\x18\x94\xc8\x16\xa3\x8a\x0b\x4d\x71\x2e\x0b\xcc\xcf\x94\x8b\x43\xaa\x51\x2a\x26\xc7\xc9\xe9\xf1\x93\x47\x91\xa7\xa1\x2e\x96\xa7\x86\x28\x02\x5b\x3a\x53\x5f\x4a\x09\xb7\x5b\x17\x9c\x1e\x27\xc5\xc4\x96\xee\x5e\xed\x02\x24\xfa\xab\x77\x89\xec\xe7\x3e\xb9\xd0\xac\xcf\xd9\x3c\x34\x81\xcf\xdb\xf7\x07\x32\x7d\x2e\xad\x29\xf3\x14\x73\x3d\xf6\xe1\xaf\x94\x3a\x99\x1f\xe7\x0a\x1a\x1a\xf3\xfe\x10\x13\x89\xcc\x4b\x73\x8a\x1a\xea\x6c\xb9\x21\xbd\xec\x14\x17\x44\x31\x18\x2a\x72\x67\x2e\xaf\x70\xa7\xe5\xf7\x10\xae\x02\xf6\xe8\x7b\xba\x2a\x7c\x12\x77\x61\x7b\x5b\x1e\x3c\xa6\x86\xde\xc6\x7c\x7c\x0f\xe8\x87\x77\x45\xfa\xfe\x2a\xea\xbf\xd9\x95\xdd\xff\x5f\xda\x65\x38\x28\xba\x40\x69\x36\xe8\x5c\x87\x4d\x6e\x8b\x1b\x29\xed\x2e\x95\x64\x52\x79\x23\xe5\x8d\x54\xd7\x06\x68\x7d\x5b\x7e\x57\x7e\x2c\xbf\x9d\x7a\x34\x04\x58\xf5\xde\xaf\xd8\x3b\x73\xac\xe1\x61\xfb\xc4\xba\xca\x03\x25\xe8\x04\x77\x1a\xcb\xbb\x1f\xa5\xb8\x3a\x9a\x8d\xa6\x09\x58\xc8\xf4\x43\xe3\x72\x50\x3a\xe8\xdc\x56\xf4\x9e\x9f\x57\xc9\xed\x9d\xa7\x86\xee\xc5\xa0\x1f\x06\x78\x0d\x5b\xf4\x32\x37\xc5\x60\xc4\x8d\xf3\x4e\xdd\xeb\xc7\xb1\x89\x63\x0d\x7f\x2c\xef\x1e\x1f\x97\x7f\x4e\xce\xf6\xec\xfb\x8e\xbe\x70\x1f\xf4\x82\x53\x8c\xb7\x98\xa5\x7e\x11\xb5\xcb\xbc\x15\x6a\x5b\x43\xb5\xc7\x54\x79\xb7\x19\x86\x9c\x7f\xf5\x0d\xb0\x98\xca\xcd\x1e\xe5\x6d\x95\x96\xe5\x24\x30\x53\x8e\xef\x92\xfc\x27\x00\x00\xff\xff\xa3\x02\xd5\x4d\xd2\x08\x00\x00")

func neuronDevicePluginYamlBytes() ([]byte, error) {
	return bindataRead(
		_neuronDevicePluginYaml,
		"neuron-device-plugin.yaml",
	)
}

func neuronDevicePluginYaml() (*asset, error) {
	bytes, err := neuronDevicePluginYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "neuron-device-plugin.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xbc, 0x78, 0x25, 0x23, 0x27, 0x5f, 0x54, 0x5c, 0x5e, 0x69, 0xa0, 0x82, 0xb0, 0xd6, 0xca, 0x32, 0xb9, 0x1c, 0x55, 0x9c, 0x4, 0x6, 0xff, 0xf5, 0x9e, 0x2, 0xf3, 0x78, 0x45, 0x56, 0x1c, 0x4b}}
	return a, nil
}

var _nvidiaDevicePluginYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x55\x41\x8f\xdb\x36\x13\xbd\xeb\x57\x3c\xd8\x97\x04\x58\x49\x9b\x5c\xbe\xaf\x0a\x7a\x70\x77\xb7\xa8\x91\x8d\x1d\xac\x37\x09\x82\xa2\x87\x31\x39\x96\x08\x53\x24\x4b\x52\xf6\xea\xdf\x17\x94\xb5\x5e\x2b\x05\x9a\x9c\x0a\x54\x17\x1b\xd2\x9b\x99\x37\x6f\xde\x90\x73\xdc\x58\xd7\x7b\x55\x37\x11\xaf\xc4\x6b\xbc\xbd\x7e\xf3\xd3\x15\x56\x9f\x97\xb7\xcb\x05\x6e\xd6\x0f\x1f\xd7\x0f\x8b\xc7\xe5\x7a\x55\x00\x0b\xad\x31\x00\x03\x3c\x07\xf6\x07\x96\x45\x36\xcf\xe6\xb8\x57\x82\x4d\x60\x89\xce\x48\xf6\x88\x0d\x63\xe1\x48\x34\xfc\xfc\xe5\x0a\x9f\xd9\x07\x65\x0d\xde\x16\xd7\x78\x95\x00\xb3\xf1\xd3\xec\xf5\xbb\x6c\x8e\xde\x76\x68\xa9\x87\xb1\x11\x5d\x60\xc4\x46\x05\xec\x94\x66\xf0\x93\x60\x17\xa1\x0c\x84\x6d\x9d\x56\x64\x04\xe3\xa8\x62\x33\x94\x19\x93\x14\xd9\x1c\x5f\xc7\x14\x76\x1b\x49\x19\x10\x84\x75\x3d\xec\xee\x12\x07\x8a\x03\xe1\xf4\x34\x31\xba\xaa\x2c\x8f\xc7\x63\x41\x03\xd9\xc2\xfa\xba\xd4\x27\x60\x28\xef\x97\x37\x77\xab\xcd\x5d\xfe\xb6\xb8\x1e\x42\x3e\x19\xcd\x21\x35\xfe\x67\xa7\x3c\x4b\x6c\x7b\x90\x73\x5a\x09\xda\x6a\x86\xa6\x23\xac\x07\xd5\x9e\x59\x22\xda\xc4\xf7\xe8\x55\x54\xa6\xbe\x42\xb0\xbb\x78\x24\xcf\xd9\x1c\x52\x85\xe8\xd5\xb6\x8b\x13\xb1\x9e\xd9\xa9\x30\x01\x58\x03\x32\x98\x2d\x36\x58\x6e\x66\xf8\x65\xb1\x59\x6e\xae\xb2\x39\xbe\x2c\x1f\x7f\x5b\x7f\x7a\xc4\x97\xc5\xc3\xc3\x62\xf5\xb8\xbc\xdb\x60\xfd\x80\x9b\xf5\xea\x76\x99\x06\xb5\xc1\xfa\x57\x2c\x56\x5f\xf1\x7e\xb9\xba\xbd\x02\xab\xd8\xb0\x07\x3f\x39\x9f\xf8\x5b\x0f\x95\x64\x1c\x46\x87\x0d\xf3\x84\xc0\xce\x9e\x08\x05\xc7\x42\xed\x94\x80\x26\x53\x77\x54\x33\x6a\x7b\x60\x6f\x94\xa9\xe1\xd8\xb7\x2a\xa4\x61\x06\x90\x91\xd9\x1c\x5a\xb5\x2a\x52\x1c\xde\xfc\xad\xa9\x22\xcb\xc8\xa9\x71\xfc\x55\xd2\x2c\x94\x87\x37\xd9\x5e\x19\x59\xe1\x96\xb8\xb5\x66\xc3\x31\x6b\x39\x92\xa4\x48\x55\x06\x18\x6a\xb9\x82\x39\x28\xa9\x28\x97\x7c\x50\x82\x73\xa7\xbb\x5a\x99\x5c\x0e\x01\x81\xe3\x08\x0b\x8e\x04\x57\xd8\x77\x5b\xce\x43\x1f\x22\xb7\x59\xe2\x9e\xb2\x04\xd6\x2c\xa2\xf5\xe9\x3f\xd0\x52\x14\xcd\x3d\x6d\x59\x87\xd3\x8b\x7f\x2e\x13\x32\xa0\x73\x92\x22\x6f\xa2\xa7\xc8\x75\x7f\x8a\x8a\xbd\xe3\x0a\x0f\x56\x6b\x65\xea\x4f\x03\x20\x03\x22\xb7\x4e\x53\xe4\xb1\xd4\x45\x2b\xe9\x99\xe3\x31\xb9\x99\x8c\xb1\x27\x95\x86\x39\xb3\xf3\x2c\x28\xb2\x2c\xf0\x3e\x19\xbc\x61\x7f\xd2\x7f\x4b\x62\x7f\x24\x2f\x07\xbf\x53\x54\x5b\xa5\x55\xec\xcf\xb9\xd2\xc8\x92\x75\x43\x55\x96\xa9\x6d\x6f\x38\x72\x28\x94\x2d\xa5\x15\xa1\x8c\x14\xf6\xa1\x24\xd9\x2a\xa3\x42\x64\x9f\x0b\xdd\xa5\xdf\xb2\xee\xc8\x93\x89\xcc\x32\x0f\xa2\x61\xd9\xa5\x0e\x72\x91\x3c\x2a\x48\xe7\x24\xa5\x35\xb9\xb3\x32\x94\x63\xa9\x17\xbe\x67\xc5\x80\x31\x94\x7d\x41\xda\x35\x54\x4c\x19\x9c\xb3\x39\x2b\x2b\xcc\x66\x63\x98\x9e\xc8\xfe\x7d\xe1\x81\xe7\x19\x0e\x92\x5b\xcd\x7e\xca\x63\x54\xf4\xe5\xcb\x7f\x43\xd1\x1c\x7b\xee\x2b\xdc\x8c\x88\x45\x02\x84\xb5\xd1\xfd\x59\x19\xeb\x52\x43\xd6\x57\xb8\x7b\x52\x21\x86\x69\xe0\x49\xb1\x42\xd8\xb6\xac\x5d\xf7\xbd\x20\x80\x77\x3b\x16\xb1\xc2\xca\x6e\xc6\xb1\x9d\x7b\xfe\x40\x7e\x7f\x3a\x64\x9d\x95\xa0\x90\x8e\xcb\x91\x16\x48\xca\xdc\x9a\x77\x38\x36\x6c\xc0\x26\x9d\x6f\xf2\x6a\x58\xe9\x6f\x20\xe7\x6c\x67\x57\x3c\xdf\x0b\xc3\x05\x61\x3b\x2f\x38\x0c\x13\xf8\x26\x30\x15\x0d\x08\x16\xb1\xa1\x98\x32\xf7\x10\xf4\x92\x6e\xcb\x29\x7c\xcc\x29\x41\xbb\xc8\x1e\x84\x1d\x29\xdd\x79\x2e\xfe\xfd\xc1\x39\xaf\xac\x57\xb1\xbf\xd1\x14\xc2\x6a\x70\xef\xec\x74\xda\xe4\xc6\x4a\x3e\x87\x3e\x1b\x5e\x58\x93\xee\x20\xf6\x67\xc7\xe6\x50\x2d\xd5\x67\xd7\x97\xfb\xff\x87\xa9\xf3\xab\xc3\x75\xf1\xbf\xe2\xfa\x47\x76\x44\x44\x7f\x86\x91\xaf\x43\x85\xdf\x67\x79\x9e\xd4\xc9\xad\xc9\x95\x51\x31\x67\xef\xad\xff\x79\x47\x3a\xf0\xec\x8f\x97\xe5\x65\xd1\x0d\x6d\x58\x13\xf9\x29\xbe\x2c\x24\x40\x5a\xdb\xe3\x47\xaf\x0e\x4a\x73\xcd\x77\x41\x90\x1e\x16\xab\xc2\x90\xe4\x02\x29\xc8\xd1\xb0\x40\x8a\xc3\x65\x06\x40\x7a\xeb\x12\x97\xc5\xfd\xfd\x45\xd1\x83\xd5\x5d\xcb\x1f\x6c\x67\xe2\x04\x9f\x8f\x2d\x4e\x7a\x9b\xe4\x6b\x53\xcc\x47\x8a\x4d\x85\xf2\x40\xbe\xd4\x6a\x3b\x8c\x59\x73\x2c\x27\x51\xcf\x8e\x3f\x95\xba\xa8\xf2\xbd\x1a\x8d\x0d\xa7\x02\x93\xba\xee\x87\x4a\x66\x7f\x05\x00\x00\xff\xff\x31\x6e\xa1\x29\x39\x09\x00\x00")

func nvidiaDevicePluginYamlBytes() ([]byte, error) {
	return bindataRead(
		_nvidiaDevicePluginYaml,
		"nvidia-device-plugin.yaml",
	)
}

func nvidiaDevicePluginYaml() (*asset, error) {
	bytes, err := nvidiaDevicePluginYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nvidia-device-plugin.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc4, 0x9d, 0xa, 0xb7, 0x22, 0x67, 0xb3, 0x61, 0xf5, 0x37, 0x17, 0x90, 0x53, 0x59, 0x85, 0x25, 0xc7, 0x1, 0x1d, 0x96, 0x1, 0x51, 0xd5, 0x17, 0xea, 0x87, 0x44, 0x28, 0xc8, 0x3a, 0x53, 0xfc}}
	return a, nil
}

var _vpcAdmissionWebhookConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x91\x4f\x6b\xf3\x30\x0c\xc6\xef\xf9\x14\x22\xf7\xa4\xf4\xf6\xe2\xdb\x4b\x29\x63\x87\xc1\x18\x63\x3b\x8c\x1d\x14\x47\x4d\x45\x62\xcb\x58\x76\x4a\xf7\xe9\x47\xfe\xb4\xac\xb0\xd5\x17\xdb\x7a\xa4\xdf\x63\xc9\x18\xf8\x8d\xa2\xb2\x78\x03\xd8\x3a\xd6\xe9\x18\xa9\x63\x4d\x11\x13\x8b\xaf\xfb\x7f\x5a\xb3\x6c\xc6\x6d\x43\x09\xb7\x45\xcf\xbe\x35\xf0\x94\x13\x26\xf6\xdd\x3b\x35\x47\x91\x7e\x27\xfe\xc0\x5d\x5e\x2a\x0a\x47\x09\x5b\x4c\x68\x0a\x00\x8f\x8e\x0c\x8c\xc1\x56\x57\x7a\x75\x5a\x8a\x2a\x7b\xe8\xd6\x0c\x0d\x68\xc9\x40\x9f\x1b\xaa\xf4\xac\x89\x5c\x01\x30\x60\x43\x83\x4e\x10\x00\x0c\xe1\x0f\x4a\xb1\xee\x73\x62\x75\xcf\xaf\x46\x87\x5f\xe2\xf1\xa4\xb5\x15\x37\x63\xed\xc0\xe4\xd3\xf2\xfa\xc5\x08\x40\x29\x8e\x6c\xe9\x72\xbd\xdb\xc2\x4d\xce\xaf\x4d\x2c\x2b\x60\x3a\x1a\x28\x37\x6e\x1a\x1b\x95\x73\x3c\xe6\x81\xf4\xe2\x52\x81\x04\x5a\xc6\xa7\x06\x3e\xa0\xdc\xbd\xec\xff\xbf\xee\x4b\xf8\xbc\x32\x30\xf0\x43\x94\x1c\x26\xbd\x2c\x6f\xe2\xeb\x0f\xce\xca\xb8\xfd\xa1\x45\x52\xc9\xd1\xd2\xac\x04\x69\x75\xd5\x0e\xc8\x43\x8e\xf4\x2c\x03\xdb\xb3\x81\xc7\xce\x4b\xa4\xe2\x3b\x00\x00\xff\xff\x49\xee\x9e\x02\x0c\x02\x00\x00")

func vpcAdmissionWebhookConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_vpcAdmissionWebhookConfigYaml,
		"vpc-admission-webhook-config.yaml",
	)
}

func vpcAdmissionWebhookConfigYaml() (*asset, error) {
	bytes, err := vpcAdmissionWebhookConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vpc-admission-webhook-config.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x7d, 0x11, 0x23, 0x17, 0x2f, 0xdc, 0x4e, 0x18, 0xaa, 0xc8, 0x66, 0xee, 0xf3, 0xc1, 0x85, 0x63, 0xb1, 0xe3, 0x53, 0x57, 0x80, 0x96, 0xe6, 0x54, 0x26, 0x46, 0x6b, 0x3f, 0x17, 0x20, 0x31, 0x8}}
	return a, nil
}

var _vpcAdmissionWebhookCsrYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8d\x4b\x4e\xc4\x30\x10\x44\xf7\x3e\x45\x5f\x20\x41\xb3\x43\xde\x72\x03\x90\xd8\x77\xec\xc2\x69\x19\x7f\x70\xb7\x83\xe6\xf6\x28\x13\xa4\xd9\x95\xaa\x9e\x5e\x71\x97\x4f\x0c\x95\x56\x3d\x05\x0c\x93\x2f\x09\x6c\xd0\x35\xbf\xea\x2a\xed\xe5\xb8\x6d\x30\xbe\xb9\x2c\x35\x7a\x7a\x7b\x12\x1f\x92\xaa\xd4\xf4\x8e\x9f\x09\x35\x57\x60\x1c\xd9\xd8\x3b\xa2\xca\x05\x9e\x8e\x1e\x16\x8e\x45\xf4\x94\x2f\xbf\xd8\xf6\xd6\xf2\x9a\xe7\x86\x45\xef\x6a\x28\x4e\x3b\xc2\xc9\xa7\xd1\x66\xd7\x33\x2d\x74\x4d\x9e\xa7\xed\xa8\xf6\x78\x8a\x8e\x68\x2a\x27\xfc\x23\x51\x92\x18\x7f\x93\x4a\xaa\x6c\x73\xe0\xd1\x66\xdc\x09\x35\x48\xdf\x31\x0a\xaa\x5d\x36\x8c\x03\x83\x4e\x9b\xfb\x0b\x00\x00\xff\xff\xf1\x7d\x42\x97\xea\x00\x00\x00")

func vpcAdmissionWebhookCsrYamlBytes() ([]byte, error) {
	return bindataRead(
		_vpcAdmissionWebhookCsrYaml,
		"vpc-admission-webhook-csr.yaml",
	)
}

func vpcAdmissionWebhookCsrYaml() (*asset, error) {
	bytes, err := vpcAdmissionWebhookCsrYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vpc-admission-webhook-csr.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x56, 0xf4, 0x54, 0x57, 0xf8, 0xbd, 0x8a, 0x1a, 0x67, 0xb2, 0x29, 0x18, 0xe2, 0x44, 0x8a, 0xc2, 0x7f, 0x44, 0x26, 0x4c, 0x52, 0xe0, 0x70, 0xe0, 0xc8, 0x5a, 0x74, 0x76, 0x2, 0xcd, 0x3e, 0xa}}
	return a, nil
}

var _vpcAdmissionWebhookDepYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x93\xbf\x8e\xdb\x30\x0c\xc6\xf7\x3c\x05\x97\xf4\x26\xdb\xcd\xe1\xd0\x41\x40\x0a\x14\x2d\xba\xb4\xbd\x04\x0d\xd0\x9d\x91\x89\x44\xb0\xfe\x41\xa4\x9d\xfa\xed\x0b\xc5\xbe\xc2\xf6\x25\x77\x1a\x3c\x90\xfc\x7d\x14\x3f\x53\x18\xcd\x1f\x4a\x6c\x82\x57\x80\x31\x72\xd5\x6d\x56\x8d\xf1\xb5\x82\x6f\x14\x6d\xe8\x1d\x79\x59\x39\x12\xac\x51\x50\xad\x00\x3c\x3a\x52\xd0\x45\x5d\x60\xed\x0c\x67\xb2\xb8\xd0\xf1\x1c\x42\x33\x66\x39\xa2\x26\x05\x4d\x7b\xa4\x82\x7b\x16\x72\x2b\x00\x8b\x47\xb2\x9c\x05\x20\xf7\xb9\xa7\xc0\x91\x74\x2e\x4a\x14\xad\xd1\xc8\x0a\x36\x2b\x00\x96\x84\x42\xa7\x7e\xc0\xa5\x8f\xa4\xe0\x37\xe9\x44\x28\x94\xd3\x64\x49\x4b\x48\x43\xda\xa1\xe8\xf3\xcf\x49\xbb\x37\x1b\x02\x08\xb9\x68\x51\x68\xa4\x27\xa3\xe6\x63\x67\x42\xef\x48\x01\xbc\xdc\x3f\x1f\x1d\xbc\xa0\xf1\x94\x26\x78\xf1\x8e\x7f\xff\xdb\xa4\xd3\x84\x1a\xc8\x42\x2c\x7f\xa5\x24\xdf\x8d\xa5\x6d\x45\xa2\xab\x91\xab\x34\x25\xe1\xeb\xb7\x8c\x57\xb7\x97\xd8\x0f\xea\xef\x51\x0d\xf5\xb7\xa0\xdd\xe1\x6a\xe1\x61\xb4\x76\xd7\x51\x4a\xa6\xa6\xed\xc5\xf8\x3a\x5c\x78\x59\x8e\x96\x83\x0d\x27\x09\x2c\x35\xa5\xb4\x4c\x77\xdb\xa7\x45\xe8\xf1\xf3\x87\xcd\x24\x64\x1c\x9e\x48\xc1\xc3\x9a\xcb\xba\x49\x25\xe9\x54\xae\xb9\x5c\x73\x45\x0d\x57\x37\xcd\x52\xdd\xc7\xf2\xb1\x7c\x7a\x58\x8a\xec\x5b\x6b\xf7\xc1\x1a\xdd\x2b\xf8\x62\x2f\xd8\x4f\xef\xda\x05\xdb\x3a\xfa\x15\x5a\x2f\xaf\xec\x1d\x7e\xcc\xa8\x5e\x5c\xcd\x99\x55\x00\xb8\xcc\xed\x51\xce\x0a\x5e\x1b\xb9\xa8\x4d\x84\xf5\xce\xdb\x5e\x81\xa4\x96\xc6\xe4\x39\xb0\x3c\x93\x5c\x42\x6a\x66\x71\x1f\x6a\x3a\xcc\x96\x38\x9f\x23\x09\x96\xf9\x15\x25\x4f\x42\x5c\x9a\x50\x05\x56\x60\x8d\x6f\xff\xbe\x55\x84\x49\x9f\x15\xa0\xab\x3f\xbd\x98\x3e\x8c\x7d\x63\x0d\xef\x4d\xcb\xf9\x71\xc9\xdc\xa1\x21\xf6\x7c\x7f\x7d\x47\x95\x7f\x01\x00\x00\xff\xff\xe1\x3a\x01\x42\x52\x04\x00\x00")

func vpcAdmissionWebhookDepYamlBytes() ([]byte, error) {
	return bindataRead(
		_vpcAdmissionWebhookDepYaml,
		"vpc-admission-webhook-dep.yaml",
	)
}

func vpcAdmissionWebhookDepYaml() (*asset, error) {
	bytes, err := vpcAdmissionWebhookDepYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vpc-admission-webhook-dep.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb9, 0xc6, 0x83, 0xd1, 0x8d, 0x3e, 0x9c, 0x6, 0xa7, 0x56, 0x5d, 0x1a, 0x1f, 0xf, 0x58, 0x39, 0x30, 0x3, 0xa4, 0x46, 0xe, 0x66, 0xb6, 0xb5, 0x9f, 0x8d, 0xa8, 0xba, 0x9e, 0x51, 0x67, 0xb6}}
	return a, nil
}

var _vpcAdmissionWebhookYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8c\x31\xae\xc2\x40\x0c\x05\xfb\x3d\x85\x2f\xe0\xe2\xeb\xa7\xf2\x29\x90\x90\xe8\x9d\xcd\x13\xac\x92\xcd\x5a\x6b\x13\xc4\xed\x51\x22\x0a\x1a\x44\x67\xf9\xcd\x0c\x33\x27\xb5\x72\x41\xf7\xd2\x56\xa1\xed\x2f\xcd\x65\x9d\x84\xce\xe8\x5b\xc9\x48\x15\xa1\x93\x86\x4a\x22\x5a\xb5\x42\x68\xb3\xcc\x3a\xd5\xe2\xbb\xc1\x0f\x8c\xb7\xd6\xe6\xf7\xea\xa6\x19\x42\xf3\x7d\x04\xfb\xd3\x03\x35\x11\x2d\x3a\x62\xf1\x3d\x40\xa4\x66\xdf\x0a\x6e\xc8\x3b\x64\xad\xc7\x41\xf3\x71\x0a\x0d\xc3\xff\xe1\x86\xf6\x2b\xe2\xf4\xf1\x73\x2c\xc8\xd1\xfa\xcf\xf6\x2b\x00\x00\xff\xff\xbc\xa7\x78\x3e\xe7\x00\x00\x00")

func vpcAdmissionWebhookYamlBytes() ([]byte, error) {
	return bindataRead(
		_vpcAdmissionWebhookYaml,
		"vpc-admission-webhook.yaml",
	)
}

func vpcAdmissionWebhookYaml() (*asset, error) {
	bytes, err := vpcAdmissionWebhookYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vpc-admission-webhook.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc0, 0x14, 0x3e, 0x68, 0x1c, 0x57, 0x26, 0x7e, 0x3b, 0xab, 0xf1, 0x55, 0x21, 0x61, 0xc3, 0xe1, 0xaf, 0x46, 0xb6, 0xf7, 0xdd, 0x11, 0x29, 0x41, 0x64, 0x57, 0xc8, 0xd4, 0xea, 0x97, 0xba, 0xff}}
	return a, nil
}

var _vpcResourceControllerDepYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x93\x41\x6b\xdc\x3e\x10\xc5\xef\xfe\x14\x73\x09\x39\xd9\xde\xcd\x3f\xff\x84\x0a\x7a\x08\x0d\xb4\x87\x52\x16\x12\x7a\x1f\x4b\x93\x58\x58\x96\xc4\xcc\xd8\x1b\xf7\xd3\x17\xb7\xd9\x5d\x9b\xa6\xd9\xce\xc9\xbc\x9f\xc6\xef\x8d\xd0\x94\x65\x59\x60\xf6\xdf\x89\xc5\xa7\x68\x00\x73\x96\x7a\xdc\x16\x9d\x8f\xce\xc0\x3d\xe5\x90\xa6\x9e\xa2\x16\x3d\x29\x3a\x54\x34\x05\x40\xc4\x9e\x0c\x8c\xd9\x96\x4c\x92\x06\xb6\x54\xda\x14\x95\x53\x08\xc4\xaf\x5c\x32\x5a\x32\xd0\x0d\x0d\x95\x32\x89\x52\x5f\x48\x26\x3b\xb7\x33\xe5\xe0\x2d\x8a\x81\x6d\x01\x20\x14\xc8\x6a\xe2\x99\x00\xf4\xa8\xb6\xfd\x8a\x0d\x05\xf9\x2d\xc0\x1c\xe9\x3d\xb3\xb9\xd4\x13\x1b\x68\xd0\x76\x14\xdd\x41\x63\xb4\x9d\x01\x51\x6c\x02\x15\x00\x4a\x7d\x0e\xa8\xf4\xea\xb3\x18\x67\xae\xb0\xb2\xfc\x27\xd3\xb7\x6d\xff\x34\x06\x38\x0c\xfe\xeb\x9b\x78\xf4\x96\xee\xac\x4d\x43\xd4\x73\x1e\xb3\x80\x3e\x12\x1f\xa3\x95\x60\x53\xdf\x63\x74\xa7\xac\x25\xd4\xe7\x92\x22\x3f\xcb\xb2\xa1\x14\x75\xc4\xac\x2d\x93\xb4\x29\xb8\x8f\x3e\x3e\xa5\x23\xf7\x3d\x3e\x93\x81\xcb\x0b\xa9\x5c\xc7\x15\x59\xae\x2e\xa4\xba\x90\x9a\x3a\xa9\xf7\x3e\xba\xb4\x97\xf2\x2f\x96\x66\xdc\x54\x57\xd5\xf5\xe5\xfa\x67\xbb\x21\x84\x5d\x0a\xde\x4e\x06\xee\xc2\x1e\x27\x39\xf2\xe0\x47\x8a\x24\xb2\xe3\xd4\xd0\x29\x23\xc0\x13\xfa\x30\x30\x3d\x1e\x32\x1a\xf8\x7f\x41\x5b\xd5\xfc\x99\x74\xd9\x00\xd0\x26\x51\x03\xdb\xab\xdb\x6a\x53\x6d\xaa\xed\x8a\x65\xd4\xd6\x40\xdd\x12\x06\x6d\x7f\xac\x51\x62\x35\x70\xb3\xbd\xbd\xfd\xb0\xd2\xc5\xb6\x34\xbf\xf4\x2f\x8f\x8f\xbb\x05\xf0\xd1\xab\xc7\x70\x4f\x01\xa7\x07\xb2\x29\x3a\x31\xf0\xdf\x66\x71\x22\x13\xfb\xe4\xde\x66\xea\x7b\x4a\x83\x1e\xe1\x69\xa8\x73\x6b\x75\x78\x40\x76\x60\xaf\xd3\xa7\x14\x95\x5e\x56\x17\x90\xd9\x8f\x3e\xd0\x33\x39\x03\xca\x03\x15\xa7\x5b\xf9\x46\xba\x4f\xdc\xad\xf4\x98\x1c\x3d\xac\xb6\x6f\xae\x86\x14\xab\x79\x6d\x39\x92\x92\x54\x3e\xd5\x49\x0c\x04\x1f\x87\x97\xf7\x0e\x21\xdb\xd6\x00\xf6\xee\xe6\xba\xf8\x19\x00\x00\xff\xff\xc0\x2f\xdb\x68\x56\x04\x00\x00")

func vpcResourceControllerDepYamlBytes() ([]byte, error) {
	return bindataRead(
		_vpcResourceControllerDepYaml,
		"vpc-resource-controller-dep.yaml",
	)
}

func vpcResourceControllerDepYaml() (*asset, error) {
	bytes, err := vpcResourceControllerDepYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vpc-resource-controller-dep.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3d, 0x55, 0x7, 0xe, 0xad, 0xaf, 0x80, 0xb9, 0xe, 0xa7, 0x8e, 0x5b, 0xaa, 0x37, 0x24, 0xda, 0x64, 0x6f, 0x9d, 0xa7, 0xb6, 0xb2, 0x58, 0x6b, 0xbb, 0x35, 0x95, 0x85, 0x39, 0xdd, 0x4d, 0xb8}}
	return a, nil
}

var _vpcResourceControllerYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x90\x31\x4f\x33\x31\x0c\x86\xf7\xfc\x8a\xa8\x7b\x5a\x7d\xdb\xa7\xdb\x80\x81\xbd\x48\xec\x3e\x9f\xdb\x9a\xe6\xe2\xc8\x76\x0e\xc1\xaf\x47\x77\xd7\xc2\x80\x44\x85\x98\xf2\xc4\xf6\x6b\x4b\x4f\x4a\x29\x40\xe5\x67\x52\x63\x29\x5d\xd4\x1e\x70\x0b\xcd\x4f\xa2\xfc\x0e\xce\x52\xb6\xe7\xff\xb6\x65\xd9\x4d\xff\xc2\x99\xcb\xd0\xc5\x87\xdc\xcc\x49\xf7\x92\x29\x8c\xe4\x30\x80\x43\x17\x62\x2c\x30\x52\x17\xa7\x8a\x49\xc9\xa4\x29\x52\x42\x29\xae\x92\x33\x69\xd0\x96\xc9\xba\x90\x22\x54\x7e\x54\x69\xd5\xe6\x4c\x8a\x9b\x4d\x88\xf1\x1a\xb8\xd4\x8a\x0c\x64\x5f\xb4\x33\x07\x6f\x6b\xa1\xca\xb0\x02\x4a\x39\xf0\x71\x84\x3a\x7f\x27\xd2\xfe\x92\x6d\x75\x00\xa7\x05\x8f\xe4\xcb\x9b\xd9\x56\x78\x05\xc7\xd3\xba\xe6\x93\x50\x69\x9e\xff\x9b\x87\x7b\x2e\x03\x97\xe3\x6f\x74\x48\xa6\x3d\x1d\xe6\xc1\xab\x90\x1f\x8e\x86\x18\xbf\xbb\xbf\x75\xc2\x5a\xff\x42\xe8\x8b\xf4\x35\xfd\x44\x3a\x31\xd2\x1d\xa2\xb4\xe2\x37\x17\xac\x7d\xab\x80\xd4\xc5\x73\xeb\x29\xd9\x9b\x39\x8d\xe1\x23\x00\x00\xff\xff\x83\x2e\x8d\x64\x35\x02\x00\x00")

func vpcResourceControllerYamlBytes() ([]byte, error) {
	return bindataRead(
		_vpcResourceControllerYaml,
		"vpc-resource-controller.yaml",
	)
}

func vpcResourceControllerYaml() (*asset, error) {
	bytes, err := vpcResourceControllerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vpc-resource-controller.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa8, 0xff, 0xd2, 0xc, 0x46, 0xfb, 0xe4, 0x4e, 0xb, 0x26, 0x4a, 0xa0, 0x7d, 0x73, 0x5e, 0xaf, 0x6d, 0xab, 0xfd, 0xe5, 0xc2, 0x4f, 0xed, 0xae, 0xb6, 0xd7, 0x17, 0x38, 0x68, 0xb0, 0xe8, 0x2d}}
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
	"neuron-device-plugin.yaml":         neuronDevicePluginYaml,
	"nvidia-device-plugin.yaml":         nvidiaDevicePluginYaml,
	"vpc-admission-webhook-config.yaml": vpcAdmissionWebhookConfigYaml,
	"vpc-admission-webhook-csr.yaml":    vpcAdmissionWebhookCsrYaml,
	"vpc-admission-webhook-dep.yaml":    vpcAdmissionWebhookDepYaml,
	"vpc-admission-webhook.yaml":        vpcAdmissionWebhookYaml,
	"vpc-resource-controller-dep.yaml":  vpcResourceControllerDepYaml,
	"vpc-resource-controller.yaml":      vpcResourceControllerYaml,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

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
	"neuron-device-plugin.yaml": {neuronDevicePluginYaml, map[string]*bintree{}},
	"nvidia-device-plugin.yaml": {nvidiaDevicePluginYaml, map[string]*bintree{}},
	"vpc-admission-webhook-config.yaml": {vpcAdmissionWebhookConfigYaml, map[string]*bintree{}},
	"vpc-admission-webhook-csr.yaml": {vpcAdmissionWebhookCsrYaml, map[string]*bintree{}},
	"vpc-admission-webhook-dep.yaml": {vpcAdmissionWebhookDepYaml, map[string]*bintree{}},
	"vpc-admission-webhook.yaml": {vpcAdmissionWebhookYaml, map[string]*bintree{}},
	"vpc-resource-controller-dep.yaml": {vpcResourceControllerDepYaml, map[string]*bintree{}},
	"vpc-resource-controller.yaml": {vpcResourceControllerYaml, map[string]*bintree{}},
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
