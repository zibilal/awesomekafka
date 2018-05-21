package publisher

type Publisher interface {
	Publish(key interface{}, data interface{}) error
}

