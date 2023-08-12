package database_wiiu

import (
	"context"
	"encoding/base64"

	"github.com/PretendoNetwork/friends-secure/database"
	"github.com/PretendoNetwork/friends-secure/globals"
	"github.com/PretendoNetwork/nex-go"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUserInfoByPID(pid uint32) *friends_wiiu_types.PrincipalBasicInfo {
	var result bson.M

	info := friends_wiiu_types.NewPrincipalBasicInfo()

	err := database.MongoCollection.FindOne(context.TODO(), bson.D{{Key: "pid", Value: pid}}, options.FindOne()).Decode(&result)

	if err != nil {
		if err != mongo.ErrNoDocuments {
			globals.Logger.Critical(err.Error())
		}

		return info
	}

	info.PID = pid
	info.NNID = result["username"].(string)
	info.Mii = friends_wiiu_types.NewMiiV2()
	info.Unknown = 2

	encodedMiiData := result["mii"].(bson.M)["data"].(string)
	decodedMiiData, _ := base64.StdEncoding.DecodeString(encodedMiiData)

	info.Mii.Name = result["mii"].(bson.M)["name"].(string)
	info.Mii.Unknown1 = 0
	info.Mii.Unknown2 = 0
	info.Mii.MiiData = decodedMiiData
	info.Mii.Datetime = nex.NewDateTime(0)

	return info
}
