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

func main() {
	bb := []byte(`
  {
	"id": 1,
    "name": "Broccoli",
    "fat_per_serv": 0.3,
    "protein_per_serv": 2.5,
    "carb_per_serv": 3.5
  }`)
	var fD Food
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