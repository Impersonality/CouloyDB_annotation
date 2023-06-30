1.hintfile是什么
答：hintfile存储了entry，entry为key:key val:pos，记录了在mergeDB中key的pos，用于恢复mergeDB时构建索引。
    如果不用hintfile，recovery时需遍历所有files构建索引。

2.为什么读用mmap，写不用
答：因为go的mmap库没有sync函数。

3.dataFile.ReadLogRecord 函数 为什么判断 offset+maxLogRecordHeaderSize > fileSize，recordHeader的长度可能不固定么？可能没有keySize和valSize？
答：recordHeader一定有key/val size，keySize和valSize使用了varInt。这样判断的原因是，recordHeader的长度<=maxLogRecordHeaderSize，正常情况使用
   maxLogRecordHeaderSize来读取recordHeader，多读一些byte不影响，如果是file末尾，recordHeader长度可能 < maxLogRecordHeaderSize。

4.b tree为什么读不上读锁
答：应该上读锁的。

5.merge.go中的merge函数，已经使用了lock，isMerging是否没必要
答：有必要，lock是对磁盘file读写是加锁，isMerging是对merge操作加锁，isMerging除了修改磁盘file，还有其他操作。

6.couloyDB的事务是如何设计的

7.couloyDB的集群是如何设计的

8.需要生成新的activeFile时（文件大小达到限制），会对old activeFile或者所有inactiveFiles进行merge操作么
答：不会。

9.merge时，inactiveFiles的数据太多该怎么处理
答：couloyDB的merge实现时构建一个新的DB，数据太多会生成多个files。

10.为什么couloyDB的读写比其他bitcask慢很多
答：写是因为默认option开启了syncWrite，所以每次写都会flush到磁盘。
   读是因为默认选择了mmap，而mmap没有实现sync，为了避免脏读，每次都是openFile后mmap，再读。多了一个打开文件并mmap到内存的步骤。