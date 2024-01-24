package main

import (
	"fmt"
	"tracer-study-grpc/common/config"
	gormConn "tracer-study-grpc/common/gorm"
	"tracer-study-grpc/common/mysql"
	"tracer-study-grpc/server"

	kabkotaModules "tracer-study-grpc/modules/kabkota"
	prodiModules "tracer-study-grpc/modules/prodi"
	provinsiModules "tracer-study-grpc/modules/provinsi"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func main() {
	cfg, cerr := config.NewConfig(".env")
	checkError(cerr)

	dsn, derr := mysql.NewPool(&cfg.MySQL)
	checkError(derr)

	db, gerr := gormConn.NewMySQLGormDB(dsn)
	checkError(gerr)

	grpcServer := server.NewGrpcServer(cfg.Port.GRPC)
	grpcConn := server.InitGRPCConn(fmt.Sprintf("127.0.0.1:%v", cfg.Port.GRPC), false, "")

	registerGrpcHandlers(grpcServer.Server, *cfg, db, grpcConn)

	_ = grpcServer.Run()
	_ = grpcServer.AwaitTermination()
}

func registerGrpcHandlers(server *grpc.Server, cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) {
	prodiModules.InitGrpc(server, cfg, db, grpcConn)
	provinsiModules.InitGrpc(server, cfg, db, grpcConn)
	kabkotaModules.InitGrpc(server, cfg, db, grpcConn)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
