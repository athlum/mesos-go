package calls

import (
	"github.com/mesos/mesos-go"
	"github.com/mesos/mesos-go/agent"
)

func LaunchNestedContainer(cid mesos.ContainerID, cmd *mesos.CommandInfo, container *mesos.ContainerInfo) *agent.Call {
	return &agent.Call{
		Type: agent.Call_LAUNCH_NESTED_CONTAINER.Enum(),
		LaunchNestedContainer: &agent.Call_LaunchNestedContainer{
			ContainerID: cid,
			Command:     cmd,
			Container:   container,
		},
	}
}

func WaitNestedContainer(cid mesos.ContainerID) *agent.Call {
	return &agent.Call{
		Type: agent.Call_WAIT_NESTED_CONTAINER.Enum(),
		WaitNestedContainer: &agent.Call_WaitNestedContainer{
			ContainerID: cid,
		},
	}
}

func KillNestedContainer(cid mesos.ContainerID) *agent.Call {
	return &agent.Call{
		Type: agent.Call_KILL_NESTED_CONTAINER.Enum(),
		KillNestedContainer: &agent.Call_KillNestedContainer{
			ContainerID: cid,
		},
	}
}

func LaunchNestedContainerSession(cid mesos.ContainerID, cmd *mesos.CommandInfo, container *mesos.ContainerInfo) *agent.Call {
	return &agent.Call{
		Type: agent.Call_LAUNCH_NESTED_CONTAINER_SESSION.Enum(),
		LaunchNestedContainerSession: &agent.Call_LaunchNestedContainerSession{
			ContainerID: cid,
			Command:     cmd,
			Container:   container,
		},
	}
}

func AttachContainerInput(ty agent.Call_AttachContainerInput_Type, containerId *mesos.ContainerID, pio *agent.ProcessIO) *agent.Call {
	return &agent.Call{
		Type: agent.Call_ATTACH_CONTAINER_INPUT.Enum(),
		AttachContainerInput: &agent.Call_AttachContainerInput{
			Type:        ty.Enum(),
			ContainerID: containerId,
			ProcessIo:   pio,
		},
	}
}

func AttachContainerOutput(containerId mesos.ContainerID) *agent.Call {
	return &agent.Call{
		Type: agent.Call_ATTACH_CONTAINER_OUTPUT.Enum(),
		AttachContainerOutput: &agent.Call_AttachContainerOutput{
			ContainerID: containerId,
		},
	}
}
