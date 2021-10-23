## Unmarshal JSON


### 1. json字段少于结构体字段

反序列化JSON的时候，如果结构体的的的字段，在json中不存在，那么会默认设置为该字段的`零值`。

比如：

```Go
type Food struct {
  Id             int     `json:"id,omitempty"`
  Name           string  `json:"name"`
  FatPerServ     float64 `json:"fat_per_serv"`
  ProteinPerServ float64 `json:"protein_per_serv"`
  CarbPerServ    float64 `json:"carb_per_serv"`
}

func main() {
  bb := []byte(`
    {
    "name": "Broccoli",
    "fat_per_serv": 0.3,
    "protein_per_serv": 2.5,
    "carb_per_serv": 3.5
  }`)
  var fD Food
  if err := json.Unmarshal(bb, &fD); err != nil {
    panic(err)
  }
  fmt.Println("unmarshaled Food:", fD)
}
```

结构体里的`Id`字段在json中不存在，因为是int值，则默认设置为0.结果如下：

```bash
unmarshaled Food: {0 Broccoli 0.3 2.5 3.5}
```

### 2. JSON字段多于结构体字段

在反序的d过程中，如果json中的字段在Go结构体中找不到对应的字段，则忽略该字段。

```Go

type Food2 struct {
  Id             int     `json:"id,omitempty"`
  Name           string  `json:"name"`
  FatPerServ     float64 `json:"fat_per_serv"`
}

func main() {
	bb := []byte(`
  {
	"id": 1,
    "name": "Broccoli",
    "fat_per_serv": 0.3,
    "protein_per_serv": 2.5,
    "carb_per_serv": 3.5
  }`)
	var fD Food2
	if err := json.Unmarshal(bb, &fD); err != nil {
		panic(err)
	}
	fmt.Printf("unmarshaled Food:\n %#v\n", fD)
}
```
输出:
```bash
unmarshaled Food:
main.Food{ID:1, Name:"Broccoli", FatPerServ:0.3}
```
json中的`protein_per_serv`和`carb_per_serv`都被忽略了。

### 3. 对于json的number类型，Go默认解析成float64
```go
func main() {
  s := []byte("1234")
  var inum interface{}
  if err := json.Unmarshal(s, &inum); err != nil {
    panic(err)
  }
  switch v := inum.(type) {
  case int:
    fmt.Println("it's an int:", v)
  case float64:
    fmt.Println("it's a float:", v)
  // other possible types enumerated...
  default:
    panic("can't figure out the type")
  }
}

```
结果是:
```shell
it's a float: 1234
```

### 4. 使用json.RawMessage进行延迟解析
```go
type Cluster struct {
	JobID   int    `json:"job_id"`
	Cluster string `json:"cluster"`
	Group   int    `json:"group"`
	Result  int    `json:"result"`
	Checkor string `json:"checkor"`
}

func main() {
	s := `{
    "CurrentDoubleCheckResult":{
        "hnb":{
            "job_id":50011,
            "cluster":"hnb",
            "group":0,
            "result":1,
            "checkor":"wiilliamcai_i"
        }
    },
    "DoubleCheckResultHistory":[
        {
            "job_id":50028,
            "cluster":"hnb",
            "group":0,
            "result":2,
            "checkor":"lfn"
        },
        {
            "job_id":50028,
            "cluster":"hnb",
            "group":0,
            "result":2,
            "checkor":"lfn"
        }
    ]}`

	var res map[string]json.RawMessage
	json.Unmarshal([]byte(s), &res)
	for k, v := range res {
		if k == "DoubleCheckResultHistory" {
			var clusters []Cluster
			json.Unmarshal(v, &clusters)
			fmt.Printf("k:%s, v:%#v\n", k, clusters)
		} else {
			var cluster map[string]Cluster
			json.Unmarshal(v, &cluster)
			fmt.Printf("k:%s, v:%#v\n", k, cluster)
		}
	}
}
```
由于对象里头的结构不同质，因此我们可以先用`json.RawMessage` 保留所有第一层的值仍然为`string`, 再慢慢解析.
结果:
```shell
k:CurrentDoubleCheckResult, v:map[string]main.Cluster{"hnb":main.Cluster{JobID:50011, Cluster:"hnb", Group:0, Result:1, Checkor:"wiilliamcai_i"}}
k:DoubleCheckResultHistory, v:[]main.Cluster{main.Cluster{JobID:50028, Cluster:"hnb", Group:0, Result:2, Checkor:"lfn"}, main.Cluster{JobID:50028, Cluster:"hnb", Group:0, Result:2, Checkor:"lfn"}}
```

### 5. 利用Marshaler interface定制专属的编解码方法

比如我们想把`int`编码为16进制:
```go
type Marshaler interface {
    MarshalJSON() ([]byte, error)
}

type Account struct {
    Id   int32
    Name string
}

func (a Account) MarshalJSON() ([]byte, error) {
    m := map[string]string{
        "id":   fmt.Sprintf("0x%08x", a.Id),
        "name": a.Name,
    }
    return json.Marshal(m)
}

func main() {
    joe := Account{Id: 123, Name: "Joe"}
    fmt.Println(joe)

    s, _ := json.Marshal(joe)
    fmt.Println(string(s))
}

```
结果:
```shell
{"id":"0x0000007b","name":"Joe"}
```

### 6. 忽略空值和忽略字段
1. 利用`omitempty` tag忽略空值
2. 利用`-` 忽略字段

## 参考文献
[Go JSON Cookbook](https://eli.thegreenplace.net/2019/go-json-cookbook/)
