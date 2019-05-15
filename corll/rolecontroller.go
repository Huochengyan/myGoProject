package corll

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"myProject/db"
	"net/http"
)

/*   获得角色  */
func GetRoles(g *gin.Context) {
	mgo := db.InitMongoDB()
	rsp := new(Rsp)
	var roles []Role
	cur, err := mgo.Collection(db.Role).Find(context.Background(), bson.D{}, nil)
	if err == nil {
		for cur.Next(context.Background()) {
			elme := new(Role)
			err := cur.Decode(elme)
			if err == nil {
				roles = append(roles, *elme)
			}
		}
		rsp.Msg = "success"
		rsp.Code = 200
		rsp.Data = roles
		g.JSON(http.StatusOK, rsp)
		return
	}
	rsp.Msg = "falid"
	rsp.Code = 201
	g.JSON(http.StatusOK, rsp)
	return
}
