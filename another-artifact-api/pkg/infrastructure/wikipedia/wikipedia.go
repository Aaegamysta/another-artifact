package wikipedia

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/aaegamysta/another-artifact/api/conf"
	"github.com/aaegamysta/another-artifact/api/pkg/model"
	"github.com/buger/jsonparser"
)

type Client interface {
	InsertEvents(model.StoryList) error
	FetchEvents(int64, int64) (model.StoryList, error)
	FetchBirths(int64, int64) (model.StoryList, error)
	FetchDeaths(int64, int64) (model.StoryList, error)
	FetchHolidays(int64, int64) (model.StoryList, error)
}

type Impl struct {
	httpClient *http.Client
	baseUrl string
}

var _ Client = &Impl{}

func New(config conf.Wikipedia) *Impl {
	return &Impl{
		httpClient: &http.Client{},
		baseUrl: config.BaseUrl,
	}
}

func (i *Impl) InsertEvents(model.StoryList) error {
	panic("unimplemented")
}

func (i *Impl) FetchEvents(day int64, month int64) (model.StoryList, error) {
	storyList := model.StoryList{
		Day: day,
		Month: month,
		Nature: model.Birth,
		Stories: make([]model.Story, 0),
	}
	err := i.doFetch(&storyList)	
	if err != nil {
		return model.StoryList{}, err
	}
	return storyList, nil
}

func (i *Impl) FetchBirths(month int64, day int64) (model.StoryList, error) {
	storyList := model.StoryList{
		Day: day,
		Month: month,
		Nature: model.Birth,
		Stories: make([]model.Story, 0),
	}
	err := i.doFetch(&storyList)	
	if err != nil {
		return model.StoryList{}, err
	}
	return storyList, nil
}

func (i *Impl) FetchDeaths(month int64, day int64) (model.StoryList, error) {
	storyList := model.StoryList{
		Day: day,
		Month: month,
		Nature: model.Birth,
		Stories: make([]model.Story, 0),
	}
	err := i.doFetch(&storyList)	
	if err != nil {
		return model.StoryList{}, err
	}
	return storyList, nil
}


func (i *Impl) FetchHolidays(month int64, day int64) (model.StoryList, error) {
	storyList := model.StoryList{
		Day: day,
		Month: month,
		Nature: model.Birth,
		Stories: make([]model.Story, 0),
	}
	err := i.doFetch(&storyList)	
	if err != nil {
		return model.StoryList{}, err
	}
	return storyList, nil
}

func (i *Impl) doFetch(storyList *model.StoryList) (error) {
	url := fmt.Sprintf("%s/%s/%v/%v", i.baseUrl, storyList.Nature ,storyList.Month, storyList.Day)
	log.Println(url)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)	
	if err != nil {
		log.Panicf("failed to create request: %v", err)
	}

	res, err := i.httpClient.Do(req)
	if err != nil {
		log.Panicf("failed to do request: %v", err)
	}

	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Panicf("failed to read response body: %v", err)
	}

	log.Println(res.Status)
	err = i.parse(storyList, bytes)
	if err != nil {
		log.Panicf("failed to read response body: %v", err)
	}	
	return nil
}

func (i *Impl) parse(storyList *model.StoryList, bytes []byte) error {
	b, _, _, _ := jsonparser.Get(bytes, storyList.Nature.String())
	_, err := jsonparser.ArrayEach(b, func(b []byte, dataType jsonparser.ValueType, offset int, parsingErr error) {
		text, parsingErr := jsonparser.GetString(b, "text")
		if parsingErr != nil {
			log.Printf("failed to retrieve text during parsing %v", parsingErr)	
			return 
		}
		description, parsingErr := jsonparser.GetString(b, "pages", "[0]", "extract")
		if parsingErr != nil {
			log.Printf("failed to retrieve description during parsing %v", parsingErr)	
		}
		// Fetching thumbnail
		thumbnailUrl, parsingErr := jsonparser.GetString(b, "pages", "[0]", "thumbnail", "source")
		if parsingErr != nil {
			log.Printf("failed to retrieve thumbnail url during parsing %v", parsingErr)	
		}
		thumbnailWidth, parsingErr := jsonparser.GetInt(b, "pages", "[0]", "thumbnail", "width")
		if parsingErr != nil {
			log.Printf("failed to retrieve thumbnail width during parsing %v", parsingErr)	
		}
		thumbnailHeight, parsingErr := jsonparser.GetInt(b, "pages", "[0]", "thumbnail", "height")
		if parsingErr != nil {
			log.Printf("failed to retrieve thumbnail height during parsing %v", parsingErr)	
		}
		year, parsingErr := jsonparser.GetInt(b, "year")
		if parsingErr != nil {
			log.Printf("failed to retrieve height during parsing %v", parsingErr)	
		}
		storyList.Stories = append(storyList.Stories, model.Story{
			Title: text,
			Description: description,
			Day: storyList.Day,
			Month: storyList.Month,
			Nature: storyList.Nature,
			Year: year,
			Media: model.Media{
				Url: thumbnailUrl,
				Width: thumbnailWidth,
				Height: thumbnailHeight,
			},
		})
	})
	if err != nil {
		return err
	}
	return nil
}