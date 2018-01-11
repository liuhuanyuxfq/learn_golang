# Go语言环境安装
## 一、windows下安装
### 下载
- 具备网络条件的同学可以去[官网](https://golang.org/)下载
- 不具备条件的同学可以去[Golang中国](https://www.golangtc.com/download)下载
### 安装
运行下载的msi安装文件，运行它，一路Next。注意路径中不要包含中文，否则可能会有问题。
### 环境变量
如果正常安装，那么GOROOT会自动写进环境变量，如果没有写进去，那么可以手动添加：

- GOROOT设置成刚才的安装目录
- GOPATH为你的项目目录，随意设置
- PATH追加;%GOROOT%\bin
### 验证安装是否成功
打开cmd后，输入：

	go version
安装成功会显示版本信息
## 二、linux下安装
### 下载
	cd /usr/local
	wget https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz
### 在/usr/local下安装程序
	tar -zxvf go1.8.3.linux-amd64.tar.gz
### 配置全局变量
	vi /etc/profile
在这个文件末尾加入下面内容：

	export GOPATH=/root/gopath
	export GOROOT=/usr/local/go
	export PATH=$PATH:$GOROOT/bin
### 加载使生效
	source /etc/profile
### 验证是否生效
	go version
### GOPATH目录结构：
	/root/gopath/
	    --/src/  存放源代码(.go .c .h 等)
	    --/pkg/　　编译后生成的文件(.a)
	    --/bin/　　编译后生成的可执行文件