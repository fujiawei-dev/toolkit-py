{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

type ExampleCreate struct {
	// 嵌套结构验证
	Array []string          `binding:"required,gt=0,dive,max=20" json:"array,omitempty"`
	Map   map[string]string `binding:"required,gt=0,dive,keys,max=10,endkeys,required,max=2" json:"map,omitempty"`

	IndexField       int `binding:"omitempty,min=5,max=128" example:"100" json:"index_field,omitempty"`
	UniqueField      int `json:"unique_field,omitempty"`
	UniqueIndexField int `json:"unique_index_field,omitempty"`

	// 限制字符串长度
	ShortStringField string `binding:"required,min=2,max=16" example:"短字符串字段" json:"short_string_field,omitempty"`
	LongStringField  string `binding:"required,min=64,max=256" example:"长字符串字段" json:"long_string_field,omitempty"`
	LongTextField    string `binding:"required,min=512" example:"超长文本字段" json:"long_text_field,omitempty"`

	// 自定义枚举值
	IntegerField int `binding:"oneof=20 50 100" example:"50" json:"integer_field,omitempty"`

	UnsignedIntegerField uint    `json:"unsigned_integer_field,omitempty"`
	Float64Field         float64 `json:"float_64_field,omitempty"`
	Float32Field         float32 `json:"float_32_field,omitempty"`
	BinaryField          []byte  `json:"binary_field,omitempty"`

	DefaultField string `example:"默认字段" json:"default_field,omitempty"`

	CheckField   int    `json:"check_field,omitempty"`
	CommentField string `binding:"required,len=5" example:"字符串定长" json:"comment_field,omitempty"`

	AllowReadAndCreate   string `binding:"omitempty,uuid4" json:"allow_read_and_create,omitempty"`
	AllowReadAndUpdate   string `json:"allow_read_and_update,omitempty"`
	AllowCreateAndUpdate string `json:"allow_create_and_update,omitempty"`
	ReadOnly             string `json:"read_only,omitempty"`
	CreateOnly           string `json:"create_only,omitempty"`
	IgnoreWriteAndRead   string `json:"ignore_write_and_read,omitempty"`
	IgnoreMigration      string `json:"ignore_migration,omitempty"`
}

type ExampleUpdate struct {
	AllowReadAndCreate   string `binding:"omitempty,uuid4" json:"allow_read_and_create,omitempty"`
	AllowReadAndUpdate   string `json:"allow_read_and_update,omitempty"`
	AllowCreateAndUpdate string `json:"allow_create_and_update,omitempty"`
	ReadOnly             string `json:"read_only,omitempty"`
	CreateOnly           string `json:"create_only,omitempty"`
	IgnoreWriteAndRead   string `json:"ignore_write_and_read,omitempty"`
	IgnoreMigration      string `json:"ignore_migration,omitempty"`
}
