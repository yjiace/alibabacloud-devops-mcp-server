使用go 1.24.0或以上版本

## 创建二进制文件
go build -o mcp-yunxiao main.go

## cp文件到$PATH
cp mcp-yunxiao /usr/local/bin/ 

## 配置mcp server
设置mcp命令： mcp-yunxiao
设置环境变量： YUNXIAO_ACCESS_TOKEN,值为云效的个人token
