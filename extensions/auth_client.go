package extensions

import (
	"context"
	"os"
	"sync"

	"github.com/SaiNageswarS/go-api-boot/logger"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	pb "github.com/kotlang/localizationGo/generated"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var auth_client *AuthClient = &AuthClient{}

type AuthClient struct {
	cached_conn        *grpc.ClientConn
	conn_creation_lock sync.Mutex
}

func (c *AuthClient) getConnection() *grpc.ClientConn {
	c.conn_creation_lock.Lock()
	defer c.conn_creation_lock.Unlock()

	if c.cached_conn == nil || c.cached_conn.GetState().String() != "READY" {
		if val, ok := os.LookupEnv("AUTH_TARGET"); ok {
			conn, err := grpc.Dial(val, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
			if err != nil {
				logger.Error("Failed getting connection with auth service", zap.Error(err))
				return nil
			}
			c.cached_conn = conn
		} else {
			logger.Error("Failed to get AUTH_TARGET env variable")
		}

	}

	return c.cached_conn
}

func IsUserAdmin(grpcContext context.Context, userId string) chan bool {
	result := make(chan bool)

	go func() {
		conn := auth_client.getConnection()
		if conn == nil {
			result <- false
			return
		}

		client := pb.NewLoginVerifiedClient(conn)

		ctx := prepareCallContext(grpcContext)
		if ctx == nil {
			result <- false
			return
		}

		resp, err := client.IsUserAdmin(ctx, &pb.IdRequest{UserId: userId})
		if err != nil {
			logger.Error("Failed getting user profile", zap.Error(err))
			result <- false
			return
		}

		result <- resp.IsAdmin
	}()

	return result
}

func prepareCallContext(grpcContext context.Context) context.Context {
	jwtToken, err := grpc_auth.AuthFromMD(grpcContext, "bearer")
	if err != nil {
		logger.Error("Failed getting jwt token", zap.Error(err))
		return nil
	}

	return metadata.AppendToOutgoingContext(context.Background(), "Authorization", "bearer "+jwtToken)
}
