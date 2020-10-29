/**
 * @Author: realpeanut
 * @Date: 2020/10/29 10:17 下午
 */
package _x00SimpleFactory

import (
	"testing"
)

func TestShowA(t *testing.T) {
	f := NewFactory(1)
	t.Log(f.DoSomeThing())
	if f.DoSomeThing() != "A" {
		t.Error("err")
	}
}

func TestShowB(t *testing.T) {
	f := NewFactory(2)
	t.Log(f.DoSomeThing())
	if f.DoSomeThing() != "B" {
		t.Error("err")
	}
}

func TestShowDefault(t *testing.T) {
	f := NewFactory(3)
	t.Log(f.DoSomeThing())
	if f.DoSomeThing() != "A" {
		t.Error("err")
	}
}