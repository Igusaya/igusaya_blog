package clock

import "time"

// SQL実行時に利用する時刻情報を制御するためのインターフェイス
type Clocker interface {
	Now() time.Time
}

type RealClocker struct{}

// アプリケーションで実際に使用する
func (r RealClocker) Now() time.Time {
	return time.Now()
}

type FixedClocker struct{}

// テスト用に固定時刻を返却する
func (fc FixedClocker) Now() time.Time {
	return time.Date(2022, 5, 10, 12, 34, 59, 0, time.UTC)
}
