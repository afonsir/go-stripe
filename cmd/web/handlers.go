package main

import "net/http"

func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["publishable_key"] = app.config.stripe.key

	if err := app.renderTemplate(w, r, "terminal", &templateData{
		StringMap: stringMap,
	}); err != nil {
		app.errorLog.Println(err)
	}
}

func (app *application) PaymentSucceeded(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	// read posted data
	cardholderName := r.Form.Get("cardholder_name")
	email := r.Form.Get("cardholder_email")
	paymentIntent := r.Form.Get("payment_intent")
	paymentMethod := r.Form.Get("payment_method")
	paymentAmount := r.Form.Get("payment_amount")
	paymentCurrency := r.Form.Get("payment_currency")

	data := make(map[string]interface{})

	data["cardholder_name"] = cardholderName
	data["cardholder_email"] = email
	data["payment_intent"] = paymentIntent
	data["payment_method"] = paymentMethod
	data["payment_amount"] = paymentAmount
	data["payment_currency"] = paymentCurrency

	if err := app.renderTemplate(w, r, "succeeded", &templateData{
		Data: data,
	}); err != nil {
		app.errorLog.Println(err)
	}
}

// ChargeOnce displays the page to buy one widget
func (app *application) ChargeOnce(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["publishable_key"] = app.config.stripe.key

	if err := app.renderTemplate(w, r, "buy-once", &templateData{
		StringMap: stringMap,
	}, "stripe-js"); err != nil {
		app.errorLog.Println(err)
	}
}
