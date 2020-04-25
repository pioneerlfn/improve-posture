package main

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)



type wantedType struct {
	RetweetNum  int   `json:"retweet_num" db:"retweet_num"`
	PublishTool string `json:"publish_tool" db:"publish_tool"`
}

func main() {
	f1 := map[string]interface{}{
		"publish_tool": []string{"亚马逊Z.cn", "天上掉的Android"},
		"user_id":      "3199224287",
	}

	db, err := sqlx.Open("mysql", "account:pw@tcp(127.0.0.1:3306)/target_db")
	if err != nil {
		panic(err)
	}
	// 注意，res要用明确的类型
	res := []wantedType{}
	// 写你自己的sqlTpl
	tpl := `SELECT retweet_num, publish_tool FROM %s WHERE `
	err = getTableRowsByField(context.Background(), db, "your_table_name", tpl, f1, &res)

	fmt.Printf("%#v\n", res)
}

// 可以作为一个通用util函数，降低代码冗余
func getTableRowsByField(ctx context.Context, db *sqlx.DB, sqlTpl, table string,
	 fields map[string]interface{}, result interface{}) error {
	rawQeury := fmt.Sprintf(sqlTpl, table) + " %s "
	var (
		fieldsValue []interface{}
		rawSql      string
	)
	if len(fields) > 0 {
		rawSql, fieldsValue = assembleQuerySQL(rawQeury, fields)
	}

	err := db.Select(result, rawSql, fieldsValue...)
	if err != nil {
		panic(err)
	}
	return nil
}

// 只处理了value是slice和非slice这两种情况
func assembleQuerySQL(sql string, fields map[string]interface{}) (string, []interface{}) {
	var (
		dest        []string
		fieldsValue []interface{}
	)
	for k, v := range fields {
		value := reflect.ValueOf(v)
		if value.Kind() == reflect.Slice {
			slen := value.Len()
			if slen > 0 {
				placeholder := []string{}
				valueholder := []interface{}{}
				for i := 0; i < slen; i++ {
					placeholder = append(placeholder, "?")
					valueholder = append(valueholder,
						value.Index(i).Interface())
				}
				dest = append(dest, fmt.Sprintf("%s IN (%s)",
					string(k), strings.Join(placeholder, ",")))
				fieldsValue = append(fieldsValue, valueholder...)
			}
		} else {
			if fv, ok := v.(string); ok {
				// string类型，直接组装
				dest = append(dest, fmt.Sprintf("%s=%s", k, fv))
			} else {
				// 其他类型，占位
				dest = append(dest, fmt.Sprintf("%s=?", k))
				fieldsValue = append(fieldsValue, v)
			}
		}
	}

	if len(dest) == 0 {
		return sql, fieldsValue
	}

	sql = fmt.Sprintf(sql, strings.Join(dest, " AND "))
	return sql, fieldsValue
}


// scanAll是sqlx中Select最后会调用的一个函数
// 在这个函数中，会通过Interface()将result的类型还原，从而将结果扫描进去

/*func scanAll(rows rowsi, dest interface{}, structOnly bool) error {
	var v, vp reflect.Value

	value := reflect.ValueOf(dest)

	// json.Unmarshal returns errors for these
	if value.Kind() != reflect.Ptr {
		return errors.New("must pass a pointer, not a value, to StructScan destination")
	}
	if value.IsNil() {
		return errors.New("nil pointer passed to StructScan destination")
	}
	direct := reflect.Indirect(value)

	slice, err := baseType(value.Type(), reflect.Slice)
	if err != nil {
		return err
	}

	isPtr := slice.Elem().Kind() == reflect.Ptr
	base := reflectx.Deref(slice.Elem())
	scannable := isScannable(base)

	if structOnly && scannable {
		return structOnlyError(base)
	}

	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	// if it's a base type make sure it only has 1 column;  if not return an error
	if scannable && len(columns) > 1 {
		return fmt.Errorf("non-struct dest type %s with >1 columns (%d)", base.Kind(), len(columns))
	}

	if !scannable {
		var values []interface{}
		var m *reflectx.Mapper

		switch rows.(type) {
		case *Rows:
			m = rows.(*Rows).Mapper
		default:
			m = mapper()
		}

		fields := m.TraversalsByName(base, columns)
		// if we are not unsafe and are missing fields, return an error
		if f, err := missingFields(fields); err != nil && !isUnsafe(rows) {
			return fmt.Errorf("missing destination name %s in %T", columns[f], dest)
		}
		values = make([]interface{}, len(columns))

		for rows.Next() {
			// create a new struct type (which returns PtrTo) and indirect it
			vp = reflect.New(base)
			v = reflect.Indirect(vp)

			err = fieldsByTraversal(v, fields, values, true)
			if err != nil {
				return err
			}

			// scan into the struct field pointers and append to our results
			err = rows.Scan(values...)
			if err != nil {
				return err
			}

			if isPtr {
				direct.Set(reflect.Append(direct, vp))
			} else {
				direct.Set(reflect.Append(direct, v))
			}
		}
	} else {
		for rows.Next() {
			vp = reflect.New(base)


			// 这里通过Interface()拿到dest的类型
			err = rows.Scan(vp.Interface())
			if err != nil {
				return err
			}
			// append
			if isPtr {
				direct.Set(reflect.Append(direct, vp))
			} else {
				direct.Set(reflect.Append(direct, reflect.Indirect(vp)))
			}
		}
	}

	return rows.Err()
}
*/