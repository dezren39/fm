package filetree

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/atotto/clipboard"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/mistakenelf/fm/filesystem"
)

type getDirectoryListingMsg struct {
	files            []DirectoryItem
	workingDirectory string
}
type errorMsg string
type copyToClipboardMsg string
type statusMessageTimeoutMsg struct{}
type editorFinishedMsg struct{ err error }
type createFileMsg struct{}
type createDirectoryMsg struct{}

// NewStatusMessageCmd sets a new status message, which will show for a limited
// amount of time. Note that this also returns a command.
func (m *Model) NewStatusMessageCmd(s string) tea.Cmd {
	m.StatusMessage = s

	if m.statusMessageTimer != nil {
		m.statusMessageTimer.Stop()
	}

	m.statusMessageTimer = time.NewTimer(m.StatusMessageLifetime)

	return func() tea.Msg {
		<-m.statusMessageTimer.C
		return statusMessageTimeoutMsg{}
	}
}

// CreateDirectoryCmd creates a directory based on the name provided.
func (m *Model) CreateDirectoryCmd(name string) tea.Cmd {
	return func() tea.Msg {
		if err := filesystem.CreateDirectory(name); err != nil {
			return errorMsg(err.Error())
		}

		return createDirectoryMsg{}
	}
}

// CreateFileCmd creates a file based on the name provided.
func (m *Model) CreateFileCmd(name string) tea.Cmd {
	return func() tea.Msg {
		if err := filesystem.CreateFile(name); err != nil {
			return errorMsg(err.Error())
		}

		return createFileMsg{}
	}
}

// GetDirectoryListingCmd updates the directory listing based on the name of the directory provided.
func (m Model) GetDirectoryListingCmd(directoryName string) tea.Cmd {
	return func() tea.Msg {
		var err error
		var directoryItems []DirectoryItem
		var files []fs.DirEntry

		if directoryName == filesystem.HomeDirectory {
			directoryName, err = filesystem.GetHomeDirectory()
			if err != nil {
				return errorMsg(err.Error())
			}
		}

		directoryInfo, err := os.Stat(directoryName)
		if err != nil {
			return errorMsg(err.Error())
		}

		if !directoryInfo.IsDir() {
			return nil
		}

		err = os.Chdir(directoryName)
		if err != nil {
			return errorMsg(err.Error())
		}

		if !m.showDirectoriesOnly && !m.showFilesOnly {
			files, err = filesystem.GetDirectoryListing(directoryName, m.showHidden)
			if err != nil {
				return errorMsg(err.Error())
			}
		} else {
			listingType := filesystem.DirectoriesListingType

			if m.showFilesOnly {
				listingType = filesystem.FilesListingType
			}

			files, err = filesystem.GetDirectoryListingByType(directoryName, listingType, m.showHidden)
			if err != nil {
				return errorMsg(err.Error())
			}
		}

		workingDirectory, err := filesystem.GetWorkingDirectory()
		if err != nil {
			return errorMsg(err.Error())
		}

		for _, file := range files {
			fileInfo, err := file.Info()
			if err != nil {
				continue
			}

			fileSize := ConvertBytesToSizeString(fileInfo.Size())

			directoryItems = append(directoryItems, DirectoryItem{
				Name:             file.Name(),
				Details:          fileInfo.Mode().String(),
				Path:             filepath.Join(workingDirectory, file.Name()),
				Extension:        filepath.Ext(fileInfo.Name()),
				IsDirectory:      fileInfo.IsDir(),
				CurrentDirectory: workingDirectory,
				FileInfo:         fileInfo,
				FileSize:         fileSize,
			})
		}

		return getDirectoryListingMsg{
			files:            directoryItems,
			workingDirectory: workingDirectory,
		}
	}
}

// deleteDirectoryItemCmd deletes a directory based on the name provided.
func deleteDirectoryItemCmd(name string, isDirectory bool) tea.Cmd {
	return func() tea.Msg {
		if isDirectory {
			if err := filesystem.DeleteDirectory(name); err != nil {
				return errorMsg(err.Error())
			}
		} else {
			if err := filesystem.DeleteFile(name); err != nil {
				return errorMsg(err.Error())
			}
		}

		return nil
	}
}

// zipDirectoryCmd zips a directory based on the name provided.
func zipDirectoryCmd(name string) tea.Cmd {
	return func() tea.Msg {
		if err := filesystem.Zip(name); err != nil {
			return errorMsg(err.Error())
		}

		return nil
	}
}

// unzipDirectoryCmd unzips a directory based on the name provided.
func unzipDirectoryCmd(name string) tea.Cmd {
	return func() tea.Msg {
		if err := filesystem.Unzip(name); err != nil {
			return errorMsg(err.Error())
		}

		return nil
	}
}

// copyDirectoryItemCmd copies a directory based on the name provided.
func copyDirectoryItemCmd(name string, isDirectory bool) tea.Cmd {
	return func() tea.Msg {
		if isDirectory {
			if err := filesystem.CopyDirectory(name); err != nil {
				return errorMsg(err.Error())
			}
		} else {
			if err := filesystem.CopyFile(name); err != nil {
				return errorMsg(err.Error())
			}
		}

		return nil
	}
}

// copyToClipboardCmd copies the provided string to the clipboard.
func copyToClipboardCmd(name string) tea.Cmd {
	return func() tea.Msg {
		workingDir, err := filesystem.GetWorkingDirectory()
		if err != nil {
			return errorMsg(err.Error())
		}

		filePath := filepath.Join(workingDir, name)
		err = clipboard.WriteAll(filePath)
		if err != nil {
			return errorMsg(err.Error())
		}

		return copyToClipboardMsg(fmt.Sprintf("%s %s %s", "Successfully copied", filePath, "to clipboard"))
	}
}

// writeSelectionPathCmd writes content to the file specified.
func writeSelectionPathCmd(selectionPath, filePath string) tea.Cmd {
	return func() tea.Msg {
		if err := filesystem.WriteToFile(selectionPath, filePath); err != nil {
			return errorMsg(err.Error())
		}

		return nil
	}
}

func openEditorCmd(file string) tea.Cmd {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}

	c := exec.Command(editor, file)

	return tea.ExecProcess(c, func(err error) tea.Msg {
		return editorFinishedMsg{err}
	})
}
