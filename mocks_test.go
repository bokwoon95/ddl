package ddl

import "bytes"

type tableinfo [2]string

func (t tableinfo) GetSchema() string { return t[0] }
func (t tableinfo) GetName() string   { return t[1] }

type field string
type blobfield struct{ field }
type booleanfield struct{ field }
type jsonfield struct{ field }
type numberfield struct{ field }
type stringfield struct{ field }
type timefield struct{ field }

func (f field) GetName() string { return string(f) }
func (f field) AppendSQLExclude(dialect string, buf *bytes.Buffer, args *[]interface{}, params map[string][]int, excludedTableQualifiers []string) error {
	buf.WriteString(string(f))
	return nil
}
func (f blobfield) GetType() string    { return "blob" }
func (f booleanfield) GetType() string { return "boolean" }
func (f jsonfield) GetType() string    { return "json" }
func (f numberfield) GetType() string  { return "number" }
func (f stringfield) GetType() string  { return "string" }
func (f timefield) GetType() string    { return "time" }

type _ACTOR struct {
	tableinfo          `ddl:"name=actor"`
	ACTOR_ID           numberfield `ddl:"type=INTEGER primarykey"`
	FIRST_NAME         stringfield `ddl:"notnull"`
	LAST_NAME          stringfield `ddl:"notnull index"`
	FULL_NAME          stringfield `ddl:"generated={{first_name || ' ' || last_name} virtual}"`
	FULL_NAME_REVERSED stringfield `ddl:"generated={{last_name || ' ' || first_name} stored}"`
	LAST_UPDATE        timefield   `ddl:"default=DATETIME('now') notnull"`
}

func new_ACTOR(dialect string) _ACTOR {
	tbl := _ACTOR{tableinfo: [2]string{"", "actor"}}
	switch dialect {
	case "postgres":
		tbl.tableinfo[0] = "public"
	case "mysql":
		tbl.tableinfo[0] = "db"
	}
	return tbl
}

func (ACTOR _ACTOR) Constraints(dialect string, t T) {
	switch dialect {
	case "postgres":
		t.Field(ACTOR.ACTOR_ID).Identity()
		t.Field(ACTOR.FULL_NAME_REVERSED).Generated("last_name || ' ' || first_name").Stored()
		t.Field(ACTOR.LAST_UPDATE).Type("TIMESTAMPTZ").Default("NOW()")
	case "mysql":
		t.Field(ACTOR.ACTOR_ID).Autoincrement()
		t.Field(ACTOR.FIRST_NAME).Type("VARCHAR(45)")
		t.Field(ACTOR.LAST_NAME).Type("VARCHAR(45)")
		t.Field(ACTOR.FULL_NAME).Type("VARCHAR(45)").Generated("CONCAT(first_name, ' ', last_name)")
		t.Field(ACTOR.FULL_NAME_REVERSED).Type("VARCHAR(45)").Generated("CONCAT(last_name, ' ', first_name)").Stored()
		t.Field(ACTOR.LAST_UPDATE).Type("TIMESTAMP").Default("CURRENT_TIMESTAMP").OnUpdateCurrentTimestamp()
	}
}

type _CATEGORY struct {
	tableinfo   `ddl:"name=category"`
	CATEGORY_ID numberfield `ddl:"type=INTEGER primarykey"`
	NAME        stringfield `ddl:"notnull"`
	LAST_UPDATE timefield   `ddl:"default=DATETIME('now') notnull"`
}

func (CATEGORY _CATEGORY) Constraints(dialect string, t T) {
	switch dialect {
	case "postgres":
		t.Field(CATEGORY.CATEGORY_ID).Identity()
		t.Field(CATEGORY.LAST_UPDATE).Type("TIMESTAMPTZ").Default("NOW()")
	case "mysql":
		t.Field(CATEGORY.CATEGORY_ID).Autoincrement()
		t.Field(CATEGORY.NAME).Type("VARCHAR(25)")
		t.Field(CATEGORY.LAST_UPDATE).Type("TIMESTAMP").Default("CURRENT_TIMESTAMP").OnUpdateCurrentTimestamp()
	}
}

type _COUNTRY struct {
	tableinfo   `ddl:"name=country"`
	COUNTRY_ID  numberfield `ddl:"type=INTEGER primarykey"`
	COUNTRY     stringfield `ddl:"notnull"`
	LAST_UPDATE timefield   `ddl:"default=DATETIME('now') notnull"`
}

func (COUNTRY _COUNTRY) Constraints(dialect string, t T) {
	switch dialect {
	case "postgres":
		t.Field(COUNTRY.COUNTRY_ID).Identity()
		t.Field(COUNTRY.LAST_UPDATE).Type("TIMESTAMPTZ").Default("NOW()")
	case "mysql":
		t.Field(COUNTRY.COUNTRY_ID).Autoincrement()
		t.Field(COUNTRY.COUNTRY).Type("VARCHAR(50)")
		t.Field(COUNTRY.LAST_UPDATE).Type("TIMESTAMP").Default("CURRENT_TIMESTAMP").OnUpdateCurrentTimestamp()
	}
}

type _CITY struct {
	tableinfo   `ddl:"name=city"`
	CITY_ID     numberfield `ddl:"type=INTEGER primarykey"`
	CITY        stringfield `ddl:"notnull"`
	COUNTRY_ID  numberfield `ddl:"notnull references={country onupdate=cascade ondelete=restrict} index"`
	LAST_UPDATE timefield   `ddl:"default=DATETIME('now') notnull"`
}

func (CITY _CITY) Constraints(dialect string, t T) {
	switch dialect {
	case "postgres":
		t.Field(CITY.CITY_ID).Identity()
		t.Field(CITY.LAST_UPDATE).Type("TIMESTAMPTZ").Default("NOW()")
	case "mysql":
		t.Field(CITY.CITY_ID).Autoincrement()
		t.Field(CITY.CITY).Type("VARCHAR(50)")
		t.Field(CITY.LAST_UPDATE).Type("TIMESTAMP").Default("CURRENT_TIMESTAMP").OnUpdateCurrentTimestamp()
	}
}

type _ADDRESS struct {
	tableinfo   `ddl:"name=address"`
	ADDRESS_ID  numberfield `ddl:"type=INTEGER primarykey"`
	ADDRESS     stringfield `ddl:"notnull"`
	ADDRESS2    stringfield
	DISTRICT    stringfield `ddl:"notnull"`
	CITY_ID     numberfield `ddl:"notnull references={city onupdate=cascade ondelete=restrict} index"`
	POSTAL_CODE stringfield
	PHONE       stringfield `ddl:"notnull"`
	LAST_UPDATE timefield   `ddl:"default=DATETIME('now') notnull"`
}

func (ADDRESS _ADDRESS) Constraints(dialect string, t T) {
	switch dialect {
	case "postgres":
		t.Field(ADDRESS.ADDRESS_ID).Identity()
		t.Field(ADDRESS.LAST_UPDATE).Type("TIMESTAMPTZ").Default("NOW()")
	case "mysql":
		t.Field(ADDRESS.ADDRESS_ID).Autoincrement()
		t.Field(ADDRESS.ADDRESS).Type("VARCHAR(50)")
		t.Field(ADDRESS.ADDRESS2).Type("VARCHAR(50)")
		t.Field(ADDRESS.DISTRICT).Type("VARCHAR(20)")
		t.Field(ADDRESS.POSTAL_CODE).Type("VARCHAR(10)")
		t.Field(ADDRESS.PHONE).Type("VARCHAR(20)")
		t.Field(ADDRESS.LAST_UPDATE).Type("TIMESTAMP").Default("CURRENT_TIMESTAMP").OnUpdateCurrentTimestamp()
	}
}

type _LANGUAGE struct {
	tableinfo   `ddl:"name=language"`
	LANGUAGE_ID numberfield `ddl:"type=INTEGER primarykey"`
	NAME        stringfield `ddl:"notnull"`
	LAST_UPDATE timefield   `ddl:"default=DATETIME('now') notnull"`
}

func (LANGUAGE _LANGUAGE) Constraints(dialect string, t T) {
	switch dialect {
	case "postgres":
		t.Field(LANGUAGE.LANGUAGE_ID).Identity()
		t.Field(LANGUAGE.LAST_UPDATE).Type("TIMESTAMPTZ").Default("NOW()")
	case "mysql":
		t.Field(LANGUAGE.LANGUAGE_ID).Autoincrement()
		t.Field(LANGUAGE.NAME).Type("CHAR(20)")
		t.Field(LANGUAGE.LAST_UPDATE).Type("TIMESTAMP").Default("CURRENT_TIMESTAMP").OnUpdateCurrentTimestamp()
	}
}

type _FILM struct {
	tableinfo            `ddl:"name=film"`
	FILM_ID              numberfield `ddl:"type=INTEGER primarykey"`
	TITLE                stringfield `ddl:"notnull index"`
	DESCRIPTION          stringfield
	RELEASE_YEAR         numberfield
	LANGUAGE_ID          numberfield `ddl:"notnull references={language onupdate=cascade ondelete=restrict} index"`
	ORIGINAL_LANGUAGE_ID numberfield `ddl:"references={language onupdate=cascade ondelete=restrict} index"`
	RENTAL_DURATION      numberfield `ddl:"default=3 notnull"`
	RENTAL_RATE          numberfield `ddl:"type=DECIMAL(4,2) default=4.99 notnull"`
	LENGTH               numberfield
	REPLACEMENT_COST     numberfield `ddl:"type=DECIMAL(5,2) default=19.99 notnull"`
	RATING               stringfield `ddl:"default='G'"`
	SPECIAL_FEATURES     jsonfield
	LAST_UPDATE          timefield   `ddl:"default=DATETIME('now') notnull"`
	FULLTEXT             stringfield `ddl:"notnull"`
}

func (FILM _FILM) Constraints(dialect string, t T) {
	switch dialect {
	case "postgres":
		t.Field(FILM.FILM_ID).Identity()
		t.Field(FILM.RELEASE_YEAR).Type("year")
		t.Field(FILM.RATING).Type("mpaa_rating").Default("'G'::mpaa_rating")
		t.Field(FILM.LAST_UPDATE).Type("TIMESTAMPTZ").Default("NOW()")
		t.Field(FILM.SPECIAL_FEATURES).Type("TEXT[]") // TODO: ArrayField
		t.Field(FILM.FULLTEXT).Type("TSVECTOR")
	case "mysql":
		t.Field(FILM.FILM_ID).Autoincrement()
		t.Field(FILM.TITLE).Type("VARCHAR(255)")
		t.Field(FILM.DESCRIPTION).Type("TEXT")
		t.Field(FILM.RATING).Type("ENUM('G','PG','PG-13','R','NC-17')")
		t.Field(FILM.LAST_UPDATE).Type("TIMESTAMP").Default("CURRENT_TIMESTAMP").OnUpdateCurrentTimestamp()
		t.Check("film_release_year_check", "%[1]s >= 1901 AND %[1]s <= 2155", FILM.RELEASE_YEAR)
	case "sqlite3":
		t.Check("film_release_year_check", "%[1]s >= 1901 AND %[1]s <= 2155", FILM.RELEASE_YEAR)
		t.Check("film_rating_check", "%s IN ('G','PG','PG-13','R','NC-17')", FILM.RATING)
	}
}

type _FILM_TEXT struct {
	tableinfo   `ddl:"name=film_text fts5={content='film' content_rowid='film_id'}"`
	FILM_ID     numberfield
	TITLE       stringfield
	DESCRIPTION stringfield
}

func (FILM_TEXT _FILM_TEXT) Constraints(dialect string, t T) {
	switch dialect {
	case "postgres": // no-op, we will ignore this table if postgres
	case "mysql":
		t.Field(FILM_TEXT.TITLE).Type("VARCHAR(255)").NotNull()
		t.Index(FILM_TEXT.TITLE, FILM_TEXT.DESCRIPTION).Using("FULLTEXT")
	case "sqlite3":
		t.Field(FILM_TEXT.FILM_ID).Ignore()
	}
}

type _FILM_ACTOR struct {
	tableinfo   `ddl:"name=film_actor index={. cols=actor_id,film_id unique}"`
	ACTOR_ID    numberfield `ddl:"notnull references={actor onupdate=cascade ondelete=restrict}"`
	FILM_ID     numberfield `ddl:"notnull references={film onupdate=cascade ondelete=restrict} index"`
	LAST_UPDATE timefield   `ddl:"default=DATETIME('now') notnull"`
}

func (FILM_ACTOR _FILM_ACTOR) Constraints(dialect string, t T) {
	t.Index(FILM_ACTOR.ACTOR_ID, FILM_ACTOR.FILM_ID).Unique()
	switch dialect {
	case "postgres":
		t.Field(FILM_ACTOR.LAST_UPDATE).Type("TIMESTAMPTZ").Default("NOW()")
	case "mysql":
		t.Field(FILM_ACTOR.LAST_UPDATE).Type("TIMESTAMP").Default("CURRENT_TIMESTAMP").OnUpdateCurrentTimestamp()
	}
}

type _FILM_CATEGORY struct {
	tableinfo   `ddl:"name=film_category"`
	FILM_ID     numberfield `ddl:"notnull references={film onupdate=cascade ondelete=restrict}"`
	CATEGORY_ID numberfield `ddl:"notnull references={category onupdate=cascade ondelete=restrict}"`
	LAST_UPDATE timefield   `ddl:"default=DATETIME('now') notnull"`
}

func (FILM_CATEGORY _FILM_CATEGORY) Constraints(dialect string, t T) {
	switch dialect {
	case "postgres":
		t.Field(FILM_CATEGORY.LAST_UPDATE).Type("TIMESTAMPTZ").Default("NOW()")
	case "mysql":
		t.Field(FILM_CATEGORY.LAST_UPDATE).Type("TIMESTAMP").Default("CURRENT_TIMESTAMP").OnUpdateCurrentTimestamp()
	}
}

type _STAFF struct {
	tableinfo   `ddl:"name=staff"`
	STAFF_ID    numberfield `ddl:"type=INTEGER primarykey"`
	FIRST_NAME  stringfield `ddl:"notnull"`
	LAST_NAME   stringfield `ddl:"notnull"`
	ADDRESS_ID  numberfield `ddl:"notnull references={address onupdate=cascade ondelete=restrict}"`
	EMAIL       stringfield
	STORE_ID    numberfield  `ddl:"references=store"`
	ACTIVE      booleanfield `ddl:"default=TRUE notnull"`
	USERNAME    stringfield  `ddl:"notnull"`
	PASSWORD    stringfield
	LAST_UPDATE timefield `ddl:"default=DATETIME('now') notnull"`
	PICTURE     blobfield
}

func (STAFF _STAFF) Constraints(dialect string, t T) {
	switch dialect {
	case "postgres":
		t.Field(STAFF.STAFF_ID).Identity()
		t.Field(STAFF.LAST_UPDATE).Type("TIMESTAMPTZ").Default("NOW()")
		t.Field(STAFF.PICTURE).Type("BYTEA")
	case "mysql":
		t.Field(STAFF.STAFF_ID).Autoincrement()
		t.Field(STAFF.FIRST_NAME).Type("VARCHAR(45)")
		t.Field(STAFF.LAST_NAME).Type("VARCHAR(45)")
		t.Field(STAFF.EMAIL).Type("VARCHAR(50)")
		t.Field(STAFF.USERNAME).Type("VARCHAR(16)")
		t.Field(STAFF.PASSWORD).Type("VARCHAR(40)")
		t.Field(STAFF.LAST_UPDATE).Type("TIMESTAMP").Default("CURRENT_TIMESTAMP").OnUpdateCurrentTimestamp()
	}
}

type _STORE struct {
	tableinfo        `ddl:"name=store"`
	STORE_ID         numberfield `ddl:"type=INTEGER primarykey"`
	MANAGER_STAFF_ID numberfield `ddl:"notnull references={staff onupdate=cascade ondelete=restrict} index={. unique}"`
	ADDRESS_ID       numberfield `ddl:"notnull references={address onupdate=cascade ondelete=restrict}"`
	LAST_UPDATE      timefield   `ddl:"default=DATETIME('now') notnull"`
}

func (STORE _STORE) Constraints(dialect string, t T) {
	switch dialect {
	case "postgres":
		t.Field(STORE.STORE_ID).Identity()
		t.Field(STORE.LAST_UPDATE).Type("TIMESTAMPTZ").Default("NOW()")
	case "mysql":
		t.Field(STORE.STORE_ID).Autoincrement()
		t.Field(STORE.LAST_UPDATE).Type("TIMESTAMP").Default("CURRENT_TIMESTAMP").OnUpdateCurrentTimestamp()
	}
}

type _CUSTOMER struct {
	tableinfo   `ddl:"name=customer unique={. cols=email,first_name,last_name}"`
	CUSTOMER_ID numberfield  `ddl:"type=INTEGER primarykey"`
	STORE_ID    numberfield  `ddl:"notnull index"`
	FIRST_NAME  stringfield  `ddl:"notnull"`
	LAST_NAME   stringfield  `ddl:"notnull index"`
	EMAIL       stringfield  `ddl:"unique"`
	ADDRESS_ID  numberfield  `ddl:"notnull references={address onupdate=cascade ondelete=restrict} index"`
	ACTIVE      booleanfield `ddl:"default=TRUE notnull"`
	DATA        jsonfield
	CREATE_DATE timefield `ddl:"default=DATETIME('now') notnull"`
	LAST_UPDATE timefield `ddl:"default=DATETIME('now')"`
}

func (CUSTOMER _CUSTOMER) Constraints(dialect string, t T) {
	switch dialect {
	case "postgres":
		t.Field(CUSTOMER.CUSTOMER_ID).Identity()
		t.Field(CUSTOMER.CREATE_DATE).Type("TIMESTAMPTZ").Default("NOW()")
		t.Field(CUSTOMER.LAST_UPDATE).Type("TIMESTAMPTZ").Default("NOW()")
	case "mysql":
		t.Field(CUSTOMER.CUSTOMER_ID).Autoincrement()
		t.Field(CUSTOMER.FIRST_NAME).Type("VARCHAR(45)")
		t.Field(CUSTOMER.LAST_NAME).Type("VARCHAR(45)")
		t.Field(CUSTOMER.EMAIL).Type("VARCHAR(50)")
		t.Field(CUSTOMER.CREATE_DATE).Type("TIMESTAMP").Default("CURRENT_TIMESTAMP")
		t.Field(CUSTOMER.LAST_UPDATE).Type("TIMESTAMP").Default("CURRENT_TIMESTAMP").OnUpdateCurrentTimestamp()
	}
}

type _INVENTORY struct {
	tableinfo    `ddl:"name=inventory index={. cols=store_id,film_id}"`
	INVENTORY_ID numberfield `ddl:"type=INTEGER primarykey"`
	FILM_ID      numberfield `ddl:"notnull references={film onupdate=cascade ondelete=restrict}"`
	STORE_ID     numberfield `ddl:"notnull references={store onupdate=cascade ondelete=restrict}"`
	LAST_UPDATE  timefield   `ddl:"default=DATETIME('now') notnull"`
}

func (INVENTORY _INVENTORY) Constraints(dialect string, t T) {
	switch dialect {
	case "postgres":
		t.Field(INVENTORY.INVENTORY_ID).Identity()
		t.Field(INVENTORY.LAST_UPDATE).Type("TIMESTAMPTZ").Default("NOW()")
	case "mysql":
		t.Field(INVENTORY.INVENTORY_ID).Autoincrement()
		t.Field(INVENTORY.LAST_UPDATE).Type("TIMESTAMP").Default("CURRENT_TIMESTAMP").OnUpdateCurrentTimestamp()
	}
}

type _RENTAL struct {
	tableinfo    `ddl:"name=rental index={. cols=rental_date,inventory_id,customer_id unique}"`
	RENTAL_ID    numberfield `ddl:"type=INTEGER primarykey"`
	RENTAL_DATE  timefield   `ddl:"notnull"`
	INVENTORY_ID numberfield `ddl:"notnull index references={inventory onupdate=cascade ondelete=restrict}"`
	CUSTOMER_ID  numberfield `ddl:"notnull index references={customer onupdate=cascade ondelete=restrict}"`
	RETURN_DATE  timefield
	STAFF_ID     numberfield `ddl:"notnull index references={staff onupdate=cascade ondelete=restrict}"`
	LAST_UPDATE  timefield   `ddl:"default=DATETIME('now') notnull"`
}

func (RENTAL _RENTAL) Constraints(dialect string, t T) {
	switch dialect {
	case "postgres":
		t.Field(RENTAL.RENTAL_ID).Identity()
		t.Field(RENTAL.RETURN_DATE).Type("TIMESTAMPTZ")
		t.Field(RENTAL.LAST_UPDATE).Type("TIMESTAMPTZ").Default("NOW()")
	case "mysql":
		t.Field(RENTAL.RENTAL_ID).Autoincrement()
		t.Field(RENTAL.RETURN_DATE).Type("TIMESTAMP")
		t.Field(RENTAL.LAST_UPDATE).Type("TIMESTAMP").Default("CURRENT_TIMESTAMP").OnUpdateCurrentTimestamp()
	}
}

type _PAYMENT struct {
	tableinfo    `ddl:"name=payment"`
	PAYMENT_ID   numberfield `ddl:"type=INTEGER primarykey"`
	CUSTOMER_ID  numberfield `ddl:"notnull index references={customer onupdate=cascade ondelete=restrict}"`
	STAFF_ID     numberfield `ddl:"notnull index references={staff onupdate=cascade ondelete=restrict}"`
	RENTAL_ID    numberfield `ddl:"references={rental onupdate=cascade ondelete=restrict}"`
	AMOUNT       numberfield `ddl:"type=DECIMAL(5,2) notnull"`
	PAYMENT_DATE timefield   `ddl:"notnull"`
}

func (PAYMENT _PAYMENT) Constraints(dialect string, t T) {
	switch dialect {
	case "postgres":
		t.Field(PAYMENT.PAYMENT_ID).Identity()
		t.Field(PAYMENT.PAYMENT_DATE).Type("TIMESTAMPTZ")
	case "mysql":
		t.Field(PAYMENT.PAYMENT_ID).Autoincrement()
		t.Field(PAYMENT.PAYMENT_DATE).Type("TIMESTAMP")
	}
}

type _DUMMY_TABLE struct {
	tableinfo `ddl:"name=dummy_table primarykey={. cols=id1,id2} unique={. cols=score,color}"`
	ID1       numberfield
	ID2       stringfield
	SCORE     numberfield
	COLOR     stringfield `ddl:"collate=nocase default='red'"`
	DATA      jsonfield
}

func (DUMMY_TABLE _DUMMY_TABLE) Constraints(dialect string, t T) {
	t.Check("dummy_table_score_positive_check", "%s > 0", DUMMY_TABLE.SCORE)
	t.Check("dummy_table_score_id1_greater_than_check", "%s > %s", DUMMY_TABLE.SCORE, DUMMY_TABLE.ID1)
	t.PrimaryKey(DUMMY_TABLE.ID1, DUMMY_TABLE.ID2)
	t.Unique(DUMMY_TABLE.SCORE, DUMMY_TABLE.COLOR)
	switch dialect {
	case "postgres":
		t.Field(DUMMY_TABLE.COLOR).Collate("C")
		t.NameIndex("dummy_table_score_color_data_idx").
			Fields(DUMMY_TABLE.SCORE).
			Expr("(%s->>'age')::INT", DUMMY_TABLE.DATA).
			Fields(DUMMY_TABLE.COLOR).
			Where("%s = 'red'", DUMMY_TABLE.COLOR)
	case "mysql":
		t.Field(DUMMY_TABLE.COLOR).Type("VARCHAR(50)").Collate("latin_swedish_ci")
		t.NameIndex("dummy_table_score_color_data_idx").
			Fields(DUMMY_TABLE.SCORE).
			Expr("CAST(%s->>'$.age' AS SIGNED)", DUMMY_TABLE.DATA).
			Fields(DUMMY_TABLE.COLOR)
	case "sqlite3":
		t.Field(DUMMY_TABLE.COLOR).Collate("nocase")
		t.NameIndex("dummy_table_score_color_data_idx").
			Fields(DUMMY_TABLE.SCORE).
			Expr("CAST(JSON_EXTRACT(%s, '$.age') AS INT)", DUMMY_TABLE.DATA).
			Fields(DUMMY_TABLE.COLOR).
			Where("%s = 'red'", DUMMY_TABLE.COLOR)
	}
}
