# OCR Dere Result

デレステの結果画面のキャプチャ画像を食わせて、諸情報をテキスト出力するコマンド。  

出力先は複数実装出来る体だけど今のところcsv文字列の標準出力だけ。  
## Getting started

* requirements: Google Cloud Vision APIが呼べる
```sh
$ export GOOGLE_APPLICATION_CREDENTIALS=/path/to/service_account_key.json
```
* `configs/coordinates.json`でOCRしたい領域を指定する
    * リポジトリに含めてあるものはiPad Pro向けに書いてある
* run
    * `app /path/to/result1.png /path/to/result2.png`
