package model

type Nature int64

const (
	Event Nature = iota
	Birth
	Death
	Holiday
)

func (s Nature) String() string {
	switch s {
	case Event:
		return "events"
	case Birth:
		return "births"
	case Death:
		return "deaths"
	case Holiday:
		return "holidays"
	default:
		return "unknown"
	}	
}

type Media struct {
	Url string
	Width int64
	Height int64
}

type Story struct {
	Id string
	Day int64
	Month int64
	Nature Nature
	Title string
	Description string
	Content string
	Year int64
	Media Media
}

type StoryList struct {
	Nature Nature
	Stories []Story
	Day int64
	Month int64
}