package main

import (
	"flag"
	"os"
	"strings"

	"github.com/korylprince/crudgen"
)

type Config struct {
	Templates  []string
	Structs    []string
	InputPath  string
	OutputPath string
}

var config *Config

func init() {
	config = new(Config)

	tmpls := flag.String("tmpls", "", "Comma-separated list of templates")
	structs := flag.String("structs", "", "Comma-separated list of struct names")
	flag.StringVar(&(config.OutputPath), "o", "crud.go", "Output path")
	flag.Parse()

	config.Templates = strings.Split(*tmpls, ",")
	for i, v := range config.Templates {
		config.Templates[i] = strings.TrimSpace(v)
	}
	config.Structs = strings.Split(*structs, ",")
	for i, v := range config.Structs {
		config.Structs[i] = strings.TrimSpace(v)
	}

	config.InputPath = flag.Arg(0)
}

func main() {

	var structs []*crudgen.Struct

	for _, n := range config.Structs {
		s, err := crudgen.Parse(config.InputPath, n)
		if err != nil {
			panic(err)
		}
		structs = append(structs, s)
	}

	src, err := crudgen.Generate(config.Templates, structs)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(config.OutputPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.Write(src)
	if err != nil {
		panic(err)
	}
}
