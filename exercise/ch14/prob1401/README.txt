ややこしいですが「関数を返す関数を返す関数」を書くことになります。関数Timeoutは次のようになります。

func Timeout(ms int) func(http.Handler) http.Handler {
    return func(h http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            ctx := r.Context()
            ctx, cancelFunc := context.WithTimeout(ctx, time.Duration(ms)*time.Millisecond)
            defer cancelFunc()
            r = r.WithContext(ctx)
            h.ServeHTTP(w, r)
        })
    }
}


シグナチャは問題の説明と一致します。本体は2つのネストされた関数を宣言します。外側の関数は func(h http.Handler) http.Handler です。この関数は func(w http.ResponseWriter, r *http.Request) を返します。これはインターフェイスhttp.Handlerを満たす http.HandlerFunc型に変換されます。

内側の関数の本体は、コンテキストに関するすべての作業を行います。
 * リクエストからコンテキストを抽出
 * それをcontext.WithTimeoutによって生成されたコンテキストにラップ
 * deferでcancelFuncを呼び出し
 * 新しいコンテキストと古いリクエストを使って、古いリクエストのメソッドWithContextで、置換する*http.Requestを構築
 * 次にミドルウェアに渡されるhttp.HandlerのメソッドServeHTTPを呼び出します。

このコードには、タイムアウトが尊重されることを保証するものがない点に注意してください。それはリクエストハンドラとビジネスロジックの仕事です。

リクエストハンドラのコードは、次のような感じになります。

func sleepy(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	message, err := doThing(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			w.WriteHeader(http.StatusGatewayTimeout)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write([]byte(message))
}


ビジネスロジックには、コンテキストをチェックして、作業に時間がかかりすぎないようにするコードを含める必要があります。

func doThing(ctx context.Context) (string, error) {
	wait := rand.Intn(200)
	select {
	case <-time.After(time.Duration(wait) * time.Millisecond):
		return "Done!", nil
	case <-ctx.Done():
		return "Too slow!", ctx.Err()
	}
}

