package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Demo struct {
	ID        uint      `xorm:"id" gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Username  string    `xorm:"username" gorm:"column:username;NOT NULL"`
}

func (m *Demo) TableName() string {
	return "demo"
}


func TestNew(t *testing.T) {
	tests := map[string]struct {
		test func(t *testing.T)
	}{

		"testGorm": { TestGorm },
		"testXorm": { TestXorm },
	}
	t.Parallel()
	for name, tt := range tests {
		t.Run(name, tt.test)
	}
}

func TestGorm(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	New(WithDsn(dsn), WithOrm("gorm"))

	demos := []Demo{
		{Username: "go"},
		{Username: "java"},
		{Username: "c++"},
		{Username: "python"},
	}
	err := GDB.CreateInBatches(demos, 4).Error
	assert.NoError(t, err, "insert data failed")


	as := make([]*Demo, 0)
	GDB.Scopes(Paginate(1, 10)).Find(&as)
	for _, v := range as {
		t.Logf("id = [%d] username = [%s]\n", v.ID, v.Username)
	}


	as2 := make([]*Demo, 0)
	GDB.Scopes(Paginate(1, 10)).
		Scopes(FilterString("username", "go", "=")).
		Find(&as2)
	for _, v := range as2 {
		t.Logf("id = [%d] username = [%s]\n", v.ID, v.Username)
	}


	as3 := make([]*Demo, 0)
	GDB.Scopes(Paginate(1, 10)).
		Scopes(InOrNotInFilter("username", []string{"go", "java"}, "=")).
		Find(&as3)
	for _, v := range as3 {
		t.Logf("id = [%d] username = [%s]\n", v.ID, v.Username)
	}

}

func TestXorm(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	New(WithDsn(dsn), WithOrm("xorm"))

	demos := make([]Demo, 0)
	err := XDB.Where("username != ''").Find(&demos)
	assert.NoError(t,err, "find demos failed")
	assert.NotEmpty(t, demos)

	for _, v := range demos {
		t.Logf("id = [%d] username = [%s]\n", v.ID, v.Username)
	}
}