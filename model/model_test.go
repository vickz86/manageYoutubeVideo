package model

import "testing"

func TestDescribe(t *testing.T) {
	// create test struct
	youtubeVideos := []YoutubeStruct{
		{Index: 1, Url: "https://example.com/1", Channel: "Channel A", Description: "This is the first video.", Status: 1},
		{Index: 2, Url: "https://example.com/2", Channel: "Channel B", Description: "This is the second video.", Status: 2},
		{Index: 3, Url: "https://example.com/3", Channel: "Channel C", Description: "This is the third video.", Status: 0},
	}

	for _, video := range youtubeVideos {
		video.Describe()

	}
}
