package stat_repository_postgres

import core_pgx_pool "github.com/Mihail4531/golang-todo/internal/core/repository/postgres/pool"

type StatRespository struct {
	pool core_pgx_pool.Pool
}

func NewStatRepository(pool core_pgx_pool.Pool) *StatRespository {
	return &StatRespository{
		pool: pool,
	}
}
