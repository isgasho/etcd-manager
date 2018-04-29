package commands

import (
	"k8s.io/kops/util/pkg/vfs"
	protoetcd "kope.io/etcd-manager/pkg/apis/etcd"
)

const CommandFilename = "_command.json"

type Store interface {
	// AddCommand adds a command to the back of the queue
	AddCommand(cmd *protoetcd.Command) error

	// ListCommands returns all the external commands that have not been removed
	ListCommands() ([]Command, error)

	// RemoveCommand marks a command as complete
	RemoveCommand(command Command) error
}

type Command interface {
	Data() protoetcd.Command
}

func NewStore(storage string) (Store, error) {
	p, err := vfs.Context.BuildVfsPath(storage)
	if err != nil {
		return nil, err
	}
	return NewVFSStore(p)
}