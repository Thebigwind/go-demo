
slice支持并发吗？

    当指定索引使用切片时，切片是支持并发读写索引区的数据的，但是索引区的数据在并发时会被覆盖的；
    当不指定索引切片时，并且切片动态扩容时，并发场景下扩容会被覆盖，所以切片是不支持并发的～

