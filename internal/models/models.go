package models

type SaveFileData struct {
	Id         string
	IsRequest  bool
	FileName   string
	StringData string
}

type Credentials struct {
	Id         string
	BaseUrl    string
	Method     string
	Headers    map[string]string
	GetParams  map[string]string
	PostFields []byte
}
