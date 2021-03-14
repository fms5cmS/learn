package gorm_related

import (
	"testing"
)

// gorm.Model 以及带有 gorm Tag 的字段才会被创建！
func TestGorm_CreateTable(t *testing.T) {
	defer db.Close()
	// 删除表
	db.DropTableIfExists(&User{})
	if !db.HasTable(&User{}) {
		if err := db.CreateTable(&User{}).Error; err != nil {
			t.Log(err)
		}
	}
}

func TestGorm_InsertRecord(t *testing.T) {
	defer db.Close()
	user := User{
		Age:  24,
		Name: "Pen",
		Sex:  "female",
	}
	// 使用 Create 插入一条数据
	if err := db.Create(&user).Error; err != nil {
		t.Error("create record error: ", err)
	}
}

func TestGorm_Select(t *testing.T) {
	defer db.Close()
	user1 := User{}
	// 条件查询
	db.First(&user1, "id=?", "2")
	t.Logf("%v", user1)
	// 查询 id、name 字段，并将值填充到 user2 的对应属性上
	user2 := User{}
	db.Select("id,name").Where("sex=?", "female").First(&user2)
	t.Log(user2)
	// Model 指定运行 db 操作的模型实例，默认接卸该结构体的名字为表名，Val 构造查询条件，Count 计算数量
	var count int
	if err := db.Model(&User{}).Where(&User{Sex: "male"}).Count(&count).Error; err != nil {
		t.Error("Count error: ", err)
	}
	t.Log(count)
	// 分页查询：Offset(pageNum).Limit(pageSize)
	users := make([]User, 0)
	db.Where("id!=?", 0).Offset(1).Limit(3).Find(&users)
	t.Log(users)
}

func TestGorm_Update(t *testing.T) {
	defer db.Close()
	user := User{}
	db.First(&user, "id=?", 3)
	t.Logf("%v", user)
	// 更新单条记录
	db.Model(&User{}).Where("id=?", 3).Update("age", 25)
	db.First(&user, "id=?", 3)
	t.Logf("%v", user)
}

func TestGorm_Delete(t *testing.T) {
	defer db.Close()
	db.Create(&User{
		Age:  18,
		Name: "testable",
		Sex:  "female",
	})
	// 使用 Val 构造查询条件，再调用 Delete 即可
	if err := db.Where("name=?", "testable").Delete(User{}).Error; err != nil {
		t.Error("delete error: ", err)
	}
	t.Log(db.First(&User{}, "name=?", "testable").RecordNotFound())
}

func TestGorm_Transaction(t *testing.T) {
	defer db.Close()
	// 开启事务
	tx := db.Begin()
	if err := tx.Create(&User{Age: 17, Name: "Xiao", Sex: "female",}).Error; err != nil {
		// 回滚
		tx.Rollback()
		t.Error("transaction insert error: ", err)
	}
	if err := tx.Model(&User{}).Where("name=?", "ZhangSan").Update("name", "fms5cmS").Error; err != nil {
		tx.Rollback()
		t.Error("transaction update error: ", err)
	}
	// 提交事务
	tx.Commit()
}
