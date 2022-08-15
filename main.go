package main;

import (
	"fmt"
	"io"
	"os"
	"time"
	"strings"
	"github.com/kkdai/youtube/v2"
)


func author() {
	str := "Created by NodirCoder"
	for _, w := range strings.Split(str, "") {
		fmt.Printf("%s", w)
		time.Sleep(200000000)
	}
	fmt.Println("\n\n")
}


func success() {
	fmt.Println("\n--------------------------------------------------")
	fmt.Println("\nVideo has been uploaded!");
}


func start() {
	author()
	var link, f_name, q string
	arr := make([]string, 1)

	fmt.Print("Enter a link: ")
	fmt.Scanln(&link)
	fmt.Print("Enter a file name: ")
	fmt.Scanln(&f_name)
	fmt.Print("Enter quality of video: ")
	fmt.Scanln(&q)

	arr = append(arr, q)
	b, err := downloadByUrl(link, f_name, arr...)
	
	if b {
		success()
	} else {
		fmt.Println("Error: " + err)
	}
}

func getFormats(url string) {
	// videoID := "x1bUOfRXt3Q"
	client := youtube.Client{}

	video, err := client.GetVideo(url)
	if err != nil {
		panic(err)
	}

	formats := video.Formats
	// count := len(formats)
	// fmt.Println("Count:", count)
	fmt.Println(formats)

	// for _, format := range formats {	
	// 	fmt.Println(format)
	// }
}


func downloadByUrl(url, f_name string, q ...string) (bool, string) {
	// videoID := "x1bUOfRXt3Q"
	q_f := "144p"
	if len(q) > 0 {
		q_f = q[0]
	}

	client := youtube.Client{}

	video, err := client.GetVideo(url)
	if err != nil {
		return false, "Error with internet!"
	}

	formats := video.Formats.FindByQuality(q_f)
	// fmt.Println(formats)
	stream, _, err := client.GetStream(video, formats)
	if err != nil {
		return false, "No such format!"
	}

	file, err := os.Create(f_name + ".mp4")
	if err != nil {
		return false, "Cannot create file!"
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		return false, "Unknown error!"
	}
	return true, "Success"
}


func main() {
	start()
	
	// getFormats()
	// recognizeFile("qrcode.png")
}
