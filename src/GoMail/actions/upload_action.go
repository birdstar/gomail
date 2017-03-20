package actions

import (
	"os"
	"net/http"
	"io"
	"html/template"
	//"os/exec"
	"fmt"
	"github.com/golang/glog"
	"gomail/utils"
	"encoding/json"
)
//var out, err = exec.Command("pwd").Output()

//Compile templates on start
var templates = template.Must(template.ParseFiles("web/upload.html"))

//Read config from toml config file
var upload_path = utils.ReadConfig("development.image_upload_folder")


type Upload_Status struct {
	Ret      string
	File_path   string
}

//Display the named template
func display(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl+".html", data)
}

//This is where the action happens.
func Upload(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//GET displays the upload form.
	case "GET":
		glog.Info("In GET upload")
		display(w, "upload", nil)
		glog.Flush()


	//POST takes the uploaded file(s) and saves it to disk.
	case "POST":
		glog.Info("In POST upload")
		//parse the multipart form in the request
		err := r.ParseMultipartForm(100000)
		if err != nil {
			glog.Errorf(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//get a ref to the parsed multipart form
		m := r.MultipartForm
		glog.Info(r.MultipartForm.File)

		//get the *fileheaders
		files := m.File["file"]
		for i, _ := range files {
			//for each fileheader, get a handle to the actual file
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				glog.Errorf(err.Error())
				fmt.Errorf(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//create destination file making sure the path is writeable.
			//dst, err := os.Create("/tmp/" + files[i].Filename)
			dst, err := os.Create(upload_path + files[i].Filename)
			defer dst.Close()
			if err != nil {
				glog.Errorf(err.Error())
				fmt.Errorf(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//copy the uploaded file to the destination file
			if _, err := io.Copy(dst, file); err != nil {
				glog.Errorf(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		}
		glog.Flush()
		w.Header().Set("Content-Type", "application/json")
		//display success message.
		//display(w, "upload", "Upload successful.")
		status := Upload_Status{
			Ret: "ok",
			File_path: files[0].Filename,
		}
		json.NewEncoder(w).Encode(status)
		//ret, _ := json.Marshal(status)
		//w.Write(ret)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
