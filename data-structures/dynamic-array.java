
public class Array<E> {

    private E[] data;
    private int size;


    public Array(int capacity){
        data = (E[])new Object[capacity];
        size = 0;
    }

    public Array(){
        this(10);
    }

    public int getCapacity(){
        return data.length;
    }

    public int getSize(){
        return size;
    }


    public boolean isEmpty(){
        return size == 0;
    }

    // 在index索引的位置插入一个新元素e , 后移
    public void add(int index, E e){

        if(index < 0 || index > size)
            throw new IllegalArgumentException("Add failed. Require index >= 0 and index <= size.");

        if(size == data.length)
            resize(2 * data.length); // 满了扩为2倍

        for(int i = size - 1; i >= index ; i --)
            data[i + 1] = data[i];

        data[index] = e;

        size ++;
    }


    public void addLast(E e){
        add(size, e);
    }


    public void addFirst(E e){
        add(0, e);
    }


    public E get(int index){
        if(index < 0 || index >= size)
            throw new IllegalArgumentException("Get failed. Index is illegal.");
        return data[index];
    }

    // 修改
    public void set(int index, E e){
        if(index < 0 || index >= size)
            throw new IllegalArgumentException("Set failed. Index is illegal.");
        data[index] = e;
    }


    public boolean contains(E e){
        for(int i = 0 ; i < size ; i ++){
            if(data[i].equals(e))
                return true;
        }
        return false;
    }


    public int find(E e){
        for(int i = 0 ; i < size ; i ++){
            if(data[i].equals(e))
                return i;
        }
        return -1;
    }


    public E remove(int index){
        if(index < 0 || index >= size)
            throw new IllegalArgumentException("Remove failed. Index is illegal.");

        //数据移动到新数组空间
        E ret = data[index];
        for(int i = index + 1 ; i < size ; i ++)
            data[i - 1] = data[i];
        size --;
        data[size] = null;

        if(size == data.length / 4 && data.length / 2 != 0) //size=cap的1/4的时候缩小cap为1/2 懒缩容,防止频繁扩容缩容
            resize(data.length / 2);
        return ret;
    }


    public E removeFirst(){
        return remove(0);
    }


    public E removeLast(){
        return remove(size - 1);
    }


    public void removeElement(E e){
        int index = find(e);
        if(index != -1)
            remove(index);
    }


    // 将数组空间的容量变成newCapacity大小
    private void resize(int newCapacity){

        E[] newData = (E[])new Object[newCapacity];
        //数据移动到新数组空间
        for(int i = 0 ; i < size ; i ++)
            newData[i] = data[i];
        data = newData;
    }
}
