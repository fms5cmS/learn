package gorm_related

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

// 注意下面的日期类型字段，在 Mysql 5.7 及以上版本中，默认是不允许设置日期的值为全0值的
// 使用 select @@sql_mode 可以看到有 NO_ZERO_IN_DATE 和 NO_ZERO_DATE
// 修改：set @@sql_mode=(select replace(@@sql_mode,'NO_ZERO_IN_DATE,NO_ZERO_DATE',''));
// 还需要修改全局的 sql_mode：set @@global.sql_mode=(select replace(@@global.sql_mode,'NO_ZERO_IN_DATE,NO_ZERO_DATE',''));
// 重连数据库，然后重新建表即可
type Article struct {
	ID        uint   `gorm:"type:tinyint(3);primary_key;AUTO_INCREMENT"`
	Title     string `gorm:"type:varchar(100);default:''"`
	Desc      string `gorm:"type:varchar(255);default:''"`
	Content   string `gorm:"type:text"`
	CreateOn  time.Time
	CreatedBy string `gorm:"type:varchar(100);default:''"`
	UpdatedOn time.Time
	UpdatedBy string `gorm:"type:varchar(100);default:''"`
	DeletedOn time.Time
}

// 实现回调函数 Callback
// 创建时设置 CreateOn、UpdatedOn 两个字段值
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().UTC()
		// 取所有字段，判断当前是否包含所需字段
		if createTimeField, ok := scope.FieldByName("CreateOn"); ok {
			// 判断该字段的值是否为空
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}
		if updatedTimeField, ok := scope.FieldByName("UpdatedOn"); ok {
			if updatedTimeField.IsBlank {
				updatedTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `UpdatedOn` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	// 根据入参获取设置了字面值的参数，这里是 gorm:update_column ，它会去查找含这个字面值的字段属性
	if _, ok := scope.Get("gorm:update_column"); !ok {
		// 假设没有指定 update_column 的字段，默认在更新回调设置 UpdatedAt 的值
		scope.SetColumn("UpdatedOn", time.Now().UTC())
	}
}

func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		// 检查是否手动指定了 delete_option
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}
		// 获取约定的删除字段，若存在则 UPDATE 软删除，若不存在则 DELETE 硬删除
		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				// 返回引用的表名，这个方法 GORM 会根据自身逻辑对表名进行一些处理
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				// AddToVars 可以添加值作为SQL的参数，也可用于防范SQL注入
				scope.AddToVars(time.Now().UTC()),
				// CombinedConditionSql 返回组合好的条件SQL
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
