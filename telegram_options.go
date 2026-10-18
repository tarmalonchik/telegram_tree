package telegram_tree

type TelegramOpt func(v *telegram)

// Style is the color of an inline button.
type Style string

const (
	StyleDanger  Style = "danger"  // red
	StyleSuccess Style = "success" // green
	StylePrimary Style = "primary" // blue
)

type Telegram interface {
	GetMessage() string
	GetTabTxt() string
	GetHideBar() bool
	GetSwitchInlineQueryCurrentChat() *string
	GetLink() *string
	GetEnablePreview() bool
	DeleteMessage() bool
	GetResendMsg() bool
	setDefaultMessage(string)
	GetColumns() int
	GetStyle() Style
	GetCopyText() *string
	GetDisabled() bool
}

type telegram struct {
	message                      string
	tabTxt                       string
	hideBar                      bool
	enablePreview                bool
	deleteMsg                    bool
	resendMsg                    bool
	switchInlineQueryCurrentChat *string
	link                         *string
	columns                      int
	style                        Style
	copyText                     *string
	disabled                     bool
}

func (t *telegram) GetMessage() string     { return t.message }
func (t *telegram) GetColumns() int        { return t.columns }
func (t *telegram) GetTabTxt() string      { return t.tabTxt }
func (t *telegram) GetHideBar() bool       { return t.hideBar }
func (t *telegram) GetEnablePreview() bool { return t.enablePreview }
func (t *telegram) DeleteMessage() bool    { return t.deleteMsg }
func (t *telegram) GetResendMsg() bool     { return t.resendMsg }
func (t *telegram) GetSwitchInlineQueryCurrentChat() *string {
	return t.switchInlineQueryCurrentChat
}
func (t *telegram) GetLink() *string {
	return t.link
}
func (t *telegram) GetStyle() Style             { return t.style }
func (t *telegram) GetCopyText() *string        { return t.copyText }
func (t *telegram) GetDisabled() bool           { return t.disabled }
func (t *telegram) setDefaultMessage(in string) { t.message = in }

func NewTelegram(options ...TelegramOpt) Telegram {
	tg := telegram{
		columns: 1,
	}
	for _, opt := range options {
		opt(&tg)
	}
	return &tg
}

func WithSwitchInline(inline string) TelegramOpt {
	return func(v *telegram) {
		v.switchInlineQueryCurrentChat = &inline
	}
}

func WithLink(link string) TelegramOpt {
	return func(v *telegram) {
		v.link = &link
	}
}

func WithHideBar() TelegramOpt {
	return func(v *telegram) {
		v.hideBar = true
	}
}

func WithMessage(msg string) TelegramOpt {
	return func(v *telegram) {
		v.message = msg
	}
}

func WithTabTxt(tabTxt string) TelegramOpt {
	return func(v *telegram) {
		v.tabTxt = tabTxt
	}
}

func WithColumns(columnsCount int) TelegramOpt {
	return func(v *telegram) {
		v.columns = columnsCount
	}
}

func DeleteMsg() TelegramOpt {
	return func(v *telegram) {
		v.deleteMsg = true
	}
}

func EnablePreview() TelegramOpt {
	return func(v *telegram) {
		v.enablePreview = true
	}
}

func ResendMsg() TelegramOpt {
	return func(v *telegram) {
		v.resendMsg = true
	}
}

func WithStyle(style Style) TelegramOpt {
	return func(v *telegram) {
		v.style = style
	}
}

// WithCopyText makes the button copy the text to the clipboard instead of sending a callback.
func WithCopyText(text string) TelegramOpt {
	return func(v *telegram) {
		v.copyText = &text
	}
}

// Disabled shows the button greyed out; pressing it does nothing.
func Disabled() TelegramOpt {
	return func(v *telegram) {
		v.disabled = true
	}
}
