# EZ配置布局源文件镜像转换工具
工具基于go语言，主要用于把ergodox-ez配置生成的源码进行左右镜像。

# 使用方法

本文基于ArchLinux操作，顺便安利一波，我认为最好用的桌面linux就是archlinux系列的，滚动式更新真的很棒。

# 准备工具
### 编译会用到的依赖
```
pacman -S avr-gcc avr-libc teensy-loader-cli
```
### 克隆本工具
```
git clone https://github.com/wqtty/ergodoxEZMirrorImageConverter.git
```

### qmk固件编译工具
用于把转换好的源文件编译烧入键盘的工具
```
git clone https://github.com/qmk/qmk_firmware.git
#编译子依赖库
make git-submodule
```

### 操作步骤
1. 使用ez配置器配置布局([地址](https://configure.ergodox-ez.com))
打开地址后点击Clone and modify this layout会生产可配置的页面，也可以自己搜索他人方案，基于他人方案上进行克隆修改，可以搜索常见tag诸如vim，linux，mac等等。
2. ez下载源代码 布局配置编译好之后 点击download source下载，会下载到一个一布局名称开头的压缩包，压缩包中会包含keymap.c文件。
文件名一般为ergodox_ez_firmware_xxxxxx.c后面6位字母。
3. 拷贝文件到镜像工具目录
4. 生成镜像文件。会生成keymap.c文件
```
go run main.go ergodox_ez_firmware_xxxxxx.c
```
5. 创建布局文件夹
```
cd ~/qmk_firmware/keyboards/ergodox_ez/keymaps
mkdir yourlayout
```
6.拷贝镜像后的文件到布局文件夹
```
cd yourlayout
cp ~/ergodoxEZMirrorImage/keymap.c ./
```
7.编译并烧入键盘
```
make ergodox_ez:yourlayout:teensy
```
编译成功后会提示
```
Waiting for Teensy device...
 (hint: press the reset button)
```
点击键盘下面的重置按钮开始烧入
读条完成后会自动重启键盘
如果仅仅只是编译的话，可以使用下面命令
```
make ergodox_ez:yourlayout
```
