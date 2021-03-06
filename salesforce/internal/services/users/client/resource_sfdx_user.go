package users

import (
	"fmt"

	"github.com/mocofound/terraform-provider-salesforce/salesforce"
	"github.com/mocofound/terraform-provider-salesforce/salesforce/internal/clients"

	"github.com/hashicorp/terraform/helper/schema"
)

const (
	UserAPIName               = "User"
	UserDefaultProfile        = "00ef2000001fVFK"
	UserDefaultRole           = ""
	UserDefaultTimeZone       = "America/Los_Angeles"
	UserDefaultLocale         = "en_US"
	UserDefaultEmailEncoding  = "ISO-8859-1"
	UserDefaultLanguageLocale = "en_US"
)

var (
	userSObjectFields = []string{
		"IsActive",
		"Alias",
		"ProfileId",
		"UserRoleId",
		"Username",
		"Email",
		"FirstName",
		"LastName",
		"TimeZoneSidKey",
		"LocaleSidKey",
		"EmailEncodingKey",
		"LanguageLocaleKey",
	}
)

type UserSObject struct {
	IsActive          bool   `json:"IsActive,omitempty"`
	Alias             string `json:"Alias,omitempty"`
	ProfileId         string `json:"ProfileId,omitempty"`
	UserRoleId        string `json:"UserRoleId,omitempty"`
	Username          string `json:"Username,omitempty"`
	Email             string `json:"Email,omitempty"`
	FirstName         string `json:"FirstName,omitempty"`
	LastName          string `json:"LastName,omitempty"`
	TimeZoneSidKey    string `json:"TimeZoneSidKey,omitempty"`
	LocaleSidKey      string `json:"LocaleSidKey,omitempty"`
	EmailEncodingKey  string `json:"EmailEncodingKey,omitempty"`
	LanguageLocaleKey string `json:"LanguageLocaleKey,omitempty"`
}

func (uso *UserSObject) ApiName() string {
	return UserAPIName
}

func (uso *UserSObject) ExternalIdApiName() string {
	return UserAPIName
}

func User() *schema.Resource {
	return &schema.Resource{
		Create: userCreate,
		Read:   userRead,
		Update: userUpdate,
		Delete: userDelete,
		Schema: map[string]*schema.Schema{
			"user_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"first_name": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"alias": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_active": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func resourceDataToUser(d *schema.ResourceData) (*UserSObject, error) {
	user := &UserSObject{
		IsActive:          d.Get("is_active").(bool),
		Alias:             d.Get("alias").(string),
		ProfileId:         UserDefaultProfile,
		UserRoleId:        UserDefaultRole,
		Username:          d.Get("username").(string),
		Email:             d.Get("email").(string),
		FirstName:         d.Get("first_name").(string),
		LastName:          d.Get("last_name").(string),
		TimeZoneSidKey:    UserDefaultTimeZone,
		LocaleSidKey:      UserDefaultLocale,
		EmailEncodingKey:  UserDefaultEmailEncoding,
		LanguageLocaleKey: UserDefaultLanguageLocale,
	}

	return user, nil
}

func userToResourceData(user *UserSObject, d *schema.ResourceData) error {
	d.Set("is_active", user.IsActive)
	d.Set("alias", user.Alias)
	d.Set("username", user.Username)
	d.Set("email", user.Email)
	d.Set("first_name", user.FirstName)
	d.Set("last_name", user.LastName)

	return nil
}

func userCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(salesforce.Client).SObjectApi
	if userId, exists := d.GetOk("user_id"); exists && len(userId.(string)) > 0 {
		return nil
	}

	user, err := resourceDataToUser(d)
	if err != nil {
		return err
	}

	if resp, err := client.InsertSObject(user); err != nil {
		return err
	} else {
		d.Set("user_id", resp.Id)
		return userRead(d, m)
	}
}

func userRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*clients.Client).SObjectApi
	userId, exists := d.GetOk("user_id")
	if !exists || len(userId.(string)) == 0 {
		return fmt.Errorf("user_id is unset")
	}

	user := &UserSObject{}
	err := client.GetSObject(userId.(string), userSObjectFields, user)
	if err != nil {
		return err
	}

	err = userToResourceData(user, d)
	if err != nil {
		return err
	}

	d.SetId(userId.(string))
	return nil
}

func userUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(salesforce.Client).SObjectApi
	userId, exists := d.GetOk("user_id")
	if !exists || len(userId.(string)) == 0 {
		return fmt.Errorf("user_id is unset")
	}

	user, err := resourceDataToUser(d)
	if err != nil {
		return err
	}

	err = client.UpdateSObject(userId.(string), user)
	if err != nil {
		return err
	}

	return userRead(d, m)
}

func userDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(salesforce.Client).SObjectApi
	userId, exists := d.GetOk("user_id")
	if !exists || len(userId.(string)) == 0 {
		return fmt.Errorf("user_id is unset")
	}

	user := &UserSObject{}
	return client.DeleteSObject(userId.(string), user)
}
