package dao

import (
	"context"
	"gorm.io/gorm"
	"time"

	"github.com/AmazingTalker/go-rpc-kit/daokit"
)

type MySqlMemberDAO struct {
	db *gorm.DB
}

func NewMySqlMemberDAO(db *gorm.DB) MySqlMemberDAO {
	return MySqlMemberDAO{db: db}
}

func (dao MySqlMemberDAO) CreateMember(ctx context.Context, member *Member, enrich ...daokit.Enrich) error {

	db, _ := daokit.UseTxOrDB(dao.db, enrich...)

	err := db.Create(member).Error

	if err != nil {
		return err
	}
	return nil
}

func (dao MySqlMemberDAO) UpdateMember(ctx context.Context, member *Member, enrich ...daokit.Enrich) (*Member, error) {
	db, _ := daokit.UseTxOrDB(dao.db, enrich...)

	targetMember := Member{}
	db.First(&targetMember, member.ID)
	db.Model(&targetMember).Updates(member)

	return &targetMember, nil
}

func (dao MySqlMemberDAO) ListMembers(ctx context.Context, birthdayBefore *time.Time, enrich ...daokit.Enrich) ([]Member, error) {
	db, _ := daokit.UseTxOrDB(dao.db, enrich...)

	var members []Member
	db.Where("birthday <= ?", &birthdayBefore).Find(&members)

	return members, nil
}

func (dao MySqlMemberDAO) DeleteMember(ctx context.Context, id int64, enrich ...daokit.Enrich) error {
	db, _ := daokit.UseTxOrDB(dao.db, enrich...)

	db.Delete(&Member{}, id)

	return nil
}
