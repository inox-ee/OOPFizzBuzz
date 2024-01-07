# FizzBuzz with TDD

## About

[ちょうぜつソフトウェア設計入門 ――PHPで理解するオブジェクト指向の活用：書籍案内｜技術評論社](https://gihyo.jp/book/2022/978-4-297-13234-7) 第６章の、FizzBuzzを例としたテスト駆動開発を実践する

## Steps

### 振る舞い駆動開発 (BDD)

- 外形となるクラス等を作成する
- メソッドは書かず、テストを書く
  - 呼び出せないが、気にしない
- その後、呼び出せるメソッドを書く
  - わざと失敗する実装を書く
  - 必ず最初に、テストが失敗することを確認する
- テストが失敗することを確認したら、テストを成功させる実装を書く
  - ちゃんとした実装は書かない
  - 次に、思いつく失敗ケースをテストに追加する
    - 失敗することを確認する
  - 最小限の実装を書く
    - テストが成功することを確認する
    - 繰り返し
- 実装コードがひどい有様になってきたらようやくリファクタを検討する

[!TIP]
TDDは実装の美しさは問わない。クールなアルゴリズムの優先度は最低。
振る舞いを満たす方針でテストと実装を進める

[!IMPORTANT]
「テスト追加」「失敗確認」「最小限の実装「成功」のサイクルを繰り返すことが、テストファースト
リファクタ=**動きを全く変えずに**プログラムコードを書き換えること。動きが全く変わらない保障は、**テストシナリオによって担保される**

### テスト駆動開発 (TDD)

- 実装内容の抽象を考え、仮説を立ててみる
  - 抽象の例=インターフェース
  - hoge(具象)は、fuga(抽象)の集約なのでは？
  - **まだ設計の整合性を確認する段階ではない**
    - テストを書く
- 境界条件をもとに、テストを書き、最小限の実装を記述する
  - スタブやモックを活用
- テストを記述している間に、仕様をより一般化できないか模索する
- イイカンジになってきたら、仮説で作成した抽象に対する具体的な実装を作成する
- 仮説のミスが見えてくる
  - 時にはインターフェースを直す

[!TIP]
スタブ=依存クラスの特定のメソッドが、決まった値を返すと仮定した疑似オブジェクト。リアルな実装無しにテストを書くことができる。
モック=使われ方を検証するための疑似オブジェクト。ダミーメソッドの呼び出しが行われたか、そのパラメータが何であったかを検証できる。

[!TIP]
構造化プログラミング的には、プロパティを持たないオブジェクトには意味がないと考えることも。
しかし、オブジェクト指向的には、プロパティを持たない、純粋関数のようなオブジェクトもあり得る。具象データにメソッドを生やすのではなく、データの有無に関係ない抽象概念をキーにすることが大事。

[!IMPORTANT]
インターフェースは実装クラスが無いと動作させることが出来ないため、スタブやモックを用いてテストを書く
動かない時には、動かすことに意識を向いて大事なことを忘れがち。**安定度の高低の向き**を意識すること。

### 依存性注入 (DI)

- 上記ステップで作成したモジュールを組み合わせたアプリケーション全体像を実装する
  - 実行できるコードになったとしても、オブジェクト指向に照らし合わせて妥当な設計になっているだろうか？
- メソッド内部で、クラスやモジュールのインスタンスを作成してはいけない
  - パッケージ間で強い依存関係が生まれてしまう
  - サブモジュールの変更影響を直接受ける上位構造になってしまう
- 具象インスタンスを外部から与えるように変更
  - 正しくインスタンス作成する責務を外部に分離する
  - 自身の責務以外に関心を払わなくて済む、閉じた設計になる

[!TIP]
モジュールの変更影響や検証しにくい動作の厄介さを切り離すには、それらを内部から追い出すのが得策

[!IMPORTANT]
依存性注入を行なうことで、依存モジュールの使うことに集中すべき実装から、生成と構築に関する知識を分離することができる
