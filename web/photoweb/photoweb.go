package main
import (
	"io"
	"os"
	"log"
	"net/http"
	"io/ioutil"
	"html/template"
	"path"
	"fmt"
)

const (
	UPLOAD_DIR = "./uploads"
	TEMPLATE_DIR = "./views"
	ListDir = 0x0001
)

var  templates map[string] *template.Template= make(map[string] *template.Template)

func init(){
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	check(err)

	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html"{
			continue
		}
		templatePath = TEMPLATE_DIR+"/" + templateName;
		t := template.Must(template.ParseFiles(templatePath))
		rlname := realName(templateName)
		log.Println("loading template:", rlname)
		templates[rlname]=t
	}
}

func realName(str string) string{
	str = path.Base(str)
	if str == "" {
		return str
	}

	for i:= 0; i<len(str); i++{
		if '.' == str[i] {
			return str[:i]
		}
	}
	return str
}

func uploadHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		readerHtml(w, "upload", nil)
	}

	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR+"/"+filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer t.Close()
		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request){
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func isExists(path string) bool{
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	return os.IsExist(err)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir("./uploads")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("failed @ listHandler")
		return
	}

	locals := make(map[string] interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr{
		if fileInfo.Name() != ".svn" {
			images = append(images, fileInfo.Name())
		}
	}
	locals["images"]=images

	readerHtml(w, "list", locals)
}

func readerHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}){
	err := templates[tmpl].Execute(w, locals)
	check(err)
}

func check (err error) {
	if err != nil {
		panic(err)
	}
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r*http.Request){
		defer func(){
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				log.Println("WARN: panic in %v - %v", fn, e)
			}
		}()
		fn(w, r)
	}
}

func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int){
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r*http.Request){
		file := staticDir + r.URL.Path[len(prefix)-1:]	
		if (flags & ListDir) == 0 {
			if exists := isExists(file); !exists{
				http.NotFound(w, r)
				fmt.Println(file, "not found")
				return
			}
		}
		fmt.Println("handle static dir")
		http.ServeFile(w, r, file)
	})
}

func main(){
	mux := http.NewServeMux()
	staticDirHandler(mux, "/assets/", "./public", 0)
	mux.HandleFunc("/", safeHandler(listHandler))
	mux.HandleFunc("/view", safeHandler(viewHandler))
	mux.HandleFunc("/upload", safeHandler(uploadHandler))
	
	err := http.ListenAndServe(":8090", mux)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}

