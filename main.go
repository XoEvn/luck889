/**
 * @Author: evnxo
 * @Description:
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2023/12/25 15:35
 */

package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.POST("/MSVACQWEB/GatewayServlet/payment", Payment)
	router.POST("/MSVACQWEB/GatewayServlet/refund", Refund)
	router.POST("/MSVACQWEB/GatewayServlet/inquiry", Inquiry)
	router.POST("/MSVACQWEB/GatewayServlet/refundInquiry", RefundInquiry)
	server := http.Server{
		Addr:    ":8080",
		Handler: myHost(router),
	}
	server.ListenAndServe()
}

type MSV_RESP struct {
	Message struct {
		Content struct {
			GatewayName  string `xml:"GATEWAY_NAME"`
			ApiCode      string `xml:"API_CODE"`
			MchId        string `xml:"MCH_ID"`
			DeviceInfo   string `xml:"DEVICE_INFO"`
			NonceStr     string `xml:"NONCE_STR"`
			FundPayType  string `xml:"FUND_PAY_TYPE"`
			Currency     string `xml:"CURRENCY"`
			TotalAmount  string `xml:"TOTAL_AMOUNT"`
			CashCurrency string `xml:"CASH_CURRENCY"`
			CashAmount   string `xml:"CASH_AMOUNT"`
			ExchangeRate string `xml:"EXCHANGE_RATE"`
			OrderNo      string `xml:"ORDER_NO"`
			MerOrderNo   string `xml:"MER_ORDER_NO"`
			PayerInfo    string `xml:"PAYER_INFO"`
			EndTime      string `xml:"END_TIME"`
			TranFlag     string `xml:"TRAN_FLAG"`
			RespCode     string `xml:"RESP_CODE"`
		} `xml:"CONTENT"`
	} `xml:"MESSAGE"`
	SignData string `xml:"SIGN_DATA"`
}

func Payment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	xmlStr := `<MSV_RESP><MESSAGE><CONTENT><GATEWAY_NAME>CIL</GATEWAY_NAME><API_CODE>MSV.ABC.PAYMENT</API_CODE><MCH_ID>30032258120000067902175</MCH_ID><DEVICE_INFO></DEVICE_INFO><NONCE_STR>54264863303646238964883286923530</NONCE_STR><FUND_PAY_TYPE></FUND_PAY_TYPE><CURRENCY>HKD</CURRENCY><TOTAL_AMOUNT>1.00</TOTAL_AMOUNT><CASH_CURRENCY>CNY</CASH_CURRENCY><CASH_AMOUNT>0.95</CASH_AMOUNT><EXCHANGE_RATE>1.0501941</EXCHANGE_RATE><ORDER_NO>175202312261437030000444916</ORDER_NO><MER_ORDER_NO>2312261437010001053158301</MER_ORDER_NO><PAYER_INFO>0041000056120190</PAYER_INFO><END_TIME>20231226143703</END_TIME><TRAN_FLAG>0002</TRAN_FLAG><RESP_CODE>00000003</RESP_CODE></CONTENT></MESSAGE><SIGN_DATA>MXKnR3pH4/0yIpY3MrnsP7FTaDy6z/RO0sp39FShasvGcrwXe6R5SD/gpq/z2s46TgmGcbHDP3HXMw2ed3k0GtEFXAYJ4kH7C5HurtCWDZO8KOYCZU3UJNx5JZ1Tov7Ivsa17tkBLikeQHtQPH9khuB73mxgceZ2bQVHH3m80AO4l0Mp8sV2Z6IteTf7VBdDgEck0fHK+JaaYsmDBxcSJouckRgUkyQhzKR1eD+9SgJ+q+ABYogQxL42Z9enExXod2FBktJIZkwcEUYDbp1v3+on64X8sZX10YLXUv3XKlffSVhWOYoSBFPCWJfkmC5JoUgaoOb0ROnxrzSJ0Xsz6Q==</SIGN_DATA></MSV_RESP>`
	var resp MSV_RESP
	err := xml.Unmarshal([]byte(xmlStr), &resp)
	output, err := xml.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/xml")
	//w.Write([]byte(xml.Header))
	w.Write(output)
}

func Refund(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	xmlStr := `<MSV_RESP><MESSAGE><CONTENT><GATEWAY_NAME>CIL</GATEWAY_NAME><API_CODE>MSV.ABC.PAYMENT</API_CODE><MCH_ID>30032258120000067902175</MCH_ID><DEVICE_INFO></DEVICE_INFO><NONCE_STR>54264863303646238964883286923530</NONCE_STR><FUND_PAY_TYPE></FUND_PAY_TYPE><CURRENCY>HKD</CURRENCY><TOTAL_AMOUNT>1.00</TOTAL_AMOUNT><CASH_CURRENCY>CNY</CASH_CURRENCY><CASH_AMOUNT>0.95</CASH_AMOUNT><EXCHANGE_RATE>1.0501941</EXCHANGE_RATE><ORDER_NO>175202312261437030000444916</ORDER_NO><MER_ORDER_NO>2312261437010001053158301</MER_ORDER_NO><PAYER_INFO>0041000056120190</PAYER_INFO><END_TIME>20231226143703</END_TIME><TRAN_FLAG>0002</TRAN_FLAG><RESP_CODE>00000003</RESP_CODE></CONTENT></MESSAGE><SIGN_DATA>MXKnR3pH4/0yIpY3MrnsP7FTaDy6z/RO0sp39FShasvGcrwXe6R5SD/gpq/z2s46TgmGcbHDP3HXMw2ed3k0GtEFXAYJ4kH7C5HurtCWDZO8KOYCZU3UJNx5JZ1Tov7Ivsa17tkBLikeQHtQPH9khuB73mxgceZ2bQVHH3m80AO4l0Mp8sV2Z6IteTf7VBdDgEck0fHK+JaaYsmDBxcSJouckRgUkyQhzKR1eD+9SgJ+q+ABYogQxL42Z9enExXod2FBktJIZkwcEUYDbp1v3+on64X8sZX10YLXUv3XKlffSVhWOYoSBFPCWJfkmC5JoUgaoOb0ROnxrzSJ0Xsz6Q==</SIGN_DATA></MSV_RESP>`
	var resp MSV_RESP
	err := xml.Unmarshal([]byte(xmlStr), &resp)
	output, err := xml.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/xml")
	//w.Write([]byte(xml.Header))
	w.Write(output)
}
func Inquiry(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	xmlStr := `<MSV_RESP><MESSAGE><CONTENT><GATEWAY_NAME>CIL</GATEWAY_NAME><API_CODE>MSV.ABC.PAYMENT</API_CODE><MCH_ID>30032258120000067902175</MCH_ID><DEVICE_INFO></DEVICE_INFO><NONCE_STR>54264863303646238964883286923530</NONCE_STR><FUND_PAY_TYPE></FUND_PAY_TYPE><CURRENCY>HKD</CURRENCY><TOTAL_AMOUNT>1.00</TOTAL_AMOUNT><CASH_CURRENCY>CNY</CASH_CURRENCY><CASH_AMOUNT>0.95</CASH_AMOUNT><EXCHANGE_RATE>1.0501941</EXCHANGE_RATE><ORDER_NO>175202312261437030000444916</ORDER_NO><MER_ORDER_NO>2312261437010001053158301</MER_ORDER_NO><PAYER_INFO>0041000056120190</PAYER_INFO><END_TIME>20231226143703</END_TIME><TRAN_FLAG>0002</TRAN_FLAG><RESP_CODE>00000003</RESP_CODE></CONTENT></MESSAGE><SIGN_DATA>MXKnR3pH4/0yIpY3MrnsP7FTaDy6z/RO0sp39FShasvGcrwXe6R5SD/gpq/z2s46TgmGcbHDP3HXMw2ed3k0GtEFXAYJ4kH7C5HurtCWDZO8KOYCZU3UJNx5JZ1Tov7Ivsa17tkBLikeQHtQPH9khuB73mxgceZ2bQVHH3m80AO4l0Mp8sV2Z6IteTf7VBdDgEck0fHK+JaaYsmDBxcSJouckRgUkyQhzKR1eD+9SgJ+q+ABYogQxL42Z9enExXod2FBktJIZkwcEUYDbp1v3+on64X8sZX10YLXUv3XKlffSVhWOYoSBFPCWJfkmC5JoUgaoOb0ROnxrzSJ0Xsz6Q==</SIGN_DATA></MSV_RESP>`
	//time.Sleep(30 * time.Second)
	var resp MSV_RESP
	err := xml.Unmarshal([]byte(xmlStr), &resp)
	output, err := xml.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/xml")
	//w.Write([]byte(xml.Header))
	w.Write(output)
}
func RefundInquiry(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	xmlStr := `<MSV_RESP><MESSAGE><CONTENT><GATEWAY_NAME>CIL</GATEWAY_NAME><API_CODE>MSV.ABC.PAYMENT</API_CODE><MCH_ID>30032258120000067902175</MCH_ID><DEVICE_INFO></DEVICE_INFO><NONCE_STR>54264863303646238964883286923530</NONCE_STR><FUND_PAY_TYPE></FUND_PAY_TYPE><CURRENCY>HKD</CURRENCY><TOTAL_AMOUNT>1.00</TOTAL_AMOUNT><CASH_CURRENCY>CNY</CASH_CURRENCY><CASH_AMOUNT>0.95</CASH_AMOUNT><EXCHANGE_RATE>1.0501941</EXCHANGE_RATE><ORDER_NO>175202312261437030000444916</ORDER_NO><MER_ORDER_NO>2312261437010001053158301</MER_ORDER_NO><PAYER_INFO>0041000056120190</PAYER_INFO><END_TIME>20231226143703</END_TIME><TRAN_FLAG>0002</TRAN_FLAG><RESP_CODE>00000003</RESP_CODE></CONTENT></MESSAGE><SIGN_DATA>MXKnR3pH4/0yIpY3MrnsP7FTaDy6z/RO0sp39FShasvGcrwXe6R5SD/gpq/z2s46TgmGcbHDP3HXMw2ed3k0GtEFXAYJ4kH7C5HurtCWDZO8KOYCZU3UJNx5JZ1Tov7Ivsa17tkBLikeQHtQPH9khuB73mxgceZ2bQVHH3m80AO4l0Mp8sV2Z6IteTf7VBdDgEck0fHK+JaaYsmDBxcSJouckRgUkyQhzKR1eD+9SgJ+q+ABYogQxL42Z9enExXod2FBktJIZkwcEUYDbp1v3+on64X8sZX10YLXUv3XKlffSVhWOYoSBFPCWJfkmC5JoUgaoOb0ROnxrzSJ0Xsz6Q==</SIGN_DATA></MSV_RESP>`
	time.Sleep(30 * time.Second)
	var resp MSV_RESP
	err := xml.Unmarshal([]byte(xmlStr), &resp)
	output, err := xml.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	w.Header().Set("Content-Type", "application/xml")
	//w.Write([]byte(xml.Header))
	w.Write(output)
}

func myHost(handler http.Handler) http.Handler {
	ourFunc := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler.ServeHTTP(w, r)
		log.Println(
			fmt.Sprintf("%s %s %s", r.Method, r.URL, time.Now().Sub(start)))
	}
	return http.HandlerFunc(ourFunc)
}
