package model

type YoutubeStruct struct {
	url         string
	channel     string
	description string
	// 0 for no started , 1 for in progress , 2 for Done
	status int
}
