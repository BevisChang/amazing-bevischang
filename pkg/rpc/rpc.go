package rpc

import (
	"context"
	"strconv"
	"unsafe"

	"github.com/AmazingTalker/go-amazing/pkg/dao"
	"github.com/AmazingTalker/go-amazing/pkg/pb"
	"github.com/AmazingTalker/go-amazing/pkg/rpc/config"
	"github.com/AmazingTalker/go-rpc-kit/logkit"
	"github.com/AmazingTalker/go-rpc-kit/metrickit"
	"github.com/AmazingTalker/go-rpc-kit/validatorkit"
)

var (
	rpcMet = metrickit.NewWithPkgName(
		metrickit.EnableAutoFillInFuncName(true),
	)
)

type GoAmazingServerOpt struct {
	Validator validatorkit.Validator
	RecordDao dao.RecordDAO
}

// GoAmazingServer 1. Implement a struct as you like.
// Generate everything with an interface named "GoAmazingRPC"
type GoAmazingServer struct {
	logkit    logkit.AmazingLogger
	validator validatorkit.Validator
	recordDao dao.RecordDAO
}

func NewGoAmazingServer(opt GoAmazingServerOpt) GoAmazingServer {
	return GoAmazingServer{
		validator: opt.Validator,
		recordDao: opt.RecordDao,
	}
}

// Health 2. Complete these methods.
func (serv GoAmazingServer) Health(_ context.Context, _ *pb.HealthReq) (*pb.HealthRes, error) {
	return &pb.HealthRes{Ok: true}, nil
}

func (serv GoAmazingServer) Config(ctx context.Context, _ *pb.ConfigReq) (*pb.ConfigRes, error) {
	cfg := config.Config()

	return &pb.ConfigRes{
		Enable: cfg.Enable,
		Num:    cfg.Num,
		Str:    cfg.Str,
	}, nil
}

func (serv GoAmazingServer) CreateRecord(ctx context.Context, req *pb.CreateRecordReq) (*pb.CreateRecordRes, error) {
	defer rpcMet.RecordDuration([]string{"time"}, map[string]string{}).End()

	r := &dao.Record{
		TheNum: req.TheNum,
		TheStr: req.TheStr,
	}

	err := serv.recordDao.CreateRecord(ctx, r)

	if err != nil {
		return nil, err
	}

	resp := pb.CreateRecordRes{Record: r.FormatPb()}
	rpcMet.SetGauge([]string{"resp_size"}, float64(unsafe.Sizeof(resp)), map[string]string{})

	return &resp, nil
}

func (serv GoAmazingServer) GetRecord(ctx context.Context, req *pb.GetRecordReq) (*pb.GetRecordRes, error) {
	defer rpcMet.RecordDuration([]string{"time"}, map[string]string{}).End()

	ctx = logkit.EnrichPayload(ctx, logkit.Payload{"id": req.ID})

	r, err := serv.recordDao.GetRecord(ctx, req.ID)

	resp := pb.GetRecordRes{Record: r.FormatPb()}
	rpcMet.SetGauge([]string{"resp_size"}, float64(unsafe.Sizeof(resp)), map[string]string{})

	return &resp, err
}

func (serv GoAmazingServer) ListRecord(ctx context.Context, req *pb.ListRecordReq) (*pb.ListRecordRes, error) {
	defer rpcMet.RecordDuration([]string{"time"}, map[string]string{}).End()

	if err := serv.validator.Valid(ctx, req); err != nil {
		return nil, err
	}

	size, _ := strconv.ParseInt(req.PageSize, 10, 32)
	page, _ := strconv.ParseInt(req.Page, 10, 32)

	// Just demo
	records, err := serv.recordDao.ListRecords(ctx, dao.ListRecordsOpt{
		Size: int(size),
		Page: int(page),
	})

	if err != nil {
		return nil, err
	}

	result := make([]*pb.Record, len(records))

	for i, r := range records {
		r := r
		result[i] = r.FormatPb()
	}

	resp := pb.ListRecordRes{Records: result}
	rpcMet.SetGauge([]string{"resp_size"}, float64(unsafe.Sizeof(resp)), map[string]string{})

	return &resp, nil
}
