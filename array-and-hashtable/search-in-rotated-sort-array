func search(nums []int, target int) int {
//0+6=6 6/2=3 nums[3]=7  7>2 , 在第一个有序数组里, 4~7 >0 , 3+6 / 2 = 4 , nums[4]=0 , return 4;
// 7 , 7>2 , 在第一个有序数组里, 4~7>3 , 3+6/2=4 , nums[4]=0 , 0<4 , 在第二个有序数组里 , 0~2<3 , 返回-1
//先找到第一个有序数组的右边界, 然后再判断在哪个有序数组中, 然后再进行bs
//case : 画图 , 看left right mid的各种落点

b:=findBoundary(nums)
left:=nums[0]
right:=nums[len(nums)-1]

res:=-1
if target>=left && target<=b {
    res=findIn(nums,left,b)
}
if target>b && target<=right {
    res=findIn(nums,b+1,right)
}
return res
}
