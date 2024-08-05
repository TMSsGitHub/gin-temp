# gin-temp

## 手动下载依赖
```shell
go mod tidy
```
## 安装go插件
### 安装swag
```shell
go install github.com/swaggo/swag/cmd/swag@latest
```

### 安装fresh
fresh用于热部署，可不安装
```shell
go install github.com/gravityblast/fresh@latest
```

## 复制配置文件
启动项目前先复制配置文件并在本地进行自定义配置，比如mysql和redis的配置
```shell
cp application_test.yml application.yml
```

## 生成swagger docs
```shell
swag init
```

## 启动项目
如果安装了fresh可使用fresh命令启动项目
```shell
fresh
```

如果没有安装则可以使用以下命令启动项目
```shell
go run main.go
```

可以在浏览器打开`/swagger/index.html`或者在goland中打开`docs/swagger.json`和`docs/swagger.yaml`其中一个都可以

>P.S.
>启动前记得建mysql表和修改自己的application.yml