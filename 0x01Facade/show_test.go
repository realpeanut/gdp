/**
 * @Author: realpeanut
 * @Date: 2020/10/29 10:32 下午
 */
package _x01Facade

import "testing"

func TestShow(t *testing.T)  {
	facade := Facade{}
	facade.FacadeMethodA()
	facade.FacadeMethodB()
	facade.FacadeMethodC()
}