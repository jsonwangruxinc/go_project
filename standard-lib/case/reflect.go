package _case

import (
	"errors"
	"fmt"
	"reflect"
)

func ReflectCase() {
	type user struct {
		ID    int64
		Name  string
		Hobby []string
	}
	type outuser struct {
		ID    int64
		Name  string
		Hobby []string
	}
	u := user{ID: 1, Name: "wang", Hobby: []string{"篮球", "乒乓球"}}
	out := outuser{}

	res := copy(&out, u)
	fmt.Println(res, out)

	listUser := []user{
		{ID: 1, Name: "wang", Hobby: []string{"赛马", "羽毛球"}},
		{ID: 2, Name: "xin", Hobby: []string{"保龄球", "棒球"}},
		{ID: 3, Name: "Ru", Hobby: []string{"足球", "高尔夫"}},
	}
	list := sliceColumn(listUser, "Hobby")
	fmt.Println(list)
}
func copy(dest interface{}, source interface{}) error {
	sTyep := reflect.TypeOf(source)
	sValue := reflect.ValueOf(source)
	//	如果为指针类型，则获取值
	if sTyep.Kind() == reflect.Ptr {
		sTyep = sTyep.Elem()
		sValue = sValue.Elem()
	}
	dType := reflect.TypeOf(dest)
	dValue := reflect.ValueOf(dest)
	if dType.Kind() != reflect.Ptr {
		return errors.New("目标对象必须为struct指针类型")
	}
	dType = dType.Elem()
	dValue = dValue.Elem()

	if sValue.Kind() != reflect.Struct {
		return errors.New("原对象必须为struct或struct的指针")
	}
	if dValue.Kind() != reflect.Struct {
		return errors.New("目标原对象必须为struct的指针")
	}
	destObject := reflect.New(dType)
	for i := 0; i < dType.NumField(); i++ {
		destFiled := dType.Field(i)
		if sourceFiled, ok := sTyep.FieldByName(destFiled.Name); ok {
			if destFiled.Type != sourceFiled.Type {
				continue
			}
			value := sValue.FieldByName(destFiled.Name)
			destObject.Elem().FieldByName(destFiled.Name).Set(value)
		}
	}
	dValue.Set(destObject.Elem())
	return nil
}

func sliceColumn(slice interface{}, column string) interface{} {
	t := reflect.TypeOf(slice)
	v := reflect.ValueOf(slice)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	if v.Kind() == reflect.Struct {
		val := v.FieldByName(column)
		return val.Interface()
	}
	if v.Kind() != reflect.Slice {
		return nil
	}
	t = t.Elem()
	if v.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	f, _ := t.FieldByName(column)
	sliceType := reflect.SliceOf(f.Type)
	s := reflect.MakeSlice(sliceType, 0, 0)
	for i := 0; i < v.Len(); i++ {
		o := v.Index(i)
		if o.Kind() == reflect.Struct {
			Val := o.FieldByName(column)
			s = reflect.Append(s, Val)
		}
		if o.Kind() == reflect.Ptr {
			v1 := o.Elem()
			val := v1.FieldByName(column)
			s = reflect.Append(s, val)
		}
	}
	return s.Interface()
}
