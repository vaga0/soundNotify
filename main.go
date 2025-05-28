package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync/atomic"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var (
	isPlaying int32 = 0
)

const sampleRate = 44100 // 固定 SampleRate，跨平台穩定

func playHandler(w http.ResponseWriter, r *http.Request) {
	if !atomic.CompareAndSwapInt32(&isPlaying, 0, 1) {
		http.Error(w, "Audio is already playing", http.StatusTooEarly)
		return
	}

	go func() {
		defer atomic.StoreInt32(&isPlaying, 0)

		// go run:
		f, err := os.Open("sound.mp3")
		// go build 可用下列
		// exePath, _ := os.Executable()
		// f, err := os.Open(filepath.Join(filepath.Dir(exePath), "sound.mp3"))
		if err != nil {
			fmt.Println("Open error:", err)
			return
		}
		defer f.Close()

		streamer, format, err := mp3.Decode(f)
		if err != nil {
			fmt.Println("Decode error:", err)
			return
		}
		defer streamer.Close()

		resampled := beep.Resample(4, format.SampleRate, beep.SampleRate(sampleRate), streamer)
		ctrl := &beep.Ctrl{Streamer: resampled, Paused: false}
		done := make(chan bool)

		speaker.Play(beep.Seq(ctrl, beep.Callback(func() {
			done <- true
		})))

		<-done // 等待播放完成
	}()

	w.Write([]byte("Playing audio..."))
}

func main() {
	// speaker 初始化（只做一次）
	sr := beep.SampleRate(sampleRate)
	speaker.Init(sr, sr.N(time.Second/10))

	// CLI 快捷鍵：輸入 q 強制中止播放
	go func() {
		fmt.Println("Type 'q' + Enter to force stop playback")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if scanner.Text() == "q" {
				fmt.Println("Force stop by CLI input")
				speaker.Clear()
				atomic.StoreInt32(&isPlaying, 0)
			}
		}
	}()

	http.HandleFunc("/", playHandler)
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
