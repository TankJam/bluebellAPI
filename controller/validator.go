package controller

import (
	"bluebellAPI/models"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
	ut "github.com/go-playground/universal-translator"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// 定义一个全局翻译器T
var trans ut.Translator

// InitTrans 初始化翻译器
func InitTrans(local string) (err error) {
	// 修改gin框架中的Validator引擎属性，实现自定制

	// 1.断言，判断当前引擎是否属于gin框架的引擎
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 2.注册一个获取 json tag 的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {  // 传入一个匿名函数
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		// 3.为SignUpParam注册自定义检验方法
		/// 校验注册
		v.RegisterStructValidation(
			SignUpParamStructLevelValidation, models.ParamSignUp{},
		)

		// 4.配置中英文翻译器
		zhT := zh.New()
		enT := zh.New()
		uni := ut.New(enT, zhT, enT)  // 可以支持多个

		// local 中文翻译器配置
		var ok bool
		trans, ok = uni.GetTranslator(local)  // 校验翻译器配置参数
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", local)
		}

		// 注册翻译器
		switch local{
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

// removeTopStruct 去除提示信息中的结构体名称
func removeTopStruct(fields map[string]string) map[string]string{
	res := map[string]string{}  // 实例化空map
	// 将错误信息循环遍历，存入map中
	for field, err := range fields{
		res[field[strings.Index(field, ".") + 1:]] = err
	}
	return res
}


// SignUpParamStructLevelValidation  自定义SignUpParam结构体校验函数
func SignUpParamStructLevelValidation(sl validator.StructLevel){
	// 断言判断当前的注册结构体
	su := sl.Current().Interface().(models.ParamSignUp)
	// 校验密码
	if su.Password != su.RePassword {
		// 输出错误提示信息，最后一个参数就是传递的param
		sl.ReportError(
			su.RePassword, "re_password", "RePassword",
			"eqfield", "password",
		)
	}
}