1.hintfile是什么
答：hintfile存储了entry，entry为key:key val:pos，记录了在mergeDB中key的pos，用于恢复mergeDB时构建索引

2.为什么读用mmap，写不用

3.dataFile.ReadLogRecord 函数 为什么判断 offset+maxLogRecordHeaderSize > fileSize，recordHeader的长度可能不固定么？可能没有keySize和valSize？

4.b tree为什么读不上读锁

5.merge.go中的merge函数，已经使用了lock，isMerging是否没必要

6.couloyDB的事务是如何设计的

7.couloyDB的集群是如何设计的

8.需要生成新的activeFile时（文件大小达到限制），会对old activeFile或者所有inactiveFiles进行merge操作么
答：不会

9.merge时，inactiveFiles的数据太多该怎么处理
答：couloyDB的merge实现时构建一个新的DB，数据太多会生成多个files