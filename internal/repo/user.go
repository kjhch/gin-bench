package repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kjhch/gin-bench/internal/config"
	"go.uber.org/zap"
	"time"
)

type User struct {
	ID         *uint64
	Username   *string
	Gender     *int
	Phone      *string
	Email      *string
	Address    *string
	CreateTime *time.Time `db:"create_time"`
	UpdateTime *time.Time `db:"update_time"`
	Creator    *string
	Modifier   *string
}

//var columns = "id,username,gender,phone,email,address,create_time,update_time,creator,modifier"

type UserRepo struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

func NewUserRepo(db *sqlx.DB, factory *config.LoggerFactory) *UserRepo {
	return &UserRepo{
		db:     db,
		logger: factory.GetDefaultLogger().Sugar(),
	}
}

func (ur *UserRepo) GetUser(id uint64) *User {
	result := new(User)
	//ur.logger.Infof("id: %v", id)
	err := ur.db.Get(result, "SELECT id,username,gender,phone,email,address,create_time,update_time,creator,modifier FROM pt_user WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	//ur.logger.Infof("result: %v", result)
	return result
}

func (ur *UserRepo) ListUsers(offset, limit uint) []*User {
	var result = make([]*User, 0)
	sql := fmt.Sprintf("SELECT id,username,gender,phone,email,address,create_time,update_time,creator,modifier FROM pt_user LIMIT %d,%d", offset, limit)
	err := ur.db.Select(&result, sql)
	if err != nil {
		panic(err)
	}
	return result
}
