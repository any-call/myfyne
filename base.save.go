package myfyne

import (
	"encoding/json"
	"os"
	"path"

	"fyne.io/fyne/v2"
)

func GetLocFile(app fyne.App, fileName string) string {
	storeLoc := app.Storage().RootURI()
	return path.Join(storeLoc.Path(), fileName)
}

func LoadFromLocFile[T any](app fyne.App, fileName string) (*T, error) {
	data, err := os.ReadFile(GetLocFile(app, fileName))
	if err != nil {
		return nil, err
	}
	var creds T
	err = json.Unmarshal(data, &creds)
	return &creds, err
}

func SaveToLocFile(app fyne.App, fileName string, obj any) error {
	// 示例中未加密，生产环境应加密 password
	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return os.WriteFile(GetLocFile(app, fileName), data, 0600)
}
