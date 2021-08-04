package service

import (
	"api-auth/model"
	"api-auth/util"
	"fmt"
	"sync"
)

func ListUser(username string, offset, limit int) ([]*model.UserInfo, int64, error) {
	infos := make([]*model.UserInfo, 0)
	users, count, err := model.ListUser(username, offset, limit)
	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, user := range users {
		ids = append(ids, user.Id)
	}

	wg := sync.WaitGroup{}
	userList := model.UserList{
		// 通过加锁保证更新同一个变量时的数据一致性
		Lock: new(sync.Mutex),
		// 查询的列表通常需要按时间顺序进行排序，一般数据库查询后的列表已经排过序了，但是为了减少延时，程序中用了并发，这时候会打乱排序，所以通过 IdMap 来记录并发处理前的顺序，处理后再重新复位
		IdMap: make(map[uint64]*model.UserInfo, len(users)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// 并发提高查询效率
	for _, u := range users {
		wg.Add(1)
		go func(u *model.UserModel) {
			defer wg.Done()

			// 生成 shordId，问下文模拟数据处理做准备
			shortId, err := util.GenShortId()
			if err != nil {
				errChan <- err
				return
			}

			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			// 通过对每个用户返回一个sqyHello字段模拟数据处理
			userList.IdMap[u.Id] = &model.UserInfo{
				Id:        u.Id,
				Username:  u.Username,
				SayHello:  fmt.Sprintf("Hello %s", shortId),
				Password:  u.Password,
				CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
			}
		}(u)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, count, err
	}
	for _, id := range ids {
		infos = append(infos, userList.IdMap[id])
	}
	return infos, count, nil
}
