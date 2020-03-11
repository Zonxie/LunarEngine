package dao

import (
	"LunarAssignment/LunarEngine/model"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (dao *Dao) GetAccount(accountID int) (*model.Account, error) {

	var resultFind *model.Account

	client, err := dao.client("accounts")
	if err != nil {
		return resultFind, err
	}

	defer client.Disconnect(client.context)

	res := client.collection.FindOne(client.context, bson.M{"accountID": accountID})
	if res.Err() != nil {
		return nil, res.Err()
	}

	res.Decode(&resultFind)

	return resultFind, nil

}

func (dao *Dao) UpdateAccount(cr *model.Account) (string, error) {

	client, err := dao.client("accounts")
	if err != nil {
		return "", err
	}

	defer client.Disconnect(client.context)

	var resultFind model.Account
	var resultUpdate model.Account

	res := client.collection.FindOne(client.context, bson.M{"accountID": cr.AccountID})
	if err != nil {
		return "", err
	}

	if res.Err() != nil {

	}
	res.Decode(&resultFind)

	filter := bson.M{"accountID": cr.AccountID}
	update := bson.M{"$set": bson.M{"value": cr.Value}}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	update = bson.M{"$set": bson.M{"value": resultFind.Value + cr.Value}}
	_ = client.collection.FindOneAndUpdate(client.context, filter, update, &opt).Decode(&resultUpdate)

	if err != nil {
		return "", err
	}

	if res.Err() != nil {
		return fmt.Sprintf("Created acccount %v with balance %v", cr.AccountID, cr.Value), nil
	}

	return fmt.Sprintf("Updated balance of acccount %v from %v to %v", cr.AccountID, resultFind.Value, resultUpdate.Value), nil
}
