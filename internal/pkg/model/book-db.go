package model

// 用於實現gorm的Tabler介面，實作TableName方法就可指定gorm去指定的table取得資料
type Tabler interface {
	TableName() string
}

type DBBook struct {
	Isbn      int    `json:"isbn"`
	Name      string `josn:"name"`
	Publisher string `json:"publisher"`
}

func (DBBook) TableName() string {
	// 指定gorm去book table取得資料
	return "book"
}
