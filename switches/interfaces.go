package switches

import "github.com/house-lee/SmartDevices/dev"

type Switch interface {
	dev.Action
	dev.Info

}

