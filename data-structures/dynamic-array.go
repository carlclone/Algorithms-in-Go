package data_structures

// 自动扩容
// 扩容策略 : 满了*2 , size=cap的1/4的时候缩小cap为1/2 懒缩容
// 实际上切片就是一个类似的动态数组....这个例子不恰当, 仅仅为了描述动态数组的一些helper method
type DynamicArray struct {
	data [1000]interface{}
	size int
}

func ArrayConstructor(cap int) *DynamicArray {
	return &DynamicArray{
		data: [1000]interface{}{},
		size: 0,
	}
}

func (t *DynamicArray) GetCapacity() int {
	return len(t.data)
}

func (t *DynamicArray) GetSize() int {
	return t.size
}

func (t *DynamicArray) IsEmpty() bool {
	return t.size == 0
}

//在index的位置插入一个新元素 , 后移之后的元素
func (t *DynamicArray) Add(index int, value interface{}) {

}

func (t *DynamicArray) AddLast() {

}

func (t *DynamicArray) AddFirst() {

}

func (t *DynamicArray) Get(index int) {

}

//修改
func (t *DynamicArray) Set(index int, value interface{}) {

}

func (t *DynamicArray) Contains(e interface{}) {

}

func (t *DynamicArray) Find(e interface{}) {

}

//删除index位置元素并前移元素
func (t *DynamicArray) Remove(e interface{}) {

}

func (t *DynamicArray) RemoveFirst() {

}

func (t *DynamicArray) RemoveLast() {

}

func (t *DynamicArray) RemoveElement(e interface{}) {

}

func (t *DynamicArray) ToString() {

}

// resize cap
func (t *DynamicArray) Resize(newCap int) {

}
