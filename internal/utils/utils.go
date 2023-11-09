package utils

import (
	"fmt"
	"os"
	"sort"
)

type Files struct {
	Text  string
	Audio string
}

func MustLoadFiles() (map[int]*Files, error) {
	var filesMap = make(map[int]*Files)
	for i := 0; i < 20; i++ {
		filesMap[i] = &Files{}
	}
	audio, err := os.Open("files/audio")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	filesAudio, err := audio.Readdir(0)
	sort.Slice(filesAudio, func(i,j int) bool{
		return filesAudio[i].Name()<(filesAudio[j].Name())
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer audio.Close()
	for i, v := range filesAudio {
		filesMap[i].Audio = "files/audio/" + v.Name()
	}

	text, err := os.Open("files/text")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer text.Close()
	filesText, err := text.Readdir(0)
	sort.Slice(filesText, func(i,j int) bool{
		return filesText[i].Name()<(filesText[j].Name())
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for i, v := range filesText {
		filesMap[i].Text = "files/text/" + v.Name()
	}

	return filesMap, nil
}
