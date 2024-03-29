package biz

import "context"

// biz 层就是业务逻辑层
// biz= 业务层的实体(User) + 业务逻辑层的接口(UserRepo)+ 业务逻辑层的实例(UserCase) + 业务逻辑层的构造函数(NewUserCase) + 业务逻辑层的业务逻辑方法(CreateUser, GetUser...)

// User 是业务逻辑层的实体
type User struct {
	ID   int
	Age  int
	Name string
}

// UserRepo 是UserCase（服务层）需要的数据层接口
// 依赖倒置：高层模块不应该依赖于底层模块，两者都应该依赖于抽象
type UserRepo interface {
	Create(ctx context.Context, u *User) (*User, error)
	FindByID(ctx context.Context, id int) (*User, error)
}

// UserCase 是业务逻辑层的实例
type UserCase struct {
	repo UserRepo
}

// NewUserCase 是UserCase的构造函数
func NewUserCase(repo UserRepo) *UserCase {
	return &UserCase{repo: repo}
}

// CreateUser 是UserCase的创建用户方法
func (uc *UserCase) CreateUser(ctx context.Context, u *User) (*User, error) {
	println("业务逻辑 ...")
	return uc.repo.Create(ctx, u)
}

// GetUser 是UserCase的获取用户方法
func (uc *UserCase) GetUser(ctx context.Context, id int) (*User, error) {
	//组装业务逻辑
	res, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	for i := 0; i < 12; i++ {
		res.Age = res.Age + 1
	}
	//返回结果
	return res, nil
}
