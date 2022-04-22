package dao

import (
	"context"
	"github.com/AmazingTalker/bevis-chang/pkg/pb"
	"time"

	"github.com/AmazingTalker/go-rpc-kit/daokit"
)

type MemberDAO interface {
	CreateMember(context.Context, *Member, ...daokit.Enrich) error
	UpdateMember(context.Context, *Member, ...daokit.Enrich) (*Member, error)
	ListMembers(context.Context, *time.Time, ...daokit.Enrich) ([]Member, error)
	DeleteMember(context.Context, int64, ...daokit.Enrich) error
}

type Member struct {
	ID        int64
	Name      string
	Birthday  *time.Time
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (m *Member) FormatPb() *pb.Member {
	return &pb.Member{
		ID:        m.ID,
		Name:      m.Name,
		Birthday:  m.Birthday,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
