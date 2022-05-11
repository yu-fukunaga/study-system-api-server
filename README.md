# Go プログラミング実践入門

## 使用技術

- Docker
- MySQL
- air
- direnv
- Makefile

## 環境構築

### direnv install

(https://github.com/direnv/direnv/blob/master/docs/hook.md)

install して、~/.zshrc に「eval "$(direnv hook bash)"」を記載する

Mac

```
$ brew install direnv
$ echo "eval \"\$(direnv hook bash)\"" >> ~/.zshrc
$ source ~/.zshrc
```

Ubuntu

```
$ sudo apt install direnv
$ echo "eval \"\$(direnv hook bash)\"" >> ~/.bashrc
$ source ~/.bashrc
```

.env ファイルを用意して下記コマンドを実行する

```
$ direnv allow .
```
