package gorm_related

import (
	"time"
)

type Hero struct {
	ID          uint64
	Name        string
	RoleMain    string
	Birthday    time.Time `gorm:"column:birthdate"`
	HpMax       float32
	MpMax       float32
	AttackMax   float32
	DefenseMax  float32
	AttackSpeed float32 `gorm:"column:attack_speed_max"`
	AttackRange string
	Test        string `gorm:"column:test"`
}

type heroArgs struct {
	ID       uint64
	Name     string
	RoleMain string
}

func (h *heroArgs) QueryFirst() (*Hero, error) {
	hero := new(Hero)
	// SELECT * FROM `heros` WHERE (id >= 1009) ORDER BY `heros`.`id` ASC LIMIT 1
	err := db.Table("heros").Where("id >= ?", h.ID).First(hero).Error
	return hero, err
}

func (h *heroArgs) QueryLast() (*Hero, error) {
	hero := new(Hero)
	// SELECT * FROM `heros`  WHERE (role_main = '法师') ORDER BY `heros`.`id` DESC LIMIT 1
	err := db.Table("heros").Where("role_main = ?", h.RoleMain).Last(hero).Error
	return hero, err
}

func (h *heroArgs) QueryFind() ([]*Hero, error) {
	heroes := make([]*Hero, 0)
	// SELECT * FROM `heros`  WHERE (role_main = '战士' AND id > 10029)
	err := db.Table("heros").Where("role_main = ? OR id > ?", h.RoleMain, h.ID).Find(&heroes).Error
	return heroes, err
}

func (h *heroArgs) QuerySub() ([]*Hero, error) {
	heroes := make([]*Hero, 0)
	err := db.Table("heros").Where("id > ?", db.Table("heros").Select("AVG(id)").SubQuery()).Find(&heroes).Error
	return heroes, err
}

func (h *heroArgs) QueryManyAndCount() ([]Hero, int, error) {
	heroes := make([]Hero, 0)
	count := 0
	err := db.Table("heros").Where("id > ?", h.ID).Find(&heroes).Count(&count).Error
	return heroes, count, err
}

func (h *heroArgs) QueryGroupByRole() ([]int, []string) {
	count, role := make([]int, 0), make([]string, 0)
	// SELECT COUNT(1) numbers, role_main FROM `heros`   GROUPBY role_main HAVING (numbers > 10)
	rows, _ := db.Table("heros").Select("COUNT(1) numbers, role_main").Group("role_main").Having("numbers > ?", 10).Rows()
	for rows.Next() {
		c, r := 0, ""
		rows.Scan(&c, &r)
		count = append(count, c)
		role = append(role, r)
	}
	return count, role
}

type Result struct {
	Num  int
	Role string
}

func (h *heroArgs) QueryGroup() ([]Result, error) {
	results := make([]Result, 0)
	// SELECT count(1) num, role_main role FROM `heros`   GROUP BY role
	// err := db.Table("heros").Select("count(1) num, role_main role").Group("role").Scan(&results).Error
	err := db.Table("heros").Select("count(1) num, role_main role").Group("role").Find(&results).Error
	return results, err
}
