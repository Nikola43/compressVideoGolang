package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	//inputFile := "/Users/paulo/go/src/github.com/nikola43/compressVideoGolang/assets/video/video1.mp4"
	//outputFile := "/Users/paulo/go/src/github.com/nikola43/compressVideoGolang/assets/video/video1Out.mp4"

	inputFile := "./assets/video/video1.mp4"
	outputFile := "./assets/video/video10.mp4"

	err := compressMP4(inputFile, outputFile)
	checkError(err)

	fmt.Print(inputFile + "->" + strconv.FormatInt(getFileSize(inputFile), 10) + " " + strconv.FormatInt(calculateCompressionPercentage(getFileSize(inputFile), getFileSize(inputFile)), 10) + "%")
	fmt.Print(outputFile + "->" + strconv.FormatInt(getFileSize(outputFile), 10) + " " + strconv.FormatInt(calculateCompressionPercentage(getFileSize(inputFile), getFileSize(inputFile)), 10) + "%")
}

func calculateCompressionPercentage(originalFileSize int64, outputFileSize int64) int64 {
	return (100 * outputFileSize) / originalFileSize
}

func compressMP4(inFile string, outFile string) error {
	// check if input file exists
	err := CheckIfFileExists(inFile)
	if err != nil {
		return err
	}

	/*
		// check if input file exists
		err = CheckIfFileExists(outFile)
		if err != nil {
			return err
		} else {
			//if exists then remove
			fmt.Println("output file called " + outFile + " already exist. Removing...")
			removeError := os.Remove(outFile)
			if removeError != nil {
				return removeError
			}
			fmt.Println("file " + outFile + " has been removed successfully")
		}
	*/

	// extract audio from video using ffmpeg library
	// ffmpeg -i input.mp4 -vcodec h264 -acodec aac output.mp4
	err = executeCommandVerbose("ffmpeg", "-i", inFile, "-vcodec", "h264", "-acodec", "mp3", outFile)
	if err != nil {
		return err
	}
	err = CheckIfFileExists(outFile)
	if err != nil {
		return err
	}
	return nil
}

/**
Check if file exist
*/
func CheckIfFileExists(f string) error {
	var err error
	if f, err := os.Stat(f); err == nil && f.Size() > 0 {
		return nil
	}
	return err
}

func getFileSize(filePath string) int64 {
	file, err := os.Stat(filePath)
	checkError(err)
	return file.Size()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("ok")
	}
}

func executeCommandVerbose(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return err
	}
	fmt.Println("Result: ")
	fmt.Println(out.String())
	return nil
}
