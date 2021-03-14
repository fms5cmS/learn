package gorm_related

import (
	"testing"
)

type Info struct {
	StrId   int64
	Status  []uint8
	DelFlag []uint8
	Name    string
}

// status、del_flag 在数据库中的数据类型为 bit 类型，插入时使用 0x00、0x01 两种值，读取的时候目标结构类型为 []uint8
func TestBit(t *testing.T) {
	info := make([]Info, 0)
	db.Table("apm_custom_kv_info").Select("str_id, status, del_flag, name").Where("status = 0x00").Find(&info)
	for _, i := range info {
		t.Logf("%+v", i)
		t.Log(len(i.Status))
		t.Log(i.Status[0])
	}
}


