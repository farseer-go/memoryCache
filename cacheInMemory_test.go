package cacheMemory

import (
	"github.com/farseer-go/cache"
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/modules"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type po struct {
	Name string
	Age  int
}

func TestCacheInMemory_Set(t *testing.T) {
	modules.StartModules(Module{})

	cache.SetProfilesInMemory[po]("test", "Name", 0)
	cacheManage := cache.GetCacheManage[po]("test")
	lst := collections.NewList(po{Name: "steden", Age: 18}, po{Name: "steden2", Age: 19})
	cacheManage.Set(lst.ToArray()...)
	lst2 := cacheManage.Get()

	assert.Equal(t, lst.Count(), lst2.Count())

	for i := 0; i < lst.Count(); i++ {
		assert.Equal(t, lst.Index(i).Name, lst2.Index(i).Name)
		assert.Equal(t, lst.Index(i).Age, lst2.Index(i).Age)
	}
}

func TestCacheInMemory_GetItem(t *testing.T) {
	modules.StartModules(Module{})
	cache.SetProfilesInMemory[po]("test", "Name", 0)
	cacheManage := cache.GetCacheManage[po]("test")
	cacheManage.Set(po{Name: "steden", Age: 18}, po{Name: "steden2", Age: 19})
	item1, _ := cacheManage.GetItem("steden")

	assert.Equal(t, item1.Name, "steden")
	assert.Equal(t, item1.Age, 18)

	item2, _ := cacheManage.GetItem("steden2")

	assert.Equal(t, item2.Name, "steden2")
	assert.Equal(t, item2.Age, 19)
}

func TestCacheInMemory_SaveItem(t *testing.T) {
	modules.StartModules(Module{})
	cache.SetProfilesInMemory[po]("test", "Name", 0)
	cacheManage := cache.GetCacheManage[po]("test")
	cacheManage.Set(po{Name: "steden", Age: 18}, po{Name: "steden2", Age: 19})
	cacheManage.SaveItem(po{Name: "steden", Age: 99})
	item1, _ := cacheManage.GetItem("steden")

	assert.Equal(t, item1.Name, "steden")
	assert.Equal(t, item1.Age, 99)

	item2, _ := cacheManage.GetItem("steden2")

	assert.Equal(t, item2.Name, "steden2")
	assert.Equal(t, item2.Age, 19)
}

func TestCacheInMemory_Remove(t *testing.T) {
	modules.StartModules(Module{})
	cache.SetProfilesInMemory[po]("test", "Name", 0)
	cacheManage := cache.GetCacheManage[po]("test")
	cacheManage.Set(po{Name: "steden", Age: 18}, po{Name: "steden2", Age: 19})
	cacheManage.Remove("steden")

	_, exists := cacheManage.GetItem("steden")
	assert.False(t, exists)

	item2, _ := cacheManage.GetItem("steden2")
	assert.Equal(t, item2.Name, "steden2")
	assert.Equal(t, item2.Age, 19)
}

func TestCacheInMemory_Clear(t *testing.T) {
	modules.StartModules(Module{})
	cache.SetProfilesInMemory[po]("test", "Name", 0)
	cacheManage := cache.GetCacheManage[po]("test")
	cacheManage.Set(po{Name: "steden", Age: 18}, po{Name: "steden2", Age: 19})
	assert.Equal(t, cacheManage.Count(), 2)
	cacheManage.Clear()
	assert.Equal(t, cacheManage.Count(), 0)
}

func TestCacheInMemory_Exists(t *testing.T) {
	modules.StartModules(Module{})
	cache.SetProfilesInMemory[po]("test", "Name", 0)
	cacheManage := cache.GetCacheManage[po]("test")
	assert.False(t, cacheManage.Exists())
	cacheManage.Set(po{Name: "steden", Age: 18}, po{Name: "steden2", Age: 19})
	assert.True(t, cacheManage.Exists())
}

func TestCacheInMemory_Ttl(t *testing.T) {
	modules.StartModules(Module{})
	cache.SetProfilesInMemory[po]("test", "Name", 10*time.Millisecond)
	cacheManage := cache.GetCacheManage[po]("test")
	lst := collections.NewList(po{Name: "steden", Age: 18}, po{Name: "steden2", Age: 19})
	cacheManage.Set(lst.ToArray()...)
	lst2 := cacheManage.Get()
	assert.Equal(t, lst.Count(), lst2.Count())
	for i := 0; i < lst.Count(); i++ {
		assert.Equal(t, lst.Index(i).Name, lst2.Index(i).Name)
		assert.Equal(t, lst.Index(i).Age, lst2.Index(i).Age)
	}
	time.Sleep(12 * time.Millisecond)
	assert.False(t, cacheManage.Exists())
}
