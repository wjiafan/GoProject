package pro_2

/**
 * Created by Administrator on 2020/8/11 0011 上午 10:22
 */

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// Hash maps bytes to unit32
// 自定义hash函数
type Hash func(data []byte) uint32

// Map constains all hashed keys
/**
	Map是一致性哈希算法的主数据结构，包含4个成员变量：
	【1】Hash 函数 hash
	【2】虚拟节点倍数 replicas
	【3】哈希环 keys
	【4】虚拟节点与真实节点的映射表hashMap,键是虚拟节点的哈希值，值是真实节点的名称
 */
type Map struct{
	hash Hash
	replicas int
	keys []int
	hashMap map[int]string
}

// New creates a Map instance
/**
	构造函数 New() 允许自定义虚拟节点倍数和 Hash 函数。
 */
func New(replicas int,fn Hash) *Map{
	m := &Map{
		replicas:	replicas,
		hash:		fn,
		hashMap:	make(map[int]string),
	}
	if m.hash == nil{
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// Add adds some keys to the hash.
/**
	实现添加真实节点/机器的 Add() 方法
 */
func (m *Map) Add(keys ...string){
	for _,key := range keys{
		for i := 0;i < m.replicas; i++ {
			//构建字符串 strconv.Itoa(i)+key
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			//fmt.Println(hash)
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	//fmt.Println(m.keys)
	sort.Ints(m.keys)
}

// Get gets the closest item in the hash to the provided key.
// Get获取散列中距离所提供的键最近的项。
func (m *Map) Get(key string) string {

	if len(m.keys) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))
	// Binary search for appropriate replica. 二分搜索
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})

	return m.hashMap[m.keys[idx%len(m.keys)]]
}






