package gen

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func configGen(s string) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	abs, err := filepath.Abs(before(dir, "src/"))
	var b []byte
	var b1 []byte
	if abs == dir {
		b, err = ioutil.ReadFile(abs + "/github.com/AliasYermukanov/gkgen/templates/config.gotxt") // just pass the file name
		b1, err = ioutil.ReadFile(abs + "/github.com/AliasYermukanov/gkgen/templates/jsonConfig.txt") // just pass the file name
	} else {
		b, err = ioutil.ReadFile(abs + "/src/github.com/AliasYermukanov/gkgen/templates/config.gotxt") // just pass the file name
		b1, err = ioutil.ReadFile(abs + "/src/github.com/AliasYermukanov/gkgen/templates/jsonConfig.txt") // just pass the file name
	}
	if err != nil {
		fmt.Print(err)
	}

	f, _ := os.Create(s + "-api/src/config/config.go")
	_, _ = f.Write([]byte(string(b)))
	f, _ = os.Create(s + "-api/src/config/config.json")
	_, _ = f.Write([]byte(string(b1)))
	_ = f.Close()
}
