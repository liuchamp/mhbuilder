package builder

import (
	"errors"
	"github.com/wxnacy/wgo/arrays"
	"go.mongodb.org/mongo-driver/bson"
)

/**
产生更新代码
更新代码格式
*/

func UserUpdateDTO(value map[string]interface{}, scope int) (updater interface{}, err error) {
	if value == nil || len(value) == 0 {
		return nil, errors.New("value nil")
	}
	sArrayxxx := []string{}
	up := bson.M{}
	for k, v := range value {
		if k == xxx && scope < xxx && arrays.ContainsString(sArrayxxx, k) != -1 {
			up[k_b] = v
		}
	}
	return bson.M{"$set": up}, nil
}
