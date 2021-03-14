package gorm_related

import (
	"log"
	"testing"
)

func TestHeroArgs_QueryFirst(t *testing.T) {
	args := heroArgs{ID: 10009}
	hero, err := args.QueryFirst()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", hero)
}

func TestHeroArgs_QueryLast(t *testing.T) {
	args := heroArgs{RoleMain: "法师"}
	hero, err := args.QueryLast()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", hero)
}

func TestHeroArgs_QueryFind(t *testing.T) {
	args := heroArgs{
		ID:       10029,
		RoleMain: "战士",
	}
	heroes, err := args.QueryFind()
	if err != nil {
		t.Fatal(err)
	}
	for _, hero := range heroes {
		t.Logf("%+v\n", hero)
	}
}

func TestHeroArgs_QuerySub(t *testing.T) {
	args := heroArgs{}
	heroes, err := args.QuerySub()
	if err != nil {
		t.Fatal(err)
	}
	for _, hero := range heroes {
		t.Logf("%+v", hero)
	}
}

func TestHeroArgs_QueryManyAndCount(t *testing.T) {
	args := heroArgs{ID: 10029}
	heroes, count, err := args.QueryManyAndCount()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(count)
	t.Log(len(heroes) == count)
	for _, hero := range heroes {
		t.Log(hero)
	}
}

func TestHeroArgs_QueryGroupByRole(t *testing.T) {
	args := heroArgs{}
	count, role := args.QueryGroupByRole()
	for i := 0; i < len(count); i++ {
		t.Logf("%s has %d", role[i], count[i])
	}
	res, err := args.QueryGroup()
	if err != nil {
		log.Fatal(err)
	}
	for _, re := range res {
		t.Logf("%s has %d", re.Role, re.Num)
	}
}

func TestTest(t *testing.T) {
	defer db.Close()
	hero := new(Hero)
	db.Table("heros").Select("*").Where("id = 10000").Find(hero)
	t.Logf("%+v", hero)
}
