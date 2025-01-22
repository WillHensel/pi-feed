package main

import (
	"flag"
	"log"
	"net/http"
	"text/template"
)

func home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./web/html/base.html",
		"./web/html/pages/home.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func main() {
	// -c:v h264 -- set the video codec to mux to as h264
	// -flags +cgop -g 150 -hls_time 5 -- Set the "Group of Pictures" size to 150 frames and the segment time to 5 seconds, assuming video is 30FPS
	// cmd := exec.Command("./bin/ffmpeg.exe", "-i ./test-data/BADKITTY.mp4 -c:v h264 -flags +cgop -g 150 -hls_time 5 out.m3u8")
	// stdout, err := cmd.StdoutPipe()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err := cmd.Run(); err != nil {
	// 	log.Fatal(err)
	// }

	// var out []byte
	// _, err = stdout.Read(out)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(string(out[:]))

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	fileServer := http.FileServer(http.Dir("./web/videos"))
	mux.Handle("/videos/", http.StripPrefix("/videos", fileServer))

	err := http.ListenAndServe(*addr, mux)

	log.Fatal(err.Error())
}
