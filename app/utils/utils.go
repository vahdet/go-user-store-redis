package utils

import (
	"github.com/vahdet/go-user-store-redis/models"
	pb "github.com/vahdet/go-user-store-redis/proto"
	"github.com/golang/protobuf/ptypes"
	log "github.com/sirupsen/logrus"
	"fmt"
)

func ConvertModelToProto(user *models.User) (*pb.User, error) {

	createdProtoTimestamp, err := ptypes.TimestampProto(user.Created)
	lastChangedProtoTimestamp, err := ptypes.TimestampProto(user.LastChanged)
	if err != nil {
		log.WithFields(log.Fields{
			"createdTime": user.Created,
			"lastChanged": user.LastChanged,
		}).Error(fmt.Sprintf("Time conversion failed: '%#v'", err))
		return nil, err
	}

	return &pb.User{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		Location:    user.Location,
		Language:    user.Language,
		Created:     createdProtoTimestamp,
		LastChanged: lastChangedProtoTimestamp,
	}, nil
}

func ConvertProtoToModel(proto *pb.User) (*models.User, error) {

	createdTimestamp, err := ptypes.Timestamp(proto.Created)
	lastChangedTimestamp, err := ptypes.Timestamp(proto.LastChanged)
	if err != nil {
		log.WithFields(log.Fields{
			"createdTime": proto.Created,
			"lastChanged": proto.LastChanged,
		}).Warn(fmt.Sprintf("Time conversion failed: '%#v'", err))
		return nil, err
	}

	return &models.User{
		Id:          proto.Id,
		Name:        proto.Name,
		Email:       proto.Email,
		Location:    proto.Location,
		Language:    proto.Language,
		Created:     createdTimestamp,
		LastChanged: lastChangedTimestamp,
	}, nil
}

func ConvertProtoCreateReqToModel(proto *pb.CreateRequest) (*models.User, error) {
	return &models.User{
		Name:        proto.Name,
		Email:       proto.Email,
		Location:    proto.Location,
		Language:    proto.Language,
	}, nil
}

func ConvertProtoUpdateReqToModel(proto *pb.UpdateRequest) (*models.User, error) {
	return &models.User{
		Id:          proto.Id,
		Name:        proto.Name,
		Email:       proto.Email,
		Location:    proto.Location,
		Language:    proto.Language,
	}, nil
}
