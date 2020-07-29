package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	//inputFile := "/Users/paulo/go/src/github.com/nikola43/compressVideoGolang/assets/video/video1.mp4"
	inputFile := "./assets/video/video1.mp4"
	//outputFile := "/Users/paulo/go/src/github.com/nikola43/compressVideoGolang/assets/video/video1Out.mp4"
	outputFile := "./assets/video/video1Out.mp4"


	cmd := exec.Command("ffmpeg", "-i", inputFile, "-vcodec", "h264", "-acodec", "acc", outputFile)

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}



	//err := compressMP4(inputFile, outputFile)
	// checkError(err)
	/*file, err := os.Open(rFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(file.Name())
	*/

	fmt.Print(inputFile + "->" + strconv.FormatInt(getFileSize(inputFile), 10))
	//fmt.Print(outputFile + "->" + strconv.FormatInt(getFileSize(outputFile), 10))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("ok")
	}
}

func getFileSize(filePath string) int64 {
	file, err := os.Stat(filePath)
	checkError(err)
	return file.Size()
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


func compressMP4(inFile string, outFile string) error {
	// check if input file exists
	inFileError := CheckIfFileExists(inFile)
	if inFileError != nil {
		log.Fatal(inFileError)
		return  inFileError
	}

	/*
	// check if output file exists
	outFileError := CheckIfFileExists(outFile)
	if outFileError != nil {
		fmt.Printf("Error: %s", outFileError.Error())
		return "", outFileError
	} else {
		// // if exists then remove
		// removeFileError := os.Remove(outFile)
		// if removeFileError != nil {
		// 	fmt.Printf("Error: %s", removeFileError.Error())
		// 	return removeFileError
		// }
	}
	*/

	// extract audio from video using ffmpeg library
	// ffmpeg -i input.mp4 -vcodec h264 -acodec aac output.mp4
	cmd := exec.Command("ffmpeg", "-i", inFile, "-vcodec", "h264", "-acodec", "acc", outFile)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

/*
func compressVideo(inputFilePath, outFilePath string) {
	// check if input file exists
	CheckIfFileExists(inputFilePath)


	// ffmpeg -i input.mp4 -vcodec h264 -acodec aac output.mp4
	// extract audio from video using ffmpeg library
	cmd := exec.Command("echo", "hola")
	//cmd := exec.Command("ffmpeg", "-i", inputFilePath, "-vcodec", "h264", "-acodec", "acc", outFilePath)
	err := cmd.Run()
	checkError(err)
	fmt.Print("no error")
}
*/
