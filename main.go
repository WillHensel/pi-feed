package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("./bin/ffmpeg.exe -h")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	var out []byte
	_, err = stdout.Read(out)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(out[:]))
}
