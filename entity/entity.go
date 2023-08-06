package entity

type Question struct {
	// Информация об устройстве, с помощью которого пользователь
	// разговаривает с Алисой.
	Meta Meta `json:"meta"`

	// Данные, полученные от пользователя.
	Request Request `json:"request"`

	// Данные о сессии.
	Session Session `json:"session"`

	// Версия протокола.
	Version string `json:"version"`
}

type Request struct {
	// Тип ввода:
	// * SimpleUtterance — голосовой ввод;
	// * ButtonPressed — нажатие кнопки.
	Type string `json:"type"`

	// Формальные характеристики реплики, которые удалось выделить
	// Яндекс.Диалогам. Отсутствует, если ни одно из вложенных свойств не
	// применимо.
	Markup Markup `json:"markup,omitempty"`

	// Текст пользовательского запроса без активационных фраз Алисы и
	// конкретного навыка.
	Command string `json:"command"`

	// Полный текст пользовательского запроса.
	OriginalUtterance string `json:"original_utterance"`

	// JSON, полученный с нажатой кнопкой от обработчика навыка (в ответе на
	// предыдущий запрос).
	Payload Payload `json:"payload,omitempty"`
}

// Markup содержит формальные характеристики реплики, которые удалось
// выделить Яндекс.Диалогам. Отсутствует, если ни одно из вложенных свойств
// не применимо.
type Markup struct {
	// Признак реплики, которая содержит криминальный подтекст (самоубийство,
	// разжигание ненависти, угрозы). Вы можете настроить навык на
	// определенную реакцию для таких случаев — например, отвечать "Не
	// понимаю, о чем вы. Пожалуйста, переформулируйте вопрос."
	//
	// Возможно только значение true. Если признак не применим, это свойство
	// не включается в ответ.
	DangerousContext bool `json:"dangerous_context,omitempty"`
}

// Session содержит данные о сессии.
type Session struct {
	// Признак новой сессии. Возможные значения:
	// * true — пользователь начал новый разговор с навыком;
	// * false — запрос отправлен в рамках уже начатого разговора.
	New bool `json:"new"`

	// Уникальный идентификатор сессии, 64 байта.
	SessionID string `json:"session_id"`

	// Идентификатор сообщения в рамках сессии. Инкрементируется с каждым
	// следующим запросом.
	MessageID int64 `json:"message_id"`

	// Идентификатор вызываемого навыка.
	SkillID string `json:"skill_id"`

	// Обфусцированный идентификатор пользователя.
	UserID string `json:"user_id"`
}

// Payload представляет собой произвольные JSON данные, идущие c кнопкой.
type Payload interface{}

type Meta struct {
	// Язык в POSIX-формате.
	Locale string `json:"locale"`

	// Название часового пояса, включая алиасы.
	TimeZone string `json:"timezone"`

	// Идентификатор устройства и приложения, в котором идет разговор.
	ClientID string `json:"client_id"`
}

type Response struct {
	EndSession bool   `json:"end_session"`
	Text       string `json:"text"`
	Tts        string `json:"tts"`
}

type ResponseEntity struct {
	Version  string   `json:"version"`
	Session  string   `json:"session"`
	Response Response `json:"response"`
}
