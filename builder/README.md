# builder

基于 MongoDB-go-driver，json，针对更新的代码简化工具。本插件只能用于 gin+mongo 的构架中。欢迎使用和建议。

## 一，软件构架和原理

1.软件原理

基于标签，注解等，生成相关代码，简化少 go 语言编写 HTTP 中各种 dto，validation 等。

2.软件构成

在开发中，需要在原有 model 中定义如下 tag，来实现控制中间代码生成。

| tag   | 可选值 | 使用范围 | 描述                                                   |
| :---- | ------ | -------- | ------------------------------------------------------ |
| build | post   | model    | 控制生成 \*AddDTO 和他的 toModel 方法                  |
|       | put    | updater  | 控制生成必须全部传输的\*UpdateDTO                      |
|       | patch  | updater  | 控制生成非必须全部传输的\*Match 方法中                 |
| scope | \*     | filter   | 控制生产*Filter 方法,控制 *UpdateDTO 和 \*Match 方法。 |
| bind  | \*     | updater  | 主要用于值校验，在*AddDTO 和 *UpdateDTO 中,是直接复制  |
|       |        |          |                                                        |

通过上面的 tag，可以控制对应对应代码。这样减少编码和设计思考。

### scop 标签

scope 标签，用于控制字段权限。用于控制 get 方法的返回值，不填表示这个字段的值是 public,具体指查阅权限设计表。scope 值直接和 role 相关。理论上来说 role 是动态可变的，但是 scope 一般都是固定的。下面是 scope 和其意义。默认大值有大权限。后续加入参数控制，使得起可控。

| scope 值 | 描述                                                                     |
| -------: | ------------------------------------------------------------------------ |
|      1xx | 公共数据，也是默认数据,                                                  |
|      2xx | 非内部人员能查阅和修改的数据，于用户相关的数据。例如用户表中用户 name 等 |
|      3xx | 代理商管理人员能查阅的数据                                               |
|      4xx | 运营人员能查阅的数据                                                     |
|      5xx | 运营人员能查阅的数据                                                     |
|      6xx | 运营人员能查阅的数据                                                     |
|      7xx | 运营人员能查阅的数据                                                     |
|      8xx | 运营人员能查阅的数据                                                     |
|      9xx | admin                                                                    |

当有较高的 scope 值时，就能操作低 scope 的 feild。从设计角度上来说，这个是不能控制权限的， 但是通过 rbac，菜单等，基本就能实现 erp 的全部权限值。这样设计的另一个好处是，不必手动写对应的 Validation 和 DTO。

## 二，软件使用

1.安装软件

```shell
go get gitlab.com/liuchamp/mhupdater
```

2.在编码完成项目后，通过如下命令生成对应代码。

```shell
mhupdater -h
```

其中 source 和 target 都是对应具体的 package。source 默认值是 models,target 默认值也是 models。参数列表和默认值如下：

3.在生成对应代码后，还需要进一步编码。默认有如下规则：

- \*AddDTO 在 Post 方法中使用。
- \*UpdateDTO 在 put 方法中使用,当值不存在时，会报错。

  | 范围 | 参数名   | 类型                   | 描述                       |
  | ---- | -------- | ---------------------- | -------------------------- |
  | i    | valueMap | map[string]interface{} | put 方法 body 里面的映射值 |
  | i    | scope    | int                    | 请求时，用户的 scope 值    |
  | o    | updater  | interface{}            | 更新操作                   |
  | o    | err      | error                  | 构架 updater 过程中的 Id   |

- \*Match 方法主要是更新部分字段，可根据参数设置 put,patch 关系。方法参数如下：

  | 范围 | 参数名   | 类型                   | 描述                           |
  | ---- | -------- | ---------------------- | ------------------------------ |
  | i    | valueMap | map[string]interface{} | put 方法 body 里面的映射值     |
  | i    | scope    | int                    | 请求时，用户的 scope 值        |
  | o    | updater  | interface{}            | 更新操作的具体值               |
  | o    | err      | error                  | 代码映射和权限构建过程中的错误 |

- \*Filter 方法在查询时，控制返回字段，也就是在 MongoDB 查询时，需要的 filter 选项。 下面是参数和使用方式:

  | 范围 | 参数名 | 类型        | 描述                           |
  | ---- | ------ | ----------- | ------------------------------ |
  | i    | scope  | int         | 请求时，用户的 scope 值        |
  | o    | filter | interface{} | 过滤条件                       |
  | o    | err    | error       | 代码映射和权限构建过程中的错误 |

当然也可以在这个之上进行拓展代码。比如生成\*Sort 方法，来控制查询结果排序等。

### \*Fliter 代码

主要用于获取数据时的过滤，限制返回数据集合。eg：

```Go
package filter
// 默认是0， 识别范围0-100
func UserFilter(scope int) interface{} {
	filter := bson.M{}
	if scope < xxx {
		filter[field] = bsonx.Int32(0)
	}
	if scope < xxx {
		filter[field] = bsonx.Int32(0)
	}
	return filter
}
```

    注意，没有 scope<0 这个条件

### \*Match 代码

```Go
package filter
func UserMatch(value map[string]interface{}, scope int) (updater interface{}, err error) {
	if value == nil || len(value) == 0 {
		return nil, errors.New("value nil")
	}
	up := bson.M{}
	for k, v := range value {
		if k == xxx && scope < xxx {
			up[k] = v
		}
	}
	return bson.M{"$set": up}, nil
}
```

### \*UpdateDTO 代码

在 put 方法中， 必须保证 required 字段上传，同时还需要限制传入遍历的限制。

```Go
import (
	"errors"
	"github.com/wxnacy/wgo/arrays"
	"go.mongodb.org/mongo-driver/bson"
	"windplatform/webbackend/server/models"
)
var (
	// key -> scope
	//  value -> json tags
	userScopeMap map[int][]string
	// key -> json tag
	// value -> bson tag
	userJBMap map[string]string
	// key -> json tag
	// value -> bind tag and options   详情查看
	userValidatorMap map[string]string
)

func init() {
	userScopeMap[6] = []string{"finderuser"}
	userJBMap["finder"] = "dsafsda"
	userValidatorMap["finder"] = "required,email"
}

func UserUpdateDTO(values map[string]interface{}, scope int) (updater interface{}, valiErr []models.ValidateErr, err error) {
	if values == nil || len(values) == 0 {
		return nil, nil, errors.New("value nil")
	}

	up := bson.M{}
	for jTag, value := range values {
		// 检查值是否存在scope中
		if checkUserValueOptInScope(jTag, scope) {
			// 值校验
			err := bsVali.Var(value, userValidatorMap[jTag])
			if err != nil {
				valiErr = append(valiErr, models.ValidateErr{FieldName: jTag, ValidatorMsg: err.Error()})
			} else {
				up[userJBMap[jTag]] = value
			}
		}
	}

	if valiErr != nil || len(valiErr) > 0 {
		return nil, valiErr, errors.New("validate error")
	}

	return bson.M{"$set": up}, nil, nil
}

// 判断值是否在操作权限内
func checkUserValueOptInScope(valueKey string, scope int) bool {
	for k, v := range userScopeMap {
		if k < scope {
			if arrays.ContainsString(v, valueKey) != -1 {
				return true
			}
		}
	}
	return false
}
```

出来可以控制 models(默认-o 参数)目录外，还可以控制某个文件的代码生成。只不过是将参数改为-f。
