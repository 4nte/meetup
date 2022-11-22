package api

import (
	"context"
	"fmt"

	coffe "github.com/4nte/meetup/coffee/proto"
	"github.com/bufbuild/connect-go"
)

type CoffeServer struct {
	sugar int32
	milk  int32
}

func (s *CoffeServer) MakeCoffe(
	ctx context.Context,
	req *connect.Request[coffe.MakeCoffeRequest],
) (*connect.Response[coffe.MakeCoffeResponse], error) {
	fmt.Printf("MakeCoffe request received")

	s.sugar -= req.Msg.GramsOfSugar

	// Check if enough sugar left
	if req.Msg.GramsOfSugar < 0 {
		return nil, connect.NewError(connect.CodeResourceExhausted, fmt.Errorf("no sugar left"))
	}

	res := connect.NewResponse(&coffe.MakeCoffeResponse{})

	return res, nil
}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
// func (s *CoffeServer) GetState(
// 	ctx context.Context,
// 	req *connect.Request[coffe.GetStateRequest],
// ) (*connect.Response[coffe.GetStateResponse], error) {
// 	fmt.Printf("GetState request received")

// 	res := connect.NewResponse(&coffe.GetStateResponse{
// 		RemainingSugar: s.sugar,
// 		RemainingMilk:  s.milk,
// 	})

// 	return res, nil
// }
