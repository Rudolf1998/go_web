package ssql

import (
	"main/model"
	"main/utils"
)

func GetDorms()([]*model.Dorm,error){
	//写sql语句
	sql := "select name, number, amount, remain from dorminfo"
	//执行
	rows, err := utils.Db.Query(sql)
	if err != nil {
		return nil, err
	}
	var dorms []*model.Dorm
	for rows.Next() {
		dorm := &model.Dorm{}
		rows.Scan(&dorm.Name, &dorm.Number, &dorm.Amount, &dorm.Remain)
		dorms = append(dorms, dorm)
	}
	return dorms, nil
}
