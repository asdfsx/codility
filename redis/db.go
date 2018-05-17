package redis1

import(
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var (
	NETWORK string
	ADDR string
	PORT int
)

func Connect() (redis.Conn, error) {
	connstr := fmt.Sprintf("%s:%d", ADDR, PORT)
	return redis.Dial(NETWORK, connstr)
}

func Set(conn redis.Conn, key string, value string) error {
	reply, err := conn.Do("SET", key, value)
	if err != nil {
		fmt.Println("redis set failed:", err)
		return err
	}
	fmt.Println(reply)
	return nil
}

func Get(conn redis.Conn, key string) (string, error) {
	result, err := redis.String(conn.Do("GET", key))
	if err != nil {
		fmt.Println("redis get failed:", err)
		return "", err
	}
	return result, nil
}

func Del(conn redis.Conn, key string) error {
	reply, err := conn.Do("DEL", key)
	if err != nil {
		fmt.Println("redis delelte failed:", err)
		return err
	}
	fmt.Println(reply)
	return nil
}

func TTL(conn redis.Conn, key string, value string, ttl int) error {
	reply, err := conn.Do("SET", key, value, "EX", ttl)
	if err != nil {
		fmt.Println("redis set failed:", err)
		return err
	}
	fmt.Println(reply)
	return nil
}

func Exists(conn redis.Conn, key string) (bool, error){
	is_exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		fmt.Println("redis exists failed:", err)
		return false, err
	}
	return is_exists, nil
}

func Lpush(conn redis.Conn, key string, value string) error {
	reply, err := conn.Do("LPUSH", key, value)
	if err != nil {
		fmt.Println("redis exists failed:", err)
		return err
	}
	fmt.Println(reply)
	return nil
}

func Llen(conn redis.Conn, key string) (length int, err error) {
    length, err = redis.Int(conn.Do("LLEN", key))
	if err != nil {
		fmt.Println("redis exists failed:", err)
	}
	return
}

func Lpop(conn redis.Conn, key string)(result string , err error){
    result, err = redis.String(conn.Do("LPOP", key))
	if err != nil {
		fmt.Println("redis exists failed:", err)
	}
	return
}