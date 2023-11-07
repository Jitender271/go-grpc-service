package db

type Table interface {
	Get(columns ...string) (stmt string, names []string)
	Insert() (stmt string, names []string)
	Update(columns ...string)(stmt string, names []string)
	Select(columns ...string)(stmt string, names []string)
	SelectAll() (stmt string, names []string)
}