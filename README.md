# soundNotify
Listen 8080 port 當被呼叫時播放音樂

## 目的
本程式僅為基本雛形，主要概念為，在本機運行後，可以讓外部系統或一些監控腳本，透過 http 發出提示聲

實作上曾在瀏覽器腳本自動化腳本中發揮作用，因為瀏覽器自動化執行時，若沒有接收到鍵盤滑鼠信號，瀏覽器不允許發出聲音，因此利用另外起一個服務的方式，讓自動化腳本以這樣的方式觸發提示聲音

## 目前於以下環境測試運行OK

- Mac mini M4
- Mac intel CPU
- Window 11

## 事前準備

事先安裝好 go

準備好一個 sound.mp3 檔案置於 git 根目錄

## Install

將 git 內容下載到資料夾後於資料夾中進行初始化
```shell
go mod init soundNotify
```

進行下列安裝
```
go get github.com/faiface/beep
go get github.com/faiface/beep/mp3
go get github.com/faiface/beep/speaker
```
或自動安裝
```
go mod tidy
```

## 啟動
```
go run main.go
```

於瀏覽器輸入 ```http://localhost:8080``` 即可觸發播放音樂
播放過程可於 cli 介面輸入 q + enter 中斷

## 打包

若你需要打包成執行黨，需自行將
```
// go run:
f, err := os.Open("sound.mp3")
// go build 可用下列
// exePath, _ := os.Executable()
// f, err := os.Open(filepath.Join(filepath.Dir(exePath), "sound.mp3"))
```
修改為
```
// go run:
// f, err := os.Open("sound.mp3")
// go build 可用下列
exePath, _ := os.Executable()
f, err := os.Open(filepath.Join(filepath.Dir(exePath), "sound.mp3"))
```

這部分不是什麼大問題，懶得寫完善了
