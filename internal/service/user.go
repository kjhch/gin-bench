package service

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/kjhch/gin-bench/internal/repo"
	"strconv"
	"time"
)

type Page[T any] struct {
	PageNum    *uint   `json:"pageNum"`
	PageSize   *uint   `json:"pageSize"`
	TotalPage  *uint   `json:"totalPage"`
	TotalCount *uint64 `json:"totalCount"`
	List       []*T    `json:"list"`
}

type User struct {
	ID         *uint64    `json:"id"`
	Username   *string    `json:"username"`
	Gender     *string    `json:"gender"`
	Phone      *string    `json:"phone"`
	Email      *string    `json:"email"`
	Address    *string    `json:"address"`
	CreateTime *time.Time `json:"createTime,omitempty"`
	UpdateTime *time.Time `json:"updateTime,omitempty"`
	Creator    *string    `json:"creator,omitempty"`
	Modifier   *string    `json:"modifier,omitempty"`
}
type UserService struct {
	userRepo *repo.UserRepo
	rdb      *redis.Client
}

func NewUserService(userRepo *repo.UserRepo, rdb *redis.Client) *UserService {
	return &UserService{userRepo: userRepo, rdb: rdb}
}

func (svc *UserService) GetUser(id uint64) *User {
	u := svc.userRepo.GetUser(id)
	return &User{
		ID:       u.ID,
		Username: u.Username,
		Gender:   newVal(strconv.Itoa(*u.Gender)),
		Phone:    u.Phone,
		Email:    u.Email,
		Address:  u.Address,
	}
}

func (svc *UserService) ListUsers(pageNum, pageSize uint) *Page[User] {
	count, err := svc.rdb.Get(context.Background(), "pt:db:count").Uint64()
	if err != nil {
		panic(nil)

	}
	var users = make([]*User, 0)
	if count < 1 {
		return &Page[User]{
			PageNum:    nil,
			PageSize:   nil,
			TotalPage:  nil,
			TotalCount: &count,
			List:       users,
		}
	}
	totalPage := uint(count / uint64(pageSize))

	if count%uint64(pageSize) != 0 {
		totalPage++
	}
	if pageNum > totalPage {
		pageNum = totalPage
	}
	userList := svc.userRepo.ListUsers((pageNum-1)*pageSize, pageSize)
	for _, u := range userList {
		users = append(users, &User{
			ID:       u.ID,
			Username: u.Username,
			Gender:   newVal(strconv.Itoa(*u.Gender)),
			Phone:    u.Phone,
			Email:    u.Email,
			Address:  u.Address,
		})
	}
	return &Page[User]{
		PageNum:    &pageNum,
		PageSize:   &pageSize,
		TotalPage:  &totalPage,
		TotalCount: &count,
		List:       users,
	}
}

func newVal[T any](a T) *T {
	return &a
}
