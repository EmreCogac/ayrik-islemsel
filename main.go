package main

import "fmt"

func main() {

	var isi int
	var isik int
	var nem int
	var sistemArızaDurumu string

	fmt.Println("Isı giriniz")
	fmt.Scan(&isi)
	fmt.Println("Işık giriniz")
	fmt.Scan(&isik)
	fmt.Println("Nem giriniz")
	fmt.Scan(&nem)
	fmt.Println("Sistem durumu giriniz çalışıyor ise 1 çalışmıyor ise 0")
	fmt.Scan(&sistemArızaDurumu)
	out := hesaplaWithIfCondition(isi, isik, nem, sistemArızaDurumu)
	println(out)
	out1 := hesaplaWithSwitchCase(isi, isik, nem, sistemArızaDurumu)
	println(out1)
}
func hesaplaWithSwitchCase(isi int, isik int, nem int, sistemDurumu string) string {
	if sistemDurumu != "0" && sistemDurumu != "1" {
		return "sistem durumu girdisi yalnızca 1 veya 0 içermelidir"
	}
	switch {
	case isi >= 15 && isi <= 20 && isik >= 10 && isik <= 15 && nem >= 5 && nem <= 12:
		return "Sistem doğru şekilde çalışıyor stabil durum" // durum 1
	case isi >= 15 && isi <= 20 && isik >= 10 && isik <= 15 && nem < 5 && sistemDurumu == "0":
		return "Alarm devreye girdi" // durum 2
	case isi >= 15 && isi <= 20 && isik >= 10 && isik <= 15 && nem < 5 && sistemDurumu == "1":
		return "nem arttırıldı" // durum 3
	case isi < 15 && isik >= 10 && isik <= 15 && nem < 5 && sistemDurumu == "1":
		return "ısı arttırıldı" // durum 4
	case isi < 15 && isik >= 10 && isik <= 15 && nem < 5 && sistemDurumu == "0":
		return "alarm çalıştı" // durum 5
	case isi > 20 && isik >= 10 && isik <= 15 && nem < 5 && sistemDurumu == "0":
		return "alarm çalıştı" // durum 6
	case isi > 20 && isik >= 10 && isik <= 15 && nem < 5 && sistemDurumu == "1":
		return "Alarm devreye girdi" // durum 7
	case isi >= 15 && isi <= 20 && isik <= 9 && isik >= 0 && nem >= 5 && nem <= 12 && sistemDurumu == "1":
		return "ışık seviyesi arttır"
	case isi >= 15 && isi <= 20 && isik <= 9 && isik >= 0 && nem >= 5 && nem <= 12 && sistemDurumu == "0":
		return "alarmı çalıştır"
	}
	return "bir sorunla karşılşıldı! tekrar deneyiniz en son durumda developer ile iletişime geçiniz  "
}
func hesaplaWithIfCondition(isi int, isik int, nem int, sistemDurumu string) string {
	if sistemDurumu != "0" && sistemDurumu != "1" {
		return "sistem durumu girdisi yalnızca 1 veya 0 içermelidir"
	}

	if isi >= 15 && isi <= 20 && isik >= 10 && isik <= 15 && nem >= 5 && nem <= 12 {
		return "Sistem doğru şekilde çalışıyor stabil durum" // durum 1
	} else if isi >= 15 && isi <= 20 && isik >= 10 && isik <= 15 && nem < 5 && sistemDurumu == "0" {
		return "Alarm devreye girdi" // durum 2
	} else if isi >= 15 && isi <= 20 && isik >= 10 && isik <= 15 && nem < 5 && sistemDurumu == "1" {
		return "nem arttırıldı" // durum 3
	} else if isi < 15 && isik >= 10 && isik <= 15 && nem < 5 && sistemDurumu == "1" {
		return "ısı arttırıldı" // durum 4
	} else if isi < 15 && isik >= 10 && isik <= 15 && nem < 5 && sistemDurumu == "0" {
		return "alarm çalıştı" // durum 5
	} else if isi > 20 && isik >= 10 && isik <= 15 && nem < 5 && sistemDurumu == "0" {
		return "alarm çalıştı" // durum 6
	} else if isi > 20 && isik >= 10 && isik <= 15 && nem < 5 && sistemDurumu == "1" {
		return "Alarm devreye girdi" // durum 7
	} else if isi >= 15 && isi <= 20 && isik <= 9 && isik >= 0 && nem >= 5 && nem <= 12 && sistemDurumu == "1" {
		return "ışık seviyesi arttır"
	} else if isi >= 15 && isi <= 20 && isik <= 9 && isik >= 0 && nem >= 5 && nem <= 12 && sistemDurumu == "0" {
		return "alarmı çalıştır"
	}
	return "bir sorunla karşılşıldı! tekrar deneyiniz en son durumda developer ile iletişime geçiniz  "
}

/*
durum 1:
ısı , 15 - 20
ışık , 10 - 15
nem , 5 - 12
arızadurumu - null
-> Sistem doğru şekilde çalışıyor stabil durum
------------------------------
durum 2:
ısı , 15 - 20
ışık , 10 - 15
nem  , <5
arızadurumu , true
-> alarmı çalıştır
------------------------------
durum 2 - arızadurumu , false -> durum3
-> nemi arttır
------------------------------
durum 4:
ısı  <15
ışık , 10 - 15
nem  , <5
arızadurumu , false
-> ısı arttır
------------------------------
durum 4 - arızadurumu , true -> durum 5
-> alarmı çalıştır
------------------------------
durum 6
ısı  >20
ışık , 10 - 15
nem  , <5
arızadurumu , true
-> alarmı çalıştır
------------------------------
durum 6 - arızadurumu false -> durum 7
-> ısı azalt
------------------------------
durum 8
ısı , 15 - 20
ışık , 0 - 9
nem , 5 - 12
arızadurumu , false
-> ışık seviyesi arttır
------------------------------
durum 8 - arızadurumu true -> durum 9
-> alarmı çalıştır
*/
