package service

import (
	"encoding/json"
	"fmt"
	"github.com/pion/webrtc"
)

type WebRTC struct {
	PeerConnection  *webrtc.PeerConnection
	DataChannel *webrtc.DataChannel
	Offer string
	isOpen bool
}

var WebRTCConfig webrtc.Configuration
var WebRTCMap map[string]WebRTC

func init() {
	WebRTCConfig = webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}
}

func (o *WebRTC) CreateDataChannel(label string) (err error){
	o.PeerConnection, err = webrtc.NewPeerConnection(WebRTCConfig)
	if err != nil {
		panic(err)
	}

	o.DataChannel, err = o.PeerConnection.CreateDataChannel(label, nil)
	if err != nil {
		panic(err)
	}

	o.DataChannel.OnOpen(func() {
		fmt.Printf("Data channel '%s'-'%d'-'%v' !!!\n", o.DataChannel.Label(), o.DataChannel.ID(), o.DataChannel.ReadyState())

		//for range time.NewTicker(5 * time.Second).C {
		//	message := signal.RandSeq(15)
		//	fmt.Printf("Sending '%s'\n", message)
		//
		//	// Send the message as text
		//	sendErr := o.DataChannel.SendText(message)
		//	if sendErr != nil {
		//		panic(sendErr)
		//	}
		//}
	})
	o.DataChannel.OnClose(func() {
		fmt.Println("DataChannel has closed")
	})
	o.DataChannel.OnMessage(func(msg webrtc.DataChannelMessage) {
		fmt.Printf("Message from DataChannel '%s': '%s'\n", o.DataChannel.Label(), string(msg.Data))
	})

	o.PeerConnection.OnICECandidate(func(c *webrtc.ICECandidate) {
		if c == nil {
			return
		}

		message := WebSocketMessage{
			Service: "Candidate",
			Data: c.ToJSON(),
		}
		b, _ := message.Encode()

		WebSocketHub.Broadcast(b)
	})
	o.PeerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		fmt.Printf("ICE Connection State has changed: %s\n", connectionState.String())
	})
	return
}

func (o *WebRTC) CreateOffer() (err error) {
	// Create an offer to send to the browser
	offer, err := o.PeerConnection.CreateOffer(nil)
	if err != nil {
		panic(err)
	}

	// Sets the LocalDescription, and starts our UDP listeners
	err = o.PeerConnection.SetLocalDescription(offer)
	if err != nil {
		panic(err)
	}

	b, _ := json.Marshal(offer)

	o.Offer = string(b)

	return
}

func (o *WebRTC) ReceiveAnswer(rs string) (err error) {
	answer := webrtc.SessionDescription{}

	err = json.Unmarshal([]byte(rs), &answer)
	if err != nil {
		panic(err)
	}
	// Apply the answer as the remote description
	err = o.PeerConnection.SetRemoteDescription(answer)
	if err != nil {
		panic(err)
	}
	return
}

func (o *WebRTC) AddCandidate(rs string) (err error) {
	if err = o.PeerConnection.AddICECandidate(webrtc.ICECandidateInit{Candidate: rs}); err != nil {
		panic(err)
	}
	return
}