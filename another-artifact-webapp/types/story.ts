enum Nature {
    Events,
    Births,
    Deaths,
    Holidays,
}

interface Media {
	Url : string,
	Width : number,
	Height : number ,
}

interface Story {
	Id : string
	Day : number
	Month : number
	Nature : Nature
	Title : string
	Description : string
	Content : string
	Year : number
	Media : Media
}

interface StoryList {
	Nature : Nature
	Stories : Story[]
	Day : number
	Month:  number
}