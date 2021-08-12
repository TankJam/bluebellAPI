package test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()  // 断言 DB.Get() 方法是否被调用

	// 获取mock对象
	m := NewMockDB(ctrl)

	// 通过mock给Get方法传入参数Tom
	m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("no exist"))

	if v := GetFromDB(m, "Tom"); v != -1{
		t.Fatal("expected -1, but got", v)
	}

}