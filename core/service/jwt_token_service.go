package coreservice

import "daijoubuteam.xyz/se214-library-management/core/entity"

type JwtTokenService interface {
	Encode(sub *entity.ID) (string, error)
	Decode(token string) (*entity.ID, error)
}
