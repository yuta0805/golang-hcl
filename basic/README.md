## parse and decodeの流れ
- hclparserにてhclファイルをparse(ParseHCL(), ParseHCLFIle)
- parseした内容を、用意したgolang構造体(スキーマ)へdecode(gohcl.DecodeBody)
これでparseからdecodeまで完了する
