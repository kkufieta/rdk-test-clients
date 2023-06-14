package main

import (
	"context"
	"math"
	"time"

	"github.com/golang/geo/r3"
	"go.viam.com/rdk/components/base"
	"go.viam.com/rdk/referenceframe"
	"go.viam.com/rdk/services/motion"
)

const (
	maxLinearVelocity = 100.0 // mm per second
)

type kinematicAckermanBase struct {
	base.Base
	localizer motion.Localizer
	model     referenceframe.Model
}

func NewKAB(b base.Base, localizer motion.Localizer, model referenceframe.Model) (base.KinematicBase, error) {
	return &kinematicAckermanBase{
		Base:      b,
		localizer: localizer,
		model:     model,
	}, nil
}

func (kab *kinematicAckermanBase) ModelFrame() referenceframe.Model {
	return kab.model
}

// desiredState = [radius in mm, angle]
func (kab *kinematicAckermanBase) GoToInputs(ctx context.Context, desiredState []referenceframe.Input) error {
	// implement GoToInputs
	// use SetVelocity here!

	// We assume we're moving in circular motions for now
	radius := desiredState[0].Value
	angle := desiredState[1].Value

	linVel := maxLinearVelocity
	angVel := linVel / radius
	dtSec := angle * radius / linVel

	if err := kab.SetVelocity(ctx, r3.Vector{0, linVel, 0}, r3.Vector{0, 0, angVel}, nil); err != nil {
		return err
	}

	time.Sleep(time.Millisecond * time.Duration(dtSec))

	// TODO: Try to return here to see if base stops moving
	// TODO: Use Stop?
	if err := kab.SetVelocity(ctx, r3.Vector{}, r3.Vector{}, nil); err != nil {
		return err
	}

	return nil
}

func (kab *kinematicAckermanBase) CurrentInputs(ctx context.Context) ([]referenceframe.Input, error) {
	// TODO(rb): make a transformation from the component reference to the base frame
	poseInFrame, err := kab.localizer.CurrentPosition(ctx)
	if err != nil {
		return nil, err
	}
	pose := poseInFrame.Pose()
	pt := pose.Point()
	// TODO: Understand this better
	theta := math.Mod(pose.Orientation().OrientationVectorRadians().Theta, 2*math.Pi) - math.Pi
	return []referenceframe.Input{{Value: pt.X}, {Value: pt.Y}, {Value: theta}}, nil
}
