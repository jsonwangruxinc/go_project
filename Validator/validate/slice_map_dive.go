package validate

type SliceStruct struct {
	OpCode int    `v:"eq=1|eq=2"`
	Op     string `v:"required"`
}

func SliceValidate() {
	//单组切片
	v := validate
	slice1 := []string{"12345", "67890", "1234210129301"}
	var err error
	err = v.Var(slice1, "gte=3,dive,required,gte=5,lte=10,number")
	outRes("slice", &err)
	//	多组切片
	slice2 := [][]string{
		{"12345", "67890", "1234210129301"},
		{"12345", "67890", "1234210129301"},
		{"123", "678"},
	}
	err = v.Var(slice2, "gte=3,dive,gte=3,dive,required,gte=5,lte=10,number")
	outRes("slice2", &err)

	slice3 := []*SliceStruct{
		{
			OpCode: 1,
			Op:     "切片操作",
		},
		{
			OpCode: 2,
			Op:     "切片操作",
		},
		{
			OpCode: 3,
			Op:     "切片操作",
		},
	}
	err = v.Var(slice3, "gte=2,dive")
	outRes("slice3", &err)
}

func MapValidate() {
	v := validate
	var err error
	mp1 := map[string]string{
		"Aa": "12345",
		"B":  "1234567899010",
		"C":  "12345",
	}
	err = v.Var(mp1, "gte=3,dive,keys,len=1,alpha,endkeys,required,gte=5,lte=10,number")
	outRes("mp1", &err)

	mp2 := map[string]map[string]string{
		"A": {
			"A": "12345",
			"B": "12345",
			"C": "12345",
		},
		"B": {
			"A": "12345",
			"B": "12345",
			"C": "12345",
		},
	}
	err = v.Var(mp2, "gte=2,dive,keys,len=1,alpha,endkeys,required,gte=3,dive,keys,len=1,alpha,endkeys,required,gte=5,lte=10,number")
	outRes("mp2", &err)

}
