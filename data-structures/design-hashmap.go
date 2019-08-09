package data_structures

//Design a HashMap without using any built-in hash table libraries.
//
//To be specific, your design should include these functions:
//
//put(key, value) : Insert a (key, value) pair into the HashMap. If the value already exists in the HashMap, update the value.
//get(key): Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key.
//remove(key) : Remove the mapping for the value key if this map contains the mapping for the key.
//
//Example:
//
//MyHashMap hashMap = new MyHashMap();
//hashMap.put(1, 1);
//hashMap.put(2, 2);
//hashMap.get(1);            // returns 1
//hashMap.get(3);            // returns -1 (not found)
//hashMap.put(2, 1);          // update the existing value
//hashMap.get(2);            // returns 1
//hashMap.remove(2);          // remove the mapping for 2
//hashMap.get(2);            // returns -1 (not found)
//
//Note:
//
//All keys and values will be in the range of [0, 1000000].
//The number of operations will be in the range of [1, 10000].
//Please do not use the built-in HashMap library.

type MyHashMap struct {
	data []int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	data := make([]int, 1000000)
	for i := 0; i <= 999999; i++ {
		data[i] = -1
	}
	return MyHashMap{data: data}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	this.data[key] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	if this.data[key] != -1 {
		return this.data[key]
	} else {
		return -1
	}
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	this.data[key] = -1
}
