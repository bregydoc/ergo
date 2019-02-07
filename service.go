package ergo

import (
	"context"

	"github.com/bregydoc/ergo/schema"
	"github.com/oklog/ulid"
	"golang.org/x/text/language"
)

type Server struct {
	ergo Ergo
}

func NewErgoServer(ergo *Ergo) *Server {
	return &Server{
		ergo: *ergo, // warning warning, pay attention
	}
}

func (s *Server) RegisterNewError(c context.Context, params *schema.ErrorSeed) (*schema.ErrorInstance, error) {
	langMessage, err := language.Parse(params.MessageLanguage)
	if err != nil {
		return nil, err
	}

	if params.Code == 0 || int64(params.Code) < 0 {
		// Register with random code...
		return s.ergo.RegisterNewError(
			params.Where,
			params.Explain,
			&UserMessage{
				Language: langMessage,
				Message:  params.MessageContent,
			},
			params.WithFeedback,
		)
	}
	return s.ergo.RegisterNewErrorWithCode(
		params.Where,
		params.Explain,
		params.Code,
		&UserMessage{
			Language: langMessage,
			Message:  params.MessageContent,
		},
		params.WithFeedback,
	)

}

func (s *Server) RegisterFullError(c context.Context, params *schema.FullErrorSeed) (*schema.ErrorInstance, error) {
	panic("unimplemented")
}

func (s *Server) ConsultErrorAsHuman(c context.Context, params *schema.ConsultAsHuman) (*schema.ErrorHuman, error) {
	var languages []language.Tag
	for _, l := range params.Languages {
		fLang, err := language.Parse(l)
		if err != nil {
			return nil, err
		}
		languages = append(languages, fLang)
	}
	return s.ergo.ConsultErrorAsHumanByID(params.ErrorID, languages...)
}

func (s *Server) ConsultErrorAsDeveloper(c context.Context, params *schema.ConsultAsDev) (*schema.ErrorDev, error) {
	return s.ergo.ConsultErrorAsDeveloperByID(params.ErrorID)
}

// Save new messages
func (s *Server) MemorizeNewMessages(c context.Context, params *schema.NewMessageParams) (*schema.UserMessages, error) {
	var messages []*UserMessage
	for _, m := range params.Messages {
		lang, err := language.Parse(m.Language)
		if err != nil {
			return nil, err
		}
		messages = append(messages, &UserMessage{Language: lang, Message: m.Message})
	}

	resultMessages, err := s.ergo.MemorizeNewMessages(params.ErrorID, params.WithAutoTranslate, messages...)
	if err != nil {
		return nil, err
	}

	return &schema.UserMessages{
		Messages: resultMessages,
	}, nil
}

// Save new feedback
func (s *Server) ReceiveFeedbackOfUser(c context.Context, params *schema.NewFeedBack) (*schema.Feedback, error) {
	var byID ulid.ULID
	copy(byID[:], params.Feedback.ByID)

	return s.ergo.ReceiveFeedbackOfUser(params.ErrorID, &UserFeedback{
		Message: params.Feedback.Message,
		By:      params.Feedback.By,
		ByID:    byID,
	})
}
