# go module

在go1.16版本发布后，go module由原来的默认值 `auto` 变为 `on` 了，这意味着后续开发中，go更推荐用go module 模式开发，而不是gopath模式开发了。

在之前，我也是大多数以go module模式进行golang开发，但至今对其不熟悉，仅仅停留在：`别人是这样做的，我跟着做就是了` ，这都算不上会使用go module， 更不必说熟悉或者精通了；在此之前，我会存在这些疑问：

* go mod文件中定义的各项内容代表什么；
* 除了常见的 `require` 、偶尔见 `replace` 关键字外，`exclude` 、`retract` （1.16）这些关键字是什么，怎么用；
* go mod文件语法格式是什么，目前除了跟着别人写，好像也不明白其中的语法
* `github.com/tal-tech/go-zero v1.1.5` 、`github.com/antlr/antlr4 v0.0.0-20210105212045-464bcbc32de2`
  、`google.golang.org/protobuf v1.25.0 // indirect` 等格式分别代表什么，为什么有的还有
  `// indirect` 修饰；
* go.mod下面为什么有一个go.sum，其有什么作用；
* ...

不知道有多少人和我一样，对go module的了解微不足道。

最近，带着这些疑惑，去学习了官方的参考手册，这些疑惑就引刃而解了。

## project

在正式进入module介绍前，有必要首先了解一下project和module的关系，相信开发过Android或者Java的同学对module有非常好的理解，通俗的讲，一个project可以有多个module组成，module可以作为独立
的project被别的project作为依赖引用，如下golang工程 `demo` 中就包含了 `foo` 和 `bar` 两个module

```text
demo
├── bar
│   └── go.mod
└── foo
    └── go.mod
```

## module介绍

go module(以下称：module、模块、工程模块)
是golang中已发布版本的package的集合，是Go管理依赖的一种方式，类似Android中的Gradle，Java中的Maven，当然，他们的管理形式肯定是大相径庭，但是目的都是一致的，对依赖进行管理。

在go.mod中，其包含了main module的module路径、module依赖及其关联信息（版本等），如果一个工程模块需要以
`go module mode`（module模式）开发，在工程模块的根目录下必须包含 `go.mod` 文件。

## module path(module路径)

module路径是一个工程模块中的名称，在go.mod中以 `module` 命令声明，其也是工程模块中package import的前缀，我们来看一下 `demo/foo` 下的module路径：

```shell
$ cat demo/foo/go.mod
```

```text
module github.com/foo

go 1.16

require github.com/tal-tech/go-zero
```

`github.com/foo` 为main module `foo` 模块的模块路径，`github.com/tal-tech/go-zero` 也是module path, 他们也是foo中package
import的前缀，我在foo下的base包下添加了一个 `Echo` 函数，然后在 `main.go` 中调用，我们观察一下其package import的前缀

目录树

```text
foo
├── base
│   └── base.go
├── go.mod
└── main.go
```

main.go

```golang
package main

import "github.com/foo/base" // github.com/foo 为 module path
import "fmt"

func main() {
	msg := base.Echo("go-zero")
	fmt.Println(msg)
}
```

module路径主要作用是描述一个工程模块的作用是什么，在哪里可以找到，因此，module path的组成元素就包含了

* repo 路径
* repo 文件夹
* 版本

其表现形式如: `{{.repo_url}}/{{.nameOfDir}}/{{.version}}`， 示例：`github.com/tal-tech/go-zero/v2` ，`github.com/tal-tech`
明确告知了 `repo` 的路径，`go-zero` 即为repo的文件夹，`v2` 即为版本号 版本一般 `v1` 一般都默认不写了，只有大于 `v1` 时则需要用来区分。

## 版本号

这里的版本号是指go.mod文件中依赖module的版本，这和上文的 `{{.version}}` 会有关联，如下示例中的 `v1.1.5` 即为本次所说的module依赖版本号。

```text
module github.com/foo

go 1.16

require github.com/tal-tech/go-zero v1.1.5
```

### 版本号组成及规则

版本号由 `major version`（主要版本）、 `minor version`（次要版本）、 `patch version`（修订版本）组成；

* major version：指module中内容作了向后不兼容的更改后，则版本会upgrade，在此版本号upgrade时，`minor version` 和 `patch version` 要归零；
* minor version：指在新的功能发布(features)或者作了向后兼容的内容变更后，此版本号会upgrade，在此版本号upgrade时，`patch version` 要归零；
* patch version：指有bug修复或者功能优化时，此版本号可以进行upgrade，在有pre-release发布需求时也可以变更此版本号

示例：`v0.0.0`、 `v1.2.3`、 `v1.2.10-pre`

> 如果一个版本的 `major version` 为 `0` 或者 `patch-version` 有版本后缀（如:pre），则认为这个版本是不稳定的，如`v0.2.0`、 `v1.5.0-pre`、 `v1.1.3-beta`
> 更多关于version语义定义可以参考[《Semantic Versioning 2.0.0》](https://semver.org/spec/v2.0.0.html)


Golang除此之外，还可以用一些标记、分支来代表某一个版本，如：`github.com/tal-tech/go-zero 39540e21d249e91f89d96d015a6e3795cfb2be44`
、 `github.com/tal-tech/go-zero v1.1.6-0.20210303091609-39540e21d249`、
`github.com/tal-tech/go-zero@master`

其中`github.com/tal-tech/go-zero v1.1.6-0.20210303091609-39540e21d249` 这种版本在golang里面称为 `Pseudo-versions`
（伪版本），其没有完全遵循上文中的版本规则，伪版本由三个部分组成：

* 版本基本前缀（ `vX.Y.Z-0` 或 `vX.0.0` ），如 `v1.1.6-0`
* 时间戳：即revision的创建时间戳，如 `20210303091609`
* revision标识符，如 `39540e21d249`

根据基本前缀的不同，伪版本会有三种形式：

* `vX.0.0-yyyymmddhhmmss-abcdefabcdef` ：在没有 `release` 版本时使用
* `vX.Y.Z-pre.0.yyyymmddhhmmss-abcdefabcdef` ：当基本版本是预发布版本时使用
* `vX.Y.(Z+1)-0.yyyymmddhhmmss-abcdefabcdef` ：当 `release` 版本类似 `vX.Y.Z`
  时使用，如 `github.com/tal-tech/go-zero v1.1.6-0.20210303091609-39540e21d249` 的 `release` 版本为 `v1.1.5`

> 伪版本不需要手动输入，其会在执行部分go命令获取某一次提交记录版本(revision)的代码作为依赖时，会自动将其转换为伪版本
>
> 上文中的 `github.com/tal-tech/go-zero v1.1.6-0.20210303091609-39540e21d249` 则是在执行 `go get github.com/tal-tech/go-zero 39540e21d249e91f89d96d015a6e3795cfb2be44` 后自动转换的结果

### 版本后缀

为了向前兼容，如果 `major version` 升级到2时，模块路径必须要指定一个版本后缀 `v2`（其数值保持和版本中的 `major version` 的值一致），在 `major version` 小于2时，不允许使用版本后缀。

我们来看一个例子，我在 `demo/foo` 模块工程中用到了 `miniRedis` 这个库，该库的 `major version` 已经升级到2了，假设我在go.mod中引用如下版本会怎么样？

```text
module github.com/foo

go 1.16

// 正确引入
// require github.com/alicebob/miniredis/v2 v2.14.1 

// 错误引入
require github.com/alicebob/miniredis v2.14.1
```

```text
invalid version: module contains a go.mod file, so major version must be compatible: should be v0 or v1, not v2
```

上面是go mod使用时的场景，我们来看一下当main module的release版本升级到 `v2.x.x` 时(前提先发一个v1.0.0的版本)，module path没有添加版本后缀，在另一个module去使用它会有什么效果：
示例module [foo](https://github.com/anqiansong/foo)
`foo` 工程目前有release版本

* `v1.0.0`
* `v2.0.1`
* ...

使用module的工程 `bar`

* 未添加版本后缀前

`foo` 工程目录树

```text
foo
├── echo
│   └── echo.go
└── go.mod

```

module path 为 `github.com/anqiansong/foo`

```text
module github.com/anqiansong/foo

go 1.16
```

在工程 `bar` 的go.mod使用 `v2.0.0` 版本

```text
require github.com/anqiansong/foo v2.0.0
```

你会发现报错内容为

```text
 require github.com/anqiansong/foo: reading https://goproxy.cn/github.com/anqiansong/foo/@v/v2.0.0.info: 404 Not Found
	server response: not found: github.com/anqiansong/foo@v2.0.0: invalid version: module contains a go.mod file, so major version must be compatible: should be v0 or v1, not v2
```

如果 `require github.com/anqiansong/foo v1.0.0` 是可以的。

* 添加版本后缀后

`foo` 工程目录树

```text
foo
├── echo
│   └── echo.go
└── go.mod

```

module path 为 `github.com/anqiansong/foo/v2`

```text
module github.com/anqiansong/foo/v2

go 1.16
```

在工程 `bar` 的go.mod使用 `v2.0.1` 版本

```text
require github.com/anqiansong/foo/v2 v2.0.1
```

`bar` 运行正常


> 如果 `major version` 升级至 `v2` 时，如果该版本没有打算向前兼容，且不想把module path添加版本后缀，则可以在build tag时以 `+incompatible` 结尾即可，
> 则别的工程引用示例为 `require github.com/anqiansong/foo v2.0.0+incompatible`

## 如何解析package中的module

Go 命令首先在构建列表中搜索具有包路径前缀的模块。例如，如果导入了包 `example.com/a/b` ，而模块 `example.com/a` 位于构建列表中， 则 go 命令将检查 `example.com/a`
是否包含目录 `b` 中的包。且该目录中至少包含一个go文件，这样才能被视为 `package` 。生成约束不应用于此目的。 如果生成列表中只有一个模块提供包，则使用该模块。如果没有模块提供包，或者有两个或多个模块提供包，则 `go`
命令报告错误。`mod=mod` 标志指示 `go` 命令尝试查找提供丢失包的新模块， 并更新 `go.mod` 和 `go.sum` 。`go get` 和 `go mod tidy` 命令会自动执行此操作。

当go命令更新或者获取module依赖时，其会检查 `GOPROXY` 环境变量，`GOPROXY` 的值是一个逗号分割的url列表，或者是关键字 `direct` 、 `off`，

* 逗号分割的具体url为代理地址，其会告知 `go` 命令以此值去发起连接
* `direct`： 指定module依赖通过版本控制系统去获取
* `off`： 表示不尝试连接获取module

如果 `GOPROXY` 设置了具体的url，假设 `go` 命令要寻找一个`github.com/tal-tech/go-zero/zrpc`的 `package`，`go` 命令会并行的去查找一下module

* `github.com/tal-tech/go-zero/zrpc`
* `github.com/tal-tech/go-zero`
* `github.com/tal-tech`
* `github.com`

如果其中有一个或者多个匹配到包含满足 `github.com/tal-tech/go-zero/zrpc` 的内容，则取最长的 `module`作为依赖，在找到合适的module和版本后，`go` 命令会向 `go.mod`
和 `go.sum` 文件中填写`require`， ，如果解析到的 `module` 不是 main module 主动引入的，则会在 `require` 的值后面添加 `// direct`
注释，如果一个都没有匹配到，则报错；如果 `GOPROXY` 有多个url代理，在前面失败的情况下，会依次 向后面代理执行上面的步骤。

## go.mod 文件

一个 `module`（模块工程）的标识是在其根目录下包含一个编码为 `UTF-8`、名称为 `go.mod` 的文本文件，`go.mod` 文件中的内容是面向`行`的，每一行包含一个指令，且每行均由一个 `关键字` 和 `参数`
组成，就像：

```text
module github.com/anqiansong/foo

go 1.16

require github.com/tal-tech/go-zero
require github.com/tal-tech/go-queue

replace go.etcd.io/etcd => go.etcd.io/etcd v0.0.0-20200402134248-51bdeb39e698

retract [v1.0.0, v1.0.1]
```

当然，拥有相同关键字的内容可以分离出来，用 `关键字` + `block`组成，就像：

```text
module github.com/anqiansong/foo

go 1.16

require (
    github.com/tal-tech/go-zero
    github.com/tal-tech/go-queue
)

replace go.etcd.io/etcd => go.etcd.io/etcd v0.0.0-20200402134248-51bdeb39e698

retract [v1.0.0, v1.0.1]
```

`go.mod`是机器可写的，像执行一些命令（如： `go get` 、 `go mod edit`）可能会自动更新 `go.mod`文件。

## module 组成元素

在解析 `go.mod`文件中的内容时，其会被解析为

* `空白符` ： 包含空格(U+0020)、制表符(U+0009)、回车(U+000D)和换行符(U+000A)
* `注释` ：注释仅支持单行注释 `//`
* `标点` ：标点符号有 `(` 、 `)` 、 `,` 、 `=>`
* `关键字` ： `go` 、 `require` 、 `replace` 、 `exclude` 、 `retract`
* `标识符` ：由非 `空白符` 组成的字符序列，如 module path、语义版本
* `字符串` ：由英文双引号 `"` （U+0022）包裹的解释字符串或者有 `<` （U+0060）包裹的原始字符串。如`"github/com/tal-tech/go-zero"`、``

> `标识符` 和 `字符串`在 `go.mod`语法中可以替换

## module 语法词法

`go.mod` 语法是通过`Extended Backus-Naur Form` (EBNF范式) 定义的，就像

```BNF
GoMod = { Directive } .
Directive = ModuleDirective |
            GoDirective |
            RequireDirective |
            ExcludeDirective |
            ReplaceDirective |
            RetractDirective .
```

### `module` 指令

`module` 关键字定义了main module的module path，在 `go.mod` 文件中有且只有一个 `module` 指定。

语法规则：

```text
ModuleDirective = "module" ( ModulePath | "(" newline ModulePath newline ")" newline .
```

示例：

```text
module github.com/tal-tech/go-zero
```

### `go` 指令

`go` 关键字定义了 `module` 设置预期使用的go语言版本，版本必须是一个有效的go版本（可以理解为符合 `version` 规则，也可以理解为为Go已经release的版本）

通过 `go` 关键字定义版本后，编译器在编译包时就知道应该使用哪个go版本去编译，除此外，`go` 关键字定义版本还可以用于 是否启用 `go`命令的一些特性，如是否自动开始vendoring在版本 `1.14` 及以后。

语法规则:

```text
GoDirective = "go" GoVersion newline .
GoVersion = string | ident .  /* valid release version; see above */
```

示例：

```text
go 1.16
```

### `require` 指令

`require` 声明了module依赖的最小版本，在 `require` 指定版本后，`go` 相关命令会根据 [MVS](https://golang.org/ref/mod#minimal-version-selection)
规则根据此值来加载依赖。

`go` 寻找依赖时，如果该依赖不是main module直接依赖的，则会在该module path 后面添加 `// direct` 注释内容。

语法规则：

```text
RequireDirective = "require" ( RequireSpec | "(" newline { RequireSpec } ")" newline ) .
RequireSpec = ModulePath Version newline .
```

示例：

```text
module github.com/tal-tech/go-zero

go 1.16

require (
    golang.org/x/crypto v1.4.5 // indirect
    golang.org/x/text v1.6.7
)
```

### `excule` 指令

`excule` 会忽略内容中的指定版本，从 `go 1.16` 后，`exclude` 指定的module会被忽略， 在 `go 1.16` 前， 如果 `require` 的module被 `exclude`
指定后，会列出并获取更改的为被 `exclude` 的版本。

语法规则：

```text
ExcludeDirective = "exclude" ( ExcludeSpec | "(" newline { ExcludeSpec } ")" ) .
ExcludeSpec = ModulePath Version newline .
```

示例：

```text
module github.com/tal-tech/go-zero

go 1.16

exclude golang.org/x/net v1.2.3

excule (
    golang.org/x/crypto v1.4.5
    golang.org/x/text v1.6.7
)
```

### `replace` 指令

`replace` 指令用于将module的指定版本或者module使用其他的module或者版本来替换，如果 `=>` 左边质指定了版本，则替换 这个版本至目标内容，否则替换替换module的所有版本至目标内容

语法规则：

```text
ReplaceDirective = "replace" ( ReplaceSpec | "(" newline { ReplaceSpec } ")" newline ")" ) .
ReplaceSpec = ModulePath [ Version ] "=>" FilePath newline
            | ModulePath [ Version ] "=>" ModulePath Version newline .
FilePath = /* platform-specific relative or absolute file path */
```

示例：

```text
replace golang.org/x/net v1.2.3 => example.com/fork/net v1.4.5
```

`=>` 右边的内容可以是有效的module path，也可以是相对或者绝对路径，如果是相对或者绝对路径，这该路径的根目录必须包含 `go.mod` 文件。

示例：

```text
require github.com/foo v1.0.0

replace github.com/foo v1.0.0 => ../bar
```

### `retract` 指令(1.16新增)

`retract` 声明的内容，用于标记某些版本或者某个版本范围（闭区间）标记为撤回，一般 `retract` 声明前需要写一条注释用于说明，当执行 `go get` 命令时， 如果引用了被标记为 `retract` 的版本，或者
在 `retract` 标记的版本范围内，则会提示一条警告（其内容为 `retract` 的注释内容），通过`go list -m -versions` 获取版本时也会隐藏该版本。

语法规则：

```text
RetractDirective = "retract" ( RetractSpec | "(" newline { RetractSpec } ")" ) .
RetractSpec = ( Version | "[" Version "," Version "]" ) newline .
```

示例：

```text
// someting wrong
retract v1.0.0

// someting wrong in range of versions => v1.0.0~v1.2.0 
retrace [v1.0.0,v1.2.0]
```

我们来看一个例子，目前 `github.com/anqiansong/retract` 已经有 `v1.0.0` 等版本了，我们 添加一行 `retract` 指令标记
`v1.0.0` 撤回：

```text
// someting wrong
retract v1.0.0
```

然后release一个版本为 `v1.0.1` ，接下来在 `github.com/anqiansong/bar` 中引用 `v1.0.0`版本

```text
require github.com/anqiansong/retract v1.0.0
```

然后执行 `go get github.com/anqiansong/retract@v1.0.0` ， 不出意外，会得到一个提示包含 `something wrong` 和 提示更新到 `v1.0.1` 的信息

```shell
$ go get github.com/anqiansong/retract@v1.0.0
```

```text
go: warning: github.com/anqiansong/retract@v1.0.0: retracted by module author: someting wrong
go: to switch to the latest unretracted version, run:
        go get github.com/anqiansong/retract@latestgo get: downgraded github.com/anqiansong/retract v1.0.1 => v1.0.0
```

获取 `github.com/anqiansong/retract` 所有module release版本

```shell
$ go list -m -versions github.com/anqiansong/retract
```

```text
github.com/anqiansong/retract v1.0.1
```

部分命令查看 `github.com/anqiansong/retract@v1.0.0`的结果：

* `go get`
    ```shell
    $ go get github.com/anqiansong/retract@v1.0.0
    ```
    ```text
    go: warning: github.com/anqiansong/retract@v1.0.0: retracted by module author: someting wrong
    go: to switch to the latest unretracted version, run:
    go get github.com/anqiansong/retract@latest
    ```
* `go list -m -u`
    ```shell
    $ go get github.com/anqiansong/retract@v1.0.0
    ```
    ```text
    github.com/anqiansong/retract v1.0.0 (retracted) [v1.0.1]
    ```

* `go list -m -versions`
  ```shell
    $ go list -m -versions github.com/anqiansong/retract
    ```
    ```text
    github.com/anqiansong/retract v1.0.1
    ```

> 说明：
>
> `retract` 控制的是main module的版本，而非依赖的module 版本。
>
> 被 `retract` 标记的版本其他module还是可以引用的，只是部分 `go` 命令执行时会有 `retract` 的不同结果，如上。

## 自动更新

如果 go.mod 缺少信息或者不能准确反映实际情况，大多数 `go` 命令都会报告错误。`go get` 、 `go mod tidy` 命令可以用来修复大多数这类问题。 此外，`-mod=mod`
标志可以用于大多数模块感知命令(`go build`、 `go test` 等) ，以指示 `go` 命令自动修复 `go.mod` 和 `go.sum` 中的问题。

## module 感知

大多数 go 命令可以在 `Module-ware` 模式或 `GOPATH` 模式下运行。在 `Module-ware` 模式下，go 命令使用 go.mod 文件来查找版本相关性，
它通常从模块缓存中加载包，如果缺少模块，则下载模块。在 `GOPATH` 模式下，go 命令忽略module； 它在 `vendor` 或 `GOPATH` 目录中查找依赖项。

在 Go 1.16中，无论是否存在 `go.mod` 文件，`Module-ware` 模式默认是启用的。在低版本中，当工作目录文件或任何父目录中存在 `go.mod` 文件时，启用 `Module-ware` 模式 。

`Module-ware` 模式可以通过GO111MODULE 环境变量控制，可以设置为 `on`、`off` 或 `auto`

* `off` ：`go` 相关命令会忽略 `go.mod`文件，然后以 `GOPATH` 模式运行
* `on` ： `on` 或者空字符串， 相关命令会 `Module-ware` 模式运行
* `auto`： 如果当前文件夹存在 `go.mod` 文件，则会以 `Module-ware` 模式运行，在 Go 1.15及更低版本，此值为默认值，

## 部分go module相关命令

这里命令必须要在 `Module-ware` 模式才有效

|命令|用法|备注|示例|
|---|---|---|---|
|go list -m|go list -m [-u] [-retracted] [-versions] [list flags] [modules]|查看module信息|go list -m all|
|go mod init|go mod init [module-path]|在工作目录初始化并创建一个go.mod文件|go mod init demo|
|go mod tidy|go mod tidy [-e] [-v]|整理go.mod文件|go mode tidy|
|go clean -modcache|go clean [-modcache]|清除module缓存|go clean -modcache|

## Proxy

模块代理是一个支持 `GET` 请求响应 的`HTTP`服务器，该请求没有 `query` 参数，甚至不需要特定的 `header` 信息，即使 该值是一个固定的文件系统站点(如：`file:// URL` )也是可以的。

模块代理的 `HTTP` 响应状态码必须包含 `200`（OK），`3xx`，`4xx`、`5xx`，`4xx`、`5xx`被认为是响应错误，`404` 和 `410`
表示所有的module请求是不可用的，注意，错误的响应的contentType 应该设置为 `text/plain`，字符集为 `utf-8` 或者 `us-ascii` 。

### URLs
`go` 命令可以通过读取 `GOPROXY` 环境变量配置来连接连接代理服务器或者版本控制系统，`GOPROXY` 接受一个逗号(,)或者竖线(|)分割的多个url值， 当以英文逗号(,)
分割时，只有响应状态码为404或者410时就会尝试后面的代理地址，如果是以竖线(|)分割，则在http出现任何错误（包含超时）都会跳过去尝试后面的代理地址。 也可以是 `direct` 或者 `off` 关键字。

下面的表格为一个代理地址必须要实现且有请求响应的path（即一个代理服务器必须要要支持一下路由的实现）

* `$base`为代理服务器地址，如：https://goproxy.cn
* `$module`为module path，如：github.com/tal-tech/go-zero
* `$version`为module 版本

<table>
  <tr>
    <td>path</td>
    <td>描述</td>
    <td>示例</td>
    <td>示例结果</td>
  </tr>
  <tr>
    <td>$base/$module/@v/list</td>
    <td>以纯文本形式返回给定模块的已知版本的列表，每行一个。此列表不应包括伪版本</td>
    <td>curl -X GET https://goproxy.cn/github.com/tal-tech/go-zero/@v/list</td>
    <td>
        v1.0.0<br>
        v1.0.1<br>
        v1.0.2<br>
        v1.0.3<br>
        v1.0.4<br>
        ...<br>
    </td>
  </tr>

  <tr>
    <td>$base/$module/c/$version.info</td>
    <td>返回关于某个模块的特定版本的 json 格式的元数据。响应必须是一个 JSON 对象，对应于下面的 Go 数据结构:
  <pre>
  type Info struct {
      Version string    // version string
      Time    time.Time // commit time
  }
  </pre>
    </td>
    <td>curl -X GET https://goproxy.cn/github.com/tal-tech/go-zero/@v/v1.1.5.info</td>
    <td>
    <pre>{
  "Version": "v1.1.5",
  "Time": "2021-03-02T03:02:57Z"
}</pre>
    </td>
  </tr>

<tr>
    <td>$base/$module/@v/$version.mod</td>
    <td>返回指定版本的go.mod中的信息/td>
    <td>curl -X GET https://goproxy.cn/github.com/tal-tech/go-zero/@v/v1.1.5.mod</td>
    <td>
    <pre>module github.com/tal-tech/go-zero

go 1.14

require (
github.com/ClickHouse/clickhouse-go v1.4.3
github.com/DATA-DOG/go-sqlmock v1.4.1
github.com/alicebob/miniredis/v2 v2.14.1
github.com/antlr/antlr4 v0.0.0-20210105212045-464bcbc32de2
....
)
</pre>
    </td>
  </tr>

<tr>
    <td>$base/$module/@v/$version.zip</td>
    <td>返回指定版本的go module的zip文件</td>
    <td>wget https://goproxy.cn/github.com/tal-tech/go-zero/@v/v1.1.5.zip</td>
    <td> v1.1.5.zip </td>
  </tr>

<tr>
    <td>$base/$module/@latest</td>
    <td>返回有关模块的最新已知版本的 json 格式的元数据</td>
    <td>curl -X GET https://goproxy.cn/github.com/tal-tech/go-zero/@latest</td>
    <td> <pre>{
  "Version": "v1.1.5",
  "Time": "2021-03-02T03:02:57Z"
}</pre> </td>
  </tr>
</table>

在获取module最新版本时， `go` 相关命令会优先请求 `$base/$module/@v/list` 地址，如果没有找到适合的版本，
则请求 `$base/$module/@latest` 获取并匹配是否满足，`go` 相关命令会按照 `release` 版本、 `pre-release` 版本、 `pseudo` 版本排序。

而 `$base/$module/$version.mod` 和 `$base/$module/$version.zip` 地址必须要提供，因为这些信息可以用于和 `go.sum` 进行数据校验。

go module下载后的内容一般会存储在 `$GOPATH/pkg/mod/cache/download` 路径下，包括版本控制系统下载的也是如此。

### direct
如果 `GOPROXY` 设置了 `direct` 值，则在执行相关 `go` 命令时会从版本控制系统(`git`、`svn`等)下载module资源，当然还需要另外两个环境
变量 `GOPRIVATE` 、 `GONOPROXY` 配合，更多环境变量信息请参考[这里](https://golang.org/ref/mod#environment-variables)

## module文件大小约束
对模块 `zip` 文件的内容有许多限制。这些约束确保可以在广泛的平台上安全和一致地提取压缩文件。

一个模块最多可以达到500MiB（不管是压缩形式还是未压缩形式文件）

`go.mod`文件要求显示在16MiB以内

## go.sum
在module的根目录下可能有一个名为 `go.sum`的文件，其是直接依附在go.mod下面的， 当执行 `go` 相关命令获取module时，
其会检查 `zip` 文件和 `go.sum` 中的值是否一致。

`go.sum` 中的值由一个 `module path`、`version` 和一个`hash`值组成，如：
```text
github.com/anqiansong/retract v1.0.1 h1:jxcsUM/6tvxM7p14/XMeZPFbql5KAAZJfFqiHG+YKxA=
```

# 总结
花了3天时间，第一天看英文原文档，第二、三天翻译成中文来验证自己的理解，并写下这篇日志，经过学习，对于go module手册的学习解开了不少疑惑，
至少开篇的几个问题算是解开了，本文也是将手册中的部分内容作了 `搬运` 结合了自己的一些理解，
以及用真实例子去验证， 其中还有很多内容我并未完全照搬，如果有兴趣的同学，可以参考官方手册，本文部分内容包含个人理解，也是结合google翻译加上
对一些翻译不合理的地方进行人工翻译，肯定存在不合理的地方，如果有理解不一致的，欢迎提出和讨论。


# 参考文档
* [《Semantic Versioning 2.0.0》](https://semver.org/spec/v2.0.0.html)
* [《Go Module官方文档》](https://golang.org/ref/mod)


