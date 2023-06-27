## HCL package
- HCLをparse, decode encode, writeする為に用いらるパッケージ

※HCLパッケージはHCLとHCL2が存在する。HCL2は現在public archiveになっている。理由としてはHCL -> HCL2へアップデートする過程で試験的に作られたpackageであり、この検証は終了しているためarchiveになった。その代わり,
HCLパッケージに統合されている。なのでHCL2を使用する(現在はterraformの基準はHCL2)場合は```github.com/hashicorp/hcl/v2```を用いること

- [HCL/V2](https://pkg.go.dev/github.com/hashicorp/hcl/v2)
