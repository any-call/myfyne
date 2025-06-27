package mybinding

import (
	"fmt"
	"reflect"
	"sync"

	"fyne.io/fyne/v2/data/binding"
)

// BindStructEx 泛型版，持有原始值
type BindStruct[T any] struct {
	value     T
	items     map[string]binding.DataItem
	computeds map[string]func(T) any
	mu        sync.RWMutex
}

// NewBindStructEx 初始化泛型版
func NewBindStructEx[T any](input T) *BindStruct[T] {
	bs := &BindStruct[T]{
		value:     input,
		items:     make(map[string]binding.DataItem),
		computeds: make(map[string]func(T) any),
	}
	bs.extractStruct("", reflect.ValueOf(input))
	return bs
}

// 新增方法：添加计算字段
func (b *BindStruct[T]) AddComputedField(name string, compute func(T) any) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.computeds == nil {
		b.computeds = make(map[string]func(T) interface{})
	}
	b.computeds[name] = compute

	// 初始化对应 DataItem
	switch val := compute(b.value).(type) {
	case int:
		it := binding.NewInt()
		it.Set(val)
		b.items[name] = it
	case int32, int64:
		it := binding.NewInt()
		it.Set(int(reflect.ValueOf(val).Int()))
		b.items[name] = it
	case float32, float64:
		it := binding.NewFloat()
		it.Set(reflect.ValueOf(val).Float())
		b.items[name] = it
	case string:
		it := binding.NewString()
		it.Set(val)
		b.items[name] = it
	// 更多类型根据需要扩展
	default:
		it := binding.NewString()
		it.Set(fmt.Sprintf("%v", val))
		b.items[name] = it
	}
}

// Get 实现 binding.Struct 接口
func (b *BindStruct[T]) GetItem(path string) (binding.DataItem, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	item, ok := b.items[path]
	if !ok {
		return nil, fmt.Errorf("no such path: %s", path)
	}
	return item, nil
}

func (b *BindStruct[T]) GetOrCreateItem(path string, createItem binding.DataItem) binding.DataItem {
	b.mu.RLock()
	defer b.mu.RUnlock()
	item, ok := b.items[path]
	if !ok {
		b.items[path] = createItem
		return createItem
	}

	return item
}

// 当前值
func (b *BindStruct[T]) Value() T {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.value
}

// SetValue 设置新值并更新绑定数据
func (b *BindStruct[T]) SetValue(newVal T) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.value = newVal
	b.updateStruct("", reflect.ValueOf(newVal))

	// 更新计算字段
	for path, compute := range b.computeds {
		result := compute(newVal)
		if item, ok := b.items[path]; ok {
			switch data := item.(type) {
			case binding.Int:
				data.Set(int(reflect.ValueOf(result).Int()))
			case binding.Float:
				data.Set(reflect.ValueOf(result).Float())
			case binding.String:
				data.Set(fmt.Sprintf("%v", result))
			}
		}
	}
}

// 递归绑定字段
func (b *BindStruct[T]) extractStruct(prefix string, v reflect.Value) {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		fieldVal := v.Field(i)
		path := field.Name
		if prefix != "" {
			path = prefix + "." + field.Name
		}
		switch fieldVal.Kind() {
		case reflect.Struct:
			b.extractStruct(path, fieldVal)
			break
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			it := binding.NewInt()
			it.Set(int(fieldVal.Int()))
			b.items[path] = it
			break
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			it := binding.NewInt()
			it.Set(int(fieldVal.Uint()))
			b.items[path] = it
			break
		case reflect.Float32, reflect.Float64:
			it := binding.NewFloat()
			it.Set(fieldVal.Float())
			b.items[path] = it
			break
		case reflect.String:
			it := binding.NewString()
			it.Set(fieldVal.String())
			b.items[path] = it
			break
		case reflect.Bool:
			it := binding.NewBool()
			it.Set(fieldVal.Bool())
			b.items[path] = it
			break
		default:
			// 忽略
		}
	}
}

// 递归更新已有字段
func (b *BindStruct[T]) updateStruct(prefix string, v reflect.Value) {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		fieldVal := v.Field(i)
		path := field.Name
		if prefix != "" {
			path = prefix + "." + field.Name
		}
		if fieldVal.Kind() == reflect.Struct {
			b.updateStruct(path, fieldVal)
		} else if item, ok := b.items[path]; ok {
			switch data := item.(type) {
			case binding.Int:
				data.Set(int(fieldVal.Int()))
			case binding.Float:
				data.Set(fieldVal.Float())
			case binding.String:
				data.Set(fieldVal.String())
			case binding.Bool:
				data.Set(fieldVal.Bool())
			}
		}
	}
}
