package models

import (
	"gopkg.in/gomail.v2"
	"math/rand"
	"time"
)

func SendEmail(email string) (string, error) {

	code := generateOTP()

	m := gomail.NewMessage()
	m.SetHeader("From", "benedictvincentrpj7@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Code Verification From Staem")
	m.SetBody("text/html", "    <div style=\"display: flex; justify-content: center; padding: 30px; align-items: center; background-color: #1b2838; width: fit-content;\">\n        <p style=\"color: white; font-weight: bold; font-size: 25px; text-align: center;\">Your Code is "+code+"</p>    \n    </div>")

	d := gomail.NewDialer("smtp.gmail.com", 587, "benedictvincentrpj7@gmail.com", "gpqkwgdbybqghjhc")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	return code, nil
}


func SendEmailTransaction(email string, firstname string, message string, sentiment string, signature string) (bool, error) {

	m := gomail.NewMessage()
	m.SetHeader("From", "benedictvincentrpj7@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Transaction Success")
	m.SetBody("text/html", " <div style=\"padding: 20px 20px; width: 300px; background-color: #1b2838;\">\n        <div style=\"margin-bottom: 20px;\">\n            <p style=\"text-align: center; color: white; font-size: 25px;\">Transaction Success</p>\n        </div>\n\n        <div style=\"color: white; display: flex; justify-content: space-between; align-items: center; margin-bottom: 10px;\">\n            <div><p>first-name: </p></div>\n            <div><p>"+firstname+"</p></div>\n        </div>\n\n        <div style=\"color: white; display: flex; justify-content: space-between; align-items: center; margin-bottom: 10px;\">\n            <div><p>Message: </p></div>\n            <div><p>"+message+"</p></div>\n        </div>\n\n        <div style=\"color: white; display: flex; justify-content: space-between; align-items: center; margin-bottom: 10px;\">\n            <div><p>Sentiment: </p></div>\n            <div><p>"+sentiment+"</p></div>\n        </div>\n\n        <div style=\"color: white; display: flex; justify-content: space-between; align-items: center; margin-bottom: 10px;\">\n            <div><p>Signature: </p></div>\n            <div><p>"+signature+"</p></div>\n        </div>\n\n    </div>")

	d := gomail.NewDialer("smtp.gmail.com", 587, "benedictvincentrpj7@gmail.com", "gpqkwgdbybqghjhc")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	return true, nil
}

func SendEmailWishlist(email string, gameName string) (bool, error) {

	m := gomail.NewMessage()
	m.SetHeader("From", "benedictvincentrpj7@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Game On Sale on Your Wishlist")
	m.SetBody("text/html", "<div style=\"background-color: rgb(23, 46, 80); width: 200px; height: 100px; display: flex; justify-content: center; align-items: center;\">\n        <div>\n            <p style=\"color: white; text-align: center; font-size: 13px; letter-spacing: 2px;\">"+gameName+" is on sale</p>\n        </div>\n    </div>")

	d := gomail.NewDialer("smtp.gmail.com", 587, "benedictvincentrpj7@gmail.com", "gpqkwgdbybqghjhc")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	return true, nil
}

func SendEmailSuccessBuy(email string) (bool, error) {

	m := gomail.NewMessage()
	m.SetHeader("From", "benedictvincentrpj7@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Successfully Buy")
	m.SetBody("text/html", "    <div style=\"display: flex; justify-content: center; padding: 30px; align-items: center; background-color: #1b2838; width: fit-content;\">\n        <p style=\"color: white; font-weight: bold; font-size: 25px; text-align: center;\">Successfully Buy</p>    \n    </div>")

	d := gomail.NewDialer("smtp.gmail.com", 587, "benedictvincentrpj7@gmail.com", "gpqkwgdbybqghjhc")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	return true, nil
}

func SendEmailSuccessSell(email string) (bool, error) {

	m := gomail.NewMessage()
	m.SetHeader("From", "benedictvincentrpj7@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Successfully Sell")
	m.SetBody("text/html", "    <div style=\"display: flex; justify-content: center; padding: 30px; align-items: center; background-color: #1b2838; width: fit-content;\">\n        <p style=\"color: white; font-weight: bold; font-size: 25px; text-align: center;\">Successfully Sold</p>    \n    </div>")

	d := gomail.NewDialer("smtp.gmail.com", 587, "benedictvincentrpj7@gmail.com", "gpqkwgdbybqghjhc")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	return true, nil
}



func generateOTP() string {
	// omitted few confusing characters
	charSet := "ABCDEFGHJKLMNPQRSTUVWXYZ123456789"
	pass := randomStringGenerator(charSet, 5)
	return pass;
}


func randomStringGenerator(charSet string, codeLength int32) string {
	rand.Seed(time.Now().UnixNano())
	code := ""
	charSetLength := int32(len(charSet))
	for i:= int32(0); i < codeLength; i++ {
		index := rand.Intn(int(charSetLength))
		code += string(charSet[index])
	}

	return code
}

func randomNumber(min, max int32) int32 {
	rand.Seed(time.Now().UnixNano())
	return min + int32(rand.Intn(int(max-min)))
}
