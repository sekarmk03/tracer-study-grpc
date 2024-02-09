package main

import (
	"fmt"
	"tracer-study-grpc/common/config"

	// errUtils "tracer-study-grpc/common/errors"
	gormConn "tracer-study-grpc/common/gorm"
	commonJwt "tracer-study-grpc/common/jwt"
	"tracer-study-grpc/common/mysql"
	"tracer-study-grpc/server"

	authModules "tracer-study-grpc/modules/auth"
	kabkotaModules "tracer-study-grpc/modules/kabkota"
	mhsbiodataModules "tracer-study-grpc/modules/mhsbiodata"
	pktsModules "tracer-study-grpc/modules/pkts"
	prodiModules "tracer-study-grpc/modules/prodi"
	provinsiModules "tracer-study-grpc/modules/provinsi"
	respondenModules "tracer-study-grpc/modules/responden"
	userstudyModules "tracer-study-grpc/modules/userstudy"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func main() {
	cfg, cerr := config.NewConfig(".env")
	checkError(cerr)

	splash(cfg)

	dsn, derr := mysql.NewPool(&cfg.MySQL)
	checkError(derr)
	// errUtils.ConvertToRestError(derr)

	db, gerr := gormConn.NewMySQLGormDB(dsn)
	checkError(gerr)
	// errUtils.ConvertToRestError(gerr)

	jwtManager := commonJwt.NewJWT(cfg.JWT.JwtSecretKey, cfg.JWT.TokenDuration)

	grpcServer := server.NewGrpcServer(cfg.Port.GRPC, jwtManager)
	grpcConn := server.InitGRPCConn(fmt.Sprintf("127.0.0.1:%v", cfg.Port.GRPC), false, "")

	registerGrpcHandlers(grpcServer.Server, *cfg, db, jwtManager, grpcConn)

	_ = grpcServer.Run()
	_ = grpcServer.AwaitTermination()
}

func registerGrpcHandlers(server *grpc.Server, cfg config.Config, db *gorm.DB, jwtManager *commonJwt.JWT, grpcConn *grpc.ClientConn) {
	prodiModules.InitGrpc(server, cfg, db, grpcConn)
	provinsiModules.InitGrpc(server, cfg, db, grpcConn)
	kabkotaModules.InitGrpc(server, cfg, db, grpcConn)
	mhsbiodataModules.InitGrpc(server, cfg, db, grpcConn)
	respondenModules.InitGrpc(server, cfg, db, grpcConn)
	pktsModules.InitGrpc(server, cfg, db, grpcConn)
	authModules.InitGrpc(server, cfg, db, jwtManager, grpcConn)
	userstudyModules.InitGrpc(server, cfg, db, grpcConn)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func splash(cfg *config.Config) {
	version := "1.0.0"
	colorReset := "\033[0m"
	// colorBlue := "\033[34m"
	colorCyan := "\033[36m"

	fmt.Printf(`
    __                                                    __                    __
  _/  |_ _______ ______   ____  ______ _______     ______/  |__  __   __       |  \ __   __
  \   __\\_  __ \\___  \_/ ___\/  __  \\_  __ \   /  ___/\   __\|  | |  |   ___|  ||  | |  |
   |  |   |  | \/  / ___\  \___\   ___/ |  | \/   \___ \  |  |  |  |_|  | /  __   ||  |_|  | >>
   |__|   |__|    (____  /\__  >\___  > |__|     /____  > |__|  | ______|(  ____  )\__   __/
                       \/    \/     \/                \/        \/        \/    \/ _ /  /    v%s
                                                                                  / ___/
	`, version)

	// fmt.Println(colorBlue, fmt.Sprintf(`⇨ REST server started on port :%s`, cfg.Port.REST))
	fmt.Println(colorCyan, fmt.Sprintf(`⇨ GRPC server started on port :%s`, cfg.Port.GRPC))
	fmt.Println(colorReset, "")
}
