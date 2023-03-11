package store

func (s *RedisStore) AddByName(name string) (int64, error) {
	return s.rdb.Incr(s.rctx, name).Result()
}
