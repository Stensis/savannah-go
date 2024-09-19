package model

import (
	"encoding/xml"
)

// Define structs to match the XML structure
type AfricasTalkingResponse struct {
	XMLName        xml.Name       `xml:"AfricasTalkingResponse"`
	SMSMessageData SMSMessageData `xml:"SMSMessageData"`
}

type SMSMessageData struct {
	Message    string     `xml:"Message"`
	Recipients Recipients `xml:"Recipients"`
}

type Recipients struct {
	Recipient []Recipient `xml:"Recipient"`
}

type Recipient struct {
	Number       string `xml:"number"`
	Cost         string `xml:"cost"`
	Status       string `xml:"status"`
	StatusCode   string `xml:"statusCode"`
	MessageID    string `xml:"messageId"`
	MessageParts string `xml:"messageParts"`
}
