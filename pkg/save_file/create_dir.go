package save_file

import "os"

type createDirPath struct {
	Path string
}

func (p *createDirPath) createDir() error {
	err := os.MkdirAll(p.Path, os.ModePerm)
	if err != nil {
		return err
	}
	return err
}
