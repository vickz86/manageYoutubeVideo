package model

type YoutubeStruct struct {
	Index       int
	Url         string
	Channel     string
	Description string
	// 0 for no started , 1 for in progress , 2 for Done
	Status int
}
