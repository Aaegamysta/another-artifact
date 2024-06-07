package cassandra

import (
	"crypto/tls"
	"log"
	"sync"

	"github.com/aaegamysta/another-artifact/api/conf"
	"github.com/aaegamysta/another-artifact/api/pkg/model"
	"github.com/stargate/stargate-grpc-go-client/stargate/pkg/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type StoriesRepository interface {
	PingPeriodicallyToPreventHibernation() error
	CreateTablesIfNotExists() error
	ListStories(int64, int64, model.Nature) ([]model.Story, error)
	InsertStory(model.Story) (model.Story, error)
}

type Impl struct {
	pool sync.Pool
}

var _ StoriesRepository = &Impl{}

func New(config conf.Cassandra) *Impl {
	impl := &Impl{
		pool: sync.Pool{
			New: func() any {
				tlsConfig := &tls.Config{
					InsecureSkipVerify: false,
				}
				log.Println(config.Url)
				log.Println(config.Token)
				conn, _ := grpc.NewClient(config.Url, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
					grpc.WithPerRPCCredentials(
						auth.NewStaticTokenProvider(config.Token),
					),
				)
				return conn
			},
		},
	}
	// impl.PingPeriodicallyToPreventHibernation()
	return impl
}

// CreateTablesIfNotExists implements StoriesRepository.
func (i *Impl) CreateTablesIfNotExists() error {
	panic("unimplemented")
}

// InsertStory implements StoriesRepository.
func (i *Impl) InsertStory(model.Story) (model.Story, error) {
	panic("unimplemented")
}

// ListStories implements StoriesRepository.
func (i *Impl) ListStories(int64, int64, model.Nature) ([]model.Story, error) {
	panic("unimplemented")
}

// PingPeriodicallyToPreventHibernation implements StoriesRepository.
func (i *Impl) PingPeriodicallyToPreventHibernation() error {
	panic("unimplemented")
}
