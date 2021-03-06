package vcard

import (
	"strings"
	"reflect"
	"testing"
)

// RFC
var testCardString = `BEGIN:VCARD
VERSION:4.0
UID:urn:uuid:4fbe8971-0bc3-424c-9c26-36c3e1eff6b1
FN;PID=1.1:J. Doe
N:Doe;J.;;;
EMAIL;PID=1.1:jdoe@example.com
CLIENTPIDMAP:1;urn:uuid:53e374d9-337e-4727-8803-a1e9c14e0556
END:VCARD`

var testCardHandmadeString = `BEGIN:VCARD
VERSION:4.0
N:Bloggs;Joe;;;
FN:Joe Bloggs
EMAIL;TYPE=home;PREF=1:me@joebloggs.com
TEL;TYPE="cell,home";PREF=1:tel:+44 20 1234 5678
ADR;TYPE=home;PREF=1:;;1 Trafalgar Square;London;;WC2N;United Kingdom
URL;TYPE=home;PREF=1:http://joebloggs.com
IMPP;TYPE=home;PREF=1:skype:joe.bloggs
X-SOCIALPROFILE;TYPE=home;PREF=1:twitter:https://twitter.com/joebloggs
END:VCARD`

// Google Contacts (15 November 2012)
var testCardGoogleString = `BEGIN:VCARD
VERSION:3.0
N:Bloggs;Joe;;;
FN:Joe Bloggs
EMAIL;TYPE=INTERNET;TYPE=HOME:me@joebloggs.com
TEL;TYPE=CELL:+44 20 1234 5678
ADR;TYPE=HOME:;;1 Trafalgar Square;London;;WC2N;United Kingdom
item1.URL:http\://joebloggs.com
item1.X-ABLabel:_$!<HomePage>!$_
X-SKYPE:joe.bloggs
item2.URL:http\://twitter.com/test
item2.X-ABLabel:Twitter
END:VCARD`

// Apple Contacts (version 7.1)
var testCardAppleString = `BEGIN:VCARD
VERSION:3.0
N:Bloggs;Joe;;;
FN:Joe Bloggs
EMAIL;type=INTERNET;type=HOME;type=pref:me@joebloggs.com
TEL;type=CELL;type=VOICE;type=pref:+44 20 1234 5678
ADR;type=HOME;type=pref:;;1 Trafalgar Square;London;;WC2N;United Kingdom
item1.URL;type=pref:http://joebloggs.com
item1.X-ABLabel:_$!<HomePage>!$_
IMPP;X-SERVICE-TYPE=Skype;type=HOME;type=pref:skype:joe.bloggs
X-SOCIALPROFILE;type=twitter:https://twitter.com/joebloggs
END:VCARD`

var decoderTests = []struct{
	s string
	card Card
}{
	{testCardString, testCard},
	{testCardHandmadeString, testCardHandmade},
	{testCardGoogleString, testCardGoogle},
	{testCardAppleString, testCardApple},
}

func TestDecoder(t *testing.T) {
	for _, test := range decoderTests {
		r := strings.NewReader(test.s)
		dec := NewDecoder(r)
		card, err := dec.Decode()
		if err != nil {
			t.Fatal("Expected no error when decoding card, got:", err)
		}
		if !reflect.DeepEqual(card, test.card) {
			t.Errorf("Invalid parsed card: expected \n%+v\n but got \n%+v", test.card, card)
		}
	}
}
