# commit-msg

一个非常轻量的 git commit-msg 钩子，用 golang 实现。

几乎零依赖。（唯一的依赖是只有一个文件的 `homedir`，无间接依赖。）

[English](./README.md) | 中文



## 安装

只需将二进制文件放到 git 管理的项目的 hook 目录里。

假设项目的根目录是 `/`, 那么钩子目录应该为 `/.git/hooks/` 。请确认文件名为 `commit-msg` 或 `commit-msg.exe` (win) ， 并且文件有可执行权限。

如果你找不到 hook 目录，请确认是否有设置隐藏目录可见。



安装之后，git 会自动调用程序完成 commit message 的检查，无需手动调用。

### 从编译好的二进制安装

你可以在 [这里](https://github.com/JayceChant/commit-msg/releases) 下载编译好的二进制文件。 不过，暂时只提供 win64 平台。



### 从源码编译并安装

#### 安装要求
- [x] `Golang`
- [x] `stringer`
```sh
go install golang.org/x/tools/cmd/stringer
```
- [x] `upx`

#### 安装方法
```sh
sudo make
```


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
    * `chore`：杂项，可能是辅助功能相关的改动
    * `build`：构建过程相关
    * `ci` ： 持续集成相关
    * `docker`：docker 镜像构建相关
    * `revert` 和 `Revert`：部分工具生成的 revert 信息首字母大写。
* `bodyRequired`：如果为 true，则提交信息必须包含信息体。（不能只有信息头）
* `lineLimit`：单行长度限制，对所有行生效，以字节为单位。如果这个值小于等于零，跳过长度检查。

如果没有任何配置文件，程序将使用以下默认配置：

```json
{
    "lang": "en",
    "scopeRequired": false,
    "bodyRequired": false,
    "lineLimit": 80
}
```

## 本地化

程序内置了两种语言的提示：英语（en） 和 中文（zh）。

你可以通过以下两种方式增加支持的语言：

### 提交代码

内置的语言支持放在 `lang` 目录下，每种语言一个 `.go` 源文件。你可以拷贝 `lang/english.go` ，修改文件名并翻译内容，然后提交，发起 Pull Request。待 PR 通过后，下一个版本就可以支持对应的语言。

当然，考虑到我的英语水平有限（中文的措辞也不见得很好），你也可以帮忙改进已有的语言支持。

### 增加语言翻译文件

但提交代码未必是最好的办法。

一方面，提交代码耗时较多。如果你不能自行编译，而是要等代码合入之后发布下一个版本，还需要等待更多的时间。另一方面，这毕竟只是一个小工具，为了避免程序变得臃肿，视乎使用者的多寡，一些语言我可能不会选择合入。

你可以选择更快捷的方式：增加语言文件。

具体的做法是，拷贝项目根目录下的 [commit-msg.en.json.sample](./commit-msg.en.json.sample) 文件，去掉文件名里的 `.sample` ，把 `en` 改为对应的语言（举例说这种语言的缩写为 `xx`，那么对应的翻译文件应该为 `commit-msg.xx.json`）。把文件内容翻译好，注意保留里面的格式化动词 `%s` 和换行符 `\n` 。然后把文件放到跟配置文件相同的目录（`home` 目录或者 `hooks` 目录）。之后记得修改语言配置为对应的语言（这里是 `xx`）。

跟配置文件不同，如果相同语言的翻译在 `home` 目录和 `hooks` 目录同时存在，并不会合并两者的内容，而是直接以项目 `hooks` 目录的翻译为准。所以每一个语言文件里的内容，都必须是完整的翻译。

## 更多信息

本项目最早受 [validate-commit-msg](https://github.com/conventional-changelog-archived-repos/validate-commit-msg) 启发。（该项目现已移至 [conventional-changelog/commitlint](https://github.com/conventional-changelog/commitlint)）

更多关于 `conventionalcommits` 的信息，请参考 https://www.conventionalcommits.org/
