package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"strings"
)

func SignClientReqHeader(req *http.Request, account *UserAccountInfo) error {
	var user, group, uid, gid, umask string

	signedHeaders := []string{}

	if _, exists := req.Header["Content-Type"]; exists {
		signedHeaders = append(signedHeaders, "Content-Type")
	}
	accHeader := make(map[string]interface{})

	user = UID_GLOBAL
	if account.Username != "" {
		user = account.Username
	}

	accHeader["user"] = user

	AesAkey := NewAesEncrypt("BIGWIND_API_KEY_SEED")
	AesSkey := NewAesEncrypt("BIGWIND_SECURITY_KEY_SEED")
	aKey, err := AesAkey.Encrypt(user)
	if err != nil {
		return err
	}

	var sKey string

	if strings.ToUpper(user) == "ROOT" {
		sKey, err = LoadAdminSkey()
	} else {
		sKey, err = AesSkey.Encrypt(user)
	}
	if err != nil {
		return err
	}

	if account.Groupname != "" {
		group = account.Groupname
		accHeader["group"] = group
	}

	if account.Uid != "" {
		uid = account.Uid
		accHeader["uid"] = uid
	}

	if account.Gid != "" {
		gid = account.Gid
		accHeader["gid"] = gid
	}

	if account.Umask != "" {
		umask = account.Umask
		accHeader["umask"] = umask
	}

	js, err := json.Marshal(accHeader)
	if err != nil {
		return err
	}

	// Encrypt the account information
	dKey := sKey
	if len(dKey) < 16 {
		dKey = fmt.Sprintf("%16s", sKey)
	}

	aesEnc := NewAesEncrypt(dKey)
	aHeader, err := aesEnc.Encrypt(string(js))
	if err != nil {
		return err
	}

	signedHeaders = append(signedHeaders, AccountHeader)
	req.Header.Add(AccountHeader, aHeader)

	options := Options{
		SignedHeaders: signedHeaders,
	}

	str, err := StringToSign(req, &options)
	if err != nil {
		return err
	}

	signature := SignString(str, sKey)

	authHeader := fmt.Sprintf("APIKey=%s,Signature=%s", aKey, signature)
	req.Header.Add("Authorization", authHeader)

	return nil
}

///////////////////////////////////////////////////////////////////

const (
	BIGWIND_CLIENT_ADMIN_PASSFILE         = "/etc/.BIGWIND_admin.key"
	BIGWIND_INITIAL_ADMIN_PASSWORD string = "b1gw1nd@x7a0"
	BIGWIND_ADMIN_PASSWORD_SEED    string = "b1gw1nd@xta01sthefuture"
	BIGWIND_ADMIN_PASSWORD_SHADOW  string = "b1gw1nd@xta0Shad0w"
)

func ClientConfDir() string {
	user, err := user.Current()
	if err != nil {
		fmt.Errorf("Fail to get current user: %s\n",
			err.Error())
		return ""
	}
	return user.HomeDir
}

/*
 * encrypt admin passwd by BIGWIND_ADMIN_PASSWORD_SHADOW at local
 * encrypt admin passwd by BIGWIND_ADMIN_PASSWORD_SEED at ETCD
 * admin skey = encrypt admin passwd by BIGWIND_SECURITY_KEY_SEED
 */

func SaveAdminPassword(pass string) error {
	tmpfile, err := ioutil.TempFile(ClientConfDir(), "bigwind_adminpass")
	if err != nil {
		fmt.Println("Fail to create temp password file", err)
		return err
	}
	defer os.Remove(tmpfile.Name())

	aesPass := NewAesEncrypt(BIGWIND_ADMIN_PASSWORD_SHADOW)
	passEnc, err := aesPass.Encrypt(pass)

	if err != nil {
		fmt.Println("Failed to encrypt password!")
		return err
	}

	if _, err := tmpfile.WriteString(passEnc); err != nil {
		fmt.Println("Write temp password error!", err)
		return err
	}

	err = os.Rename(tmpfile.Name(), BIGWIND_CLIENT_ADMIN_PASSFILE)
	if err != nil {
		fmt.Println("Failed to rename to bigwind client admin passfile:%s!\n", err.Error())
		return err
	}

	tmpfile.Close()

	return nil
}

func LoadAdminPassword() (string, error) {
	_, err := os.Stat(BIGWIND_CLIENT_ADMIN_PASSFILE)
	if os.IsNotExist(err) {
		return "", nil
	}

	o, err := ioutil.ReadFile(BIGWIND_CLIENT_ADMIN_PASSFILE)
	if err != nil {
		return "", errors.New("Failed to read password file!")
	}

	aesPass := NewAesEncrypt(BIGWIND_ADMIN_PASSWORD_SHADOW)
	pass, err := aesPass.Decrypt(string(o))
	if err != nil {
		return "", errors.New("Failed to decrypt password.")
	}

	return pass, nil
}

func LoadAdminSkey() (string, error) {

	_, err := os.Stat(BIGWIND_CLIENT_ADMIN_PASSFILE)
	if os.IsNotExist(err) {
		return BIGWIND_INITIAL_ADMIN_PASSWORD, nil
	}

	o, err := ioutil.ReadFile(BIGWIND_CLIENT_ADMIN_PASSFILE)
	if err != nil {
		fmt.Println("Failed to read password file.", err.Error())
		return "", errors.New("Failed to read password file.")
	}

	aesPass := NewAesEncrypt(BIGWIND_ADMIN_PASSWORD_SHADOW)
	pass, err := aesPass.Decrypt(string(o))
	if err != nil {
		fmt.Println("Failed to decrypt shadow passward.", err.Error())
		return "", errors.New("Failed to decrypt shadow password.")
	}

	aesSkey := NewAesEncrypt(BIGWIND_SECURITY_KEY_SEED)
	skey, err := aesSkey.Encrypt(pass)
	if err != nil {
		fmt.Println("Failed to get skey for admin.", err.Error())
		return "", errors.New("Failed to get security key for admin.")
	}

	return skey, err
}
