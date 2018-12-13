package assets

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Configsec27b42f4e737cebcd9245ac050bb73926b63135 = "driver: mysql\nprotocol: tcp(gogollellero_db:3306)\ndevelopment:\n  user: miyanokomiya\n  password: miyanokomiya\n  db: gogollellero\ntest:\n  user: miyanokomiya\n  password: miyanokomiya\n  db: gogollellero_test\n"

// Configs returns go-assets FileSystem
var Configs = assets.NewFileSystem(map[string][]string{"/configs": []string{"db.yml"}, "/": []string{"configs"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1543937523, 1543937523000000000),
		Data:     nil,
	}, "/configs": &assets.File{
		Path:     "/configs",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1543749886, 1543749886000000000),
		Data:     nil,
	}, "/configs/db.yml": &assets.File{
		Path:     "/configs/db.yml",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1543750272, 1543750272000000000),
		Data:     []byte(_Configsec27b42f4e737cebcd9245ac050bb73926b63135),
	}}, "")
