package builder

/**
产生更新代码
更新代码格式
*/

//1. 收集包含特定scope的字段集合
//2. 数据集合做下面判断
//3. 在编码过冲中，将数据做如下判断
/*func UserUpdateDTO(value map[string]interface{}, scope int) (updater interface{}, err error) {
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
*/
