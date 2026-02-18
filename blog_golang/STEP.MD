# 系统搭建过程
github 的 .gitignore
*.db


# 启动main路由程序
pwd
blog/
vim setup.go

go mod init "blog"
go mod tidy
go run setup.go

# 添加基本的models
pwd
models/
user.go
post.go
comment.go

# 读取配置文件
pwd
config/
vim config.go


# 数据库迁移程序，系统启动前运行
pwd
test/
vim database_test.go

迁移数据库
go test ./test/ -run TestInitDB -v

插入测试数据
go test ./test/ -run TestCreateData -v
