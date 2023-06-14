1.couloyDB的entry只存储了key,val,type(标记是否为删除数据)，entryHead没有暴露给上层；read时显然不需要用到entryHead，只需要offset即可；
 write时下层处理entryHead。 rosedb也是这样做的，tiny-bitcask的entry包括了entryHead，entry中的代码更简洁，不过还是couloy的设计更好。

 2.couloyDB使用了索引interface，实现方式有b tree, art, map。rosedb以前string使用了跳表，其余结构是map，现在都使用了art。

3.couloyDB的merge操作：
    1.加锁
    2.activeFile设置为oldFile，将oldFiles加入mergeFiles，并新建一个空的activeFile
    3.删除mergePath，并新建mergePath，新建mergeDB，新建hintFile
    4.遍历mergeFiles，遍历file中的entry，如何和索引中相同，说明其是最新数据，加入mergeDB
    5.sync mergeDB和hintFile
    6.在mergePath中加入一个mergeFinishedFile，并写入一条entry，value是当前新activeFile的id，标识merge完成


4.rosedb不同类型的数据，都有着不同的存储对象，比如string,set,hash有着不同的file,discard等。delete或者update时，会调用sendDiscard。
 discard记录了每个file的有效率（旧数据无效）。在merge时，将有效率不符合标准的file取出遍历，file中有效的entry重新写入当前activeFile，
 然后删除file。
 couloyDB的merge是创建了一个mergeDB，遍历所有files，将最新的数据（索引中的数据）写入mergeDB。在初始化db时，将mergeDB的file和hintfile的
 索引加载，还有构建fid > mergeDB的文件的索引。
 我觉得couloyDB的merge实现不好，因为merge后db不重启，还是使用旧数据，这样下次merge重新做了一遍重复操作，并且没有解决inactiveFiles数据
 过多的问题。

