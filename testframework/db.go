package testframework

type Repository interface {
	Create(key string, value []byte) error
	Retrieve(key string) ([]byte, error)
	Update(key string, value []byte) error
	Delete(key string) error
}

var repo Repository = newRedisRepo()

type redisrepo struct{

}

func newRedisRepo() Repository {
	return redisrepo{}
}

func(repo redisrepo) Create(key string, value []byte) error{
	return nil
}

func(repo redisrepo) Retrieve(key string) ([]byte, error) {
	return nil, nil
}

func(repo redisrepo) Update(key string, value []byte) error {
	return nil
}

func(repo redisrepo) Delete(key string) error {
	return nil
}