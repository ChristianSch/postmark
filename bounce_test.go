package postmark

import (
	"net/http"
	"testing"

	"goji.io/pat"
)

func TestGetDeliveryStats(t *testing.T) {
	tMux.HandleFunc(pat.Get("/deliverystats"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`{
      "InactiveMails": 192,
      "Bounces": [
        {
          "Name": "All",
          "Count": 253
        },
        {
          "Type": "HardBounce",
          "Name": "Hard bounce",
          "Count": 195
        },
        {
          "Type": "Transient",
          "Name": "Message delayed",
          "Count": 10
        },
        {
          "Type": "AutoResponder",
          "Name": "Auto responder",
          "Count": 14
        },
        {
          "Type": "SpamNotification",
          "Name": "Spam notification",
          "Count": 3
        },
        {
          "Type": "SoftBounce",
          "Name": "Soft bounce",
          "Count": 30
        },
        {
          "Type": "SpamComplaint",
          "Name": "Spam complaint",
          "Count": 1
        }
		]}`))
	})

	res, err := client.GetDeliveryStats()
	if err != nil {
		t.Fatalf("GetDeliveryStats: %s", err.Error())
	}

	if res.InactiveMails != 192 {
		t.Fatalf("GetDeliveryStats: wrong inactive mail count %d", res.InactiveMails)
	}
}

func TestGetBounces(t *testing.T) {
	tMux.HandleFunc(pat.Get("/bounces"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`{
      "TotalCount": 253,
      "Bounces": [
        {
          "ID": 692560173,
          "Type": "HardBounce",
          "TypeCode": 1,
          "Name": "Hard bounce",
          "Tag": "Invitation",
          "MessageID": "2c1b63fe-43f2-4db5-91b0-8bdfa44a9316",
          "Description": "The server was unable to deliver your message (ex: unknown user, mailbox not found).",
          "Details": "action: failed\r\n",
          "Email": "anything@blackhole.postmarkap.com",
          "BouncedAt": "2014-01-15T16:09:19.6421112-05:00",
          "DumpAvailable": false,
          "Inactive": false,
          "CanActivate": true,
          "Subject": "SC API5 Test"
        },
        {
          "ID": 676862817,
          "Type": "HardBounce",
          "TypeCode": 1,
          "Name": "Hard bounce",
          "Tag": "Invitation",
          "MessageID": "623b2e90-82d0-4050-ae9e-2c3a734ba091",
          "Description": "The server was unable to deliver your message (ex: unknown user, mailbox not found).",
          "Details": "smtp;554 delivery error: dd This user doesn't have a yahoo.com account (vicelcown@yahoo.com) [0] - mta1543.mail.ne1.yahoo.com",
          "Email": "vicelcown@yahoo.com",
          "BouncedAt": "2013-10-18T09:49:59.8253577-04:00",
          "DumpAvailable": false,
          "Inactive": true,
          "CanActivate": true,
          "Subject": "Production API Test"
        }
		  ]
    }`))
	})

	_, total, err := client.GetBounces(100, 0, map[string]interface{}{
		"tag": "Invitation",
	})

	if err != nil {
		t.Fatalf("GetBounces: %s", err.Error())
	}

	if total != 253 {
		t.Fatalf("GetBounces: wrong total (%d)", total)
	}
}

func TestGetBounce(t *testing.T) {
	tMux.HandleFunc(pat.Get("/bounces/:bounceID"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`{
      "ID": 692560173,
      "Type": "HardBounce",
      "TypeCode": 1,
      "Name": "Hard bounce",
      "Tag": "Invitation",
      "MessageID": "2c1b63fe-43f2-4db5-91b0-8bdfa44a9316",
      "Description": "The server was unable to deliver your message (ex: unknown user, mailbox not found).",
      "Details": "action: failed\r\n",
      "Email": "anything@blackhole.postmarkap.com",
      "BouncedAt": "2014-01-15T16:09:19.6421112-05:00",
      "DumpAvailable": false,
      "Inactive": false,
      "CanActivate": true,
      "Subject": "SC API5 Test",
      "Content": "Return-Path: <>\r\nReceived: …"
    }`))
	})

	res, err := client.GetBounce(692560173)

	if err != nil {
		t.Fatalf("GetBounces: %s", err.Error())
	}

	if res.ID != 692560173 {
		t.Fatalf("GetBounce: wrong ID (%v)", res.ID)
	}
}

func TestGetBounceDump(t *testing.T) {
	tMux.HandleFunc(pat.Get("/bounces/:bounceID/dump"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`{
	      "Body": "..."
	    }`))
	})

	res, err := client.GetBounceDump(692560173)

	if err != nil {
		t.Fatalf("GetBounceDump: %s", err.Error())
	}

	if res != "..." {
		t.Fatalf("GetBounceDump: wrong dump body (%v)", res)
	}
}
