package service

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models/errors"
)

const (
	SID         = "SID"
	PrefixUser  = "user"
	PrefixAdmin = "admin"
)

type Session struct {
	UserID     int `redis:"user_id"`
	DeviceType int `redis:"device_type"`
}

func SetOutline(sessionID string, prefix string, uid int) error {
	cli := common.RedisPool.Get()
	defer cli.Close()

	sessionKey := fmt.Sprintf("%s_session_%s", prefix, sessionID)
	metaKey := fmt.Sprintf("%s_meta_%d", prefix, uid)

	cli.Send("MULTI")
	cli.Send("DEL", sessionKey)
	cli.Send("SREM", metaKey, sessionKey)
	_, err := cli.Do("EXEC")
	if err != nil {
		return err
	}
	return err
}

func GetSessionInfo(sessionID string, prefix string) (*Session, error) {
	cli := common.RedisPool.Get()
	defer cli.Close()

	var sObj = Session{}

	sessionKey := fmt.Sprintf("%s_session_%s", prefix, sessionID)

	exists, err := redis.Bool(cli.Do("EXISTS", sessionKey))
	if err != nil {
		return &sObj, err
	}
	if exists == false {
		return &sObj, errors.New("已经不存在")
	}

	obj, err := redis.Values(cli.Do("HGETALL", sessionKey))
	if err != nil {
		return &sObj, err
	}
	err = redis.ScanStruct(obj, &sObj)

	return &sObj, err
}

func SetOnline(sessionID string, prefix string, uid int, deviceType int) error {
	cli := common.RedisPool.Get()
	defer cli.Close()

	sessionKey := fmt.Sprintf("%s_session_%s", prefix, sessionID)
	metaKey := fmt.Sprintf("%s_meta_%d", prefix, uid)

	cli.Send("MULTI")
	cli.Send("SADD", metaKey, sessionKey)
	cli.Send("HMSET", sessionKey, "user_id", uid, "device_type", deviceType)
	cli.Send("EXPIRE", sessionKey, 60*60*24*15)
	_, err := cli.Do("EXEC")
	if err != nil {
		return err
	}
	SessionKeyList, err := redis.Strings(cli.Do("SMEMBERS", metaKey))
	if err != nil {
		return err
	}

	for i := range SessionKeyList {
		sKey := SessionKeyList[i]

		if sKey == sessionKey {
			continue
		}

		obj, err := redis.Values(cli.Do("HGETALL", sKey))
		if err != nil {
			return err
		}

		var sObj = Session{}
		if err := redis.ScanStruct(obj, &sObj); err != nil {
			return err
		}

		if sObj.DeviceType == deviceType {
			cli.Send("MULTI")
			cli.Send("DEL", sKey)
			cli.Send("SREM", metaKey, sKey)
			_, err := cli.Do("EXEC")
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func SetAdminOnline(sessionID string, uid int, deviceType int) error {
	return SetOnline(sessionID, PrefixAdmin, uid, deviceType)
}

func SetUserOnline(sessionID string, uid int, deviceType int) error {
	return SetOnline(sessionID, PrefixUser, uid, deviceType)
}

func GetAdminSessionInfo(sessionID string) (*Session, error) {
	return GetSessionInfo(sessionID, PrefixAdmin)
}

func GetUserSessionInfo(sessionID string) (*Session, error) {
	return GetSessionInfo(sessionID, PrefixUser)
}
