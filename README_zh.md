# commit-msg

一个 go 实现的 git commit-msg 钩子。



## 安装

只需将二进制文件放到 git 管理的项目的 hook 目录里。

假设项目的根目录是 `/`, 那么钩子目录应该为 `/.git/hooks/` 。请确认文件名为 `commit-msg` 或 `commit-msg.exe` (win) ， 并且文件有可执行权限。

如果你找不到 hook 目录，请确认是否有设置隐藏目录可见。



安装之后，git 会自动调用程序完成 commit message 的检查，无需手动调用。

### 从编译好的二进制安装

你可以在 [这里](https://github.com/JayceChant/commit-msg/releases) 下载编译好的二进制文件。 不过，暂时只提供 win64 平台。



### 从源码编译并安装

请自行配置 go 开发环境并编译安装。

因为时间关系，抱歉没有更详细的指引。