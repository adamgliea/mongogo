package bson

const (
	_ = iota
	DoubleType
	StringType
	EmDocType
	ArrayType
	BinaryType
	undefType
	MongoIdType
	BoolType
	DateTimeType
	NullType
	RegularType
	DBPointerType
	JsCodeType
	SymbolType
	JsCodeWithScopeType
	Int32Type
	TimestampType
	Int64Type
	MinKeyType = 255
	MaxKeyType = 127
)

type bsonTypeInterface interface {
	pack() string
	unpack(content string)
}

type BsonDouble float64
func (self *BsonDouble) pack() string {
	

	return ""
}
func (self *BsonDouble) unpack(content string) {
}
