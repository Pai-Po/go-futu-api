# 富途 OpenAPI Golang SDK

## 简介

- Go 语言封装的[富途牛牛 OpenAPI](https://openapi.futunn.com/futu-api-doc/)。
- 尽量接近 Python 版本的使用方法。
- 利用 Go 语言特性，例如 channel，goroutine 等。

## 使用方法

1. import

   ```
   import "github.com/Pai-Po/go-futu-api"
   ```

1. 创建 API 实例

   ```
   ft := futuapi.NewFutuAPI()
   ```

1. 设置连接参数 (可选)

   ```
   ft.SetClientInfo("MyFutuAPI", 0)
   ```

1. 连接 FutuOpenD

   ```
   ft.Connect(context.Background(), ":11111")
   ```

1. 调用获取方法

   ```
   sub, err := ft.QuerySubscription(context.Background(), true)
   ```

1. 接收推送

   ```
   ch, err := api.UpdateTicker(context.Background())
   // ch 为 channel类型，<- ch接收推送
   ```
