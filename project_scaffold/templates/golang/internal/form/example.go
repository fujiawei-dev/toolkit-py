{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import "database/sql"

// ExampleCreate 表单验证示例
// https://github.com/go-playground/validator
type ExampleCreate struct {
	// 限制字符串长度
	ShortStringField string `binding:"required,max=16" example:"短字符串字段" json:"short_string_field"`
	LongStringField  string `binding:"required,min=16,max=512" example:"长字符串字段" json:"long_string_field"`
	LongTextField    string `binding:"required,min=512" example:"超长文本字段" json:"long_text_field"`
	CommentField     string `binding:"required,len=5" example:"定长字符串" json:"comment_field"`

	// 数值比较
	// 大于 & 小于
	IndexField int `binding:"gt=5,lt=128" example:"100" json:"index_field"`
	// 大于等于 & 小于等于
	UniqueField int `binding:"gte=5,lte=100" json:"unique_field"`
	// 不等于
	UniqueIndexField int `binding:"ne=1" json:"unique_index_field"`

	// 嵌套结构验证
	// 表示列表不为空，且元素的长度不超过 20
	StringList []string `binding:"required,gt=0,dive,max=20" json:"string_list"`
	// 表示字典不为空，且键的长度不超过 10，值的长度不超过 20
	StringDict map[string]string `binding:"required,gt=0,dive,keys,max=10,endkeys,required,max=20" json:"string_dict"`

	// 自定义枚举值
	IntegerField int `binding:"oneof=20 50 100" example:"50" json:"integer_field"`

	// 字段的默认值
	DefaultField string `example:"默认字段" json:"default_field,default=ignore"`

	// 字段间关系
	UnsignedIntegerField uint    `json:"unsigned_integer_field"`
	Float64Field         float64 `json:"float_64_field"`
	// 在 UnsignedIntegerField 或者 Float32Field 存在时，必须存在
	Float32Field float32 `binding:"required_with=UnsignedIntegerField Float32Field" json:"float_32_field"`
	// 在 UnsignedIntegerField 或者 Float32Field 不存在时，必须存在
	BinaryField []byte `binding:"required_without_all=UnsignedIntegerField Float32Field" json:"binary_field"`
	// 必须不等于 UniqueField 的值
	CheckField string `binding:"required,nefield=UniqueField" json:"check_field"`

	ExampleUpdate
}

type ExampleUpdate struct {
	Ignored int `swaggerignore:"true"` // 排除字段

	NotNullField sql.NullBool `binding:"required" swaggertype:"boolean" example:"1" json:"not_null_field"` // 禁止空值字段

	// 必填项/条件可填
	// omitempty 表示变量可以不填，但是填的时候必须满足条件
	AllowReadAndCreate   string `binding:"required,ip4_addr" json:"allow_read_and_create"`
	AllowReadAndUpdate   string `binding:"omitempty,url" json:"allow_read_and_update"`
	AllowCreateAndUpdate string `binding:"omitempty,email" json:"allow_create_and_update"`
	ReadOnly             string `binding:"required,alphanum" json:"read_only"`
	CreateOnly           string `binding:"required,uuid4" json:"create_only"`
	IgnoreWriteAndRead   string `binding:"omitempty,startswith=go" json:"ignore_write_and_read"`
	IgnoreMigration      string `binding:"omitempty,lowercase" json:"ignore_migration"`
}
