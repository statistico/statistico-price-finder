package mock

import (
	"context"
	"github.com/statistico/statistico-bet-finder/internal/app/grpc/proto"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type OddsCompilerServiceClient struct {
	mock.Mock
}

func (o OddsCompilerServiceClient) GetOverUnderGoalsForFixture (ctx context.Context, req *proto.OverUnderRequest, opts ...grpc.CallOption) (*proto.OverUnderGoalsResponse, error) {
	args := o.Called(ctx, req, opts)
	return args.Get(0).(*proto.OverUnderGoalsResponse), args.Error(1)
}

type FixtureServiceClient struct {
	mock.Mock
}

func (f FixtureServiceClient) FixtureByID(ctx context.Context, in *proto.FixtureRequest, opts ...grpc.CallOption) (*proto.Fixture, error) {
	args := f.Called(ctx, in, opts)
	return args.Get(0).(*proto.Fixture), args.Error(1)
}
