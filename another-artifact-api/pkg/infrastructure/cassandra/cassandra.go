package cassandra

import (
	"crypto/tls"
	"log"
	"sync"
	"time"

	"github.com/aaegamysta/another-artifact/api/conf"
	"github.com/aaegamysta/another-artifact/api/pkg/model"
	"github.com/stargate/stargate-grpc-go-client/stargate/pkg/auth"
	"github.com/stargate/stargate-grpc-go-client/stargate/pkg/client"
	"github.com/stargate/stargate-grpc-go-client/stargate/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type StoriesRepository interface {
	Pulsate() 
	CreateTablesIfNotExists() error
	ListStories(int64, int64, model.Nature) ([]model.Story, error)
	InsertStory(model.Story) (model.Story, error)
}

type Impl struct {
	pool sync.Pool
	enablePulse bool
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
		enablePulse: config.EnablePulse,
	}
	// impl.Pulsate()
	return impl
}

func (i *Impl) CreateTablesIfNotExists() error {
	conn := i.pool.Get().(*grpc.ClientConn)
	defer func() {
		i.pool.Put(conn)
	}()
	client, err := client.NewStargateClientWithConn(conn)
	if err != nil {
		log.Panicf("failed to create a client to create tables: %v", err)
	}
	_ , err = client.ExecuteQuery(&proto.Query{
		Cql: `CREATE TABLE main.stories (id text, day int, month int, nature text, title text, content text, year int, media_url text, media_width int, media_height int, PRIMARY KEY((day, month, nature), id));`,
	})
	if err != nil {
		log.Panicf("failed to create tables: %v", err)
	}		
	return nil
}

func (i *Impl) InsertStory(model.Story) (model.Story, error) {
	panic("unimplemented")
}

func (i *Impl) ListStories(int64, int64, model.Nature) ([]model.Story, error) {
	panic("unimplemented")
}

func (i *Impl) Pulsate() {
	panic("unimplemented")
}
