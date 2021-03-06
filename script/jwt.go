package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"math/rand"
)

func GenSymmetricKey(bits int) (k []byte, err error) {
	KeySizeError := "The size of key is invalid!"
	if bits <= 0 || bits%8 != 0 {
		return nil, errors.New(KeySizeError)
	}

	size := bits / 8
	k = make([]byte, size)
	if _, err = rand.Read(k); err != nil {
		return nil, err
	}

	return k, nil
}

func GenJWTString(hmacSampleSecret []byte) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"xusr": "helios",
		"xpwd": "Helios12$",
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)
	//fmt.Println(tokenString, err)
	return tokenString, err
}

func ParseJWTString(tokenString string, hmacSampleSecret []byte) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("xusr", claims["xusr"])
		fmt.Println("xpwd", claims["xpwd"])
	} else {
		fmt.Println(err)
	}
}

func main() {
	test_b := []byte("Here is a string....")
	fmt.Println(test_b)
	key, err := GenSymmetricKey(256)
	if err != nil {
		log.Println("Failed to generate symmetric key")
		panic(err)
		return
	} else {
		fmt.Println(key)
	}

	jwt_token, err := GenJWTString(key)
	if err != nil {
		log.Println("Failed to generate jwt token")
		panic(err)
		return
	} else {
		fmt.Println(jwt_token)
	}
	ParseJWTString(jwt_token, key)
}
