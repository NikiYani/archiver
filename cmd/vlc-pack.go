package cmd

import (
	"archiver/lib/vlc"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var vlcPackCmd = &cobra.Command{
	Use:   "vlc",
	Short: "pack file using variable-length code",
	Run:   pack,
}

const packedExtension = "vlc"

var ErrEmptyPath = errors.New("path ot file is not specified")

func pack(_ *cobra.Command, args []string) {
	if len(args) < 1 || args[0] == "" {
		handleError(ErrEmptyPath) 
	}

	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleError(err)
	}

	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		handleError(err)
	}

	packed := vlc.Encode(string(data))

	err = os.WriteFile(packedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleError(err)
	}

}

func packedFileName(path string) string {
	// /path/to/file/myFile.txt -> myFile.vlc
	fileName := filepath.Base(path)

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedExtension
}

func init() {
	packCmd.AddCommand(vlcPackCmd)
}
