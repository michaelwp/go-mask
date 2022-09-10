package go_mask

import (
	"fmt"
	"testing"
)

func TestMaskingString(t *testing.T) {
	name := "ANTON BOY"
	phoneNumber := "081234567890"
	address := "j"
	postalCode := ""

	sentence := fmt.Sprintf(
		"my name is %s, i live at %s - %s, here is my phone number %s",
		name, address, postalCode, phoneNumber,
	)

	type args struct {
		sentence *string
		opt      []*Option
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "masking string",
			args: args{
				sentence: &sentence,
				opt: []*Option{
					{
						Mask: &Mask{
							Char:    "*",
							Length:  len(name) - 1,
							Prepend: string(name[0]),
							Append:  string(name[len(name)-1]),
						},
						Word: &name,
					},
					{
						Mask: &Mask{
							Char:    "?",
							Length:  len(phoneNumber) - 2,
							Prepend: phoneNumber[:2],
						},
						Word: &phoneNumber,
					},
					{
						Mask: &Mask{
							Char:   "#",
							Length: len(address),
						},
						Word: &address,
					},
					{
						Mask: &Mask{
							Char:   "x",
							Length: len(postalCode),
						},
						Word: &postalCode,
					},
				}},
		},
		{
			name: "masking string nil",
			args: args{sentence: nil, opt: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MaskingString(tt.args.sentence, tt.args.opt...)
		})

		fmt.Println()

		fmt.Println("name:", name)
		fmt.Println("phone number:", phoneNumber)
		fmt.Println("address:", address)
		fmt.Println("postal code:", postalCode)
		fmt.Println(sentence)
	}
}

func TestMaskingJSON(t *testing.T) {
	jVA := JSON(`{"bca": "3971294005314411455", "bri": "103908 202 4499019", "cimb": "422903 202 4499019", "permata": "877308 202 4499019", "alfamart": "317060 202 4499019", "bukalapak": "1020319094000716", "indomaret": "317066 202 4499019", "tokopedia": "1010319094000716", "mandiri_open": "94005314411455"}`)

	type args struct {
		jsonData []*JSON
		opt      []*Option
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "masking json va nil keys",
			args: args{
				jsonData: []*JSON{&jVA},
				opt:      nil,
			},
			wantErr: false,
		},
		{
			name: "masking json va",
			args: args{
				jsonData: []*JSON{&jVA},
				opt: []*Option{
					{Json: &Json{Key: "bca"}},
					{Json: &Json{Key: "bri"}},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MaskingJSON(tt.args.jsonData, tt.args.opt...); (err != nil) != tt.wantErr {
				t.Errorf("MaskingJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		fmt.Println(string(jVA))
	}
}

func TestMaskingJSONSlice(t *testing.T) {
	jContacts := JSON(`[{"name": "GILANG CAHYA PRATAMA", "type": "USER", "number": "085314411455", "source": "REGISTRATION"}, {"name": "GILANG CAHYA PRATAMA", "type": "REGISTRATION", "number": "085314411455", "source": null}, {"name": "GILANG CAHYA PRATAMA", "type": "USER_VERIFIED", "number": "085314411455", "source": null}]`)
	jAddresses := JSON(`[{"type": "GPS", "address": "Perumahan Bumi Mulyo Permai Blok E2 No. 1, Sambiroto, Karangtanjung, Kec. Candi, Kabupaten Sidoarjo, Jawa Timur 61272, Indonesia . Location details : -", "latitude": "-7.4899863", "longitude": "112.6939529", "postal_code": "61272", "area_level_1": null, "area_level_2": null, "area_level_3": null, "area_level_4": null, "area_level_5": null}, {"type": "GPS", "address": "Jl. Raya Tandes Lor No.15, Tandes, Kec. Tandes, Kota SBY, Jawa Timur 60187, Indonesia . Location details : -", "latitude": "-7.2610123", "longitude": "112.68772179999999", "postal_code": "60187", "area_level_1": null, "area_level_2": null, "area_level_3": null, "area_level_4": null, "area_level_5": null}, {"type": "GPS", "address": "Perumahan Bumi Mulyo Permai Blok E2 No. 1, Sambiroto, Karangtanjung, Kec. Candi, Kabupaten Sidoarjo, Jawa Timur 61272, Indonesia . Location details : -", "latitude": "-7.4899863", "longitude": "112.6939529", "postal_code": "61272", "area_level_1": null, "area_level_2": null, "area_level_3": null, "area_level_4": null, "area_level_5": null}, {"type": "LEGAL", "address": "GANG V, 001/004, GOSARI, UJUNG PANGKAH, GRESIK, JAWA TIMUR, 61154. Location details : -", "latitude": null, "longitude": null, "postal_code": "61154", "area_level_1": "JAWA TIMUR", "area_level_2": "GRESIK", "area_level_3": "UJUNG PANGKAH", "area_level_4": "GOSARI", "area_level_5": "001/004"}, {"type": "OUTSIDE_AREA", "address": "Gg. X, Gosari, Ujungpangkah, Kabupaten Gresik, Jawa Timur 61154, Indonesia. Location details : -", "latitude": null, "longitude": null, "postal_code": "61154", "area_level_1": "JAWA TIMUR", "area_level_2": "GRESIK", "area_level_3": "UJUNG PANGKAH", "area_level_4": "GOSARI", "area_level_5": null}, {"type": "PEFINDO", "address": "banten tangerang sepatan karet, tangerang, kab.. Location details : -", "latitude": null, "longitude": null, "postal_code": null, "area_level_1": null, "area_level_2": null, "area_level_3": null, "area_level_4": null, "area_level_5": null}, {"type": "RESIDENCE", "address": "Jl. Raya Tandes Lor No.11A, Tandes, Kec. Tandes, Kota SBY, Jawa Timur 60186, Indonesia. Location details : -", "latitude": null, "longitude": null, "postal_code": "60186", "area_level_1": "JAWA TIMUR", "area_level_2": "SURABAYA", "area_level_3": "TANDES", "area_level_4": "BALONGSARI", "area_level_5": null}, {"type": "USER", "address": "jl genteng, KETABANG, GENTENG, SURABAYA, JAWA TIMUR, 60272. Location details : rt4 rw3", "latitude": "-7.2590879", "longitude": "112.7479862", "postal_code": "60272", "area_level_1": "JAWA TIMUR", "area_level_2": "SURABAYA", "area_level_3": "GENTENG", "area_level_4": "KETABANG", "area_level_5": null}, {"type": "USER", "address": "Jl. Raya Tandes Lor No.11A, Tandes, Kec. Tandes, Kota SBY, Jawa Timur 60186, Indonesia. Location details : -", "latitude": "-7.2587319999999993", "longitude": "112.683803", "postal_code": "60186", "area_level_1": "JAWA TIMUR", "area_level_2": "SURABAYA", "area_level_3": "TANDES", "area_level_4": "BALONGSARI", "area_level_5": null}, {"type": "USER", "address": "Jl. Raya Tandes Lor No.11A, Tandes, Kec. Tandes, Kota SBY, Jawa Timur 60186, Indonesia. Location details : tandes lor", "latitude": "-7.2587319999999993", "longitude": "112.683803", "postal_code": "60186", "area_level_1": "JAWA TIMUR", "area_level_2": "SURABAYA", "area_level_3": "TANDES", "area_level_4": "BALONGSARI", "area_level_5": null}, {"type": "USER", "address": "jl genteng, KETABANG, GENTENG, SURABAYA, JAWA TIMUR, 60272. Location details : rt4 rw3", "latitude": "-7.2590879", "longitude": "112.7479862", "postal_code": "60272", "area_level_1": "JAWA TIMUR", "area_level_2": "SURABAYA", "area_level_3": "GENTENG", "area_level_4": "KETABANG", "area_level_5": null}, {"type": "USER", "address": "Jl. Buntaran No.2, Manukan Wetan, Kec. Tandes, Kota SBY, Jawa Timur 60184, Indonesia. Location details : gang v", "latitude": "-7.2544204999999993", "longitude": "112.6738735", "postal_code": "60184", "area_level_1": "JAWA TIMUR", "area_level_2": "SURABAYA", "area_level_3": "ASEMROWO", "area_level_4": "TAMBAK LANGON", "area_level_5": null}, {"type": "USER", "address": "Gang V, RT.1/RW.4, Gosari, Ujung Pangkah  Rt 1 rw 4 , KAB. GRESIK, UJUNG PANGKAH, JAWA TIMUR, ID, 61154. Location details : Rt 1 Rw 4 Gosari", "latitude": "0", "longitude": "0", "postal_code": "61154", "area_level_1": null, "area_level_2": null, "area_level_3": null, "area_level_4": null, "area_level_5": null}, {"type": "USER", "address": "jl genteng, KETABANG, GENTENG, SURABAYA, JAWA TIMUR, 60272. Location details : rt4 rw3", "latitude": "-7.2590879", "longitude": "112.7479862", "postal_code": "60272", "area_level_1": "JAWA TIMUR", "area_level_2": "SURABAYA", "area_level_3": "GENTENG", "area_level_4": "KETABANG", "area_level_5": null}, {"type": "USER", "address": "rt 1 rw 4, KETABANG, GENTENG, SURABAYA, JAWA TIMUR, 60272. Location details : rt 1 rw 4", "latitude": "-7.2590879", "longitude": "112.7479862", "postal_code": "60272", "area_level_1": "JAWA TIMUR", "area_level_2": "SURABAYA", "area_level_3": "GENTENG", "area_level_4": "KETABANG", "area_level_5": null}, {"type": "USER", "address": "Gg. X, Gosari, Ujungpangkah, Kabupaten Gresik, Jawa Timur 61154, Indonesia. Location details : rt 1 rw 4", "latitude": "-6.9364785", "longitude": "112.5103968", "postal_code": "61154", "area_level_1": "JAWA TIMUR", "area_level_2": "GRESIK", "area_level_3": "UJUNG PANGKAH", "area_level_4": "GOSARI", "area_level_5": null}, {"type": "USER", "address": "Gg. X, Gosari, Ujungpangkah, Kabupaten Gresik, Jawa Timur 61154, Indonesia. Location details : rt 01 rw 04", "latitude": "-6.9364785", "longitude": "112.5103968", "postal_code": "61154", "area_level_1": "JAWA TIMUR", "area_level_2": "GRESIK", "area_level_3": "UJUNG PANGKAH", "area_level_4": "GOSARI", "area_level_5": null}]`)

	type args struct {
		jsonData []*JSON
		opt      []*Option
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		//{
		//	name: "masking slice json nil keys",
		//	args: args{
		//		jsonData: []*JSON{&jContacts, &jAddresses},
		//		opt:      nil,
		//	},
		//	wantErr: false,
		//},
		{
			name: "masking slice json",
			args: args{
				jsonData:    []*JSON{&jContacts, &jAddresses},
				opt: []*Option{
					{Json: &Json{Key: "number"}},
					{Json: &Json{Key: "address"}},
					{Json: &Json{Key: "latitude"}},
					{Json: &Json{Key: "longitude"}},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MaskingJSONSlice(tt.args.jsonData, tt.args.opt...); (err != nil) != tt.wantErr {
				t.Errorf("MaskingSliceJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		fmt.Println("contacts: ", string(jContacts))
		fmt.Println("address: ", string(jAddresses))
	}
}
