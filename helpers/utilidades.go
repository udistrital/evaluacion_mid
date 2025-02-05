package helpers

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/utils_oas/xray"
)

func sendJson(url string, trequest string, target interface{}, datajson interface{}) error {
	b := new(bytes.Buffer)
	if datajson != nil {
		if err := json.NewEncoder(b).Encode(datajson); err != nil {
			beego.Error(err)
		}
		fmt.Println("Body JSON enviado:", b.String())
	}
	client := &http.Client{}
	req, err := http.NewRequest(trequest, url, b)
	r, err := client.Do(req)
	if err != nil {
		beego.Error("error", err)
		return err
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			beego.Error(err)
		}
	}()

	return json.NewDecoder(r.Body).Decode(target)
}

func SendJson2(url string, trequest string, target interface{}, datajson interface{}) error {
	b := new(bytes.Buffer)
	if datajson != nil {
		if err := json.NewEncoder(b).Encode(datajson); err != nil {
			beego.Error(err)
		}
	}

	client := &http.Client{}
	req, _ := http.NewRequest(trequest, url, b)
	seg := xray.BeginSegmentSec(req)
	defer func() {
		//Catch
		if r := recover(); r != nil {
			client := &http.Client{}
			resp, err := client.Do(req)
			xray.UpdateSegment(resp, err, seg)
			if err != nil {
				beego.Error("Error reading response. ", err)
			}

			defer resp.Body.Close()
			mensaje, err := io.ReadAll(resp.Body)
			if err != nil {
				beego.Error("Error converting response. ", err)
			}
			bodyreq, err := io.ReadAll(req.Body)
			if err != nil {
				beego.Error("Error converting response. ", err)
			}
			respuesta := map[string]interface{}{"request": map[string]interface{}{"url": req.URL.String(), "header": req.Header, "body": bodyreq}, "body": mensaje, "statusCode": resp.StatusCode, "status": resp.Status}
			e, err := json.Marshal(respuesta)
			if err != nil {
				logs.Error(err)
			}
			json.Unmarshal(e, &target)
		}
	}()

	req.Header.Set("Authorization", "")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("accept", "*/*")
	r, err := client.Do(req)
	xray.UpdateSegment(r, err, seg)
	if err != nil {
		beego.Error("error", err)
		return err
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			beego.Error(err)
		}
	}()

	fmt.Println("Respuesta del servidor:", json.NewDecoder(r.Body).Decode(target))

	return json.NewDecoder(r.Body).Decode(target)
}

func sendJsonAutenticacion(url string, trequest string, target interface{}, datajson interface{}) error {
	b := new(bytes.Buffer)
	if datajson != nil {
		if err := json.NewEncoder(b).Encode(datajson); err != nil {
			beego.Error("Error al codificar el JSON:", err)
			return err
		}
		fmt.Println("Body JSON enviado:", b.String())
	}

	client := &http.Client{}
	req, err := http.NewRequest(trequest, url, b)
	if err != nil {
		beego.Error("Error al crear la solicitud:", err)
		return err
	}
	// Agregar encabezados necesarios
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	r, err := client.Do(req)
	if err != nil {
		beego.Error("Error en la solicitud:", err)
		return err
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			beego.Error("Error al cerrar el cuerpo de la respuesta:", err)
		}
	}()

	// Leer y mostrar la respuesta del servidor
	bodyBytes, _ := ioutil.ReadAll(r.Body)

	// Decodificar JSON en el objeto de destino
	if err := json.Unmarshal(bodyBytes, target); err != nil {
		beego.Error("Error al decodificar el JSON:", err)
		return err
	}

	return nil
}

func sendJsonWithToken(url string, trequest string, target interface{}, datajson interface{}, token string) error {
	b := new(bytes.Buffer)
	if datajson != nil {
		if err := json.NewEncoder(b).Encode(datajson); err != nil {
			beego.Error("Error al codificar el JSON:", err)
			return err
		}
	}

	client := &http.Client{}
	req, err := http.NewRequest(trequest, url, b)
	if err != nil {
		beego.Error("Error al crear la solicitud:", err)
		return err
	}

	// Agregar encabezados necesarios
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	r, err := client.Do(req)
	if err != nil {
		beego.Error("Error en la solicitud:", err)
		return err
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			beego.Error("Error al cerrar el cuerpo de la respuesta:", err)
		}
	}()

	// Leer y mostrar la respuesta del servidor
	bodyBytes, _ := ioutil.ReadAll(r.Body)

	// Decodificar JSON en el objeto de destino
	if err := json.Unmarshal(bodyBytes, target); err != nil {
		beego.Error("Error al decodificar el JSON:", err)
		return err
	}

	return nil
}

func getJsonTest(url string, target interface{}) (status int, err error) {
	r, err := http.Get(url)
	if err != nil {
		return r.StatusCode, err
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			beego.Error(err)
		}
	}()

	return r.StatusCode, json.NewDecoder(r.Body).Decode(target)
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			beego.Error(err)
		}
	}()

	return json.NewDecoder(r.Body).Decode(target)
}

// func getJsonTest(w http.ResponseWriter,r *http.Request){
// 	err := r.ParseForm()
// 	if err != nil {
// 	   log.Fatal("parse form error ",err)
// 	}
// 	// 初始化请求变量结构
// 	formData := make(map[string]interface{})
// 	// 调用json包的解析，解析请求body
// 	json.NewDecoder(r.Body).Decode(&formData)
// 	for key,value := range formData{
// 	   log.Println("key:",key," => value :",value)
// 	}
//  }

func getXml(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			beego.Error(err)
		}
	}()

	return xml.NewDecoder(r.Body).Decode(target)
}

func getJsonWSO2(urlp string, target interface{}) error {
	b := new(bytes.Buffer)
	client := &http.Client{}
	req, err := http.NewRequest("GET", urlp, b)
	req.Header.Set("Accept", "application/json")
	r, err := client.Do(req)
	if err != nil {
		beego.Error("error", err)
		return err
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			beego.Error(err)
		}
	}()

	return json.NewDecoder(r.Body).Decode(target)
}

func getJsonWSO2Test(urlp string, target interface{}) (status int, err error) {
	b := new(bytes.Buffer)
	client := &http.Client{}
	req, err := http.NewRequest("GET", urlp, b)
	req.Header.Set("Accept", "application/json")
	r, err := client.Do(req)
	if err != nil {
		beego.Error("error", err)
		return r.StatusCode, err
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			beego.Error(nil, err)
		}
	}()

	return r.StatusCode, json.NewDecoder(r.Body).Decode(target)
}

func diff(a, b time.Time) (year, month, day int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	oneDay := time.Hour * 5
	a = a.Add(oneDay)
	b = b.Add(oneDay)
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)

	// Normalize negative values

	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}

func FormatMoney(value interface{}, Precision int) string {
	formattedNumber := FormatNumber(value, Precision, ",", ".")
	return FormatMoneyString(formattedNumber, Precision)
}

func FormatMoneyString(formattedNumber string, Precision int) string {
	var format string

	zero := "0"
	if Precision > 0 {
		zero += "." + strings.Repeat("0", Precision)
	}

	format = "%s%v"
	result := strings.Replace(format, "%s", "$", -1)
	result = strings.Replace(result, "%v", formattedNumber, -1)

	return result
}

func FormatNumber(value interface{}, precision int, thousand string, decimal string) string {
	v := reflect.ValueOf(value)
	var x string
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		x = fmt.Sprintf("%d", v.Int())
		if precision > 0 {
			x += "." + strings.Repeat("0", precision)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		x = fmt.Sprintf("%d", v.Uint())
		if precision > 0 {
			x += "." + strings.Repeat("0", precision)
		}
	case reflect.Float32, reflect.Float64:
		x = fmt.Sprintf(fmt.Sprintf("%%.%df", precision), v.Float())
	case reflect.Ptr:
		switch v.Type().String() {
		case "*big.Rat":
			x = value.(*big.Rat).FloatString(precision)

		default:
			panic("Unsupported type - " + v.Type().String())
		}
	default:
		panic("Unsupported type - " + v.Kind().String())
	}

	return formatNumberString(x, precision, thousand, decimal)
}

func formatNumberString(x string, precision int, thousand string, decimal string) string {
	lastIndex := strings.Index(x, ".") - 1
	if lastIndex < 0 {
		lastIndex = len(x) - 1
	}

	var buffer []byte
	var strBuffer bytes.Buffer

	j := 0
	for i := lastIndex; i >= 0; i-- {
		j++
		buffer = append(buffer, x[i])

		if j == 3 && i > 0 && !(i == 1 && x[0] == '-') {
			buffer = append(buffer, ',')
			j = 0
		}
	}

	for i := len(buffer) - 1; i >= 0; i-- {
		strBuffer.WriteByte(buffer[i])
	}
	result := strBuffer.String()

	if thousand != "," {
		result = strings.Replace(result, ",", thousand, -1)
	}

	extra := x[lastIndex+1:]
	if decimal != "." {
		extra = strings.Replace(extra, ".", decimal, 1)
	}

	return result + extra
}

func CrearQueryContratoGeneral(proveedor, numeroContrato, vigencia, supervisor, tipoContrato string) string {

	var query []string
	if proveedor != "0" {
		query = append(query, "Contratista:"+proveedor)
	}

	if numeroContrato != "0" {
		query = append(query, "ContratoSuscrito__NumeroContratoSuscrito:"+numeroContrato)
	}

	if vigencia != "0" {
		query = append(query, "VigenciaContrato:"+vigencia)
	}

	if supervisor != "0" {
		query = append(query, "Supervisor__Documento:"+supervisor)
	}

	if tipoContrato != "" {
		if strings.HasPrefix(tipoContrato, "notin:") || strings.HasPrefix(tipoContrato, "in:") {
			prefix := strings.SplitN(tipoContrato, ":", 2)
			tipoContrato = "__" + prefix[0] + ":" + url.QueryEscape(prefix[1])
		} else {
			tipoContrato = ":" + url.QueryEscape(tipoContrato)
		}

		query = append(query, "TipoContrato__TipoContrato"+tipoContrato)
	}

	query_ := strings.Join(query, ",")
	query_ = query_ + "&limit=0"
	return query_
}

func CrearQueryNovedadesCesion(proveedor, numeroContrato, vigencia string) string {

	var query = []string{"IdTipoPropiedad__Nombre:Cesionario", "IdNovedadesPoscontractuales__TipoNovedad:2"}
	if proveedor != "0" {
		query = append(query, "Propiedad:"+proveedor)
	}

	if numeroContrato != "0" {
		query = append(query, "IdNovedadesPoscontractuales__ContratoId:"+numeroContrato)
	}

	if vigencia != "0" {
		query = append(query, "IdNovedadesPoscontractuales__Vigencia:"+vigencia)
	}

	query_ := strings.Join(query, ",")
	return query_
}

// Funcion para agregar los datos a un slice
func StringToSlice(cadena string) (slice []string) {
	parts := strings.Split(cadena, ",")

	if cadena != "" {
		for _, part := range parts {
			slice = append(slice, part)
		}
	}
	return slice
}

func LimpiezaRespuestaRefactor(respuesta map[string]interface{}, v interface{}) {
	b, err := json.Marshal(respuesta["Data"])
	if err != nil {
		panic(err)
	}

	json.Unmarshal(b, v)
}
