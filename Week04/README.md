### Week04 作业：
Q: 按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。

A: 
```bash
☁  Week04 [main]  tree
.
├── README.md
├── api
│   └── pic
│       └── v1
│           ├── pic.pb.go
│           └── pic.proto
├── cmd
│   └── pic
│       └── main.go
├── configs
│   └── server.yaml
├── go.mod
├── go.sum
├── internal
│   ├── biz
│   │   └── pic.go
│   ├── data
│   │   └── pic.go
│   ├── pkg
│   │   └── server
│   │       └── pic.go
│   └── service
│       └── pic.go
└── test
    └── pic
        └── client.go
```