package gorm_related

import (
	"testing"
)

func TestCreateArticle(t *testing.T) {
	defer db.Close()
	db.DropTableIfExists(&Article{})
	if err := db.CreateTable(&Article{}).Error; err != nil {
		t.Error("Create Table error: ", err)
	}
}

func TestCallBack(t *testing.T) {
	defer db.Close()
	article := &Article{
		ID:        1,
		Title:     "zzk",
		Desc:      "myself",
		Content:   "I like animation and songs, they are interesting and great",
		CreatedBy: "fms5cmS",
		UpdatedBy: "fms5cmS",
	}
	if err := db.Create(&article).Error; err != nil {
		t.Error("insert error: ", err)
	}
	// 测试删除的回调函数
	db.Delete(article)
}
