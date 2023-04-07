package grpc_client

import (
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/config"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type DialFactory interface {
	GetEventsProcessingPlatformConn() (*grpc.ClientConn, error)
	Close(conn *grpc.ClientConn)
}

type DialFactoryImpl struct {
	conf                         *config.Config
	eventsProcessingPlatformConn *grpc.ClientConn
}

func (dfi DialFactoryImpl) Close(conn *grpc.ClientConn) {
	err := conn.Close()
	if err != nil {
		logrus.Printf("Error closing connection: %v", err)
	}
}

func (dfi DialFactoryImpl) GetEventsProcessingPlatformConn() (*grpc.ClientConn, error) {
	if dfi.eventsProcessingPlatformConn != nil {
		return dfi.eventsProcessingPlatformConn, nil
	}
	// TODO: alexb investigate for required dial options
	conn, err := grpc.Dial(dfi.conf.Service.EventsProcessingPlatformUrl, grpc.WithInsecure())
	dfi.eventsProcessingPlatformConn = conn
	return conn, err
}

func NewDialFactory(conf *config.Config) DialFactory {
	dfi := new(DialFactoryImpl)
	dfi.conf = conf
	return *dfi
}
