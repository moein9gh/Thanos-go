package main

import (
	"os"
	"path/filepath"

	"github.com/thanos-go/cmd"
	"github.com/thanos-go/config"

	_ "github.com/thanos-go/docs"
)

// @title         Thanos
// @version       1.0
// @description   This is Thanos  api docs.
// @contact.name  API Support
// @license.name  Apache 2.0
// @host          185.255.88.17:8095
// @BasePath      /v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name authorization
func main() {
	cmd.Execute()
}

func init() {
	cfg := config.Get()
	pathList := []string{cfg.Static.StaticFilePath, cfg.Static.CharacterFilePath}

	err := createStaticFolders(pathList)
	if err != nil {
		panic(err)
	}
}

func createStaticFolders(pathList []string) error {
	for _, fp := range pathList {
		if _, err := os.Stat(fp); os.IsNotExist(err) {
			if err := os.MkdirAll(filepath.Dir(fp), 0755); err != nil {
				return err
			}
		}
	}

	return nil
}
