package ergo

import (
	"github.com/bregydoc/ergo/schema"
	"golang.org/x/text/language"
)

// var unknownErrorInstance = &schema.ErrorInstance{
// 	Id:   []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
// 	Code: uint64(100),
// 	Type: schema.ErrorType_ONLY_READ,
// }

var unknownErrorForHumans = &schema.ErrorHuman{
	Id:   []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	Type: schema.ErrorType_ONLY_READ,
	Action: []*schema.Action{
		{
			Message: "Contact with us",
			Link:    "/support",
		},
	},
	Image: "error_not_found.jpg",
	Messages: []*schema.UserMessage{
		{
			Language: language.English.String(),
			Message:  "Sorry, we're not register this error in our systems. If you want, you can send a feedback.",
		},
	},
}

var unknownErrorForDevelopers = &schema.ErrorDev{
	Id:       []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	Type:     schema.ErrorType_ONLY_READ,
	Code:     uint64(100),
	Explain:  "unknown error, ergo could not found",
	Raw:      "unknown error, ergo could not found",
	Feedback: []*schema.Feedback{},
	Where:    "ergo",
}

var ergoIsNotWorkingForHumans = &schema.ErrorHuman{
	Id:   []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	Type: schema.ErrorType_ONLY_READ,
	Action: []*schema.Action{
		{
			Message: "Contact with us",
			Link:    "/support",
		},
	},
	Image: "error_panic.jpg",
	Messages: []*schema.UserMessage{
		{
			Language: language.English.String(),
			Message:  "Our systems are not working correctly, please contact with us",
		},
	},
}

// var ergoIsNotWorkingForDevelopers = &schema.ErrorDev{
// 	Id:       []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
// 	Type:     schema.ErrorType_ONLY_READ,
// 	Code:     uint64(101),
// 	Explain:  "ergo is not working, please check the service",
// 	Raw:      "ergo is not working, please check the service",
// 	Feedback: []*schema.Feedback{},
// 	Where:    "ergo",
// }
