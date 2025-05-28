# soundNotify
Listen 8080 port 當被呼叫時播放音樂

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
