まず、Level とその定数を定義します。


type Level string

const (
	Debug Level = "debug"
	Info  Level = "info"
)


次に、ログレベルのコンテキストを管理する関数を作成する必要があります。まず、キーの型に対してエクスポートされない型を定義し、キーに対してエクスポートされない定数を定義します。


type logKey int

const (
	_ logKey = iota
	key
)


定義した2つの型を使ってコンテキストの値を管理する関数を書きます。


func ContextWithLevel(ctx context.Context, level Level) context.Context {
	return context.WithValue(ctx, key, level)
}

func LevelFromContext(ctx context.Context) (Level, bool) {
	level, ok := ctx.Value(key).(Level)
	return level, ok
}


これで、関数ContextWithLevelと型Levelを使ってミドルウェアを記述できます。


func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		level := r.URL.Query().Get("log_level")
		ctx := ContextWithLevel(r.Context(), Level(level))
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}


最後に、LevelFromContextを使って関数Logの不足している行を埋めます。


	inLevel, ok := LevelFromContext(ctx)
	if !ok {
		return
	}

