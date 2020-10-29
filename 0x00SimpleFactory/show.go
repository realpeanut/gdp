/**
 * @Author: realpeanut
 * @Date: 2020/10/29 10:08 下午
 */
package _x00SimpleFactory

/**
	定义：定义一个工厂类，他可以根据参数的不同返回不同类的实例，被创建的实例通常都具有共同的父类
 */

/**
	NewFactory(status int) 可以根据不同的status返回不同的类的实体,就像一个工厂，源源不断的生产所需的产品
*/

/**
	定义工厂接口，类似PHP、JAVA的class
 */
type Factory interface {
	DoSomeThing()string
}

/**
	类似construct(),返回不同的类
 */
func NewFactory(status int) Factory {
	switch status {
	case 1:
		//也可以return &ClassA{}
		return ClassA{}
	case 2:
		return ClassB{}
	default:
		return ClassA{}
	}
}

type ClassA struct {

}

type ClassB struct {

}

func (a ClassA) DoSomeThing()string {
	return "A"
}

func (b ClassB) DoSomeThing()string {
	return "B"
}




