package db

type RedisMI struct {
}

var Redis = RedisMI{}

func (r RedisMI) Set(key string, value string) error {
	return nil
}

func (r RedisMI) Get(key string) (string, error) {
	return "", nil
}

func (r RedisMI) Delete(key string) error {
	return nil
}
