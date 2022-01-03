package webserver

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/nanu-c/axolotl/app/config"
	"github.com/nanu-c/axolotl/app/store"
	"github.com/signal-golang/textsecure/profiles"
	log "github.com/sirupsen/logrus"
)

func attachmentsHandler(w http.ResponseWriter, r *http.Request) {
	Filename := r.URL.Query().Get("file")
	if Filename == "" {
		//Get not set, send a 400 bad request
		http.Error(w, "Get 'file' not specified in url.", 400)
		return
	}

	//Check if file exists and open
	filename := strings.Split(Filename, "/")
	path := config.AttachDir + "/" + filename[len(filename)-1]
	log.Debugln("[axolotl] open file: " + path)
	Openfile, err := os.Open(path)
	defer Openfile.Close() //Close after function return
	if err != nil {
		//File not found, send 404
		http.Error(w, "File not found.", 404)
		return
	}
	//File is found, create and send the correct headers

	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	Openfile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	w.Header().Set("Content-Disposition", "attachment; filename="+Filename)
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)

	//Send the file
	//We read 512 bytes from the file already, so we reset the offset back to 0
	Openfile.Seek(0, 0)
	io.Copy(w, Openfile) //'Copy' the file to the client
}
func avatarsGroupHandler(w http.ResponseWriter, r *http.Request) {

	Filename := r.URL.Query().Get("file")
	if Filename == "" {

		//Get not set, send a 400 bad request
		http.Error(w, "Get 'file' not specified in url.", 400)
		return
	}
	//handle group avatars
	log.Debugln("[axolotl] open group avatar file: " + Filename)
	if len(Filename) > 30 {
		group := store.GetGroupById(Filename)
		if group == nil {
			//File not found, send 404
			http.Error(w, "File not found.", 404)
			return
		}
		FileContentType := http.DetectContentType(group.Avatar)
		w.Header().Set("Content-Disposition", "attachment; filename="+Filename+".png")
		w.Header().Set("Content-Type", FileContentType)
		w.Write(group.Avatar)
	} else {

		http.Error(w, "File not found.", 404)
		return
	}
}
func avatarsProfileHandler(w http.ResponseWriter, r *http.Request) {

	Filename := r.URL.Query().Get("file")
	if Filename == "" {

		//Get not set, send a 400 bad request
		http.Error(w, "Get 'file' not specified in url.", 400)
		return
	}
	//handle group avatars
	log.Debugln("[axolotl] open profile avatar file: " + Filename)
	if len(Filename) > 30 {

		// log.Debugln("[axolotl] non group")
		// Filename = "+" + Filename[1:]
		contact := store.GetContactForUUID(Filename)
		_, err := profiles.GetProfileUUID(Filename)
		if err != nil {
			log.Errorln("[axolotl] error getting profile: " + err.Error())
			http.Error(w, "File not found.", 404)
			return
		}
		if contact == nil {
			http.Error(w, "File not found.", 404)
			return
		}
		if !contact.Avatar {
			//File not found, send 404
			http.Error(w, "File not found.", 404)
			return
		}
		log.Debugln("[axolotl] avatar len >0")
		avatar, err := ioutil.ReadFile(config.AttachDir + "/" + contact.UUID + ".png")
		if err != nil {
			//File not found, send 404
			http.Error(w, "File not found.", 404)
			return
		}
		FileContentType := http.DetectContentType(avatar)
		log.Debugln("[axolotl] file content type", FileContentType)
		w.Header().Set("Content-Disposition", "attachment; filename="+Filename+".png")
		w.Header().Set("Content-Type", FileContentType)
		w.Write(avatar)
	}
}
