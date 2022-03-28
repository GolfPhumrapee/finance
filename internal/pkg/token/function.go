package token

import "time"

func (s *service) getAccessExpireAt() time.Duration {
	return (time.Hour * (s.config.JWT.Access.ExpireTime.Day * 24)) +
		(time.Hour * s.config.JWT.Access.ExpireTime.Hour) +
		(time.Minute * s.config.JWT.Access.ExpireTime.Minute)
}

func (s *service) getRefreshExpireAt() time.Duration {
	return (time.Hour * (s.config.JWT.Refresh.ExpireTime.Day * 24)) +
		(time.Hour * s.config.JWT.Refresh.ExpireTime.Hour) +
		(time.Minute * s.config.JWT.Refresh.ExpireTime.Minute)
}
