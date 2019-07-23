# builder


基于 MongoDB-go-driver，json，针对更新的代码简化工具。本插件只能用于 gin+mongo 的构架中。欢迎使用和建议。

## 一，软件构架和原理

1.软件原理

基于标签，注解等，生成相关代码，简化少 go 语言编写 HTTP 中各种 dto，validation 等。

2.软件构成

在开发中，需要在原有 model 中定义如下 tag，来实现控制中间代码生成。

| tag     | 可选值 | 使用范围 | 描述                                                   |
| :------ | ------ | -------- | ------------------------------------------------------ |
| builder | post   | model    | 控制生成 \*AddDTO 和他的 toModel 方法                  |
|         | put    | updater  | 控制生成必须全部传输的\*UpdateDTO                      |
|         | patch  | updater  | 控制生成非必须全部传输的\*Match 方法中                 |
|         |        |          |                                                        |
| scope   | \*     | filter   | 控制生产*Filter 方法,控制 *UpdateDTO 和 \*Match 方法。 |
| bind    | \*     | updater  | 主要用于值校验，在*AddDTO 和 *UpdateDTO 中,是直接复制  |
|         |        |          |                                                        |

通过上面的 tg，可以控制对应对应代码。这样减少编码和设计思考。

### scop 标签

scope 标签，用于控制字段权限。用于控制 get 方法的返回值，不填表示这个字段的值是 public,具体指查阅权限设计表。scope 值直接和 role 相关。理论上来说 role 是动态可变的，但是 scope 一般都是固定的。下面是 scope 和其意义。

| scope 值 | 描述                                                                     |
| -------: | ------------------------------------------------------------------------ |
| public/0 | 公共数据，也是默认数据                                                   |
|        1 | 非内部人员能查阅和修改的数据，于用户相关的数据。例如用户表中用户 name 等 |
|        2 | 运营人员能查阅的数据                                                     |

当有较高的 scope 值时，就能操作低 scope 的 feild。从设计角度上来说，这个是不能控制权限的， 但是通过 rbac，菜单等，基本就能实现 erp 的全部权限值。这样设计的另一个好处是，不必手动写对应的 Validation 和 DTO。

## 二，软件使用

1.安装软件

```shell
go get gitlab.com/liuchamp/mhupdater
```

2.在编码完成项目后，通过如下命令生成对应代码。

```shell
mhupdater init -f source -o target
```

其中 source 和 target 都是对应具体的 package。source 默认值是 models,target 默认值也是 models。参数列表和默认值如下：

3.在生成对应代码后，还需要进一步编码。默认有如下规则：

- \*AddDTO 在 Post 方法中使用。
- \*UpdateDTO 在 put 方法中使用,当值不存在时，会报错。
- \*Match 方法主要是更新部分字段，可根据参数设置 put,patch 关系。方法参数如下：

  | 范围 | 参数名   | 类型                   | 描述                           |
  | ---- | -------- | ---------------------- | ------------------------------ |
  | i    | valueMap | map[string]interface{} | put 方法 body 里面的映射值     |
  | i    | scopes   | []string               | 请求时，用户的 scope 值        |
  | o    | updater  | interface{}            | 更新操作的具体值               |
  | o    | err      | error                  | 代码映射和权限构建过程中的错误 |

- \*Filter 方法在查询时，控制返回字段，也就是在 MongoDB 查询时，需要的 filter 选项。 下面是参数和使用方式:

  | 范围 | 参数名 | 类型        | 描述                           |
  | ---- | ------ | ----------- | ------------------------------ |
  | i    | scopes | []string    | 请求时，用户的 scope 值        |
  | o    | filter | interface{} | 过滤条件                       |
  | o    | err    | error       | 代码映射和权限构建过程中的错误 |

当然也可以在这个之上进行拓展代码。比如生成\*Sort 方法，来控制查询结果排序等。

出来可以控制 models(默认-o 参数)目录外，还可以控制某个文件的代码生成。只不过是将参数改为-f。
