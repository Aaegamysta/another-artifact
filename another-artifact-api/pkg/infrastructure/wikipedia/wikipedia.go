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

// FetchBirths implements Client.
func (i *Impl) FetchBirths(int64, int64) (model.StoryList, error) {
	panic("unimplemented")
}

// FetchDeaths implements Client.
func (i *Impl) FetchDeaths(int64, int64) (model.StoryList, error) {
	panic("unimplemented")
}

// FetchEvents implements Client.
func (i *Impl) FetchEvents(day int64, month int64) (model.StoryList, error) {
	panic("unimplemented")
}

// FetchHolidays implements Client.
func (i *Impl) FetchHolidays(int64, int64) (model.StoryList, error) {
	panic("unimplemented")
}

// InsertEvents implements Client.
func (i *Impl) InsertEvents(model.StoryList) error {
	panic("unimplemented")
}

