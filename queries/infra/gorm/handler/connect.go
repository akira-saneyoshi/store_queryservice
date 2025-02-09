package handler

import (
	"fmt"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"gorm.io/driver/mysql" // MySQL ドライバ
	"gorm.io/gorm"         // GORM
)

// データベース接続情報（DBConfig型は変更なし）
type DBConfig struct {
	Dbname string `toml:"dbname"` //	データベース名
	Host   string `toml:"host"`   //	ホスト名
	Port   int64  `toml:"port"`   //	ポート番号
	User   string `toml:"user"`   //	ユーザー名
	Pass   string `toml:"pass"`   //	パスワード
}

// database.tomlから接続情報を取得してDbConfig型で返す（tomlRead関数は変更なし）
func tomlRead() (*DBConfig, error) {
	// 環境変数からファイルパスを取得する
	path := os.Getenv("DATABSE_TOML_PATH")
	if path == "" {
		// 環境変数が無い場合のパスを設定する
		path = "infra/gorm/config/database.toml"
	}
	// database.tomlを読取りDBConfigにマッピングする
	m := map[string]DBConfig{}
	_, err := toml.DecodeFile(path, &m)
	if err != nil {
		return nil, err
	}
	config := m["mysql"]
	return &config, nil
}

// GORMを使用してデータベースに接続する関数
func ConnectDB() (*gorm.DB, error) {
	config, err := tomlRead() // database.tomlの定義内容を読み取る
	if err != nil {
		return nil, DBErrHandler(err) // エラーハンドリングは元の関数を再利用
	}

	// 接続文字列を生成する (GORM用に少し変更)
	connectString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Pass,
		config.Host,
		config.Port,
		config.Dbname,
	)

	// GORMを使用してデータベースに接続
	db, err := gorm.Open(mysql.Open(connectString), &gorm.Config{})
	if err != nil {
		return nil, DBErrHandler(err) // エラーハンドリングは元の関数を再利用
	}

	sqlDB, err := db.DB() // gorm.DB から sql.DB を取得
	if err != nil {
		return nil, DBErrHandler(err) // sql.DB取得時のエラーハンドリング
	}

	MAX_IDLE_CONNS := 10                   // 初期接続数
	MAX_OPEN_CONNS := 100                  // 最大接続数
	CONN_MAX_LIFETIME := 300 * time.Second // 最大生存期間

	// コネクションプールの設定 (sql.DBに対して設定)
	sqlDB.SetMaxIdleConns(MAX_IDLE_CONNS)
	sqlDB.SetMaxOpenConns(MAX_OPEN_CONNS)
	sqlDB.SetConnMaxLifetime(CONN_MAX_LIFETIME)

	// デバッグモード (GORMのLoggerでSQL出力)
	// 必要であれば、ロガーを設定してSQLログを出力できます。
	// 例：db.Logger = logger.Default.LogMode(logger.Info)

	return db, nil // GORMのDBインスタンスを返す
}
