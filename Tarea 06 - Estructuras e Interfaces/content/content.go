package content

import "fmt"

//Interface

type Multimedia interface {
	Show()
}

//Structs

type Image struct {
	Title string
	Format string
	Channels string
}

type Audio struct {
	Title string
	Format string
	Duration int64
}

type Video struct {
	Title string
	Format string
	Fps int64
}

type WebContent struct {
	Content []Multimedia
}

//Functions

func (img *Image) Show() {
	fmt.Println(img.Title + "." + img.Format + " | " + img.Channels)
}

func (a *Audio) Show() {
	fmt.Printf("%s.%s Duration: %d seconds\n", a.Title, a.Format, a.Duration)
}

func (v *Video) Show() {
	fmt.Printf("%s.%s %d Fps\n", v.Title, v.Format, v.Fps)
}

func (c *WebContent) Show() {
	for i, v := range c.Content {
		fmt.Printf("%d. ", i + 1)
		v.Show()
		for i := 0; i < 25; i++ {
			fmt.Print("-")
		}
		fmt.Println()
	}
}