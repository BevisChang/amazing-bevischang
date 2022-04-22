package rpc

import (
	"context"
	"strconv"
	"time"
	"unsafe"

	"github.com/AmazingTalker/bevis-chang/pkg/dao"
	"github.com/AmazingTalker/bevis-chang/pkg/pb"
	"github.com/AmazingTalker/bevis-chang/pkg/rpc/config"
	"github.com/AmazingTalker/go-rpc-kit/logkit"
	"github.com/AmazingTalker/go-rpc-kit/metrickit"
	"github.com/AmazingTalker/go-rpc-kit/validatorkit"
)

var (
	rpcMet = metrickit.NewWithPkgName(
		metrickit.EnableAutoFillInFuncName(true),
	)
)

type BevisChangServerOpt struct {
	Validator validatorkit.Validator
	RecordDao dao.RecordDAO
	MemberDao dao.MemberDAO
}

// BevisChangServer 1. Implement a struct as you like.
// Generate everything with an interface named "BevisChangRPC"
type BevisChangServer struct {
	logkit    logkit.AmazingLogger
	validator validatorkit.Validator
	recordDao dao.RecordDAO
	memberDao dao.MemberDAO
}

func NewBevisChangServer(opt BevisChangServerOpt) BevisChangServer {
	return BevisChangServer{
		validator: opt.Validator,
		recordDao: opt.RecordDao,
		memberDao: opt.MemberDao,
	}
}

// Health 2. Complete these methods.
func (serv BevisChangServer) Health(_ context.Context, _ *pb.HealthReq) (*pb.HealthRes, error) {
	return &pb.HealthRes{Ok: true}, nil
}

func (serv BevisChangServer) Config(ctx context.Context, _ *pb.ConfigReq) (*pb.ConfigRes, error) {
	cfg := config.Config()

	return &pb.ConfigRes{
		Enable: cfg.Enable,
		Num:    cfg.Num,
		Str:    cfg.Str,
	}, nil
}

func (serv BevisChangServer) CreateRecord(ctx context.Context, req *pb.CreateRecordReq) (*pb.CreateRecordRes, error) {
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

func (serv BevisChangServer) GetRecord(ctx context.Context, req *pb.GetRecordReq) (*pb.GetRecordRes, error) {
	defer rpcMet.RecordDuration([]string{"time"}, map[string]string{}).End()

	ctx = logkit.EnrichPayload(ctx, logkit.Payload{"id": req.ID})

	r, err := serv.recordDao.GetRecord(ctx, req.ID)

	resp := pb.GetRecordRes{Record: r.FormatPb()}
	rpcMet.SetGauge([]string{"resp_size"}, float64(unsafe.Sizeof(resp)), map[string]string{})

	return &resp, err
}

func (serv BevisChangServer) ListRecord(ctx context.Context, req *pb.ListRecordReq) (*pb.ListRecordRes, error) {
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

func (serv BevisChangServer) CreateMember(ctx context.Context, req *pb.CreateMemberReq) (*pb.CreateMemberRes, error) {
	m := &dao.Member{
		Name:     req.Name,
		Birthday: req.Birthday,
	}

	err := serv.memberDao.CreateMember(ctx, m)

	if err != nil {
		return nil, err
	}

	resp := pb.CreateMemberRes{Member: m.FormatPb()}

	return &resp, nil
}

func (serv BevisChangServer) UpdateMember(ctx context.Context, req *pb.UpdateMemberReq) (*pb.UpdateMemberRes, error) {
	id, _ := strconv.ParseInt(req.ID, 10, 64)
	m := &dao.Member{
		ID:       id,
		Name:     req.Name,
		Birthday: req.Birthday,
	}

	updatedMember, err := serv.memberDao.UpdateMember(ctx, m)

	if err != nil {
		return nil, err
	}

	resp := pb.UpdateMemberRes{Member: updatedMember.FormatPb()}

	return &resp, nil
}

func (serv BevisChangServer) ListMembers(ctx context.Context, req *pb.ListMembersReq) (*pb.ListMembersRes, error) {
	rpcMet.IncrCount([]string{"list_member_call_count"}, float64(1), map[string]string{})

	birthdayBefore, _ := time.Parse(time.RFC3339, req.BirthdayBefore)
	members, err := serv.memberDao.ListMembers(ctx, &birthdayBefore)

	if err != nil {
		return nil, err
	}

	result := make([]*pb.Member, len(members))

	for i, m := range members {
		m := m
		result[i] = m.FormatPb()
	}

	resp := pb.ListMembersRes{Member: result}

	return &resp, nil
}

func (serv BevisChangServer) DeleteMember(ctx context.Context, req *pb.DeleteMemberReq) (*pb.DeleteMemberRes, error) {
	id, _ := strconv.ParseInt(req.ID, 10, 64)

	err := serv.memberDao.DeleteMember(ctx, id)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
