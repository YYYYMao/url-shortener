# url-shortener

#### 執行方法

    1. docker compose up -d
    2. http://localhost:8080/swagger/index.html

#### api doc
./api 內使用 gin-swagger 產出 docs/swagger  

#### unit test

    ENV=test go test ./...  

#### API Flow Chart

* create url

``` flow
start=>start: Client
verifyParam=>operation: 參數驗證
genUrlId=>operation: 產生 urlId
isSuccess=>condition: 驗證參數正確
isInsert=>condition: 資料建立正確
respondSuccess=>operation: 成功
response400Failure=>operation: 回傳 status 400 參數有誤
response500Failure=>operation: 回傳 status 500 internal server error
end=>end: Client

start->verifyParam->isSuccess
isSuccess(yes)->genUrlId->isInsert(yes)->respondSuccess->end
isSuccess(no)->response400Failure->end
isInsert(no)->response500Failure->end
```

* redirect url

``` flow
start=>start: Client
verifyUrlId=>operation: urlId 驗證
isUrlId=>condition: 驗證urlId正確
getFromCache=>operation: 從快取撈取資料
isCacheExist=>condition: 是否存在
getFromDB=>operation: 從資料庫撈取資料
isExist=>condition: 是否存在
checkData=>operation: 驗證資料
checkData2=>operation: 驗證資料
isNotCacheExpire=>condition: 是否在期限內
isNotExpire=>condition: 是否在期限內
respondSuccess=>operation: 成功導流
response404Failure=>operation: 回傳 status 404 url not found
response500Failure=>operation: 回傳 status 500 internal server error
end=>end: Client

start->verifyUrlId->isUrlId
isUrlId(yes)->getFromCache->isCacheExist
isCacheExist(yes)->checkData->isNotCacheExpire
isCacheExist(no)->getFromDB->isExist
isNotCacheExpire(yes)->respondSuccess->end
isNotCacheExpire(no)->response404Failure->end
isExist(yes)->isNotExpire
isExist(no,left)->response404Failure->end
isNotExpire(no,left)->response404Failure->end
isNotExpire(yes)->respondSuccess->end
isUrlId(no)->response404Failure->end
```

* 檔案架構

        -  api 
           實作 api 層函式
           router 實作 api 接口
           controller 將資料轉換成合適格式輸入輸出 
           service 處理 api 業務邏輯
        -  docs
           文件
        -  repositories
           實作資料庫溝通 interface
        -  utils
           其餘功能性函式

* 選擇 postgreSQL

      以目前場景 postgreSQL 和 mySQL 都是可行解決方案
      SQLite 不提供網路訪問 當有多台 server 時則不適用
      選用 postgreSQL 則是考慮未來資料量大後需要擴展 postgreSQL 提供更好的穩定性與一致性

* 選擇 redis

      以目前場景 Memcached 和 Redis 都是可行解決方案
      redis 支援持久化 針對短網址服務需要大量的讀取
      如果當 redis 需要重啟 可快速回復快取資料
      並且 redis 提供更多種 data type 與功能 在未來開發新功能更加彈性

* 選擇 gin

      gin 是目前 golang 主流框架擁有優秀的性能表現 


* urlId 設計方式

      shortUrl 設計為 a-zA-Z0-9 隨機產生長度6的字串 
      在使用當時產生的隨機數字的加總和為檢查碼 避免惡意嘗試
      ex
        隨機字串 abcdef  
        亂數和 (0+1+2+3+4+5) = 15
        最終 url_id = abcdef15
        後續驗證 url_id 會檢查前6碼計算結果與亂數是否一致


* 對於大量api同時操作情境,這部分可以使用 

      1. redis cache 減少 db 操作
      2. 使用 rate limiter 
        1. 架構可使用 api gateway 做限制
        2. 程式端可實作 rate limiter 在 middleware層針對api 接口做限制 

* 對於不存在的 url_id 這部分可以使用 

      1. 針對 url_id 做驗證減少資料庫存取

* 後續優化

      1. server auto scale
      2. 當資料多後 可以將 db 設計成讀寫分離 資料表也可依 usrId sharding
      3. 確認 url_id 是否可用避免 db collision 可用 bloom filter
      4. 統一 logger format



