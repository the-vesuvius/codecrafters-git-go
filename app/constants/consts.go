package constants

import "path"

const (
	GitDirName = ".git"

	ObjectsDirName   = "objects"
	RefsDirName      = "refs"
	HEADFileName     = "HEAD"
	RefsHeadsDirName = "heads"
)

var (
	GitDirPath     = GitDirName
	ObjectsDirPath = path.Join(GitDirName, ObjectsDirName)
	RefsDirPath    = path.Join(GitDirName, RefsDirName)
	HEADFilePath   = path.Join(GitDirName, HEADFileName)
)
