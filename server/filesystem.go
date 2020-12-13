package server

import (
	"github.com/pterodactyl/wings/server/filesystem"
	"os"
)

func (s *Server) Filesystem() *filesystem.Filesystem {
	return s.fs
}

// Ensures that the data directory for the server instance exists.
func (s *Server) EnsureDataDirectoryExists() error {
	if _, err := os.Stat(s.fs.Path()); err != nil && !os.IsNotExist(err) {
		return err
	} else if err != nil {
		// Create the server data directory because it does not currently exist
		// on the system.
		if err := os.MkdirAll(s.fs.Path(), 0700); err != nil {
			return err
		}

		if err := s.fs.Chown("/"); err != nil {
			s.Log().WithField("error", err).Warn("failed to chown server data directory")
		}
	}

	return nil
}

func (s *Server) EnsureShadowNodeDirectoriesExists() error {
	if _, err := os.Stat("scripts"); err != nil && !os.IsNotExist(err) {
		return err
	} else if _, err := os.Stat("snapshots"); err != nil && !os.IsNotExist(err) {
		return err
	} else if _, err := os.Stat("backups"); err != nil && !os.IsNotExist(err) {
		return err
	} else if err != nil {
		// Create the server data directory because it does not currently exist
		// on the system.
		if err := os.MkdirAll("scripts", 0700); err != nil {
			return err
		}
		if err := os.MkdirAll("snapshots", 0700); err != nil {
			return err
		}
		if err := os.MkdirAll("backups", 0700); err != nil {
			return err
		}

		if err := s.fs.Chown("/"); err != nil {
			s.Log().WithField("error", err).Warn("failed to chown server scripts, snapshots and backups directories")
		}
	}

	return nil
}
