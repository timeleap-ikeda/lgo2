package env

import "testing"

// ProcessEnvVarsは環境変数を処理し、OutputFormatのフィールドをもつ構造体を返す関数  //liststart1
func TestEnvVarProcess(t *testing.T) {
	t.Setenv("OUTPUT_FORMAT", "JSON")  // 環境変数OUTPUT_FORMATを設定
	cfg := ProcessEnvVars()  // テストする関数を呼び出し
	if cfg.OutputFormat != "JSON" {  // 確認
		t.Error("OutputFormat not set correctly")
	}
	// テスト関数が終了するとOUTPUT_FORMATの値がリセットされる
}  //listend1
