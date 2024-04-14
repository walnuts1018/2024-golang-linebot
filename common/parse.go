// このファイルには、授業情報を文字列からパースするための関数が定義されています。
package common

import (
	"fmt"
	"regexp"
	"strconv"
)

// 「月1 パワーエレクトロニクス 電総大」などの文字列にヒットする正規表現です。
var regex = regexp.MustCompile(`^(?P<weekday>[月火水木金土日])(?P<period>\d)\s(?P<name>.+)\s(?P<room>.+)\s?$`)

// ParseFromString は、文字列から授業情報を取り出します。
func ParseFromString(s string) (Subject, error) {
	// 上で定義した正規表現を使って、文字列から授業情報を取り出します。
	params := regex.FindStringSubmatch(s)
	// 👆paramsの中身は["月1 パワーエレクトロニクス 電総大", "月", "1", "パワーエレクトロニクス", "電総大"]のような配列になります。

	// ここで、正規表現にマッチしなかった場合はエラーを返します。
	// paramsの長さが5であれば、正しくすべてのグループがマッチしているということです。
	if len(params) != regex.NumSubexp()+1 {
		return Subject{}, fmt.Errorf("形式が間違っています。「月1 パワーエレクトロニクス 電総大」などと入力してください")
	}

	// マッチしたグループを取り出します。
	weekday := params[regex.SubexpIndex("weekday")] //曜日
	name := params[regex.SubexpIndex("name")]       // 授業名
	room := params[regex.SubexpIndex("room")]       // 教室

	// 文字列を数字に変換します
	period, err := strconv.Atoi(params[regex.SubexpIndex("period")]) // 時限
	if err != nil {
		return Subject{}, fmt.Errorf("時間の取得に失敗しました: %v", err)
	}

	// Subject構造体を作成して返します。
	return Subject{
		Name:    name,                  // 授業名
		Weekday: ParseWeekday(weekday), // 曜日, ParseWeekday関数は文字列から曜日型に変換する関数です。
		Period:  period,                // 時限
		Room:    room,                  // 教室
	}, nil
}
