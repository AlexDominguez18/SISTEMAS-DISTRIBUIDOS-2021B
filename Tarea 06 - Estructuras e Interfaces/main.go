package main

import (
	"bufio"
	"fmt"
	"os"
	"./content"
)

const (
	ADD_IMG = 1
	ADD_AUDIO = 2
	ADD_VIDEO = 3
	SHOW_CONTENT = 4
	EXIT = 5
)

func showMenu(){
	fmt.Println("MENU")
	fmt.Println("1. Add image")
	fmt.Println("2. Add audio")
	fmt.Println("3. Add video")
	fmt.Println("4. Show content")
	fmt.Println("5. Exit")
	fmt.Print("\nOption = ")
}

func addImage() content.Image {
	var img content.Image
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Title = ")
	scanner.Scan()
	img.Title = scanner.Text()

	fmt.Print("Format = ")
	scanner.Scan()
	img.Format = scanner.Text()

	fmt.Print("Channels = ")
	scanner.Scan()
	img.Channels = scanner.Text()

	return img
}

func addAudio() content.Audio {
	var audio content.Audio
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Title = ")
	scanner.Scan()
	audio.Title = scanner.Text()

	fmt.Print("Format = ")
	scanner.Scan()
	audio.Format = scanner.Text()

	fmt.Print("Duration = ")
	fmt.Scan(&audio.Duration)

	return audio
}

func addVideo() content.Video{
	var video content.Video
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Title = ")
	scanner.Scan()
	video.Title = scanner.Text()

	fmt.Print("Format = ")
	scanner.Scan()
	video.Format = scanner.Text()

	fmt.Print("Frames per second = ")
	fmt.Scan(&video.Fps)

	return video
}

func main() {
	var option int
	var webContent content.WebContent
	
	for isRunning := true; isRunning; {
		fmt.Print("\033[H\033[2J")
		showMenu()
		fmt.Scan(&option)
		fmt.Print("\033[H\033[2J")
		switch option {
			case ADD_IMG:
				img := addImage()
				webContent.Content= append(webContent.Content, &img)
			case ADD_AUDIO:
				audio := addAudio()
				webContent.Content= append(webContent.Content, &audio)
			case ADD_VIDEO:
				video := addVideo()
				webContent.Content= append(webContent.Content, &video)
			case SHOW_CONTENT:
				webContent.Show()
			case EXIT:
				isRunning = false
			default:
				fmt.Println("Invalid option!")
		}
		if isRunning {
			fmt.Println("Presiona ENTER para continuar...")
			fmt.Scanln()
		}
	}
}