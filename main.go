package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/shirobrak/app-review-slackbot/adapters"
	"github.com/shirobrak/app-review-slackbot/repositories"
	"github.com/shirobrak/app-review-slackbot/usecases"
)

func main() {
	// バッチ開始時間
	startTime := time.Now().Format("2006/1/2 15:04:05")
	_, err := fmt.Fprintf(os.Stdout, "[%s] Batch Start...\n", startTime)
	if err != nil {
		log.Fatal(err)
	}

	// 環境変数の読み込み
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Repository の作成
	iosReviewRepository := repositories.NewAppStoreRssRepository()
	slackWebhookRepository := repositories.NewSlackWebhookRepository()
	updateLogRepository := repositories.NewUpdateLogRepository()
	// Adapter の作成
	reviewGetter := adapters.NewReviewGetter(iosReviewRepository)
	reviewSender := adapters.NewReviewSender(slackWebhookRepository)
	batchManager := adapters.NewBatchManager(updateLogRepository)
	// UserCase の作成
	usecase := usecases.NewSendReviewUseCase(reviewGetter, reviewSender, batchManager)
	// ユースケース実行
	resp, err := usecase.Run()
	if err != nil {
		log.Fatal(err)
	}

	// バッチ処理結果出力
	finishTime := time.Now().Format("2006/1/2 15:04:05")
	_, err = fmt.Fprintf(os.Stdout, "[%s] Batch Finish Successfully!! %d reviews were posted.\n", finishTime, resp)
	if err != nil {
		log.Fatal(err)
	}
}
