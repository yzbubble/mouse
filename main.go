package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/russross/blackfriday.v2"
)

var config = struct {
	root         string
	addr         string
	templatePath string
	indexPath    string
}{
	root:         ".",
	addr:         ":8080",
	templatePath: "",
	indexPath:    "index.md",
}

func main() {
	loadEnvConfig()
	loadFlagConfig()
	log.Printf("Mouse listen and serve on %q ...", config.addr)
	if err := http.ListenAndServe(config.addr, http.HandlerFunc(render)); err != nil {
		log.Fatalln(err)
	}
}

func loadEnvConfig() {
	if v, ok := os.LookupEnv("mouse_root"); ok {
		config.root = v
	}
	if v, ok := os.LookupEnv("mouse_addr"); ok {
		config.addr = v
	}
	if v, ok := os.LookupEnv("mouse_template_path"); ok {
		config.templatePath = v
	}
	if v, ok := os.LookupEnv("mouse_index_path"); ok {
		config.indexPath = v
	}
}

func loadFlagConfig() {
	root := flag.String("root", ".", "resources root path")
	addr := flag.String("addr", ":8080", "http addr")
	templatePath := flag.String("template", "", "template file path")
	indexPath := flag.String("index", "index.md", "default index file path")
	flag.Parse()
	if *root != "" {
		config.root = *root
	}
	if *addr != "" {
		config.addr = *addr
	}
	if *templatePath != "" {
		if *templatePath == "nil" {
			config.templatePath = ""
		} else {
			config.templatePath = *templatePath
		}
	}
	templateFlag := flag.Lookup("template")
	if templateFlag == nil {
		log.Println("templateFlag is nil")
	}
	if *indexPath != "" {
		config.indexPath = *indexPath
	}
}

func render(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "" {
		path = "/" + config.indexPath
	}
	if strings.HasSuffix(path, "/") {
		path += config.indexPath
	}
	if filepath.Ext(path) == "" {
		path += ".md"
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	path = config.root + path
	fileName := filepath.Base(path)
	log.Printf("[debug]: url path map: %q => %q\n", r.URL.String(), path)
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ext := filepath.Ext(path)
	switch ext {
	case "", ".md":
		w.Header().Set("Content-Type", "text/html")
		context := struct {
			Title   string
			Content template.HTML
		}{
			Title:   fileName,
			Content: template.HTML(string(blackfriday.Run(bs))),
		}
		var tpl *template.Template
		if config.templatePath == "" {
			tpl, err = template.New("default").Parse(defaultTemplate)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			tpl, err = template.ParseFiles(config.templatePath)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		if err := tpl.Execute(w, context); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case ".htm", ".html":
		w.Header().Set("Content-Type", "text/html")
		w.Write(bs)
		return
	default:
		w.Write(bs)
		return
	}
}
