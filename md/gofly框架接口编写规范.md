# Gofly框架开发规范文档

## 目录
1. [数据类型](#1-数据类型)
   - [基础类型](#11-基础类型)
   - [Map类型](#12-map类型)
   - [List类型](#13-list类型)
   - [Slice类型](#14-slice类型)
2. [类型转换](#2-类型转换)
   - [基础类型转换](#21-基础类型转换)
   - [ORM数据转换](#22-orm数据转换)
3. [接口开发](#3-接口开发)
   - [接口注册](#31-接口注册)
   - [参数处理](#32-参数处理)
   - [响应处理](#33-响应处理)
   - [最佳实践](#34-最佳实践)
4. [接口安全](#4-接口安全)
   - [接口加密验证](#41-接口加密验证)
   - [Token验证](#42-token验证)
   - [权限验证](#43-权限验证)
   - [跨域处理](#44-跨域处理)
5. [Token详细使用](#5-token详细使用)
   - [Token生成](#51-token生成)
   - [Token解析](#52-token解析)
   - [Token刷新](#53-token刷新)
   - [Token传递](#54-token传递)
   - [登录退出](#55-登录退出)
6. [GF工具函数](#6-gf工具函数)
   - [数组操作](#61-数组操作)
   - [IP处理](#62-ip处理)
   - [树形数据](#63-树形数据)
   - [文件操作](#64-文件操作)
   - [数据库操作](#65-数据库操作)
   - [时间处理](#66-时间处理)
   - [字符串处理](#67-字符串处理)

## 1. 数据类型

### 1.1 基础类型

框架在Go基础数据类型（如`int`、`float`、`string`、`bool`、`interface{}`）基础上，封装了以下常用类型，调用方式：`gf.类型`

| 类型      | 说明            |
| --------- | --------------- |
| gf.GinCtx | Gin上下文       |
| gf.GinObj | Gin返回数据类型 |
| gf.Var    | 通用的变量接口  |
| gf.Ctx    | 上下文类型      |
| gf.Meta   | 元数据          |

### 1.2 Map类型

Map类型提供了多种键值对组合，使用示例：`gf.Map{"name":"张三","age":25}`

| 类型       | 对应Go类型                  | 说明                                    |
| ---------- | --------------------------- | --------------------------------------- |
| Map        | map[string]interface{}      | 通用Map类型                             |
| MapAnyAny  | map[interface{}]interface{} | 键值均为interface{}                     |
| MapAnyStr  | map[interface{}]string      | 键为interface{}，值为string             |
| MapAnyInt  | map[interface{}]int         | 键为interface{}，值为int                |
| MapStrStr  | map[string]string           | 键值均为string                          |
| MapStrInt  | map[string]int              | 键为string，值为int                     |
| MapIntAny  | map[int]interface{}         | 键为int，值为interface{}                |
| MapIntStr  | map[int]string              | 键为int，值为string                     |
| MapIntInt  | map[int]int                 | 键值均为int                             |
| MapAnyBool | map[interface{}]bool        | 键为interface{}，值为bool               |
| MapStrBool | map[string]bool             | 键为string，值为bool                    |
| MapIntBool | map[int]bool                | 键为int，值为bool                       |

### 1.3 List类型

List类型是Map类型的切片集合

| 类型        | 对应Go类型                    | 说明                                    |
| ----------- | ---------------------------- | --------------------------------------- |
| List        | []map[string]interface{}      | 通用List类型                            |
| ListAnyAny  | []map[interface{}]interface{} | 键值均为interface{}的列表                |
| ListAnyStr  | []map[interface{}]string      | 键为interface{}，值为string的列表        |
| ListAnyInt  | []map[interface{}]int         | 键为interface{}，值为int的列表           |
| ListStrStr  | []map[string]string           | 键值均为string的列表                     |
| ListStrInt  | []map[string]int              | 键为string，值为int的列表                |
| ListIntAny  | []map[int]interface{}         | 键为int，值为interface{}的列表           |
| ListIntStr  | []map[int]string              | 键为int，值为string的列表                |
| ListIntInt  | []map[int]int                 | 键值均为int的列表                        |
| ListAnyBool | []map[interface{}]bool        | 键为interface{}，值为bool的列表          |
| ListStrBool | []map[string]bool             | 键为string，值为bool的列表               |
| ListIntBool | []map[int]bool                | 键为int，值为bool的列表                  |

### 1.4 Slice类型

Slice是动态数组，长度可变，使用示例：`gf.SliceInt{1,2,3}`

| 类型     | 对应Go类型     | 说明                  |
| -------- | -------------- | --------------------- |
| Slice    | []interface{}  | 通用切片类型          |
| SliceStr | []string       | 字符串切片            |
| SliceInt | []int          | 整数切片              |

## 2. 类型转换

### 2.1 基础类型转换

框架提供的类型转换方法（any为interface{}类型）：

| 方法              | 说明                                    |
| ----------------- | --------------------------------------- |
| gf.Int()          | 转换为int                               |
| gf.Int8()         | 转换为int8                              |
| gf.Int16()        | 转换为int16                             |
| gf.Int32()        | 转换为int32                             |
| gf.Int64()        | 转换为int64                             |
| gf.Uint()         | 转换为uint                              |
| gf.Uint8()        | 转换为uint8                             |
| gf.Uint16()       | 转换为uint16                            |
| gf.Uint32()       | 转换为uint32                            |
| gf.Uint64()       | 转换为uint64                            |
| gf.Float32()      | 转换为float32                           |
| gf.Float64()      | 转换为float64                           |
| gf.Bool()         | 转换为bool                              |
| gf.String()       | 转换为string                            |
| gf.Bytes()        | 转换为[]byte                            |
| gf.Strings()      | 转换为[]string                          |
| gf.Ints()         | 转换为[]int                             |
| gf.Floats()       | 转换为[]float64                         |
| gf.Interfaces()   | 转换为[]interface{}                     |
| gf.StringToJSON() | 字符串转JSON（返回interface{}）         |
| gf.JSONToString() | JSON转字符串                            |

### 2.2 ORM数据转换

ORM查询结果的字段类型转换方法：

```go
// 示例：从查询结果中获取并转换数据类型
data, err := gf.Model("createcode_product").Where("id", id).Find()
price := data["price"].Int()  // 获取int类型的价格
```

| 方法                | 说明                |
| ------------------- | ------------------- |
| Int()               | 转换为int           |
| Int8()              | 转换为int8          |
| Int16()             | 转换为int16         |
| Int32()             | 转换为int32         |
| Int64()             | 转换为int64         |
| Uint()              | 转换为uint          |
| Uint8()             | 转换为uint8         |
| Uint16()            | 转换为uint16        |
| Uint32()            | 转换为uint32        |
| Uint64()            | 转换为uint64        |
| Float32()           | 转换为float32       |
| Float64()           | 转换为float64       |
| Time("Y-m-d H:i:s") | 转换为time.Time     |
| Duration()          | 转换为time.Duration |
| Val()               | 转换为通用val       |
| Bytes()             | 转换为[]byte        |
| String()            | 转换为string        |
| Bool()              | 转换为bool          |

> 注意：如需修改查询结果中的值，请使用`gf.VarNew()`方法：
> ```go
> data["price"] = gf.VarNew(12)  // 正确的赋值方式
> ```

## 3. 接口开发

### 3.1 接口注册

使用`gf.Register`注册接口模块：

```go
type Monitor struct {}

func init() {
    fpath := Monitor{}
    gf.Register(&fpath, fpath)
    fmt.Println("Monitor模块已注册，路由: /business/dashboard/monitor/*")
}
```

### 3.2 参数处理

使用`gf.RequestParam`获取请求参数：

```go
// 简单用法
param, _ := gf.RequestParam(c)

// 带错误处理
param, err := gf.RequestParam(c)
if err != nil {
    gf.Failed().SetMsg("参数解析失败").SetData(err).Regin(c)
    return
}
```

### 3.3 响应处理

使用`gf.Success`和`gf.Failed`返回响应：

```go
// 成功响应
gf.Success().SetMsg("操作成功").SetData(data).Regin(c)

// 失败响应
gf.Failed().SetMsg("操作失败").SetData(err).Regin(c)
```

### 3.4 最佳实践

完整的接口示例：

```go
package article

import "gofly/utils/gf"

type Cate struct{}

func init() {
    fpath := Cate{}
    gf.Register(&fpath, fpath)
}

func (api *Cate) Save(c *gf.GinCtx) {
    // 1. 获取参数
    param, err := gf.RequestParam(c)
    if err != nil {
        gf.Failed().SetMsg("参数解析失败").SetData(err).Regin(c)
        return
    }

    // 2. 业务处理
    data, err := gf.Model("gf_article_cate").Save(param)
    
    // 3. 返回结果
    if err != nil {
        gf.Failed().SetMsg("保存失败").SetData(err).Regin(c)
    } else {
        gf.Success().SetMsg("保存成功").SetData(data).Regin(c)
    }
}
```

## 4. 接口安全

### 4.1 接口加密验证

框架采用三层验证机制保障接口安全：
- API接口对称加密验证
- 账号登录Token验证
- 管理后台RBAC权限验证

API接口请求需要在请求头中携带合法性参数，用于验证请求是否来自授权应用，防止接口被第三方盗用。

### 4.2 Token验证

Token验证主要用于用户身份认证，框架使用JWT机制。基本流程：
1. 用户登录后获取Token
2. 请求接口时在Header中携带Token
3. 服务端验证Token有效性

### 4.3 权限验证

后台采用RBAC权限管理模型，可以细粒度控制用户权限。

### 4.4 跨域处理

由于前后端分离架构，需要在后台配置允许访问的域名，避免跨域请求被拦截。

## 5. Token详细使用

### 5.1 Token生成

框架提供了两种生成Token的方式：

```go
// 方式一：使用UserClaims
token, err := routeuse.GenerateToken(&routeuse.UserClaims{
    ID:         res["id"].Int64(),
    Accountid:  res["accountID"].Int64(),
    BusinessID: res["businessID"].Int64(),
})

// 方式二：使用Map（2.2.6版本后）
token, err := gf.CreateToken(gf.Map{
    "ID":          data["id"].Int64(),
    "account_id":  data["account_id"].Int64(),
    "business_id": data["business_id"].Int64(),
})
```

Token中可存储的用户信息：
```go
ID         int64  `json:"id"`         // 用户数据id
Accountid  int64  `json:"accountid"`  // A端主账号id
BusinessID int64  `json:"businessID"` // B端主账号id
```

### 5.2 Token解析

根据接口是否需要Token验证，有不同的解析方式：

1. 需要Token验证的接口：
```go
// 推荐方式
getuser, _ := c.Get("user")
user := getuser.(gf.UserObj)
userID := user.ID
accountID := user.AccountID
businessID := user.BusinessID

// 或直接获取ID
userID := c.GetInt64("userID")
accountID := c.GetInt64("accountID")
businessID := c.GetInt64("businessID")

// 带错误处理
user, err := gf.GetUserInfoErr(c)
if err != nil {
    gf.Failed().SetMsg(err.Error()).Regin(c)
    return
}
```

2. 不需要Token验证的接口：
```go
// 解析并验证Token
user, err := gf.ParseTokenGetInfo(c)
if err != nil {
    gf.Failed().SetMsg("登录失效请重新登录！").SetData(err).Regin(c)
}

// 仅解析Token不验证
data, err := gf.ParseTokenNoValid(c)
```

### 5.3 Token刷新

框架会在响应成功时自动刷新Token，也可以手动刷新：

```go
token := c.Request.Header.Get("Authorization")
newtoken, err := routeuse.Refresh(token)
if err != nil {
    gf.Failed().SetMsg(err.Error()).Regin(c)
} else {
    gf.Success().SetMsg("刷新token").SetData(newtoken).Regin(c)
}
```

### 5.4 Token传递

前端请求时需要在请求头中携带Token：

```javascript
headers = {
    Authorization: "token字符串"
}
```

### 5.5 登录退出

退出登录示例：

```go
func (api *Index) Logout(c *gf.GinCtx) {
    user, err := gf.ParseTokenNoValid(c)
    if err == nil {
        // 更新用户登录状态
        gf.Model("business_account").Where("id", user.ID).Data(gf.Map{
            "loginstatus": 0
        }).Update()
    }
    gf.Success().SetMsg("退出登录").SetData(true).Regin(c)
}
```

接口忽略Token验证配置：

```go
type Index struct {
    NoNeedLogin []string // 忽略登录接口配置，忽略全部传[*]
}

func init() {
    fpath := Index{NoNeedLogin: []string{"login", "logout"}}
    gf.Register(&fpath, fpath)
}
```

> 注意：
> 1. Token验证失败会直接返回错误，验证通过才会继续执行
> 2. 框架不强制使用Redis，退出登录时建议前端清除Token
> 3. Token包含的用户信息应根据业务需求合理设置

## 6. GF工具函数

框架提供了丰富的工具函数，通过 `gf.*` 方法调用。使用前需要导入：

```go
import "gofly/utils/gf"
```

### 6.1 数组操作

#### 数组查找
```go
// 判断元素是否存在数组中
bool := gf.IsContain(array []interface{}, str string)

// 判断元素是否存在数组中-泛型
bool := gf.IsContainVal(array []gvar.Var, str gvar.Var)

// 判断字符串是否在字符串数组中
bool := gf.IsContainStr(array []string, str string)

// 判断整数是否在Var数组中
bool := gf.IntInVarArray(target int, arr []*gvar.Var)

// 判断整数是否在interface数组中
bool := gf.IntInInterfaceArray(target int, arr []interface{})
```

#### 数组合并
```go
// 合并两个数组
result := gf.MergeArr(a []*gvar.Var, b []interface{})

// 合并多个map数组
list := gf.ArrayMerge_x(a, b, c, d ...Map)

// 合并两个interface数组
result := gf.MergeArr_interface(a, b []interface{})
```

#### 数组转换
```go
// 逗号分隔字符串转数组
arr := gf.Axplode(data string)

// 数组去重
unique := gf.UniqueArr(datas []interface{})

// 将带逗号的数组字符串转为数组
result := gf.ArraymoreMerge(data []*gvar.Var)
```

### 6.2 IP处理

```go
// 获取客户端IP
ip := gf.GetIp(c *gin.Context)

// 获取服务器IP
ip := gf.LocalIP()
```

### 6.3 树形数据

```go
// 获取树状数组
tree := gf.GetTreeArray(list gform.Result, pid int64, itemprefix string)

// 树形数组转二维数组
list := gf.GetTreeToList(list []Map, field string)

// 二维数组转children树形结构
tree := gf.GetRuleTreeArrayByPack(list gform.Result, pid interface{})

// 获取子菜单包含的父级ID
pids := gf.GetRulesID(tablename, field string, menus interface{})

// 获取所有子节点ID
ids := gf.GetAllChilIds(tablename string, pids interface{})
```

### 6.4 文件操作

```go
// 批量删除文件
gf.Del_file(file_list []*gvar.Var)

// 删除单个文件
gf.DelOneFile(file_path string)
```

### 6.5 数据库操作

```go
// 判断表是否存在字段
exists := gf.DbHaseField(tablename, fields string)

// 获取表字段值
val := gf.GetTalbeFieldVal(tablename, field string, id interface{})

// 获取字典字段值
val := gf.GetDicFieldVal(group_id, val interface{})
```

### 6.6 时间处理

```go
// 日期时间转时间戳
// timetype: datetime=日期时间 datesecond=日期时间秒 date=日期
timestamp := gf.StringTimestamp(timeLayout, timetype string)

// 时间戳转日期字符串
datetime := gf.TimestampString(timedata interface{}, timetype string)
```

### 6.7 字符串处理

```go
// JSON与字符串转换
str := gf.JSONToString(data interface{})
json := gf.StringToJSON(val interface{})

// 字符串包含检查
contains := gf.StrContains(str, filed string)

// 字符串分割与合并
arr := gf.SplitAndStr(str, step string)
str := gf.ArrayToStr(data interface{}, step string)

// 敏感信息隐藏
// strtype: email、mobile
hidden := gf.HideStrInfo(strtype, val string)
```

> 注意：
> 1. 所有工具函数都在 `gofly/utils/gf` 包中
> 2. 部分函数可能需要特定的数据类型转换
> 3. 使用前请确保了解函数的参数要求和返回值类型


## gf对象管理

框架封装了一些常用的数据类型以及对象获取方法，通过​​gf.*​​​方法获取。​

使用方式：
```go
import "gofly/utils/gf"
```

### 数据类型
```go
type (
    GinCtx = gin.Context     // Gin 
    Var    = gvar.Var        // Var is a universal variable interface, like generics.
    Ctx    = context.Context // Ctx is alias of frequently-used type context.Context.
    Meta   = gmeta.Meta      // Meta is alias of frequently-used type gmeta.Meta.
)
​
type (
    Map        = map[string]interface{}      // Map is alias of frequently-used map type map[string]interface{}.
    MapAnyStr  = map[interface{}]string      // MapAnyStr is alias of frequently-used map type map[interface{}]string.
    MapAnyInt  = map[interface{}]int         // MapAnyInt is alias of frequently-used map type map[interface{}]int.
    MapStrAny  = map[string]interface{}      // MapStrAny is alias of frequently-used map type map[string]interface{}.
    MapStrStr  = map[string]string           // MapStrStr is alias of frequently-used map type map[string]string.
    MapStrInt  = map[string]int              // MapStrInt is alias of frequently-used map type map[string]int.
    MapIntAny  = map[int]interface{}         // MapIntAny is alias of frequently-used map type map[int]interface{}.
    MapIntStr  = map[int]string              // MapIntStr is alias of frequently-used map type map[int]string.
    MapIntInt  = map[int]int                 // MapIntInt is alias of frequently-used map type map[int]int.
    MapAnyBool = map[interface{}]bool        // MapAnyBool is alias of frequently-used map type map[interface{}]bool.
    MapStrBool = map[string]bool             // MapStrBool is alias of frequently-used map type map[string]bool.
    MapIntBool = map[int]bool                // MapIntBool is alias of frequently-used map type map[int]bool.
)
​
type (
    List        = []Map        // List is alias of frequently-used slice type []Map.
    ListAnyStr  = []MapAnyStr  // ListAnyStr is alias of frequently-used slice type []MapAnyStr.
    ListAnyInt  = []MapAnyInt  // ListAnyInt is alias of frequently-used slice type []MapAnyInt.
    ListStrAny  = []MapStrAny  // ListStrAny is alias of frequently-used slice type []MapStrAny.
    ListStrStr  = []MapStrStr  // ListStrStr is alias of frequently-used slice type []MapStrStr.
    ListStrInt  = []MapStrInt  // ListStrInt is alias of frequently-used slice type []MapStrInt.
    ListIntAny  = []MapIntAny  // ListIntAny is alias of frequently-used slice type []MapIntAny.
    ListIntStr  = []MapIntStr  // ListIntStr is alias of frequently-used slice type []MapIntStr.
    ListIntInt  = []MapIntInt  // ListIntInt is alias of frequently-used slice type []MapIntInt.
    ListAnyBool = []MapAnyBool // ListAnyBool is alias of frequently-used slice type []MapAnyBool.
    ListStrBool = []MapStrBool // ListStrBool is alias of frequently-used slice type []MapStrBool.
    ListIntBool = []MapIntBool // ListIntBool is alias of frequently-used slice type []MapIntBool.
)

type (
    Slice    = []interface{} // Slice is alias of frequently-used slice type []interface{}.
    SliceStr = []string      // SliceStr is alias of frequently-used slice type []string.
    SliceInt = []int         // SliceInt is alias of frequently-used slice type []int.
)
```

### 函数对象
#### IsContain
说明：判断元素是否存在数组中
参数：array是[]interface{}，str是string
返回：布尔类
格式：​
```go   
bool:=gf.IsContain(array,str)
```

#### IsContainVal
说明：判断元素是否存在数组中-泛型val
参数：array是• array是[]gvar.Var ，str是gvar.Var
返回：布尔类
格式：
```go   
bool:=gf.IsContainVal(array,str)
```

#### IsContainStr
说明：判断元素是否存在数组中-字符串类型
参数：array是[]string ，str是string 
返回：布尔类
格式：
```go   
bool:=gf.IsContainStr(array,str)
```

#### IntInVarArray
说明：判断元素是否存在数组中-int类型
参数： target是int，arr是[]*gvar.Var
返回：布尔类
格式：
```go   
data_arr:=gf.IntInVarArray(target, arr)
```

#### IntInInterfaceArray
说明：Int类型是否存在interface数组中
参数： target是int，arr是[]interface{}
返回：布尔类
格式：
```go   
bool:=gf.IntInInterfaceArray(target, arr)
```

#### GetIp
说明：获取客户访问ip函数，​​可以使用"ip地址解析属地区"插件解析IP属地​​
参数： c是gin上下文
返回：字符串
格式：
```go   
ip:=gf.GetIp("39.128.22.15")
```

#### LocalIP
说明：获取后端服务ip
返回：字符串 
格式：
```go   
ip:=gf.LocalIP()
```

#### GetAllChilIds
说明：批量获取子节点id
​参数： tablename是数据表名称，pids是父级id，即数据的pid
返回：子集数据id
格式：
```go   
ids:=gf.GetAllChilIds(tablename,pids)
```

#### GetAllChilId
说明：获取所有子级ID
​参数： tablename是数据表名称，pid是父级id，即数据的pid
返回：子集数据id，类型是[]interface{}
格式：
```go   
ids:=gf.GetAllChilId(tablename,pid)
```

#### MergeArr
说明：合并数组-两个数组合并为一个数组
参数： a是[]*gvar.Var，b是[]interface{}
返回：[]interface{}​
格式：
```go   
:=gf.MergeArr(a,b)
```

#### ArrayMerge
说明：把逗号拼接的字符串数据组，打散成数组-后台权限数据
参数： data是[]*gvar.Var
返回：[]interface{}
格式：
```go   
data_arr:=gf.ArrayMerge(data)
```

#### Axplode
说明：把逗号拼接的字符串，打散成数组
参数： data是string
返回：[]interface{}
格式：
```go   
data_arr:=gf.Axplode(data)
```

#### GetDataAuthor
说明：获取账号的数据权限-acount_id绑定的数据权限
参数： c是Gin框架请求上下文
返回：acount_ids是[]interface{}类，账号id数组；filter是bool类型，true需要权限过滤，false不需权限过滤​
格式：
```go   
acount_ids,filter:=gf.GetDataAuthor(c)
//使用
if filter { //需要权限过滤
  whereMap.Set("account_id IN(?)", account_id)
}
```

#### JSONToString
说明：转JSON编码为字符串
参数： data是interface{}​
返回：string
格式：
```go   
data_str:=gf.JSONToString(data)
```

#### StringToJSON
说明：字符串转JSON编码
参数： val是interface{}
返回：interface{}
格式：
```go   
data_str:=gf.StringToJSON(val)
```

#### GetTreeArray
说明：获取树状数组
参数：list是gform.Result类型数据，pid是最上级id（int64）,itemprefix是新标题加的前缀（string）
返回：gform.Result
格式：
```go   
list:=gf.GetTreeArray(list,pid,itemprefix)
```

#### GetTreeToList
说明：将getTreeArray的得到树形数组转为二维数组
参数： list是[]Map，field是标题字段​
返回：[]Map
格式：
```go   
list:=gf.GetTreeToList(list,field)
```

#### ArrayMerge_x
说明：多个map数组数据拼接
参数： 可添加多个map 不限个数，如：gf.ArrayMerge_x(a,b,c,d)
返回：[]Map
格式：
1
list:=gf.ArrayMerge_x(a,b...)
```

#### GetRuleTreeArrayByPack
说明：通父级pid把二维数据组转children的树形数组。
参数：list是gform.Result，pid是父级id
返回：gform.Result
格式：
```go   
tree:=gf.GetRuleTreeArrayByPack(list,pid)
```

#### GetRulesID
说明：获取子菜单包含的父级ID
参数：tablename是表名称（string），field是父级字段如：pid，menus子菜单id数组（interface{}）
返回：interface{}​
格式：
```go   
pids:=gf.GetRulesID(tablename,field,menus)
```

```

#### getParentID
说明：获取单条数据所有父级ID
参数：tablename是表名称（string），field是父级字段如：pid，id子菜单id（interface{}）
返回：interface{}
格式：
1
pids:=gf.getParentID(tablename,field,menus)
```

#### UniqueArr
说明：数组数据去重
参数：datas是[]interface{}
返回：[]interface{}
格式：
```go   
pids:=gf.UniqueArr(datas)
```

#### MergeArr_interface
说明：合并两个数组数据
参数：a、b是[]interface{}
返回：[]interface{}
格式：
```go   
pids:=gf.MergeArr_interface(a,b)
```

#### ArraymoreMerge
说明：将带有逗号的数组中字符串差分合并为数组
参数：data是[]*gvar.Var
返回：[]interface{}
格式：
```go   
data:=gf.ArraymoreMerge(data)
```

#### Del_file
说明：批量删除本地附件
参数：file_list是附件路径数组（[]*gvar.Var）
返回：无
格式：
```go   
gf.Del_file(file_list)
```

#### DelOneFile
说明：删除本地附件
参数：file_path是附件路径（string）
返回：无
格式：
```go   
gf.DelOneFile(file_path)
```

#### DbHaseField
说明：判断数据表是否存在某个字段
参数：tablename是表名称（string），fields是字段名（string）
返回：无
格式：
```go   
bool:=gf.DbHaseField(tablename, fields)
```

#### GetEditId
说明：获取请求参数id-用于数据保存或更新
参数：idstr是id参数（interface{}）
返回：float64​
格式：
```go   
f_id:=gf.GetEditId(idstr)
```

#### StringTimestamp
说明：日期时间转时间戳
参数：timetype时间格式类型  datetime=日期时间 datesecond=日期时间秒date=日期
返回：int64
格式：
```go   
timeStamp:=gf.StringTimestamp(timeLayout, timetype)
```

#### TimestampString 
说明： 时间戳格式化为日期字符串
参数：timedata是时间戳（interface{}），timetype时间格式类型 date=日期 datetime=日期时间 datesecond=日期时间秒
返回：string​
格式：
```go   
datetime:=gf.TimestampString(timedata, timetype)
```

#### GetTalbeFieldVal 
说明： 获取某数据表指定id值的对应的字段值
参数：tablename是表名称，field是字段，id需要传递的主键id值
返回：*gvar.Var
格式：
```go   
datetime:=gf.GetTalbeFieldVal(tablename, field, id)
```

#### GetDicFieldVal
说明： 获取某数据表指定id值的对应的字段值
参数：group_id是分组id，val是对keyvalue值
返回：*gvar.Var
格式：
```go   
data:=gf.GetDicFieldVal(group_id, val)
```

#### StrContains
说明：判断字符串是否包含指定元素
参数：str是字符串，filed是指定元素
返回：bool
格式：
```go   
booldata:=gf.StrContains(str, filed)
```

#### SplitAndStr
说明：把字符串打散为数组
参数：str是字符串，step是差分字符
返回：[]string
格式：
```go   
data_arr:=gf.SplitAndStr(str, step)
```

#### ArrayToStr
说明：把数组按指定字符拼接成字符串
参数：data是数组（interface{}类型），step是差分字符
返回：string​
格式：
```go   
data_arr:=gf.ArrayToStr(str, step)
```

#### StrInArray
说明：判断一个字符串是存在一个字符数组中
参数：target是判断的字符（string类型），str_array是字符串数组
返回：bool
格式：
```go   
bool:=gf.StrInArray(target, str_array)
```

#### CateAllChilId
说明：获取分类下全部子id
参数：tablename是表名称（string类型），cid是分类id
返回： []interface{}​  
格式：
```go   
bool:=gf.CateAllChilId(tablename, cid)
```

#### IsModelPath
说明：判断请求路由是否是该模块
参数：path是路径（string类型），model是模块名称
返回：bool
格式：
```go   
bool:=gf.IsModelPath(path, model)
```

#### ​HideStrInfo 
说明：隐藏手机号等敏感信息用*替换展示
参数：strtype是类型（email、mobile），vall是内容值
返回：string
格式：  
```go   
bool:=gf.HideStrInfo(strtype, vall)
```
