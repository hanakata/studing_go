package lexer

type Lexer struct {
	input        string
	position     int  //現在検査中のバイトchの位置を指し示す
	readPosition int  //入力における次の位置を指し示す
	ch           byte //現在検査中の文字
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

//次の一文字を読んでinput文字列の現在位置を進める
func (l *Lexer) readChar() {
	//終端チェック
	if l.readPosition >= len(l.input) {
		//終端まで来てれば0
		l.ch = 0
	} else {
		//途中なら次の文字へ
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
