package dao

import (
	"context"
	"github.com/AmazingTalker/go-cache"
	"github.com/AmazingTalker/go-rpc-kit/daokit"
	"gorm.io/gorm"
)

type MemberImpl struct {
	mysql MySqlMemberDAO
}

func NewMemberDAO(db *gorm.DB, cacheSrv cache.Service) MemberDAO {
	return &MemberImpl{mysql: NewMySqlMemberDAO(db)}
}

func (im *MemberImpl) CreateMember(ctx context.Context, member *Member, enrich ...daokit.Enrich) error {
	return im.mysql.CreateMember(ctx, member, enrich...)
}

func (im *MemberImpl) UpdateMember(ctx context.Context, member *Member, enrich ...daokit.Enrich) (Member, error) {
	return im.mysql.UpdateMember(ctx, member, enrich...)
}

func (im *MemberImpl) ListMembers(ctx context.Context, birthdayBefore int64, enrich ...daokit.Enrich) ([]Member, error) {
	return im.mysql.ListMembers(ctx, birthdayBefore, enrich)
}

func (im *MemberImpl) DeleteMember(ctx context.Context, id int64, enrich ...daokit.Enrich) error {
	return im.mysql.DeleteMember(ctx, id, enrich)
}
