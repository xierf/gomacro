// this file was generated by gomacro command: import _b "database/sql/driver"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"context"
	"database/sql/driver"
	"reflect"
)

// reflection: allow interpreted code to import "database/sql/driver"
func init() {
	Packages["database/sql/driver"] = Package{
	Binds: map[string]Value{
		"Bool":	ValueOf(&driver.Bool).Elem(),
		"DefaultParameterConverter":	ValueOf(&driver.DefaultParameterConverter).Elem(),
		"ErrBadConn":	ValueOf(&driver.ErrBadConn).Elem(),
		"ErrSkip":	ValueOf(&driver.ErrSkip).Elem(),
		"Int32":	ValueOf(&driver.Int32).Elem(),
		"IsScanValue":	ValueOf(driver.IsScanValue),
		"IsValue":	ValueOf(driver.IsValue),
		"ResultNoRows":	ValueOf(&driver.ResultNoRows).Elem(),
		"String":	ValueOf(&driver.String).Elem(),
	},
	Types: map[string]Type{
		"ColumnConverter":	TypeOf((*driver.ColumnConverter)(nil)).Elem(),
		"Conn":	TypeOf((*driver.Conn)(nil)).Elem(),
		"ConnBeginTx":	TypeOf((*driver.ConnBeginTx)(nil)).Elem(),
		"ConnPrepareContext":	TypeOf((*driver.ConnPrepareContext)(nil)).Elem(),
		"Driver":	TypeOf((*driver.Driver)(nil)).Elem(),
		"Execer":	TypeOf((*driver.Execer)(nil)).Elem(),
		"ExecerContext":	TypeOf((*driver.ExecerContext)(nil)).Elem(),
		"IsolationLevel":	TypeOf((*driver.IsolationLevel)(nil)).Elem(),
		"NamedValue":	TypeOf((*driver.NamedValue)(nil)).Elem(),
		"NotNull":	TypeOf((*driver.NotNull)(nil)).Elem(),
		"Null":	TypeOf((*driver.Null)(nil)).Elem(),
		"Pinger":	TypeOf((*driver.Pinger)(nil)).Elem(),
		"Queryer":	TypeOf((*driver.Queryer)(nil)).Elem(),
		"QueryerContext":	TypeOf((*driver.QueryerContext)(nil)).Elem(),
		"Result":	TypeOf((*driver.Result)(nil)).Elem(),
		"Rows":	TypeOf((*driver.Rows)(nil)).Elem(),
		"RowsAffected":	TypeOf((*driver.RowsAffected)(nil)).Elem(),
		"RowsColumnTypeDatabaseTypeName":	TypeOf((*driver.RowsColumnTypeDatabaseTypeName)(nil)).Elem(),
		"RowsColumnTypeLength":	TypeOf((*driver.RowsColumnTypeLength)(nil)).Elem(),
		"RowsColumnTypeNullable":	TypeOf((*driver.RowsColumnTypeNullable)(nil)).Elem(),
		"RowsColumnTypePrecisionScale":	TypeOf((*driver.RowsColumnTypePrecisionScale)(nil)).Elem(),
		"RowsColumnTypeScanType":	TypeOf((*driver.RowsColumnTypeScanType)(nil)).Elem(),
		"RowsNextResultSet":	TypeOf((*driver.RowsNextResultSet)(nil)).Elem(),
		"Stmt":	TypeOf((*driver.Stmt)(nil)).Elem(),
		"StmtExecContext":	TypeOf((*driver.StmtExecContext)(nil)).Elem(),
		"StmtQueryContext":	TypeOf((*driver.StmtQueryContext)(nil)).Elem(),
		"Tx":	TypeOf((*driver.Tx)(nil)).Elem(),
		"TxOptions":	TypeOf((*driver.TxOptions)(nil)).Elem(),
		"Value":	TypeOf((*driver.Value)(nil)).Elem(),
		"ValueConverter":	TypeOf((*driver.ValueConverter)(nil)).Elem(),
		"Valuer":	TypeOf((*driver.Valuer)(nil)).Elem(),
	},
	Proxies: map[string]Type{
		"ColumnConverter":	TypeOf((*ColumnConverter_database_sql_driver)(nil)).Elem(),
		"Conn":	TypeOf((*Conn_database_sql_driver)(nil)).Elem(),
		"ConnBeginTx":	TypeOf((*ConnBeginTx_database_sql_driver)(nil)).Elem(),
		"ConnPrepareContext":	TypeOf((*ConnPrepareContext_database_sql_driver)(nil)).Elem(),
		"Driver":	TypeOf((*Driver_database_sql_driver)(nil)).Elem(),
		"Execer":	TypeOf((*Execer_database_sql_driver)(nil)).Elem(),
		"ExecerContext":	TypeOf((*ExecerContext_database_sql_driver)(nil)).Elem(),
		"Pinger":	TypeOf((*Pinger_database_sql_driver)(nil)).Elem(),
		"Queryer":	TypeOf((*Queryer_database_sql_driver)(nil)).Elem(),
		"QueryerContext":	TypeOf((*QueryerContext_database_sql_driver)(nil)).Elem(),
		"Result":	TypeOf((*Result_database_sql_driver)(nil)).Elem(),
		"Rows":	TypeOf((*Rows_database_sql_driver)(nil)).Elem(),
		"RowsColumnTypeDatabaseTypeName":	TypeOf((*RowsColumnTypeDatabaseTypeName_database_sql_driver)(nil)).Elem(),
		"RowsColumnTypeLength":	TypeOf((*RowsColumnTypeLength_database_sql_driver)(nil)).Elem(),
		"RowsColumnTypeNullable":	TypeOf((*RowsColumnTypeNullable_database_sql_driver)(nil)).Elem(),
		"RowsColumnTypePrecisionScale":	TypeOf((*RowsColumnTypePrecisionScale_database_sql_driver)(nil)).Elem(),
		"RowsColumnTypeScanType":	TypeOf((*RowsColumnTypeScanType_database_sql_driver)(nil)).Elem(),
		"RowsNextResultSet":	TypeOf((*RowsNextResultSet_database_sql_driver)(nil)).Elem(),
		"Stmt":	TypeOf((*Stmt_database_sql_driver)(nil)).Elem(),
		"StmtExecContext":	TypeOf((*StmtExecContext_database_sql_driver)(nil)).Elem(),
		"StmtQueryContext":	TypeOf((*StmtQueryContext_database_sql_driver)(nil)).Elem(),
		"Tx":	TypeOf((*Tx_database_sql_driver)(nil)).Elem(),
		"Value":	TypeOf((*Value_database_sql_driver)(nil)).Elem(),
		"ValueConverter":	TypeOf((*ValueConverter_database_sql_driver)(nil)).Elem(),
		"Valuer":	TypeOf((*Valuer_database_sql_driver)(nil)).Elem(),
	},
	Wrappers: map[string][]string{
	} }
}

// --------------- proxy for database/sql/driver.ColumnConverter ---------------
type ColumnConverter_database_sql_driver struct {
	Object	interface{}
	ColumnConverter_	func(idx int) driver.ValueConverter
}
func (Proxy *ColumnConverter_database_sql_driver) ColumnConverter(idx int) driver.ValueConverter {
	return Proxy.ColumnConverter_(idx)
}

// --------------- proxy for database/sql/driver.Conn ---------------
type Conn_database_sql_driver struct {
	Object	interface{}
	Begin_	func() (driver.Tx, error)
	Close_	func() error
	Prepare_	func(query string) (driver.Stmt, error)
}
func (Proxy *Conn_database_sql_driver) Begin() (driver.Tx, error) {
	return Proxy.Begin_()
}
func (Proxy *Conn_database_sql_driver) Close() error {
	return Proxy.Close_()
}
func (Proxy *Conn_database_sql_driver) Prepare(query string) (driver.Stmt, error) {
	return Proxy.Prepare_(query)
}

// --------------- proxy for database/sql/driver.ConnBeginTx ---------------
type ConnBeginTx_database_sql_driver struct {
	Object	interface{}
	BeginTx_	func(ctx context.Context, opts driver.TxOptions) (driver.Tx, error)
}
func (Proxy *ConnBeginTx_database_sql_driver) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	return Proxy.BeginTx_(ctx, opts)
}

// --------------- proxy for database/sql/driver.ConnPrepareContext ---------------
type ConnPrepareContext_database_sql_driver struct {
	Object	interface{}
	PrepareContext_	func(ctx context.Context, query string) (driver.Stmt, error)
}
func (Proxy *ConnPrepareContext_database_sql_driver) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	return Proxy.PrepareContext_(ctx, query)
}

// --------------- proxy for database/sql/driver.Driver ---------------
type Driver_database_sql_driver struct {
	Object	interface{}
	Open_	func(name string) (driver.Conn, error)
}
func (Proxy *Driver_database_sql_driver) Open(name string) (driver.Conn, error) {
	return Proxy.Open_(name)
}

// --------------- proxy for database/sql/driver.Execer ---------------
type Execer_database_sql_driver struct {
	Object	interface{}
	Exec_	func(query string, args []driver.Value) (driver.Result, error)
}
func (Proxy *Execer_database_sql_driver) Exec(query string, args []driver.Value) (driver.Result, error) {
	return Proxy.Exec_(query, args)
}

// --------------- proxy for database/sql/driver.ExecerContext ---------------
type ExecerContext_database_sql_driver struct {
	Object	interface{}
	ExecContext_	func(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error)
}
func (Proxy *ExecerContext_database_sql_driver) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	return Proxy.ExecContext_(ctx, query, args)
}

// --------------- proxy for database/sql/driver.Pinger ---------------
type Pinger_database_sql_driver struct {
	Object	interface{}
	Ping_	func(ctx context.Context) error
}
func (Proxy *Pinger_database_sql_driver) Ping(ctx context.Context) error {
	return Proxy.Ping_(ctx)
}

// --------------- proxy for database/sql/driver.Queryer ---------------
type Queryer_database_sql_driver struct {
	Object	interface{}
	Query_	func(query string, args []driver.Value) (driver.Rows, error)
}
func (Proxy *Queryer_database_sql_driver) Query(query string, args []driver.Value) (driver.Rows, error) {
	return Proxy.Query_(query, args)
}

// --------------- proxy for database/sql/driver.QueryerContext ---------------
type QueryerContext_database_sql_driver struct {
	Object	interface{}
	QueryContext_	func(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error)
}
func (Proxy *QueryerContext_database_sql_driver) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	return Proxy.QueryContext_(ctx, query, args)
}

// --------------- proxy for database/sql/driver.Result ---------------
type Result_database_sql_driver struct {
	Object	interface{}
	LastInsertId_	func() (int64, error)
	RowsAffected_	func() (int64, error)
}
func (Proxy *Result_database_sql_driver) LastInsertId() (int64, error) {
	return Proxy.LastInsertId_()
}
func (Proxy *Result_database_sql_driver) RowsAffected() (int64, error) {
	return Proxy.RowsAffected_()
}

// --------------- proxy for database/sql/driver.Rows ---------------
type Rows_database_sql_driver struct {
	Object	interface{}
	Close_	func() error
	Columns_	func() []string
	Next_	func(dest []driver.Value) error
}
func (Proxy *Rows_database_sql_driver) Close() error {
	return Proxy.Close_()
}
func (Proxy *Rows_database_sql_driver) Columns() []string {
	return Proxy.Columns_()
}
func (Proxy *Rows_database_sql_driver) Next(dest []driver.Value) error {
	return Proxy.Next_(dest)
}

// --------------- proxy for database/sql/driver.RowsColumnTypeDatabaseTypeName ---------------
type RowsColumnTypeDatabaseTypeName_database_sql_driver struct {
	Object	interface{}
	Close_	func() error
	ColumnTypeDatabaseTypeName_	func(index int) string
	Columns_	func() []string
	Next_	func(dest []driver.Value) error
}
func (Proxy *RowsColumnTypeDatabaseTypeName_database_sql_driver) Close() error {
	return Proxy.Close_()
}
func (Proxy *RowsColumnTypeDatabaseTypeName_database_sql_driver) ColumnTypeDatabaseTypeName(index int) string {
	return Proxy.ColumnTypeDatabaseTypeName_(index)
}
func (Proxy *RowsColumnTypeDatabaseTypeName_database_sql_driver) Columns() []string {
	return Proxy.Columns_()
}
func (Proxy *RowsColumnTypeDatabaseTypeName_database_sql_driver) Next(dest []driver.Value) error {
	return Proxy.Next_(dest)
}

// --------------- proxy for database/sql/driver.RowsColumnTypeLength ---------------
type RowsColumnTypeLength_database_sql_driver struct {
	Object	interface{}
	Close_	func() error
	ColumnTypeLength_	func(index int) (length int64, ok bool)
	Columns_	func() []string
	Next_	func(dest []driver.Value) error
}
func (Proxy *RowsColumnTypeLength_database_sql_driver) Close() error {
	return Proxy.Close_()
}
func (Proxy *RowsColumnTypeLength_database_sql_driver) ColumnTypeLength(index int) (length int64, ok bool) {
	return Proxy.ColumnTypeLength_(index)
}
func (Proxy *RowsColumnTypeLength_database_sql_driver) Columns() []string {
	return Proxy.Columns_()
}
func (Proxy *RowsColumnTypeLength_database_sql_driver) Next(dest []driver.Value) error {
	return Proxy.Next_(dest)
}

// --------------- proxy for database/sql/driver.RowsColumnTypeNullable ---------------
type RowsColumnTypeNullable_database_sql_driver struct {
	Object	interface{}
	Close_	func() error
	ColumnTypeNullable_	func(index int) (nullable bool, ok bool)
	Columns_	func() []string
	Next_	func(dest []driver.Value) error
}
func (Proxy *RowsColumnTypeNullable_database_sql_driver) Close() error {
	return Proxy.Close_()
}
func (Proxy *RowsColumnTypeNullable_database_sql_driver) ColumnTypeNullable(index int) (nullable bool, ok bool) {
	return Proxy.ColumnTypeNullable_(index)
}
func (Proxy *RowsColumnTypeNullable_database_sql_driver) Columns() []string {
	return Proxy.Columns_()
}
func (Proxy *RowsColumnTypeNullable_database_sql_driver) Next(dest []driver.Value) error {
	return Proxy.Next_(dest)
}

// --------------- proxy for database/sql/driver.RowsColumnTypePrecisionScale ---------------
type RowsColumnTypePrecisionScale_database_sql_driver struct {
	Object	interface{}
	Close_	func() error
	ColumnTypePrecisionScale_	func(index int) (precision int64, scale int64, ok bool)
	Columns_	func() []string
	Next_	func(dest []driver.Value) error
}
func (Proxy *RowsColumnTypePrecisionScale_database_sql_driver) Close() error {
	return Proxy.Close_()
}
func (Proxy *RowsColumnTypePrecisionScale_database_sql_driver) ColumnTypePrecisionScale(index int) (precision int64, scale int64, ok bool) {
	return Proxy.ColumnTypePrecisionScale_(index)
}
func (Proxy *RowsColumnTypePrecisionScale_database_sql_driver) Columns() []string {
	return Proxy.Columns_()
}
func (Proxy *RowsColumnTypePrecisionScale_database_sql_driver) Next(dest []driver.Value) error {
	return Proxy.Next_(dest)
}

// --------------- proxy for database/sql/driver.RowsColumnTypeScanType ---------------
type RowsColumnTypeScanType_database_sql_driver struct {
	Object	interface{}
	Close_	func() error
	ColumnTypeScanType_	func(index int) reflect.Type
	Columns_	func() []string
	Next_	func(dest []driver.Value) error
}
func (Proxy *RowsColumnTypeScanType_database_sql_driver) Close() error {
	return Proxy.Close_()
}
func (Proxy *RowsColumnTypeScanType_database_sql_driver) ColumnTypeScanType(index int) reflect.Type {
	return Proxy.ColumnTypeScanType_(index)
}
func (Proxy *RowsColumnTypeScanType_database_sql_driver) Columns() []string {
	return Proxy.Columns_()
}
func (Proxy *RowsColumnTypeScanType_database_sql_driver) Next(dest []driver.Value) error {
	return Proxy.Next_(dest)
}

// --------------- proxy for database/sql/driver.RowsNextResultSet ---------------
type RowsNextResultSet_database_sql_driver struct {
	Object	interface{}
	Close_	func() error
	Columns_	func() []string
	HasNextResultSet_	func() bool
	Next_	func(dest []driver.Value) error
	NextResultSet_	func() error
}
func (Proxy *RowsNextResultSet_database_sql_driver) Close() error {
	return Proxy.Close_()
}
func (Proxy *RowsNextResultSet_database_sql_driver) Columns() []string {
	return Proxy.Columns_()
}
func (Proxy *RowsNextResultSet_database_sql_driver) HasNextResultSet() bool {
	return Proxy.HasNextResultSet_()
}
func (Proxy *RowsNextResultSet_database_sql_driver) Next(dest []driver.Value) error {
	return Proxy.Next_(dest)
}
func (Proxy *RowsNextResultSet_database_sql_driver) NextResultSet() error {
	return Proxy.NextResultSet_()
}

// --------------- proxy for database/sql/driver.Stmt ---------------
type Stmt_database_sql_driver struct {
	Object	interface{}
	Close_	func() error
	Exec_	func(args []driver.Value) (driver.Result, error)
	NumInput_	func() int
	Query_	func(args []driver.Value) (driver.Rows, error)
}
func (Proxy *Stmt_database_sql_driver) Close() error {
	return Proxy.Close_()
}
func (Proxy *Stmt_database_sql_driver) Exec(args []driver.Value) (driver.Result, error) {
	return Proxy.Exec_(args)
}
func (Proxy *Stmt_database_sql_driver) NumInput() int {
	return Proxy.NumInput_()
}
func (Proxy *Stmt_database_sql_driver) Query(args []driver.Value) (driver.Rows, error) {
	return Proxy.Query_(args)
}

// --------------- proxy for database/sql/driver.StmtExecContext ---------------
type StmtExecContext_database_sql_driver struct {
	Object	interface{}
	ExecContext_	func(ctx context.Context, args []driver.NamedValue) (driver.Result, error)
}
func (Proxy *StmtExecContext_database_sql_driver) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	return Proxy.ExecContext_(ctx, args)
}

// --------------- proxy for database/sql/driver.StmtQueryContext ---------------
type StmtQueryContext_database_sql_driver struct {
	Object	interface{}
	QueryContext_	func(ctx context.Context, args []driver.NamedValue) (driver.Rows, error)
}
func (Proxy *StmtQueryContext_database_sql_driver) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	return Proxy.QueryContext_(ctx, args)
}

// --------------- proxy for database/sql/driver.Tx ---------------
type Tx_database_sql_driver struct {
	Object	interface{}
	Commit_	func() error
	Rollback_	func() error
}
func (Proxy *Tx_database_sql_driver) Commit() error {
	return Proxy.Commit_()
}
func (Proxy *Tx_database_sql_driver) Rollback() error {
	return Proxy.Rollback_()
}

// --------------- proxy for database/sql/driver.Value ---------------
type Value_database_sql_driver struct {
	Object	interface{}
}

// --------------- proxy for database/sql/driver.ValueConverter ---------------
type ValueConverter_database_sql_driver struct {
	Object	interface{}
	ConvertValue_	func(v interface{}) (driver.Value, error)
}
func (Proxy *ValueConverter_database_sql_driver) ConvertValue(v interface{}) (driver.Value, error) {
	return Proxy.ConvertValue_(v)
}

// --------------- proxy for database/sql/driver.Valuer ---------------
type Valuer_database_sql_driver struct {
	Object	interface{}
	Value_	func() (driver.Value, error)
}
func (Proxy *Valuer_database_sql_driver) Value() (driver.Value, error) {
	return Proxy.Value_()
}
