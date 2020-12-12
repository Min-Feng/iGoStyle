package domain

type LanguageMapTableRepo interface {
	LanguageMapTable() (LanguageMapTable, error)
}

// 從 sql 文件
// 無法確定 LanguageID 與 LanguageSlug 的關係
// 先猜是 1對1
type LanguageMapTable struct {
	MapSlugToID map[LanguageSlug]LanguageID
}

func (t LanguageMapTable) SlugToID(slug LanguageSlug) LanguageID {
	return t.MapSlugToID[slug]
}

type LanguageID int
type LanguageSlug string
