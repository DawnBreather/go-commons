package file

import (
	"commons/env_var"
	"commons/logger"
	"commons/placeholder"
	"encoding/base64"
	"io/ioutil"
)

var _logger = logger.New()

type File struct{
	path    string
	content []byte
	base64  string
}

func (f File) GetPath() string{
	return f.path
}
func (f File) GetContent() []byte{
	return f.content
}
func (f File) GetBase64() string{
	return f.base64
}

func (f *File) SetPath(path string) *File{
	f.path = path
	return f
}
func (f *File) SetContent(content []byte) *File{
	f.content = content
	return f
}
func (f *File) SetBase64(b64string string) *File{
	f.base64 = b64string
	return f
}


func(f *File) ParseContentToBase64() *File{
	f.base64 = base64.StdEncoding.EncodeToString(f.content)
	return f
}

func(f *File) ParseBase64ToContent() *File{
	var err error
	f.content, err = base64.StdEncoding.DecodeString(f.base64)
	if err != nil {
		_logger.Errorf("Unable to decode base64 encoded string into file content: %v", err)
	}
	return f
}

func(f *File) ReadContent() *File {
	var err error
	f.content, err = ioutil.ReadFile(f.path)
	if err != nil {
		_logger.Errorf("Error reading file {%s}: %v", f.path, err)
	}
	return f
}

func (f *File) SaveTo(dstPath string) *File{
	var err error

	if dstPath != "" {
		err = ioutil.WriteFile(dstPath, f.content, 0644)
		if err != nil {
			_logger.Errorf("Unable to save file content into {%s}: %v", dstPath, err)
		}
		if f.path == "" {
			f.path = dstPath
		}
	} else {
		_logger.Errorf("Unable to save file content: %v", "no destination path provided")
	}

	return f
}

func (f *File) Save() *File{
	return f.SaveTo(f.path)
}

func (f *File) ReplaceEnvVarsPlaceholder(prefix, suffix string) *File{
	var p = placeholder.New(prefix, suffix, env_var.REGEX_MASK)
	f.content = []byte(p.ReplacePlaceholdersWithEnvVars(string(f.content)))

	return f
}