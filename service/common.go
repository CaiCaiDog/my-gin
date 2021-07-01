package service

import (
	"database/sql"
	"fmt"
	"reflect"
)

// 获取查询结果(单行)
func GetRow(row *sql.Rows) map[string]string {
	// 字段名组成的切片
	col, _ := row.Columns()
	// 存放数据
	vals := make([][]byte, len(col))
	// 做scan的参数，将vals的值保存在scans中
	scans := make([]interface{}, len(col))
	for k := range col {
		scans[k] = &vals[k]
	}
	result := make(map[string]string)
	for row.Next() {
		// 将查询的结果保存到scans中，也是保存在vals中(切片展开)
		row.Scan(scans...)
		// 将单行的byte数据循环转换成string类型
		for k, v := range vals {
			key := col[k]
			result[key] = string(v)
		}
	}
	return result
}

// 获取查询结果（多行）
func GetAllRows(rows *sql.Rows) map[int]map[string]string {
	// 字段名组成的切片
	columns, _ := rows.Columns()
	// 存放取出来的数据结果，表示一行所有列的值，长度表示行数
	vals := make([][]byte, len(columns))
	// 做rows.scan的参数，将扫描后的数据存储在scan中
	scans := make([]interface{}, len(columns))
	// 将每一行填充到vals中
	for k, _ := range vals {
		scans[k] = &vals[k]
	}
	result := make(map[int]map[string]string)
	i := 0
	for rows.Next() {
		// 查询结果存放到scan[]中，也是存放到vals变量中(切片展开多参数传入)
		rows.Scan(scans...)
		// 循环取出转换成字符串
		row := make(map[string]string)
		for k, v := range vals {
			key := columns[k]
			row[key] = string(v)
		}
		result[i] = row
		i ++
	}
	return result
}

func GetMethod(class interface{}) {
	value := reflect.ValueOf(class)
    typ := value.Type()
    for i := 0; i < value.NumMethod(); i++ {
        fmt.Println(fmt.Sprintf("method[%d]%s and type is %v", i, typ.Method(i).Name, typ.Method(i).Type))
    }
}