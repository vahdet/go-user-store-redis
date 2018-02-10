package grpcserver

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	pb "github.com/vahdet/tafalk-user-store/proto"
	"golang.org/x/net/context"

	"github.com/vahdet/tafalk-user-store/services/interfaces"
	"github.com/vahdet/tafalk-user-store/app/utils"
)

type UserServer struct {
	Service interfaces.UserService
}

func (s *UserServer) Get(ctx context.Context, in *pb.UserId) (*pb.User, error) {
	res, err := s.Service.Get(in.Value)
	if err != nil {
		log.WithFields(log.Fields{
			"id": in.Value,
		}).Error(fmt.Sprintf("Getting failed: '%#v'", err))
		return nil, err
	}

	return utils.ConvertModelToProto(res)
}

func (s *UserServer) Create(ctx context.Context, in *pb.CreateRequest) (*pb.UserId, error) {
	user, err := utils.ConvertProtoCreateReqToModel(in)
	if err != nil {
		log.WithFields(log.Fields{
			"name": in.Name,
			"email": in.Email,
		}).Error(fmt.Sprintf("Conversion failed: '%#v'", err))
		return nil, err
	}

	res, err := s.Service.Create(user)
	if err != nil {
		log.WithFields(log.Fields{
			"name": in.Name,
			"email": in.Email,
		}).Error(fmt.Sprintf("Creation failed: '%#v'", err))
		return nil, err
	}

	return &pb.UserId{
		Value: res.Id,
	}, nil
}

func (s *UserServer) Update(ctx context.Context, in *pb.UpdateRequest) (*pb.UserId, error) {
	user, err := utils.ConvertProtoUpdateReqToModel(in)
	if err != nil {
		log.WithFields(log.Fields{
			"name": in.Name,
			"email": in.Email,
		}).Error(fmt.Sprintf("Conversion failed: '%#v'", err))
		return nil, err
	}

	res, err := s.Service.Update(user.Id, user)
	if err != nil {
		log.WithFields(log.Fields{
			"id": in.Id,
		}).Error(fmt.Sprintf("Update failed: '%#v'", err))
		return nil, err
	}

	return &pb.UserId{
		Value: res.Id,
	}, nil
}

func (s *UserServer) Delete(ctx context.Context, in *pb.UserId) (*pb.UserId, error) {
	res, err := s.Service.Delete(in.Value)
	if err != nil {
		log.WithFields(log.Fields{
			"id": in.Value,
		}).Error(fmt.Sprintf("Deletion failed: '%#v'", err))
		return nil, err
	}

	return &pb.UserId{
		Value: res.Id,
	}, nil
}
