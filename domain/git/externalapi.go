package git

import (
	"errors"
)

var (
	ErrRepositoryNotExists = errors.New("repository does not exist")
)

type GitCommand interface {
	Clone(path Path, remoteUrl string, targetBranch string, recursive bool) error
	Fetch(workingDir Path) error
	DiffExists(workingDir Path, to string) (exists bool, err error)
	Pull(workingDir Path, force bool) (refCommitId string, err error)
	IsGitDirectory(workingDir Path) (bool, error)
	RefCommitId(workingDir Path) (string, error)
	RefBranchName(workingDir Path) (string, error)
	FindHashByTagRegex(workingDir Path, regex string) (hash string, err error)
	CheckoutBranch(workingDir Path, branch string) error
	CheckoutHash(workingDir Path, hash string) error
	GetRemoteUrl(workingDir Path, remoteName string) (string, error)
	SetRemoteUrl(workingDir Path, remoteName string, url string) error
}
