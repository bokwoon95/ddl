package ddl

type T struct {
}

type TField struct {
	// NOTE: if all TField does is write back to some pointer field, then it could function just a value receiver just like T does
}

func (t T) Field(field Field) TField {
	return TField{}
}

func (f TField) Ignore() {
}

func (f TField) Type(typ string) TField {
	return f
}

func (f TField) Generated(expr string, fields ...Field) TField {
	return f
}

func (f TField) Stored() TField {
	return f
}

func (f TField) Default(expr string, fields ...Field) TField {
	return f
}

func (f TField) Autoincrement() TField {
	return f
}

func (f TField) Identity() TField {
	return f
}

func (f TField) AlwaysIdentity() TField {
	return f
}

func (f TField) OnUpdateCurrentTimestamp() TField {
	return f
}

func (f TField) NotNull() TField {
	return f
}

func (f TField) PrimaryKey() TField {
	return f
}

func (f TField) Unique() TField {
	return f
}

func (f TField) Collate(collation string) TField {
	return f
}

func (t T) Check(name string, expr string, fields ...Field) {
}

func (t T) Unique(fields ...Field) {
}

func (t T) PrimaryKey(fields ...Field) {
}

func (t T) NameUnique(name string, fields ...Field) {
}

func (t T) NamePrimaryKey(name string, fields ...Field) {
}

type TIndex struct {
}

func (t T) Index(fields ...Field) TIndex {
	return TIndex{}
}

func (t T) NameIndex(name string) TIndex {
	return TIndex{}
}

func (i TIndex) Fields(fields ...Field) TIndex {
	return i
}

func (i TIndex) Expr(expr string, fields ...Field) TIndex {
	return i
}

func (i TIndex) Unique() TIndex {
	return i
}

func (i TIndex) Schema(schema string) TIndex {
	return i
}

func (i TIndex) Using(method string) TIndex {
	return i
}

func (i TIndex) Where(expr string, fields ...Field) TIndex {
	return i
}

func (i TIndex) Include(fields ...Field) TIndex {
	return i
}
