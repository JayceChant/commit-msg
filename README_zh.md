# commit-msg

一个 go 实现的 git commit-msg 钩子。

[English](./README.md) | 中文



## 安装

只需将二进制文件放到 git 管理的项目的 hook 目录里。

假设项目的根目录是 `/`, 那么钩子目录应该为 `/.git/hooks/` 。请确认文件名为 `commit-msg` 或 `commit-msg.exe` (win) ， 并且文件有可执行权限。

如果你找不到 hook 目录，请确认是否有设置隐藏目录可见。



安装之后，git 会自动调用程序完成 commit message 的检查，无需手动调用。

### 从编译好的二进制安装

你可以在 [这里](https://github.com/JayceChant/commit-msg/releases) 下载编译好的二进制文件。 不过，暂时只提供 win64 平台。



### 从源码编译并安装

对于 win64 以外的平台，请自行配置 go 开发环境并编译安装。

因为时间关系，抱歉没有更详细的指引。



## 配置

配置文件可以放在两个地方：

* 全局配置：`$HOME/.commit-msg.json`
* 项目配置：`project/.git/hooks/.commit-msg.json`

两个配置文件可以都不设置或者都设置，也可以选择只有其中一个。程序会先尝试加载全局配置，再加载项目配置，两个配置文件都设置的项，以项目配置为准。

先看一下默认的 commit-message 格式：

```
<type>(<scope>): <subject>
// 空行
<body>
// 空行
<footer>
```

例如：

```
feat(model): some changes to model layer

* change 1
* change 2
other notes

BREAKING CHANGE: some change break the compatibility

the way to migrate:
...
```

支持的配置项如下：

* `lang`：提示语言。目前只支持内置的 `en` 和 `zh` ，后续会支持添加自定义语言。
* `scopeRequired`：如果为 true，`(<scope>)` 则为必填项。
* `scopes`：是一个字符串列表，如果不为空，则 `scope` 必须从列表中取值。该设置项仅当 `scope` 非空时生效，与 `scopeRequired` 互相独立。
* `types` 和 `denyTypes`：均为字符串列表。`types` 列表的关键字会加入到默认关键字列表；`denyTypes` 的关键字则会被删除（如果有）。如果一个关键字同时出现在两个列表里，由于是先添加后删除，以 `denyTypes` 为准。
    默认的 `type` 关键字列表为
    * `feat`：新功能
    * `fix`：bug 修复
    * `docs`：文档
    * `style`：代码风格，不影响运行逻辑的修改
    * `refactor`：重构（既没有新增功能，也没有修复错误的逻辑改动）
    * `perf`：性能相关的改动
    * `test`：测试相关的改动
    * `chore`：杂项，构建过程或辅助功能相关的改动
    * `revert` 和 `Revert`：部分工具生成的 revert 信息首字母大写。
* `bodyRequired`：如果为 true，则提交信息必须包含信息体。（不能只有信息头）
* `lineLimit`：单行长度限制，对所有行生效，以字节为单位。

如果没有任何配置文件，程序将使用以下默认配置：

```json
{
    "lang": "en",
    "scopeRequired": false,
    "bodyRequired": false,
    "lineLimit": 80
}
```
