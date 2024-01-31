package handler

import (
	"context"
	"net/http"
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/pkts/entity"
	"tracer-study-grpc/modules/pkts/service"
	"tracer-study-grpc/pb"
)

type PKTSHandler struct {
	pb.UnimplementedPKTSServiceServer
	config  config.Config
	PKTSSvc service.PKTSServiceUseCase
}

func NewPKTSHandler(config config.Config, pktsService service.PKTSServiceUseCase) *PKTSHandler {
	return &PKTSHandler{
		config:  config,
		PKTSSvc: pktsService,
	}
}

func (ph *PKTSHandler) GetAllPKTS(ctx context.Context, req *pb.EmptyPKTSRequest) (*pb.GetAllPKTSResponse, error) {
	pkts, err := ph.PKTSSvc.FindAll(ctx, req)
	if err != nil {
		return nil, err
	}

	var pktsArr []*pb.PKTS
	for _, p := range pkts {
		pktsProto := entity.ConvertEntityToProto(p)
		pktsArr = append(pktsArr, pktsProto)
	}

	return &pb.GetAllPKTSResponse{
		Code:    uint32(http.StatusOK),
		Message: "get all pkts success",
		Data:    pktsArr,
	}, nil
}

func (ph *PKTSHandler) GetPKTSByNim(ctx context.Context, req *pb.GetPKTSByNimRequest) (*pb.GetPKTSResponse, error) {
	pkt, err := ph.PKTSSvc.FindByNim(ctx, req.Nim)
	if err != nil {
		return nil, err
	}

	pktProto := entity.ConvertEntityToProto(pkt)

	return &pb.GetPKTSResponse{
		Code:    uint32(http.StatusOK),
		Message: "get pkts by nim success",
		Data:    pktProto,
	}, nil
}

func (ph *PKTSHandler) CreatePKTS(ctx context.Context, req *pb.PKTS) (*pb.GetPKTSResponse, error) {
	// get data from api siak 1 -> kodeprodi, thn_sidang

	pkts, err := ph.PKTSSvc.Create(ctx, req.GetNim(), req.GetKodeprodi(), req.GetThnSidang(), req.GetF8(), req.GetF5_04(), req.GetF5_02(), req.GetF5_06(), req.GetF5_05(), req.GetF5A1(), req.GetF5A2(), req.GetF11_01(), req.GetF11_02(), req.GetF5B(), req.GetF5C(), req.GetF5D(), req.GetF18A(), req.GetF18B(), req.GetF18C(), req.GetF18D(), req.GetF12_01(), req.GetF12_02(), req.GetF14(), req.GetF15(), req.GetF1761(), req.GetF1762(), req.GetF1763(), req.GetF1764(), req.GetF1765(), req.GetF1766(), req.GetF1767(), req.GetF1768(), req.GetF1769(), req.GetF1770(), req.GetF1771(), req.GetF1772(), req.GetF1773(), req.GetF1774(), req.GetF21(), req.GetF22(), req.GetF23(), req.GetF24(), req.GetF25(), req.GetF26(), req.GetF27(), req.GetF301(), req.GetF302(), req.GetF303(), req.GetF4_01(), req.GetF4_02(), req.GetF4_03(), req.GetF4_04(), req.GetF4_05(), req.GetF4_06(), req.GetF4_07(), req.GetF4_08(), req.GetF4_09(), req.GetF4_10(), req.GetF4_11(), req.GetF4_12(), req.GetF4_13(), req.GetF4_14(), req.GetF4_15(), req.GetF4_16(), req.GetF6(), req.GetF7(), req.GetF7A(), req.GetF10_01(), req.GetF10_02(), req.GetF16_01(), req.GetF16_02(), req.GetF16_03(), req.GetF16_04(), req.GetF16_05(), req.GetF16_06(), req.GetF16_07(), req.GetF16_08(), req.GetF16_09(), req.GetF16_10(), req.GetF16_11(), req.GetF16_12(), req.GetF16_13(), req.GetF16_14(), req.GetNamaAtasan(), req.GetHpAtasan(), req.GetEmailAtasan(), req.GetTinggalSelamaKuliah())

	if err != nil {
		return nil, err
	}

	pktProto := entity.ConvertEntityToProto(pkts)
	
	return &pb.GetPKTSResponse{
		Code:    uint32(http.StatusOK),
		Message: "create pkts success",
		Data:    pktProto,
	}, nil
}