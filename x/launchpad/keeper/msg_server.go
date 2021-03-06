package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/nghuyenthevinh2000/nebula/x/launchpad/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (server msgServer) CreateProject(goCtx context.Context, msg *types.MsgCreateProjectRequest) (*types.MsgCreateProjectResponse, error) {
	// get ctx SDK context
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get project owner
	project_owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	// invoke logic CreateProject
	project_id, err := server.Keeper.CreateProject(ctx, project_owner, msg)
	if err != nil {
		return nil, err
	}

	// emit event
	ctx.EventManager().EmitEvents(sdk.Events{
		// an event to signify a project created
		sdk.NewEvent(
			types.TypeProjectCreated,
			sdk.NewAttribute(types.AttributeProjectID, strconv.FormatUint(project_id, 10)),
		),
		// an event to signify the event comes from which module and which signer
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner),
		),
	})

	// return result to gRPC server
	return &types.MsgCreateProjectResponse{
		ProjectId: project_id,
	}, nil
}

func (server msgServer) DeleteProject(goCtx context.Context, msg *types.MsgDeleteProjectRequest) (*types.MsgDeleteProjectResponse, error) {
	// get ctx SDK context
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if owner address is valid
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	// invoke logic CreateProject
	project_id, err := server.Keeper.DeleteProject(ctx, msg)
	if err != nil {
		return nil, err
	}

	// emit event
	ctx.EventManager().EmitEvents(sdk.Events{
		// an event to signify a project created
		sdk.NewEvent(
			types.TypeProjectModified,
			sdk.NewAttribute(types.AttributeProjectID, strconv.FormatUint(project_id, 10)),
		),
		// an event to signify the event comes from which module and which signer
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner),
		),
	})

	return &types.MsgDeleteProjectResponse{}, nil
}

func (server msgServer) ModifyStartTime(goCtx context.Context, msg *types.MsgModifyStartTimeRequest) (*types.MsgModifyStartTimeResponse, error) {
	// get ctx SDK context
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if owner address is valid
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	// invoke logic CreateProject
	project_id, err := server.Keeper.ModifyStartTime(ctx, msg)
	if err != nil {
		return nil, err
	}

	// emit event
	ctx.EventManager().EmitEvents(sdk.Events{
		// an event to signify a project created
		sdk.NewEvent(
			types.TypeProjectModified,
			sdk.NewAttribute(types.AttributeProjectID, strconv.FormatUint(project_id, 10)),
		),
		// an event to signify the event comes from which module and which signer
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner),
		),
	})

	return &types.MsgModifyStartTimeResponse{}, nil
}

func (server msgServer) ModifyProjectInformation(goCtx context.Context, msg *types.MsgModifyProjectInformationRequest) (*types.MsgModifyProjectInformationResponse, error) {
	// get ctx SDK context
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if owner address is valid
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	// invoke logic CreateProject
	project_id, err := server.Keeper.ModifyProjectInformation(ctx, msg)
	if err != nil {
		return nil, err
	}

	// emit event
	ctx.EventManager().EmitEvents(sdk.Events{
		// an event to signify a project created
		sdk.NewEvent(
			types.TypeProjectModified,
			sdk.NewAttribute(types.AttributeProjectID, strconv.FormatUint(project_id, 10)),
		),
		// an event to signify the event comes from which module and which signer
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner),
		),
	})

	return &types.MsgModifyProjectInformationResponse{}, nil
}
