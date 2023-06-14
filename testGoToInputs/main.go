package main

import (
	"context"

	"github.com/edaniels/golog"
	"go.viam.com/rdk/components/base"
	"go.viam.com/rdk/components/board"
	"go.viam.com/rdk/components/camera"
	"go.viam.com/rdk/components/encoder"
	"go.viam.com/rdk/components/motor"
	"go.viam.com/rdk/robot/client"
	"go.viam.com/rdk/utils"
	"go.viam.com/utils/rpc"
)

func main() {
	logger := golog.NewDevelopmentLogger("client")
	robot, err := client.New(
		context.Background(),
		"catmotion-main.sjkw3nonb9.viam.cloud",
		logger,
		client.WithDialOptions(rpc.WithCredentials(rpc.Credentials{
			Type:    utils.CredentialsTypeRobotLocationSecret,
			Payload: "fvrf51544m5wzivpms42vq405ry5oex1yz80iiz88dzk6xyl",
		})),
	)
	if err != nil {
		logger.Fatal(err)
	}

	sampleFunctionality(robot, logger)

	// viam_base
	viamBaseComponent, err := base.FromRobot(robot, "viam_base")
	if err != nil {
		logger.Error(err)
	}
	viamBaseReturnValue, err := viamBaseComponent.IsMoving(context.Background())
	if err != nil {
		logger.Error(err)
	}
	logger.Debugf("viam_base IsMoving return value: %v", viamBaseReturnValue)

	kab, err := NewKAB(viamBaseComponent, localizer, model)
	kab.GoToInputs(context.Background(), [100, 90])
}

func sampleFunctionality(robot *client.RobotClient, logger golog.Logger) {

	// rplidar
	rplidarComponent, err := camera.FromRobot(robot, "rplidar")
	if err != nil {
		logger.Error(err)
	}
	rplidarReturnValue, err := rplidarComponent.Properties(context.Background())
	if err != nil {
		logger.Error(err)
	}
	logger.Debugf("rplidar Properties return value: %v", rplidarReturnValue)

	// Note that the pin supplied is a placeholder. Please change this to a valid pin.
	// local
	localComponent, err := board.FromRobot(robot, "local")
	if err != nil {
		logger.Error(err)
	}
	localReturnValue, err := localComponent.GPIOPinByName("16")
	if err != nil {
		logger.Error(err)
	}
	logger.Debugf("local GPIOPinByName return value: %v", localReturnValue)

	// right
	rightComponent, err := motor.FromRobot(robot, "right")
	if err != nil {
		logger.Error(err)
	}
	rightReturnValue, err := rightComponent.IsMoving(context.Background())
	if err != nil {
		logger.Error(err)
	}
	logger.Debugf("right IsMoving return value: %v", rightReturnValue)

	// left
	leftComponent, err := motor.FromRobot(robot, "left")
	if err != nil {
		logger.Error(err)
	}
	leftReturnValue, err := leftComponent.IsMoving(context.Background())
	if err != nil {
		logger.Error(err)
	}
	logger.Debugf("left IsMoving return value: %v", leftReturnValue)

	// viam_base
	viamBaseComponent, err := base.FromRobot(robot, "viam_base")
	if err != nil {
		logger.Error(err)
	}
	viamBaseReturnValue, err := viamBaseComponent.IsMoving(context.Background())
	if err != nil {
		logger.Error(err)
	}
	logger.Debugf("viam_base IsMoving return value: %v", viamBaseReturnValue)

	// Renc
	rencComponent, err := encoder.FromRobot(robot, "Renc")
	if err != nil {
		logger.Error(err)
	}
	rencReturnValue, err := rencComponent.Properties(context.Background(), map[string]interface{}{})
	if err != nil {
		logger.Error(err)
	}
	logger.Debugf("Renc Properties return value: %v", rencReturnValue)

	// Lenc
	lencComponent, err := encoder.FromRobot(robot, "Lenc")
	if err != nil {
		logger.Error(err)
	}
	lencReturnValue, err := lencComponent.Properties(context.Background(), map[string]interface{}{})
	if err != nil {
		logger.Error(err)
	}
	logger.Debugf("Lenc Properties return value: %v", lencReturnValue)

	defer robot.Close(context.Background())
	logger.Info("Resources:")
	logger.Info(robot.ResourceNames())
}
