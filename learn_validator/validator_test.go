package learn_validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
	"testing"
	"time"
)

type User struct {
	Name string `validate:"ne=admin"` // ne 不等于
	Age  int    `validate:"gte=18"`   // gte  大于等于
	// oneof 只能是列举出的值其中一个，这些值必须是数值或字符串，以空格分隔，如果字符串中有空格，将字符串用单引号包围
	Sex string `validate:"oneof=male female"`
	// 注意如果字段类型是time.Time，使用 gt/gte/lt/lte 等约束时不用指定参数值，默认与当前的 UTC 时间比较
	RegTime time.Time `validate:"lte"` // lte 小于等于
}

func TestValidatorConstraint(t *testing.T) {
	// 创建验证器，这个验证器可以指定选项、添加自定义约束
	validate := validator.New()
	// 使用 Struct() 验证各种结构对象的字段是否符合定义的约束
	u1 := User{Name: "dj", Age: 18, Sex: "male", RegTime: time.Now().UTC()}
	err := validate.Struct(u1)
	if err != nil {
		fmt.Println("User1 validate error: \n", err)
	}
	
	u2 := User{Name: "admin", Age: 15, Sex: "none", RegTime: time.Now().UTC().Add(1 * time.Hour)}
	err = validate.Struct(u2)
	if err != nil {
		fmt.Println("User2 validate error: \n", err)
	}
}

type RegisterForm struct {
	Name     string `validate:"min=2"`
	Age      int    `validate:"min=18"`
	Password string `validate:"min=10"`
	// 该字段的值要和 Password 字段的值相等
	Pwd      string `validate:"eqfield=Password"`
}

func TestCrossFieldConstraint(t *testing.T) {
	validate := validator.New()
	f1 := RegisterForm{
		Name:     "zz",
		Age:      18,
		Password: "123456789",
		Pwd:      "123456789",
	}
	err := validate.Struct(f1)
	if err != nil {
		fmt.Println("f1 validate error: \n", err)
	}
	f2 := RegisterForm{
		Name:     "dj",
		Age:      18,
		Password: "1234567890",
		Pwd:      "123",
	}
	err = validate.Struct(f2)
	if err != nil {
		fmt.Println("f2 validate error: \n", err)
	}
}

func TestVarWithValue(t *testing.T) {
	str1 := "my name is fms5cms"
	str2 := "ms5cms"
	validate := validator.New()
	t.Log(validate.VarWithValue(str1,str2,"eqfield"))
	t.Log(validate.VarWithValue(str1,str2,"nefield"))
}

type Information struct {
	Info string `validate:"startwithfS"`
}

func CheckStartWithfS(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return strings.HasPrefix(value,"fS")
}

func TestCustomize(t *testing.T) {
	validate := validator.New()
	// 注册校验规则
	validate.RegisterValidation("startwithfS",CheckStartWithfS)
	question := Information{"fS is what?"}
	answer := Information{"my favorite group"}
	t.Log(validate.Struct(question))
	t.Log(validate.Struct(answer))
}