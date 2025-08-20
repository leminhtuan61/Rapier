package ds

import "time"

type Object struct {
	Value interface{}
}

type Dictionary struct {
	objects map[string]*Object
	expired map[string]uint64
}

func NewDictionary() *Dictionary {
	return &Dictionary{
		objects: make(map[string]*Object),
		expired: make(map[string]uint64),
	}
}
func (d *Dictionary) Set(key string, value interface{}, expire uint64) {
	d.objects[key] = &Object{Value: value}
	if expire > 0 {
		d.expired[key] = expire
	} else {
		delete(d.expired, key)
	}
}
func (d *Dictionary) Get(key string) (interface{}, bool) {
	if obj, exists := d.objects[key]; exists {
		if d.HasExpired(key) {
			d.Delete(key) // Clean up expired key
			return nil, false
		}
		return obj.Value, true
	}
	return nil, false
}

func (d *Dictionary) Delete(key string) {
	if _, exists := d.objects[key]; exists {
		delete(d.objects, key)
		delete(d.expired, key)
	}
}
func (d *Dictionary) GetExpireDict() map[string]uint64 {
	return d.expired

}

func (d *Dictionary) SetExpireDict() map[string]uint64 {
	return d.expired
}

func (d *Dictionary) HasExpired(key string) bool {
	expired, exists := d.expired[key]
	if !exists {
		return false
	}
	return expired <= uint64(time.Now().UnixMilli())
}
